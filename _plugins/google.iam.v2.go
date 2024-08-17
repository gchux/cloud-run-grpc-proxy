// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_iam_v2_iampb "cloud.google.com/go/iam/apiv2/iampb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.iam.v2.Policies", "iam.googleapis.com")
	method2RequestType.Store("google.iam.v2.Policies/CreatePolicy", reflect.TypeOf((*google_iam_v2_iampb.CreatePolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.v2.Policies/CreatePolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.iam.v2.Policies/DeletePolicy", reflect.TypeOf((*google_iam_v2_iampb.DeletePolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.v2.Policies/DeletePolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.iam.v2.Policies/GetPolicy", reflect.TypeOf((*google_iam_v2_iampb.GetPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.v2.Policies/GetPolicy", reflect.TypeOf((*google_iam_v2_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.iam.v2.Policies/ListPolicies", reflect.TypeOf((*google_iam_v2_iampb.ListPoliciesRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.v2.Policies/ListPolicies", reflect.TypeOf((*google_iam_v2_iampb.ListPoliciesResponse)(nil)).Elem())
	method2RequestType.Store("google.iam.v2.Policies/UpdatePolicy", reflect.TypeOf((*google_iam_v2_iampb.UpdatePolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.v2.Policies/UpdatePolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
