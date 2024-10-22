// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_visionai_v1_platformpb "google.golang.org/genproto/googleapis/cloud/visionai/v1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.visionai.v1.AppPlatform", "visionai.googleapis.com")
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/AddApplicationStreamInput", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.AddApplicationStreamInputRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/AddApplicationStreamInput", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/CreateApplication", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.CreateApplicationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/CreateApplication", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/CreateApplicationInstances", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.CreateApplicationInstancesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/CreateApplicationInstances", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/CreateDraft", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.CreateDraftRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/CreateDraft", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/CreateProcessor", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.CreateProcessorRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/CreateProcessor", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/DeleteApplication", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.DeleteApplicationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/DeleteApplication", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/DeleteApplicationInstances", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.DeleteApplicationInstancesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/DeleteApplicationInstances", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/DeleteDraft", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.DeleteDraftRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/DeleteDraft", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/DeleteProcessor", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.DeleteProcessorRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/DeleteProcessor", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/DeployApplication", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.DeployApplicationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/DeployApplication", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/GetApplication", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.GetApplicationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/GetApplication", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.Application)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/GetDraft", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.GetDraftRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/GetDraft", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.Draft)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/GetInstance", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.GetInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/GetInstance", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.Instance)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/GetProcessor", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.GetProcessorRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/GetProcessor", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.Processor)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/ListApplications", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.ListApplicationsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/ListApplications", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.ListApplicationsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/ListDrafts", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.ListDraftsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/ListDrafts", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.ListDraftsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/ListInstances", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.ListInstancesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/ListInstances", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.ListInstancesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/ListPrebuiltProcessors", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.ListPrebuiltProcessorsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/ListPrebuiltProcessors", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.ListPrebuiltProcessorsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/ListProcessors", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.ListProcessorsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/ListProcessors", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.ListProcessorsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/RemoveApplicationStreamInput", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.RemoveApplicationStreamInputRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/RemoveApplicationStreamInput", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/UndeployApplication", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.UndeployApplicationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/UndeployApplication", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/UpdateApplication", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.UpdateApplicationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/UpdateApplication", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/UpdateApplicationInstances", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.UpdateApplicationInstancesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/UpdateApplicationInstances", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/UpdateApplicationStreamInput", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.UpdateApplicationStreamInputRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/UpdateApplicationStreamInput", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/UpdateDraft", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.UpdateDraftRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/UpdateDraft", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.AppPlatform/UpdateProcessor", reflect.TypeOf((*google_cloud_visionai_v1_platformpb.UpdateProcessorRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.AppPlatform/UpdateProcessor", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
