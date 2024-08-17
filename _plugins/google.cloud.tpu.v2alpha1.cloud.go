// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_tpu_v2alpha1_cloud_tpupb "google.golang.org/genproto/googleapis/cloud/tpu/v2alpha1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.tpu.v2alpha1.Tpu", "tpu.googleapis.com")
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/CreateNode", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.CreateNodeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/CreateNode", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/CreateQueuedResource", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.CreateQueuedResourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/CreateQueuedResource", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/DeleteNode", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.DeleteNodeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/DeleteNode", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/DeleteQueuedResource", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.DeleteQueuedResourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/DeleteQueuedResource", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/GenerateServiceIdentity", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.GenerateServiceIdentityRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/GenerateServiceIdentity", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.GenerateServiceIdentityResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/GetAcceleratorType", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.GetAcceleratorTypeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/GetAcceleratorType", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.AcceleratorType)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/GetGuestAttributes", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.GetGuestAttributesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/GetGuestAttributes", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.GetGuestAttributesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/GetNode", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.GetNodeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/GetNode", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.Node)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/GetQueuedResource", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.GetQueuedResourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/GetQueuedResource", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.QueuedResource)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/GetRuntimeVersion", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.GetRuntimeVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/GetRuntimeVersion", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.RuntimeVersion)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/ListAcceleratorTypes", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.ListAcceleratorTypesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/ListAcceleratorTypes", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.ListAcceleratorTypesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/ListNodes", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.ListNodesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/ListNodes", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.ListNodesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/ListQueuedResources", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.ListQueuedResourcesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/ListQueuedResources", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.ListQueuedResourcesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/ListRuntimeVersions", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.ListRuntimeVersionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/ListRuntimeVersions", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.ListRuntimeVersionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/ResetQueuedResource", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.ResetQueuedResourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/ResetQueuedResource", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/SimulateMaintenanceEvent", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.SimulateMaintenanceEventRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/SimulateMaintenanceEvent", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/StartNode", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.StartNodeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/StartNode", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/StopNode", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.StopNodeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/StopNode", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.tpu.v2alpha1.Tpu/UpdateNode", reflect.TypeOf((*google_cloud_tpu_v2alpha1_cloud_tpupb.UpdateNodeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.tpu.v2alpha1.Tpu/UpdateNode", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
