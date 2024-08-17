// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_shopping_merchant_quota_v1beta_quotapb "cloud.google.com/go/shopping/merchant/quota/apiv1beta/quotapb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.shopping.merchant.quota.v1beta.QuotaService", "merchantapi.googleapis.com")
	method2RequestType.Store("google.shopping.merchant.quota.v1beta.QuotaService/ListQuotaGroups", reflect.TypeOf((*google_shopping_merchant_quota_v1beta_quotapb.ListQuotaGroupsRequest)(nil)).Elem())
	method2ResponseType.Store("google.shopping.merchant.quota.v1beta.QuotaService/ListQuotaGroups", reflect.TypeOf((*google_shopping_merchant_quota_v1beta_quotapb.ListQuotaGroupsResponse)(nil)).Elem())
}
