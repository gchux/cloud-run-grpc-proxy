// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_redis_cluster_v1_clusterpb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.redis.cluster.v1.CloudRedisCluster", "redis.googleapis.com")
	method2RequestType.Store("google.cloud.redis.cluster.v1.CloudRedisCluster/CreateCluster", reflect.TypeOf((*google_cloud_redis_cluster_v1_clusterpb.CreateClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.redis.cluster.v1.CloudRedisCluster/CreateCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.redis.cluster.v1.CloudRedisCluster/DeleteCluster", reflect.TypeOf((*google_cloud_redis_cluster_v1_clusterpb.DeleteClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.redis.cluster.v1.CloudRedisCluster/DeleteCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.redis.cluster.v1.CloudRedisCluster/GetCluster", reflect.TypeOf((*google_cloud_redis_cluster_v1_clusterpb.GetClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.redis.cluster.v1.CloudRedisCluster/GetCluster", reflect.TypeOf((*google_cloud_redis_cluster_v1_clusterpb.Cluster)(nil)).Elem())
	method2RequestType.Store("google.cloud.redis.cluster.v1.CloudRedisCluster/GetClusterCertificateAuthority", reflect.TypeOf((*google_cloud_redis_cluster_v1_clusterpb.GetClusterCertificateAuthorityRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.redis.cluster.v1.CloudRedisCluster/GetClusterCertificateAuthority", reflect.TypeOf((*google_cloud_redis_cluster_v1_clusterpb.CertificateAuthority)(nil)).Elem())
	method2RequestType.Store("google.cloud.redis.cluster.v1.CloudRedisCluster/ListClusters", reflect.TypeOf((*google_cloud_redis_cluster_v1_clusterpb.ListClustersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.redis.cluster.v1.CloudRedisCluster/ListClusters", reflect.TypeOf((*google_cloud_redis_cluster_v1_clusterpb.ListClustersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.redis.cluster.v1.CloudRedisCluster/UpdateCluster", reflect.TypeOf((*google_cloud_redis_cluster_v1_clusterpb.UpdateClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.redis.cluster.v1.CloudRedisCluster/UpdateCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
