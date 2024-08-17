// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_dialogflow_cx_v3beta1_cxpb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.dialogflow.cx.v3beta1.Flows", "dialogflow.googleapis.com")
	method2RequestType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/CreateFlow", reflect.TypeOf((*google_cloud_dialogflow_cx_v3beta1_cxpb.CreateFlowRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/CreateFlow", reflect.TypeOf((*google_cloud_dialogflow_cx_v3beta1_cxpb.Flow)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/DeleteFlow", reflect.TypeOf((*google_cloud_dialogflow_cx_v3beta1_cxpb.DeleteFlowRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/DeleteFlow", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/ExportFlow", reflect.TypeOf((*google_cloud_dialogflow_cx_v3beta1_cxpb.ExportFlowRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/ExportFlow", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/GetFlow", reflect.TypeOf((*google_cloud_dialogflow_cx_v3beta1_cxpb.GetFlowRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/GetFlow", reflect.TypeOf((*google_cloud_dialogflow_cx_v3beta1_cxpb.Flow)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/GetFlowValidationResult", reflect.TypeOf((*google_cloud_dialogflow_cx_v3beta1_cxpb.GetFlowValidationResultRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/GetFlowValidationResult", reflect.TypeOf((*google_cloud_dialogflow_cx_v3beta1_cxpb.FlowValidationResult)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/ImportFlow", reflect.TypeOf((*google_cloud_dialogflow_cx_v3beta1_cxpb.ImportFlowRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/ImportFlow", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/ListFlows", reflect.TypeOf((*google_cloud_dialogflow_cx_v3beta1_cxpb.ListFlowsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/ListFlows", reflect.TypeOf((*google_cloud_dialogflow_cx_v3beta1_cxpb.ListFlowsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/TrainFlow", reflect.TypeOf((*google_cloud_dialogflow_cx_v3beta1_cxpb.TrainFlowRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/TrainFlow", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/UpdateFlow", reflect.TypeOf((*google_cloud_dialogflow_cx_v3beta1_cxpb.UpdateFlowRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/UpdateFlow", reflect.TypeOf((*google_cloud_dialogflow_cx_v3beta1_cxpb.Flow)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/ValidateFlow", reflect.TypeOf((*google_cloud_dialogflow_cx_v3beta1_cxpb.ValidateFlowRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.cx.v3beta1.Flows/ValidateFlow", reflect.TypeOf((*google_cloud_dialogflow_cx_v3beta1_cxpb.FlowValidationResult)(nil)).Elem())
}
