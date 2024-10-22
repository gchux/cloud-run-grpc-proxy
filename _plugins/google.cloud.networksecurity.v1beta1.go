// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_networksecurity_v1beta1_networksecuritypb "cloud.google.com/go/networksecurity/apiv1beta1/networksecuritypb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity", "networksecurity.googleapis.com")
	method2RequestType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/CreateAuthorizationPolicy", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.CreateAuthorizationPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/CreateAuthorizationPolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/CreateClientTlsPolicy", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.CreateClientTlsPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/CreateClientTlsPolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/CreateServerTlsPolicy", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.CreateServerTlsPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/CreateServerTlsPolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/DeleteAuthorizationPolicy", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.DeleteAuthorizationPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/DeleteAuthorizationPolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/DeleteClientTlsPolicy", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.DeleteClientTlsPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/DeleteClientTlsPolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/DeleteServerTlsPolicy", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.DeleteServerTlsPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/DeleteServerTlsPolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/GetAuthorizationPolicy", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.GetAuthorizationPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/GetAuthorizationPolicy", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.AuthorizationPolicy)(nil)).Elem())
	method2RequestType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/GetClientTlsPolicy", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.GetClientTlsPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/GetClientTlsPolicy", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.ClientTlsPolicy)(nil)).Elem())
	method2RequestType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/GetServerTlsPolicy", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.GetServerTlsPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/GetServerTlsPolicy", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.ServerTlsPolicy)(nil)).Elem())
	method2RequestType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/ListAuthorizationPolicies", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.ListAuthorizationPoliciesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/ListAuthorizationPolicies", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.ListAuthorizationPoliciesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/ListClientTlsPolicies", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.ListClientTlsPoliciesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/ListClientTlsPolicies", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.ListClientTlsPoliciesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/ListServerTlsPolicies", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.ListServerTlsPoliciesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/ListServerTlsPolicies", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.ListServerTlsPoliciesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/UpdateAuthorizationPolicy", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.UpdateAuthorizationPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/UpdateAuthorizationPolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/UpdateClientTlsPolicy", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.UpdateClientTlsPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/UpdateClientTlsPolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/UpdateServerTlsPolicy", reflect.TypeOf((*google_cloud_networksecurity_v1beta1_networksecuritypb.UpdateServerTlsPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.networksecurity.v1beta1.NetworkSecurity/UpdateServerTlsPolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
