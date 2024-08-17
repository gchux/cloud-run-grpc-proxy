// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_ai_generativelanguage_v1beta_generativelanguagepb "cloud.google.com/go/ai/generativelanguage/apiv1beta/generativelanguagepb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.ai.generativelanguage.v1beta.PermissionService", "generativelanguage.googleapis.com")
	method2RequestType.Store("google.ai.generativelanguage.v1beta.PermissionService/CreatePermission", reflect.TypeOf((*google_ai_generativelanguage_v1beta_generativelanguagepb.CreatePermissionRequest)(nil)).Elem())
	method2ResponseType.Store("google.ai.generativelanguage.v1beta.PermissionService/CreatePermission", reflect.TypeOf((*google_ai_generativelanguage_v1beta_generativelanguagepb.Permission)(nil)).Elem())
	method2RequestType.Store("google.ai.generativelanguage.v1beta.PermissionService/DeletePermission", reflect.TypeOf((*google_ai_generativelanguage_v1beta_generativelanguagepb.DeletePermissionRequest)(nil)).Elem())
	method2ResponseType.Store("google.ai.generativelanguage.v1beta.PermissionService/DeletePermission", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.ai.generativelanguage.v1beta.PermissionService/GetPermission", reflect.TypeOf((*google_ai_generativelanguage_v1beta_generativelanguagepb.GetPermissionRequest)(nil)).Elem())
	method2ResponseType.Store("google.ai.generativelanguage.v1beta.PermissionService/GetPermission", reflect.TypeOf((*google_ai_generativelanguage_v1beta_generativelanguagepb.Permission)(nil)).Elem())
	method2RequestType.Store("google.ai.generativelanguage.v1beta.PermissionService/ListPermissions", reflect.TypeOf((*google_ai_generativelanguage_v1beta_generativelanguagepb.ListPermissionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.ai.generativelanguage.v1beta.PermissionService/ListPermissions", reflect.TypeOf((*google_ai_generativelanguage_v1beta_generativelanguagepb.ListPermissionsResponse)(nil)).Elem())
	method2RequestType.Store("google.ai.generativelanguage.v1beta.PermissionService/TransferOwnership", reflect.TypeOf((*google_ai_generativelanguage_v1beta_generativelanguagepb.TransferOwnershipRequest)(nil)).Elem())
	method2ResponseType.Store("google.ai.generativelanguage.v1beta.PermissionService/TransferOwnership", reflect.TypeOf((*google_ai_generativelanguage_v1beta_generativelanguagepb.TransferOwnershipResponse)(nil)).Elem())
	method2RequestType.Store("google.ai.generativelanguage.v1beta.PermissionService/UpdatePermission", reflect.TypeOf((*google_ai_generativelanguage_v1beta_generativelanguagepb.UpdatePermissionRequest)(nil)).Elem())
	method2ResponseType.Store("google.ai.generativelanguage.v1beta.PermissionService/UpdatePermission", reflect.TypeOf((*google_ai_generativelanguage_v1beta_generativelanguagepb.Permission)(nil)).Elem())
}