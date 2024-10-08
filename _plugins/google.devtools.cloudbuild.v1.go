// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_devtools_cloudbuild_v1_cloudbuildpb "cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.devtools.cloudbuild.v1.CloudBuild", "cloudbuild.googleapis.com")
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/ApproveBuild", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.ApproveBuildRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/ApproveBuild", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/CancelBuild", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.CancelBuildRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/CancelBuild", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.Build)(nil)).Elem())
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/CreateBuild", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.CreateBuildRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/CreateBuild", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/CreateBuildTrigger", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.CreateBuildTriggerRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/CreateBuildTrigger", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.BuildTrigger)(nil)).Elem())
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/CreateWorkerPool", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.CreateWorkerPoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/CreateWorkerPool", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/DeleteBuildTrigger", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.DeleteBuildTriggerRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/DeleteBuildTrigger", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/DeleteWorkerPool", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.DeleteWorkerPoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/DeleteWorkerPool", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/GetBuild", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.GetBuildRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/GetBuild", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.Build)(nil)).Elem())
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/GetBuildTrigger", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.GetBuildTriggerRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/GetBuildTrigger", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.BuildTrigger)(nil)).Elem())
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/GetWorkerPool", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.GetWorkerPoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/GetWorkerPool", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.WorkerPool)(nil)).Elem())
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/ListBuildTriggers", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.ListBuildTriggersRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/ListBuildTriggers", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.ListBuildTriggersResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/ListBuilds", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.ListBuildsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/ListBuilds", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.ListBuildsResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/ListWorkerPools", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.ListWorkerPoolsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/ListWorkerPools", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.ListWorkerPoolsResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/ReceiveTriggerWebhook", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.ReceiveTriggerWebhookRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/ReceiveTriggerWebhook", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.ReceiveTriggerWebhookResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/RetryBuild", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.RetryBuildRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/RetryBuild", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/RunBuildTrigger", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.RunBuildTriggerRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/RunBuildTrigger", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/UpdateBuildTrigger", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.UpdateBuildTriggerRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/UpdateBuildTrigger", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.BuildTrigger)(nil)).Elem())
	method2RequestType.Store("google.devtools.cloudbuild.v1.CloudBuild/UpdateWorkerPool", reflect.TypeOf((*google_devtools_cloudbuild_v1_cloudbuildpb.UpdateWorkerPoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.cloudbuild.v1.CloudBuild/UpdateWorkerPool", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
