package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"net/netip"
	"os"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/Jeffail/gabs/v2"
	"github.com/gchux/grpc-proxy/pkg/proxy"
	"github.com/segmentio/fasthash/fnv1a"
	"github.com/wissance/stringFormatter"
	"github.com/zhangyunhao116/skipmap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/encoding/protojson"

	"golang.org/x/oauth2"
	auth "golang.org/x/oauth2/google"
	"golang.org/x/sys/unix"
)

type (
	connectionFactoryProvider = func() connectionFactory
	connectionFactory         = func() *grpc.ClientConn
)

const (
	defaultProxyPort = 51051
)

var (
	projectID  = os.Getenv("PROJECT_ID")
	projectNum = os.Getenv("PROJECT_NUM")
	kService   = os.Getenv("K_SERVICE")
	kRevision  = os.Getenv("K_REVISION")
	proxyPort  = os.Getenv("GRPC_PROXY_PORT")
	targetHost = os.Getenv("GRPC_PROXY_TARGET_HOST")
	targetPort = os.Getenv("GRPC_PROXY_TARGET_PORT")
)

var (
	xForwardedFor = stringFormatter.Format("/apis/serving.knative.dev/v1/namespaces/{0}/revisions/{1}", projectNum, kRevision)
	logsProducer  = stringFormatter.Format("{0}/{1}", xForwardedFor, "sidecar/grpc-proxy")
	target        = stringFormatter.Format("{0}:{1}", targetHost, targetPort)
	targerAddr, _ = netip.ParseAddrPort(target)
)

var scopes = []string{
	"https://www.googleapis.com/auth/cloud-platform",
}

var (
	// https://pkg.go.dev/syscall#Linger
	_SO_LINGER = &syscall.Linger{Onoff: 1, Linger: 0} // skip TIME_WAIT

	dialer = &net.Dialer{
		DualStack: true,
		Control: func(network, address string, conn syscall.RawConn) error {
			return connectionControl(network, address, conn)
		},
	}

	tlsCreds = credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})

	keepAliveEnforcementPolicy = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: true,            // Allow pings even when there are no active streams
	}

	keepAliveServerParams = keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
		MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
		MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
		Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
	}

	ipToIfaceMap map[string]net.Interface

	connectionFactories *skipmap.OrderedMap[string, connectionFactory]
	connections         *skipmap.OrderedMap[string, *grpc.ClientConn]
	flows               *skipmap.OrderedMap[uint64, *proxy.ProxyFlow]

	defaultClientConnFactory         = newConnectionFactory(&target)
	defaultConnectionFactoryProvider = func() connectionFactory { return defaultClientConnFactory }

	xCloudTraceContextRegexp = regexp.MustCompile(`^(?P<trace>.+?)/(?P<span>.+?)(?:;o=.*)?$`)
)

func setSocketOptions(fd uintptr, err *error) {
	// https://golang.google.cn/src/internal/poll/sockopt.go
	// https://pkg.go.dev/syscall#SetsockoptLinger
	lingerErr := syscall.SetsockoptLinger(int(fd), unix.SOL_SOCKET, unix.SO_LINGER, _SO_LINGER)
	*err = errors.Join(lingerErr)
}

func connectionControl(network, address string, conn syscall.RawConn) error {
	var operr error
	if err := conn.Control(func(fd uintptr) {
		setSocketOptions(fd, &operr)
	}); err != nil {
		return err
	}
	return operr
}

func getToken() (*oauth2.Token, error) {
	ctx := context.Background()
	credentials, err := auth.FindDefaultCredentials(ctx, scopes...)
	credentials.ProjectID = projectID
	if err == nil {
		return credentials.TokenSource.Token()
	}
	return nil, err
}

func hashL3(srcIP, dstIP *netip.Addr) uint64 {
	l3Hash := uint64(4) // IPv4 == 4
	if srcIP.Is6() {
		src := srcIP.As16()
		dst := dstIP.As16()
		return (uint64(41 /* IPv6 == 41 */) + fnv1a.HashBytes64(src[:]) + fnv1a.HashBytes64(dst[:]))
	} else {
		src := srcIP.As4()
		dst := dstIP.As4()
		l3Hash += (fnv1a.HashBytes64(src[:]) + fnv1a.HashBytes64(dst[:]))
	}
	return fnv1a.HashUint64(l3Hash)
}

