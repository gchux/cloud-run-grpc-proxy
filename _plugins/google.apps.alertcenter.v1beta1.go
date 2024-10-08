// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_apps_alertcenter_v1beta1_alertcenterpb "google.golang.org/genproto/googleapis/apps/alertcenter/v1beta1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.apps.alertcenter.v1beta1.AlertCenterService", "alertcenter.googleapis.com")
	method2RequestType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/BatchDeleteAlerts", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.BatchDeleteAlertsRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/BatchDeleteAlerts", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.BatchDeleteAlertsResponse)(nil)).Elem())
	method2RequestType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/BatchUndeleteAlerts", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.BatchUndeleteAlertsRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/BatchUndeleteAlerts", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.BatchUndeleteAlertsResponse)(nil)).Elem())
	method2RequestType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/CreateAlertFeedback", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.CreateAlertFeedbackRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/CreateAlertFeedback", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.AlertFeedback)(nil)).Elem())
	method2RequestType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/DeleteAlert", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.DeleteAlertRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/DeleteAlert", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/GetAlert", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.GetAlertRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/GetAlert", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.Alert)(nil)).Elem())
	method2RequestType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/GetAlertMetadata", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.GetAlertMetadataRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/GetAlertMetadata", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.AlertMetadata)(nil)).Elem())
	method2RequestType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/GetSettings", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.GetSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/GetSettings", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.Settings)(nil)).Elem())
	method2RequestType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/ListAlertFeedback", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.ListAlertFeedbackRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/ListAlertFeedback", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.ListAlertFeedbackResponse)(nil)).Elem())
	method2RequestType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/ListAlerts", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.ListAlertsRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/ListAlerts", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.ListAlertsResponse)(nil)).Elem())
	method2RequestType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/UndeleteAlert", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.UndeleteAlertRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/UndeleteAlert", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.Alert)(nil)).Elem())
	method2RequestType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/UpdateSettings", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.UpdateSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.alertcenter.v1beta1.AlertCenterService/UpdateSettings", reflect.TypeOf((*google_apps_alertcenter_v1beta1_alertcenterpb.Settings)(nil)).Elem())
}
