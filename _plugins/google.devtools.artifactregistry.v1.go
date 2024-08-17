// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_devtools_artifactregistry_v1_artifactregistrypb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	google_iam_v1_iampb "cloud.google.com/go/iam/apiv1/iampb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.devtools.artifactregistry.v1.ArtifactRegistry", "artifactregistry.googleapis.com")
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/BatchDeleteVersions", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.BatchDeleteVersionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/BatchDeleteVersions", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/CreateRepository", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.CreateRepositoryRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/CreateRepository", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/CreateTag", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.CreateTagRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/CreateTag", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.Tag)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/DeletePackage", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.DeletePackageRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/DeletePackage", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/DeleteRepository", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.DeleteRepositoryRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/DeleteRepository", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/DeleteTag", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.DeleteTagRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/DeleteTag", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/DeleteVersion", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.DeleteVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/DeleteVersion", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetDockerImage", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.GetDockerImageRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetDockerImage", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.DockerImage)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetFile", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.GetFileRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetFile", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.File)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.GetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetMavenArtifact", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.GetMavenArtifactRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetMavenArtifact", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.MavenArtifact)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetNpmPackage", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.GetNpmPackageRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetNpmPackage", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.NpmPackage)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetPackage", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.GetPackageRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetPackage", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.Package)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetProjectSettings", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.GetProjectSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetProjectSettings", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ProjectSettings)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetPythonPackage", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.GetPythonPackageRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetPythonPackage", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.PythonPackage)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetRepository", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.GetRepositoryRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetRepository", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.Repository)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetTag", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.GetTagRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetTag", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.Tag)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetVPCSCConfig", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.GetVPCSCConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetVPCSCConfig", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.VPCSCConfig)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetVersion", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.GetVersionRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/GetVersion", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.Version)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ImportAptArtifacts", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ImportAptArtifactsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ImportAptArtifacts", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ImportYumArtifacts", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ImportYumArtifactsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ImportYumArtifacts", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListDockerImages", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListDockerImagesRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListDockerImages", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListDockerImagesResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListFiles", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListFilesRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListFiles", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListFilesResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListMavenArtifacts", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListMavenArtifactsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListMavenArtifacts", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListMavenArtifactsResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListNpmPackages", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListNpmPackagesRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListNpmPackages", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListNpmPackagesResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListPackages", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListPackagesRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListPackages", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListPackagesResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListPythonPackages", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListPythonPackagesRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListPythonPackages", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListPythonPackagesResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListRepositories", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListRepositoriesRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListRepositories", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListRepositoriesResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListTags", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListTagsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListTags", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListTagsResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListVersions", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListVersionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/ListVersions", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ListVersionsResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.SetIamPolicyRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/SetIamPolicy", reflect.TypeOf((*google_iam_v1_iampb.Policy)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/TestIamPermissions", reflect.TypeOf((*google_iam_v1_iampb.TestIamPermissionsResponse)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/UpdateProjectSettings", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.UpdateProjectSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/UpdateProjectSettings", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.ProjectSettings)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/UpdateRepository", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.UpdateRepositoryRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/UpdateRepository", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.Repository)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/UpdateTag", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.UpdateTagRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/UpdateTag", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.Tag)(nil)).Elem())
	method2RequestType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/UpdateVPCSCConfig", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.UpdateVPCSCConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.devtools.artifactregistry.v1.ArtifactRegistry/UpdateVPCSCConfig", reflect.TypeOf((*google_devtools_artifactregistry_v1_artifactregistrypb.VPCSCConfig)(nil)).Elem())
}