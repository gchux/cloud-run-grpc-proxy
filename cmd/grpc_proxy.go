package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"net/netip"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/Jeffail/gabs/v2"
	"github.com/gchux/grpc-proxy/pkg/proxy"
	"github.com/segmentio/fasthash/fnv1a"
	"github.com/wissance/stringFormatter"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	_ "google.golang.org/grpc/status"

	"golang.org/x/oauth2"
	auth "golang.org/x/oauth2/google"
	"golang.org/x/sys/unix"
)

const (
	defaultProxyPort = 5001
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
)

var ipToIfaceMap map[string]net.Interface

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
	var l3Hash uint64 = 4 // IPv4 == 4
	if srcIP.Is6() {
		l3Hash += 37 // IPv6 == 41
	}
	srcBytes := srcIP.As16()
	dstBytes := dstIP.As16()
	return l3Hash + fnv1a.HashBytes64(srcBytes[:]) + fnv1a.HashBytes64(dstBytes[:])
}

func hashL4(srcPort, dstPort *uint64) uint64 {
	// [ToDo]: use network to discover proto
	return 6 + *srcPort + *dstPort
}

func hash(srcIP, dstIP *netip.Addr, srcPort, dstPort, ifaceIndex *uint64) uint64 {
	hash := fnv1a.AddUint64(fnv1a.Init64, *ifaceIndex)
	hash = fnv1a.AddUint64(hash, hashL3(srcIP, dstIP))
	return fnv1a.AddUint64(hash, hashL4(srcPort, dstPort))
}

func newStreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		// [ToDo]: add atomic RPC counter
		clientStream, err := streamer(ctx, desc, cc, method, opts...)
		timestamp := time.Now()

		mdIn, _ := metadata.FromIncomingContext(ctx)
		mdOut, _ := metadata.FromOutgoingContext(clientStream.Context())
		authority := mdIn.Get(":authority")[0]
		authorityParts := strings.Split(authority, ":")

		// connection from client to proxy
		server, _ := peer.FromContext(ctx)
		serverAddr, _ := netip.ParseAddrPort(server.Addr.String())
		serverLocalAddr, _ := netip.ParseAddrPort(server.LocalAddr.String())
		serverIP := serverAddr.Addr()
		serverIPStr := serverIP.String()
		serverPort := uint64(serverAddr.Port())
		serverLocalIP := serverLocalAddr.Addr()
		serverLocalIPStr := serverLocalIP.String()
		serverIfaceIndex := uint64(ipToIfaceMap[serverLocalIPStr].Index)
		serverLocalPort := uint64(serverLocalAddr.Port())

		// connection from proxy to server
		client, _ := peer.FromContext(clientStream.Context())
		clientAddr, _ := netip.ParseAddrPort(client.Addr.String())
		clientLocalAddr, _ := netip.ParseAddrPort(client.LocalAddr.String())
		clientIP := clientAddr.Addr()
		clientIPStr := clientIP.String()
		clientPort := uint64(clientAddr.Port())
		clientLocalIP := clientLocalAddr.Addr()
		clientLocalIPStr := clientLocalIP.String()
		clientIfaceIndex := uint64(ipToIfaceMap[clientLocalIPStr].Index)
		clientLocalPort := uint64(clientLocalAddr.Port())

		serverFlow := hash(&serverIP, &serverLocalIP, &serverPort, &serverLocalPort, &serverIfaceIndex)
		clientFlow := hash(&clientLocalIP, &clientIP, &clientLocalPort, &clientPort, &clientIfaceIndex)
		proxyFlow := fnv1a.HashUint64(serverFlow + clientFlow)

		json := gabs.New()

		json.Set(proxyFlow, "flow")

		mdJSON, _ := json.Object("metadata")
		mdJSON.Set(mdIn, "in")
		mdJSON.Set(mdOut, "out")

		inJSON, _ := json.Object("in")
		inJSON.Set(serverFlow, "flow")
		inSrcJSON, _ := inJSON.Object("src")
		inSrcJSON.Set(serverIPStr, "ip")
		inSrcJSON.Set(serverPort, "port")
		inSrcJSON.Set(server.Addr.Network(), "net")
		inDstJSON, _ := inJSON.Object("dst")
		inDstJSON.Set(serverLocalIPStr, "ip")
		inDstJSON.Set(serverLocalPort, "port")
		inDstJSON.Set(server.LocalAddr.Network(), "net")

		outJSON, _ := json.Object("out")
		outJSON.Set(clientFlow, "flow")
		outSrcJSON, _ := outJSON.Object("src")
		outSrcJSON.Set(clientLocalIPStr, "ip")
		outSrcJSON.Set(clientLocalPort, "port")
		outSrcJSON.Set(client.LocalAddr.Network(), "net")
		outDstJSON, _ := outJSON.Object("dst")
		outDstJSON.Set(clientIPStr, "ip")
		outDstJSON.Set(clientPort, "port")
		outDstJSON.Set(client.Addr.Network(), "net")

		rpcJSON, _ := json.Object("rpc")
		rpcJSON.Set(method, "method")
		authorityJSON, _ := rpcJSON.Object("authority")
		authorityJSON.Set(authorityParts[0], "src")
		authorityJSON.Set(targetHost, "dst")

		timestampJSON, _ := json.Object("timestamp")
		timestampJSON.Set(timestamp.Unix(), "seconds")
		timestampJSON.Set(timestamp.Nanosecond(), "nanos")

		srcConn := stringFormatter.Format("{0}:{1} > {2}:{3}", serverIPStr, serverPort, serverLocalIPStr, serverLocalPort)
		dstConn := stringFormatter.Format("{0}:{1} > {2}:{3}", clientLocalIPStr, clientLocalPort, clientIPStr, clientPort)

		operation, _ := json.Object("logging.googleapis.com/operation")
		operation.Set(stringFormatter.Format("grpc-proxy/src/{0}:{1}/dst/{2}:{3}/rpc:{5}", serverIPStr, serverPort, targetHost, clientPort, targetHost, method), "producer")
		operation.Set(stringFormatter.Format("{0}/{1}/{2}", serverFlow, clientFlow, proxyFlow), "id")

		labels, _ := json.Object("logging.googleapis.com/labels")
		labels.Set("grpc-proxy", "tools.chux.dev/tool")
		labels.Set(authorityParts[0], "tools.chux.dev/grpc-proxy/authority/src")
		labels.Set(targetHost, "tools.chux.dev/grpc-proxy/authority/dst")
		labels.Set(method, "tools.chux.dev/grpc/proxy/method")

		message := stringFormatter.Format("src[{0}] >> dst[{1}] | rpc[{2}{3}]", srcConn, dstConn, targetHost, method)
		json.Set(message, "message")
		fmt.Println(json.String())

		return clientStream, err
	}
}

