// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_marketingplatform_admin_v1alpha_admin "google.golang.org/genproto/googleapis/marketingplatform/admin/v1alpha"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.marketingplatform.admin.v1alpha.MarketingplatformAdminService", "marketingplatformadmin.googleapis.com")
	method2RequestType.Store("google.marketingplatform.admin.v1alpha.MarketingplatformAdminService/CreateAnalyticsAccountLink", reflect.TypeOf((*google_marketingplatform_admin_v1alpha_admin.CreateAnalyticsAccountLinkRequest)(nil)).Elem())
	method2ResponseType.Store("google.marketingplatform.admin.v1alpha.MarketingplatformAdminService/CreateAnalyticsAccountLink", reflect.TypeOf((*google_marketingplatform_admin_v1alpha_admin.AnalyticsAccountLink)(nil)).Elem())
	method2RequestType.Store("google.marketingplatform.admin.v1alpha.MarketingplatformAdminService/DeleteAnalyticsAccountLink", reflect.TypeOf((*google_marketingplatform_admin_v1alpha_admin.DeleteAnalyticsAccountLinkRequest)(nil)).Elem())
	method2ResponseType.Store("google.marketingplatform.admin.v1alpha.MarketingplatformAdminService/DeleteAnalyticsAccountLink", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.marketingplatform.admin.v1alpha.MarketingplatformAdminService/GetOrganization", reflect.TypeOf((*google_marketingplatform_admin_v1alpha_admin.GetOrganizationRequest)(nil)).Elem())
	method2ResponseType.Store("google.marketingplatform.admin.v1alpha.MarketingplatformAdminService/GetOrganization", reflect.TypeOf((*google_marketingplatform_admin_v1alpha_admin.Organization)(nil)).Elem())
	method2RequestType.Store("google.marketingplatform.admin.v1alpha.MarketingplatformAdminService/ListAnalyticsAccountLinks", reflect.TypeOf((*google_marketingplatform_admin_v1alpha_admin.ListAnalyticsAccountLinksRequest)(nil)).Elem())
	method2ResponseType.Store("google.marketingplatform.admin.v1alpha.MarketingplatformAdminService/ListAnalyticsAccountLinks", reflect.TypeOf((*google_marketingplatform_admin_v1alpha_admin.ListAnalyticsAccountLinksResponse)(nil)).Elem())
	method2RequestType.Store("google.marketingplatform.admin.v1alpha.MarketingplatformAdminService/SetPropertyServiceLevel", reflect.TypeOf((*google_marketingplatform_admin_v1alpha_admin.SetPropertyServiceLevelRequest)(nil)).Elem())
	method2ResponseType.Store("google.marketingplatform.admin.v1alpha.MarketingplatformAdminService/SetPropertyServiceLevel", reflect.TypeOf((*google_marketingplatform_admin_v1alpha_admin.SetPropertyServiceLevelResponse)(nil)).Elem())
}
