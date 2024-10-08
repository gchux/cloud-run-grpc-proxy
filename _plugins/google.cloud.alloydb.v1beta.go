// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_alloydb_v1beta_alloydbpb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin", "alloydb.googleapis.com")
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/BatchCreateInstances", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.BatchCreateInstancesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/BatchCreateInstances", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/CreateBackup", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.CreateBackupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/CreateBackup", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/CreateCluster", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.CreateClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/CreateCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/CreateInstance", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.CreateInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/CreateInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/CreateSecondaryCluster", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.CreateSecondaryClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/CreateSecondaryCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/CreateSecondaryInstance", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.CreateSecondaryInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/CreateSecondaryInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/CreateUser", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.CreateUserRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/CreateUser", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.User)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/DeleteBackup", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.DeleteBackupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/DeleteBackup", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/DeleteCluster", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.DeleteClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/DeleteCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/DeleteInstance", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.DeleteInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/DeleteInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/DeleteUser", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.DeleteUserRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/DeleteUser", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/FailoverInstance", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.FailoverInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/FailoverInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/GenerateClientCertificate", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.GenerateClientCertificateRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/GenerateClientCertificate", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.GenerateClientCertificateResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/GetBackup", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.GetBackupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/GetBackup", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.Backup)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/GetCluster", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.GetClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/GetCluster", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.Cluster)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/GetConnectionInfo", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.GetConnectionInfoRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/GetConnectionInfo", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.ConnectionInfo)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/GetInstance", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.GetInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/GetInstance", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.Instance)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/GetUser", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.GetUserRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/GetUser", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.User)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/InjectFault", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.InjectFaultRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/InjectFault", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/ListBackups", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.ListBackupsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/ListBackups", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.ListBackupsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/ListClusters", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.ListClustersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/ListClusters", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.ListClustersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/ListDatabases", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.ListDatabasesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/ListDatabases", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.ListDatabasesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/ListInstances", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.ListInstancesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/ListInstances", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.ListInstancesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/ListSupportedDatabaseFlags", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.ListSupportedDatabaseFlagsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/ListSupportedDatabaseFlags", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.ListSupportedDatabaseFlagsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/ListUsers", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.ListUsersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/ListUsers", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.ListUsersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/PromoteCluster", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.PromoteClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/PromoteCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/RestartInstance", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.RestartInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/RestartInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/RestoreCluster", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.RestoreClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/RestoreCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/UpdateBackup", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.UpdateBackupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/UpdateBackup", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/UpdateCluster", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.UpdateClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/UpdateCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/UpdateInstance", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.UpdateInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/UpdateInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/UpdateUser", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.UpdateUserRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.alloydb.v1beta.AlloyDBAdmin/UpdateUser", reflect.TypeOf((*google_cloud_alloydb_v1beta_alloydbpb.User)(nil)).Elem())
}
