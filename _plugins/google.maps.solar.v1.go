// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_maps_solar_v1_solarpb "cloud.google.com/go/maps/solar/apiv1/solarpb"
	google_api_httpbodypb "google.golang.org/genproto/googleapis/api/httpbody"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.maps.solar.v1.Solar", "solar.googleapis.com")
	method2RequestType.Store("google.maps.solar.v1.Solar/FindClosestBuildingInsights", reflect.TypeOf((*google_maps_solar_v1_solarpb.FindClosestBuildingInsightsRequest)(nil)).Elem())
	method2ResponseType.Store("google.maps.solar.v1.Solar/FindClosestBuildingInsights", reflect.TypeOf((*google_maps_solar_v1_solarpb.BuildingInsights)(nil)).Elem())
	method2RequestType.Store("google.maps.solar.v1.Solar/GetDataLayers", reflect.TypeOf((*google_maps_solar_v1_solarpb.GetDataLayersRequest)(nil)).Elem())
	method2ResponseType.Store("google.maps.solar.v1.Solar/GetDataLayers", reflect.TypeOf((*google_maps_solar_v1_solarpb.DataLayers)(nil)).Elem())
	method2RequestType.Store("google.maps.solar.v1.Solar/GetGeoTiff", reflect.TypeOf((*google_maps_solar_v1_solarpb.GetGeoTiffRequest)(nil)).Elem())
	method2ResponseType.Store("google.maps.solar.v1.Solar/GetGeoTiff", reflect.TypeOf((*google_api_httpbodypb.HttpBody)(nil)).Elem())
}
