// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_iam_v1_iampb "cloud.google.com/go/iam/apiv1/iampb"
	google_devtools_containeranalysis_v1_containeranalysispb "google.golang.org/genproto/googleapis/devtools/containeranalysis/v1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.devtools.containeranalysis.v1.ContainerAnalysis", "containeranalysis.googleapis.com")
	method2RequestType.Store("google.devtools.containeranalysis.v1.ContainerAnalysis/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.GetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.containeranalysis.v1.ContainerAnalysis/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.devtools.containeranalysis.v1.ContainerAnalysis/GetVulnerabilityOccurrencesSummary", reflect.TypeOf((*google_devtools_containeranalysis_v1_containeranalysispb.GetVulnerabilityOccurrencesSummaryRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.containeranalysis.v1.ContainerAnalysis/GetVulnerabilityOccurrencesSummary", reflect.TypeOf((*google_devtools_containeranalysis_v1_containeranalysispb.VulnerabilityOccurrencesSummary)(nil)).Elem())
	method2RequestType.Store("google.devtools.containeranalysis.v1.ContainerAnalysis/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.SetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.containeranalysis.v1.ContainerAnalysis/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.devtools.containeranalysis.v1.ContainerAnalysis/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.containeranalysis.v1.ContainerAnalysis/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsResponse)(nil)).Elem())
}
