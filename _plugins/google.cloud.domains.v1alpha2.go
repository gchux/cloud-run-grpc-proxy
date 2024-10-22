// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_domains_v1alpha2_domainspb "google.golang.org/genproto/googleapis/cloud/domains/v1alpha2"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.domains.v1alpha2.Domains", "domains.googleapis.com")
	method2RequestType.Store("google.cloud.domains.v1alpha2.Domains/ConfigureContactSettings", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.ConfigureContactSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.domains.v1alpha2.Domains/ConfigureContactSettings", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.domains.v1alpha2.Domains/ConfigureDnsSettings", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.ConfigureDnsSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.domains.v1alpha2.Domains/ConfigureDnsSettings", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.domains.v1alpha2.Domains/ConfigureManagementSettings", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.ConfigureManagementSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.domains.v1alpha2.Domains/ConfigureManagementSettings", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.domains.v1alpha2.Domains/DeleteRegistration", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.DeleteRegistrationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.domains.v1alpha2.Domains/DeleteRegistration", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.domains.v1alpha2.Domains/ExportRegistration", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.ExportRegistrationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.domains.v1alpha2.Domains/ExportRegistration", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.domains.v1alpha2.Domains/GetRegistration", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.GetRegistrationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.domains.v1alpha2.Domains/GetRegistration", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.Registration)(nil)).Elem())
	method2RequestType.Store("google.cloud.domains.v1alpha2.Domains/ListRegistrations", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.ListRegistrationsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.domains.v1alpha2.Domains/ListRegistrations", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.ListRegistrationsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.domains.v1alpha2.Domains/RegisterDomain", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.RegisterDomainRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.domains.v1alpha2.Domains/RegisterDomain", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.domains.v1alpha2.Domains/ResetAuthorizationCode", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.ResetAuthorizationCodeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.domains.v1alpha2.Domains/ResetAuthorizationCode", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.AuthorizationCode)(nil)).Elem())
	method2RequestType.Store("google.cloud.domains.v1alpha2.Domains/RetrieveAuthorizationCode", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.RetrieveAuthorizationCodeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.domains.v1alpha2.Domains/RetrieveAuthorizationCode", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.AuthorizationCode)(nil)).Elem())
	method2RequestType.Store("google.cloud.domains.v1alpha2.Domains/RetrieveRegisterParameters", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.RetrieveRegisterParametersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.domains.v1alpha2.Domains/RetrieveRegisterParameters", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.RetrieveRegisterParametersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.domains.v1alpha2.Domains/RetrieveTransferParameters", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.RetrieveTransferParametersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.domains.v1alpha2.Domains/RetrieveTransferParameters", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.RetrieveTransferParametersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.domains.v1alpha2.Domains/SearchDomains", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.SearchDomainsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.domains.v1alpha2.Domains/SearchDomains", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.SearchDomainsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.domains.v1alpha2.Domains/TransferDomain", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.TransferDomainRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.domains.v1alpha2.Domains/TransferDomain", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.domains.v1alpha2.Domains/UpdateRegistration", reflect.TypeOf((*google_cloud_domains_v1alpha2_domainspb.UpdateRegistrationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.domains.v1alpha2.Domains/UpdateRegistration", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
