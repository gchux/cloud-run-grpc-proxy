// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_workstations_v1beta_workstationspb "cloud.google.com/go/workstations/apiv1beta/workstationspb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.workstations.v1beta.Workstations", "workstations.googleapis.com")
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/CreateWorkstation", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.CreateWorkstationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/CreateWorkstation", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/CreateWorkstationCluster", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.CreateWorkstationClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/CreateWorkstationCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/CreateWorkstationConfig", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.CreateWorkstationConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/CreateWorkstationConfig", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/DeleteWorkstation", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.DeleteWorkstationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/DeleteWorkstation", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/DeleteWorkstationCluster", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.DeleteWorkstationClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/DeleteWorkstationCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/DeleteWorkstationConfig", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.DeleteWorkstationConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/DeleteWorkstationConfig", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/GenerateAccessToken", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.GenerateAccessTokenRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/GenerateAccessToken", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.GenerateAccessTokenResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/GetWorkstation", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.GetWorkstationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/GetWorkstation", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.Workstation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/GetWorkstationCluster", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.GetWorkstationClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/GetWorkstationCluster", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.WorkstationCluster)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/GetWorkstationConfig", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.GetWorkstationConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/GetWorkstationConfig", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.WorkstationConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/ListUsableWorkstationConfigs", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.ListUsableWorkstationConfigsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/ListUsableWorkstationConfigs", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.ListUsableWorkstationConfigsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/ListUsableWorkstations", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.ListUsableWorkstationsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/ListUsableWorkstations", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.ListUsableWorkstationsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/ListWorkstationClusters", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.ListWorkstationClustersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/ListWorkstationClusters", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.ListWorkstationClustersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/ListWorkstationConfigs", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.ListWorkstationConfigsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/ListWorkstationConfigs", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.ListWorkstationConfigsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/ListWorkstations", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.ListWorkstationsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/ListWorkstations", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.ListWorkstationsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/StartWorkstation", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.StartWorkstationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/StartWorkstation", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/StopWorkstation", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.StopWorkstationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/StopWorkstation", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/UpdateWorkstation", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.UpdateWorkstationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/UpdateWorkstation", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/UpdateWorkstationCluster", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.UpdateWorkstationClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/UpdateWorkstationCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workstations.v1beta.Workstations/UpdateWorkstationConfig", reflect.TypeOf((*google_cloud_workstations_v1beta_workstationspb.UpdateWorkstationConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workstations.v1beta.Workstations/UpdateWorkstationConfig", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}