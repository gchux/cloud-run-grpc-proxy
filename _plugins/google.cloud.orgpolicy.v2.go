// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_orgpolicy_v2_orgpolicypb "cloud.google.com/go/orgpolicy/apiv2/orgpolicypb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.orgpolicy.v2.OrgPolicy", "orgpolicy.googleapis.com")
	method2RequestType.Store("google.cloud.orgpolicy.v2.OrgPolicy/CreateCustomConstraint", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.CreateCustomConstraintRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orgpolicy.v2.OrgPolicy/CreateCustomConstraint", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.CustomConstraint)(nil)).Elem())
	method2RequestType.Store("google.cloud.orgpolicy.v2.OrgPolicy/CreatePolicy", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.CreatePolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orgpolicy.v2.OrgPolicy/CreatePolicy", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.Policy)(nil)).Elem())
	method2RequestType.Store("google.cloud.orgpolicy.v2.OrgPolicy/DeleteCustomConstraint", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.DeleteCustomConstraintRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orgpolicy.v2.OrgPolicy/DeleteCustomConstraint", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.orgpolicy.v2.OrgPolicy/DeletePolicy", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.DeletePolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orgpolicy.v2.OrgPolicy/DeletePolicy", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.orgpolicy.v2.OrgPolicy/GetCustomConstraint", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.GetCustomConstraintRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orgpolicy.v2.OrgPolicy/GetCustomConstraint", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.CustomConstraint)(nil)).Elem())
	method2RequestType.Store("google.cloud.orgpolicy.v2.OrgPolicy/GetEffectivePolicy", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.GetEffectivePolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orgpolicy.v2.OrgPolicy/GetEffectivePolicy", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.Policy)(nil)).Elem())
	method2RequestType.Store("google.cloud.orgpolicy.v2.OrgPolicy/GetPolicy", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.GetPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orgpolicy.v2.OrgPolicy/GetPolicy", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.Policy)(nil)).Elem())
	method2RequestType.Store("google.cloud.orgpolicy.v2.OrgPolicy/ListConstraints", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.ListConstraintsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orgpolicy.v2.OrgPolicy/ListConstraints", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.ListConstraintsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.orgpolicy.v2.OrgPolicy/ListCustomConstraints", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.ListCustomConstraintsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orgpolicy.v2.OrgPolicy/ListCustomConstraints", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.ListCustomConstraintsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.orgpolicy.v2.OrgPolicy/ListPolicies", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.ListPoliciesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orgpolicy.v2.OrgPolicy/ListPolicies", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.ListPoliciesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.orgpolicy.v2.OrgPolicy/UpdateCustomConstraint", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.UpdateCustomConstraintRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orgpolicy.v2.OrgPolicy/UpdateCustomConstraint", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.CustomConstraint)(nil)).Elem())
	method2RequestType.Store("google.cloud.orgpolicy.v2.OrgPolicy/UpdatePolicy", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.UpdatePolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orgpolicy.v2.OrgPolicy/UpdatePolicy", reflect.TypeOf((*google_cloud_orgpolicy_v2_orgpolicypb.Policy)(nil)).Elem())
}