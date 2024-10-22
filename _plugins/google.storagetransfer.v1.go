// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_storagetransfer_v1_storagetransferpb "cloud.google.com/go/storagetransfer/apiv1/storagetransferpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.storagetransfer.v1.StorageTransferService", "storagetransfer.googleapis.com")
	method2RequestType.Store("google.storagetransfer.v1.StorageTransferService/CreateAgentPool", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.CreateAgentPoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.storagetransfer.v1.StorageTransferService/CreateAgentPool", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.AgentPool)(nil)).Elem())
	method2RequestType.Store("google.storagetransfer.v1.StorageTransferService/CreateTransferJob", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.CreateTransferJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.storagetransfer.v1.StorageTransferService/CreateTransferJob", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.TransferJob)(nil)).Elem())
	method2RequestType.Store("google.storagetransfer.v1.StorageTransferService/DeleteAgentPool", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.DeleteAgentPoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.storagetransfer.v1.StorageTransferService/DeleteAgentPool", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.storagetransfer.v1.StorageTransferService/DeleteTransferJob", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.DeleteTransferJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.storagetransfer.v1.StorageTransferService/DeleteTransferJob", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.storagetransfer.v1.StorageTransferService/GetAgentPool", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.GetAgentPoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.storagetransfer.v1.StorageTransferService/GetAgentPool", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.AgentPool)(nil)).Elem())
	method2RequestType.Store("google.storagetransfer.v1.StorageTransferService/GetGoogleServiceAccount", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.GetGoogleServiceAccountRequest)(nil)).Elem())
	method2ResponseType.Store("google.storagetransfer.v1.StorageTransferService/GetGoogleServiceAccount", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.GoogleServiceAccount)(nil)).Elem())
	method2RequestType.Store("google.storagetransfer.v1.StorageTransferService/GetTransferJob", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.GetTransferJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.storagetransfer.v1.StorageTransferService/GetTransferJob", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.TransferJob)(nil)).Elem())
	method2RequestType.Store("google.storagetransfer.v1.StorageTransferService/ListAgentPools", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.ListAgentPoolsRequest)(nil)).Elem())
	method2ResponseType.Store("google.storagetransfer.v1.StorageTransferService/ListAgentPools", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.ListAgentPoolsResponse)(nil)).Elem())
	method2RequestType.Store("google.storagetransfer.v1.StorageTransferService/ListTransferJobs", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.ListTransferJobsRequest)(nil)).Elem())
	method2ResponseType.Store("google.storagetransfer.v1.StorageTransferService/ListTransferJobs", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.ListTransferJobsResponse)(nil)).Elem())
	method2RequestType.Store("google.storagetransfer.v1.StorageTransferService/PauseTransferOperation", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.PauseTransferOperationRequest)(nil)).Elem())
	method2ResponseType.Store("google.storagetransfer.v1.StorageTransferService/PauseTransferOperation", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.storagetransfer.v1.StorageTransferService/ResumeTransferOperation", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.ResumeTransferOperationRequest)(nil)).Elem())
	method2ResponseType.Store("google.storagetransfer.v1.StorageTransferService/ResumeTransferOperation", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.storagetransfer.v1.StorageTransferService/RunTransferJob", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.RunTransferJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.storagetransfer.v1.StorageTransferService/RunTransferJob", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.storagetransfer.v1.StorageTransferService/UpdateAgentPool", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.UpdateAgentPoolRequest)(nil)).Elem())
	method2ResponseType.Store("google.storagetransfer.v1.StorageTransferService/UpdateAgentPool", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.AgentPool)(nil)).Elem())
	method2RequestType.Store("google.storagetransfer.v1.StorageTransferService/UpdateTransferJob", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.UpdateTransferJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.storagetransfer.v1.StorageTransferService/UpdateTransferJob", reflect.TypeOf((*google_storagetransfer_v1_storagetransferpb.TransferJob)(nil)).Elem())
}
