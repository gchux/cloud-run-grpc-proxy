// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_iam_v1_iampb "cloud.google.com/go/iam/apiv1/iampb"
	google_storage_v2_storagepb "google.golang.org/genproto/googleapis/storage/v2"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.storage.v2.Storage", "storage.googleapis.com")
	method2RequestType.Store("google.storage.v2.Storage/BidiWriteObject", reflect.TypeOf((*google_storage_v2_storagepb.BidiWriteObjectRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/BidiWriteObject", reflect.TypeOf((*google_storage_v2_storagepb.BidiWriteObjectResponse)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/CancelResumableWrite", reflect.TypeOf((*google_storage_v2_storagepb.CancelResumableWriteRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/CancelResumableWrite", reflect.TypeOf((*google_storage_v2_storagepb.CancelResumableWriteResponse)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/ComposeObject", reflect.TypeOf((*google_storage_v2_storagepb.ComposeObjectRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/ComposeObject", reflect.TypeOf((*google_storage_v2_storagepb.Object)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/CreateBucket", reflect.TypeOf((*google_storage_v2_storagepb.CreateBucketRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/CreateBucket", reflect.TypeOf((*google_storage_v2_storagepb.Bucket)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/CreateHmacKey", reflect.TypeOf((*google_storage_v2_storagepb.CreateHmacKeyRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/CreateHmacKey", reflect.TypeOf((*google_storage_v2_storagepb.CreateHmacKeyResponse)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/CreateNotificationConfig", reflect.TypeOf((*google_storage_v2_storagepb.CreateNotificationConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/CreateNotificationConfig", reflect.TypeOf((*google_storage_v2_storagepb.NotificationConfig)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/DeleteBucket", reflect.TypeOf((*google_storage_v2_storagepb.DeleteBucketRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/DeleteBucket", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/DeleteHmacKey", reflect.TypeOf((*google_storage_v2_storagepb.DeleteHmacKeyRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/DeleteHmacKey", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/DeleteNotificationConfig", reflect.TypeOf((*google_storage_v2_storagepb.DeleteNotificationConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/DeleteNotificationConfig", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/DeleteObject", reflect.TypeOf((*google_storage_v2_storagepb.DeleteObjectRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/DeleteObject", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/GetBucket", reflect.TypeOf((*google_storage_v2_storagepb.GetBucketRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/GetBucket", reflect.TypeOf((*google_storage_v2_storagepb.Bucket)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/GetHmacKey", reflect.TypeOf((*google_storage_v2_storagepb.GetHmacKeyRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/GetHmacKey", reflect.TypeOf((*google_storage_v2_storagepb.HmacKeyMetadata)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.GetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/GetNotificationConfig", reflect.TypeOf((*google_storage_v2_storagepb.GetNotificationConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/GetNotificationConfig", reflect.TypeOf((*google_storage_v2_storagepb.NotificationConfig)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/GetObject", reflect.TypeOf((*google_storage_v2_storagepb.GetObjectRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/GetObject", reflect.TypeOf((*google_storage_v2_storagepb.Object)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/GetServiceAccount", reflect.TypeOf((*google_storage_v2_storagepb.GetServiceAccountRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/GetServiceAccount", reflect.TypeOf((*google_storage_v2_storagepb.ServiceAccount)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/ListBuckets", reflect.TypeOf((*google_storage_v2_storagepb.ListBucketsRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/ListBuckets", reflect.TypeOf((*google_storage_v2_storagepb.ListBucketsResponse)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/ListHmacKeys", reflect.TypeOf((*google_storage_v2_storagepb.ListHmacKeysRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/ListHmacKeys", reflect.TypeOf((*google_storage_v2_storagepb.ListHmacKeysResponse)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/ListNotificationConfigs", reflect.TypeOf((*google_storage_v2_storagepb.ListNotificationConfigsRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/ListNotificationConfigs", reflect.TypeOf((*google_storage_v2_storagepb.ListNotificationConfigsResponse)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/ListObjects", reflect.TypeOf((*google_storage_v2_storagepb.ListObjectsRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/ListObjects", reflect.TypeOf((*google_storage_v2_storagepb.ListObjectsResponse)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/LockBucketRetentionPolicy", reflect.TypeOf((*google_storage_v2_storagepb.LockBucketRetentionPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/LockBucketRetentionPolicy", reflect.TypeOf((*google_storage_v2_storagepb.Bucket)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/QueryWriteStatus", reflect.TypeOf((*google_storage_v2_storagepb.QueryWriteStatusRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/QueryWriteStatus", reflect.TypeOf((*google_storage_v2_storagepb.QueryWriteStatusResponse)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/ReadObject", reflect.TypeOf((*google_storage_v2_storagepb.ReadObjectRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/ReadObject", reflect.TypeOf((*google_storage_v2_storagepb.ReadObjectResponse)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/RestoreObject", reflect.TypeOf((*google_storage_v2_storagepb.RestoreObjectRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/RestoreObject", reflect.TypeOf((*google_storage_v2_storagepb.Object)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/RewriteObject", reflect.TypeOf((*google_storage_v2_storagepb.RewriteObjectRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/RewriteObject", reflect.TypeOf((*google_storage_v2_storagepb.RewriteResponse)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.SetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/StartResumableWrite", reflect.TypeOf((*google_storage_v2_storagepb.StartResumableWriteRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/StartResumableWrite", reflect.TypeOf((*google_storage_v2_storagepb.StartResumableWriteResponse)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsResponse)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/UpdateBucket", reflect.TypeOf((*google_storage_v2_storagepb.UpdateBucketRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/UpdateBucket", reflect.TypeOf((*google_storage_v2_storagepb.Bucket)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/UpdateHmacKey", reflect.TypeOf((*google_storage_v2_storagepb.UpdateHmacKeyRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/UpdateHmacKey", reflect.TypeOf((*google_storage_v2_storagepb.HmacKeyMetadata)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/UpdateObject", reflect.TypeOf((*google_storage_v2_storagepb.UpdateObjectRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/UpdateObject", reflect.TypeOf((*google_storage_v2_storagepb.Object)(nil)).Elem())
	method2RequestType.Store("google.storage.v2.Storage/WriteObject", reflect.TypeOf((*google_storage_v2_storagepb.WriteObjectRequest)(nil)).Elem())
	method2ResponseType.Store("google.storage.v2.Storage/WriteObject", reflect.TypeOf((*google_storage_v2_storagepb.WriteObjectResponse)(nil)).Elem())
}
