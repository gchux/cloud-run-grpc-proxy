// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_discoveryengine_v1_discoveryenginepb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService", "discoveryengine.googleapis.com")
	method2RequestType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/BatchCreateTargetSites", reflect.TypeOf((*google_cloud_discoveryengine_v1_discoveryenginepb.BatchCreateTargetSitesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/BatchCreateTargetSites", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/BatchVerifyTargetSites", reflect.TypeOf((*google_cloud_discoveryengine_v1_discoveryenginepb.BatchVerifyTargetSitesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/BatchVerifyTargetSites", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/CreateTargetSite", reflect.TypeOf((*google_cloud_discoveryengine_v1_discoveryenginepb.CreateTargetSiteRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/CreateTargetSite", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/DeleteTargetSite", reflect.TypeOf((*google_cloud_discoveryengine_v1_discoveryenginepb.DeleteTargetSiteRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/DeleteTargetSite", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/DisableAdvancedSiteSearch", reflect.TypeOf((*google_cloud_discoveryengine_v1_discoveryenginepb.DisableAdvancedSiteSearchRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/DisableAdvancedSiteSearch", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/EnableAdvancedSiteSearch", reflect.TypeOf((*google_cloud_discoveryengine_v1_discoveryenginepb.EnableAdvancedSiteSearchRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/EnableAdvancedSiteSearch", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/FetchDomainVerificationStatus", reflect.TypeOf((*google_cloud_discoveryengine_v1_discoveryenginepb.FetchDomainVerificationStatusRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/FetchDomainVerificationStatus", reflect.TypeOf((*google_cloud_discoveryengine_v1_discoveryenginepb.FetchDomainVerificationStatusResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/GetSiteSearchEngine", reflect.TypeOf((*google_cloud_discoveryengine_v1_discoveryenginepb.GetSiteSearchEngineRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/GetSiteSearchEngine", reflect.TypeOf((*google_cloud_discoveryengine_v1_discoveryenginepb.SiteSearchEngine)(nil)).Elem())
	method2RequestType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/GetTargetSite", reflect.TypeOf((*google_cloud_discoveryengine_v1_discoveryenginepb.GetTargetSiteRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/GetTargetSite", reflect.TypeOf((*google_cloud_discoveryengine_v1_discoveryenginepb.TargetSite)(nil)).Elem())
	method2RequestType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/ListTargetSites", reflect.TypeOf((*google_cloud_discoveryengine_v1_discoveryenginepb.ListTargetSitesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/ListTargetSites", reflect.TypeOf((*google_cloud_discoveryengine_v1_discoveryenginepb.ListTargetSitesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/RecrawlUris", reflect.TypeOf((*google_cloud_discoveryengine_v1_discoveryenginepb.RecrawlUrisRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/RecrawlUris", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/UpdateTargetSite", reflect.TypeOf((*google_cloud_discoveryengine_v1_discoveryenginepb.UpdateTargetSiteRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.discoveryengine.v1.SiteSearchEngineService/UpdateTargetSite", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
