// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_confidentialcomputing_v1_confidentialcomputingpb "cloud.google.com/go/confidentialcomputing/apiv1/confidentialcomputingpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.confidentialcomputing.v1.ConfidentialComputing", "confidentialcomputing.googleapis.com")
	method2RequestType.Store("google.cloud.confidentialcomputing.v1.ConfidentialComputing/CreateChallenge", reflect.TypeOf((*google_cloud_confidentialcomputing_v1_confidentialcomputingpb.CreateChallengeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.confidentialcomputing.v1.ConfidentialComputing/CreateChallenge", reflect.TypeOf((*google_cloud_confidentialcomputing_v1_confidentialcomputingpb.Challenge)(nil)).Elem())
	method2RequestType.Store("google.cloud.confidentialcomputing.v1.ConfidentialComputing/VerifyAttestation", reflect.TypeOf((*google_cloud_confidentialcomputing_v1_confidentialcomputingpb.VerifyAttestationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.confidentialcomputing.v1.ConfidentialComputing/VerifyAttestation", reflect.TypeOf((*google_cloud_confidentialcomputing_v1_confidentialcomputingpb.VerifyAttestationResponse)(nil)).Elem())
}
