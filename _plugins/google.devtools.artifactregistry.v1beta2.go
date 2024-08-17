// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_devtools_artifactregistry_v1beta2_artifactregistrypb "cloud.google.com/go/artifactregistry/apiv1beta2/artifactregistrypb"
	google_iam_v1_iampb "cloud.google.com/go/iam/apiv1/iampb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry", "artifactregistry.googleapis.com")
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/CreateRepository", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.CreateRepositoryRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/CreateRepository", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/CreateTag", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.CreateTagRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/CreateTag", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.Tag)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/DeletePackage", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.DeletePackageRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/DeletePackage", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/DeleteRepository", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.DeleteRepositoryRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/DeleteRepository", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/DeleteTag", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.DeleteTagRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/DeleteTag", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/DeleteVersion", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.DeleteVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/DeleteVersion", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/GetFile", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.GetFileRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/GetFile", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.File)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.GetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/GetPackage", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.GetPackageRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/GetPackage", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.Package)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/GetProjectSettings", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.GetProjectSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/GetProjectSettings", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.ProjectSettings)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/GetRepository", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.GetRepositoryRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/GetRepository", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.Repository)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/GetTag", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.GetTagRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/GetTag", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.Tag)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/GetVersion", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.GetVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/GetVersion", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.Version)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/ImportAptArtifacts", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.ImportAptArtifactsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/ImportAptArtifacts", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/ImportYumArtifacts", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.ImportYumArtifactsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/ImportYumArtifacts", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/ListFiles", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.ListFilesRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/ListFiles", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.ListFilesResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/ListPackages", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.ListPackagesRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/ListPackages", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.ListPackagesResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/ListRepositories", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.ListRepositoriesRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/ListRepositories", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.ListRepositoriesResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/ListTags", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.ListTagsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/ListTags", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.ListTagsResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/ListVersions", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.ListVersionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/ListVersions", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.ListVersionsResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.SetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/UpdateProjectSettings", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.UpdateProjectSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/UpdateProjectSettings", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.ProjectSettings)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/UpdateRepository", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.UpdateRepositoryRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/UpdateRepository", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.Repository)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/UpdateTag", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.UpdateTagRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1beta2.ArtifactRegistry/UpdateTag", reflect.TypeOf((*google_devtools_artifactregistry_v1beta2_artifactregistrypb.Tag)(nil)).Elem())
}