func hashL4(p *peer.Peer, srcPort, dstPort *uint64) uint64 {
	proto := uint64(6) // TCP(6)
	if strings.HasPrefix(p.LocalAddr.Network(), "udp") {
		proto += uint64(11) // UDP(17)
	}
	return fnv1a.HashUint64(proto + *srcPort + *dstPort)
}

func hash(p *peer.Peer,
	srcIP, dstIP *netip.Addr,
	srcPort, dstPort, ifaceIndex *uint64,
) uint64 {
	hash := fnv1a.AddUint64(fnv1a.Init64, *ifaceIndex)
	hash = fnv1a.AddUint64(hash, hashL3(srcIP, dstIP))
	return fnv1a.AddUint64(hash, hashL4(p, srcPort, dstPort))
}

func getProxyFlowFromMetadata(md metadata.MD) (*proxy.ProxyFlow, error) {
	grpcProxyID := md.Get("x-grpc-proxy-id")

	if len(grpcProxyID) == 0 {
		return nil, errors.New("RPC flow not found")
	}

	rpc, err := strconv.ParseUint(grpcProxyID[0], 10, 64)
	if err != nil {
		return nil, err
	}

	flow, ok := flows.Load(rpc)
	if !ok {
		return nil, errors.New("RPC flow not found: " + grpcProxyID[0])
	}

	return flow, nil
}

