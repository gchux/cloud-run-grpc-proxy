// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_notebooks_v1_notebookspb "cloud.google.com/go/notebooks/apiv1/notebookspb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.notebooks.v1.ManagedNotebookService", "notebooks.googleapis.com")
	method2RequestType.Store("google.cloud.notebooks.v1.ManagedNotebookService/CreateRuntime", reflect.TypeOf((*google_cloud_notebooks_v1_notebookspb.CreateRuntimeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.notebooks.v1.ManagedNotebookService/CreateRuntime", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.notebooks.v1.ManagedNotebookService/DeleteRuntime", reflect.TypeOf((*google_cloud_notebooks_v1_notebookspb.DeleteRuntimeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.notebooks.v1.ManagedNotebookService/DeleteRuntime", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.notebooks.v1.ManagedNotebookService/DiagnoseRuntime", reflect.TypeOf((*google_cloud_notebooks_v1_notebookspb.DiagnoseRuntimeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.notebooks.v1.ManagedNotebookService/DiagnoseRuntime", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.notebooks.v1.ManagedNotebookService/GetRuntime", reflect.TypeOf((*google_cloud_notebooks_v1_notebookspb.GetRuntimeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.notebooks.v1.ManagedNotebookService/GetRuntime", reflect.TypeOf((*google_cloud_notebooks_v1_notebookspb.Runtime)(nil)).Elem())
	method2RequestType.Store("google.cloud.notebooks.v1.ManagedNotebookService/ListRuntimes", reflect.TypeOf((*google_cloud_notebooks_v1_notebookspb.ListRuntimesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.notebooks.v1.ManagedNotebookService/ListRuntimes", reflect.TypeOf((*google_cloud_notebooks_v1_notebookspb.ListRuntimesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.notebooks.v1.ManagedNotebookService/RefreshRuntimeTokenInternal", reflect.TypeOf((*google_cloud_notebooks_v1_notebookspb.RefreshRuntimeTokenInternalRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.notebooks.v1.ManagedNotebookService/RefreshRuntimeTokenInternal", reflect.TypeOf((*google_cloud_notebooks_v1_notebookspb.RefreshRuntimeTokenInternalResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.notebooks.v1.ManagedNotebookService/ReportRuntimeEvent", reflect.TypeOf((*google_cloud_notebooks_v1_notebookspb.ReportRuntimeEventRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.notebooks.v1.ManagedNotebookService/ReportRuntimeEvent", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.notebooks.v1.ManagedNotebookService/ResetRuntime", reflect.TypeOf((*google_cloud_notebooks_v1_notebookspb.ResetRuntimeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.notebooks.v1.ManagedNotebookService/ResetRuntime", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.notebooks.v1.ManagedNotebookService/StartRuntime", reflect.TypeOf((*google_cloud_notebooks_v1_notebookspb.StartRuntimeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.notebooks.v1.ManagedNotebookService/StartRuntime", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.notebooks.v1.ManagedNotebookService/StopRuntime", reflect.TypeOf((*google_cloud_notebooks_v1_notebookspb.StopRuntimeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.notebooks.v1.ManagedNotebookService/StopRuntime", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.notebooks.v1.ManagedNotebookService/SwitchRuntime", reflect.TypeOf((*google_cloud_notebooks_v1_notebookspb.SwitchRuntimeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.notebooks.v1.ManagedNotebookService/SwitchRuntime", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.notebooks.v1.ManagedNotebookService/UpdateRuntime", reflect.TypeOf((*google_cloud_notebooks_v1_notebookspb.UpdateRuntimeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.notebooks.v1.ManagedNotebookService/UpdateRuntime", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.notebooks.v1.ManagedNotebookService/UpgradeRuntime", reflect.TypeOf((*google_cloud_notebooks_v1_notebookspb.UpgradeRuntimeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.notebooks.v1.ManagedNotebookService/UpgradeRuntime", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
