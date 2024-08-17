// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_asset_v1p2beta1_assetpb "cloud.google.com/go/asset/apiv1p2beta1/assetpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.asset.v1p2beta1.AssetService", "cloudasset.googleapis.com")
	method2RequestType.Store("google.cloud.asset.v1p2beta1.AssetService/CreateFeed", reflect.TypeOf((*google_cloud_asset_v1p2beta1_assetpb.CreateFeedRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.asset.v1p2beta1.AssetService/CreateFeed", reflect.TypeOf((*google_cloud_asset_v1p2beta1_assetpb.Feed)(nil)).Elem())
	method2RequestType.Store("google.cloud.asset.v1p2beta1.AssetService/DeleteFeed", reflect.TypeOf((*google_cloud_asset_v1p2beta1_assetpb.DeleteFeedRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.asset.v1p2beta1.AssetService/DeleteFeed", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.asset.v1p2beta1.AssetService/GetFeed", reflect.TypeOf((*google_cloud_asset_v1p2beta1_assetpb.GetFeedRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.asset.v1p2beta1.AssetService/GetFeed", reflect.TypeOf((*google_cloud_asset_v1p2beta1_assetpb.Feed)(nil)).Elem())
	method2RequestType.Store("google.cloud.asset.v1p2beta1.AssetService/ListFeeds", reflect.TypeOf((*google_cloud_asset_v1p2beta1_assetpb.ListFeedsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.asset.v1p2beta1.AssetService/ListFeeds", reflect.TypeOf((*google_cloud_asset_v1p2beta1_assetpb.ListFeedsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.asset.v1p2beta1.AssetService/UpdateFeed", reflect.TypeOf((*google_cloud_asset_v1p2beta1_assetpb.UpdateFeedRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.asset.v1p2beta1.AssetService/UpdateFeed", reflect.TypeOf((*google_cloud_asset_v1p2beta1_assetpb.Feed)(nil)).Elem())
}