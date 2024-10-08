// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_workflows_executions_v1_executionspb "cloud.google.com/go/workflows/executions/apiv1/executionspb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.workflows.executions.v1.Executions", "workflowexecutions.googleapis.com")
	method2RequestType.Store("google.cloud.workflows.executions.v1.Executions/CancelExecution", reflect.TypeOf((*google_cloud_workflows_executions_v1_executionspb.CancelExecutionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workflows.executions.v1.Executions/CancelExecution", reflect.TypeOf((*google_cloud_workflows_executions_v1_executionspb.Execution)(nil)).Elem())
	method2RequestType.Store("google.cloud.workflows.executions.v1.Executions/CreateExecution", reflect.TypeOf((*google_cloud_workflows_executions_v1_executionspb.CreateExecutionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workflows.executions.v1.Executions/CreateExecution", reflect.TypeOf((*google_cloud_workflows_executions_v1_executionspb.Execution)(nil)).Elem())
	method2RequestType.Store("google.cloud.workflows.executions.v1.Executions/GetExecution", reflect.TypeOf((*google_cloud_workflows_executions_v1_executionspb.GetExecutionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workflows.executions.v1.Executions/GetExecution", reflect.TypeOf((*google_cloud_workflows_executions_v1_executionspb.Execution)(nil)).Elem())
	method2RequestType.Store("google.cloud.workflows.executions.v1.Executions/ListExecutions", reflect.TypeOf((*google_cloud_workflows_executions_v1_executionspb.ListExecutionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workflows.executions.v1.Executions/ListExecutions", reflect.TypeOf((*google_cloud_workflows_executions_v1_executionspb.ListExecutionsResponse)(nil)).Elem())
}
