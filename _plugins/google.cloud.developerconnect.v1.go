// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_developerconnect_v1_developerconnectpb "cloud.google.com/go/developerconnect/apiv1/developerconnectpb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.developerconnect.v1.DeveloperConnect", "developerconnect.googleapis.com")
	method2RequestType.Store("google.cloud.developerconnect.v1.DeveloperConnect/CreateConnection", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.CreateConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.developerconnect.v1.DeveloperConnect/CreateConnection", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.developerconnect.v1.DeveloperConnect/CreateGitRepositoryLink", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.CreateGitRepositoryLinkRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.developerconnect.v1.DeveloperConnect/CreateGitRepositoryLink", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.developerconnect.v1.DeveloperConnect/DeleteConnection", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.DeleteConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.developerconnect.v1.DeveloperConnect/DeleteConnection", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.developerconnect.v1.DeveloperConnect/DeleteGitRepositoryLink", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.DeleteGitRepositoryLinkRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.developerconnect.v1.DeveloperConnect/DeleteGitRepositoryLink", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.developerconnect.v1.DeveloperConnect/FetchGitHubInstallations", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.FetchGitHubInstallationsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.developerconnect.v1.DeveloperConnect/FetchGitHubInstallations", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.FetchGitHubInstallationsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.developerconnect.v1.DeveloperConnect/FetchGitRefs", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.FetchGitRefsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.developerconnect.v1.DeveloperConnect/FetchGitRefs", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.FetchGitRefsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.developerconnect.v1.DeveloperConnect/FetchLinkableGitRepositories", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.FetchLinkableGitRepositoriesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.developerconnect.v1.DeveloperConnect/FetchLinkableGitRepositories", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.FetchLinkableGitRepositoriesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.developerconnect.v1.DeveloperConnect/FetchReadToken", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.FetchReadTokenRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.developerconnect.v1.DeveloperConnect/FetchReadToken", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.FetchReadTokenResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.developerconnect.v1.DeveloperConnect/FetchReadWriteToken", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.FetchReadWriteTokenRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.developerconnect.v1.DeveloperConnect/FetchReadWriteToken", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.FetchReadWriteTokenResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.developerconnect.v1.DeveloperConnect/GetConnection", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.GetConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.developerconnect.v1.DeveloperConnect/GetConnection", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.Connection)(nil)).Elem())
	method2RequestType.Store("google.cloud.developerconnect.v1.DeveloperConnect/GetGitRepositoryLink", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.GetGitRepositoryLinkRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.developerconnect.v1.DeveloperConnect/GetGitRepositoryLink", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.GitRepositoryLink)(nil)).Elem())
	method2RequestType.Store("google.cloud.developerconnect.v1.DeveloperConnect/ListConnections", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.ListConnectionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.developerconnect.v1.DeveloperConnect/ListConnections", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.ListConnectionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.developerconnect.v1.DeveloperConnect/ListGitRepositoryLinks", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.ListGitRepositoryLinksRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.developerconnect.v1.DeveloperConnect/ListGitRepositoryLinks", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.ListGitRepositoryLinksResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.developerconnect.v1.DeveloperConnect/UpdateConnection", reflect.TypeOf((*google_cloud_developerconnect_v1_developerconnectpb.UpdateConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.developerconnect.v1.DeveloperConnect/UpdateConnection", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
