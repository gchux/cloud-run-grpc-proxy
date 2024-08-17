// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_api_serviceusage_v1beta1_serviceusagepb "google.golang.org/genproto/googleapis/api/serviceusage/v1beta1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.api.serviceusage.v1beta1.ServiceUsage", "serviceusage.googleapis.com")
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/BatchEnableServices", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.BatchEnableServicesRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/BatchEnableServices", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/CreateAdminOverride", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.CreateAdminOverrideRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/CreateAdminOverride", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/CreateConsumerOverride", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.CreateConsumerOverrideRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/CreateConsumerOverride", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/DeleteAdminOverride", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.DeleteAdminOverrideRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/DeleteAdminOverride", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/DeleteConsumerOverride", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.DeleteConsumerOverrideRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/DeleteConsumerOverride", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/DisableService", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.DisableServiceRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/DisableService", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/EnableService", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.EnableServiceRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/EnableService", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/GenerateServiceIdentity", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.GenerateServiceIdentityRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/GenerateServiceIdentity", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/GetConsumerQuotaLimit", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.GetConsumerQuotaLimitRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/GetConsumerQuotaLimit", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.ConsumerQuotaLimit)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/GetConsumerQuotaMetric", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.GetConsumerQuotaMetricRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/GetConsumerQuotaMetric", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.ConsumerQuotaMetric)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/GetService", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.GetServiceRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/GetService", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.Service)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/ImportAdminOverrides", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.ImportAdminOverridesRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/ImportAdminOverrides", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/ImportConsumerOverrides", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.ImportConsumerOverridesRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/ImportConsumerOverrides", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/ListAdminOverrides", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.ListAdminOverridesRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/ListAdminOverrides", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.ListAdminOverridesResponse)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/ListConsumerOverrides", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.ListConsumerOverridesRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/ListConsumerOverrides", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.ListConsumerOverridesResponse)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/ListConsumerQuotaMetrics", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.ListConsumerQuotaMetricsRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/ListConsumerQuotaMetrics", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.ListConsumerQuotaMetricsResponse)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/ListServices", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.ListServicesRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/ListServices", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.ListServicesResponse)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/UpdateAdminOverride", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.UpdateAdminOverrideRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/UpdateAdminOverride", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.api.serviceusage.v1beta1.ServiceUsage/UpdateConsumerOverride", reflect.TypeOf((*google_api_serviceusage_v1beta1_serviceusagepb.UpdateConsumerOverrideRequest)(nil)).Elem())
	method2ResponseType.Store("google.api.serviceusage.v1beta1.ServiceUsage/UpdateConsumerOverride", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}