// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_recaptchaenterprise_v1beta1_recaptchaenterprisepb "cloud.google.com/go/recaptchaenterprise/v2/apiv1beta1/recaptchaenterprisepb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.recaptchaenterprise.v1beta1.RecaptchaEnterpriseServiceV1Beta1", "recaptchaenterprise.googleapis.com")
	method2RequestType.Store("google.cloud.recaptchaenterprise.v1beta1.RecaptchaEnterpriseServiceV1Beta1/AnnotateAssessment", reflect.TypeOf((*google_cloud_recaptchaenterprise_v1beta1_recaptchaenterprisepb.AnnotateAssessmentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.recaptchaenterprise.v1beta1.RecaptchaEnterpriseServiceV1Beta1/AnnotateAssessment", reflect.TypeOf((*google_cloud_recaptchaenterprise_v1beta1_recaptchaenterprisepb.AnnotateAssessmentResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.recaptchaenterprise.v1beta1.RecaptchaEnterpriseServiceV1Beta1/CreateAssessment", reflect.TypeOf((*google_cloud_recaptchaenterprise_v1beta1_recaptchaenterprisepb.CreateAssessmentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.recaptchaenterprise.v1beta1.RecaptchaEnterpriseServiceV1Beta1/CreateAssessment", reflect.TypeOf((*google_cloud_recaptchaenterprise_v1beta1_recaptchaenterprisepb.Assessment)(nil)).Elem())
}
