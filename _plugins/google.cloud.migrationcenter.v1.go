// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_migrationcenter_v1_migrationcenterpb "cloud.google.com/go/migrationcenter/apiv1/migrationcenterpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.migrationcenter.v1.MigrationCenter", "migrationcenter.googleapis.com")
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/AddAssetsToGroup", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.AddAssetsToGroupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/AddAssetsToGroup", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/AggregateAssetsValues", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.AggregateAssetsValuesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/AggregateAssetsValues", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.AggregateAssetsValuesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/BatchDeleteAssets", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.BatchDeleteAssetsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/BatchDeleteAssets", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/BatchUpdateAssets", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.BatchUpdateAssetsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/BatchUpdateAssets", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.BatchUpdateAssetsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/CreateGroup", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.CreateGroupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/CreateGroup", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/CreateImportDataFile", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.CreateImportDataFileRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/CreateImportDataFile", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/CreateImportJob", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.CreateImportJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/CreateImportJob", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/CreatePreferenceSet", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.CreatePreferenceSetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/CreatePreferenceSet", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/CreateReport", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.CreateReportRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/CreateReport", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/CreateReportConfig", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.CreateReportConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/CreateReportConfig", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/CreateSource", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.CreateSourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/CreateSource", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/DeleteAsset", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.DeleteAssetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/DeleteAsset", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/DeleteGroup", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.DeleteGroupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/DeleteGroup", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/DeleteImportDataFile", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.DeleteImportDataFileRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/DeleteImportDataFile", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/DeleteImportJob", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.DeleteImportJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/DeleteImportJob", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/DeletePreferenceSet", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.DeletePreferenceSetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/DeletePreferenceSet", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/DeleteReport", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.DeleteReportRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/DeleteReport", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/DeleteReportConfig", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.DeleteReportConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/DeleteReportConfig", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/DeleteSource", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.DeleteSourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/DeleteSource", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetAsset", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.GetAssetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetAsset", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.Asset)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetErrorFrame", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.GetErrorFrameRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetErrorFrame", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ErrorFrame)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetGroup", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.GetGroupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetGroup", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.Group)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetImportDataFile", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.GetImportDataFileRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetImportDataFile", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ImportDataFile)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetImportJob", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.GetImportJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetImportJob", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ImportJob)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetPreferenceSet", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.GetPreferenceSetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetPreferenceSet", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.PreferenceSet)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetReport", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.GetReportRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetReport", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.Report)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetReportConfig", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.GetReportConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetReportConfig", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ReportConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetSettings", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.GetSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetSettings", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.Settings)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetSource", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.GetSourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/GetSource", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.Source)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListAssets", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListAssetsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListAssets", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListAssetsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListErrorFrames", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListErrorFramesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListErrorFrames", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListErrorFramesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListGroups", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListGroupsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListGroups", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListGroupsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListImportDataFiles", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListImportDataFilesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListImportDataFiles", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListImportDataFilesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListImportJobs", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListImportJobsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListImportJobs", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListImportJobsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListPreferenceSets", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListPreferenceSetsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListPreferenceSets", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListPreferenceSetsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListReportConfigs", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListReportConfigsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListReportConfigs", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListReportConfigsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListReports", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListReportsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListReports", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListReportsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListSources", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListSourcesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ListSources", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ListSourcesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/RemoveAssetsFromGroup", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.RemoveAssetsFromGroupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/RemoveAssetsFromGroup", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ReportAssetFrames", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ReportAssetFramesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ReportAssetFrames", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ReportAssetFramesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/RunImportJob", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.RunImportJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/RunImportJob", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/UpdateAsset", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.UpdateAssetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/UpdateAsset", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.Asset)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/UpdateGroup", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.UpdateGroupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/UpdateGroup", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/UpdateImportJob", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.UpdateImportJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/UpdateImportJob", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/UpdatePreferenceSet", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.UpdatePreferenceSetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/UpdatePreferenceSet", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/UpdateSettings", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.UpdateSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/UpdateSettings", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/UpdateSource", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.UpdateSourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/UpdateSource", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ValidateImportJob", reflect.TypeOf((*google_cloud_migrationcenter_v1_migrationcenterpb.ValidateImportJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.migrationcenter.v1.MigrationCenter/ValidateImportJob", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}