// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_shopping_merchant_reports_v1beta_reportspb "cloud.google.com/go/shopping/merchant/reports/apiv1beta/reportspb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.shopping.merchant.reports.v1beta.ReportService", "merchantapi.googleapis.com")
	method2RequestType.Store("google.shopping.merchant.reports.v1beta.ReportService/Search", reflect.TypeOf((*google_shopping_merchant_reports_v1beta_reportspb.SearchRequest)(nil)).Elem())
	method2ResponseType.Store("google.shopping.merchant.reports.v1beta.ReportService/Search", reflect.TypeOf((*google_shopping_merchant_reports_v1beta_reportspb.SearchResponse)(nil)).Elem())
}
