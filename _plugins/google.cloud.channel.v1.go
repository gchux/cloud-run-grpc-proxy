// Code generated by grpc-proxy/proto/cmd/codegen.go; DO NOT EDIT.

package main

import (
	"reflect"

	"github.com/zhangyunhao116/skipmap"

	google_cloud_channel_v1_channelpb "cloud.google.com/go/channel/apiv1/channelpb"
	google_longrunning_longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func Load(service2Host *skipmap.OrderedMap[string, string], method2RequestType, method2ResponseType *skipmap.OrderedMap[string, reflect.Type]) {
	service2Host.Store("google.cloud.channel.v1.CloudChannelService", "cloudchannel.googleapis.com")
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ActivateEntitlement", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ActivateEntitlementRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ActivateEntitlement", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/CancelEntitlement", reflect.TypeOf((*google_cloud_channel_v1_channelpb.CancelEntitlementRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/CancelEntitlement", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ChangeOffer", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ChangeOfferRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ChangeOffer", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ChangeParameters", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ChangeParametersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ChangeParameters", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ChangeRenewalSettings", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ChangeRenewalSettingsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ChangeRenewalSettings", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/CheckCloudIdentityAccountsExist", reflect.TypeOf((*google_cloud_channel_v1_channelpb.CheckCloudIdentityAccountsExistRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/CheckCloudIdentityAccountsExist", reflect.TypeOf((*google_cloud_channel_v1_channelpb.CheckCloudIdentityAccountsExistResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/CreateChannelPartnerLink", reflect.TypeOf((*google_cloud_channel_v1_channelpb.CreateChannelPartnerLinkRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/CreateChannelPartnerLink", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ChannelPartnerLink)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/CreateCustomer", reflect.TypeOf((*google_cloud_channel_v1_channelpb.CreateCustomerRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/CreateCustomer", reflect.TypeOf((*google_cloud_channel_v1_channelpb.Customer)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/CreateCustomerRepricingConfig", reflect.TypeOf((*google_cloud_channel_v1_channelpb.CreateCustomerRepricingConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/CreateCustomerRepricingConfig", reflect.TypeOf((*google_cloud_channel_v1_channelpb.CustomerRepricingConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/CreateEntitlement", reflect.TypeOf((*google_cloud_channel_v1_channelpb.CreateEntitlementRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/CreateEntitlement", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/DeleteCustomer", reflect.TypeOf((*google_cloud_channel_v1_channelpb.DeleteCustomerRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/DeleteCustomer", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/DeleteCustomerRepricingConfig", reflect.TypeOf((*google_cloud_channel_v1_channelpb.DeleteCustomerRepricingConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/DeleteCustomerRepricingConfig", reflect.TypeOf((*emptypb.Empty)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/GetChannelPartnerLink", reflect.TypeOf((*google_cloud_channel_v1_channelpb.GetChannelPartnerLinkRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/GetChannelPartnerLink", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ChannelPartnerLink)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/GetChannelPartnerRepricingConfig", reflect.TypeOf((*google_cloud_channel_v1_channelpb.GetChannelPartnerRepricingConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/GetChannelPartnerRepricingConfig", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ChannelPartnerRepricingConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/GetCustomer", reflect.TypeOf((*google_cloud_channel_v1_channelpb.GetCustomerRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/GetCustomer", reflect.TypeOf((*google_cloud_channel_v1_channelpb.Customer)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/GetCustomerRepricingConfig", reflect.TypeOf((*google_cloud_channel_v1_channelpb.GetCustomerRepricingConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/GetCustomerRepricingConfig", reflect.TypeOf((*google_cloud_channel_v1_channelpb.CustomerRepricingConfig)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/GetEntitlement", reflect.TypeOf((*google_cloud_channel_v1_channelpb.GetEntitlementRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/GetEntitlement", reflect.TypeOf((*google_cloud_channel_v1_channelpb.Entitlement)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ImportCustomer", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ImportCustomerRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ImportCustomer", reflect.TypeOf((*google_cloud_channel_v1_channelpb.Customer)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ListChannelPartnerLinks", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListChannelPartnerLinksRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ListChannelPartnerLinks", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListChannelPartnerLinksResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ListCustomerRepricingConfigs", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListCustomerRepricingConfigsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ListCustomerRepricingConfigs", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListCustomerRepricingConfigsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ListCustomers", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListCustomersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ListCustomers", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListCustomersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ListEntitlementChanges", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListEntitlementChangesRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ListEntitlementChanges", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListEntitlementChangesResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ListEntitlements", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListEntitlementsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ListEntitlements", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListEntitlementsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ListOffers", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListOffersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ListOffers", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListOffersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ListProducts", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListProductsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ListProducts", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListProductsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ListPurchasableOffers", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListPurchasableOffersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ListPurchasableOffers", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListPurchasableOffersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ListPurchasableSkus", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListPurchasableSkusRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ListPurchasableSkus", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListPurchasableSkusResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ListSkuGroupBillableSkus", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListSkuGroupBillableSkusRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ListSkuGroupBillableSkus", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListSkuGroupBillableSkusResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ListSkuGroups", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListSkuGroupsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ListSkuGroups", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListSkuGroupsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ListSkus", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListSkusRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ListSkus", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListSkusResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ListSubscribers", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListSubscribersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ListSubscribers", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListSubscribersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ListTransferableOffers", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListTransferableOffersRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ListTransferableOffers", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListTransferableOffersResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ListTransferableSkus", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListTransferableSkusRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ListTransferableSkus", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ListTransferableSkusResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/LookupOffer", reflect.TypeOf((*google_cloud_channel_v1_channelpb.LookupOfferRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/LookupOffer", reflect.TypeOf((*google_cloud_channel_v1_channelpb.Offer)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/ProvisionCloudIdentity", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ProvisionCloudIdentityRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/ProvisionCloudIdentity", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/QueryEligibleBillingAccounts", reflect.TypeOf((*google_cloud_channel_v1_channelpb.QueryEligibleBillingAccountsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/QueryEligibleBillingAccounts", reflect.TypeOf((*google_cloud_channel_v1_channelpb.QueryEligibleBillingAccountsResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/RegisterSubscriber", reflect.TypeOf((*google_cloud_channel_v1_channelpb.RegisterSubscriberRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/RegisterSubscriber", reflect.TypeOf((*google_cloud_channel_v1_channelpb.RegisterSubscriberResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/StartPaidService", reflect.TypeOf((*google_cloud_channel_v1_channelpb.StartPaidServiceRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/StartPaidService", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/SuspendEntitlement", reflect.TypeOf((*google_cloud_channel_v1_channelpb.SuspendEntitlementRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/SuspendEntitlement", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/TransferEntitlements", reflect.TypeOf((*google_cloud_channel_v1_channelpb.TransferEntitlementsRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/TransferEntitlements", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/TransferEntitlementsToGoogle", reflect.TypeOf((*google_cloud_channel_v1_channelpb.TransferEntitlementsToGoogleRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/TransferEntitlementsToGoogle", reflect.TypeOf((*google_longrunning_longrunningpb.Operation)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/UnregisterSubscriber", reflect.TypeOf((*google_cloud_channel_v1_channelpb.UnregisterSubscriberRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/UnregisterSubscriber", reflect.TypeOf((*google_cloud_channel_v1_channelpb.UnregisterSubscriberResponse)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/UpdateChannelPartnerLink", reflect.TypeOf((*google_cloud_channel_v1_channelpb.UpdateChannelPartnerLinkRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/UpdateChannelPartnerLink", reflect.TypeOf((*google_cloud_channel_v1_channelpb.ChannelPartnerLink)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/UpdateCustomer", reflect.TypeOf((*google_cloud_channel_v1_channelpb.UpdateCustomerRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/UpdateCustomer", reflect.TypeOf((*google_cloud_channel_v1_channelpb.Customer)(nil)).Elem())
	method2RequestType.Store("google.cloud.channel.v1.CloudChannelService/UpdateCustomerRepricingConfig", reflect.TypeOf((*google_cloud_channel_v1_channelpb.UpdateCustomerRepricingConfigRequest)(nil)).Elem())
	method2ResponseType.Store("google.cloud.channel.v1.CloudChannelService/UpdateCustomerRepricingConfig", reflect.TypeOf((*google_cloud_channel_v1_channelpb.CustomerRepricingConfig)(nil)).Elem())
}
