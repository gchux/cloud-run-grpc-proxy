// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_apps_drive_labels_v2_label_servicepb "google.golang.org/genproto/googleapis/apps/drive/labels/v2"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.apps.drive.labels.v2.LabelService", "drivelabels.googleapis.com")
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/BatchDeleteLabelPermissions", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.BatchDeleteLabelPermissionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/BatchDeleteLabelPermissions", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/BatchUpdateLabelPermissions", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.BatchUpdateLabelPermissionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/BatchUpdateLabelPermissions", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.BatchUpdateLabelPermissionsResponse)(nil)).Elem())
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/CreateLabel", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.CreateLabelRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/CreateLabel", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.Label)(nil)).Elem())
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/CreateLabelPermission", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.CreateLabelPermissionRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/CreateLabelPermission", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.LabelPermission)(nil)).Elem())
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/DeleteLabel", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.DeleteLabelRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/DeleteLabel", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/DeleteLabelPermission", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.DeleteLabelPermissionRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/DeleteLabelPermission", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/DeltaUpdateLabel", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.DeltaUpdateLabelRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/DeltaUpdateLabel", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.DeltaUpdateLabelResponse)(nil)).Elem())
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/DisableLabel", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.DisableLabelRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/DisableLabel", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.Label)(nil)).Elem())
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/EnableLabel", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.EnableLabelRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/EnableLabel", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.Label)(nil)).Elem())
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/GetLabel", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.GetLabelRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/GetLabel", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.Label)(nil)).Elem())
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/GetLabelLimits", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.GetLabelLimitsRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/GetLabelLimits", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.LabelLimits)(nil)).Elem())
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/GetUserCapabilities", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.GetUserCapabilitiesRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/GetUserCapabilities", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.UserCapabilities)(nil)).Elem())
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/ListLabelLocks", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.ListLabelLocksRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/ListLabelLocks", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.ListLabelLocksResponse)(nil)).Elem())
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/ListLabelPermissions", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.ListLabelPermissionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/ListLabelPermissions", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.ListLabelPermissionsResponse)(nil)).Elem())
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/ListLabels", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.ListLabelsRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/ListLabels", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.ListLabelsResponse)(nil)).Elem())
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/PublishLabel", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.PublishLabelRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/PublishLabel", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.Label)(nil)).Elem())
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/UpdateLabelCopyMode", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.UpdateLabelCopyModeRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/UpdateLabelCopyMode", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.Label)(nil)).Elem())
	method2RequestType.Store("google.apps.drive.labels.v2.LabelService/UpdateLabelPermission", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.UpdateLabelPermissionRequest)(nil)).Elem())
	method2ResponseType.Store("google.apps.drive.labels.v2.LabelService/UpdateLabelPermission", reflect.TypeOf((*google_apps_drive_labels_v2_label_servicepb.LabelPermission)(nil)).Elem())
}
