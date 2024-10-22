// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_networkmanagement_v1_networkmanagementpb "cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.networkmanagement.v1.ReachabilityService", "networkmanagement.googleapis.com")
	method2RequestType.Store("google.cloud.networkmanagement.v1.ReachabilityService/CreateConnectivityTest", reflect.TypeOf((*google_cloud_networkmanagement_v1_networkmanagementpb.CreateConnectivityTestRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networkmanagement.v1.ReachabilityService/CreateConnectivityTest", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.networkmanagement.v1.ReachabilityService/DeleteConnectivityTest", reflect.TypeOf((*google_cloud_networkmanagement_v1_networkmanagementpb.DeleteConnectivityTestRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networkmanagement.v1.ReachabilityService/DeleteConnectivityTest", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.networkmanagement.v1.ReachabilityService/GetConnectivityTest", reflect.TypeOf((*google_cloud_networkmanagement_v1_networkmanagementpb.GetConnectivityTestRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networkmanagement.v1.ReachabilityService/GetConnectivityTest", reflect.TypeOf((*google_cloud_networkmanagement_v1_networkmanagementpb.ConnectivityTest)(nil)).Elem())
	method2RequestType.Store("google.cloud.networkmanagement.v1.ReachabilityService/ListConnectivityTests", reflect.TypeOf((*google_cloud_networkmanagement_v1_networkmanagementpb.ListConnectivityTestsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networkmanagement.v1.ReachabilityService/ListConnectivityTests", reflect.TypeOf((*google_cloud_networkmanagement_v1_networkmanagementpb.ListConnectivityTestsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.networkmanagement.v1.ReachabilityService/RerunConnectivityTest", reflect.TypeOf((*google_cloud_networkmanagement_v1_networkmanagementpb.RerunConnectivityTestRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networkmanagement.v1.ReachabilityService/RerunConnectivityTest", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.networkmanagement.v1.ReachabilityService/UpdateConnectivityTest", reflect.TypeOf((*google_cloud_networkmanagement_v1_networkmanagementpb.UpdateConnectivityTestRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networkmanagement.v1.ReachabilityService/UpdateConnectivityTest", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
