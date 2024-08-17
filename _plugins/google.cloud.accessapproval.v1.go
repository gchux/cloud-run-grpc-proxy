// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_accessapproval_v1_accessapprovalpb "cloud.google.com/go/accessapproval/apiv1/accessapprovalpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.accessapproval.v1.AccessApproval", "accessapproval.googleapis.com")
	method2RequestType.Store("google.cloud.accessapproval.v1.AccessApproval/ApproveApprovalRequest", reflect.TypeOf((*google_cloud_accessapproval_v1_accessapprovalpb.ApproveApprovalRequestMessage)(nil)).Elem())
	method2ResponseType.Store("google.cloud.accessapproval.v1.AccessApproval/ApproveApprovalRequest", reflect.TypeOf((*google_cloud_accessapproval_v1_accessapprovalpb.ApprovalRequest)(nil)).Elem())
	method2RequestType.Store("google.cloud.accessapproval.v1.AccessApproval/DeleteAccessApprovalSettings", reflect.TypeOf((*google_cloud_accessapproval_v1_accessapprovalpb.DeleteAccessApprovalSettingsMessage)(nil)).Elem())
	method2ResponseType.Store("google.cloud.accessapproval.v1.AccessApproval/DeleteAccessApprovalSettings", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.accessapproval.v1.AccessApproval/DismissApprovalRequest", reflect.TypeOf((*google_cloud_accessapproval_v1_accessapprovalpb.DismissApprovalRequestMessage)(nil)).Elem())
	method2ResponseType.Store("google.cloud.accessapproval.v1.AccessApproval/DismissApprovalRequest", reflect.TypeOf((*google_cloud_accessapproval_v1_accessapprovalpb.ApprovalRequest)(nil)).Elem())
	method2RequestType.Store("google.cloud.accessapproval.v1.AccessApproval/GetAccessApprovalServiceAccount", reflect.TypeOf((*google_cloud_accessapproval_v1_accessapprovalpb.GetAccessApprovalServiceAccountMessage)(nil)).Elem())
	method2ResponseType.Store("google.cloud.accessapproval.v1.AccessApproval/GetAccessApprovalServiceAccount", reflect.TypeOf((*google_cloud_accessapproval_v1_accessapprovalpb.AccessApprovalServiceAccount)(nil)).Elem())
	method2RequestType.Store("google.cloud.accessapproval.v1.AccessApproval/GetAccessApprovalSettings", reflect.TypeOf((*google_cloud_accessapproval_v1_accessapprovalpb.GetAccessApprovalSettingsMessage)(nil)).Elem())
	method2ResponseType.Store("google.cloud.accessapproval.v1.AccessApproval/GetAccessApprovalSettings", reflect.TypeOf((*google_cloud_accessapproval_v1_accessapprovalpb.AccessApprovalSettings)(nil)).Elem())
	method2RequestType.Store("google.cloud.accessapproval.v1.AccessApproval/GetApprovalRequest", reflect.TypeOf((*google_cloud_accessapproval_v1_accessapprovalpb.GetApprovalRequestMessage)(nil)).Elem())
	method2ResponseType.Store("google.cloud.accessapproval.v1.AccessApproval/GetApprovalRequest", reflect.TypeOf((*google_cloud_accessapproval_v1_accessapprovalpb.ApprovalRequest)(nil)).Elem())
	method2RequestType.Store("google.cloud.accessapproval.v1.AccessApproval/InvalidateApprovalRequest", reflect.TypeOf((*google_cloud_accessapproval_v1_accessapprovalpb.InvalidateApprovalRequestMessage)(nil)).Elem())
	method2ResponseType.Store("google.cloud.accessapproval.v1.AccessApproval/InvalidateApprovalRequest", reflect.TypeOf((*google_cloud_accessapproval_v1_accessapprovalpb.ApprovalRequest)(nil)).Elem())
	method2RequestType.Store("google.cloud.accessapproval.v1.AccessApproval/ListApprovalRequests", reflect.TypeOf((*google_cloud_accessapproval_v1_accessapprovalpb.ListApprovalRequestsMessage)(nil)).Elem())
	method2ResponseType.Store("google.cloud.accessapproval.v1.AccessApproval/ListApprovalRequests", reflect.TypeOf((*google_cloud_accessapproval_v1_accessapprovalpb.ListApprovalRequestsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.accessapproval.v1.AccessApproval/UpdateAccessApprovalSettings", reflect.TypeOf((*google_cloud_accessapproval_v1_accessapprovalpb.UpdateAccessApprovalSettingsMessage)(nil)).Elem())
	method2ResponseType.Store("google.cloud.accessapproval.v1.AccessApproval/UpdateAccessApprovalSettings", reflect.TypeOf((*google_cloud_accessapproval_v1_accessapprovalpb.AccessApprovalSettings)(nil)).Elem())
}