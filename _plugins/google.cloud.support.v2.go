// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_support_v2_supportpb "cloud.google.com/go/support/apiv2/supportpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.support.v2.CommentService", "cloudsupport.googleapis.com")
	method2RequestType.Store("google.cloud.support.v2.CommentService/CreateComment", reflect.TypeOf((*google_cloud_support_v2_supportpb.CreateCommentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.support.v2.CommentService/CreateComment", reflect.TypeOf((*google_cloud_support_v2_supportpb.Comment)(nil)).Elem())
	method2RequestType.Store("google.cloud.support.v2.CommentService/ListComments", reflect.TypeOf((*google_cloud_support_v2_supportpb.ListCommentsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.support.v2.CommentService/ListComments", reflect.TypeOf((*google_cloud_support_v2_supportpb.ListCommentsResponse)(nil)).Elem())
}
