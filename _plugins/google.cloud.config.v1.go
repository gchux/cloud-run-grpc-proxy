// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_config_v1_configpb "cloud.google.com/go/config/apiv1/configpb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.config.v1.Config", "config.googleapis.com")
	method2RequestType.Store("google.cloud.config.v1.Config/CreateDeployment", reflect.TypeOf((*google_cloud_config_v1_configpb.CreateDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/CreateDeployment", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/CreatePreview", reflect.TypeOf((*google_cloud_config_v1_configpb.CreatePreviewRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/CreatePreview", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/DeleteDeployment", reflect.TypeOf((*google_cloud_config_v1_configpb.DeleteDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/DeleteDeployment", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/DeletePreview", reflect.TypeOf((*google_cloud_config_v1_configpb.DeletePreviewRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/DeletePreview", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/DeleteStatefile", reflect.TypeOf((*google_cloud_config_v1_configpb.DeleteStatefileRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/DeleteStatefile", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/ExportDeploymentStatefile", reflect.TypeOf((*google_cloud_config_v1_configpb.ExportDeploymentStatefileRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/ExportDeploymentStatefile", reflect.TypeOf((*google_cloud_config_v1_configpb.Statefile)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/ExportLockInfo", reflect.TypeOf((*google_cloud_config_v1_configpb.ExportLockInfoRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/ExportLockInfo", reflect.TypeOf((*google_cloud_config_v1_configpb.LockInfo)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/ExportPreviewResult", reflect.TypeOf((*google_cloud_config_v1_configpb.ExportPreviewResultRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/ExportPreviewResult", reflect.TypeOf((*google_cloud_config_v1_configpb.ExportPreviewResultResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/ExportRevisionStatefile", reflect.TypeOf((*google_cloud_config_v1_configpb.ExportRevisionStatefileRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/ExportRevisionStatefile", reflect.TypeOf((*google_cloud_config_v1_configpb.Statefile)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/GetDeployment", reflect.TypeOf((*google_cloud_config_v1_configpb.GetDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/GetDeployment", reflect.TypeOf((*google_cloud_config_v1_configpb.Deployment)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/GetPreview", reflect.TypeOf((*google_cloud_config_v1_configpb.GetPreviewRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/GetPreview", reflect.TypeOf((*google_cloud_config_v1_configpb.Preview)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/GetResource", reflect.TypeOf((*google_cloud_config_v1_configpb.GetResourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/GetResource", reflect.TypeOf((*google_cloud_config_v1_configpb.Resource)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/GetRevision", reflect.TypeOf((*google_cloud_config_v1_configpb.GetRevisionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/GetRevision", reflect.TypeOf((*google_cloud_config_v1_configpb.Revision)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/GetTerraformVersion", reflect.TypeOf((*google_cloud_config_v1_configpb.GetTerraformVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/GetTerraformVersion", reflect.TypeOf((*google_cloud_config_v1_configpb.TerraformVersion)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/ImportStatefile", reflect.TypeOf((*google_cloud_config_v1_configpb.ImportStatefileRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/ImportStatefile", reflect.TypeOf((*google_cloud_config_v1_configpb.Statefile)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/ListDeployments", reflect.TypeOf((*google_cloud_config_v1_configpb.ListDeploymentsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/ListDeployments", reflect.TypeOf((*google_cloud_config_v1_configpb.ListDeploymentsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/ListPreviews", reflect.TypeOf((*google_cloud_config_v1_configpb.ListPreviewsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/ListPreviews", reflect.TypeOf((*google_cloud_config_v1_configpb.ListPreviewsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/ListResources", reflect.TypeOf((*google_cloud_config_v1_configpb.ListResourcesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/ListResources", reflect.TypeOf((*google_cloud_config_v1_configpb.ListResourcesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/ListRevisions", reflect.TypeOf((*google_cloud_config_v1_configpb.ListRevisionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/ListRevisions", reflect.TypeOf((*google_cloud_config_v1_configpb.ListRevisionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/ListTerraformVersions", reflect.TypeOf((*google_cloud_config_v1_configpb.ListTerraformVersionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/ListTerraformVersions", reflect.TypeOf((*google_cloud_config_v1_configpb.ListTerraformVersionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/LockDeployment", reflect.TypeOf((*google_cloud_config_v1_configpb.LockDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/LockDeployment", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/UnlockDeployment", reflect.TypeOf((*google_cloud_config_v1_configpb.UnlockDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/UnlockDeployment", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.config.v1.Config/UpdateDeployment", reflect.TypeOf((*google_cloud_config_v1_configpb.UpdateDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.config.v1.Config/UpdateDeployment", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
