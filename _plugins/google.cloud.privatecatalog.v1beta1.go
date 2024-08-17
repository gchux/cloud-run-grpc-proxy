// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_privatecatalog_v1beta1_privatecatalogpb "cloud.google.com/go/privatecatalog/apiv1beta1/privatecatalogpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.privatecatalog.v1beta1.PrivateCatalog", "cloudprivatecatalog.googleapis.com")
	method2RequestType.Store("google.cloud.privatecatalog.v1beta1.PrivateCatalog/SearchCatalogs", reflect.TypeOf((*google_cloud_privatecatalog_v1beta1_privatecatalogpb.SearchCatalogsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.privatecatalog.v1beta1.PrivateCatalog/SearchCatalogs", reflect.TypeOf((*google_cloud_privatecatalog_v1beta1_privatecatalogpb.SearchCatalogsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.privatecatalog.v1beta1.PrivateCatalog/SearchProducts", reflect.TypeOf((*google_cloud_privatecatalog_v1beta1_privatecatalogpb.SearchProductsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.privatecatalog.v1beta1.PrivateCatalog/SearchProducts", reflect.TypeOf((*google_cloud_privatecatalog_v1beta1_privatecatalogpb.SearchProductsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.privatecatalog.v1beta1.PrivateCatalog/SearchVersions", reflect.TypeOf((*google_cloud_privatecatalog_v1beta1_privatecatalogpb.SearchVersionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.privatecatalog.v1beta1.PrivateCatalog/SearchVersions", reflect.TypeOf((*google_cloud_privatecatalog_v1beta1_privatecatalogpb.SearchVersionsResponse)(nil)).Elem())
}
