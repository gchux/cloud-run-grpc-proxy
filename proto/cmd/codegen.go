package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
	"text/template"

	mapset "github.com/deckarep/golang-set/v2"
)

var (
	codegenSrc = flag.String("codegen_src", "", "")
	codegenDst = flag.String("codegen_dst", "", "")
)

const (
	googleCloudGoBasePackage = "cloud.google.com/go/"

	tmpl = `// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/genproto/googleapis/rpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/zhangyunhao116/skipmap"

{{- range $alias, $import := .imports }}
	{{$alias}} "{{$import -}}"
{{- end }}
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
{{- range $service, $host := .svcHost }}
	service2Host.Store("{{$service}}", "{{$host}}")
{{- end }}

{{- range $method, $request := .rpcReq }}
	method2RequestType.Store("{{$method}}", reflect.TypeOf((*{{ index $.rpcReqGoPkg $method }}.{{$request}})(nil)).Elem())
	method2ResponseType.Store("{{$method}}", reflect.TypeOf((*{{ index $.rpcResGoPkg $method }}.{{ index $.rpcRes $method }})(nil)).Elem())
{{- end }}
}
`
)

// GeneratedCodeRegexp checks the code generation standard
// defined at https://golang.org/s/generatedcode.
var GeneratedCodeRegexp = regexp.MustCompile(`(?m:^// Code generated .* DO NOT EDIT\.$)`)

// CreateFile runs the "text/template".Template with data, pass it through gofmt
// and saves it to filePath.
func CreateFile(filePath string, t *template.Template, data interface{}) (err error) {
	return (&CodeTemplate{Template: t}).CreateFile(filePath, data)
}

// CodeTemplate is the precompiled template for generating one or multiple Go source files.
type CodeTemplate struct {
	Template *template.Template // See "text/template"
	Buffer   bytes.Buffer       // Used for sharing allocated memory between multiple CreateFile calls
	mu       sync.Mutex
}

// Parse creates a CodeTemplate from a "text/template" source.
//
// The expansion of the template is expected to be valid a Go source file
// containing the code generation standard tag. See GeneratedCodeRegexp.
func Parse(codeTemplate string) (*CodeTemplate, error) {
	t, err := template.New("").Parse(codeTemplate)
	if err != nil {
		return nil, err
	}
	return &CodeTemplate{Template: t}, nil
}

// MustParse wraps Parse throwing errors as exception.
func MustParse(codeTemplate string) *CodeTemplate {
	tmpl, err := Parse(codeTemplate)
	if err != nil {
		panic(err)
	}
	return tmpl
}

// CreateFile runs the template with data, pass it through gofmt
// and saves it to filePath.
//
// The code generation standard at https://golang.org/s/generatedcode is enforced.
func (t *CodeTemplate) CreateFile(filePath string, data interface{}) error {
	// This anonymous function exists just to wrap the mutex protected block
	out, err := func() ([]byte, error) {
		// To protect t.Buffer
		t.mu.Lock()
		defer t.mu.Unlock()

		t.Buffer.Reset()

		if err := t.Template.Execute(&t.Buffer, data); err != nil {
			return nil, err
		}

		code := t.Buffer.Bytes()

		// Enforce code generation standard https://golang.org/s/generatedcode
		if !GeneratedCodeRegexp.Match(code) {
			return nil, errors.New("output does not follow standard defined at https://golang.org/s/generatedcode")
		}

		return format.Source(code)
		//  return code, nil
	}()
	if err != nil {
		return err
	}

	f, err := os.Create(filePath)
	if err == nil {
		defer f.Close()
		_, err = f.Write(out)
	}
	return err
}

func processMessages(pkg, ns *string, messages []*Message, messageType2namespace map[string]string) {
	if len(messages) == 0 {
		return
	}
	for _, msg := range messages {
		messageType := *pkg + "." + *msg.Type
		messageType2namespace[messageType] = *ns
		processMessages(&messageType, ns, msg.Messages, messageType2namespace)
	}
}

