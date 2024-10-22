// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_vision_v1p1beta1_visionpb "cloud.google.com/go/vision/v2/apiv1p1beta1/visionpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.vision.v1p1beta1.ImageAnnotator", "vision.googleapis.com")
	method2RequestType.Store("google.cloud.vision.v1p1beta1.ImageAnnotator/BatchAnnotateImages", reflect.TypeOf((*google_cloud_vision_v1p1beta1_visionpb.BatchAnnotateImagesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.vision.v1p1beta1.ImageAnnotator/BatchAnnotateImages", reflect.TypeOf((*google_cloud_vision_v1p1beta1_visionpb.BatchAnnotateImagesResponse)(nil)).Elem())
}
