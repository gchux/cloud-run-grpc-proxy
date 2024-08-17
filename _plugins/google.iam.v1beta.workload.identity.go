// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_iam_v1beta_workload_identity_poolpb "google.golang.org/genproto/googleapis/iam/v1beta"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.iam.v1beta.WorkloadIdentityPools", "iam.googleapis.com")
	method2RequestType.Store("google.iam.v1beta.WorkloadIdentityPools/CreateWorkloadIdentityPool", reflect.TypeOf((*google_iam_v1beta_workload_identity_poolpb.CreateWorkloadIdentityPoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.v1beta.WorkloadIdentityPools/CreateWorkloadIdentityPool", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.iam.v1beta.WorkloadIdentityPools/CreateWorkloadIdentityPoolProvider", reflect.TypeOf((*google_iam_v1beta_workload_identity_poolpb.CreateWorkloadIdentityPoolProviderRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.v1beta.WorkloadIdentityPools/CreateWorkloadIdentityPoolProvider", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.iam.v1beta.WorkloadIdentityPools/DeleteWorkloadIdentityPool", reflect.TypeOf((*google_iam_v1beta_workload_identity_poolpb.DeleteWorkloadIdentityPoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.v1beta.WorkloadIdentityPools/DeleteWorkloadIdentityPool", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.iam.v1beta.WorkloadIdentityPools/DeleteWorkloadIdentityPoolProvider", reflect.TypeOf((*google_iam_v1beta_workload_identity_poolpb.DeleteWorkloadIdentityPoolProviderRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.v1beta.WorkloadIdentityPools/DeleteWorkloadIdentityPoolProvider", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.iam.v1beta.WorkloadIdentityPools/GetWorkloadIdentityPool", reflect.TypeOf((*google_iam_v1beta_workload_identity_poolpb.GetWorkloadIdentityPoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.v1beta.WorkloadIdentityPools/GetWorkloadIdentityPool", reflect.TypeOf((*google_iam_v1beta_workload_identity_poolpb.WorkloadIdentityPool)(nil)).Elem())
	method2RequestType.Store("google.iam.v1beta.WorkloadIdentityPools/GetWorkloadIdentityPoolProvider", reflect.TypeOf((*google_iam_v1beta_workload_identity_poolpb.GetWorkloadIdentityPoolProviderRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.v1beta.WorkloadIdentityPools/GetWorkloadIdentityPoolProvider", reflect.TypeOf((*google_iam_v1beta_workload_identity_poolpb.WorkloadIdentityPoolProvider)(nil)).Elem())
	method2RequestType.Store("google.iam.v1beta.WorkloadIdentityPools/ListWorkloadIdentityPoolProviders", reflect.TypeOf((*google_iam_v1beta_workload_identity_poolpb.ListWorkloadIdentityPoolProvidersRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.v1beta.WorkloadIdentityPools/ListWorkloadIdentityPoolProviders", reflect.TypeOf((*google_iam_v1beta_workload_identity_poolpb.ListWorkloadIdentityPoolProvidersResponse)(nil)).Elem())
	method2RequestType.Store("google.iam.v1beta.WorkloadIdentityPools/ListWorkloadIdentityPools", reflect.TypeOf((*google_iam_v1beta_workload_identity_poolpb.ListWorkloadIdentityPoolsRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.v1beta.WorkloadIdentityPools/ListWorkloadIdentityPools", reflect.TypeOf((*google_iam_v1beta_workload_identity_poolpb.ListWorkloadIdentityPoolsResponse)(nil)).Elem())
	method2RequestType.Store("google.iam.v1beta.WorkloadIdentityPools/UndeleteWorkloadIdentityPool", reflect.TypeOf((*google_iam_v1beta_workload_identity_poolpb.UndeleteWorkloadIdentityPoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.v1beta.WorkloadIdentityPools/UndeleteWorkloadIdentityPool", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.iam.v1beta.WorkloadIdentityPools/UndeleteWorkloadIdentityPoolProvider", reflect.TypeOf((*google_iam_v1beta_workload_identity_poolpb.UndeleteWorkloadIdentityPoolProviderRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.v1beta.WorkloadIdentityPools/UndeleteWorkloadIdentityPoolProvider", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.iam.v1beta.WorkloadIdentityPools/UpdateWorkloadIdentityPool", reflect.TypeOf((*google_iam_v1beta_workload_identity_poolpb.UpdateWorkloadIdentityPoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.v1beta.WorkloadIdentityPools/UpdateWorkloadIdentityPool", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.iam.v1beta.WorkloadIdentityPools/UpdateWorkloadIdentityPoolProvider", reflect.TypeOf((*google_iam_v1beta_workload_identity_poolpb.UpdateWorkloadIdentityPoolProviderRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.v1beta.WorkloadIdentityPools/UpdateWorkloadIdentityPoolProvider", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}