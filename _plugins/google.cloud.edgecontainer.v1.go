// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_edgecontainer_v1_edgecontainerpb "cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.edgecontainer.v1.EdgeContainer", "edgecontainer.googleapis.com")
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/CreateCluster", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.CreateClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/CreateCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/CreateNodePool", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.CreateNodePoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/CreateNodePool", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/CreateVpnConnection", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.CreateVpnConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/CreateVpnConnection", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/DeleteCluster", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.DeleteClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/DeleteCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/DeleteNodePool", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.DeleteNodePoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/DeleteNodePool", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/DeleteVpnConnection", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.DeleteVpnConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/DeleteVpnConnection", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/GenerateAccessToken", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.GenerateAccessTokenRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/GenerateAccessToken", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.GenerateAccessTokenResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/GenerateOfflineCredential", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.GenerateOfflineCredentialRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/GenerateOfflineCredential", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.GenerateOfflineCredentialResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/GetCluster", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.GetClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/GetCluster", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.Cluster)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/GetMachine", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.GetMachineRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/GetMachine", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.Machine)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/GetNodePool", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.GetNodePoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/GetNodePool", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.NodePool)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/GetServerConfig", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.GetServerConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/GetServerConfig", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.ServerConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/GetVpnConnection", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.GetVpnConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/GetVpnConnection", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.VpnConnection)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/ListClusters", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.ListClustersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/ListClusters", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.ListClustersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/ListMachines", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.ListMachinesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/ListMachines", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.ListMachinesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/ListNodePools", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.ListNodePoolsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/ListNodePools", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.ListNodePoolsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/ListVpnConnections", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.ListVpnConnectionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/ListVpnConnections", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.ListVpnConnectionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/UpdateCluster", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.UpdateClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/UpdateCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/UpdateNodePool", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.UpdateNodePoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/UpdateNodePool", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.edgecontainer.v1.EdgeContainer/UpgradeCluster", reflect.TypeOf((*google_cloud_edgecontainer_v1_edgecontainerpb.UpgradeClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.edgecontainer.v1.EdgeContainer/UpgradeCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
