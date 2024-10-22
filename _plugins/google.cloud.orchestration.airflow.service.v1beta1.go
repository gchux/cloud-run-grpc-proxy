// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_orchestration_airflow_service_v1beta1_environmentspb "google.golang.org/genproto/googleapis/cloud/orchestration/airflow/service/v1beta1"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments", "composer.googleapis.com")
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/CheckUpgrade", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.CheckUpgradeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/CheckUpgrade", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/CreateEnvironment", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.CreateEnvironmentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/CreateEnvironment", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/CreateUserWorkloadsConfigMap", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.CreateUserWorkloadsConfigMapRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/CreateUserWorkloadsConfigMap", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.UserWorkloadsConfigMap)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/CreateUserWorkloadsSecret", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.CreateUserWorkloadsSecretRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/CreateUserWorkloadsSecret", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.UserWorkloadsSecret)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/DatabaseFailover", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.DatabaseFailoverRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/DatabaseFailover", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/DeleteEnvironment", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.DeleteEnvironmentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/DeleteEnvironment", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/DeleteUserWorkloadsConfigMap", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.DeleteUserWorkloadsConfigMapRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/DeleteUserWorkloadsConfigMap", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/DeleteUserWorkloadsSecret", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.DeleteUserWorkloadsSecretRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/DeleteUserWorkloadsSecret", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/ExecuteAirflowCommand", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.ExecuteAirflowCommandRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/ExecuteAirflowCommand", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.ExecuteAirflowCommandResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/FetchDatabaseProperties", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.FetchDatabasePropertiesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/FetchDatabaseProperties", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.FetchDatabasePropertiesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/GetEnvironment", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.GetEnvironmentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/GetEnvironment", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.Environment)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/GetUserWorkloadsConfigMap", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.GetUserWorkloadsConfigMapRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/GetUserWorkloadsConfigMap", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.UserWorkloadsConfigMap)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/GetUserWorkloadsSecret", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.GetUserWorkloadsSecretRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/GetUserWorkloadsSecret", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.UserWorkloadsSecret)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/ListEnvironments", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.ListEnvironmentsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/ListEnvironments", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.ListEnvironmentsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/ListUserWorkloadsConfigMaps", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.ListUserWorkloadsConfigMapsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/ListUserWorkloadsConfigMaps", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.ListUserWorkloadsConfigMapsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/ListUserWorkloadsSecrets", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.ListUserWorkloadsSecretsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/ListUserWorkloadsSecrets", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.ListUserWorkloadsSecretsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/ListWorkloads", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.ListWorkloadsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/ListWorkloads", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.ListWorkloadsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/LoadSnapshot", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.LoadSnapshotRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/LoadSnapshot", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/PollAirflowCommand", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.PollAirflowCommandRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/PollAirflowCommand", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.PollAirflowCommandResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/RestartWebServer", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.RestartWebServerRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/RestartWebServer", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/SaveSnapshot", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.SaveSnapshotRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/SaveSnapshot", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/StopAirflowCommand", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.StopAirflowCommandRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/StopAirflowCommand", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.StopAirflowCommandResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/UpdateEnvironment", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.UpdateEnvironmentRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/UpdateEnvironment", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/UpdateUserWorkloadsConfigMap", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.UpdateUserWorkloadsConfigMapRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/UpdateUserWorkloadsConfigMap", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.UserWorkloadsConfigMap)(nil)).Elem())
	method2RequestType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/UpdateUserWorkloadsSecret", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.UpdateUserWorkloadsSecretRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.orchestration.airflow.service.v1beta1.Environments/UpdateUserWorkloadsSecret", reflect.TypeOf((*google_cloud_orchestration_airflow_service_v1beta1_environmentspb.UserWorkloadsSecret)(nil)).Elem())
}
