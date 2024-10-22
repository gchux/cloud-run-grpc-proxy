// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_firestore_admin_v1beta1_firestore_adminpb "google.golang.org/genproto/googleapis/firestore/admin/v1beta1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.firestore.admin.v1beta1.FirestoreAdmin", "firestore.googleapis.com")
	method2RequestType.Store("google.firestore.admin.v1beta1.FirestoreAdmin/CreateIndex", reflect.TypeOf((*google_firestore_admin_v1beta1_firestore_adminpb.CreateIndexRequest)(nil)).Elem())
	method2ResponseType.Store("google.firestore.admin.v1beta1.FirestoreAdmin/CreateIndex", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.firestore.admin.v1beta1.FirestoreAdmin/DeleteIndex", reflect.TypeOf((*google_firestore_admin_v1beta1_firestore_adminpb.DeleteIndexRequest)(nil)).Elem())
	method2ResponseType.Store("google.firestore.admin.v1beta1.FirestoreAdmin/DeleteIndex", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.firestore.admin.v1beta1.FirestoreAdmin/ExportDocuments", reflect.TypeOf((*google_firestore_admin_v1beta1_firestore_adminpb.ExportDocumentsRequest)(nil)).Elem())
	method2ResponseType.Store("google.firestore.admin.v1beta1.FirestoreAdmin/ExportDocuments", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.firestore.admin.v1beta1.FirestoreAdmin/GetIndex", reflect.TypeOf((*google_firestore_admin_v1beta1_firestore_adminpb.GetIndexRequest)(nil)).Elem())
	method2ResponseType.Store("google.firestore.admin.v1beta1.FirestoreAdmin/GetIndex", reflect.TypeOf((*google_firestore_admin_v1beta1_firestore_adminpb.Index)(nil)).Elem())
	method2RequestType.Store("google.firestore.admin.v1beta1.FirestoreAdmin/ImportDocuments", reflect.TypeOf((*google_firestore_admin_v1beta1_firestore_adminpb.ImportDocumentsRequest)(nil)).Elem())
	method2ResponseType.Store("google.firestore.admin.v1beta1.FirestoreAdmin/ImportDocuments", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.firestore.admin.v1beta1.FirestoreAdmin/ListIndexes", reflect.TypeOf((*google_firestore_admin_v1beta1_firestore_adminpb.ListIndexesRequest)(nil)).Elem())
	method2ResponseType.Store("google.firestore.admin.v1beta1.FirestoreAdmin/ListIndexes", reflect.TypeOf((*google_firestore_admin_v1beta1_firestore_adminpb.ListIndexesResponse)(nil)).Elem())
}
