// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_discoveryengine_v1beta_discoveryenginepb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.discoveryengine.v1beta.GroundedGenerationService", "discoveryengine.googleapis.com")
	method2RequestType.Store("google.cloud.discoveryengine.v1beta.GroundedGenerationService/CheckGrounding", reflect.TypeOf((*google_cloud_discoveryengine_v1beta_discoveryenginepb.CheckGroundingRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.discoveryengine.v1beta.GroundedGenerationService/CheckGrounding", reflect.TypeOf((*google_cloud_discoveryengine_v1beta_discoveryenginepb.CheckGroundingResponse)(nil)).Elem())
}