func main_codegen() {
	jsonFile, err := openFile(codegenSrc)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// blacklist only for types
	blacklist := mapset.NewSet[string]()
	blacklist.Add("StreamRawPredictRequest")
	blacklist.Add("UpdateDeploymentResourcePoolRequest")

	// [ToDo]: load whitelists from files; match by:
	//   - regular expressions
	//   - literal names
	pkgWhitelist := mapset.NewSet[string]()

	svcWhitelist := mapset.NewSet[string]()
	rpcWhitelist := mapset.NewSet[string]()

	namespaces := mapset.NewSet[string]()
	messageType2namespace := make(map[string]string)
	namespace2goPkg := make(map[string]string)

	// https://github.com/googleapis/google-cloud-go/blob/main/internal/gapicgen/generator/genproto.go#L56-L140

	namespaces.Add("emptypb")
	messageType2namespace["google.protobuf.Empty"] = "emptypb"
	namespace2goPkg["emptypb"] = "google.golang.org/protobuf/types/known/emptypb"

	namespaces.Add("httpbodypb")
	messageType2namespace["google.api.HttpBody"] = "httpbodypb"
	namespace2goPkg["httpbodypb"] = "google.golang.org/genproto/googleapis/api/httpbody"

	namespaces.Add("metricpb")
	messageType2namespace["google.api.MetricDescriptor"] = "metricpb"
	namespace2goPkg["metricpb"] = "google.golang.org/genproto/googleapis/api/metric"

	namespaces.Add("monitoredrespb")
	messageType2namespace["google.api.MonitoredResourceDescriptor"] = "monitoredrespb"
	namespace2goPkg["monitoredrespb"] = "google.golang.org/genproto/googleapis/api/monitoredres"

	namespaces.Add("serviceconfigpb")
	namespace2goPkg["serviceconfigpb"] = "google.golang.org/genproto/googleapis/api/serviceconfig"
	messageType2namespace["google.api.Service"] = "serviceconfigpb"

	namespaces.Add("google_longrunning_longrunningpb")
	messageType2namespace["google.longrunning.Operation"] = "google_longrunning_longrunningpb"
	namespace2goPkg["google_longrunning_longrunningpb"] = "cloud.google.com/go/longrunning/autogen/longrunningpb"

	namespaces.Add("oslogin_common_commonpb")
	messageType2namespace["google.cloud.oslogin.common.SshPublicKey"] = "oslogin_common_commonpb"
	namespace2goPkg["oslogin_common_commonpb"] = "cloud.google.com/go/oslogin/common/commonpb"

	protos := make([]*ProtoDef, 0)
	d := json.NewDecoder(jsonFile)
	protosCount := 0
	for {
		var v ProtoDef
		err := d.Decode(&v)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(2)
		}

		if strings.HasPrefix(*v.Package, "google.ads.") ||
			strings.HasPrefix(*v.Package, "google.analytics.") ||
			len(v.Services)+len(v.Messages) == 0 {
			// skipping:
			//   - google ads/analytics stuff (sorry...)
			//   - protos with 0 side effects
			continue
		}

		protosCount += 1

		namespace := strings.ReplaceAll(*v.Package, ".", "_") + "_" + *v.Go.Package
		namespaces.Add(namespace)

		v.Go.Namespace = &namespace

		if !v.Go.Migrated && v.Go.LegacyPackageFull != nil {
			namespace2goPkg[namespace] = *v.Go.LegacyPackageFull
		} else {
			namespace2goPkg[namespace] = *v.Go.PackageFull
		}

		protos = append(protos, &v)
	}

	for _, proto := range protos {
		ns := *proto.Go.Namespace

		for _, svc := range proto.Services {
			for _, rpc := range svc.RPCs {

				requestType := *rpc.Request.Type
				responseType := *rpc.Response.Type

				if !strings.HasPrefix(requestType, "google.") {
					messageType2namespace[*proto.Package+"."+requestType] = ns
				} else if strings.HasPrefix(requestType, *proto.Package) {
					messageType2namespace[requestType] = ns
				}

				if !strings.HasPrefix(responseType, "google.") {
					messageType2namespace[*proto.Package+"."+responseType] = ns
				} else if strings.HasPrefix(responseType, *proto.Package) {
					messageType2namespace[responseType] = ns
				}
			}
		}

		processMessages(proto.Package, &ns, proto.Messages, messageType2namespace)
	}

	for _, proto := range protos {

		imports := make(map[string]string)
		rpcMethod2rpcReq := make(map[string]string)
		rpcMethod2rpcRes := make(map[string]string)
		rpcMethod2rpcReqGoPkg := make(map[string]string)
		rpcMethod2rpcResGoPkg := make(map[string]string)
		service2host := make(map[string]string)

		if len(proto.Services) == 0 {
			continue
		}

		// fmt.Fprintln(os.Stderr, *v.Package)
		if !pkgWhitelist.IsEmpty() && !pkgWhitelist.Contains(*proto.Package) {
			continue
		}

		goPkgNamespace := *proto.Go.Namespace

		if !proto.Go.Migrated && proto.Go.LegacyPackageFull != nil {
			imports[goPkgNamespace] = *proto.Go.LegacyPackageFull
		} else {
			imports[goPkgNamespace] = *proto.Go.PackageFull
		}

		for _, svc := range proto.Services {
			if !svcWhitelist.IsEmpty() && !svcWhitelist.Contains(*svc.Service) {
				continue
			}

			pkgAndSvc := *proto.Package + "." + *svc.Service

			if svc.Host != nil {
				// yes... this can happen:
				//   - sample: https://github.com/googleapis/googleapis/blob/master/google/firebase/fcm/connection/v1alpha1/connection_api.proto
				service2host[pkgAndSvc] = *svc.Host
			}

			for _, rpc := range svc.RPCs {
				if !rpcWhitelist.IsEmpty() && !rpcWhitelist.Contains(*rpc.Method) {
					continue
				}

				requestType := *rpc.Request.Type
				responseType := *rpc.Response.Type

				if blacklist.ContainsAny(requestType, responseType) {
					continue
				}

				fullMethodName := pkgAndSvc + "/" + *rpc.Method

				if strings.HasPrefix(requestType, "google.") {
					namespace := messageType2namespace[requestType]
					requestTypeParts := strings.Split(requestType, ".")
					rpcMethod2rpcReq[fullMethodName] = requestTypeParts[len(requestTypeParts)-1]
					rpcMethod2rpcReqGoPkg[fullMethodName] = namespace
					if goPkg, ok := namespace2goPkg[namespace]; ok {
						imports[namespace] = goPkg
					}
				} else {
					rpcMethod2rpcReq[fullMethodName] = requestType
					rpcMethod2rpcReqGoPkg[fullMethodName] = goPkgNamespace
				}

				if strings.HasPrefix(responseType, "google.") {
					namespace := messageType2namespace[responseType]
					responseTypeParts := strings.Split(responseType, ".")
					rpcMethod2rpcRes[fullMethodName] = responseTypeParts[len(responseTypeParts)-1]
					rpcMethod2rpcResGoPkg[fullMethodName] = namespace
					if goPkg, ok := namespace2goPkg[namespace]; ok {
						imports[namespace] = goPkg
					}
				} else {
					rpcMethod2rpcRes[fullMethodName] = *rpc.Response.Type
					rpcMethod2rpcResGoPkg[fullMethodName] = goPkgNamespace
				}

			}
		}

		goNamespaceParts := strings.Split(*proto.Go.Namespace, "_")
		sizeOfGoNamespaceParts := len(goNamespaceParts)
		goFileName := strings.Join(goNamespaceParts[:sizeOfGoNamespaceParts-1], ".") + ".go"

		MustParse(tmpl).CreateFile(
			*codegenDst+"/"+goFileName,
			map[string]any{
				"imports":     imports,
				"svcHost":     service2host,
				"rpcReq":      rpcMethod2rpcReq,
				"rpcRes":      rpcMethod2rpcRes,
				"rpcReqGoPkg": rpcMethod2rpcReqGoPkg,
				"rpcResGoPkg": rpcMethod2rpcResGoPkg,
			},
		)
	}
}
