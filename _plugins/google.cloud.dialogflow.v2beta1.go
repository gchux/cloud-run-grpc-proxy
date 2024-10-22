// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_dialogflow_v2beta1_dialogflowpb "cloud.google.com/go/dialogflow/apiv2beta1/dialogflowpb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.dialogflow.v2beta1.EntityTypes", "dialogflow.googleapis.com")
	method2RequestType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/BatchCreateEntities", reflect.TypeOf((*google_cloud_dialogflow_v2beta1_dialogflowpb.BatchCreateEntitiesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/BatchCreateEntities", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/BatchDeleteEntities", reflect.TypeOf((*google_cloud_dialogflow_v2beta1_dialogflowpb.BatchDeleteEntitiesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/BatchDeleteEntities", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/BatchDeleteEntityTypes", reflect.TypeOf((*google_cloud_dialogflow_v2beta1_dialogflowpb.BatchDeleteEntityTypesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/BatchDeleteEntityTypes", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/BatchUpdateEntities", reflect.TypeOf((*google_cloud_dialogflow_v2beta1_dialogflowpb.BatchUpdateEntitiesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/BatchUpdateEntities", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/BatchUpdateEntityTypes", reflect.TypeOf((*google_cloud_dialogflow_v2beta1_dialogflowpb.BatchUpdateEntityTypesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/BatchUpdateEntityTypes", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/CreateEntityType", reflect.TypeOf((*google_cloud_dialogflow_v2beta1_dialogflowpb.CreateEntityTypeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/CreateEntityType", reflect.TypeOf((*google_cloud_dialogflow_v2beta1_dialogflowpb.EntityType)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/DeleteEntityType", reflect.TypeOf((*google_cloud_dialogflow_v2beta1_dialogflowpb.DeleteEntityTypeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/DeleteEntityType", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/GetEntityType", reflect.TypeOf((*google_cloud_dialogflow_v2beta1_dialogflowpb.GetEntityTypeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/GetEntityType", reflect.TypeOf((*google_cloud_dialogflow_v2beta1_dialogflowpb.EntityType)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/ListEntityTypes", reflect.TypeOf((*google_cloud_dialogflow_v2beta1_dialogflowpb.ListEntityTypesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/ListEntityTypes", reflect.TypeOf((*google_cloud_dialogflow_v2beta1_dialogflowpb.ListEntityTypesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/UpdateEntityType", reflect.TypeOf((*google_cloud_dialogflow_v2beta1_dialogflowpb.UpdateEntityTypeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.dialogflow.v2beta1.EntityTypes/UpdateEntityType", reflect.TypeOf((*google_cloud_dialogflow_v2beta1_dialogflowpb.EntityType)(nil)).Elem())
}
