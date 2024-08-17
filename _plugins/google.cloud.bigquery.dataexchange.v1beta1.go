// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb "cloud.google.com/go/bigquery/dataexchange/apiv1beta1/dataexchangepb"
	google_iam_v1_iampb "cloud.google.com/go/iam/apiv1/iampb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService", "analyticshub.googleapis.com")
	method2RequestType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/CreateDataExchange", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.CreateDataExchangeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/CreateDataExchange", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.DataExchange)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/CreateListing", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.CreateListingRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/CreateListing", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.Listing)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/DeleteDataExchange", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.DeleteDataExchangeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/DeleteDataExchange", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/DeleteListing", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.DeleteListingRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/DeleteListing", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/GetDataExchange", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.GetDataExchangeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/GetDataExchange", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.DataExchange)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.GetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/GetListing", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.GetListingRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/GetListing", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.Listing)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/ListDataExchanges", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.ListDataExchangesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/ListDataExchanges", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.ListDataExchangesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/ListListings", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.ListListingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/ListListings", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.ListListingsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/ListOrgDataExchanges", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.ListOrgDataExchangesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/ListOrgDataExchanges", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.ListOrgDataExchangesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.SetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/SubscribeListing", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.SubscribeListingRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/SubscribeListing", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.SubscribeListingResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/UpdateDataExchange", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.UpdateDataExchangeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/UpdateDataExchange", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.DataExchange)(nil)).Elem())
	method2RequestType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/UpdateListing", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.UpdateListingRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.bigquery.dataexchange.v1beta1.AnalyticsHubService/UpdateListing", reflect.TypeOf((*google_cloud_bigquery_dataexchange_v1beta1_dataexchangepb.Listing)(nil)).Elem())
}