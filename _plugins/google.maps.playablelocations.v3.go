// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_maps_playablelocations_v3_playablelocationspb "google.golang.org/genproto/googleapis/maps/playablelocations/v3"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.maps.playablelocations.v3.PlayableLocations", "playablelocations.googleapis.com")
	method2RequestType.Store("google.maps.playablelocations.v3.PlayableLocations/LogImpressions", reflect.TypeOf((*google_maps_playablelocations_v3_playablelocationspb.LogImpressionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.maps.playablelocations.v3.PlayableLocations/LogImpressions", reflect.TypeOf((*google_maps_playablelocations_v3_playablelocationspb.LogImpressionsResponse)(nil)).Elem())
	method2RequestType.Store("google.maps.playablelocations.v3.PlayableLocations/LogPlayerReports", reflect.TypeOf((*google_maps_playablelocations_v3_playablelocationspb.LogPlayerReportsRequest)(nil)).Elem())
	method2ResponseType.Store("google.maps.playablelocations.v3.PlayableLocations/LogPlayerReports", reflect.TypeOf((*google_maps_playablelocations_v3_playablelocationspb.LogPlayerReportsResponse)(nil)).Elem())
	method2RequestType.Store("google.maps.playablelocations.v3.PlayableLocations/SamplePlayableLocations", reflect.TypeOf((*google_maps_playablelocations_v3_playablelocationspb.SamplePlayableLocationsRequest)(nil)).Elem())
	method2ResponseType.Store("google.maps.playablelocations.v3.PlayableLocations/SamplePlayableLocations", reflect.TypeOf((*google_maps_playablelocations_v3_playablelocationspb.SamplePlayableLocationsResponse)(nil)).Elem())
}
