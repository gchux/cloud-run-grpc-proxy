// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_vmwareengine_v1_vmwareenginepb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.vmwareengine.v1.VmwareEngine", "vmwareengine.googleapis.com")
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateCluster", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.CreateClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateExternalAccessRule", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.CreateExternalAccessRuleRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateExternalAccessRule", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateExternalAddress", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.CreateExternalAddressRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateExternalAddress", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateHcxActivationKey", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.CreateHcxActivationKeyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateHcxActivationKey", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateLoggingServer", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.CreateLoggingServerRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateLoggingServer", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateManagementDnsZoneBinding", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.CreateManagementDnsZoneBindingRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateManagementDnsZoneBinding", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateNetworkPeering", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.CreateNetworkPeeringRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateNetworkPeering", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateNetworkPolicy", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.CreateNetworkPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateNetworkPolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreatePrivateCloud", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.CreatePrivateCloudRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreatePrivateCloud", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreatePrivateConnection", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.CreatePrivateConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreatePrivateConnection", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateVmwareEngineNetwork", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.CreateVmwareEngineNetworkRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/CreateVmwareEngineNetwork", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeleteCluster", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.DeleteClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeleteCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeleteExternalAccessRule", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.DeleteExternalAccessRuleRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeleteExternalAccessRule", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeleteExternalAddress", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.DeleteExternalAddressRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeleteExternalAddress", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeleteLoggingServer", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.DeleteLoggingServerRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeleteLoggingServer", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeleteManagementDnsZoneBinding", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.DeleteManagementDnsZoneBindingRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeleteManagementDnsZoneBinding", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeleteNetworkPeering", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.DeleteNetworkPeeringRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeleteNetworkPeering", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeleteNetworkPolicy", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.DeleteNetworkPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeleteNetworkPolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeletePrivateCloud", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.DeletePrivateCloudRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeletePrivateCloud", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeletePrivateConnection", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.DeletePrivateConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeletePrivateConnection", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeleteVmwareEngineNetwork", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.DeleteVmwareEngineNetworkRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/DeleteVmwareEngineNetwork", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetCluster", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.GetClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetCluster", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.Cluster)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetDnsBindPermission", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.GetDnsBindPermissionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetDnsBindPermission", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.DnsBindPermission)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetDnsForwarding", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.GetDnsForwardingRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetDnsForwarding", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.DnsForwarding)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetExternalAccessRule", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.GetExternalAccessRuleRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetExternalAccessRule", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ExternalAccessRule)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetExternalAddress", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.GetExternalAddressRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetExternalAddress", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ExternalAddress)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetHcxActivationKey", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.GetHcxActivationKeyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetHcxActivationKey", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.HcxActivationKey)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetLoggingServer", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.GetLoggingServerRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetLoggingServer", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.LoggingServer)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetManagementDnsZoneBinding", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.GetManagementDnsZoneBindingRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetManagementDnsZoneBinding", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ManagementDnsZoneBinding)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetNetworkPeering", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.GetNetworkPeeringRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetNetworkPeering", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.NetworkPeering)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetNetworkPolicy", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.GetNetworkPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetNetworkPolicy", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.NetworkPolicy)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetNode", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.GetNodeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetNode", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.Node)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetNodeType", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.GetNodeTypeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetNodeType", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.NodeType)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetPrivateCloud", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.GetPrivateCloudRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetPrivateCloud", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.PrivateCloud)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetPrivateConnection", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.GetPrivateConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetPrivateConnection", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.PrivateConnection)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetSubnet", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.GetSubnetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetSubnet", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.Subnet)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetVmwareEngineNetwork", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.GetVmwareEngineNetworkRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GetVmwareEngineNetwork", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.VmwareEngineNetwork)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GrantDnsBindPermission", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.GrantDnsBindPermissionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/GrantDnsBindPermission", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListClusters", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListClustersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListClusters", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListClustersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListExternalAccessRules", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListExternalAccessRulesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListExternalAccessRules", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListExternalAccessRulesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListExternalAddresses", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListExternalAddressesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListExternalAddresses", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListExternalAddressesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListHcxActivationKeys", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListHcxActivationKeysRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListHcxActivationKeys", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListHcxActivationKeysResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListLoggingServers", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListLoggingServersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListLoggingServers", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListLoggingServersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListManagementDnsZoneBindings", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListManagementDnsZoneBindingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListManagementDnsZoneBindings", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListManagementDnsZoneBindingsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListNetworkPeerings", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListNetworkPeeringsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListNetworkPeerings", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListNetworkPeeringsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListNetworkPolicies", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListNetworkPoliciesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListNetworkPolicies", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListNetworkPoliciesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListNodeTypes", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListNodeTypesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListNodeTypes", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListNodeTypesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListNodes", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListNodesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListNodes", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListNodesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListPeeringRoutes", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListPeeringRoutesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListPeeringRoutes", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListPeeringRoutesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListPrivateClouds", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListPrivateCloudsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListPrivateClouds", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListPrivateCloudsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListPrivateConnections", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListPrivateConnectionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListPrivateConnections", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListPrivateConnectionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListSubnets", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListSubnetsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListSubnets", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListSubnetsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListVmwareEngineNetworks", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListVmwareEngineNetworksRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ListVmwareEngineNetworks", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ListVmwareEngineNetworksResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/RepairManagementDnsZoneBinding", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.RepairManagementDnsZoneBindingRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/RepairManagementDnsZoneBinding", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ResetNsxCredentials", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ResetNsxCredentialsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ResetNsxCredentials", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ResetVcenterCredentials", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ResetVcenterCredentialsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ResetVcenterCredentials", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/RevokeDnsBindPermission", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.RevokeDnsBindPermissionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/RevokeDnsBindPermission", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ShowNsxCredentials", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ShowNsxCredentialsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ShowNsxCredentials", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.Credentials)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ShowVcenterCredentials", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.ShowVcenterCredentialsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/ShowVcenterCredentials", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.Credentials)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UndeletePrivateCloud", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.UndeletePrivateCloudRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UndeletePrivateCloud", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateCluster", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.UpdateClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateDnsForwarding", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.UpdateDnsForwardingRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateDnsForwarding", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateExternalAccessRule", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.UpdateExternalAccessRuleRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateExternalAccessRule", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateExternalAddress", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.UpdateExternalAddressRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateExternalAddress", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateLoggingServer", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.UpdateLoggingServerRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateLoggingServer", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateManagementDnsZoneBinding", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.UpdateManagementDnsZoneBindingRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateManagementDnsZoneBinding", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateNetworkPeering", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.UpdateNetworkPeeringRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateNetworkPeering", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateNetworkPolicy", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.UpdateNetworkPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateNetworkPolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdatePrivateCloud", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.UpdatePrivateCloudRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdatePrivateCloud", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdatePrivateConnection", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.UpdatePrivateConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdatePrivateConnection", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateSubnet", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.UpdateSubnetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateSubnet", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateVmwareEngineNetwork", reflect.TypeOf((*google_cloud_vmwareengine_v1_vmwareenginepb.UpdateVmwareEngineNetworkRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vmwareengine.v1.VmwareEngine/UpdateVmwareEngineNetwork", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
