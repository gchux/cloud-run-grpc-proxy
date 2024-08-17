// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_devtools_testing_v1_application_detailspb "google.golang.org/genproto/googleapis/devtools/testing/v1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.devtools.testing.v1.ApplicationDetailService", "testing.googleapis.com")
	method2RequestType.Store("google.devtools.testing.v1.ApplicationDetailService/GetApkDetails", reflect.TypeOf((*google_devtools_testing_v1_application_detailspb.GetApkDetailsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.testing.v1.ApplicationDetailService/GetApkDetails", reflect.TypeOf((*google_devtools_testing_v1_application_detailspb.GetApkDetailsResponse)(nil)).Elem())
}