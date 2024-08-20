package proxy

import (
	"context"
	"fmt"
	"io"
	"os"
	"plugin"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zhangyunhao116/skipmap"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type (
	handler struct {
		director StreamDirector
		stats    *ProxyStats
		registry *ProxyRegistry
	}

	RPCLogger    = func(request, response *protoreflect.ProtoMessage, status *spb.Status, start, end *time.Time)
	ProtoFactory = func(bool, bool) (protoreflect.ProtoMessage, protoreflect.ProtoMessage)
	PluginLoader = func(
		*skipmap.OrderedMap[string, string],
		*skipmap.OrderedMap[string, reflect.Type],
		*skipmap.OrderedMap[string, reflect.Type],
	)
)

var clientStreamDescForProxying = &grpc.StreamDesc{
	ServerStreams: true,
	ClientStreams: true,
}

func newHandler(director StreamDirector) *handler {
	streamer := &handler{
		director: director,
		registry: &ProxyRegistry{
			hosts:         skipmap.New[string, string](),
			requestTypes:  skipmap.New[string, reflect.Type](),
			responseTypes: skipmap.New[string, reflect.Type](),
			pluginsKeys:   skipmap.New[string, *string](),
			plugins:       skipmap.New[string, *plugin.Plugin](),
		},
		stats: &ProxyStats{
			Counters: &ProxyCounters{
				Flows:     new(atomic.Uint64),
				Messages:  new(atomic.Uint64),
				ByTarget:  skipmap.New[string, *atomic.Uint64](),
				ByMethod:  skipmap.New[string, *atomic.Uint64](),
				ByMessage: skipmap.New[string, *atomic.Uint64](),
			},
		},
	}
	return streamer
}

// RegisterService sets up a proxy handler for a particular gRPC service and method.
// The behaviour is the same as if you were registering a handler method, e.g. from a generated pb.go file.
func RegisterService(server *grpc.Server, director StreamDirector, serviceName string, methodNames ...string) {
	streamer := newHandler(director)

	fakeDesc := &grpc.ServiceDesc{
		ServiceName: serviceName,
		HandlerType: (*interface{})(nil),
	}

	for _, m := range methodNames {
		streamDesc := grpc.StreamDesc{
			StreamName:    m,
			Handler:       streamer.handler,
			ServerStreams: true,
			ClientStreams: true,
		}
		fakeDesc.Streams = append(fakeDesc.Streams, streamDesc)
	}

	server.RegisterService(fakeDesc, streamer)
}

// TransparentHandler returns a handler that attempts to proxy all requests that are not registered in the server.
// The indented use here is as a transparent proxy, where the server doesn't know about the services implemented by the
// backends. It should be used as a `grpc.UnknownServiceHandler`.
func TransparentHandler(director StreamDirector) grpc.StreamHandler {
	streamer := newHandler(director)
	return streamer.handler
}

func newProto(
	method *string, regsitry *skipmap.OrderedMap[string, reflect.Type],
) (protoreflect.ProtoMessage, error) {
	protoType, loaded := regsitry.Load(*method)
	if !loaded {
		return &emptypb.Empty{},
			fmt.Errorf("no proto not foound for: %s", *method)
	}

	if proto, ok := reflect.New(protoType).
		Interface().(protoreflect.ProtoMessage); ok {
		return proto, nil
	}
	return &emptypb.Empty{},
		fmt.Errorf("proto is not 'protoreflect.ProtoMessage': %s", protoType.Name())
}

func (s *handler) newProtosForMethod(
	method *string, includeRequest, includeResponse bool,
) (protoreflect.ProtoMessage, protoreflect.ProtoMessage) {
	var request, response protoreflect.ProtoMessage = nil, nil
	var requestErr, responseErr error = nil, nil
	if includeRequest {
		if request, requestErr = newProto(method, s.registry.requestTypes); requestErr != nil {
			io.WriteString(os.Stderr, requestErr.Error()+"\n")
		}
	}
	if includeResponse {
		if response, responseErr = newProto(method, s.registry.responseTypes); responseErr != nil {
			io.WriteString(os.Stderr, responseErr.Error()+"\n")
		}
	}
	return request, response
}

