// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_datafusion_v1_datafusionpb "cloud.google.com/go/datafusion/apiv1/datafusionpb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.datafusion.v1.DataFusion", "datafusion.googleapis.com")
	method2RequestType.Store("google.cloud.datafusion.v1.DataFusion/CreateInstance", reflect.TypeOf((*google_cloud_datafusion_v1_datafusionpb.CreateInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datafusion.v1.DataFusion/CreateInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.datafusion.v1.DataFusion/DeleteInstance", reflect.TypeOf((*google_cloud_datafusion_v1_datafusionpb.DeleteInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datafusion.v1.DataFusion/DeleteInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.datafusion.v1.DataFusion/GetInstance", reflect.TypeOf((*google_cloud_datafusion_v1_datafusionpb.GetInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datafusion.v1.DataFusion/GetInstance", reflect.TypeOf((*google_cloud_datafusion_v1_datafusionpb.Instance)(nil)).Elem())
	method2RequestType.Store("google.cloud.datafusion.v1.DataFusion/ListAvailableVersions", reflect.TypeOf((*google_cloud_datafusion_v1_datafusionpb.ListAvailableVersionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datafusion.v1.DataFusion/ListAvailableVersions", reflect.TypeOf((*google_cloud_datafusion_v1_datafusionpb.ListAvailableVersionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datafusion.v1.DataFusion/ListInstances", reflect.TypeOf((*google_cloud_datafusion_v1_datafusionpb.ListInstancesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datafusion.v1.DataFusion/ListInstances", reflect.TypeOf((*google_cloud_datafusion_v1_datafusionpb.ListInstancesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datafusion.v1.DataFusion/RestartInstance", reflect.TypeOf((*google_cloud_datafusion_v1_datafusionpb.RestartInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datafusion.v1.DataFusion/RestartInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.datafusion.v1.DataFusion/UpdateInstance", reflect.TypeOf((*google_cloud_datafusion_v1_datafusionpb.UpdateInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datafusion.v1.DataFusion/UpdateInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
