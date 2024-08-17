// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_bigquery_storage_v1_storagepb "cloud.google.com/go/bigquery/storage/apiv1/storagepb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.bigquery.storage.v1.BigQueryRead", "bigquerystorage.googleapis.com")
	service2Host.Store("google.cloud.bigquery.storage.v1.BigQueryWrite", "bigquerystorage.googleapis.com")
	method2RequestType.Store("google.cloud.bigquery.storage.v1.BigQueryRead/CreateReadSession", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.CreateReadSessionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.storage.v1.BigQueryRead/CreateReadSession", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.ReadSession)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.storage.v1.BigQueryRead/ReadRows", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.ReadRowsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.storage.v1.BigQueryRead/ReadRows", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.ReadRowsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.storage.v1.BigQueryRead/SplitReadStream", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.SplitReadStreamRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.storage.v1.BigQueryRead/SplitReadStream", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.SplitReadStreamResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.storage.v1.BigQueryWrite/AppendRows", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.AppendRowsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.storage.v1.BigQueryWrite/AppendRows", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.AppendRowsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.storage.v1.BigQueryWrite/BatchCommitWriteStreams", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.BatchCommitWriteStreamsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.storage.v1.BigQueryWrite/BatchCommitWriteStreams", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.BatchCommitWriteStreamsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.storage.v1.BigQueryWrite/CreateWriteStream", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.CreateWriteStreamRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.storage.v1.BigQueryWrite/CreateWriteStream", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.WriteStream)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.storage.v1.BigQueryWrite/FinalizeWriteStream", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.FinalizeWriteStreamRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.storage.v1.BigQueryWrite/FinalizeWriteStream", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.FinalizeWriteStreamResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.storage.v1.BigQueryWrite/FlushRows", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.FlushRowsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.storage.v1.BigQueryWrite/FlushRows", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.FlushRowsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.storage.v1.BigQueryWrite/GetWriteStream", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.GetWriteStreamRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.storage.v1.BigQueryWrite/GetWriteStream", reflect.TypeOf((*google_cloud_bigquery_storage_v1_storagepb.WriteStream)(nil)).Elem())
}
