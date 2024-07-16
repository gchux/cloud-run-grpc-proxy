package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/netip"
	"os"
	"strconv"
	"time"

	"github.com/Jeffail/gabs/v2"
	"github.com/mwitkow/grpc-proxy/proxy"
	"github.com/wissance/stringFormatter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"

	"golang.org/x/oauth2"
	auth "golang.org/x/oauth2/google"
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

func getToken() (*oauth2.Token, error) {
	ctx := context.Background()
	credentials, err := auth.FindDefaultCredentials(ctx, scopes...)
	credentials.ProjectID = projectID
	if err == nil {
		return credentials.TokenSource.Token()
	}
	return nil, err
}

func newDirector(cc *grpc.ClientConn) proxy.StreamDirector {
	return func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		md, _ := metadata.FromIncomingContext(ctx)

		client, _ := peer.FromContext(ctx)
		clientAddr, _ := netip.ParseAddrPort(client.Addr.String())

		md.Set("X-Forwarded-For", xForwardedFor)
		md.Set("x-goog-user-project", projectID)

		json := gabs.New()

		targetJSON, _ := json.Object("target")
		targetJSON.Set(targetHost, "host")
		targetJSON.Set(targetPort, "port")
		targetJSON.Set(fullMethodName, "method")

		json.Set(md, "metadata")

		clientJSON, _ := json.Object("client")
		clientJSON.Set(clientAddr.Addr(), "ip")
		clientJSON.Set(clientAddr.Port(), "port")

		timestamp := time.Now()
		timestampJSON, _ := json.Object("timestamp")
		timestampJSON.Set(timestamp.Unix(), "seconds")
		timestampJSON.Set(timestamp.Nanosecond(), "nanos")

		clientStr := clientAddr.String()

		operation, _ := json.Object("logging.googleapis.com/operation")
		operation.Set(logsProducer, "producer")
		operation.Set(stringFormatter.Format("{0}/{1}", clientStr, target), "id")

		labels, _ := json.Object("logging.googleapis.com/labels")
		labels.Set(targetHost, "target")
		labels.Set(fullMethodName, "method")

		message := stringFormatter.Format("{0} => {1}{2}", clientStr, target, fullMethodName)

		token, err := getToken()
		if err != nil {
			json.Set("ERROR", "severity")
			json.Set(stringFormatter.Format("{0} | ERROR: {1}", message, err.Error()), "message")
			fmt.Println(json.String())
			return ctx, cc, err
		}

		json.Set(message, "message")
		fmt.Println(json.String())

		md.Set("Authorization", "Bearer "+token.AccessToken)

		ctx = metadata.NewOutgoingContext(ctx, md.Copy())
		return ctx, cc, nil
	}
}

func main() {
	proxyPort, proxyPortError := strconv.Atoi(proxyPort)
	if proxyPortError != nil {
		proxyPort = defaultProxyPort
	}

	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	cc, err := grpc.DialContext(ctx, target, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("dialing %s: %v", target, err)
		cancel()
		os.Exit(1)
	}

	streamHandler := proxy.TransparentHandler(newDirector(cc))
	opt := grpc.UnknownServiceHandler(streamHandler)
	opts := []grpc.ServerOption{opt}

	// proxy := proxy.NewProxy(cc)
	proxy := grpc.NewServer(opts...)
	serverListener, _ := net.Listen("tcp", fmt.Sprintf(":%d", proxyPort))
	proxy.Serve(serverListener)
	cancel()
}
