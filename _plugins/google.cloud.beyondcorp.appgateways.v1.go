// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_beyondcorp_appgateways_v1_appgatewayspb "cloud.google.com/go/beyondcorp/appgateways/apiv1/appgatewayspb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.beyondcorp.appgateways.v1.AppGatewaysService", "beyondcorp.googleapis.com")
	method2RequestType.Store("google.cloud.beyondcorp.appgateways.v1.AppGatewaysService/CreateAppGateway", reflect.TypeOf((*google_cloud_beyondcorp_appgateways_v1_appgatewayspb.CreateAppGatewayRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.beyondcorp.appgateways.v1.AppGatewaysService/CreateAppGateway", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.beyondcorp.appgateways.v1.AppGatewaysService/DeleteAppGateway", reflect.TypeOf((*google_cloud_beyondcorp_appgateways_v1_appgatewayspb.DeleteAppGatewayRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.beyondcorp.appgateways.v1.AppGatewaysService/DeleteAppGateway", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.beyondcorp.appgateways.v1.AppGatewaysService/GetAppGateway", reflect.TypeOf((*google_cloud_beyondcorp_appgateways_v1_appgatewayspb.GetAppGatewayRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.beyondcorp.appgateways.v1.AppGatewaysService/GetAppGateway", reflect.TypeOf((*google_cloud_beyondcorp_appgateways_v1_appgatewayspb.AppGateway)(nil)).Elem())
	method2RequestType.Store("google.cloud.beyondcorp.appgateways.v1.AppGatewaysService/ListAppGateways", reflect.TypeOf((*google_cloud_beyondcorp_appgateways_v1_appgatewayspb.ListAppGatewaysRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.beyondcorp.appgateways.v1.AppGatewaysService/ListAppGateways", reflect.TypeOf((*google_cloud_beyondcorp_appgateways_v1_appgatewayspb.ListAppGatewaysResponse)(nil)).Elem())
}
