// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_workstations_v1_workstationspb "cloud.google.com/go/workstations/apiv1/workstationspb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.workstations.v1.Workstations", "workstations.googleapis.com")
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/CreateWorkstation", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.CreateWorkstationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/CreateWorkstation", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/CreateWorkstationCluster", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.CreateWorkstationClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/CreateWorkstationCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/CreateWorkstationConfig", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.CreateWorkstationConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/CreateWorkstationConfig", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/DeleteWorkstation", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.DeleteWorkstationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/DeleteWorkstation", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/DeleteWorkstationCluster", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.DeleteWorkstationClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/DeleteWorkstationCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/DeleteWorkstationConfig", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.DeleteWorkstationConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/DeleteWorkstationConfig", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/GenerateAccessToken", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.GenerateAccessTokenRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/GenerateAccessToken", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.GenerateAccessTokenResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/GetWorkstation", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.GetWorkstationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/GetWorkstation", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.Workstation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/GetWorkstationCluster", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.GetWorkstationClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/GetWorkstationCluster", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.WorkstationCluster)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/GetWorkstationConfig", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.GetWorkstationConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/GetWorkstationConfig", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.WorkstationConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/ListUsableWorkstationConfigs", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.ListUsableWorkstationConfigsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/ListUsableWorkstationConfigs", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.ListUsableWorkstationConfigsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/ListUsableWorkstations", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.ListUsableWorkstationsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/ListUsableWorkstations", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.ListUsableWorkstationsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/ListWorkstationClusters", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.ListWorkstationClustersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/ListWorkstationClusters", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.ListWorkstationClustersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/ListWorkstationConfigs", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.ListWorkstationConfigsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/ListWorkstationConfigs", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.ListWorkstationConfigsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/ListWorkstations", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.ListWorkstationsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/ListWorkstations", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.ListWorkstationsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/StartWorkstation", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.StartWorkstationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/StartWorkstation", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/StopWorkstation", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.StopWorkstationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/StopWorkstation", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/UpdateWorkstation", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.UpdateWorkstationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/UpdateWorkstation", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/UpdateWorkstationCluster", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.UpdateWorkstationClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/UpdateWorkstationCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1.Workstations/UpdateWorkstationConfig", reflect.TypeOf((*google_cloud_workstations_v1_workstationspb.UpdateWorkstationConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1.Workstations/UpdateWorkstationConfig", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}