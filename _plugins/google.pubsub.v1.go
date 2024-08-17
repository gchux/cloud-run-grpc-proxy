// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_pubsub_v1_pubsubpb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.pubsub.v1.SchemaService", "pubsub.googleapis.com")
	method2RequestType.Store("google.pubsub.v1.SchemaService/CommitSchema", reflect.TypeOf((*google_pubsub_v1_pubsubpb.CommitSchemaRequest)(nil)).Elem())
	method2ResponseType.Store("google.pubsub.v1.SchemaService/CommitSchema", reflect.TypeOf((*google_pubsub_v1_pubsubpb.Schema)(nil)).Elem())
	method2RequestType.Store("google.pubsub.v1.SchemaService/CreateSchema", reflect.TypeOf((*google_pubsub_v1_pubsubpb.CreateSchemaRequest)(nil)).Elem())
	method2ResponseType.Store("google.pubsub.v1.SchemaService/CreateSchema", reflect.TypeOf((*google_pubsub_v1_pubsubpb.Schema)(nil)).Elem())
	method2RequestType.Store("google.pubsub.v1.SchemaService/DeleteSchema", reflect.TypeOf((*google_pubsub_v1_pubsubpb.DeleteSchemaRequest)(nil)).Elem())
	method2ResponseType.Store("google.pubsub.v1.SchemaService/DeleteSchema", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.pubsub.v1.SchemaService/DeleteSchemaRevision", reflect.TypeOf((*google_pubsub_v1_pubsubpb.DeleteSchemaRevisionRequest)(nil)).Elem())
	method2ResponseType.Store("google.pubsub.v1.SchemaService/DeleteSchemaRevision", reflect.TypeOf((*google_pubsub_v1_pubsubpb.Schema)(nil)).Elem())
	method2RequestType.Store("google.pubsub.v1.SchemaService/GetSchema", reflect.TypeOf((*google_pubsub_v1_pubsubpb.GetSchemaRequest)(nil)).Elem())
	method2ResponseType.Store("google.pubsub.v1.SchemaService/GetSchema", reflect.TypeOf((*google_pubsub_v1_pubsubpb.Schema)(nil)).Elem())
	method2RequestType.Store("google.pubsub.v1.SchemaService/ListSchemaRevisions", reflect.TypeOf((*google_pubsub_v1_pubsubpb.ListSchemaRevisionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.pubsub.v1.SchemaService/ListSchemaRevisions", reflect.TypeOf((*google_pubsub_v1_pubsubpb.ListSchemaRevisionsResponse)(nil)).Elem())
	method2RequestType.Store("google.pubsub.v1.SchemaService/ListSchemas", reflect.TypeOf((*google_pubsub_v1_pubsubpb.ListSchemasRequest)(nil)).Elem())
	method2ResponseType.Store("google.pubsub.v1.SchemaService/ListSchemas", reflect.TypeOf((*google_pubsub_v1_pubsubpb.ListSchemasResponse)(nil)).Elem())
	method2RequestType.Store("google.pubsub.v1.SchemaService/RollbackSchema", reflect.TypeOf((*google_pubsub_v1_pubsubpb.RollbackSchemaRequest)(nil)).Elem())
	method2ResponseType.Store("google.pubsub.v1.SchemaService/RollbackSchema", reflect.TypeOf((*google_pubsub_v1_pubsubpb.Schema)(nil)).Elem())
	method2RequestType.Store("google.pubsub.v1.SchemaService/ValidateMessage", reflect.TypeOf((*google_pubsub_v1_pubsubpb.ValidateMessageRequest)(nil)).Elem())
	method2ResponseType.Store("google.pubsub.v1.SchemaService/ValidateMessage", reflect.TypeOf((*google_pubsub_v1_pubsubpb.ValidateMessageResponse)(nil)).Elem())
	method2RequestType.Store("google.pubsub.v1.SchemaService/ValidateSchema", reflect.TypeOf((*google_pubsub_v1_pubsubpb.ValidateSchemaRequest)(nil)).Elem())
	method2ResponseType.Store("google.pubsub.v1.SchemaService/ValidateSchema", reflect.TypeOf((*google_pubsub_v1_pubsubpb.ValidateSchemaResponse)(nil)).Elem())
}
