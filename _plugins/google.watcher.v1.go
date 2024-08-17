// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_watcher_v1_watchpb "google.golang.org/genproto/googleapis/watcher/v1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	method2RequestType.Store("google.watcher.v1.Watcher/Watch", reflect.TypeOf((*google_watcher_v1_watchpb.Request)(nil)).Elem())
	method2ResponseType.Store("google.watcher.v1.Watcher/Watch", reflect.TypeOf((*google_watcher_v1_watchpb.ChangeBatch)(nil)).Elem())
}
