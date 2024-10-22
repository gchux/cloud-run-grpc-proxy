// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_visionai_v1alpha1_streaming_servicepb "google.golang.org/genproto/googleapis/cloud/visionai/v1alpha1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.visionai.v1alpha1.StreamingService", "visionai.googleapis.com")
	method2RequestType.Store("google.cloud.visionai.v1alpha1.StreamingService/AcquireLease", reflect.TypeOf((*google_cloud_visionai_v1alpha1_streaming_servicepb.AcquireLeaseRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1alpha1.StreamingService/AcquireLease", reflect.TypeOf((*google_cloud_visionai_v1alpha1_streaming_servicepb.Lease)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1alpha1.StreamingService/ReceiveEvents", reflect.TypeOf((*google_cloud_visionai_v1alpha1_streaming_servicepb.ReceiveEventsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1alpha1.StreamingService/ReceiveEvents", reflect.TypeOf((*google_cloud_visionai_v1alpha1_streaming_servicepb.ReceiveEventsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1alpha1.StreamingService/ReceivePackets", reflect.TypeOf((*google_cloud_visionai_v1alpha1_streaming_servicepb.ReceivePacketsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1alpha1.StreamingService/ReceivePackets", reflect.TypeOf((*google_cloud_visionai_v1alpha1_streaming_servicepb.ReceivePacketsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1alpha1.StreamingService/ReleaseLease", reflect.TypeOf((*google_cloud_visionai_v1alpha1_streaming_servicepb.ReleaseLeaseRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1alpha1.StreamingService/ReleaseLease", reflect.TypeOf((*google_cloud_visionai_v1alpha1_streaming_servicepb.ReleaseLeaseResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1alpha1.StreamingService/RenewLease", reflect.TypeOf((*google_cloud_visionai_v1alpha1_streaming_servicepb.RenewLeaseRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1alpha1.StreamingService/RenewLease", reflect.TypeOf((*google_cloud_visionai_v1alpha1_streaming_servicepb.Lease)(nil)).Elem())
	method2RequestType.Store("google.cloud.visionai.v1alpha1.StreamingService/SendPackets", reflect.TypeOf((*google_cloud_visionai_v1alpha1_streaming_servicepb.SendPacketsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.visionai.v1alpha1.StreamingService/SendPackets", reflect.TypeOf((*google_cloud_visionai_v1alpha1_streaming_servicepb.SendPacketsResponse)(nil)).Elem())
}
