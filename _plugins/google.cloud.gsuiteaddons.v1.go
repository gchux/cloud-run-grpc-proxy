// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_gsuiteaddons_v1_gsuiteaddonspb "cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns", "gsuiteaddons.googleapis.com")
	method2RequestType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/CreateDeployment", reflect.TypeOf((*google_cloud_gsuiteaddons_v1_gsuiteaddonspb.CreateDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/CreateDeployment", reflect.TypeOf((*google_cloud_gsuiteaddons_v1_gsuiteaddonspb.Deployment)(nil)).Elem())
	method2RequestType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/DeleteDeployment", reflect.TypeOf((*google_cloud_gsuiteaddons_v1_gsuiteaddonspb.DeleteDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/DeleteDeployment", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/GetAuthorization", reflect.TypeOf((*google_cloud_gsuiteaddons_v1_gsuiteaddonspb.GetAuthorizationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/GetAuthorization", reflect.TypeOf((*google_cloud_gsuiteaddons_v1_gsuiteaddonspb.Authorization)(nil)).Elem())
	method2RequestType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/GetDeployment", reflect.TypeOf((*google_cloud_gsuiteaddons_v1_gsuiteaddonspb.GetDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/GetDeployment", reflect.TypeOf((*google_cloud_gsuiteaddons_v1_gsuiteaddonspb.Deployment)(nil)).Elem())
	method2RequestType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/GetInstallStatus", reflect.TypeOf((*google_cloud_gsuiteaddons_v1_gsuiteaddonspb.GetInstallStatusRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/GetInstallStatus", reflect.TypeOf((*google_cloud_gsuiteaddons_v1_gsuiteaddonspb.InstallStatus)(nil)).Elem())
	method2RequestType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/InstallDeployment", reflect.TypeOf((*google_cloud_gsuiteaddons_v1_gsuiteaddonspb.InstallDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/InstallDeployment", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/ListDeployments", reflect.TypeOf((*google_cloud_gsuiteaddons_v1_gsuiteaddonspb.ListDeploymentsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/ListDeployments", reflect.TypeOf((*google_cloud_gsuiteaddons_v1_gsuiteaddonspb.ListDeploymentsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/ReplaceDeployment", reflect.TypeOf((*google_cloud_gsuiteaddons_v1_gsuiteaddonspb.ReplaceDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/ReplaceDeployment", reflect.TypeOf((*google_cloud_gsuiteaddons_v1_gsuiteaddonspb.Deployment)(nil)).Elem())
	method2RequestType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/UninstallDeployment", reflect.TypeOf((*google_cloud_gsuiteaddons_v1_gsuiteaddonspb.UninstallDeploymentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.gsuiteaddons.v1.GSuiteAddOns/UninstallDeployment", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
}