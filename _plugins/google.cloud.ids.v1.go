// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_ids_v1_idspb "cloud.google.com/go/ids/apiv1/idspb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.ids.v1.IDS", "ids.googleapis.com")
	method2RequestType.Store("google.cloud.ids.v1.IDS/CreateEndpoint", reflect.TypeOf((*google_cloud_ids_v1_idspb.CreateEndpointRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.ids.v1.IDS/CreateEndpoint", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.ids.v1.IDS/DeleteEndpoint", reflect.TypeOf((*google_cloud_ids_v1_idspb.DeleteEndpointRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.ids.v1.IDS/DeleteEndpoint", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.ids.v1.IDS/GetEndpoint", reflect.TypeOf((*google_cloud_ids_v1_idspb.GetEndpointRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.ids.v1.IDS/GetEndpoint", reflect.TypeOf((*google_cloud_ids_v1_idspb.Endpoint)(nil)).Elem())
	method2RequestType.Store("google.cloud.ids.v1.IDS/ListEndpoints", reflect.TypeOf((*google_cloud_ids_v1_idspb.ListEndpointsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.ids.v1.IDS/ListEndpoints", reflect.TypeOf((*google_cloud_ids_v1_idspb.ListEndpointsResponse)(nil)).Elem())
}
