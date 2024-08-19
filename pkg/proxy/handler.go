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
	"google.golang.org/protobuf/types/known/emptypb"
)

type (
	handler struct {
		director StreamDirector
		stats    *ProxyStats
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

var (
	flowCounter atomic.Uint64
	rpcCounter  atomic.Uint64
)

var (
	service2HostRegistry        = skipmap.New[string, string]()
	method2RequestTypeRegistry  = skipmap.New[string, reflect.Type]()
	method2ResponseTypeRegistry = skipmap.New[string, reflect.Type]()

	rpc2pluginKey = skipmap.New[string, *string]()
	protoPlugins  = skipmap.New[string, *plugin.Plugin]()
)

func newHandler(director StreamDirector) *handler {
	streamer := handler{
		director,
		&ProxyStats{
			Counters: &ProxyCounters{
				ByMethod:  skipmap.New[string, *atomic.Uint64](),
				ByMessage: skipmap.New[string, *atomic.Uint64](),
			},
		},
	}
	return &streamer
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

func newProtosForMethod(
	method *string, includeRequest, includeResponse bool,
) (protoreflect.ProtoMessage, protoreflect.ProtoMessage) {
	var request, response protoreflect.ProtoMessage = nil, nil
	if includeRequest {
		request, _ = newProto(method, method2RequestTypeRegistry)
	}
	if includeResponse {
		response, _ = newProto(method, method2ResponseTypeRegistry)
	}
	return request, response
}

func loadProtoPluginForMethod(method *string) (*string, bool) {
	pluginKey, _ := rpc2pluginKey.
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
	protoPlugin, _ := protoPlugins.LoadOrStoreLazy(
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
				service2HostRegistry,
				method2RequestTypeRegistry,
				method2ResponseTypeRegistry,
			)

			return protoPlugin
		})

	return pluginKey, protoPlugin != nil
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

	serial := flowCounter.Add(1)

	// `ProxyFlow` is the main mechanism for state propagation
	flow := &ProxyFlow{
		Serial: &serial,
		Stats: &ProxyStats{
			Counters: &ProxyCounters{},
		},
		TsProxyReceived: &tsProxyReceived,
		Method:          &fullMethodName,
	}

	method := fullMethodName[1:]
	if serviceKey, loaded := loadProtoPluginForMethod(&method); loaded {
		if endpoint, loaded := service2HostRegistry.Load(*serviceKey); loaded {
			flow.Endpoint = &endpoint
		}
	}

	counterByMethod, _ := s.stats.Counters.ByMethod.
		LoadOrStoreLazy(method, func() *atomic.Uint64 {
			var counterByMethod atomic.Uint64
			return &counterByMethod
		})
	currentCountByMethod := counterByMethod.Add(1)

	// We require that the director's returned context inherits from the serverStream.Context().
	outgoingCtx, backendConn, logger, onStreamEnd, err := s.director(serverStream.Context(), flow)
	if err != nil {
		return err
	}

	clientCtx, clientCancel := context.WithCancel(outgoingCtx)
	defer clientCancel()

	clientStream, err := grpc.NewClientStream(clientCtx, clientStreamDescForProxying, backendConn, fullMethodName)
	if err != nil {
		return err
	}

	protoFactory := func(includeRequest, includeResponse bool) (
		protoreflect.ProtoMessage, protoreflect.ProtoMessage,
	) {
		return newProtosForMethod(&method, includeRequest, includeResponse)
	}

	var requestCounter, responseCounter atomic.Uint64
	var wg sync.WaitGroup

	rpcLogger := func(
		request, response *protoreflect.ProtoMessage,
		status *spb.Status, start, end *time.Time,
	) {
		wg.Add(1)
		currentCountByRPC := rpcCounter.Add(1)

		// do NOT block streams
		// [ToDo]: use go-routines pool
		go func(currentCountByMethod, currentCountByRPC *uint64) {
			rpc := &RPC{
				Endpoint:    flow.Endpoint,
				Method:      flow.Method,
				TsStart:     start,
				TsEnd:       end,
				StatusProto: status,
			}

			var protoMessage protoreflect.ProtoMessage

			requestCount := requestCounter.Load()
			responseCount := responseCounter.Load()

			if request != nil && response == nil {
				requestCount = requestCounter.Add(1)
				protoMessage = *request
				rpc.MessageProto = request
				rpc.IsRequest = true
				rpc.IsResponse = false
			} else if response != nil && request == nil {
				responseCount = responseCounter.Add(1)
				protoMessage = *response
				rpc.MessageProto = response
				rpc.IsRequest = false
				rpc.IsResponse = true
			} else {
				rpc.IsRequest = false
				rpc.IsResponse = false
			}

			if rpc.IsRequest || rpc.IsResponse {
				// `protoreflect.FullName` is a `string`
				// see: https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect#ProtoMessage
				messageFullName := protoMessage.ProtoReflect().Type().Descriptor().FullName()
				rpc.MessageFullName = &messageFullName
				counterByMessage, _ := s.stats.Counters.ByMessage.
					LoadOrStoreLazy(string(messageFullName),
						func() *atomic.Uint64 {
							var counterByMethod atomic.Uint64
							return &counterByMethod
						})
				counterByMessage.Add(1)
			}

			rpc.Stats = &RPCStats{
				Counters: &RPCCounters{
					ForMethod: currentCountByMethod,
					ForRPC:    currentCountByRPC,
					Requests:  &requestCount,
					Responses: &responseCount,
				},
			}

			go logger(clientCtx, clientStream.Context(), flow, rpc)

			wg.Done()
		}(&currentCountByMethod, &currentCountByRPC)
	}

	tsStreamStart := time.Now()

	// Explicitly *do not close* s2cErrChan and c2sErrChan, otherwise the select below will not terminate.
	// Channels do not have to be closed, it is just a control flow mechanism, see
	// https://groups.google.com/forum/#!msg/golang-nuts/pZwdYRGxCIk/qpbHxRRPJdUJ
	s2cErrChan := s.forwardServerToClient(serverStream, clientStream, protoFactory, rpcLogger)
	c2sErrChan := s.forwardClientToServer(clientStream, serverStream, protoFactory, rpcLogger)

	defer func() {
		tsStreamEnd := time.Now()
		// do NOT block
		go func() {
			wg.Wait()
			flow.TsStreamStart = &tsStreamStart
			flow.TsStreamEnd = &tsStreamEnd
			_ = rpcCounter.Load()
			_ = requestCounter.Load()
			_ = responseCounter.Load()
			onStreamEnd(clientCtx, clientStream.Context(), flow)
		}()
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

			var statusProto *spb.Status
			start := time.Now()
			// receiving `response` from the actual `server`
			if err := src.RecvMsg(protoResponse); err != nil {
				end := time.Now()
				ret <- err // this can be io.EOF which is happy case
				if rpcStatus, ok := status.FromError(err); ok {
					statusProto = rpcStatus.Proto()
				}
				logger(nil, nil, statusProto, &start, &end)
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

			var statusProto *spb.Status = nil
			start := time.Now()
			// forwarding `request` to the actual `server`
			if err := dst.SendMsg(protoRequest); err != nil {
				end := time.Now()
				ret <- err
				if rpcStatus, ok := status.FromError(err); ok {
					statusProto = rpcStatus.Proto()
				}
				logger(nil, nil, statusProto, &start, &end)
				break
			}
			end := time.Now()
			logger(&protoRequest, nil, nil, &start, &end)
		}
	}()
	return ret
}
