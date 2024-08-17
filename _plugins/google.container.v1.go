// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_container_v1_containerpb "cloud.google.com/go/container/apiv1/containerpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.container.v1.ClusterManager", "container.googleapis.com")
	method2RequestType.Store("google.container.v1.ClusterManager/CancelOperation", reflect.TypeOf((*google_container_v1_containerpb.CancelOperationRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/CancelOperation", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/CheckAutopilotCompatibility", reflect.TypeOf((*google_container_v1_containerpb.CheckAutopilotCompatibilityRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/CheckAutopilotCompatibility", reflect.TypeOf((*google_container_v1_containerpb.CheckAutopilotCompatibilityResponse)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/CompleteIPRotation", reflect.TypeOf((*google_container_v1_containerpb.CompleteIPRotationRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/CompleteIPRotation", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/CompleteNodePoolUpgrade", reflect.TypeOf((*google_container_v1_containerpb.CompleteNodePoolUpgradeRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/CompleteNodePoolUpgrade", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/CreateCluster", reflect.TypeOf((*google_container_v1_containerpb.CreateClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/CreateCluster", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/CreateNodePool", reflect.TypeOf((*google_container_v1_containerpb.CreateNodePoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/CreateNodePool", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/DeleteCluster", reflect.TypeOf((*google_container_v1_containerpb.DeleteClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/DeleteCluster", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/DeleteNodePool", reflect.TypeOf((*google_container_v1_containerpb.DeleteNodePoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/DeleteNodePool", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/GetCluster", reflect.TypeOf((*google_container_v1_containerpb.GetClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/GetCluster", reflect.TypeOf((*google_container_v1_containerpb.Cluster)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/GetJSONWebKeys", reflect.TypeOf((*google_container_v1_containerpb.GetJSONWebKeysRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/GetJSONWebKeys", reflect.TypeOf((*google_container_v1_containerpb.GetJSONWebKeysResponse)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/GetNodePool", reflect.TypeOf((*google_container_v1_containerpb.GetNodePoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/GetNodePool", reflect.TypeOf((*google_container_v1_containerpb.NodePool)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/GetOperation", reflect.TypeOf((*google_container_v1_containerpb.GetOperationRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/GetOperation", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/GetServerConfig", reflect.TypeOf((*google_container_v1_containerpb.GetServerConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/GetServerConfig", reflect.TypeOf((*google_container_v1_containerpb.ServerConfig)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/ListClusters", reflect.TypeOf((*google_container_v1_containerpb.ListClustersRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/ListClusters", reflect.TypeOf((*google_container_v1_containerpb.ListClustersResponse)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/ListNodePools", reflect.TypeOf((*google_container_v1_containerpb.ListNodePoolsRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/ListNodePools", reflect.TypeOf((*google_container_v1_containerpb.ListNodePoolsResponse)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/ListOperations", reflect.TypeOf((*google_container_v1_containerpb.ListOperationsRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/ListOperations", reflect.TypeOf((*google_container_v1_containerpb.ListOperationsResponse)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/ListUsableSubnetworks", reflect.TypeOf((*google_container_v1_containerpb.ListUsableSubnetworksRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/ListUsableSubnetworks", reflect.TypeOf((*google_container_v1_containerpb.ListUsableSubnetworksResponse)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/RollbackNodePoolUpgrade", reflect.TypeOf((*google_container_v1_containerpb.RollbackNodePoolUpgradeRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/RollbackNodePoolUpgrade", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/SetAddonsConfig", reflect.TypeOf((*google_container_v1_containerpb.SetAddonsConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/SetAddonsConfig", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/SetLabels", reflect.TypeOf((*google_container_v1_containerpb.SetLabelsRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/SetLabels", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/SetLegacyAbac", reflect.TypeOf((*google_container_v1_containerpb.SetLegacyAbacRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/SetLegacyAbac", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/SetLocations", reflect.TypeOf((*google_container_v1_containerpb.SetLocationsRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/SetLocations", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/SetLoggingService", reflect.TypeOf((*google_container_v1_containerpb.SetLoggingServiceRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/SetLoggingService", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/SetMaintenancePolicy", reflect.TypeOf((*google_container_v1_containerpb.SetMaintenancePolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/SetMaintenancePolicy", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/SetMasterAuth", reflect.TypeOf((*google_container_v1_containerpb.SetMasterAuthRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/SetMasterAuth", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/SetMonitoringService", reflect.TypeOf((*google_container_v1_containerpb.SetMonitoringServiceRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/SetMonitoringService", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/SetNetworkPolicy", reflect.TypeOf((*google_container_v1_containerpb.SetNetworkPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/SetNetworkPolicy", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/SetNodePoolAutoscaling", reflect.TypeOf((*google_container_v1_containerpb.SetNodePoolAutoscalingRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/SetNodePoolAutoscaling", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/SetNodePoolManagement", reflect.TypeOf((*google_container_v1_containerpb.SetNodePoolManagementRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/SetNodePoolManagement", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/SetNodePoolSize", reflect.TypeOf((*google_container_v1_containerpb.SetNodePoolSizeRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/SetNodePoolSize", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/StartIPRotation", reflect.TypeOf((*google_container_v1_containerpb.StartIPRotationRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/StartIPRotation", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/UpdateCluster", reflect.TypeOf((*google_container_v1_containerpb.UpdateClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/UpdateCluster", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/UpdateMaster", reflect.TypeOf((*google_container_v1_containerpb.UpdateMasterRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/UpdateMaster", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.container.v1.ClusterManager/UpdateNodePool", reflect.TypeOf((*google_container_v1_containerpb.UpdateNodePoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.container.v1.ClusterManager/UpdateNodePool", reflect.TypeOf((*google_container_v1_containerpb.Operation)(nil)).Elem())
}