func (s *handler) loadProtoPluginForMethod(method *string) (*string, bool) {
	pluginKey, _ := s.registry.pluginsKeys.
		LoadOrStoreLazy(*method, func() *string {
			rpcParts := strings.Split(*method, ".")
			// `pluginKey` is everything except the method:
			//   - a Proto package contains all the services ( thus: all the `rcp`s/`method`s )
			key := strings.Join(rpcParts[:len(rpcParts)-1], ".")
			return &key
		})

	// loading `Proto Plugin` from cache only for the side-effects...
	// the error is not important as `newProto` returns `emptypb.Empty` if the actual type is not available;
	// all that's been lost is the capacity to show the actual proto2json translations.
	protoPlugin, _ := s.registry.plugins.LoadOrStoreLazy(
		*pluginKey, func() *plugin.Plugin {
			protoPlugin, err := plugin.Open(*pluginKey + ".so")
			if err != nil {
				io.WriteString(os.Stderr, err.Error()+"\n")
				return nil
			}

			LoadFn, err := protoPlugin.Lookup("Load")
			if err != nil {
				io.WriteString(os.Stderr, err.Error()+"\n")
				return nil
			}

			// a `Plugin` registers all the types for request/response
			// for all the `rpc`s/`method`s available in a PRoto package.
			LoadFn.(PluginLoader)(
				s.registry.hosts,
				s.registry.requestTypes,
				s.registry.responseTypes,
			)

			return protoPlugin
		})

	return pluginKey, protoPlugin != nil
}

func (s *handler) getCounter(name *string,
	holder *skipmap.OrderedMap[string, *atomic.Uint64],
) *atomic.Uint64 {
	counter, _ := holder.
		LoadOrStoreLazy(*name,
			func() *atomic.Uint64 {
				var counter atomic.Uint64
				return &counter
			})
	return counter
}

