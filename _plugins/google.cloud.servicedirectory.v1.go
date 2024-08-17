// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_servicedirectory_v1_servicedirectorypb "cloud.google.com/go/servicedirectory/apiv1/servicedirectorypb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.servicedirectory.v1.LookupService", "servicedirectory.googleapis.com")
	method2RequestType.Store("google.cloud.servicedirectory.v1.LookupService/ResolveService", reflect.TypeOf((*google_cloud_servicedirectory_v1_servicedirectorypb.ResolveServiceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.servicedirectory.v1.LookupService/ResolveService", reflect.TypeOf((*google_cloud_servicedirectory_v1_servicedirectorypb.ResolveServiceResponse)(nil)).Elem())
}
