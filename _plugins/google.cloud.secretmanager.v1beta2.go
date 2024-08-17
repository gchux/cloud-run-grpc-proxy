// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_iam_v1_iampb "cloud.google.com/go/iam/apiv1/iampb"
	google_cloud_secretmanager_v1beta2_secretmanagerpb "cloud.google.com/go/secretmanager/apiv1beta2/secretmanagerpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.secretmanager.v1beta2.SecretManagerService", "secretmanager.googleapis.com")
	method2RequestType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/AccessSecretVersion", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.AccessSecretVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/AccessSecretVersion", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.AccessSecretVersionResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/AddSecretVersion", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.AddSecretVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/AddSecretVersion", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.SecretVersion)(nil)).Elem())
	method2RequestType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/CreateSecret", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.CreateSecretRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/CreateSecret", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.Secret)(nil)).Elem())
	method2RequestType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/DeleteSecret", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.DeleteSecretRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/DeleteSecret", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/DestroySecretVersion", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.DestroySecretVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/DestroySecretVersion", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.SecretVersion)(nil)).Elem())
	method2RequestType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/DisableSecretVersion", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.DisableSecretVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/DisableSecretVersion", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.SecretVersion)(nil)).Elem())
	method2RequestType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/EnableSecretVersion", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.EnableSecretVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/EnableSecretVersion", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.SecretVersion)(nil)).Elem())
	method2RequestType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.GetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/GetSecret", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.GetSecretRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/GetSecret", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.Secret)(nil)).Elem())
	method2RequestType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/GetSecretVersion", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.GetSecretVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/GetSecretVersion", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.SecretVersion)(nil)).Elem())
	method2RequestType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/ListSecretVersions", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.ListSecretVersionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/ListSecretVersions", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.ListSecretVersionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/ListSecrets", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.ListSecretsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/ListSecrets", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.ListSecretsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.SetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/UpdateSecret", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.UpdateSecretRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.secretmanager.v1beta2.SecretManagerService/UpdateSecret", reflect.TypeOf((*google_cloud_secretmanager_v1beta2_secretmanagerpb.Secret)(nil)).Elem())
}
