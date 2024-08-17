// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_workflows_executions_v1beta_executionspb "cloud.google.com/go/workflows/executions/apiv1beta/executionspb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.workflows.executions.v1beta.Executions", "workflowexecutions.googleapis.com")
	method2RequestType.Store("google.cloud.workflows.executions.v1beta.Executions/CancelExecution", reflect.TypeOf((*google_cloud_workflows_executions_v1beta_executionspb.CancelExecutionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workflows.executions.v1beta.Executions/CancelExecution", reflect.TypeOf((*google_cloud_workflows_executions_v1beta_executionspb.Execution)(nil)).Elem())
	method2RequestType.Store("google.cloud.workflows.executions.v1beta.Executions/CreateExecution", reflect.TypeOf((*google_cloud_workflows_executions_v1beta_executionspb.CreateExecutionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workflows.executions.v1beta.Executions/CreateExecution", reflect.TypeOf((*google_cloud_workflows_executions_v1beta_executionspb.Execution)(nil)).Elem())
	method2RequestType.Store("google.cloud.workflows.executions.v1beta.Executions/GetExecution", reflect.TypeOf((*google_cloud_workflows_executions_v1beta_executionspb.GetExecutionRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workflows.executions.v1beta.Executions/GetExecution", reflect.TypeOf((*google_cloud_workflows_executions_v1beta_executionspb.Execution)(nil)).Elem())
	method2RequestType.Store("google.cloud.workflows.executions.v1beta.Executions/ListExecutions", reflect.TypeOf((*google_cloud_workflows_executions_v1beta_executionspb.ListExecutionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workflows.executions.v1beta.Executions/ListExecutions", reflect.TypeOf((*google_cloud_workflows_executions_v1beta_executionspb.ListExecutionsResponse)(nil)).Elem())
}