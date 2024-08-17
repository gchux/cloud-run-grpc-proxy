// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_gkehub_v1beta1_gkehubpb "cloud.google.com/go/gkehub/apiv1beta1/gkehubpb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.gkehub.v1beta1.GkeHubMembershipService", "gkehub.googleapis.com")
	method2RequestType.Store("google.cloud.gkehub.v1beta1.GkeHubMembershipService/CreateMembership", reflect.TypeOf((*google_cloud_gkehub_v1beta1_gkehubpb.CreateMembershipRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gkehub.v1beta1.GkeHubMembershipService/CreateMembership", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.gkehub.v1beta1.GkeHubMembershipService/DeleteMembership", reflect.TypeOf((*google_cloud_gkehub_v1beta1_gkehubpb.DeleteMembershipRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gkehub.v1beta1.GkeHubMembershipService/DeleteMembership", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.gkehub.v1beta1.GkeHubMembershipService/GenerateConnectManifest", reflect.TypeOf((*google_cloud_gkehub_v1beta1_gkehubpb.GenerateConnectManifestRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gkehub.v1beta1.GkeHubMembershipService/GenerateConnectManifest", reflect.TypeOf((*google_cloud_gkehub_v1beta1_gkehubpb.GenerateConnectManifestResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.gkehub.v1beta1.GkeHubMembershipService/GenerateExclusivityManifest", reflect.TypeOf((*google_cloud_gkehub_v1beta1_gkehubpb.GenerateExclusivityManifestRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gkehub.v1beta1.GkeHubMembershipService/GenerateExclusivityManifest", reflect.TypeOf((*google_cloud_gkehub_v1beta1_gkehubpb.GenerateExclusivityManifestResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.gkehub.v1beta1.GkeHubMembershipService/GetMembership", reflect.TypeOf((*google_cloud_gkehub_v1beta1_gkehubpb.GetMembershipRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gkehub.v1beta1.GkeHubMembershipService/GetMembership", reflect.TypeOf((*google_cloud_gkehub_v1beta1_gkehubpb.Membership)(nil)).Elem())
	method2RequestType.Store("google.cloud.gkehub.v1beta1.GkeHubMembershipService/ListMemberships", reflect.TypeOf((*google_cloud_gkehub_v1beta1_gkehubpb.ListMembershipsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gkehub.v1beta1.GkeHubMembershipService/ListMemberships", reflect.TypeOf((*google_cloud_gkehub_v1beta1_gkehubpb.ListMembershipsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.gkehub.v1beta1.GkeHubMembershipService/UpdateMembership", reflect.TypeOf((*google_cloud_gkehub_v1beta1_gkehubpb.UpdateMembershipRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gkehub.v1beta1.GkeHubMembershipService/UpdateMembership", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.gkehub.v1beta1.GkeHubMembershipService/ValidateExclusivity", reflect.TypeOf((*google_cloud_gkehub_v1beta1_gkehubpb.ValidateExclusivityRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gkehub.v1beta1.GkeHubMembershipService/ValidateExclusivity", reflect.TypeOf((*google_cloud_gkehub_v1beta1_gkehubpb.ValidateExclusivityResponse)(nil)).Elem())
}