// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_documentai_v1_documentaipb "cloud.google.com/go/documentai/apiv1/documentaipb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.documentai.v1.DocumentProcessorService", "documentai.googleapis.com")
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/BatchProcessDocuments", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.BatchProcessRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/BatchProcessDocuments", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/CreateProcessor", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.CreateProcessorRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/CreateProcessor", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.Processor)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/DeleteProcessor", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.DeleteProcessorRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/DeleteProcessor", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/DeleteProcessorVersion", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.DeleteProcessorVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/DeleteProcessorVersion", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/DeployProcessorVersion", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.DeployProcessorVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/DeployProcessorVersion", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/DisableProcessor", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.DisableProcessorRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/DisableProcessor", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/EnableProcessor", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.EnableProcessorRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/EnableProcessor", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/EvaluateProcessorVersion", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.EvaluateProcessorVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/EvaluateProcessorVersion", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/FetchProcessorTypes", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.FetchProcessorTypesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/FetchProcessorTypes", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.FetchProcessorTypesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/GetEvaluation", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.GetEvaluationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/GetEvaluation", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.Evaluation)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/GetProcessor", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.GetProcessorRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/GetProcessor", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.Processor)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/GetProcessorType", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.GetProcessorTypeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/GetProcessorType", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.ProcessorType)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/GetProcessorVersion", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.GetProcessorVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/GetProcessorVersion", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.ProcessorVersion)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/ListEvaluations", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.ListEvaluationsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/ListEvaluations", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.ListEvaluationsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/ListProcessorTypes", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.ListProcessorTypesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/ListProcessorTypes", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.ListProcessorTypesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/ListProcessorVersions", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.ListProcessorVersionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/ListProcessorVersions", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.ListProcessorVersionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/ListProcessors", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.ListProcessorsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/ListProcessors", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.ListProcessorsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/ProcessDocument", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.ProcessRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/ProcessDocument", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.ProcessResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/ReviewDocument", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.ReviewDocumentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/ReviewDocument", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/SetDefaultProcessorVersion", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.SetDefaultProcessorVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/SetDefaultProcessorVersion", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/TrainProcessorVersion", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.TrainProcessorVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/TrainProcessorVersion", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.documentai.v1.DocumentProcessorService/UndeployProcessorVersion", reflect.TypeOf((*google_cloud_documentai_v1_documentaipb.UndeployProcessorVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.documentai.v1.DocumentProcessorService/UndeployProcessorVersion", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}