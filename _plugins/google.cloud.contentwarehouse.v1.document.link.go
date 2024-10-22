// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_contentwarehouse_v1_document_link_servicepb "google.golang.org/genproto/googleapis/cloud/contentwarehouse/v1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.contentwarehouse.v1.DocumentLinkService", "contentwarehouse.googleapis.com")
	method2RequestType.Store("google.cloud.contentwarehouse.v1.DocumentLinkService/CreateDocumentLink", reflect.TypeOf((*google_cloud_contentwarehouse_v1_document_link_servicepb.CreateDocumentLinkRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contentwarehouse.v1.DocumentLinkService/CreateDocumentLink", reflect.TypeOf((*google_cloud_contentwarehouse_v1_document_link_servicepb.DocumentLink)(nil)).Elem())
	method2RequestType.Store("google.cloud.contentwarehouse.v1.DocumentLinkService/DeleteDocumentLink", reflect.TypeOf((*google_cloud_contentwarehouse_v1_document_link_servicepb.DeleteDocumentLinkRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contentwarehouse.v1.DocumentLinkService/DeleteDocumentLink", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.contentwarehouse.v1.DocumentLinkService/ListLinkedSources", reflect.TypeOf((*google_cloud_contentwarehouse_v1_document_link_servicepb.ListLinkedSourcesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contentwarehouse.v1.DocumentLinkService/ListLinkedSources", reflect.TypeOf((*google_cloud_contentwarehouse_v1_document_link_servicepb.ListLinkedSourcesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.contentwarehouse.v1.DocumentLinkService/ListLinkedTargets", reflect.TypeOf((*google_cloud_contentwarehouse_v1_document_link_servicepb.ListLinkedTargetsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contentwarehouse.v1.DocumentLinkService/ListLinkedTargets", reflect.TypeOf((*google_cloud_contentwarehouse_v1_document_link_servicepb.ListLinkedTargetsResponse)(nil)).Elem())
}
