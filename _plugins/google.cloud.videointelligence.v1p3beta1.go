// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_videointelligence_v1p3beta1_videointelligencepb "cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.videointelligence.v1p3beta1.StreamingVideoIntelligenceService", "videointelligence.googleapis.com")
	service2Host.Store("google.cloud.videointelligence.v1p3beta1.VideoIntelligenceService", "videointelligence.googleapis.com")
	method2RequestType.Store("google.cloud.videointelligence.v1p3beta1.StreamingVideoIntelligenceService/StreamingAnnotateVideo", reflect.TypeOf((*google_cloud_videointelligence_v1p3beta1_videointelligencepb.StreamingAnnotateVideoRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.videointelligence.v1p3beta1.StreamingVideoIntelligenceService/StreamingAnnotateVideo", reflect.TypeOf((*google_cloud_videointelligence_v1p3beta1_videointelligencepb.StreamingAnnotateVideoResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.videointelligence.v1p3beta1.VideoIntelligenceService/AnnotateVideo", reflect.TypeOf((*google_cloud_videointelligence_v1p3beta1_videointelligencepb.AnnotateVideoRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.videointelligence.v1p3beta1.VideoIntelligenceService/AnnotateVideo", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
