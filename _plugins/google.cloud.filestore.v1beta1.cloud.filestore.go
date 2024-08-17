// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_filestore_v1beta1_cloud_filestore_servicepb "google.golang.org/genproto/googleapis/cloud/filestore/v1beta1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager", "file.googleapis.com")
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/CreateBackup", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.CreateBackupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/CreateBackup", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/CreateInstance", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.CreateInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/CreateInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/CreateShare", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.CreateShareRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/CreateShare", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/CreateSnapshot", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.CreateSnapshotRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/CreateSnapshot", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/DeleteBackup", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.DeleteBackupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/DeleteBackup", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/DeleteInstance", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.DeleteInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/DeleteInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/DeleteShare", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.DeleteShareRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/DeleteShare", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/DeleteSnapshot", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.DeleteSnapshotRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/DeleteSnapshot", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/GetBackup", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.GetBackupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/GetBackup", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.Backup)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/GetInstance", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.GetInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/GetInstance", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.Instance)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/GetShare", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.GetShareRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/GetShare", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.Share)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/GetSnapshot", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.GetSnapshotRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/GetSnapshot", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.Snapshot)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/ListBackups", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.ListBackupsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/ListBackups", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.ListBackupsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/ListInstances", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.ListInstancesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/ListInstances", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.ListInstancesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/ListShares", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.ListSharesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/ListShares", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.ListSharesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/ListSnapshots", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.ListSnapshotsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/ListSnapshots", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.ListSnapshotsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/RestoreInstance", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.RestoreInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/RestoreInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/RevertInstance", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.RevertInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/RevertInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/UpdateBackup", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.UpdateBackupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/UpdateBackup", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/UpdateInstance", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.UpdateInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/UpdateInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/UpdateShare", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.UpdateShareRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/UpdateShare", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/UpdateSnapshot", reflect.TypeOf((*google_cloud_filestore_v1beta1_cloud_filestore_servicepb.UpdateSnapshotRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.filestore.v1beta1.CloudFilestoreManager/UpdateSnapshot", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
