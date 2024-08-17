// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_shopping_merchant_notifications_v1beta_notificationspb "cloud.google.com/go/shopping/merchant/notifications/apiv1beta/notificationspb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.shopping.merchant.notifications.v1beta.NotificationsApiService", "merchantapi.googleapis.com")
	method2RequestType.Store("google.shopping.merchant.notifications.v1beta.NotificationsApiService/CreateNotificationSubscription", reflect.TypeOf((*google_shopping_merchant_notifications_v1beta_notificationspb.CreateNotificationSubscriptionRequest)(nil)).Elem())
	method2ResponseType.Store("google.shopping.merchant.notifications.v1beta.NotificationsApiService/CreateNotificationSubscription", reflect.TypeOf((*google_shopping_merchant_notifications_v1beta_notificationspb.NotificationSubscription)(nil)).Elem())
	method2RequestType.Store("google.shopping.merchant.notifications.v1beta.NotificationsApiService/DeleteNotificationSubscription", reflect.TypeOf((*google_shopping_merchant_notifications_v1beta_notificationspb.DeleteNotificationSubscriptionRequest)(nil)).Elem())
	method2ResponseType.Store("google.shopping.merchant.notifications.v1beta.NotificationsApiService/DeleteNotificationSubscription", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.shopping.merchant.notifications.v1beta.NotificationsApiService/GetNotificationSubscription", reflect.TypeOf((*google_shopping_merchant_notifications_v1beta_notificationspb.GetNotificationSubscriptionRequest)(nil)).Elem())
	method2ResponseType.Store("google.shopping.merchant.notifications.v1beta.NotificationsApiService/GetNotificationSubscription", reflect.TypeOf((*google_shopping_merchant_notifications_v1beta_notificationspb.NotificationSubscription)(nil)).Elem())
	method2RequestType.Store("google.shopping.merchant.notifications.v1beta.NotificationsApiService/ListNotificationSubscriptions", reflect.TypeOf((*google_shopping_merchant_notifications_v1beta_notificationspb.ListNotificationSubscriptionsRequest)(nil)).Elem())
	method2ResponseType.Store("google.shopping.merchant.notifications.v1beta.NotificationsApiService/ListNotificationSubscriptions", reflect.TypeOf((*google_shopping_merchant_notifications_v1beta_notificationspb.ListNotificationSubscriptionsResponse)(nil)).Elem())
	method2RequestType.Store("google.shopping.merchant.notifications.v1beta.NotificationsApiService/UpdateNotificationSubscription", reflect.TypeOf((*google_shopping_merchant_notifications_v1beta_notificationspb.UpdateNotificationSubscriptionRequest)(nil)).Elem())
	method2ResponseType.Store("google.shopping.merchant.notifications.v1beta.NotificationsApiService/UpdateNotificationSubscription", reflect.TypeOf((*google_shopping_merchant_notifications_v1beta_notificationspb.NotificationSubscription)(nil)).Elem())
}