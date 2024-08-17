// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_datastream_v1_datastreampb "cloud.google.com/go/datastream/apiv1/datastreampb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.datastream.v1.Datastream", "datastream.googleapis.com")
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/CreateConnectionProfile", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.CreateConnectionProfileRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/CreateConnectionProfile", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/CreatePrivateConnection", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.CreatePrivateConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/CreatePrivateConnection", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/CreateRoute", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.CreateRouteRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/CreateRoute", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/CreateStream", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.CreateStreamRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/CreateStream", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/DeleteConnectionProfile", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.DeleteConnectionProfileRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/DeleteConnectionProfile", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/DeletePrivateConnection", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.DeletePrivateConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/DeletePrivateConnection", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/DeleteRoute", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.DeleteRouteRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/DeleteRoute", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/DeleteStream", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.DeleteStreamRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/DeleteStream", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/DiscoverConnectionProfile", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.DiscoverConnectionProfileRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/DiscoverConnectionProfile", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.DiscoverConnectionProfileResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/FetchStaticIps", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.FetchStaticIpsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/FetchStaticIps", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.FetchStaticIpsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/GetConnectionProfile", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.GetConnectionProfileRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/GetConnectionProfile", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.ConnectionProfile)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/GetPrivateConnection", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.GetPrivateConnectionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/GetPrivateConnection", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.PrivateConnection)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/GetRoute", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.GetRouteRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/GetRoute", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.Route)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/GetStream", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.GetStreamRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/GetStream", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.Stream)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/GetStreamObject", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.GetStreamObjectRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/GetStreamObject", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.StreamObject)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/ListConnectionProfiles", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.ListConnectionProfilesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/ListConnectionProfiles", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.ListConnectionProfilesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/ListPrivateConnections", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.ListPrivateConnectionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/ListPrivateConnections", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.ListPrivateConnectionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/ListRoutes", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.ListRoutesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/ListRoutes", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.ListRoutesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/ListStreamObjects", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.ListStreamObjectsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/ListStreamObjects", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.ListStreamObjectsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/ListStreams", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.ListStreamsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/ListStreams", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.ListStreamsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/LookupStreamObject", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.LookupStreamObjectRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/LookupStreamObject", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.StreamObject)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/StartBackfillJob", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.StartBackfillJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/StartBackfillJob", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.StartBackfillJobResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/StopBackfillJob", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.StopBackfillJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/StopBackfillJob", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.StopBackfillJobResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/UpdateConnectionProfile", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.UpdateConnectionProfileRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/UpdateConnectionProfile", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.datastream.v1.Datastream/UpdateStream", reflect.TypeOf((*google_cloud_datastream_v1_datastreampb.UpdateStreamRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datastream.v1.Datastream/UpdateStream", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
