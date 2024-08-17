// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_memcache_v1beta2_memcachepb "cloud.google.com/go/memcache/apiv1beta2/memcachepb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.memcache.v1beta2.CloudMemcache", "memcache.googleapis.com")
	method2RequestType.Store("google.cloud.memcache.v1beta2.CloudMemcache/ApplyParameters", reflect.TypeOf((*google_cloud_memcache_v1beta2_memcachepb.ApplyParametersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.memcache.v1beta2.CloudMemcache/ApplyParameters", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.memcache.v1beta2.CloudMemcache/ApplySoftwareUpdate", reflect.TypeOf((*google_cloud_memcache_v1beta2_memcachepb.ApplySoftwareUpdateRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.memcache.v1beta2.CloudMemcache/ApplySoftwareUpdate", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.memcache.v1beta2.CloudMemcache/CreateInstance", reflect.TypeOf((*google_cloud_memcache_v1beta2_memcachepb.CreateInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.memcache.v1beta2.CloudMemcache/CreateInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.memcache.v1beta2.CloudMemcache/DeleteInstance", reflect.TypeOf((*google_cloud_memcache_v1beta2_memcachepb.DeleteInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.memcache.v1beta2.CloudMemcache/DeleteInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.memcache.v1beta2.CloudMemcache/GetInstance", reflect.TypeOf((*google_cloud_memcache_v1beta2_memcachepb.GetInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.memcache.v1beta2.CloudMemcache/GetInstance", reflect.TypeOf((*google_cloud_memcache_v1beta2_memcachepb.Instance)(nil)).Elem())
	method2RequestType.Store("google.cloud.memcache.v1beta2.CloudMemcache/ListInstances", reflect.TypeOf((*google_cloud_memcache_v1beta2_memcachepb.ListInstancesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.memcache.v1beta2.CloudMemcache/ListInstances", reflect.TypeOf((*google_cloud_memcache_v1beta2_memcachepb.ListInstancesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.memcache.v1beta2.CloudMemcache/RescheduleMaintenance", reflect.TypeOf((*google_cloud_memcache_v1beta2_memcachepb.RescheduleMaintenanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.memcache.v1beta2.CloudMemcache/RescheduleMaintenance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.memcache.v1beta2.CloudMemcache/UpdateInstance", reflect.TypeOf((*google_cloud_memcache_v1beta2_memcachepb.UpdateInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.memcache.v1beta2.CloudMemcache/UpdateInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.memcache.v1beta2.CloudMemcache/UpdateParameters", reflect.TypeOf((*google_cloud_memcache_v1beta2_memcachepb.UpdateParametersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.memcache.v1beta2.CloudMemcache/UpdateParameters", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}