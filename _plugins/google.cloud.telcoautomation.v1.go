// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_telcoautomation_v1_telcoautomationpb "cloud.google.com/go/telcoautomation/apiv1/telcoautomationpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.telcoautomation.v1.TelcoAutomation", "telcoautomation.googleapis.com")
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ApplyDeployment", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ApplyDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ApplyDeployment", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.Deployment)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ApplyHydratedDeployment", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ApplyHydratedDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ApplyHydratedDeployment", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.HydratedDeployment)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ApproveBlueprint", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ApproveBlueprintRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ApproveBlueprint", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.Blueprint)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ComputeDeploymentStatus", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ComputeDeploymentStatusRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ComputeDeploymentStatus", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ComputeDeploymentStatusResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/CreateBlueprint", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.CreateBlueprintRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/CreateBlueprint", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.Blueprint)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/CreateDeployment", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.CreateDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/CreateDeployment", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.Deployment)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/CreateEdgeSlm", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.CreateEdgeSlmRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/CreateEdgeSlm", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/CreateOrchestrationCluster", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.CreateOrchestrationClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/CreateOrchestrationCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/DeleteBlueprint", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.DeleteBlueprintRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/DeleteBlueprint", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/DeleteEdgeSlm", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.DeleteEdgeSlmRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/DeleteEdgeSlm", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/DeleteOrchestrationCluster", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.DeleteOrchestrationClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/DeleteOrchestrationCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/DiscardBlueprintChanges", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.DiscardBlueprintChangesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/DiscardBlueprintChanges", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.DiscardBlueprintChangesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/DiscardDeploymentChanges", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.DiscardDeploymentChangesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/DiscardDeploymentChanges", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.DiscardDeploymentChangesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/GetBlueprint", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.GetBlueprintRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/GetBlueprint", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.Blueprint)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/GetDeployment", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.GetDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/GetDeployment", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.Deployment)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/GetEdgeSlm", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.GetEdgeSlmRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/GetEdgeSlm", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.EdgeSlm)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/GetHydratedDeployment", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.GetHydratedDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/GetHydratedDeployment", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.HydratedDeployment)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/GetOrchestrationCluster", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.GetOrchestrationClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/GetOrchestrationCluster", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.OrchestrationCluster)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/GetPublicBlueprint", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.GetPublicBlueprintRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/GetPublicBlueprint", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.PublicBlueprint)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ListBlueprintRevisions", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ListBlueprintRevisionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ListBlueprintRevisions", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ListBlueprintRevisionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ListBlueprints", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ListBlueprintsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ListBlueprints", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ListBlueprintsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ListDeploymentRevisions", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ListDeploymentRevisionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ListDeploymentRevisions", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ListDeploymentRevisionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ListDeployments", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ListDeploymentsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ListDeployments", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ListDeploymentsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ListEdgeSlms", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ListEdgeSlmsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ListEdgeSlms", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ListEdgeSlmsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ListHydratedDeployments", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ListHydratedDeploymentsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ListHydratedDeployments", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ListHydratedDeploymentsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ListOrchestrationClusters", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ListOrchestrationClustersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ListOrchestrationClusters", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ListOrchestrationClustersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ListPublicBlueprints", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ListPublicBlueprintsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ListPublicBlueprints", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ListPublicBlueprintsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ProposeBlueprint", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.ProposeBlueprintRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/ProposeBlueprint", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.Blueprint)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/RejectBlueprint", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.RejectBlueprintRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/RejectBlueprint", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.Blueprint)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/RemoveDeployment", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.RemoveDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/RemoveDeployment", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/RollbackDeployment", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.RollbackDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/RollbackDeployment", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.Deployment)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/SearchBlueprintRevisions", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.SearchBlueprintRevisionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/SearchBlueprintRevisions", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.SearchBlueprintRevisionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/SearchDeploymentRevisions", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.SearchDeploymentRevisionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/SearchDeploymentRevisions", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.SearchDeploymentRevisionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/UpdateBlueprint", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.UpdateBlueprintRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/UpdateBlueprint", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.Blueprint)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/UpdateDeployment", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.UpdateDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/UpdateDeployment", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.Deployment)(nil)).Elem())
	method2RequestType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/UpdateHydratedDeployment", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.UpdateHydratedDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.telcoautomation.v1.TelcoAutomation/UpdateHydratedDeployment", reflect.TypeOf((*google_cloud_telcoautomation_v1_telcoautomationpb.HydratedDeployment)(nil)).Elem())
}
