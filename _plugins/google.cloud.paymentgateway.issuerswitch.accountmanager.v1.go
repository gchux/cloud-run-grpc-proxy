// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_paymentgateway_issuerswitch_accountmanager_v1_accountmanagerpb "cloud.google.com/go/paymentgateway/issuerswitch/accountmanager/apiv1/accountmanagerpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.paymentgateway.issuerswitch.accountmanager.v1.ManagedAccounts", "issuerswitch.googleapis.com")
	method2RequestType.Store("google.cloud.paymentgateway.issuerswitch.accountmanager.v1.ManagedAccounts/GetManagedAccount", reflect.TypeOf((*google_cloud_paymentgateway_issuerswitch_accountmanager_v1_accountmanagerpb.GetManagedAccountRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.paymentgateway.issuerswitch.accountmanager.v1.ManagedAccounts/GetManagedAccount", reflect.TypeOf((*google_cloud_paymentgateway_issuerswitch_accountmanager_v1_accountmanagerpb.ManagedAccount)(nil)).Elem())
}
