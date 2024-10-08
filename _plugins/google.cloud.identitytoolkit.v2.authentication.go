// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_identitytoolkit_v2_authentication_servicepb "google.golang.org/genproto/googleapis/cloud/identitytoolkit/v2"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.identitytoolkit.v2.AuthenticationService", "identitytoolkit.googleapis.com")
	method2RequestType.Store("google.cloud.identitytoolkit.v2.AuthenticationService/FinalizeMfaSignIn", reflect.TypeOf((*google_cloud_identitytoolkit_v2_authentication_servicepb.FinalizeMfaSignInRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.identitytoolkit.v2.AuthenticationService/FinalizeMfaSignIn", reflect.TypeOf((*google_cloud_identitytoolkit_v2_authentication_servicepb.FinalizeMfaSignInResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.identitytoolkit.v2.AuthenticationService/StartMfaSignIn", reflect.TypeOf((*google_cloud_identitytoolkit_v2_authentication_servicepb.StartMfaSignInRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.identitytoolkit.v2.AuthenticationService/StartMfaSignIn", reflect.TypeOf((*google_cloud_identitytoolkit_v2_authentication_servicepb.StartMfaSignInResponse)(nil)).Elem())
}
