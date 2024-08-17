// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_osconfig_agentendpoint_v1beta_agentendpointpb "cloud.google.com/go/osconfig/agentendpoint/apiv1beta/agentendpointpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.osconfig.agentendpoint.v1beta.AgentEndpointService", "osconfig.googleapis.com")
	method2RequestType.Store("google.cloud.osconfig.agentendpoint.v1beta.AgentEndpointService/LookupEffectiveGuestPolicy", reflect.TypeOf((*google_cloud_osconfig_agentendpoint_v1beta_agentendpointpb.LookupEffectiveGuestPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.agentendpoint.v1beta.AgentEndpointService/LookupEffectiveGuestPolicy", reflect.TypeOf((*google_cloud_osconfig_agentendpoint_v1beta_agentendpointpb.EffectiveGuestPolicy)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.agentendpoint.v1beta.AgentEndpointService/ReceiveTaskNotification", reflect.TypeOf((*google_cloud_osconfig_agentendpoint_v1beta_agentendpointpb.ReceiveTaskNotificationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.agentendpoint.v1beta.AgentEndpointService/ReceiveTaskNotification", reflect.TypeOf((*google_cloud_osconfig_agentendpoint_v1beta_agentendpointpb.ReceiveTaskNotificationResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.agentendpoint.v1beta.AgentEndpointService/RegisterAgent", reflect.TypeOf((*google_cloud_osconfig_agentendpoint_v1beta_agentendpointpb.RegisterAgentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.agentendpoint.v1beta.AgentEndpointService/RegisterAgent", reflect.TypeOf((*google_cloud_osconfig_agentendpoint_v1beta_agentendpointpb.RegisterAgentResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.agentendpoint.v1beta.AgentEndpointService/ReportTaskComplete", reflect.TypeOf((*google_cloud_osconfig_agentendpoint_v1beta_agentendpointpb.ReportTaskCompleteRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.agentendpoint.v1beta.AgentEndpointService/ReportTaskComplete", reflect.TypeOf((*google_cloud_osconfig_agentendpoint_v1beta_agentendpointpb.ReportTaskCompleteResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.agentendpoint.v1beta.AgentEndpointService/ReportTaskProgress", reflect.TypeOf((*google_cloud_osconfig_agentendpoint_v1beta_agentendpointpb.ReportTaskProgressRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.agentendpoint.v1beta.AgentEndpointService/ReportTaskProgress", reflect.TypeOf((*google_cloud_osconfig_agentendpoint_v1beta_agentendpointpb.ReportTaskProgressResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.osconfig.agentendpoint.v1beta.AgentEndpointService/StartNextTask", reflect.TypeOf((*google_cloud_osconfig_agentendpoint_v1beta_agentendpointpb.StartNextTaskRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.osconfig.agentendpoint.v1beta.AgentEndpointService/StartNextTask", reflect.TypeOf((*google_cloud_osconfig_agentendpoint_v1beta_agentendpointpb.StartNextTaskResponse)(nil)).Elem())
}