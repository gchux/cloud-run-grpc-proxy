// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_policytroubleshooter_iam_v3_iampb "cloud.google.com/go/policytroubleshooter/iam/apiv3/iampb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.policytroubleshooter.iam.v3.PolicyTroubleshooter", "policytroubleshooter.googleapis.com")
	method2RequestType.Store("google.cloud.policytroubleshooter.iam.v3.PolicyTroubleshooter/TroubleshootIamPolicy", reflect.TypeOf((*google_cloud_policytroubleshooter_iam_v3_iampb.TroubleshootIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.policytroubleshooter.iam.v3.PolicyTroubleshooter/TroubleshootIamPolicy", reflect.TypeOf((*google_cloud_policytroubleshooter_iam_v3_iampb.TroubleshootIamPolicyResponse)(nil)).Elem())
}
