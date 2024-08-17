// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_gkemulticloud_v1_gkemulticloudpb "cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.gkemulticloud.v1.AttachedClusters", "gkemulticloud.googleapis.com")
	method2RequestType.Store("google.cloud.gkemulticloud.v1.AttachedClusters/CreateAttachedCluster", reflect.TypeOf((*google_cloud_gkemulticloud_v1_gkemulticloudpb.CreateAttachedClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gkemulticloud.v1.AttachedClusters/CreateAttachedCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.gkemulticloud.v1.AttachedClusters/DeleteAttachedCluster", reflect.TypeOf((*google_cloud_gkemulticloud_v1_gkemulticloudpb.DeleteAttachedClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gkemulticloud.v1.AttachedClusters/DeleteAttachedCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.gkemulticloud.v1.AttachedClusters/GetAttachedCluster", reflect.TypeOf((*google_cloud_gkemulticloud_v1_gkemulticloudpb.GetAttachedClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gkemulticloud.v1.AttachedClusters/GetAttachedCluster", reflect.TypeOf((*google_cloud_gkemulticloud_v1_gkemulticloudpb.AttachedCluster)(nil)).Elem())
	method2RequestType.Store("google.cloud.gkemulticloud.v1.AttachedClusters/GetAttachedServerConfig", reflect.TypeOf((*google_cloud_gkemulticloud_v1_gkemulticloudpb.GetAttachedServerConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gkemulticloud.v1.AttachedClusters/GetAttachedServerConfig", reflect.TypeOf((*google_cloud_gkemulticloud_v1_gkemulticloudpb.AttachedServerConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.gkemulticloud.v1.AttachedClusters/ImportAttachedCluster", reflect.TypeOf((*google_cloud_gkemulticloud_v1_gkemulticloudpb.ImportAttachedClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gkemulticloud.v1.AttachedClusters/ImportAttachedCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.gkemulticloud.v1.AttachedClusters/ListAttachedClusters", reflect.TypeOf((*google_cloud_gkemulticloud_v1_gkemulticloudpb.ListAttachedClustersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gkemulticloud.v1.AttachedClusters/ListAttachedClusters", reflect.TypeOf((*google_cloud_gkemulticloud_v1_gkemulticloudpb.ListAttachedClustersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.gkemulticloud.v1.AttachedClusters/UpdateAttachedCluster", reflect.TypeOf((*google_cloud_gkemulticloud_v1_gkemulticloudpb.UpdateAttachedClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gkemulticloud.v1.AttachedClusters/UpdateAttachedCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
