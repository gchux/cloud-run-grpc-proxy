// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_securitycenter_settings_v1beta1_settingspb "cloud.google.com/go/securitycenter/settings/apiv1beta1/settingspb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService", "securitycenter.googleapis.com")
	method2RequestType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/BatchCalculateEffectiveSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.BatchCalculateEffectiveSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/BatchCalculateEffectiveSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.BatchCalculateEffectiveSettingsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/BatchGetSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.BatchGetSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/BatchGetSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.BatchGetSettingsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/CalculateEffectiveComponentSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.CalculateEffectiveComponentSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/CalculateEffectiveComponentSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.ComponentSettings)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/CalculateEffectiveSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.CalculateEffectiveSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/CalculateEffectiveSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.Settings)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/GetComponentSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.GetComponentSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/GetComponentSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.ComponentSettings)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/GetServiceAccount", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.GetServiceAccountRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/GetServiceAccount", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.ServiceAccount)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/GetSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.GetSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/GetSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.Settings)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/ListComponents", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.ListComponentsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/ListComponents", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.ListComponentsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/ListDetectors", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.ListDetectorsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/ListDetectors", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.ListDetectorsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/ResetComponentSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.ResetComponentSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/ResetComponentSettings", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/ResetSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.ResetSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/ResetSettings", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/UpdateComponentSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.UpdateComponentSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/UpdateComponentSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.ComponentSettings)(nil)).Elem())
	method2RequestType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/UpdateSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.UpdateSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.securitycenter.settings.v1beta1.SecurityCenterSettingsService/UpdateSettings", reflect.TypeOf((*google_cloud_securitycenter_settings_v1beta1_settingspb.Settings)(nil)).Elem())
}