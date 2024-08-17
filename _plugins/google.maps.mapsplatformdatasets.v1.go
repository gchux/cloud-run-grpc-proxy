// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_maps_mapsplatformdatasets_v1_mapsplatformdatasetspb "cloud.google.com/go/maps/mapsplatformdatasets/apiv1/mapsplatformdatasetspb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.maps.mapsplatformdatasets.v1.MapsPlatformDatasets", "mapsplatformdatasets.googleapis.com")
	method2RequestType.Store("google.maps.mapsplatformdatasets.v1.MapsPlatformDatasets/CreateDataset", reflect.TypeOf((*google_maps_mapsplatformdatasets_v1_mapsplatformdatasetspb.CreateDatasetRequest)(nil)).Elem())
	method2ResponseType.Store("google.maps.mapsplatformdatasets.v1.MapsPlatformDatasets/CreateDataset", reflect.TypeOf((*google_maps_mapsplatformdatasets_v1_mapsplatformdatasetspb.Dataset)(nil)).Elem())
	method2RequestType.Store("google.maps.mapsplatformdatasets.v1.MapsPlatformDatasets/DeleteDataset", reflect.TypeOf((*google_maps_mapsplatformdatasets_v1_mapsplatformdatasetspb.DeleteDatasetRequest)(nil)).Elem())
	method2ResponseType.Store("google.maps.mapsplatformdatasets.v1.MapsPlatformDatasets/DeleteDataset", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.maps.mapsplatformdatasets.v1.MapsPlatformDatasets/FetchDatasetErrors", reflect.TypeOf((*google_maps_mapsplatformdatasets_v1_mapsplatformdatasetspb.FetchDatasetErrorsRequest)(nil)).Elem())
	method2ResponseType.Store("google.maps.mapsplatformdatasets.v1.MapsPlatformDatasets/FetchDatasetErrors", reflect.TypeOf((*google_maps_mapsplatformdatasets_v1_mapsplatformdatasetspb.FetchDatasetErrorsResponse)(nil)).Elem())
	method2RequestType.Store("google.maps.mapsplatformdatasets.v1.MapsPlatformDatasets/GetDataset", reflect.TypeOf((*google_maps_mapsplatformdatasets_v1_mapsplatformdatasetspb.GetDatasetRequest)(nil)).Elem())
	method2ResponseType.Store("google.maps.mapsplatformdatasets.v1.MapsPlatformDatasets/GetDataset", reflect.TypeOf((*google_maps_mapsplatformdatasets_v1_mapsplatformdatasetspb.Dataset)(nil)).Elem())
	method2RequestType.Store("google.maps.mapsplatformdatasets.v1.MapsPlatformDatasets/ListDatasets", reflect.TypeOf((*google_maps_mapsplatformdatasets_v1_mapsplatformdatasetspb.ListDatasetsRequest)(nil)).Elem())
	method2ResponseType.Store("google.maps.mapsplatformdatasets.v1.MapsPlatformDatasets/ListDatasets", reflect.TypeOf((*google_maps_mapsplatformdatasets_v1_mapsplatformdatasetspb.ListDatasetsResponse)(nil)).Elem())
	method2RequestType.Store("google.maps.mapsplatformdatasets.v1.MapsPlatformDatasets/UpdateDatasetMetadata", reflect.TypeOf((*google_maps_mapsplatformdatasets_v1_mapsplatformdatasetspb.UpdateDatasetMetadataRequest)(nil)).Elem())
	method2ResponseType.Store("google.maps.mapsplatformdatasets.v1.MapsPlatformDatasets/UpdateDatasetMetadata", reflect.TypeOf((*google_maps_mapsplatformdatasets_v1_mapsplatformdatasetspb.Dataset)(nil)).Elem())
}
