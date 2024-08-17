// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_blockchainnodeengine_v1_blockchainnodeenginepb "cloud.google.com/go/blockchainnodeengine/apiv1/blockchainnodeenginepb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.blockchainnodeengine.v1.BlockchainNodeEngine", "blockchainnodeengine.googleapis.com")
	method2RequestType.Store("google.cloud.blockchainnodeengine.v1.BlockchainNodeEngine/CreateBlockchainNode", reflect.TypeOf((*google_cloud_blockchainnodeengine_v1_blockchainnodeenginepb.CreateBlockchainNodeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.blockchainnodeengine.v1.BlockchainNodeEngine/CreateBlockchainNode", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.blockchainnodeengine.v1.BlockchainNodeEngine/DeleteBlockchainNode", reflect.TypeOf((*google_cloud_blockchainnodeengine_v1_blockchainnodeenginepb.DeleteBlockchainNodeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.blockchainnodeengine.v1.BlockchainNodeEngine/DeleteBlockchainNode", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.blockchainnodeengine.v1.BlockchainNodeEngine/GetBlockchainNode", reflect.TypeOf((*google_cloud_blockchainnodeengine_v1_blockchainnodeenginepb.GetBlockchainNodeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.blockchainnodeengine.v1.BlockchainNodeEngine/GetBlockchainNode", reflect.TypeOf((*google_cloud_blockchainnodeengine_v1_blockchainnodeenginepb.BlockchainNode)(nil)).Elem())
	method2RequestType.Store("google.cloud.blockchainnodeengine.v1.BlockchainNodeEngine/ListBlockchainNodes", reflect.TypeOf((*google_cloud_blockchainnodeengine_v1_blockchainnodeenginepb.ListBlockchainNodesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.blockchainnodeengine.v1.BlockchainNodeEngine/ListBlockchainNodes", reflect.TypeOf((*google_cloud_blockchainnodeengine_v1_blockchainnodeenginepb.ListBlockchainNodesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.blockchainnodeengine.v1.BlockchainNodeEngine/UpdateBlockchainNode", reflect.TypeOf((*google_cloud_blockchainnodeengine_v1_blockchainnodeenginepb.UpdateBlockchainNodeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.blockchainnodeengine.v1.BlockchainNodeEngine/UpdateBlockchainNode", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}