// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_deploy_v1_deploypb "cloud.google.com/go/deploy/apiv1/deploypb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.deploy.v1.CloudDeploy", "clouddeploy.googleapis.com")
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/AbandonRelease", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.AbandonReleaseRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/AbandonRelease", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.AbandonReleaseResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/AdvanceRollout", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.AdvanceRolloutRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/AdvanceRollout", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.AdvanceRolloutResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/ApproveRollout", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ApproveRolloutRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/ApproveRollout", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ApproveRolloutResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/CancelAutomationRun", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.CancelAutomationRunRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/CancelAutomationRun", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.CancelAutomationRunResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/CancelRollout", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.CancelRolloutRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/CancelRollout", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.CancelRolloutResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/CreateAutomation", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.CreateAutomationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/CreateAutomation", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/CreateCustomTargetType", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.CreateCustomTargetTypeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/CreateCustomTargetType", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/CreateDeliveryPipeline", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.CreateDeliveryPipelineRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/CreateDeliveryPipeline", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/CreateDeployPolicy", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.CreateDeployPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/CreateDeployPolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/CreateRelease", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.CreateReleaseRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/CreateRelease", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/CreateRollout", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.CreateRolloutRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/CreateRollout", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/CreateTarget", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.CreateTargetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/CreateTarget", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/DeleteAutomation", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.DeleteAutomationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/DeleteAutomation", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/DeleteCustomTargetType", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.DeleteCustomTargetTypeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/DeleteCustomTargetType", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/DeleteDeliveryPipeline", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.DeleteDeliveryPipelineRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/DeleteDeliveryPipeline", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/DeleteDeployPolicy", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.DeleteDeployPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/DeleteDeployPolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/DeleteTarget", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.DeleteTargetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/DeleteTarget", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/GetAutomation", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.GetAutomationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/GetAutomation", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.Automation)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/GetAutomationRun", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.GetAutomationRunRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/GetAutomationRun", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.AutomationRun)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/GetConfig", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.GetConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/GetConfig", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.Config)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/GetCustomTargetType", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.GetCustomTargetTypeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/GetCustomTargetType", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.CustomTargetType)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/GetDeliveryPipeline", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.GetDeliveryPipelineRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/GetDeliveryPipeline", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.DeliveryPipeline)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/GetDeployPolicy", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.GetDeployPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/GetDeployPolicy", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.DeployPolicy)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/GetJobRun", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.GetJobRunRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/GetJobRun", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.JobRun)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/GetRelease", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.GetReleaseRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/GetRelease", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.Release)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/GetRollout", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.GetRolloutRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/GetRollout", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.Rollout)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/GetTarget", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.GetTargetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/GetTarget", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.Target)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/IgnoreJob", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.IgnoreJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/IgnoreJob", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.IgnoreJobResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/ListAutomationRuns", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListAutomationRunsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/ListAutomationRuns", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListAutomationRunsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/ListAutomations", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListAutomationsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/ListAutomations", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListAutomationsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/ListCustomTargetTypes", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListCustomTargetTypesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/ListCustomTargetTypes", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListCustomTargetTypesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/ListDeliveryPipelines", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListDeliveryPipelinesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/ListDeliveryPipelines", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListDeliveryPipelinesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/ListDeployPolicies", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListDeployPoliciesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/ListDeployPolicies", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListDeployPoliciesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/ListJobRuns", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListJobRunsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/ListJobRuns", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListJobRunsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/ListReleases", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListReleasesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/ListReleases", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListReleasesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/ListRollouts", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListRolloutsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/ListRollouts", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListRolloutsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/ListTargets", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListTargetsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/ListTargets", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.ListTargetsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/RetryJob", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.RetryJobRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/RetryJob", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.RetryJobResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/RollbackTarget", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.RollbackTargetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/RollbackTarget", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.RollbackTargetResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/TerminateJobRun", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.TerminateJobRunRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/TerminateJobRun", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.TerminateJobRunResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/UpdateAutomation", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.UpdateAutomationRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/UpdateAutomation", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/UpdateCustomTargetType", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.UpdateCustomTargetTypeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/UpdateCustomTargetType", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/UpdateDeliveryPipeline", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.UpdateDeliveryPipelineRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/UpdateDeliveryPipeline", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/UpdateDeployPolicy", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.UpdateDeployPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/UpdateDeployPolicy", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.deploy.v1.CloudDeploy/UpdateTarget", reflect.TypeOf((*google_cloud_deploy_v1_deploypb.UpdateTargetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.deploy.v1.CloudDeploy/UpdateTarget", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}
