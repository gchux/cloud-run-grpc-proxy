// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_partner_aistreams_v1alpha1_aistreamspb "google.golang.org/genproto/googleapis/partner/aistreams/v1alpha1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.partner.aistreams.v1alpha1.AIStreams", "aistreams.googleapis.com")
	method2RequestType.Store("google.partner.aistreams.v1alpha1.AIStreams/CreateCluster", reflect.TypeOf((*google_partner_aistreams_v1alpha1_aistreamspb.CreateClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.partner.aistreams.v1alpha1.AIStreams/CreateCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.partner.aistreams.v1alpha1.AIStreams/CreateStream", reflect.TypeOf((*google_partner_aistreams_v1alpha1_aistreamspb.CreateStreamRequest)(nil)).Elem())
	method2ResponseType.Store("google.partner.aistreams.v1alpha1.AIStreams/CreateStream", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.partner.aistreams.v1alpha1.AIStreams/DeleteCluster", reflect.TypeOf((*google_partner_aistreams_v1alpha1_aistreamspb.DeleteClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.partner.aistreams.v1alpha1.AIStreams/DeleteCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.partner.aistreams.v1alpha1.AIStreams/DeleteStream", reflect.TypeOf((*google_partner_aistreams_v1alpha1_aistreamspb.DeleteStreamRequest)(nil)).Elem())
	method2ResponseType.Store("google.partner.aistreams.v1alpha1.AIStreams/DeleteStream", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.partner.aistreams.v1alpha1.AIStreams/GetCluster", reflect.TypeOf((*google_partner_aistreams_v1alpha1_aistreamspb.GetClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.partner.aistreams.v1alpha1.AIStreams/GetCluster", reflect.TypeOf((*google_partner_aistreams_v1alpha1_aistreamspb.Cluster)(nil)).Elem())
	method2RequestType.Store("google.partner.aistreams.v1alpha1.AIStreams/GetStream", reflect.TypeOf((*google_partner_aistreams_v1alpha1_aistreamspb.GetStreamRequest)(nil)).Elem())
	method2ResponseType.Store("google.partner.aistreams.v1alpha1.AIStreams/GetStream", reflect.TypeOf((*google_partner_aistreams_v1alpha1_aistreamspb.Stream)(nil)).Elem())
	method2RequestType.Store("google.partner.aistreams.v1alpha1.AIStreams/ListClusters", reflect.TypeOf((*google_partner_aistreams_v1alpha1_aistreamspb.ListClustersRequest)(nil)).Elem())
	method2ResponseType.Store("google.partner.aistreams.v1alpha1.AIStreams/ListClusters", reflect.TypeOf((*google_partner_aistreams_v1alpha1_aistreamspb.ListClustersResponse)(nil)).Elem())
	method2RequestType.Store("google.partner.aistreams.v1alpha1.AIStreams/ListStreams", reflect.TypeOf((*google_partner_aistreams_v1alpha1_aistreamspb.ListStreamsRequest)(nil)).Elem())
	method2ResponseType.Store("google.partner.aistreams.v1alpha1.AIStreams/ListStreams", reflect.TypeOf((*google_partner_aistreams_v1alpha1_aistreamspb.ListStreamsResponse)(nil)).Elem())
	method2RequestType.Store("google.partner.aistreams.v1alpha1.AIStreams/UpdateCluster", reflect.TypeOf((*google_partner_aistreams_v1alpha1_aistreamspb.UpdateClusterRequest)(nil)).Elem())
	method2ResponseType.Store("google.partner.aistreams.v1alpha1.AIStreams/UpdateCluster", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.partner.aistreams.v1alpha1.AIStreams/UpdateStream", reflect.TypeOf((*google_partner_aistreams_v1alpha1_aistreamspb.UpdateStreamRequest)(nil)).Elem())
	method2ResponseType.Store("google.partner.aistreams.v1alpha1.AIStreams/UpdateStream", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
