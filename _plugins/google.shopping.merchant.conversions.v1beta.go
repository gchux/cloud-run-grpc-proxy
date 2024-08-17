// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_shopping_merchant_conversions_v1beta_conversionspb "cloud.google.com/go/shopping/merchant/conversions/apiv1beta/conversionspb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.shopping.merchant.conversions.v1beta.ConversionSourcesService", "merchantapi.googleapis.com")
	method2RequestType.Store("google.shopping.merchant.conversions.v1beta.ConversionSourcesService/CreateConversionSource", reflect.TypeOf((*google_shopping_merchant_conversions_v1beta_conversionspb.CreateConversionSourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.shopping.merchant.conversions.v1beta.ConversionSourcesService/CreateConversionSource", reflect.TypeOf((*google_shopping_merchant_conversions_v1beta_conversionspb.ConversionSource)(nil)).Elem())
	method2RequestType.Store("google.shopping.merchant.conversions.v1beta.ConversionSourcesService/DeleteConversionSource", reflect.TypeOf((*google_shopping_merchant_conversions_v1beta_conversionspb.DeleteConversionSourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.shopping.merchant.conversions.v1beta.ConversionSourcesService/DeleteConversionSource", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.shopping.merchant.conversions.v1beta.ConversionSourcesService/GetConversionSource", reflect.TypeOf((*google_shopping_merchant_conversions_v1beta_conversionspb.GetConversionSourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.shopping.merchant.conversions.v1beta.ConversionSourcesService/GetConversionSource", reflect.TypeOf((*google_shopping_merchant_conversions_v1beta_conversionspb.ConversionSource)(nil)).Elem())
	method2RequestType.Store("google.shopping.merchant.conversions.v1beta.ConversionSourcesService/ListConversionSources", reflect.TypeOf((*google_shopping_merchant_conversions_v1beta_conversionspb.ListConversionSourcesRequest)(nil)).Elem())
	method2ResponseType.Store("google.shopping.merchant.conversions.v1beta.ConversionSourcesService/ListConversionSources", reflect.TypeOf((*google_shopping_merchant_conversions_v1beta_conversionspb.ListConversionSourcesResponse)(nil)).Elem())
	method2RequestType.Store("google.shopping.merchant.conversions.v1beta.ConversionSourcesService/UndeleteConversionSource", reflect.TypeOf((*google_shopping_merchant_conversions_v1beta_conversionspb.UndeleteConversionSourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.shopping.merchant.conversions.v1beta.ConversionSourcesService/UndeleteConversionSource", reflect.TypeOf((*google_shopping_merchant_conversions_v1beta_conversionspb.ConversionSource)(nil)).Elem())
	method2RequestType.Store("google.shopping.merchant.conversions.v1beta.ConversionSourcesService/UpdateConversionSource", reflect.TypeOf((*google_shopping_merchant_conversions_v1beta_conversionspb.UpdateConversionSourceRequest)(nil)).Elem())
	method2ResponseType.Store("google.shopping.merchant.conversions.v1beta.ConversionSourcesService/UpdateConversionSource", reflect.TypeOf((*google_shopping_merchant_conversions_v1beta_conversionspb.ConversionSource)(nil)).Elem())
}