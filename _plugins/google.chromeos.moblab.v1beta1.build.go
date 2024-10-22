// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_chromeos_moblab_v1beta1_build_servicepb "google.golang.org/genproto/googleapis/chromeos/moblab/v1beta1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.chromeos.moblab.v1beta1.BuildService", "chromeosmoblab.googleapis.com")
	method2RequestType.Store("google.chromeos.moblab.v1beta1.BuildService/CheckBuildStageStatus", reflect.TypeOf((*google_chromeos_moblab_v1beta1_build_servicepb.CheckBuildStageStatusRequest)(nil)).Elem())
	method2ResponseType.Store("google.chromeos.moblab.v1beta1.BuildService/CheckBuildStageStatus", reflect.TypeOf((*google_chromeos_moblab_v1beta1_build_servicepb.CheckBuildStageStatusResponse)(nil)).Elem())
	method2RequestType.Store("google.chromeos.moblab.v1beta1.BuildService/FindMostStableBuild", reflect.TypeOf((*google_chromeos_moblab_v1beta1_build_servicepb.FindMostStableBuildRequest)(nil)).Elem())
	method2ResponseType.Store("google.chromeos.moblab.v1beta1.BuildService/FindMostStableBuild", reflect.TypeOf((*google_chromeos_moblab_v1beta1_build_servicepb.FindMostStableBuildResponse)(nil)).Elem())
	method2RequestType.Store("google.chromeos.moblab.v1beta1.BuildService/ListBuildTargets", reflect.TypeOf((*google_chromeos_moblab_v1beta1_build_servicepb.ListBuildTargetsRequest)(nil)).Elem())
	method2ResponseType.Store("google.chromeos.moblab.v1beta1.BuildService/ListBuildTargets", reflect.TypeOf((*google_chromeos_moblab_v1beta1_build_servicepb.ListBuildTargetsResponse)(nil)).Elem())
	method2RequestType.Store("google.chromeos.moblab.v1beta1.BuildService/ListBuilds", reflect.TypeOf((*google_chromeos_moblab_v1beta1_build_servicepb.ListBuildsRequest)(nil)).Elem())
	method2ResponseType.Store("google.chromeos.moblab.v1beta1.BuildService/ListBuilds", reflect.TypeOf((*google_chromeos_moblab_v1beta1_build_servicepb.ListBuildsResponse)(nil)).Elem())
	method2RequestType.Store("google.chromeos.moblab.v1beta1.BuildService/ListModels", reflect.TypeOf((*google_chromeos_moblab_v1beta1_build_servicepb.ListModelsRequest)(nil)).Elem())
	method2ResponseType.Store("google.chromeos.moblab.v1beta1.BuildService/ListModels", reflect.TypeOf((*google_chromeos_moblab_v1beta1_build_servicepb.ListModelsResponse)(nil)).Elem())
	method2RequestType.Store("google.chromeos.moblab.v1beta1.BuildService/StageBuild", reflect.TypeOf((*google_chromeos_moblab_v1beta1_build_servicepb.StageBuildRequest)(nil)).Elem())
	method2ResponseType.Store("google.chromeos.moblab.v1beta1.BuildService/StageBuild", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
