// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_workflows_v1beta_workflowspb "cloud.google.com/go/workflows/apiv1beta/workflowspb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.workflows.v1beta.Workflows", "workflows.googleapis.com")
	method2RequestType.Store("google.cloud.workflows.v1beta.Workflows/CreateWorkflow", reflect.TypeOf((*google_cloud_workflows_v1beta_workflowspb.CreateWorkflowRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workflows.v1beta.Workflows/CreateWorkflow", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workflows.v1beta.Workflows/DeleteWorkflow", reflect.TypeOf((*google_cloud_workflows_v1beta_workflowspb.DeleteWorkflowRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workflows.v1beta.Workflows/DeleteWorkflow", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.workflows.v1beta.Workflows/GetWorkflow", reflect.TypeOf((*google_cloud_workflows_v1beta_workflowspb.GetWorkflowRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workflows.v1beta.Workflows/GetWorkflow", reflect.TypeOf((*google_cloud_workflows_v1beta_workflowspb.Workflow)(nil)).Elem())
	method2RequestType.Store("google.cloud.workflows.v1beta.Workflows/ListWorkflows", reflect.TypeOf((*google_cloud_workflows_v1beta_workflowspb.ListWorkflowsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workflows.v1beta.Workflows/ListWorkflows", reflect.TypeOf((*google_cloud_workflows_v1beta_workflowspb.ListWorkflowsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.workflows.v1beta.Workflows/UpdateWorkflow", reflect.TypeOf((*google_cloud_workflows_v1beta_workflowspb.UpdateWorkflowRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.workflows.v1beta.Workflows/UpdateWorkflow", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}