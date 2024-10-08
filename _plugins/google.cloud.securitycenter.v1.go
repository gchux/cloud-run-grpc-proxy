// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_iam_v1_iampb "cloud.google.com/go/iam/apiv1/iampb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_securitycenter_v1_securitycenterpb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.securitycenter.v1.SecurityCenter", "securitycenter.googleapis.com")
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/BatchCreateResourceValueConfigs", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.BatchCreateResourceValueConfigsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/BatchCreateResourceValueConfigs", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.BatchCreateResourceValueConfigsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/BulkMuteFindings", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.BulkMuteFindingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/BulkMuteFindings", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/CreateBigQueryExport", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.CreateBigQueryExportRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/CreateBigQueryExport", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.BigQueryExport)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/CreateFinding", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.CreateFindingRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/CreateFinding", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.Finding)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/CreateMuteConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.CreateMuteConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/CreateMuteConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.MuteConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/CreateNotificationConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.CreateNotificationConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/CreateNotificationConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.NotificationConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/CreateSource", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.CreateSourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/CreateSource", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.Source)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/DeleteBigQueryExport", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.DeleteBigQueryExportRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/DeleteBigQueryExport", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/DeleteMuteConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.DeleteMuteConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/DeleteMuteConfig", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/DeleteNotificationConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.DeleteNotificationConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/DeleteNotificationConfig", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/DeleteResourceValueConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.DeleteResourceValueConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/DeleteResourceValueConfig", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetBigQueryExport", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.GetBigQueryExportRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetBigQueryExport", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.BigQueryExport)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.GetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetMuteConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.GetMuteConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetMuteConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.MuteConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetNotificationConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.GetNotificationConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetNotificationConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.NotificationConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetOrganizationSettings", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.GetOrganizationSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetOrganizationSettings", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.OrganizationSettings)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetResourceValueConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.GetResourceValueConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetResourceValueConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ResourceValueConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetSimulation", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.GetSimulationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetSimulation", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.Simulation)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetSource", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.GetSourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetSource", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.Source)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetValuedResource", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.GetValuedResourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/GetValuedResource", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ValuedResource)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/GroupAssets", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.GroupAssetsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/GroupAssets", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.GroupAssetsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/GroupFindings", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.GroupFindingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/GroupFindings", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.GroupFindingsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListAssets", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListAssetsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListAssets", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListAssetsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListAttackPaths", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListAttackPathsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListAttackPaths", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListAttackPathsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListBigQueryExports", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListBigQueryExportsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListBigQueryExports", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListBigQueryExportsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListFindings", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListFindingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListFindings", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListFindingsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListMuteConfigs", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListMuteConfigsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListMuteConfigs", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListMuteConfigsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListNotificationConfigs", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListNotificationConfigsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListNotificationConfigs", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListNotificationConfigsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListResourceValueConfigs", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListResourceValueConfigsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListResourceValueConfigs", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListResourceValueConfigsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListSources", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListSourcesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListSources", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListSourcesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListValuedResources", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListValuedResourcesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/ListValuedResources", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ListValuedResourcesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/RunAssetDiscovery", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.RunAssetDiscoveryRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/RunAssetDiscovery", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/SetFindingState", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.SetFindingStateRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/SetFindingState", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.Finding)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.SetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/SetMute", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.SetMuteRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/SetMute", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.Finding)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateBigQueryExport", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.UpdateBigQueryExportRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateBigQueryExport", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.BigQueryExport)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateExternalSystem", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.UpdateExternalSystemRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateExternalSystem", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ExternalSystem)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateFinding", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.UpdateFindingRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateFinding", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.Finding)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateMuteConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.UpdateMuteConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateMuteConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.MuteConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateNotificationConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.UpdateNotificationConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateNotificationConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.NotificationConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateOrganizationSettings", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.UpdateOrganizationSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateOrganizationSettings", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.OrganizationSettings)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateResourceValueConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.UpdateResourceValueConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateResourceValueConfig", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.ResourceValueConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateSecurityMarks", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.UpdateSecurityMarksRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateSecurityMarks", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.SecurityMarks)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateSource", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.UpdateSourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.v1.SecurityCenter/UpdateSource", reflect.TypeOf((*google_cloud_securitycenter_v1_securitycenterpb.Source)(nil)).Elem())
}