// handler is where the real magic of proxying happens.
// It is invoked like any gRPC server stream and uses the emptypb.Empty type server
// to proxy calls between the input and output streams.
func (s *handler) handler(srv interface{}, serverStream grpc.ServerStream) error {
	tsProxyReceived := time.Now()

	// little bit of gRPC internals never hurt anyone
	fullMethodName, ok := grpc.MethodFromServerStream(serverStream)
	if !ok {
		return status.Errorf(codes.Internal, "lowLevelServerStream not exists in context")
	}

	serial := s.stats.Counters.Flows.Add(1)

	// `ProxyFlow` is the main mechanism for state propagation at the STREAM level
	flow := &ProxyFlow{
		Serial: &serial,
		Stats: &ProxyStats{
			Counters: &ProxyCounters{},
		},
		Timestamps: &ProxyFlowTimestamps{
			ProxyReceived: &tsProxyReceived,
		},
		Method: &fullMethodName,
	}

	method := fullMethodName[1:]
	if serviceKey, loaded := s.loadProtoPluginForMethod(&method); loaded {
		if host, loaded := s.registry.hosts.Load(*serviceKey); loaded {
			flow.Target = &host
		}
	}

	// We require that the director's returned context inherits from the serverStream.Context().
	outgoingCtx, backendConn, logger, handlers, err := s.director(serverStream.Context(), flow)
	if err != nil {
		return err
	}

	clientCtx, clientCancel := context.WithCancel(outgoingCtx)
	defer clientCancel()

	clientStream, err := grpc.NewClientStream(clientCtx, clientStreamDescForProxying, backendConn, fullMethodName)
	if err != nil {
		return err
	}

	// Proxy is now commited to handle the RPC
	s.getCounter(flow.Target, s.stats.Counters.ByTarget).Add(1)
	s.getCounter(&method, s.stats.Counters.ByMethod).Add(1)

	protoFactory := func(includeRequest, includeResponse bool) (
		protoreflect.ProtoMessage, protoreflect.ProtoMessage,
	) {
		return s.newProtosForMethod(&method, includeRequest, includeResponse)
	}

	var requestCounter, responseCounter, errorCounter, messageCounter atomic.Uint64
	var wg sync.WaitGroup

	rpcLogger := func(
		request, response *protoreflect.ProtoMessage,
		status *spb.Status, start, end *time.Time,
	) {
		wg.Add(1)

		s.stats.Counters.Messages.Add(1)

		// do NOT block streams
		// [ToDo]: use go-routines pool
		go func(messageCounter, requestCounter, responseCounter, errorCounter *atomic.Uint64) {
			if request == nil && response == nil && status == nil {
				return // no RPC error, must be EOF with status 0
			}

			messagesCount := messageCounter.Add(1)
			requestsCount := requestCounter.Load()
			responsesCount := responseCounter.Load()
			errorsCount := errorCounter.Load()

			rpc := &RPC{
				Target: flow.Target,
				Method: flow.Method,
				Timestamps: &RPCTimestamps{
					Start: start,
					End:   end,
				},
				IsRequest:  false,
				IsResponse: false,
				Stats: &RPCStats{
					Counters: &RPCCounters{
						Messages:  &messagesCount,
						Requests:  &requestsCount,
						Responses: &responsesCount,
						Errors:    &errorsCount,
					},
				},
			}

			var protoMessage protoreflect.ProtoMessage

			if request != nil && response == nil {
				requestsCount = requestCounter.Add(1)
				protoMessage = *request
				rpc.MessageProto = request
				rpc.IsRequest = true
				rpc.IsResponse = false
				rpc.Stats.Counters.Requests = &requestsCount
			} else if response != nil && request == nil {
				responsesCount = responseCounter.Add(1)
				protoMessage = *response
				rpc.MessageProto = response
				rpc.IsRequest = false
				rpc.IsResponse = true
				rpc.Stats.Counters.Responses = &responsesCount
			}

			if status != nil {
				errorsCount = errorCounter.Add(1)
				rpc.StatusProto = status
				rpc.Stats.Counters.Errors = &errorsCount
			}

			if rpc.IsRequest || rpc.IsResponse {
				// `protoreflect.FullName` is a `string`
				// see: https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect#ProtoMessage
				any, _ := anypb.New(protoMessage)
				// messageFullName := protoMessage.ProtoReflect().Type().Descriptor().FullName()
				messageFullName := any.MessageName()
				rpc.MessageFullName = &messageFullName
				messageFullNameStr := string(messageFullName)
				s.getCounter(&messageFullNameStr, s.stats.Counters.ByMessage).Add(1)
			}

			go logger(clientCtx, clientStream.Context(), flow, rpc)

			wg.Done()
		}(&messageCounter, &requestCounter, &responseCounter, &errorCounter)
	}

	tsStreamStart := time.Now()
	flow.Timestamps.StreamStart = &tsStreamStart
	// Explicitly *do not close* s2cErrChan and c2sErrChan, otherwise the select below will not terminate.
	// Channels do not have to be closed, it is just a control flow mechanism, see
	// https://groups.google.com/forum/#!msg/golang-nuts/pZwdYRGxCIk/qpbHxRRPJdUJ
	s2cErrChan := s.forwardServerToClient(serverStream, clientStream, protoFactory, rpcLogger)
	c2sErrChan := s.forwardClientToServer(clientStream, serverStream, protoFactory, rpcLogger)

	flow.Stats = s.stats // live view of `ProxyStats`

	go handlers.OnStreamStart(clientCtx, clientStream.Context(), flow)
	defer func() {
		tsStreamEnd := time.Now()
		// do NOT block
		go func(messageCounter, requestCounter, responseCounter, errorCounter *atomic.Uint64) {
			wg.Wait()

			flow.Timestamps.StreamEnd = &tsStreamEnd

			// capture global proxy stats on stream end
			overallFlowsCount := s.stats.Counters.Flows.Load()
			overallMessagesCount := s.stats.Counters.Messages.Load()
			overallCountForTarget := s.getCounter(flow.Target, s.stats.Counters.ByTarget).Load()
			overallCountForMethod := s.getCounter(&method, s.stats.Counters.ByMethod).Load()

			messagesCount := messageCounter.Load()
			requestsCount := requestCounter.Load()
			responsesCount := responseCounter.Load()
			errorsCount := errorCounter.Load()

			stats := &RPCStats{
				Counters: &RPCCounters{
					Overall: &Counters{ // immutable current view of overall Proxy state
						Flows:     &overallFlowsCount,
						Messages:  &overallMessagesCount,
						ForTarget: &overallCountForTarget,
						ForMethod: &overallCountForMethod,
					},
					Messages:  &messagesCount,
					Requests:  &requestsCount,
					Responses: &responsesCount,
					Errors:    &errorsCount,
				},
			}
			handlers.OnStreamEnd(clientCtx, clientStream.Context(), flow, stats)
		}(&messageCounter, &requestCounter, &responseCounter, &errorCounter)
	}()

	// We don't know which side is going to stop sending first, so we need a select between the two.
	for i := 0; i < 2; i++ {
		select {
		case s2cErr := <-s2cErrChan:
			if s2cErr == io.EOF {
				// this is the happy case where the sender has encountered io.EOF, and won't be sending anymore./
				// the clientStream>serverStream may continue pumping though.
				clientStream.CloseSend()
			} else {
				// however, we may have gotten a receive error (stream disconnected, a read error etc) in which case we need
				// to cancel the clientStream to the backend, let all of its goroutines be freed up by the CancelFunc and
				// exit with an error to the stack
				clientCancel()
				return status.Errorf(codes.Internal, "failed proxying s2c: %v", s2cErr)
			}
		case c2sErr := <-c2sErrChan:
			// This happens when the clientStream has nothing else to offer (io.EOF), returned a gRPC error. In those two
			// cases we may have received Trailers as part of the call. In case of other errors (stream closed) the trailers
			// will be nil.
			md := clientStream.Trailer()
			serverStream.SetTrailer(md)
			// c2sErr will contain RPC error from client code. If not io.EOF return the RPC error as server stream error.
			if c2sErr != io.EOF {
				return c2sErr
			}
			return nil
		}
	}

	return status.Errorf(codes.Internal, "gRPC proxying should never reach this stage.")
}

