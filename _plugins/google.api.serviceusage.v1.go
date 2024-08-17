// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_api_serviceusage_v1_serviceusagepb "cloud.google.com/go/serviceusage/apiv1/serviceusagepb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.api.serviceusage.v1.ServiceUsage", "serviceusage.googleapis.com")
	method2RequestType.Store("google.api.serviceusage.v1.ServiceUsage/BatchEnableServices", reflect.TypeOf((*google_api_serviceusage_v1_serviceusagepb.BatchEnableServicesRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1.ServiceUsage/BatchEnableServices", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1.ServiceUsage/BatchGetServices", reflect.TypeOf((*google_api_serviceusage_v1_serviceusagepb.BatchGetServicesRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1.ServiceUsage/BatchGetServices", reflect.TypeOf((*google_api_serviceusage_v1_serviceusagepb.BatchGetServicesResponse)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1.ServiceUsage/DisableService", reflect.TypeOf((*google_api_serviceusage_v1_serviceusagepb.DisableServiceRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1.ServiceUsage/DisableService", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1.ServiceUsage/EnableService", reflect.TypeOf((*google_api_serviceusage_v1_serviceusagepb.EnableServiceRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1.ServiceUsage/EnableService", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1.ServiceUsage/GetService", reflect.TypeOf((*google_api_serviceusage_v1_serviceusagepb.GetServiceRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1.ServiceUsage/GetService", reflect.TypeOf((*google_api_serviceusage_v1_serviceusagepb.Service)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1.ServiceUsage/ListServices", reflect.TypeOf((*google_api_serviceusage_v1_serviceusagepb.ListServicesRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1.ServiceUsage/ListServices", reflect.TypeOf((*google_api_serviceusage_v1_serviceusagepb.ListServicesResponse)(nil)).Elem())
}