package proxy

import (
	"context"
	"fmt"
	"io"
	"reflect"
	"sync/atomic"
	"time"

	"github.com/zhangyunhao116/skipmap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"
)

var clientStreamDescForProxying = &grpc.StreamDesc{
	ServerStreams: true,
	ClientStreams: true,
}
var counter atomic.Uint64

// RegisterService sets up a proxy handler for a particular gRPC service and method.
// The behaviour is the same as if you were registering a handler method, e.g. from a generated pb.go file.
func RegisterService(server *grpc.Server, director StreamDirector, serviceName string, methodNames ...string) {
	streamer := &handler{director}
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
	streamer := &handler{director: director}
	return streamer.handler
}

type handler struct {
	director StreamDirector
}

func newProto(rpc *string, regsitry *skipmap.OrderedMap[string, reflect.Type]) (protoreflect.ProtoMessage, error) {
	protoType, loaded := regsitry.Load(*rpc)
	if !loaded {
		return &emptypb.Empty{},
			fmt.Errorf("no proto not foound for: %s", *rpc)
	}

	if proto, ok := reflect.New(protoType).Interface().(protoreflect.ProtoMessage); ok {
		return proto, nil
	}
	return &emptypb.Empty{},
		fmt.Errorf("proto is not 'protoreflect.ProtoMessage': %s", protoType.Name())
}

func getProtosForRPC(rpc *string) (protoreflect.ProtoMessage, protoreflect.ProtoMessage) {
	requestProto, _ := newProto(rpc, method2RequestType)
	responseProto, _ := newProto(rpc, method2ResponseType)
	return requestProto, responseProto
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

	serial := counter.Add(1)

	// `ProxyFlow` is the main mechanism for state propagation
	flow := &ProxyFlow{
		Serial: &serial,
		Stats: &ProxyStats{
			Counters: &ProxyCounters{},
		},
		TsProxyReceived: &tsProxyReceived,
		Method:          &fullMethodName,
	}

	// We require that the director's returned context inherits from the serverStream.Context().
	outgoingCtx, backendConn, onStreamEnd, err := s.director(serverStream.Context(), flow)
	if err != nil {
		return err
	}

	clientCtx, clientCancel := context.WithCancel(outgoingCtx)
	defer clientCancel()

	clientStream, err := grpc.NewClientStream(clientCtx, clientStreamDescForProxying, backendConn, fullMethodName)
	if err != nil {
		return err
	}

	// [ToDo]: split stream timestamps into: server{start|end} and client{start|end}
	tsStreamStart := time.Now()
	defer func() {
		tsStreamEnd := time.Now()
		flow.TsStreamStart = &tsStreamStart
		flow.TsStreamEnd = &tsStreamEnd
		go onStreamEnd(clientCtx, clientStream.Context(), flow)
	}()

	rpcKey := fullMethodName[1:]
	flow.ProtoRequest, flow.ProtoResponse = getProtosForRPC(&rpcKey)

	// Explicitly *do not close* s2cErrChan and c2sErrChan, otherwise the select below will not terminate.
	// Channels do not have to be closed, it is just a control flow mechanism, see
	// https://groups.google.com/forum/#!msg/golang-nuts/pZwdYRGxCIk/qpbHxRRPJdUJ
	s2cErrChan := s.forwardServerToClient(serverStream, clientStream, flow)
	c2sErrChan := s.forwardClientToServer(clientStream, serverStream, flow)

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
			if rpcStatus, ok := status.FromError(c2sErr); ok {
				flow.StatusProto = rpcStatus.Proto()
			}
			// c2sErr will contain RPC error from client code. If not io.EOF return the RPC error as server stream error.
			if c2sErr != io.EOF {
				return c2sErr
			}
			return nil
		}
	}

	return status.Errorf(codes.Internal, "gRPC proxying should never reach this stage.")
}

func (s *handler) forwardClientToServer(src grpc.ClientStream, dst grpc.ServerStream, flow *ProxyFlow) chan error {
	ret := make(chan error, 1)
	go func() {
		// f := &emptypb.Empty{}
		for i := 0; ; i++ {
			if err := src.RecvMsg(flow.ProtoResponse); err != nil {
				ret <- err // this can be io.EOF which is happy case
				break
			}
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

			if err := dst.SendMsg(flow.ProtoResponse); err != nil {
				ret <- err
				break
			}
		}
	}()
	return ret
}

func (s *handler) forwardServerToClient(src grpc.ServerStream, dst grpc.ClientStream, flow *ProxyFlow) chan error {
	ret := make(chan error, 1)
	go func() {
		// f := &emptypb.Empty{}
		for i := 0; ; i++ {
			if err := src.RecvMsg(flow.ProtoRequest); err != nil {
				ret <- err // this can be io.EOF which is happy case
				break
			}

			if err := dst.SendMsg(flow.ProtoRequest); err != nil {
				ret <- err
				break
			}
		}
	}()
	return ret
}
