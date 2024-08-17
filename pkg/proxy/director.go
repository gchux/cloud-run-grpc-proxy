package proxy

import (
	"context"
	"time"

	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type (
	ProxyCounters struct {
		ByMethod *uint64
	}
	ProxyStats struct {
		Counters *ProxyCounters
	}
	ProxyFlow struct {
		ClientConn             *grpc.ClientConn
		Serial                 *uint64
		Stats                  *ProxyStats
		ProjectID              *string
		Endpoint, Method       *string
		XCloudTraceContext     *string
		ProtoRequest           protoreflect.ProtoMessage
		ProtoResponse          protoreflect.ProtoMessage
		StatusProto            *spb.Status
		Client, Server         *peer.Peer
		TsProxyReceived        *time.Time
		TsBeforeStreamCreation *time.Time
		TsAfterStreamCreation  *time.Time
		TsDirectorStart        *time.Time
		TsDirectorEnd          *time.Time
		TsOauth2Start          *time.Time
		TsOauth2End            *time.Time
		TsStreamStart          *time.Time
		TsStreamEnd            *time.Time
	}

	Logger func(serverCtx, clientCtx context.Context, flow *ProxyFlow, request, response *protoreflect.ProtoMessage, start, end *time.Time, isStreamEnd bool)

	// StreamDirector returns a gRPC ClientConn to be used to forward the call to.
	//
	// The presence of the `Context` allows for rich filtering, e.g. based on Metadata (headers).
	// If no handling is meant to be done, a `codes.NotImplemented` gRPC error should be returned.
	//
	// The context returned from this function should be the context for the *outgoing* (to backend) call. In case you want
	// to forward any Metadata between the inbound request and outbound requests, you should do it manually. However, you
	// *must* propagate the cancel function (`context.WithCancel`) of the inbound context to the one returned.
	//
	// It is worth noting that the StreamDirector will be fired *after* all server-side stream interceptors
	// are invoked. So decisions around authorization, monitoring etc. are better to be handled there.
	StreamDirector func(ctx context.Context, flow *ProxyFlow) (context.Context, *grpc.ClientConn, Logger, error)
)
