// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	google_cloud_speech_v2_speechpb "cloud.google.com/go/speech/apiv2/speechpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.speech.v2.Speech", "speech.googleapis.com")
	method2RequestType.Store("google.cloud.speech.v2.Speech/BatchRecognize", reflect.TypeOf((*google_cloud_speech_v2_speechpb.BatchRecognizeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/BatchRecognize", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/CreateCustomClass", reflect.TypeOf((*google_cloud_speech_v2_speechpb.CreateCustomClassRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/CreateCustomClass", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/CreatePhraseSet", reflect.TypeOf((*google_cloud_speech_v2_speechpb.CreatePhraseSetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/CreatePhraseSet", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/CreateRecognizer", reflect.TypeOf((*google_cloud_speech_v2_speechpb.CreateRecognizerRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/CreateRecognizer", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/DeleteCustomClass", reflect.TypeOf((*google_cloud_speech_v2_speechpb.DeleteCustomClassRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/DeleteCustomClass", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/DeletePhraseSet", reflect.TypeOf((*google_cloud_speech_v2_speechpb.DeletePhraseSetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/DeletePhraseSet", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/DeleteRecognizer", reflect.TypeOf((*google_cloud_speech_v2_speechpb.DeleteRecognizerRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/DeleteRecognizer", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/GetConfig", reflect.TypeOf((*google_cloud_speech_v2_speechpb.GetConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/GetConfig", reflect.TypeOf((*google_cloud_speech_v2_speechpb.Config)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/GetCustomClass", reflect.TypeOf((*google_cloud_speech_v2_speechpb.GetCustomClassRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/GetCustomClass", reflect.TypeOf((*google_cloud_speech_v2_speechpb.CustomClass)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/GetPhraseSet", reflect.TypeOf((*google_cloud_speech_v2_speechpb.GetPhraseSetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/GetPhraseSet", reflect.TypeOf((*google_cloud_speech_v2_speechpb.PhraseSet)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/GetRecognizer", reflect.TypeOf((*google_cloud_speech_v2_speechpb.GetRecognizerRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/GetRecognizer", reflect.TypeOf((*google_cloud_speech_v2_speechpb.Recognizer)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/ListCustomClasses", reflect.TypeOf((*google_cloud_speech_v2_speechpb.ListCustomClassesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/ListCustomClasses", reflect.TypeOf((*google_cloud_speech_v2_speechpb.ListCustomClassesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/ListPhraseSets", reflect.TypeOf((*google_cloud_speech_v2_speechpb.ListPhraseSetsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/ListPhraseSets", reflect.TypeOf((*google_cloud_speech_v2_speechpb.ListPhraseSetsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/ListRecognizers", reflect.TypeOf((*google_cloud_speech_v2_speechpb.ListRecognizersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/ListRecognizers", reflect.TypeOf((*google_cloud_speech_v2_speechpb.ListRecognizersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/Recognize", reflect.TypeOf((*google_cloud_speech_v2_speechpb.RecognizeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/Recognize", reflect.TypeOf((*google_cloud_speech_v2_speechpb.RecognizeResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/StreamingRecognize", reflect.TypeOf((*google_cloud_speech_v2_speechpb.StreamingRecognizeRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/StreamingRecognize", reflect.TypeOf((*google_cloud_speech_v2_speechpb.StreamingRecognizeResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/UndeleteCustomClass", reflect.TypeOf((*google_cloud_speech_v2_speechpb.UndeleteCustomClassRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/UndeleteCustomClass", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/UndeletePhraseSet", reflect.TypeOf((*google_cloud_speech_v2_speechpb.UndeletePhraseSetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/UndeletePhraseSet", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/UndeleteRecognizer", reflect.TypeOf((*google_cloud_speech_v2_speechpb.UndeleteRecognizerRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/UndeleteRecognizer", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/UpdateConfig", reflect.TypeOf((*google_cloud_speech_v2_speechpb.UpdateConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/UpdateConfig", reflect.TypeOf((*google_cloud_speech_v2_speechpb.Config)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/UpdateCustomClass", reflect.TypeOf((*google_cloud_speech_v2_speechpb.UpdateCustomClassRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/UpdateCustomClass", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/UpdatePhraseSet", reflect.TypeOf((*google_cloud_speech_v2_speechpb.UpdatePhraseSetRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/UpdatePhraseSet", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.speech.v2.Speech/UpdateRecognizer", reflect.TypeOf((*google_cloud_speech_v2_speechpb.UpdateRecognizerRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.speech.v2.Speech/UpdateRecognizer", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
}