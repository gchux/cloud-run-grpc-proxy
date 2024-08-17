// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_datacatalog_v1beta1_datacatalogpb "cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb"
	google_iam_v1_iampb "cloud.google.com/go/iam/apiv1/iampb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.datacatalog.v1beta1.DataCatalog", "datacatalog.googleapis.com")
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/CreateEntry", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.CreateEntryRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/CreateEntry", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.Entry)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/CreateEntryGroup", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.CreateEntryGroupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/CreateEntryGroup", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.EntryGroup)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/CreateTag", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.CreateTagRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/CreateTag", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.Tag)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/CreateTagTemplate", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.CreateTagTemplateRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/CreateTagTemplate", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.TagTemplate)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/CreateTagTemplateField", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.CreateTagTemplateFieldRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/CreateTagTemplateField", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.TagTemplateField)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/DeleteEntry", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.DeleteEntryRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/DeleteEntry", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/DeleteEntryGroup", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.DeleteEntryGroupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/DeleteEntryGroup", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/DeleteTag", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.DeleteTagRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/DeleteTag", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/DeleteTagTemplate", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.DeleteTagTemplateRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/DeleteTagTemplate", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/DeleteTagTemplateField", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.DeleteTagTemplateFieldRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/DeleteTagTemplateField", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/GetEntry", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.GetEntryRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/GetEntry", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.Entry)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/GetEntryGroup", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.GetEntryGroupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/GetEntryGroup", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.EntryGroup)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.GetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/GetTagTemplate", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.GetTagTemplateRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/GetTagTemplate", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.TagTemplate)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/ListEntries", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.ListEntriesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/ListEntries", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.ListEntriesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/ListEntryGroups", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.ListEntryGroupsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/ListEntryGroups", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.ListEntryGroupsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/ListTags", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.ListTagsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/ListTags", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.ListTagsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/LookupEntry", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.LookupEntryRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/LookupEntry", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.Entry)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/RenameTagTemplateField", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.RenameTagTemplateFieldRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/RenameTagTemplateField", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.TagTemplateField)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/RenameTagTemplateFieldEnumValue", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.RenameTagTemplateFieldEnumValueRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/RenameTagTemplateFieldEnumValue", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.TagTemplateField)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/SearchCatalog", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.SearchCatalogRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/SearchCatalog", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.SearchCatalogResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.SetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/UpdateEntry", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.UpdateEntryRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/UpdateEntry", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.Entry)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/UpdateEntryGroup", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.UpdateEntryGroupRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/UpdateEntryGroup", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.EntryGroup)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/UpdateTag", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.UpdateTagRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/UpdateTag", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.Tag)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/UpdateTagTemplate", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.UpdateTagTemplateRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/UpdateTagTemplate", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.TagTemplate)(nil)).Elem())
	method2RequestType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/UpdateTagTemplateField", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.UpdateTagTemplateFieldRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.datacatalog.v1beta1.DataCatalog/UpdateTagTemplateField", reflect.TypeOf((*google_cloud_datacatalog_v1beta1_datacatalogpb.TagTemplateField)(nil)).Elem())
}