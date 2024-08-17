// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_managedidentities_v1_managedidentitiespb "cloud.google.com/go/managedidentities/apiv1/managedidentitiespb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService", "managedidentities.googleapis.com")
	method2RequestType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/AttachTrust", reflect.TypeOf((*google_cloud_managedidentities_v1_managedidentitiespb.AttachTrustRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/AttachTrust", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/CreateMicrosoftAdDomain", reflect.TypeOf((*google_cloud_managedidentities_v1_managedidentitiespb.CreateMicrosoftAdDomainRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/CreateMicrosoftAdDomain", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/DeleteDomain", reflect.TypeOf((*google_cloud_managedidentities_v1_managedidentitiespb.DeleteDomainRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/DeleteDomain", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/DetachTrust", reflect.TypeOf((*google_cloud_managedidentities_v1_managedidentitiespb.DetachTrustRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/DetachTrust", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/GetDomain", reflect.TypeOf((*google_cloud_managedidentities_v1_managedidentitiespb.GetDomainRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/GetDomain", reflect.TypeOf((*google_cloud_managedidentities_v1_managedidentitiespb.Domain)(nil)).Elem())
	method2RequestType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/ListDomains", reflect.TypeOf((*google_cloud_managedidentities_v1_managedidentitiespb.ListDomainsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/ListDomains", reflect.TypeOf((*google_cloud_managedidentities_v1_managedidentitiespb.ListDomainsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/ReconfigureTrust", reflect.TypeOf((*google_cloud_managedidentities_v1_managedidentitiespb.ReconfigureTrustRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/ReconfigureTrust", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/ResetAdminPassword", reflect.TypeOf((*google_cloud_managedidentities_v1_managedidentitiespb.ResetAdminPasswordRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/ResetAdminPassword", reflect.TypeOf((*google_cloud_managedidentities_v1_managedidentitiespb.ResetAdminPasswordResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/UpdateDomain", reflect.TypeOf((*google_cloud_managedidentities_v1_managedidentitiespb.UpdateDomainRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/UpdateDomain", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/ValidateTrust", reflect.TypeOf((*google_cloud_managedidentities_v1_managedidentitiespb.ValidateTrustRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.managedidentities.v1.ManagedIdentitiesService/ValidateTrust", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}