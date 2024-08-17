// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_bigtable_v2_bigtablepb "cloud.google.com/go/bigtable/apiv2/bigtablepb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.bigtable.v2.Bigtable", "bigtable.googleapis.com")
	method2RequestType.Store("google.bigtable.v2.Bigtable/CheckAndMutateRow", reflect.TypeOf((*google_bigtable_v2_bigtablepb.CheckAndMutateRowRequest)(nil)).Elem())
	method2ResponseType.Store("google.bigtable.v2.Bigtable/CheckAndMutateRow", reflect.TypeOf((*google_bigtable_v2_bigtablepb.CheckAndMutateRowResponse)(nil)).Elem())
	method2RequestType.Store("google.bigtable.v2.Bigtable/ExecuteQuery", reflect.TypeOf((*google_bigtable_v2_bigtablepb.ExecuteQueryRequest)(nil)).Elem())
	method2ResponseType.Store("google.bigtable.v2.Bigtable/ExecuteQuery", reflect.TypeOf((*google_bigtable_v2_bigtablepb.ExecuteQueryResponse)(nil)).Elem())
	method2RequestType.Store("google.bigtable.v2.Bigtable/MutateRow", reflect.TypeOf((*google_bigtable_v2_bigtablepb.MutateRowRequest)(nil)).Elem())
	method2ResponseType.Store("google.bigtable.v2.Bigtable/MutateRow", reflect.TypeOf((*google_bigtable_v2_bigtablepb.MutateRowResponse)(nil)).Elem())
	method2RequestType.Store("google.bigtable.v2.Bigtable/MutateRows", reflect.TypeOf((*google_bigtable_v2_bigtablepb.MutateRowsRequest)(nil)).Elem())
	method2ResponseType.Store("google.bigtable.v2.Bigtable/MutateRows", reflect.TypeOf((*google_bigtable_v2_bigtablepb.MutateRowsResponse)(nil)).Elem())
	method2RequestType.Store("google.bigtable.v2.Bigtable/PingAndWarm", reflect.TypeOf((*google_bigtable_v2_bigtablepb.PingAndWarmRequest)(nil)).Elem())
	method2ResponseType.Store("google.bigtable.v2.Bigtable/PingAndWarm", reflect.TypeOf((*google_bigtable_v2_bigtablepb.PingAndWarmResponse)(nil)).Elem())
	method2RequestType.Store("google.bigtable.v2.Bigtable/ReadChangeStream", reflect.TypeOf((*google_bigtable_v2_bigtablepb.ReadChangeStreamRequest)(nil)).Elem())
	method2ResponseType.Store("google.bigtable.v2.Bigtable/ReadChangeStream", reflect.TypeOf((*google_bigtable_v2_bigtablepb.ReadChangeStreamResponse)(nil)).Elem())
	method2RequestType.Store("google.bigtable.v2.Bigtable/ReadModifyWriteRow", reflect.TypeOf((*google_bigtable_v2_bigtablepb.ReadModifyWriteRowRequest)(nil)).Elem())
	method2ResponseType.Store("google.bigtable.v2.Bigtable/ReadModifyWriteRow", reflect.TypeOf((*google_bigtable_v2_bigtablepb.ReadModifyWriteRowResponse)(nil)).Elem())
	method2RequestType.Store("google.bigtable.v2.Bigtable/ReadRows", reflect.TypeOf((*google_bigtable_v2_bigtablepb.ReadRowsRequest)(nil)).Elem())
	method2ResponseType.Store("google.bigtable.v2.Bigtable/ReadRows", reflect.TypeOf((*google_bigtable_v2_bigtablepb.ReadRowsResponse)(nil)).Elem())
	method2RequestType.Store("google.bigtable.v2.Bigtable/SampleRowKeys", reflect.TypeOf((*google_bigtable_v2_bigtablepb.SampleRowKeysRequest)(nil)).Elem())
	method2ResponseType.Store("google.bigtable.v2.Bigtable/SampleRowKeys", reflect.TypeOf((*google_bigtable_v2_bigtablepb.SampleRowKeysResponse)(nil)).Elem())
}