// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_memcache_v1_memcachepb "cloud.google.com/go/memcache/apiv1/memcachepb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.memcache.v1.CloudMemcache", "memcache.googleapis.com")
	method2RequestType.Store("google.cloud.memcache.v1.CloudMemcache/ApplyParameters", reflect.TypeOf((*google_cloud_memcache_v1_memcachepb.ApplyParametersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.memcache.v1.CloudMemcache/ApplyParameters", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.memcache.v1.CloudMemcache/CreateInstance", reflect.TypeOf((*google_cloud_memcache_v1_memcachepb.CreateInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.memcache.v1.CloudMemcache/CreateInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.memcache.v1.CloudMemcache/DeleteInstance", reflect.TypeOf((*google_cloud_memcache_v1_memcachepb.DeleteInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.memcache.v1.CloudMemcache/DeleteInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.memcache.v1.CloudMemcache/GetInstance", reflect.TypeOf((*google_cloud_memcache_v1_memcachepb.GetInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.memcache.v1.CloudMemcache/GetInstance", reflect.TypeOf((*google_cloud_memcache_v1_memcachepb.Instance)(nil)).Elem())
	method2RequestType.Store("google.cloud.memcache.v1.CloudMemcache/ListInstances", reflect.TypeOf((*google_cloud_memcache_v1_memcachepb.ListInstancesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.memcache.v1.CloudMemcache/ListInstances", reflect.TypeOf((*google_cloud_memcache_v1_memcachepb.ListInstancesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.memcache.v1.CloudMemcache/RescheduleMaintenance", reflect.TypeOf((*google_cloud_memcache_v1_memcachepb.RescheduleMaintenanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.memcache.v1.CloudMemcache/RescheduleMaintenance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.memcache.v1.CloudMemcache/UpdateInstance", reflect.TypeOf((*google_cloud_memcache_v1_memcachepb.UpdateInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.memcache.v1.CloudMemcache/UpdateInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.memcache.v1.CloudMemcache/UpdateParameters", reflect.TypeOf((*google_cloud_memcache_v1_memcachepb.UpdateParametersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.memcache.v1.CloudMemcache/UpdateParameters", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
