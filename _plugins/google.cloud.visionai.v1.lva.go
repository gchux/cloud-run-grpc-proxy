// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_visionai_v1_lva_servicepb "google.golang.org/genproto/googleapis/cloud/visionai/v1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.visionai.v1.LiveVideoAnalytics", "visionai.googleapis.com")
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/BatchRunProcess", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.BatchRunProcessRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/BatchRunProcess", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/CreateAnalysis", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.CreateAnalysisRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/CreateAnalysis", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/CreateOperator", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.CreateOperatorRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/CreateOperator", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/CreateProcess", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.CreateProcessRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/CreateProcess", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/DeleteAnalysis", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.DeleteAnalysisRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/DeleteAnalysis", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/DeleteOperator", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.DeleteOperatorRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/DeleteOperator", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/DeleteProcess", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.DeleteProcessRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/DeleteProcess", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/GetAnalysis", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.GetAnalysisRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/GetAnalysis", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.Analysis)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/GetOperator", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.GetOperatorRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/GetOperator", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.Operator)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/GetProcess", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.GetProcessRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/GetProcess", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.Process)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/ListAnalyses", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.ListAnalysesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/ListAnalyses", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.ListAnalysesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/ListOperators", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.ListOperatorsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/ListOperators", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.ListOperatorsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/ListProcesses", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.ListProcessesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/ListProcesses", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.ListProcessesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/ListPublicOperators", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.ListPublicOperatorsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/ListPublicOperators", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.ListPublicOperatorsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/ResolveOperatorInfo", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.ResolveOperatorInfoRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/ResolveOperatorInfo", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.ResolveOperatorInfoResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/UpdateAnalysis", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.UpdateAnalysisRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/UpdateAnalysis", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/UpdateOperator", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.UpdateOperatorRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/UpdateOperator", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/UpdateProcess", reflect.TypeOf((*google_cloud_visionai_v1_lva_servicepb.UpdateProcessRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1.LiveVideoAnalytics/UpdateProcess", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
