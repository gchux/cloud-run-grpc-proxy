// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_chat_v1_chatpb "cloud.google.com/go/chat/apiv1/chatpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.chat.v1.ChatService", "chat.googleapis.com")
	method2RequestType.Store("google.chat.v1.ChatService/CompleteImportSpace", reflect.TypeOf((*google_chat_v1_chatpb.CompleteImportSpaceRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/CompleteImportSpace", reflect.TypeOf((*google_chat_v1_chatpb.CompleteImportSpaceResponse)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/CreateMembership", reflect.TypeOf((*google_chat_v1_chatpb.CreateMembershipRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/CreateMembership", reflect.TypeOf((*google_chat_v1_chatpb.Membership)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/CreateMessage", reflect.TypeOf((*google_chat_v1_chatpb.CreateMessageRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/CreateMessage", reflect.TypeOf((*google_chat_v1_chatpb.Message)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/CreateReaction", reflect.TypeOf((*google_chat_v1_chatpb.CreateReactionRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/CreateReaction", reflect.TypeOf((*google_chat_v1_chatpb.Reaction)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/CreateSpace", reflect.TypeOf((*google_chat_v1_chatpb.CreateSpaceRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/CreateSpace", reflect.TypeOf((*google_chat_v1_chatpb.Space)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/DeleteMembership", reflect.TypeOf((*google_chat_v1_chatpb.DeleteMembershipRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/DeleteMembership", reflect.TypeOf((*google_chat_v1_chatpb.Membership)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/DeleteMessage", reflect.TypeOf((*google_chat_v1_chatpb.DeleteMessageRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/DeleteMessage", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/DeleteReaction", reflect.TypeOf((*google_chat_v1_chatpb.DeleteReactionRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/DeleteReaction", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/DeleteSpace", reflect.TypeOf((*google_chat_v1_chatpb.DeleteSpaceRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/DeleteSpace", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/FindDirectMessage", reflect.TypeOf((*google_chat_v1_chatpb.FindDirectMessageRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/FindDirectMessage", reflect.TypeOf((*google_chat_v1_chatpb.Space)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/GetAttachment", reflect.TypeOf((*google_chat_v1_chatpb.GetAttachmentRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/GetAttachment", reflect.TypeOf((*google_chat_v1_chatpb.Attachment)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/GetMembership", reflect.TypeOf((*google_chat_v1_chatpb.GetMembershipRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/GetMembership", reflect.TypeOf((*google_chat_v1_chatpb.Membership)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/GetMessage", reflect.TypeOf((*google_chat_v1_chatpb.GetMessageRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/GetMessage", reflect.TypeOf((*google_chat_v1_chatpb.Message)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/GetSpace", reflect.TypeOf((*google_chat_v1_chatpb.GetSpaceRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/GetSpace", reflect.TypeOf((*google_chat_v1_chatpb.Space)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/GetSpaceEvent", reflect.TypeOf((*google_chat_v1_chatpb.GetSpaceEventRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/GetSpaceEvent", reflect.TypeOf((*google_chat_v1_chatpb.SpaceEvent)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/GetSpaceReadState", reflect.TypeOf((*google_chat_v1_chatpb.GetSpaceReadStateRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/GetSpaceReadState", reflect.TypeOf((*google_chat_v1_chatpb.SpaceReadState)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/GetThreadReadState", reflect.TypeOf((*google_chat_v1_chatpb.GetThreadReadStateRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/GetThreadReadState", reflect.TypeOf((*google_chat_v1_chatpb.ThreadReadState)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/ListMemberships", reflect.TypeOf((*google_chat_v1_chatpb.ListMembershipsRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/ListMemberships", reflect.TypeOf((*google_chat_v1_chatpb.ListMembershipsResponse)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/ListMessages", reflect.TypeOf((*google_chat_v1_chatpb.ListMessagesRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/ListMessages", reflect.TypeOf((*google_chat_v1_chatpb.ListMessagesResponse)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/ListReactions", reflect.TypeOf((*google_chat_v1_chatpb.ListReactionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/ListReactions", reflect.TypeOf((*google_chat_v1_chatpb.ListReactionsResponse)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/ListSpaceEvents", reflect.TypeOf((*google_chat_v1_chatpb.ListSpaceEventsRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/ListSpaceEvents", reflect.TypeOf((*google_chat_v1_chatpb.ListSpaceEventsResponse)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/ListSpaces", reflect.TypeOf((*google_chat_v1_chatpb.ListSpacesRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/ListSpaces", reflect.TypeOf((*google_chat_v1_chatpb.ListSpacesResponse)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/SetUpSpace", reflect.TypeOf((*google_chat_v1_chatpb.SetUpSpaceRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/SetUpSpace", reflect.TypeOf((*google_chat_v1_chatpb.Space)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/UpdateMembership", reflect.TypeOf((*google_chat_v1_chatpb.UpdateMembershipRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/UpdateMembership", reflect.TypeOf((*google_chat_v1_chatpb.Membership)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/UpdateMessage", reflect.TypeOf((*google_chat_v1_chatpb.UpdateMessageRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/UpdateMessage", reflect.TypeOf((*google_chat_v1_chatpb.Message)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/UpdateSpace", reflect.TypeOf((*google_chat_v1_chatpb.UpdateSpaceRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/UpdateSpace", reflect.TypeOf((*google_chat_v1_chatpb.Space)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/UpdateSpaceReadState", reflect.TypeOf((*google_chat_v1_chatpb.UpdateSpaceReadStateRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/UpdateSpaceReadState", reflect.TypeOf((*google_chat_v1_chatpb.SpaceReadState)(nil)).Elem())
	method2RequestType.Store("google.chat.v1.ChatService/UploadAttachment", reflect.TypeOf((*google_chat_v1_chatpb.UploadAttachmentRequest)(nil)).Elem())
	method2ResponseType.Store("google.chat.v1.ChatService/UploadAttachment", reflect.TypeOf((*google_chat_v1_chatpb.UploadAttachmentResponse)(nil)).Elem())
}