func (s *handler) forwardClientToServer(src grpc.ClientStream, dst grpc.ServerStream, factory ProtoFactory, logger RPCLogger) chan error {
	ret := make(chan error, 1)
	go func() {
		// f := &emptypb.Empty{}
		for i := uint32(0); ; i++ {

			_, protoResponse := factory(false /* includeRequest */, true /* includeResponse */)

			start := time.Now()
			// receiving `response` from the actual `server`
			if err := src.RecvMsg(protoResponse); err != nil {
				end := time.Now()
				ret <- err // this can be io.EOF which is happy case
				if rpcStatus, ok := status.FromError(err); ok {
					logger(nil, nil, rpcStatus.Proto(), &start, &end)
				}
				break
			}
			end := time.Now()

			if i == 0 {
				// This is a bit of a hack, but client to server headers are only readable after first client msg is
				// received but must be written to server stream before the first msg is flushed.
				// This is the only place to do it nicely.
				md, err := src.Header()
				if err != nil {
					ret <- err
					break
				}
				if err := dst.SendHeader(md); err != nil {
					ret <- err
					break
				}
			}

			// forwarding `response` to the actual `client`
			if err := dst.SendMsg(protoResponse); err != nil {
				ret <- err
				break
			}

			logger(nil, &protoResponse, nil, &start, &end)
		}
	}()
	return ret
}

func (s *handler) forwardServerToClient(src grpc.ServerStream, dst grpc.ClientStream, factory ProtoFactory, logger RPCLogger) chan error {
	ret := make(chan error, 1)
	go func() {
		// f := &emptypb.Empty{}
		for i := uint32(0); ; i++ {

			protoRequest, _ := factory(true /* includeRequest */, false /* includeResponse */)

			// receiving `request` from the actual `client`
			if err := src.RecvMsg(protoRequest); err != nil {
				ret <- err // this can be io.EOF which is happy case
				break
			}

			start := time.Now()
			// forwarding `request` to the actual `server`
			if err := dst.SendMsg(protoRequest); err != nil {
				end := time.Now()
				ret <- err
				if rpcStatus, ok := status.FromError(err); ok {
					logger(nil, &protoRequest, rpcStatus.Proto(), &start, &end)
				}
				break
			}
			end := time.Now()
			logger(&protoRequest, nil, nil, &start, &end)
		}
	}()
	return ret
}
