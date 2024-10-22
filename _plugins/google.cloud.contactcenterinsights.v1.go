// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_contactcenterinsights_v1_contactcenterinsightspb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights", "contactcenterinsights.googleapis.com")
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/BulkAnalyzeConversations", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.BulkAnalyzeConversationsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/BulkAnalyzeConversations", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/BulkDeleteConversations", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.BulkDeleteConversationsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/BulkDeleteConversations", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/CalculateIssueModelStats", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.CalculateIssueModelStatsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/CalculateIssueModelStats", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.CalculateIssueModelStatsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/CalculateStats", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.CalculateStatsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/CalculateStats", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.CalculateStatsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/CreateAnalysis", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.CreateAnalysisRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/CreateAnalysis", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/CreateConversation", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.CreateConversationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/CreateConversation", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.Conversation)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/CreateIssueModel", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.CreateIssueModelRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/CreateIssueModel", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/CreatePhraseMatcher", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.CreatePhraseMatcherRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/CreatePhraseMatcher", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.PhraseMatcher)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/CreateView", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.CreateViewRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/CreateView", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.View)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/DeleteAnalysis", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.DeleteAnalysisRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/DeleteAnalysis", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/DeleteConversation", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.DeleteConversationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/DeleteConversation", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/DeleteIssue", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.DeleteIssueRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/DeleteIssue", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/DeleteIssueModel", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.DeleteIssueModelRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/DeleteIssueModel", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/DeletePhraseMatcher", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.DeletePhraseMatcherRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/DeletePhraseMatcher", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/DeleteView", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.DeleteViewRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/DeleteView", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/DeployIssueModel", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.DeployIssueModelRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/DeployIssueModel", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/ExportInsightsData", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.ExportInsightsDataRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/ExportInsightsData", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/GetAnalysis", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.GetAnalysisRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/GetAnalysis", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.Analysis)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/GetConversation", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.GetConversationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/GetConversation", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.Conversation)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/GetIssue", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.GetIssueRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/GetIssue", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.Issue)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/GetIssueModel", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.GetIssueModelRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/GetIssueModel", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.IssueModel)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/GetPhraseMatcher", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.GetPhraseMatcherRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/GetPhraseMatcher", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.PhraseMatcher)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/GetSettings", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.GetSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/GetSettings", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.Settings)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/GetView", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.GetViewRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/GetView", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.View)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/IngestConversations", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.IngestConversationsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/IngestConversations", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/ListAnalyses", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.ListAnalysesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/ListAnalyses", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.ListAnalysesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/ListConversations", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.ListConversationsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/ListConversations", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.ListConversationsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/ListIssueModels", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.ListIssueModelsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/ListIssueModels", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.ListIssueModelsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/ListIssues", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.ListIssuesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/ListIssues", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.ListIssuesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/ListPhraseMatchers", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.ListPhraseMatchersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/ListPhraseMatchers", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.ListPhraseMatchersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/ListViews", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.ListViewsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/ListViews", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.ListViewsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/UndeployIssueModel", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.UndeployIssueModelRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/UndeployIssueModel", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/UpdateConversation", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.UpdateConversationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/UpdateConversation", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.Conversation)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/UpdateIssue", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.UpdateIssueRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/UpdateIssue", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.Issue)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/UpdateIssueModel", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.UpdateIssueModelRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/UpdateIssueModel", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.IssueModel)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/UpdatePhraseMatcher", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.UpdatePhraseMatcherRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/UpdatePhraseMatcher", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.PhraseMatcher)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/UpdateSettings", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.UpdateSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/UpdateSettings", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.Settings)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/UpdateView", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.UpdateViewRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/UpdateView", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.View)(nil)).Elem())
	method2RequestType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/UploadConversation", reflect.TypeOf((*google_cloud_contactcenterinsights_v1_contactcenterinsightspb.UploadConversationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.contactcenterinsights.v1.ContactCenterInsights/UploadConversation", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
