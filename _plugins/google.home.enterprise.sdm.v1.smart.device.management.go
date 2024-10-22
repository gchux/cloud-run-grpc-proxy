// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_home_enterprise_sdm_v1_smart_device_management_servicepb "google.golang.org/genproto/googleapis/home/enterprise/sdm/v1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.home.enterprise.sdm.v1.SmartDeviceManagementService", "smartdevicemanagement.googleapis.com")
	method2RequestType.Store("google.home.enterprise.sdm.v1.SmartDeviceManagementService/ExecuteDeviceCommand", reflect.TypeOf((*google_home_enterprise_sdm_v1_smart_device_management_servicepb.ExecuteDeviceCommandRequest)(nil)).Elem())
	method2ResponseType.Store("google.home.enterprise.sdm.v1.SmartDeviceManagementService/ExecuteDeviceCommand", reflect.TypeOf((*google_home_enterprise_sdm_v1_smart_device_management_servicepb.ExecuteDeviceCommandResponse)(nil)).Elem())
	method2RequestType.Store("google.home.enterprise.sdm.v1.SmartDeviceManagementService/GetDevice", reflect.TypeOf((*google_home_enterprise_sdm_v1_smart_device_management_servicepb.GetDeviceRequest)(nil)).Elem())
	method2ResponseType.Store("google.home.enterprise.sdm.v1.SmartDeviceManagementService/GetDevice", reflect.TypeOf((*google_home_enterprise_sdm_v1_smart_device_management_servicepb.Device)(nil)).Elem())
	method2RequestType.Store("google.home.enterprise.sdm.v1.SmartDeviceManagementService/GetRoom", reflect.TypeOf((*google_home_enterprise_sdm_v1_smart_device_management_servicepb.GetRoomRequest)(nil)).Elem())
	method2ResponseType.Store("google.home.enterprise.sdm.v1.SmartDeviceManagementService/GetRoom", reflect.TypeOf((*google_home_enterprise_sdm_v1_smart_device_management_servicepb.Room)(nil)).Elem())
	method2RequestType.Store("google.home.enterprise.sdm.v1.SmartDeviceManagementService/GetStructure", reflect.TypeOf((*google_home_enterprise_sdm_v1_smart_device_management_servicepb.GetStructureRequest)(nil)).Elem())
	method2ResponseType.Store("google.home.enterprise.sdm.v1.SmartDeviceManagementService/GetStructure", reflect.TypeOf((*google_home_enterprise_sdm_v1_smart_device_management_servicepb.Structure)(nil)).Elem())
	method2RequestType.Store("google.home.enterprise.sdm.v1.SmartDeviceManagementService/ListDevices", reflect.TypeOf((*google_home_enterprise_sdm_v1_smart_device_management_servicepb.ListDevicesRequest)(nil)).Elem())
	method2ResponseType.Store("google.home.enterprise.sdm.v1.SmartDeviceManagementService/ListDevices", reflect.TypeOf((*google_home_enterprise_sdm_v1_smart_device_management_servicepb.ListDevicesResponse)(nil)).Elem())
	method2RequestType.Store("google.home.enterprise.sdm.v1.SmartDeviceManagementService/ListRooms", reflect.TypeOf((*google_home_enterprise_sdm_v1_smart_device_management_servicepb.ListRoomsRequest)(nil)).Elem())
	method2ResponseType.Store("google.home.enterprise.sdm.v1.SmartDeviceManagementService/ListRooms", reflect.TypeOf((*google_home_enterprise_sdm_v1_smart_device_management_servicepb.ListRoomsResponse)(nil)).Elem())
	method2RequestType.Store("google.home.enterprise.sdm.v1.SmartDeviceManagementService/ListStructures", reflect.TypeOf((*google_home_enterprise_sdm_v1_smart_device_management_servicepb.ListStructuresRequest)(nil)).Elem())
	method2ResponseType.Store("google.home.enterprise.sdm.v1.SmartDeviceManagementService/ListStructures", reflect.TypeOf((*google_home_enterprise_sdm_v1_smart_device_management_servicepb.ListStructuresResponse)(nil)).Elem())
}
