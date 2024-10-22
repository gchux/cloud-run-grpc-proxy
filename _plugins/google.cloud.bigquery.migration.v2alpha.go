// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_bigquery_migration_v2alpha_migrationpb "cloud.google.com/go/bigquery/migration/apiv2alpha/migrationpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.bigquery.migration.v2alpha.SqlTranslationService", "bigquerymigration.googleapis.com")
	method2RequestType.Store("google.cloud.bigquery.migration.v2alpha.SqlTranslationService/TranslateQuery", reflect.TypeOf((*google_cloud_bigquery_migration_v2alpha_migrationpb.TranslateQueryRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.migration.v2alpha.SqlTranslationService/TranslateQuery", reflect.TypeOf((*google_cloud_bigquery_migration_v2alpha_migrationpb.TranslateQueryResponse)(nil)).Elem())
}