func newDirector(cc *grpc.ClientConn) proxy.StreamDirector {
	return func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		md, _ := metadata.FromIncomingContext(ctx)

		md.Set(":authority", targetHost)
		md.Set("X-Forwarded-For", xForwardedFor)
		md.Set("x-goog-user-project", projectID)

		md.Append("user-agent", "grpc-proxy/1.0.0")

		token, err := getToken()
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			return ctx, nil, err
		}
		md.Set("Authorization", "Bearer "+token.AccessToken)

		ctx = metadata.NewOutgoingContext(ctx, md.Copy())
		return ctx, cc, nil
	}
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

	ipToIfaceMap = make(map[string]net.Interface)
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			ipToIfaceMap[addr.String()] = iface
		}
	}

	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})

	cc, err := grpc.NewClient(target,
		grpc.WithTransportCredentials(creds),
		grpc.WithContextDialer(contextDialer),
		grpc.WithUserAgent("grpc-proxy/1.0.0"),
		grpc.WithStreamInterceptor(newStreamClientInterceptor()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`))
	if err != nil {
		log.Fatalf("dialing %s: %v", target, err)
		os.Exit(1)
	}

	streamHandler := proxy.TransparentHandler(newDirector(cc))
	opt := grpc.UnknownServiceHandler(streamHandler)
	opts := []grpc.ServerOption{opt}

	proxy := grpc.NewServer(opts...)
	serverListener, _ := net.Listen("tcp", fmt.Sprintf(":%d", proxyPort))
	proxy.Serve(serverListener)
}
