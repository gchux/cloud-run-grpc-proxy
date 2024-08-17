// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_servicehealth_v1_servicehealthpb "cloud.google.com/go/servicehealth/apiv1/servicehealthpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.servicehealth.v1.ServiceHealth", "servicehealth.googleapis.com")
	method2RequestType.Store("google.cloud.servicehealth.v1.ServiceHealth/GetEvent", reflect.TypeOf((*google_cloud_servicehealth_v1_servicehealthpb.GetEventRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.servicehealth.v1.ServiceHealth/GetEvent", reflect.TypeOf((*google_cloud_servicehealth_v1_servicehealthpb.Event)(nil)).Elem())
	method2RequestType.Store("google.cloud.servicehealth.v1.ServiceHealth/GetOrganizationEvent", reflect.TypeOf((*google_cloud_servicehealth_v1_servicehealthpb.GetOrganizationEventRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.servicehealth.v1.ServiceHealth/GetOrganizationEvent", reflect.TypeOf((*google_cloud_servicehealth_v1_servicehealthpb.OrganizationEvent)(nil)).Elem())
	method2RequestType.Store("google.cloud.servicehealth.v1.ServiceHealth/GetOrganizationImpact", reflect.TypeOf((*google_cloud_servicehealth_v1_servicehealthpb.GetOrganizationImpactRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.servicehealth.v1.ServiceHealth/GetOrganizationImpact", reflect.TypeOf((*google_cloud_servicehealth_v1_servicehealthpb.OrganizationImpact)(nil)).Elem())
	method2RequestType.Store("google.cloud.servicehealth.v1.ServiceHealth/ListEvents", reflect.TypeOf((*google_cloud_servicehealth_v1_servicehealthpb.ListEventsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.servicehealth.v1.ServiceHealth/ListEvents", reflect.TypeOf((*google_cloud_servicehealth_v1_servicehealthpb.ListEventsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.servicehealth.v1.ServiceHealth/ListOrganizationEvents", reflect.TypeOf((*google_cloud_servicehealth_v1_servicehealthpb.ListOrganizationEventsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.servicehealth.v1.ServiceHealth/ListOrganizationEvents", reflect.TypeOf((*google_cloud_servicehealth_v1_servicehealthpb.ListOrganizationEventsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.servicehealth.v1.ServiceHealth/ListOrganizationImpacts", reflect.TypeOf((*google_cloud_servicehealth_v1_servicehealthpb.ListOrganizationImpactsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.servicehealth.v1.ServiceHealth/ListOrganizationImpacts", reflect.TypeOf((*google_cloud_servicehealth_v1_servicehealthpb.ListOrganizationImpactsResponse)(nil)).Elem())
}