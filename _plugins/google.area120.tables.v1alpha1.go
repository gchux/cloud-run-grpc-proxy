// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_area120_tables_v1alpha1_tablespb "cloud.google.com/go/area120/tables/apiv1alpha1/tablespb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.area120.tables.v1alpha1.TablesService", "area120tables.googleapis.com")
	method2RequestType.Store("google.area120.tables.v1alpha1.TablesService/BatchCreateRows", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.BatchCreateRowsRequest)(nil)).Elem())
	method2ResponseType.Store("google.area120.tables.v1alpha1.TablesService/BatchCreateRows", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.BatchCreateRowsResponse)(nil)).Elem())
	method2RequestType.Store("google.area120.tables.v1alpha1.TablesService/BatchDeleteRows", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.BatchDeleteRowsRequest)(nil)).Elem())
	method2ResponseType.Store("google.area120.tables.v1alpha1.TablesService/BatchDeleteRows", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.area120.tables.v1alpha1.TablesService/BatchUpdateRows", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.BatchUpdateRowsRequest)(nil)).Elem())
	method2ResponseType.Store("google.area120.tables.v1alpha1.TablesService/BatchUpdateRows", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.BatchUpdateRowsResponse)(nil)).Elem())
	method2RequestType.Store("google.area120.tables.v1alpha1.TablesService/CreateRow", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.CreateRowRequest)(nil)).Elem())
	method2ResponseType.Store("google.area120.tables.v1alpha1.TablesService/CreateRow", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.Row)(nil)).Elem())
	method2RequestType.Store("google.area120.tables.v1alpha1.TablesService/DeleteRow", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.DeleteRowRequest)(nil)).Elem())
	method2ResponseType.Store("google.area120.tables.v1alpha1.TablesService/DeleteRow", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.area120.tables.v1alpha1.TablesService/GetRow", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.GetRowRequest)(nil)).Elem())
	method2ResponseType.Store("google.area120.tables.v1alpha1.TablesService/GetRow", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.Row)(nil)).Elem())
	method2RequestType.Store("google.area120.tables.v1alpha1.TablesService/GetTable", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.GetTableRequest)(nil)).Elem())
	method2ResponseType.Store("google.area120.tables.v1alpha1.TablesService/GetTable", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.Table)(nil)).Elem())
	method2RequestType.Store("google.area120.tables.v1alpha1.TablesService/GetWorkspace", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.GetWorkspaceRequest)(nil)).Elem())
	method2ResponseType.Store("google.area120.tables.v1alpha1.TablesService/GetWorkspace", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.Workspace)(nil)).Elem())
	method2RequestType.Store("google.area120.tables.v1alpha1.TablesService/ListRows", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.ListRowsRequest)(nil)).Elem())
	method2ResponseType.Store("google.area120.tables.v1alpha1.TablesService/ListRows", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.ListRowsResponse)(nil)).Elem())
	method2RequestType.Store("google.area120.tables.v1alpha1.TablesService/ListTables", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.ListTablesRequest)(nil)).Elem())
	method2ResponseType.Store("google.area120.tables.v1alpha1.TablesService/ListTables", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.ListTablesResponse)(nil)).Elem())
	method2RequestType.Store("google.area120.tables.v1alpha1.TablesService/ListWorkspaces", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.ListWorkspacesRequest)(nil)).Elem())
	method2ResponseType.Store("google.area120.tables.v1alpha1.TablesService/ListWorkspaces", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.ListWorkspacesResponse)(nil)).Elem())
	method2RequestType.Store("google.area120.tables.v1alpha1.TablesService/UpdateRow", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.UpdateRowRequest)(nil)).Elem())
	method2ResponseType.Store("google.area120.tables.v1alpha1.TablesService/UpdateRow", reflect.TypeOf((*google_area120_tables_v1alpha1_tablespb.Row)(nil)).Elem())
}
