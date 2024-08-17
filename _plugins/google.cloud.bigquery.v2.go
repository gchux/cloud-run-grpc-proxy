// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_bigquery_v2_bigquerypb "cloud.google.com/go/bigquery/apiv2/bigquerypb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.bigquery.v2.RoutineService", "bigquery.googleapis.com")
	method2RequestType.Store("google.cloud.bigquery.v2.RoutineService/DeleteRoutine", reflect.TypeOf((*google_cloud_bigquery_v2_bigquerypb.DeleteRoutineRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.v2.RoutineService/DeleteRoutine", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.v2.RoutineService/GetRoutine", reflect.TypeOf((*google_cloud_bigquery_v2_bigquerypb.GetRoutineRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.v2.RoutineService/GetRoutine", reflect.TypeOf((*google_cloud_bigquery_v2_bigquerypb.Routine)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.v2.RoutineService/InsertRoutine", reflect.TypeOf((*google_cloud_bigquery_v2_bigquerypb.InsertRoutineRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.v2.RoutineService/InsertRoutine", reflect.TypeOf((*google_cloud_bigquery_v2_bigquerypb.Routine)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.v2.RoutineService/ListRoutines", reflect.TypeOf((*google_cloud_bigquery_v2_bigquerypb.ListRoutinesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.v2.RoutineService/ListRoutines", reflect.TypeOf((*google_cloud_bigquery_v2_bigquerypb.ListRoutinesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.v2.RoutineService/PatchRoutine", reflect.TypeOf((*google_cloud_bigquery_v2_bigquerypb.PatchRoutineRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.v2.RoutineService/PatchRoutine", reflect.TypeOf((*google_cloud_bigquery_v2_bigquerypb.Routine)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.v2.RoutineService/UpdateRoutine", reflect.TypeOf((*google_cloud_bigquery_v2_bigquerypb.UpdateRoutineRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.v2.RoutineService/UpdateRoutine", reflect.TypeOf((*google_cloud_bigquery_v2_bigquerypb.Routine)(nil)).Elem())
}
