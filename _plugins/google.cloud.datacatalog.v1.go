// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_datacatalog_v1_datacatalogpb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.datacatalog.v1.PolicyTagManagerSerialization", "datacatalog.googleapis.com")
	method2RequestType.Store("google.cloud.datacatalog.v1.PolicyTagManagerSerialization/ExportTaxonomies", reflect.TypeOf((*google_cloud_datacatalog_v1_datacatalogpb.ExportTaxonomiesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1.PolicyTagManagerSerialization/ExportTaxonomies", reflect.TypeOf((*google_cloud_datacatalog_v1_datacatalogpb.ExportTaxonomiesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1.PolicyTagManagerSerialization/ImportTaxonomies", reflect.TypeOf((*google_cloud_datacatalog_v1_datacatalogpb.ImportTaxonomiesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1.PolicyTagManagerSerialization/ImportTaxonomies", reflect.TypeOf((*google_cloud_datacatalog_v1_datacatalogpb.ImportTaxonomiesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1.PolicyTagManagerSerialization/ReplaceTaxonomy", reflect.TypeOf((*google_cloud_datacatalog_v1_datacatalogpb.ReplaceTaxonomyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1.PolicyTagManagerSerialization/ReplaceTaxonomy", reflect.TypeOf((*google_cloud_datacatalog_v1_datacatalogpb.Taxonomy)(nil)).Elem())
}