// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_kms_v1_kmspb "cloud.google.com/go/kms/apiv1/kmspb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.kms.v1.EkmService", "cloudkms.googleapis.com")
	method2RequestType.Store("google.cloud.kms.v1.EkmService/CreateEkmConnection", reflect.TypeOf((*google_cloud_kms_v1_kmspb.CreateEkmConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.kms.v1.EkmService/CreateEkmConnection", reflect.TypeOf((*google_cloud_kms_v1_kmspb.EkmConnection)(nil)).Elem())
	method2RequestType.Store("google.cloud.kms.v1.EkmService/GetEkmConfig", reflect.TypeOf((*google_cloud_kms_v1_kmspb.GetEkmConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.kms.v1.EkmService/GetEkmConfig", reflect.TypeOf((*google_cloud_kms_v1_kmspb.EkmConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.kms.v1.EkmService/GetEkmConnection", reflect.TypeOf((*google_cloud_kms_v1_kmspb.GetEkmConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.kms.v1.EkmService/GetEkmConnection", reflect.TypeOf((*google_cloud_kms_v1_kmspb.EkmConnection)(nil)).Elem())
	method2RequestType.Store("google.cloud.kms.v1.EkmService/ListEkmConnections", reflect.TypeOf((*google_cloud_kms_v1_kmspb.ListEkmConnectionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.kms.v1.EkmService/ListEkmConnections", reflect.TypeOf((*google_cloud_kms_v1_kmspb.ListEkmConnectionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.kms.v1.EkmService/UpdateEkmConfig", reflect.TypeOf((*google_cloud_kms_v1_kmspb.UpdateEkmConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.kms.v1.EkmService/UpdateEkmConfig", reflect.TypeOf((*google_cloud_kms_v1_kmspb.EkmConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.kms.v1.EkmService/UpdateEkmConnection", reflect.TypeOf((*google_cloud_kms_v1_kmspb.UpdateEkmConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.kms.v1.EkmService/UpdateEkmConnection", reflect.TypeOf((*google_cloud_kms_v1_kmspb.EkmConnection)(nil)).Elem())
	method2RequestType.Store("google.cloud.kms.v1.EkmService/VerifyConnectivity", reflect.TypeOf((*google_cloud_kms_v1_kmspb.VerifyConnectivityRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.kms.v1.EkmService/VerifyConnectivity", reflect.TypeOf((*google_cloud_kms_v1_kmspb.VerifyConnectivityResponse)(nil)).Elem())
}