func newConnectionFactory(target *string) connectionFactory {
	return func() *grpc.ClientConn {
		connection, _ := connections.
			LoadOrStoreLazy(*target, func() *grpc.ClientConn {
				// [ToDo]: validate `target`
				cc, err := grpc.NewClient(*target,
					grpc.WithTransportCredentials(tlsCreds),
					grpc.WithContextDialer(contextDialer),
					grpc.WithUserAgent("grpc-proxy/1.0.0"),
					grpc.WithStreamInterceptor(streamClientInterceptor),
					grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`))
				if err != nil {
					io.WriteString(os.Stderr, err.Error()+"\n")
					return nil
				}
				return cc
			})
		return connection
	}
}

func newJSONLog(
	serverCtx, clientCtx context.Context,
	flow *proxy.ProxyFlow, ts *time.Time,
) (*gabs.Container, *string) {
	serial := *flow.Serial
	projectID := *flow.ProjectID
	target := *flow.Target
	method := *flow.Method

	mdIn, _ := metadata.FromIncomingContext(serverCtx)
	mdOut, _ := metadata.FromOutgoingContext(clientCtx)

	mdIn.Delete("authorization")
	mdOut.Delete("authorization")

	authority := mdIn.Get(":authority")[0]
	authorityParts := strings.Split(authority, ":")

	// connection from client to proxy
	server := flow.Server
	// connection from proxy to server
	client := flow.Client

	serverAddr, _ := netip.ParseAddrPort(server.Addr.String())
	serverLocalAddr, _ := netip.ParseAddrPort(server.LocalAddr.String())
	serverIP := serverAddr.Addr()
	serverIPStr := serverIP.String()
	serverPort := uint64(serverAddr.Port())
	serverLocalIP := serverLocalAddr.Addr()
	serverLocalIPStr := serverLocalIP.String()
	serverIfaceIndex := uint64(ipToIfaceMap[serverLocalIPStr].Index)
	serverLocalPort := uint64(serverLocalAddr.Port())

	clientAddr, _ := netip.ParseAddrPort(client.Addr.String())
	clientLocalAddr, _ := netip.ParseAddrPort(client.LocalAddr.String())
	clientIP := clientAddr.Addr()
	clientIPStr := clientIP.String()
	clientPort := uint64(clientAddr.Port())
	clientLocalIP := clientLocalAddr.Addr()
	clientLocalIPStr := clientLocalIP.String()
	clientIfaceIndex := uint64(ipToIfaceMap[clientLocalIPStr].Index)
	clientLocalPort := uint64(clientLocalAddr.Port())

	serverFlow := hash(server, &serverIP, &serverLocalIP, &serverPort, &serverLocalPort, &serverIfaceIndex)
	clientFlow := hash(client, &clientLocalIP, &clientIP, &clientLocalPort, &clientPort, &clientIfaceIndex)
	proxyFlow := fnv1a.HashUint64(serial + serverFlow + clientFlow)

	serverFlowStr := strconv.FormatUint(serverFlow, 10)
	clientFlowStr := strconv.FormatUint(clientFlow, 10)
	proxyFlowStr := strconv.FormatUint(proxyFlow, 10)

	json := gabs.New()

	json.Set(serial, "serial")
	json.Set(proxyFlowStr, "flow")
	json.Set(*flow.ProjectID, "project")

	mdJSON, _ := json.Object("metadata")
	mdJSON.Set(mdIn, "in")
	mdJSON.Set(mdOut, "out")

	inJSON, _ := json.Object("in")
	outJSON, _ := json.Object("out")

	inJSON.Set(serverFlowStr, "flow")
	inSrcJSON, _ := inJSON.Object("src")
	inSrcJSON.Set(serverIPStr, "ip")
	inSrcJSON.Set(serverPort, "port")
	inSrcJSON.Set(server.Addr.Network(), "net")

	inDstJSON, _ := inJSON.Object("dst")
	inDstJSON.Set(serverLocalIPStr, "ip")
	inDstJSON.Set(serverLocalPort, "port")
	inDstJSON.Set(server.LocalAddr.Network(), "net")

	outJSON.Set(clientFlowStr, "flow")
	outSrcJSON, _ := outJSON.Object("src")
	outSrcJSON.Set(clientLocalIPStr, "ip")
	outSrcJSON.Set(clientLocalPort, "port")
	outSrcJSON.Set(client.LocalAddr.Network(), "net")

	outDstJSON, _ := outJSON.Object("dst")
	outDstJSON.Set(clientIPStr, "ip")
	outDstJSON.Set(clientPort, "port")
	outDstJSON.Set(client.Addr.Network(), "net")

	authorityJSON, _ := json.Object("authority")
	authorityJSON.Set(authorityParts[0], "src")
	authorityJSON.Set(target, "dst")

	operation, _ := json.Object("logging.googleapis.com/operation")
	operation.Set(stringFormatter.Format("src/{0}:{1}/pxy/{2}:{3}/dst/{4}:{5}/pid/{6}/rpc:{7}{8}",
		serverIPStr, serverPort, clientLocalIPStr, clientLocalPort, clientIPStr, clientPort, projectID, target, method), "producer")
	operation.Set(stringFormatter.Format("{0}/{1}/{2}", serverFlow, clientFlow, proxyFlow), "id")

	labelsJSON, _ := json.Object("logging.googleapis.com/labels")
	labelsJSON.Set("grpc-proxy", "tools.chux.dev/tool")
	labelsJSON.Set(target, "tools.chux.dev/grpc-proxy/authority")
	labelsJSON.Set(method, "tools.chux.dev/grpc/proxy/method")

	if flow.XCloudTraceContext != nil && *flow.XCloudTraceContext != "" {
		if traceAndSpan := xCloudTraceContextRegexp.FindStringSubmatch(*flow.XCloudTraceContext); traceAndSpan != nil {
			json.Set(stringFormatter.Format("projects/{0}/traces/{1}", projectID, traceAndSpan[1]), "logging.googleapis.com/trace")
			json.Set(traceAndSpan[2], "logging.googleapis.com/spanId")
			json.Set(true, "logging.googleapis.com/trace_sampled")
		}
	}

	timestampsJSON, _ := json.Object("timestamps")
	timestampsJSON.Set(flow.Timestamps.ProxyReceived.Format(time.RFC3339Nano), "proxyStart")
	timestampsJSON.Set(ts.Format(time.RFC3339Nano), "logger")
	timestampsJSON.Set(flow.Timestamps.DirectorStart.Format(time.RFC3339Nano), "directorStart")
	timestampsJSON.Set(flow.Timestamps.DirectorEnd.Format(time.RFC3339Nano), "directorEnd")
	timestampsJSON.Set(flow.Timestamps.Oauth2Start.Format(time.RFC3339Nano), "oauth2Start")
	timestampsJSON.Set(flow.Timestamps.Oauth2End.Format(time.RFC3339Nano), "oauth2End")
	timestampsJSON.Set(flow.Timestamps.StreamCreationStart.Format(time.RFC3339Nano), "streamCreationStart")
	timestampsJSON.Set(flow.Timestamps.StreamCreationEnd.Format(time.RFC3339Nano), "streamCreationEnd")
	timestampsJSON.Set(flow.Timestamps.StreamStart.Format(time.RFC3339Nano), "streamStart")

	latencyJSON, _ := json.Object("latency")
	latencyJSON.Set(flow.Timestamps.DirectorEnd.Sub(*flow.Timestamps.DirectorStart).Milliseconds(), "director")
	latencyJSON.Set(flow.Timestamps.Oauth2End.Sub(*flow.Timestamps.Oauth2Start).Milliseconds(), "oauth2")

	streamSetupLatency := flow.Timestamps.StreamCreationEnd.Sub(*flow.Timestamps.StreamCreationStart).Milliseconds()
	latencyJSON.Set(streamSetupLatency, "setup")

	srcConn := stringFormatter.Format("{0}:{1} > {2}:{3}", serverIPStr, serverPort, serverLocalIPStr, serverLocalPort)
	dstConn := stringFormatter.Format("{0}:{1} > {2}:{3}", clientLocalIPStr, clientLocalPort, clientIPStr, clientPort)

	// pre-populate some of the message to be shown in Cloud Logging
	message := stringFormatter.Format("#:{0} | src[{1}] >> dst[{2}] | project:{3} | rpc:{4}{5}",
		serial, srcConn, dstConn, projectID, target, method, streamSetupLatency)

	return json, &message
}

func onStreamStart(
	serverCtx, clientCtx context.Context, flow *proxy.ProxyFlow,
) {
	timestamp := time.Now()

	e2eLatencyMS := flow.Timestamps.StreamStart.Sub(*flow.Timestamps.ProxyReceived).Milliseconds()

	json, message := newJSONLog(serverCtx, clientCtx, flow, &timestamp)

	timestampJSON, _ := json.Object("timestamp")
	timestampJSON.Set(flow.Timestamps.StreamStart.Unix(), "seconds")
	timestampJSON.Set(flow.Timestamps.StreamStart.Nanosecond(), "nanos")

	rpcJSON, _ := json.Object("rpc")
	rpcJSON.Set(*flow.Target, "target")
	rpcJSON.Set(*flow.Method, "method")

	latencyJSON := json.S("latency")
	latencyJSON.Set(e2eLatencyMS, "e2e")

	logMessage := stringFormatter.Format("{0} | STREAM_START | {1}ms", *message, e2eLatencyMS)
	json.Set(logMessage, "message")

	io.WriteString(os.Stdout, json.String()+"\n")
}

func onStreamEnd(
	serverCtx, clientCtx context.Context,
	flow *proxy.ProxyFlow, stats *proxy.RPCStats,
) {
	timestamp := time.Now()

	serial := *flow.Serial
	defer flows.Delete(serial)

	e2eLatencyMS := flow.Timestamps.StreamEnd.Sub(*flow.Timestamps.ProxyReceived).Milliseconds()
	streamatencyMS := flow.Timestamps.StreamEnd.Sub(*flow.Timestamps.StreamStart).Milliseconds()

	json, message := newJSONLog(serverCtx, clientCtx, flow, &timestamp)

	timestampJSON, _ := json.Object("timestamp")
	timestampJSON.Set(flow.Timestamps.StreamEnd.Unix(), "seconds")
	timestampJSON.Set(flow.Timestamps.StreamEnd.Nanosecond(), "nanos")

	rpcJSON, _ := json.Object("rpc")
	rpcJSON.Set(*flow.Target, "target")
	rpcJSON.Set(*flow.Method, "method")

	timestampsJSON := json.S("timestamps")
	timestampsJSON.Set(flow.Timestamps.StreamEnd.Format(time.RFC3339Nano), "streamEnd")

	latencyJSON := json.S("latency")
	latencyJSON.Set(e2eLatencyMS, "e2e")
	latencyJSON.Set(streamatencyMS, "stream")

	statsJSON, _ := json.Object("stats")

	countersJSON, _ := statsJSON.Object("counters")
	countersJSON.Set(stats.Counters.Messages, "messages")
	countersJSON.Set(stats.Counters.Requests, "reqiests")
	countersJSON.Set(stats.Counters.Responses, "responses")
	countersJSON.Set(stats.Counters.Errors, "errors")

	overallCountersJSON, _ := countersJSON.Object("overall")
	overallCountersJSON.Set(stats.Counters.Overall.Flows, "flows")
	overallCountersJSON.Set(stats.Counters.Overall.Messages, "messages")
	overallCountersJSON.Set(stats.Counters.Overall.ForTarget, "forHost")
	overallCountersJSON.Set(stats.Counters.Overall.ForMethod, "forMethod")

	logMessage := stringFormatter.Format("{0} | STREAM_END | {1}ms", *message, streamatencyMS)
	json.Set(logMessage, "message")

	io.WriteString(os.Stdout, json.String()+"\n")
}

func logger(
	serverCtx, clientCtx context.Context,
	flow *proxy.ProxyFlow, rpc *proxy.RPC,
) {
	timestamp := time.Now()

	rpcLatency := rpc.Timestamps.End.Sub(*rpc.Timestamps.Start).Milliseconds()
	e2eLatencyMS := rpc.Timestamps.End.Sub(*flow.Timestamps.ProxyReceived).Milliseconds()

	json, message := newJSONLog(serverCtx, clientCtx, flow, &timestamp)

	rpcJSON, _ := json.Object("rpc")

	rpcJSON.Set(*rpc.Target, "target")
	rpcJSON.Set(*flow.Method, "method")

	logMessage := ""

	if rpc.IsRequest || rpc.IsResponse {

		var rpcMessageJSON *gabs.Container
		if rpc.IsRequest {
			rpcMessageJSON, _ = rpcJSON.Object("request")
		} else if rpc.IsResponse {
			rpcMessageJSON, _ = rpcJSON.Object("response")
		}

		protoJSONstr := protojson.Format(*rpc.MessageProto)
		protoJSON, _ := gabs.ParseJSON([]byte(protoJSONstr))

		messageFullName := string(*rpc.MessageFullName)

		rpcMessageJSON.Set(protoJSON, "proto")
		rpcMessageJSON.Set(messageFullName, "type")

		labelsJSON := json.S("logging.googleapis.com/labels")
		labelsJSON.Set(messageFullName, "tools.chux.dev/grpc/proxy/message")

		logMessage = stringFormatter.Format("{0} | msg:{1} | {2}ms",
			*message, messageFullName, rpcLatency)

	}

	if rpc.StatusProto != nil {
		json.Set("ERROR", "severity")

		statusJSON, _ := rpcJSON.Object("status")
		statusJSON.Set(rpc.StatusProto.Code, "code")
		statusJSON.Set(rpc.StatusProto.Message, "message")

		if rpc.IsRequest || rpc.IsResponse {
			logMessage = stringFormatter.Format("{0} | code:{1} | {2}",
				logMessage, rpc.StatusProto.GetCode(), rpc.StatusProto.GetMessage())
		} else {
			logMessage = stringFormatter.Format("{0} | code:{1} | {2} | {3}ms",
				*message, rpc.StatusProto.GetCode(), rpc.StatusProto.GetMessage(), rpcLatency)
		}
	}

	json.Set(logMessage, "message")

	timestampJSON, _ := json.Object("timestamp")
	timestampJSON.Set(rpc.Timestamps.Start.Unix(), "seconds")
	timestampJSON.Set(rpc.Timestamps.Start.Nanosecond(), "nanos")

	timestampsJSON := json.S("timestamps")
	timestampsJSON.Set(rpc.Timestamps.Start.Format(time.RFC3339Nano), "rpcStart")
	timestampsJSON.Set(rpc.Timestamps.End.Format(time.RFC3339Nano), "rpcEnd")

	latencyJSON := json.S("latency")
	latencyJSON.Set(e2eLatencyMS, "e2e")
	latencyJSON.Set(rpcLatency, "rpc")
	latencyJSON.Set(rpc.Timestamps.Start.Sub(*flow.Timestamps.ProxyReceived).Milliseconds(), "proxy")

	statsJSON, _ := json.Object("stats")
	countersJSON, _ := statsJSON.Object("counters")
	countersJSON.Set(rpc.Stats.Counters.Messages, "messages")
	countersJSON.Set(rpc.Stats.Counters.Requests, "reqiests")
	countersJSON.Set(rpc.Stats.Counters.Responses, "responses")
	countersJSON.Set(rpc.Stats.Counters.Errors, "errors")

	io.WriteString(os.Stdout, json.String()+"\n")
}

func streamClientInterceptor(
	serverCtx context.Context,
	desc *grpc.StreamDesc,
	cc *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	tsBeforeStreamCreation := time.Now()
	clientStream, err := streamer(serverCtx, desc, cc, method, opts...)
	tsAfterStreamCreation := time.Now()

	if err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		return clientStream, err
	}

	mdIn, _ := metadata.FromOutgoingContext(serverCtx)
	flow, flowErr := getProxyFlowFromMetadata(mdIn)
	if flowErr != nil {
		fmt.Fprint(os.Stderr, flowErr.Error()+"\n")
		return clientStream, err
	}

	clientCtx := clientStream.Context()
	// connection from proxy to server
	client, _ := peer.FromContext(clientCtx)

	flow.Client = client
	flow.Timestamps.StreamCreationStart = &tsBeforeStreamCreation
	flow.Timestamps.StreamCreationEnd = &tsAfterStreamCreation

	return clientStream, err
}

func rpcTrafficDirector(
	serverCtx context.Context,
	flow *proxy.ProxyFlow,
) (
	context.Context,
	*grpc.ClientConn,
	proxy.ProxyLogger,
	*proxy.ProxyFlowHandlers,
	error,
) {
	tsDirectorStart := time.Now()

	serial := *flow.Serial
	flows.Store(serial, flow)

	md, _ := metadata.FromIncomingContext(serverCtx)

	_target := target
	_connectionFactoryProvider := defaultConnectionFactoryProvider

	// [ToDo]:
	//   - allow whitelisting endpoints and methods
	//   - ratelimit? on max-concurrent-rpc per project/host/method
	rpcEndpointHeader := md.Get("x-grpc-proxy-endpoint")
	if len(rpcEndpointHeader) > 0 && rpcEndpointHeader[0] != target { // override
		_target = rpcEndpointHeader[0]
	} else if flow.Target != nil {
		_target = *flow.Target + ":443"
	}

	if _target != target {
		_connectionFactoryProvider = func() connectionFactory {
			return newConnectionFactory(&_target)
		}
	}

	var _connectionFactory connectionFactory
	connectionFactoryLoaded := false

	_connectionFactory, connectionFactoryLoaded = connectionFactories.LoadOrStoreLazy(_target, _connectionFactoryProvider)
	if !connectionFactoryLoaded {
		errorMsg := stringFormatter.Format("failed to create connection factory for: {0} ", _target)
		io.WriteString(os.Stderr, errorMsg+"\n")
		return serverCtx, nil, nil, nil, errors.New(errorMsg)
	}

	tsOauth2Start := time.Now()
	token, err := getToken()
	tsOauth2End := time.Now()
	if err != nil {
		io.WriteString(os.Stderr, err.Error()+"\n")
		return serverCtx, nil, nil, nil, err
	}
	md.Set("Authorization", "Bearer "+token.AccessToken)

	rpcConnection := _connectionFactory()
	if rpcConnection == nil {
		errorMsg := stringFormatter.Format("failed to create connection for: {0}", _target)
		io.WriteString(os.Stderr, errorMsg+"\n")
		return serverCtx, nil, nil, nil, errors.New(errorMsg)
	}
	rpcTarget := rpcConnection.Target()
	flow.Target = &rpcTarget

	// connection from client to proxy
	server, _ := peer.FromContext(serverCtx)
	flow.Server = server

	flow.ClientConn = rpcConnection
	flow.XCloudTraceContext = nil
	flow.Timestamps.DirectorStart = &tsDirectorStart
	flow.Timestamps.Oauth2Start = &tsOauth2Start
	flow.Timestamps.Oauth2End = &tsOauth2End

	projectIDHeader := md.Get("x-grpc-proxy-project")
	if len(projectIDHeader) > 0 {
		_projectID := projectIDHeader[0]
		md.Set("x-goog-user-project", _projectID)
		flow.ProjectID = &_projectID
	} else {
		md.Set("x-goog-user-project", projectID)
		flow.ProjectID = &projectID
	}

	xCloudTraceContextHeader := md.Get("x-cloud-trace-context")
	if len(xCloudTraceContextHeader) > 0 {
		flow.XCloudTraceContext = &xCloudTraceContextHeader[0]
	}

	md.Set("x-grpc-proxy-id", strconv.FormatUint(serial, 10))
	md.Set("X-Forwarded-For", xForwardedFor)
	md.Append("user-agent", "grpc-proxy/1.0.0")

	ctx := metadata.NewOutgoingContext(serverCtx, md.Copy())

	tsDirectorEnd := time.Now()
	flow.Timestamps.DirectorEnd = &tsDirectorEnd

	handlers := &proxy.ProxyFlowHandlers{
		OnStreamStart: onStreamStart,
		OnStreamEnd:   onStreamEnd,
	}

	return ctx, rpcConnection, logger, handlers, nil
}

func contextDialer(ctx context.Context, addr string) (net.Conn, error) {
	dst, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	tcpConn, tcpErr := dialer.DialContext(ctx, "tcp", dst.String())
	if tcpErr != nil {
		return nil, tcpErr
	}
	return tcpConn, tcpErr
}

func main() {
	proxyPort, proxyPortError := strconv.Atoi(proxyPort)
	if proxyPortError != nil {
		proxyPort = defaultProxyPort
	}

	serverListener, listenerErr := net.Listen("tcp", fmt.Sprintf(":%d", proxyPort))

	if listenerErr != nil {
		fmt.Fprintf(os.Stderr, "failed to start gRPC proxy at: %d\n", proxyPort)
		os.Exit(1)
	} else {
		fmt.Fprintf(os.Stderr, "gRPC proxy listening at: %d\n", proxyPort)
	}

	ipToIfaceMap = make(map[string]net.Interface)
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			ipToIfaceMap[addr.String()] = iface
		}
	}

	connectionFactories = skipmap.New[string, connectionFactory]()
	connections = skipmap.New[string, *grpc.ClientConn]()
	flows = skipmap.New[uint64, *proxy.ProxyFlow]()

	encoding.RegisterCodec(proxy.Codec(encoding.GetCodec("proto")))
	streamHandler := proxy.TransparentHandler(rpcTrafficDirector)

	proxyServer := grpc.NewServer(
		grpc.UnknownServiceHandler(streamHandler),
		grpc.MaxConcurrentStreams(100),
		// grpc.KeepaliveParams(keepAliveServerParams),
		grpc.KeepaliveEnforcementPolicy(keepAliveEnforcementPolicy),
	)

	// [ToDo]: listener must be a facade to handle connections:
	//           - intercept H2 keepalives ( Ping frames )
	proxyServer.Serve(serverListener)
}
