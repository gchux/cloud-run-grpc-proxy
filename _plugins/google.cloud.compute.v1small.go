// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_compute_v1small_compute "google.golang.org/genproto/googleapis/cloud/compute/v1small"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.compute.v1small.Addresses", "compute.googleapis.com")
	service2Host.Store("google.cloud.compute.v1small.RegionOperations", "compute.googleapis.com")
	method2RequestType.Store("google.cloud.compute.v1small.Addresses/AggregatedList", reflect.TypeOf((*google_cloud_compute_v1small_compute.AggregatedListAddressesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.compute.v1small.Addresses/AggregatedList", reflect.TypeOf((*google_cloud_compute_v1small_compute.AddressAggregatedList)(nil)).Elem())
	method2RequestType.Store("google.cloud.compute.v1small.Addresses/Delete", reflect.TypeOf((*google_cloud_compute_v1small_compute.DeleteAddressRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.compute.v1small.Addresses/Delete", reflect.TypeOf((*google_cloud_compute_v1small_compute.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.compute.v1small.Addresses/Insert", reflect.TypeOf((*google_cloud_compute_v1small_compute.InsertAddressRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.compute.v1small.Addresses/Insert", reflect.TypeOf((*google_cloud_compute_v1small_compute.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.compute.v1small.Addresses/List", reflect.TypeOf((*google_cloud_compute_v1small_compute.ListAddressesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.compute.v1small.Addresses/List", reflect.TypeOf((*google_cloud_compute_v1small_compute.AddressList)(nil)).Elem())
	method2RequestType.Store("google.cloud.compute.v1small.RegionOperations/Get", reflect.TypeOf((*google_cloud_compute_v1small_compute.GetRegionOperationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.compute.v1small.RegionOperations/Get", reflect.TypeOf((*google_cloud_compute_v1small_compute.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.compute.v1small.RegionOperations/Wait", reflect.TypeOf((*google_cloud_compute_v1small_compute.WaitRegionOperationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.compute.v1small.RegionOperations/Wait", reflect.TypeOf((*google_cloud_compute_v1small_compute.Operation)(nil)).Elem())
}
