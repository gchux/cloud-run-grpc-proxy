// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_privacy_dlp_v2_dlppb "cloud.google.com/go/dlp/apiv2/dlppb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.privacy.dlp.v2.DlpService", "dlp.googleapis.com")
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/ActivateJobTrigger", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ActivateJobTriggerRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/ActivateJobTrigger", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DlpJob)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/CancelDlpJob", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.CancelDlpJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/CancelDlpJob", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/CreateConnection", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.CreateConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/CreateConnection", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.Connection)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/CreateDeidentifyTemplate", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.CreateDeidentifyTemplateRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/CreateDeidentifyTemplate", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DeidentifyTemplate)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/CreateDiscoveryConfig", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.CreateDiscoveryConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/CreateDiscoveryConfig", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DiscoveryConfig)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/CreateDlpJob", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.CreateDlpJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/CreateDlpJob", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DlpJob)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/CreateInspectTemplate", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.CreateInspectTemplateRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/CreateInspectTemplate", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.InspectTemplate)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/CreateJobTrigger", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.CreateJobTriggerRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/CreateJobTrigger", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.JobTrigger)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/CreateStoredInfoType", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.CreateStoredInfoTypeRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/CreateStoredInfoType", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.StoredInfoType)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/DeidentifyContent", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DeidentifyContentRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/DeidentifyContent", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DeidentifyContentResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/DeleteConnection", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DeleteConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/DeleteConnection", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/DeleteDeidentifyTemplate", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DeleteDeidentifyTemplateRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/DeleteDeidentifyTemplate", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/DeleteDiscoveryConfig", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DeleteDiscoveryConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/DeleteDiscoveryConfig", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/DeleteDlpJob", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DeleteDlpJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/DeleteDlpJob", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/DeleteFileStoreDataProfile", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DeleteFileStoreDataProfileRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/DeleteFileStoreDataProfile", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/DeleteInspectTemplate", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DeleteInspectTemplateRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/DeleteInspectTemplate", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/DeleteJobTrigger", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DeleteJobTriggerRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/DeleteJobTrigger", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/DeleteStoredInfoType", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DeleteStoredInfoTypeRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/DeleteStoredInfoType", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/DeleteTableDataProfile", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DeleteTableDataProfileRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/DeleteTableDataProfile", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/FinishDlpJob", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.FinishDlpJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/FinishDlpJob", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/GetColumnDataProfile", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.GetColumnDataProfileRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/GetColumnDataProfile", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ColumnDataProfile)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/GetConnection", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.GetConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/GetConnection", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.Connection)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/GetDeidentifyTemplate", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.GetDeidentifyTemplateRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/GetDeidentifyTemplate", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DeidentifyTemplate)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/GetDiscoveryConfig", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.GetDiscoveryConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/GetDiscoveryConfig", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DiscoveryConfig)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/GetDlpJob", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.GetDlpJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/GetDlpJob", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DlpJob)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/GetFileStoreDataProfile", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.GetFileStoreDataProfileRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/GetFileStoreDataProfile", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.FileStoreDataProfile)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/GetInspectTemplate", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.GetInspectTemplateRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/GetInspectTemplate", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.InspectTemplate)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/GetJobTrigger", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.GetJobTriggerRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/GetJobTrigger", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.JobTrigger)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/GetProjectDataProfile", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.GetProjectDataProfileRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/GetProjectDataProfile", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ProjectDataProfile)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/GetStoredInfoType", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.GetStoredInfoTypeRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/GetStoredInfoType", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.StoredInfoType)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/GetTableDataProfile", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.GetTableDataProfileRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/GetTableDataProfile", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.TableDataProfile)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/HybridInspectDlpJob", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.HybridInspectDlpJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/HybridInspectDlpJob", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.HybridInspectResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/HybridInspectJobTrigger", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.HybridInspectJobTriggerRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/HybridInspectJobTrigger", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.HybridInspectResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/InspectContent", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.InspectContentRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/InspectContent", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.InspectContentResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/ListColumnDataProfiles", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListColumnDataProfilesRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/ListColumnDataProfiles", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListColumnDataProfilesResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/ListConnections", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListConnectionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/ListConnections", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListConnectionsResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/ListDeidentifyTemplates", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListDeidentifyTemplatesRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/ListDeidentifyTemplates", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListDeidentifyTemplatesResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/ListDiscoveryConfigs", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListDiscoveryConfigsRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/ListDiscoveryConfigs", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListDiscoveryConfigsResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/ListDlpJobs", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListDlpJobsRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/ListDlpJobs", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListDlpJobsResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/ListFileStoreDataProfiles", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListFileStoreDataProfilesRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/ListFileStoreDataProfiles", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListFileStoreDataProfilesResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/ListInfoTypes", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListInfoTypesRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/ListInfoTypes", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListInfoTypesResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/ListInspectTemplates", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListInspectTemplatesRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/ListInspectTemplates", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListInspectTemplatesResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/ListJobTriggers", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListJobTriggersRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/ListJobTriggers", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListJobTriggersResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/ListProjectDataProfiles", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListProjectDataProfilesRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/ListProjectDataProfiles", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListProjectDataProfilesResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/ListStoredInfoTypes", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListStoredInfoTypesRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/ListStoredInfoTypes", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListStoredInfoTypesResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/ListTableDataProfiles", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListTableDataProfilesRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/ListTableDataProfiles", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ListTableDataProfilesResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/RedactImage", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.RedactImageRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/RedactImage", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.RedactImageResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/ReidentifyContent", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ReidentifyContentRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/ReidentifyContent", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.ReidentifyContentResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/SearchConnections", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.SearchConnectionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/SearchConnections", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.SearchConnectionsResponse)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/UpdateConnection", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.UpdateConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/UpdateConnection", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.Connection)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/UpdateDeidentifyTemplate", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.UpdateDeidentifyTemplateRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/UpdateDeidentifyTemplate", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DeidentifyTemplate)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/UpdateDiscoveryConfig", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.UpdateDiscoveryConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/UpdateDiscoveryConfig", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.DiscoveryConfig)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/UpdateInspectTemplate", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.UpdateInspectTemplateRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/UpdateInspectTemplate", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.InspectTemplate)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/UpdateJobTrigger", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.UpdateJobTriggerRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/UpdateJobTrigger", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.JobTrigger)(nil)).Elem())
	method2RequestType.Store("google.privacy.dlp.v2.DlpService/UpdateStoredInfoType", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.UpdateStoredInfoTypeRequest)(nil)).Elem())
	method2ResponseType.Store("google.privacy.dlp.v2.DlpService/UpdateStoredInfoType", reflect.TypeOf((*google_privacy_dlp_v2_dlppb.StoredInfoType)(nil)).Elem())
}