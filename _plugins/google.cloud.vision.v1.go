// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_vision_v1_visionpb "cloud.google.com/go/vision/v2/apiv1/visionpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.vision.v1.ImageAnnotator", "vision.googleapis.com")
	method2RequestType.Store("google.cloud.vision.v1.ImageAnnotator/AsyncBatchAnnotateFiles", reflect.TypeOf((*google_cloud_vision_v1_visionpb.AsyncBatchAnnotateFilesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vision.v1.ImageAnnotator/AsyncBatchAnnotateFiles", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vision.v1.ImageAnnotator/AsyncBatchAnnotateImages", reflect.TypeOf((*google_cloud_vision_v1_visionpb.AsyncBatchAnnotateImagesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vision.v1.ImageAnnotator/AsyncBatchAnnotateImages", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.vision.v1.ImageAnnotator/BatchAnnotateFiles", reflect.TypeOf((*google_cloud_vision_v1_visionpb.BatchAnnotateFilesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vision.v1.ImageAnnotator/BatchAnnotateFiles", reflect.TypeOf((*google_cloud_vision_v1_visionpb.BatchAnnotateFilesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.vision.v1.ImageAnnotator/BatchAnnotateImages", reflect.TypeOf((*google_cloud_vision_v1_visionpb.BatchAnnotateImagesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vision.v1.ImageAnnotator/BatchAnnotateImages", reflect.TypeOf((*google_cloud_vision_v1_visionpb.BatchAnnotateImagesResponse)(nil)).Elem())
}
