package proxy

import (
	"context"
	"plugin"
	"reflect"
	"sync/atomic"
	"time"

	"github.com/zhangyunhao116/skipmap"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type (
	ProxyLogger = func(serverCtx, clientCtx context.Context, flow *ProxyFlow, rpc *RPC)

	OnStreamStartHamdler = func(serverCtx, clientCtx context.Context, flow *ProxyFlow)
	OnStreamEndHandler   = func(serverCtx, clientCtx context.Context, flow *ProxyFlow, stats *RPCStats)

	ProxyRegistry struct {
		hosts         *skipmap.OrderedMap[string, string]
		requestTypes  *skipmap.OrderedMap[string, reflect.Type]
		responseTypes *skipmap.OrderedMap[string, reflect.Type]
		pluginsKeys   *skipmap.OrderedMap[string, *string]
		plugins       *skipmap.OrderedMap[string, *plugin.Plugin]
	}

	ProxyPlugins struct {
		keys *skipmap.OrderedMap[string, *string]
	}

	ProxyCounters struct {
		Flows     *atomic.Uint64
		Messages  *atomic.Uint64
		ByTarget  *skipmap.OrderedMap[string, *atomic.Uint64]
		ByMethod  *skipmap.OrderedMap[string, *atomic.Uint64]
		ByMessage *skipmap.OrderedMap[string, *atomic.Uint64]
	}

	ProxyStats struct {
		Counters *ProxyCounters
	}

	ProxyFlowTimestamps struct {
		ProxyReceived       *time.Time
		StreamCreationStart *time.Time
		StreamCreationEnd   *time.Time
		DirectorStart       *time.Time
		DirectorEnd         *time.Time
		Oauth2Start         *time.Time
		Oauth2End           *time.Time
		StreamStart         *time.Time
		StreamEnd           *time.Time
	}

	ProxyFlowHandlers struct {
		OnStreamStart OnStreamStartHamdler
		OnStreamEnd   OnStreamEndHandler
	}

	ProxyFlow struct {
		ClientConn         *grpc.ClientConn
		Serial             *uint64
		Count              *uint32
		ProjectID          *string
		Target, Method     *string
		XCloudTraceContext *string
		Client, Server     *peer.Peer
		Handlers           *ProxyFlowHandlers
		Timestamps         *ProxyFlowTimestamps
		Stats              *ProxyStats
	}

	RPCTimestamps struct {
		Start *time.Time
		End   *time.Time
	}

	RPC struct {
		Target          *string
		Method          *string
		Timestamps      *RPCTimestamps
		MessageProto    *protoreflect.ProtoMessage
		MessageFullName *protoreflect.FullName
		StatusProto     *spb.Status
		IsRequest       bool
		IsResponse      bool
		Stats           *RPCStats
	}

	Counters struct {
		Flows     *uint64
		Messages  *uint64
		ForTarget *uint64
		ForMethod *uint64
	}

	RPCCounters struct {
		Overall   *Counters
		Messages  *uint64
		Requests  *uint64
		Responses *uint64
		Errors    *uint64
	}

	RPCStats struct {
		Counters *RPCCounters
	}

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
	StreamDirector = func(ctx context.Context, flow *ProxyFlow) (context.Context, *grpc.ClientConn, ProxyLogger, *ProxyFlowHandlers, error)
)
