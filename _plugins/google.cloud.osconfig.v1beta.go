// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_osconfig_v1beta_osconfigpb "cloud.google.com/go/osconfig/apiv1beta/osconfigpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.osconfig.v1beta.OsConfigService", "osconfig.googleapis.com")
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/CancelPatchJob", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.CancelPatchJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/CancelPatchJob", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.PatchJob)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/CreateGuestPolicy", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.CreateGuestPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/CreateGuestPolicy", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.GuestPolicy)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/CreatePatchDeployment", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.CreatePatchDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/CreatePatchDeployment", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.PatchDeployment)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/DeleteGuestPolicy", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.DeleteGuestPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/DeleteGuestPolicy", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/DeletePatchDeployment", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.DeletePatchDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/DeletePatchDeployment", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/ExecutePatchJob", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.ExecutePatchJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/ExecutePatchJob", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.PatchJob)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/GetGuestPolicy", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.GetGuestPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/GetGuestPolicy", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.GuestPolicy)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/GetPatchDeployment", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.GetPatchDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/GetPatchDeployment", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.PatchDeployment)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/GetPatchJob", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.GetPatchJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/GetPatchJob", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.PatchJob)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/ListGuestPolicies", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.ListGuestPoliciesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/ListGuestPolicies", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.ListGuestPoliciesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/ListPatchDeployments", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.ListPatchDeploymentsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/ListPatchDeployments", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.ListPatchDeploymentsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/ListPatchJobInstanceDetails", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.ListPatchJobInstanceDetailsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/ListPatchJobInstanceDetails", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.ListPatchJobInstanceDetailsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/ListPatchJobs", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.ListPatchJobsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/ListPatchJobs", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.ListPatchJobsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/LookupEffectiveGuestPolicy", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.LookupEffectiveGuestPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/LookupEffectiveGuestPolicy", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.EffectiveGuestPolicy)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/PausePatchDeployment", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.PausePatchDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/PausePatchDeployment", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.PatchDeployment)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/ResumePatchDeployment", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.ResumePatchDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/ResumePatchDeployment", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.PatchDeployment)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/UpdateGuestPolicy", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.UpdateGuestPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/UpdateGuestPolicy", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.GuestPolicy)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.v1beta.OsConfigService/UpdatePatchDeployment", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.UpdatePatchDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.v1beta.OsConfigService/UpdatePatchDeployment", reflect.TypeOf((*google_cloud_osconfig_v1beta_osconfigpb.PatchDeployment)(nil)).Elem())
}
