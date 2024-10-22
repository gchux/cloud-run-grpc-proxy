// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_iam_v1_iampb "cloud.google.com/go/iam/apiv1/iampb"
	google_cloud_iap_v1beta1_servicepb "google.golang.org/genproto/googleapis/cloud/iap/v1beta1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1", "iap.googleapis.com")
	method2RequestType.Store("google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.GetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.SetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.iap.v1beta1.IdentityAwareProxyAdminV1Beta1/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsResponse)(nil)).Elem())
}
