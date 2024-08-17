// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_asset_v1p7beta1_asset_servicepb "google.golang.org/genproto/googleapis/cloud/asset/v1p7beta1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.asset.v1p7beta1.AssetService", "cloudasset.googleapis.com")
	method2RequestType.Store("google.cloud.asset.v1p7beta1.AssetService/ExportAssets", reflect.TypeOf((*google_cloud_asset_v1p7beta1_asset_servicepb.ExportAssetsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.asset.v1p7beta1.AssetService/ExportAssets", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
