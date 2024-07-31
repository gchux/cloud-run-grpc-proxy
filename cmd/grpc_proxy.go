package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/netip"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/Jeffail/gabs/v2"
	"github.com/alphadose/haxmap"
	"github.com/gchux/grpc-proxy/pkg/proxy"
	"github.com/segmentio/fasthash/fnv1a"
	"github.com/wissance/stringFormatter"

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
	ProxyFlow struct {
		Serial                 *uint64
		Method                 *string
		Client, Server         *peer.Peer
		TsProxyReceived        *time.Time
		TsBeforeStreamCreation *time.Time
		TsAfterStreamCreation  *time.Time
	}

	clientConnFactory func() *grpc.ClientConn
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

	rpcCounters *haxmap.Map[string, *atomic.Uint64]

	endpoints *haxmap.Map[string, *grpc.ClientConn]
	flows     *haxmap.Map[uint64, *proxy.ProxyFlow]

	defaultClientConnFactory = newClientConnFactory(&target)

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
		l3Hash += 37 // IPv6 == 41
	}
	srcBytes := srcIP.As16()
	dstBytes := dstIP.As16()
	return l3Hash +
		fnv1a.HashBytes64(srcBytes[:]) +
		fnv1a.HashBytes64(dstBytes[:])
}

func hashL4(p *peer.Peer, srcPort, dstPort *uint64) uint64 {
	proto := uint64(6) // TCP(6)
	if strings.HasPrefix(p.LocalAddr.Network(), "udp") {
		proto += uint64(11) // UDP(17)
	}
	return proto + *srcPort + *dstPort
}

func hash(p *peer.Peer, srcIP, dstIP *netip.Addr, srcPort, dstPort, ifaceIndex *uint64) uint64 {
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

	flow, ok := flows.Get(rpc)
	if !ok {
		return nil, errors.New("RPC flow not found: " + grpcProxyID[0])
	}

	return flow, nil
}

