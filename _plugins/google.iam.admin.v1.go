// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_iam_admin_v1_adminpb "cloud.google.com/go/iam/admin/apiv1/adminpb"
	google_iam_v1_iampb "cloud.google.com/go/iam/apiv1/iampb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.iam.admin.v1.IAM", "iam.googleapis.com")
	method2RequestType.Store("google.iam.admin.v1.IAM/CreateRole", reflect.TypeOf((*google_iam_admin_v1_adminpb.CreateRoleRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/CreateRole", reflect.TypeOf((*google_iam_admin_v1_adminpb.Role)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/CreateServiceAccount", reflect.TypeOf((*google_iam_admin_v1_adminpb.CreateServiceAccountRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/CreateServiceAccount", reflect.TypeOf((*google_iam_admin_v1_adminpb.ServiceAccount)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/CreateServiceAccountKey", reflect.TypeOf((*google_iam_admin_v1_adminpb.CreateServiceAccountKeyRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/CreateServiceAccountKey", reflect.TypeOf((*google_iam_admin_v1_adminpb.ServiceAccountKey)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/DeleteRole", reflect.TypeOf((*google_iam_admin_v1_adminpb.DeleteRoleRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/DeleteRole", reflect.TypeOf((*google_iam_admin_v1_adminpb.Role)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/DeleteServiceAccount", reflect.TypeOf((*google_iam_admin_v1_adminpb.DeleteServiceAccountRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/DeleteServiceAccount", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/DeleteServiceAccountKey", reflect.TypeOf((*google_iam_admin_v1_adminpb.DeleteServiceAccountKeyRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/DeleteServiceAccountKey", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/DisableServiceAccount", reflect.TypeOf((*google_iam_admin_v1_adminpb.DisableServiceAccountRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/DisableServiceAccount", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/DisableServiceAccountKey", reflect.TypeOf((*google_iam_admin_v1_adminpb.DisableServiceAccountKeyRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/DisableServiceAccountKey", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/EnableServiceAccount", reflect.TypeOf((*google_iam_admin_v1_adminpb.EnableServiceAccountRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/EnableServiceAccount", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/EnableServiceAccountKey", reflect.TypeOf((*google_iam_admin_v1_adminpb.EnableServiceAccountKeyRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/EnableServiceAccountKey", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.GetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/GetRole", reflect.TypeOf((*google_iam_admin_v1_adminpb.GetRoleRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/GetRole", reflect.TypeOf((*google_iam_admin_v1_adminpb.Role)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/GetServiceAccount", reflect.TypeOf((*google_iam_admin_v1_adminpb.GetServiceAccountRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/GetServiceAccount", reflect.TypeOf((*google_iam_admin_v1_adminpb.ServiceAccount)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/GetServiceAccountKey", reflect.TypeOf((*google_iam_admin_v1_adminpb.GetServiceAccountKeyRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/GetServiceAccountKey", reflect.TypeOf((*google_iam_admin_v1_adminpb.ServiceAccountKey)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/LintPolicy", reflect.TypeOf((*google_iam_admin_v1_adminpb.LintPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/LintPolicy", reflect.TypeOf((*google_iam_admin_v1_adminpb.LintPolicyResponse)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/ListRoles", reflect.TypeOf((*google_iam_admin_v1_adminpb.ListRolesRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/ListRoles", reflect.TypeOf((*google_iam_admin_v1_adminpb.ListRolesResponse)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/ListServiceAccountKeys", reflect.TypeOf((*google_iam_admin_v1_adminpb.ListServiceAccountKeysRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/ListServiceAccountKeys", reflect.TypeOf((*google_iam_admin_v1_adminpb.ListServiceAccountKeysResponse)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/ListServiceAccounts", reflect.TypeOf((*google_iam_admin_v1_adminpb.ListServiceAccountsRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/ListServiceAccounts", reflect.TypeOf((*google_iam_admin_v1_adminpb.ListServiceAccountsResponse)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/PatchServiceAccount", reflect.TypeOf((*google_iam_admin_v1_adminpb.PatchServiceAccountRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/PatchServiceAccount", reflect.TypeOf((*google_iam_admin_v1_adminpb.ServiceAccount)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/QueryAuditableServices", reflect.TypeOf((*google_iam_admin_v1_adminpb.QueryAuditableServicesRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/QueryAuditableServices", reflect.TypeOf((*google_iam_admin_v1_adminpb.QueryAuditableServicesResponse)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/QueryGrantableRoles", reflect.TypeOf((*google_iam_admin_v1_adminpb.QueryGrantableRolesRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/QueryGrantableRoles", reflect.TypeOf((*google_iam_admin_v1_adminpb.QueryGrantableRolesResponse)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/QueryTestablePermissions", reflect.TypeOf((*google_iam_admin_v1_adminpb.QueryTestablePermissionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/QueryTestablePermissions", reflect.TypeOf((*google_iam_admin_v1_adminpb.QueryTestablePermissionsResponse)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.SetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/SignBlob", reflect.TypeOf((*google_iam_admin_v1_adminpb.SignBlobRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/SignBlob", reflect.TypeOf((*google_iam_admin_v1_adminpb.SignBlobResponse)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/SignJwt", reflect.TypeOf((*google_iam_admin_v1_adminpb.SignJwtRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/SignJwt", reflect.TypeOf((*google_iam_admin_v1_adminpb.SignJwtResponse)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsResponse)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/UndeleteRole", reflect.TypeOf((*google_iam_admin_v1_adminpb.UndeleteRoleRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/UndeleteRole", reflect.TypeOf((*google_iam_admin_v1_adminpb.Role)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/UndeleteServiceAccount", reflect.TypeOf((*google_iam_admin_v1_adminpb.UndeleteServiceAccountRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/UndeleteServiceAccount", reflect.TypeOf((*google_iam_admin_v1_adminpb.UndeleteServiceAccountResponse)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/UpdateRole", reflect.TypeOf((*google_iam_admin_v1_adminpb.UpdateRoleRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/UpdateRole", reflect.TypeOf((*google_iam_admin_v1_adminpb.Role)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/UpdateServiceAccount", reflect.TypeOf((*google_iam_admin_v1_adminpb.ServiceAccount)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/UpdateServiceAccount", reflect.TypeOf((*google_iam_admin_v1_adminpb.ServiceAccount)(nil)).Elem())
	method2RequestType.Store("google.iam.admin.v1.IAM/UploadServiceAccountKey", reflect.TypeOf((*google_iam_admin_v1_adminpb.UploadServiceAccountKeyRequest)(nil)).Elem())
	method2ResponseType.Store("google.iam.admin.v1.IAM/UploadServiceAccountKey", reflect.TypeOf((*google_iam_admin_v1_adminpb.ServiceAccountKey)(nil)).Elem())
}
