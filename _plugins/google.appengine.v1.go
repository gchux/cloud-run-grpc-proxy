// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_appengine_v1_appenginepb "cloud.google.com/go/appengine/apiv1/appenginepb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.appengine.v1.Applications", "appengine.googleapis.com")
	service2Host.Store("google.appengine.v1.AuthorizedCertificates", "appengine.googleapis.com")
	service2Host.Store("google.appengine.v1.AuthorizedDomains", "appengine.googleapis.com")
	service2Host.Store("google.appengine.v1.DomainMappings", "appengine.googleapis.com")
	service2Host.Store("google.appengine.v1.Firewall", "appengine.googleapis.com")
	service2Host.Store("google.appengine.v1.Instances", "appengine.googleapis.com")
	service2Host.Store("google.appengine.v1.Services", "appengine.googleapis.com")
	service2Host.Store("google.appengine.v1.Versions", "appengine.googleapis.com")
	method2RequestType.Store("google.appengine.v1.Applications/CreateApplication", reflect.TypeOf((*google_appengine_v1_appenginepb.CreateApplicationRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Applications/CreateApplication", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Applications/GetApplication", reflect.TypeOf((*google_appengine_v1_appenginepb.GetApplicationRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Applications/GetApplication", reflect.TypeOf((*google_appengine_v1_appenginepb.Application)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Applications/RepairApplication", reflect.TypeOf((*google_appengine_v1_appenginepb.RepairApplicationRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Applications/RepairApplication", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Applications/UpdateApplication", reflect.TypeOf((*google_appengine_v1_appenginepb.UpdateApplicationRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Applications/UpdateApplication", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.AuthorizedCertificates/CreateAuthorizedCertificate", reflect.TypeOf((*google_appengine_v1_appenginepb.CreateAuthorizedCertificateRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.AuthorizedCertificates/CreateAuthorizedCertificate", reflect.TypeOf((*google_appengine_v1_appenginepb.AuthorizedCertificate)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.AuthorizedCertificates/DeleteAuthorizedCertificate", reflect.TypeOf((*google_appengine_v1_appenginepb.DeleteAuthorizedCertificateRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.AuthorizedCertificates/DeleteAuthorizedCertificate", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.AuthorizedCertificates/GetAuthorizedCertificate", reflect.TypeOf((*google_appengine_v1_appenginepb.GetAuthorizedCertificateRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.AuthorizedCertificates/GetAuthorizedCertificate", reflect.TypeOf((*google_appengine_v1_appenginepb.AuthorizedCertificate)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.AuthorizedCertificates/ListAuthorizedCertificates", reflect.TypeOf((*google_appengine_v1_appenginepb.ListAuthorizedCertificatesRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.AuthorizedCertificates/ListAuthorizedCertificates", reflect.TypeOf((*google_appengine_v1_appenginepb.ListAuthorizedCertificatesResponse)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.AuthorizedCertificates/UpdateAuthorizedCertificate", reflect.TypeOf((*google_appengine_v1_appenginepb.UpdateAuthorizedCertificateRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.AuthorizedCertificates/UpdateAuthorizedCertificate", reflect.TypeOf((*google_appengine_v1_appenginepb.AuthorizedCertificate)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.AuthorizedDomains/ListAuthorizedDomains", reflect.TypeOf((*google_appengine_v1_appenginepb.ListAuthorizedDomainsRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.AuthorizedDomains/ListAuthorizedDomains", reflect.TypeOf((*google_appengine_v1_appenginepb.ListAuthorizedDomainsResponse)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.DomainMappings/CreateDomainMapping", reflect.TypeOf((*google_appengine_v1_appenginepb.CreateDomainMappingRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.DomainMappings/CreateDomainMapping", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.DomainMappings/DeleteDomainMapping", reflect.TypeOf((*google_appengine_v1_appenginepb.DeleteDomainMappingRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.DomainMappings/DeleteDomainMapping", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.DomainMappings/GetDomainMapping", reflect.TypeOf((*google_appengine_v1_appenginepb.GetDomainMappingRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.DomainMappings/GetDomainMapping", reflect.TypeOf((*google_appengine_v1_appenginepb.DomainMapping)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.DomainMappings/ListDomainMappings", reflect.TypeOf((*google_appengine_v1_appenginepb.ListDomainMappingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.DomainMappings/ListDomainMappings", reflect.TypeOf((*google_appengine_v1_appenginepb.ListDomainMappingsResponse)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.DomainMappings/UpdateDomainMapping", reflect.TypeOf((*google_appengine_v1_appenginepb.UpdateDomainMappingRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.DomainMappings/UpdateDomainMapping", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Firewall/BatchUpdateIngressRules", reflect.TypeOf((*google_appengine_v1_appenginepb.BatchUpdateIngressRulesRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Firewall/BatchUpdateIngressRules", reflect.TypeOf((*google_appengine_v1_appenginepb.BatchUpdateIngressRulesResponse)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Firewall/CreateIngressRule", reflect.TypeOf((*google_appengine_v1_appenginepb.CreateIngressRuleRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Firewall/CreateIngressRule", reflect.TypeOf((*google_appengine_v1_appenginepb.FirewallRule)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Firewall/DeleteIngressRule", reflect.TypeOf((*google_appengine_v1_appenginepb.DeleteIngressRuleRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Firewall/DeleteIngressRule", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Firewall/GetIngressRule", reflect.TypeOf((*google_appengine_v1_appenginepb.GetIngressRuleRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Firewall/GetIngressRule", reflect.TypeOf((*google_appengine_v1_appenginepb.FirewallRule)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Firewall/ListIngressRules", reflect.TypeOf((*google_appengine_v1_appenginepb.ListIngressRulesRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Firewall/ListIngressRules", reflect.TypeOf((*google_appengine_v1_appenginepb.ListIngressRulesResponse)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Firewall/UpdateIngressRule", reflect.TypeOf((*google_appengine_v1_appenginepb.UpdateIngressRuleRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Firewall/UpdateIngressRule", reflect.TypeOf((*google_appengine_v1_appenginepb.FirewallRule)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Instances/DebugInstance", reflect.TypeOf((*google_appengine_v1_appenginepb.DebugInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Instances/DebugInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Instances/DeleteInstance", reflect.TypeOf((*google_appengine_v1_appenginepb.DeleteInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Instances/DeleteInstance", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Instances/GetInstance", reflect.TypeOf((*google_appengine_v1_appenginepb.GetInstanceRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Instances/GetInstance", reflect.TypeOf((*google_appengine_v1_appenginepb.Instance)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Instances/ListInstances", reflect.TypeOf((*google_appengine_v1_appenginepb.ListInstancesRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Instances/ListInstances", reflect.TypeOf((*google_appengine_v1_appenginepb.ListInstancesResponse)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Services/DeleteService", reflect.TypeOf((*google_appengine_v1_appenginepb.DeleteServiceRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Services/DeleteService", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Services/GetService", reflect.TypeOf((*google_appengine_v1_appenginepb.GetServiceRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Services/GetService", reflect.TypeOf((*google_appengine_v1_appenginepb.Service)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Services/ListServices", reflect.TypeOf((*google_appengine_v1_appenginepb.ListServicesRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Services/ListServices", reflect.TypeOf((*google_appengine_v1_appenginepb.ListServicesResponse)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Services/UpdateService", reflect.TypeOf((*google_appengine_v1_appenginepb.UpdateServiceRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Services/UpdateService", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Versions/CreateVersion", reflect.TypeOf((*google_appengine_v1_appenginepb.CreateVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Versions/CreateVersion", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Versions/DeleteVersion", reflect.TypeOf((*google_appengine_v1_appenginepb.DeleteVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Versions/DeleteVersion", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Versions/GetVersion", reflect.TypeOf((*google_appengine_v1_appenginepb.GetVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Versions/GetVersion", reflect.TypeOf((*google_appengine_v1_appenginepb.Version)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Versions/ListVersions", reflect.TypeOf((*google_appengine_v1_appenginepb.ListVersionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Versions/ListVersions", reflect.TypeOf((*google_appengine_v1_appenginepb.ListVersionsResponse)(nil)).Elem())
	method2RequestType.Store("google.appengine.v1.Versions/UpdateVersion", reflect.TypeOf((*google_appengine_v1_appenginepb.UpdateVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.appengine.v1.Versions/UpdateVersion", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
