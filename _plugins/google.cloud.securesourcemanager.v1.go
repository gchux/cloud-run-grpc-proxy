// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_iam_v1_iampb "cloud.google.com/go/iam/apiv1/iampb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_securesourcemanager_v1_securesourcemanagerpb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.securesourcemanager.v1.SecureSourceManager", "securesourcemanager.googleapis.com")
	method2RequestType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/CreateInstance", reflect.TypeOf((*google_cloud_securesourcemanager_v1_securesourcemanagerpb.CreateInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/CreateInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/CreateRepository", reflect.TypeOf((*google_cloud_securesourcemanager_v1_securesourcemanagerpb.CreateRepositoryRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/CreateRepository", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/DeleteInstance", reflect.TypeOf((*google_cloud_securesourcemanager_v1_securesourcemanagerpb.DeleteInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/DeleteInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/DeleteRepository", reflect.TypeOf((*google_cloud_securesourcemanager_v1_securesourcemanagerpb.DeleteRepositoryRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/DeleteRepository", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/GetIamPolicyRepo", reflect.TypeOf((*google_iam_v1_iampb.GetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/GetIamPolicyRepo", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/GetInstance", reflect.TypeOf((*google_cloud_securesourcemanager_v1_securesourcemanagerpb.GetInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/GetInstance", reflect.TypeOf((*google_cloud_securesourcemanager_v1_securesourcemanagerpb.Instance)(nil)).Elem())
	method2RequestType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/GetRepository", reflect.TypeOf((*google_cloud_securesourcemanager_v1_securesourcemanagerpb.GetRepositoryRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/GetRepository", reflect.TypeOf((*google_cloud_securesourcemanager_v1_securesourcemanagerpb.Repository)(nil)).Elem())
	method2RequestType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/ListInstances", reflect.TypeOf((*google_cloud_securesourcemanager_v1_securesourcemanagerpb.ListInstancesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/ListInstances", reflect.TypeOf((*google_cloud_securesourcemanager_v1_securesourcemanagerpb.ListInstancesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/ListRepositories", reflect.TypeOf((*google_cloud_securesourcemanager_v1_securesourcemanagerpb.ListRepositoriesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/ListRepositories", reflect.TypeOf((*google_cloud_securesourcemanager_v1_securesourcemanagerpb.ListRepositoriesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/SetIamPolicyRepo", reflect.TypeOf((*google_iam_v1_iampb.SetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/SetIamPolicyRepo", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/TestIamPermissionsRepo", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securesourcemanager.v1.SecureSourceManager/TestIamPermissionsRepo", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsResponse)(nil)).Elem())
}
