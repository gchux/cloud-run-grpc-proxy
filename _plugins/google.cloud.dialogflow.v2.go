// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_dialogflow_v2_dialogflowpb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.dialogflow.v2.ConversationDatasets", "dialogflow.googleapis.com")
	method2RequestType.Store("google.cloud.dialogflow.v2.ConversationDatasets/CreateConversationDataset", reflect.TypeOf((*google_cloud_dialogflow_v2_dialogflowpb.CreateConversationDatasetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.v2.ConversationDatasets/CreateConversationDataset", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.v2.ConversationDatasets/DeleteConversationDataset", reflect.TypeOf((*google_cloud_dialogflow_v2_dialogflowpb.DeleteConversationDatasetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.v2.ConversationDatasets/DeleteConversationDataset", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.v2.ConversationDatasets/GetConversationDataset", reflect.TypeOf((*google_cloud_dialogflow_v2_dialogflowpb.GetConversationDatasetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.v2.ConversationDatasets/GetConversationDataset", reflect.TypeOf((*google_cloud_dialogflow_v2_dialogflowpb.ConversationDataset)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.v2.ConversationDatasets/ImportConversationData", reflect.TypeOf((*google_cloud_dialogflow_v2_dialogflowpb.ImportConversationDataRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.v2.ConversationDatasets/ImportConversationData", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.v2.ConversationDatasets/ListConversationDatasets", reflect.TypeOf((*google_cloud_dialogflow_v2_dialogflowpb.ListConversationDatasetsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.v2.ConversationDatasets/ListConversationDatasets", reflect.TypeOf((*google_cloud_dialogflow_v2_dialogflowpb.ListConversationDatasetsResponse)(nil)).Elem())
}
