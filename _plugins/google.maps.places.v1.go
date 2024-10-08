// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_maps_places_v1_placespb "cloud.google.com/go/maps/places/apiv1/placespb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.maps.places.v1.Places", "places.googleapis.com")
	method2RequestType.Store("google.maps.places.v1.Places/AutocompletePlaces", reflect.TypeOf((*google_maps_places_v1_placespb.AutocompletePlacesRequest)(nil)).Elem())
	method2ResponseType.Store("google.maps.places.v1.Places/AutocompletePlaces", reflect.TypeOf((*google_maps_places_v1_placespb.AutocompletePlacesResponse)(nil)).Elem())
	method2RequestType.Store("google.maps.places.v1.Places/GetPhotoMedia", reflect.TypeOf((*google_maps_places_v1_placespb.GetPhotoMediaRequest)(nil)).Elem())
	method2ResponseType.Store("google.maps.places.v1.Places/GetPhotoMedia", reflect.TypeOf((*google_maps_places_v1_placespb.PhotoMedia)(nil)).Elem())
	method2RequestType.Store("google.maps.places.v1.Places/GetPlace", reflect.TypeOf((*google_maps_places_v1_placespb.GetPlaceRequest)(nil)).Elem())
	method2ResponseType.Store("google.maps.places.v1.Places/GetPlace", reflect.TypeOf((*google_maps_places_v1_placespb.Place)(nil)).Elem())
	method2RequestType.Store("google.maps.places.v1.Places/SearchNearby", reflect.TypeOf((*google_maps_places_v1_placespb.SearchNearbyRequest)(nil)).Elem())
	method2ResponseType.Store("google.maps.places.v1.Places/SearchNearby", reflect.TypeOf((*google_maps_places_v1_placespb.SearchNearbyResponse)(nil)).Elem())
	method2RequestType.Store("google.maps.places.v1.Places/SearchText", reflect.TypeOf((*google_maps_places_v1_placespb.SearchTextRequest)(nil)).Elem())
	method2ResponseType.Store("google.maps.places.v1.Places/SearchText", reflect.TypeOf((*google_maps_places_v1_placespb.SearchTextResponse)(nil)).Elem())
}