func newClientConnFactory(target *string) clientConnFactory {
	return func() *grpc.ClientConn {
		// [ToDo]: validate `target`
		cc, err := grpc.NewClient(*target,
			grpc.WithTransportCredentials(tlsCreds),
			grpc.WithContextDialer(contextDialer),
			grpc.WithUserAgent("grpc-proxy/1.0.0"),
			grpc.WithStreamInterceptor(streamClientInterceptor),
			grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`))
		if err != nil {
			return nil
		}
		return cc
	}
}

func onStreamEnd(
	serverCtx, clientCtx context.Context,
	flow *proxy.ProxyFlow,
) {
	timestamp := time.Now()

	serial := *flow.Serial
	defer flows.Del(serial)

	_projectID := *flow.ProjectID
	endpoint := *flow.Endpoint
	method := *flow.Method

	countByMethod := *flow.Stats.Counters.ByMethod

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
	proxyFlow := fnv1a.HashUint64(serverFlow + clientFlow)

	serverFlowStr := strconv.FormatUint(serverFlow, 10)
	clientFlowStr := strconv.FormatUint(clientFlow, 10)
	proxyFlowStr := strconv.FormatUint(proxyFlow, 10)

	json := gabs.New()

	json.Set(serial, "serial")
	json.Set(proxyFlowStr, "flow")

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

	rpcJSON, _ := json.Object("rpc")
	rpcJSON.Set(serial, "serial")
	rpcJSON.Set(_projectID, "project")

	methodJSON, _ := rpcJSON.Object("method")
	methodJSON.Set(method, "id")
	methodJSON.Set(strconv.FormatUint(countByMethod, 10), "serial")

	authorityJSON, _ := rpcJSON.Object("authority")
	authorityJSON.Set(authorityParts[0], "src")
	authorityJSON.Set(endpoint, "dst")

	rpcRequestJSON, _ := rpcJSON.Object("request")
	jsonRquest := protojson.Format(flow.ProtoRequest)
	protoRequest, _ := gabs.ParseJSON([]byte(jsonRquest))
	rpcRequestJSON.Set(protoRequest, "proto")

	rpcResponseJSON, _ := rpcJSON.Object("response")
	if flow.StatusProto != nil {
		status := flow.StatusProto
		rpcResponseJSON.Set(status.GetCode(), "status")
		rpcResponseJSON.Set(status.GetMessage(), "message")
	} else {
		rpcResponseJSON.Set(0, "status")
		jsonResponse := protojson.Format(flow.ProtoResponse)
		protoResponse, _ := gabs.ParseJSON([]byte(jsonResponse))
		rpcResponseJSON.Set(protoResponse, "proto")
	}

	timestampJSON, _ := json.Object("timestamp")
	timestampJSON.Set(flow.TsProxyReceived.Unix(), "seconds")
	timestampJSON.Set(flow.TsProxyReceived.Nanosecond(), "nanos")

	tsBeforeStreamCreation := *flow.TsBeforeStreamCreation
	tsAfterStreamCreation := *flow.TsAfterStreamCreation
	streamSetupLatency := tsAfterStreamCreation.Sub(tsBeforeStreamCreation).Milliseconds()
	rpcLatency := flow.TsStreamEnd.Sub(*flow.TsStreamStart).Milliseconds()
	e2eLatencyMS := flow.TsStreamEnd.Sub(*flow.TsProxyReceived).Milliseconds()

	timestampsJSON, _ := json.Object("timestamps")
	// [ToDo]: consider creating an object for each pair
	timestampsJSON.Set(flow.TsProxyReceived.String(), "proxyStart")
	timestampsJSON.Set(timestamp.String(), "proxyEnd")
	timestampsJSON.Set(flow.TsDirectorStart.String(), "directorStart")
	timestampsJSON.Set(flow.TsDirectorEnd.String(), "directorEnd")
	timestampsJSON.Set(flow.TsOauth2Start.String(), "oauth2Start")
	timestampsJSON.Set(flow.TsOauth2End.String(), "oauth2End")
	timestampsJSON.Set(flow.TsStreamStart.String(), "streamStart")
	timestampsJSON.Set(flow.TsStreamEnd.String(), "streamEnd")
	timestampsJSON.Set(flow.TsBeforeStreamCreation.String(), "streamCreationStart")
	timestampsJSON.Set(flow.TsAfterStreamCreation.String(), "streamCreationEnd")

	latencyJSON, _ := json.Object("latency")
	latencyJSON.Set(e2eLatencyMS, "e2e")
	latencyJSON.Set(rpcLatency, "rpc")
	latencyJSON.Set(flow.TsDirectorEnd.Sub(*flow.TsDirectorStart).Milliseconds(), "director")
	latencyJSON.Set(flow.TsOauth2End.Sub(*flow.TsOauth2Start).Milliseconds(), "oauth2")
	latencyJSON.Set(flow.TsStreamStart.Sub(*flow.TsProxyReceived).Milliseconds(), "proxy") // overhead
	latencyJSON.Set(streamSetupLatency, "setup")

	srcConn := stringFormatter.Format("{0}:{1} > {2}:{3}", serverIPStr, serverPort, serverLocalIPStr, serverLocalPort)
	dstConn := stringFormatter.Format("{0}:{1} > {2}:{3}", clientLocalIPStr, clientLocalPort, clientIPStr, clientPort)

	operation, _ := json.Object("logging.googleapis.com/operation")
	operation.Set(stringFormatter.Format("src/{0}:{1}/pxy/{2}:{3}/dst/{4}:{5}/pid/{6}/rpc:{7}{8}",
		serverIPStr, serverPort, clientLocalIPStr, clientLocalPort, clientIPStr, clientPort, _projectID, endpoint, method), "producer")
	operation.Set(stringFormatter.Format("{0}/{1}/{2}", serverFlow, clientFlow, proxyFlow), "id")

	labels, _ := json.Object("logging.googleapis.com/labels")
	labels.Set("grpc-proxy", "tools.chux.dev/tool")
	labels.Set(authorityParts[0], "tools.chux.dev/grpc-proxy/authority/src")
	labels.Set(endpoint, "tools.chux.dev/grpc-proxy/authority/dst")
	labels.Set(method, "tools.chux.dev/grpc/proxy/method")

	if flow.XCloudTraceContext != nil && *flow.XCloudTraceContext != "" {
		if traceAndSpan := xCloudTraceContextRegexp.FindStringSubmatch(*flow.XCloudTraceContext); traceAndSpan != nil {
			json.Set(stringFormatter.Format("projects/{0}/traces/{1}", _projectID, traceAndSpan[1]), "logging.googleapis.com/trace")
			json.Set(traceAndSpan[2], "logging.googleapis.com/spanId")
			json.Set(true, "logging.googleapis.com/trace_sampled")
		}
	}

	message := stringFormatter.Format("#:{0}/{1} | src[{2}] >> dst[{3}] | project:{4} | rpc:{5}{6} | latency[setup:{7}|rpc:{8}|e2e:{9}]ms",
		serial, countByMethod, srcConn, dstConn, _projectID, endpoint, method, streamSetupLatency, rpcLatency, e2eLatencyMS)
	json.Set(message, "message")
	fmt.Println(json.String())
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
	flow.TsBeforeStreamCreation = &tsBeforeStreamCreation
	flow.TsAfterStreamCreation = &tsAfterStreamCreation

	return clientStream, err
}

func rpcTrafficDirector(
	serverCtx context.Context,
	flow *proxy.ProxyFlow,
) (
	context.Context,
	*grpc.ClientConn,
	proxy.OnStreamEnd,
	error,
) {
	tsDirectorStart := time.Now()

	serial := *flow.Serial
	flows.Set(serial, flow)

	// count RPCs by method
	// [ToDO]: consider counting by project And method
	byMethodCounter, _ := rpcCounters.GetOrCompute(*flow.Method,
		func() *atomic.Uint64 {
			return new(atomic.Uint64)
		})

	byMethodCount := byMethodCounter.Add(1)

	md, _ := metadata.FromIncomingContext(serverCtx)

	var rpcConn *grpc.ClientConn
	rpcConnLoaded := false

	// [ToDo]:
	//   - allow whitelisting endpoints and methods
	//   - ratelimit on max-concurrent-rpc per project/host/method
	rpcEndpointHeader := md.Get("x-grpc-proxy-endpoint")
	if len(rpcEndpointHeader) > 0 && rpcEndpointHeader[0] != target {
		rpcEndpoint := rpcEndpointHeader[0]
		rpcConn, rpcConnLoaded = endpoints.GetOrCompute(rpcEndpoint, newClientConnFactory(&rpcEndpoint))
	} else {
		rpcConn, rpcConnLoaded = endpoints.GetOrCompute(target, defaultClientConnFactory)
	}

	if !rpcConnLoaded && rpcConn == nil {
		return serverCtx, nil, nil, errors.New("failed to create connection")
	}

	tsOauth2Start := time.Now()
	token, err := getToken()
	tsOauth2End := time.Now()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		return serverCtx, nil, nil, err
	}
	md.Set("Authorization", "Bearer "+token.AccessToken)

	rpcEndpoint := rpcConn.Target()

	// connection from client to proxy
	server, _ := peer.FromContext(serverCtx)

	flow.ClientConn = rpcConn
	flow.Endpoint = &rpcEndpoint
	flow.Server = server
	flow.XCloudTraceContext = nil
	flow.TsDirectorStart = &tsDirectorStart
	flow.TsOauth2Start = &tsOauth2Start
	flow.TsOauth2End = &tsOauth2End
	flow.Stats.Counters.ByMethod = &byMethodCount

	md.Set(":authority", strings.SplitN(rpcEndpoint, ":", 2)[0])

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
	flow.TsDirectorEnd = &tsDirectorEnd

	return ctx, rpcConn, onStreamEnd, nil
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

	endpoints = haxmap.New[string, *grpc.ClientConn]()
	flows = haxmap.New[uint64, *proxy.ProxyFlow]()
	rpcCounters = haxmap.New[string, *atomic.Uint64]()

	encoding.RegisterCodec(proxy.Codec(encoding.GetCodec("proto")))
	streamHandler := proxy.TransparentHandler(rpcTrafficDirector)

	proxyServer := grpc.NewServer(
		grpc.UnknownServiceHandler(streamHandler),
		grpc.MaxConcurrentStreams(100),
		// grpc.KeepaliveParams(keepAliveServerParams),
		grpc.KeepaliveEnforcementPolicy(keepAliveEnforcementPolicy),
	)

	proxyServer.Serve(serverListener)
}
