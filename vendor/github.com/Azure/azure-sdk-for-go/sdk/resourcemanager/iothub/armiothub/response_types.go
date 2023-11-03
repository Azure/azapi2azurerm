//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armiothub

// CertificatesClientCreateOrUpdateResponse contains the response from method CertificatesClient.CreateOrUpdate.
type CertificatesClientCreateOrUpdateResponse struct {
	CertificateDescription
}

// CertificatesClientDeleteResponse contains the response from method CertificatesClient.Delete.
type CertificatesClientDeleteResponse struct {
	// placeholder for future response values
}

// CertificatesClientGenerateVerificationCodeResponse contains the response from method CertificatesClient.GenerateVerificationCode.
type CertificatesClientGenerateVerificationCodeResponse struct {
	CertificateWithNonceDescription
}

// CertificatesClientGetResponse contains the response from method CertificatesClient.Get.
type CertificatesClientGetResponse struct {
	CertificateDescription
}

// CertificatesClientListByIotHubResponse contains the response from method CertificatesClient.ListByIotHub.
type CertificatesClientListByIotHubResponse struct {
	CertificateListDescription
}

// CertificatesClientVerifyResponse contains the response from method CertificatesClient.Verify.
type CertificatesClientVerifyResponse struct {
	CertificateDescription
}

// ClientManualFailoverResponse contains the response from method Client.BeginManualFailover.
type ClientManualFailoverResponse struct {
	// placeholder for future response values
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.List.
type PrivateEndpointConnectionsClientListResponse struct {
	// The list of private endpoint connections for an IotHub
	PrivateEndpointConnectionArray []*PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientUpdateResponse contains the response from method PrivateEndpointConnectionsClient.BeginUpdate.
type PrivateEndpointConnectionsClientUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	GroupIDInformation
}

// PrivateLinkResourcesClientListResponse contains the response from method PrivateLinkResourcesClient.List.
type PrivateLinkResourcesClientListResponse struct {
	PrivateLinkResources
}

// ResourceClientCheckNameAvailabilityResponse contains the response from method ResourceClient.CheckNameAvailability.
type ResourceClientCheckNameAvailabilityResponse struct {
	NameAvailabilityInfo
}

// ResourceClientCreateEventHubConsumerGroupResponse contains the response from method ResourceClient.CreateEventHubConsumerGroup.
type ResourceClientCreateEventHubConsumerGroupResponse struct {
	EventHubConsumerGroupInfo
}

// ResourceClientCreateOrUpdateResponse contains the response from method ResourceClient.BeginCreateOrUpdate.
type ResourceClientCreateOrUpdateResponse struct {
	Description
}

// ResourceClientDeleteEventHubConsumerGroupResponse contains the response from method ResourceClient.DeleteEventHubConsumerGroup.
type ResourceClientDeleteEventHubConsumerGroupResponse struct {
	// placeholder for future response values
}

// ResourceClientDeleteResponse contains the response from method ResourceClient.BeginDelete.
type ResourceClientDeleteResponse struct {
	Description
}

// ResourceClientExportDevicesResponse contains the response from method ResourceClient.ExportDevices.
type ResourceClientExportDevicesResponse struct {
	JobResponse
}

// ResourceClientGetEndpointHealthResponse contains the response from method ResourceClient.NewGetEndpointHealthPager.
type ResourceClientGetEndpointHealthResponse struct {
	EndpointHealthDataListResult
}

// ResourceClientGetEventHubConsumerGroupResponse contains the response from method ResourceClient.GetEventHubConsumerGroup.
type ResourceClientGetEventHubConsumerGroupResponse struct {
	EventHubConsumerGroupInfo
}

// ResourceClientGetJobResponse contains the response from method ResourceClient.GetJob.
type ResourceClientGetJobResponse struct {
	JobResponse
}

// ResourceClientGetKeysForKeyNameResponse contains the response from method ResourceClient.GetKeysForKeyName.
type ResourceClientGetKeysForKeyNameResponse struct {
	SharedAccessSignatureAuthorizationRule
}

// ResourceClientGetQuotaMetricsResponse contains the response from method ResourceClient.NewGetQuotaMetricsPager.
type ResourceClientGetQuotaMetricsResponse struct {
	QuotaMetricInfoListResult
}

// ResourceClientGetResponse contains the response from method ResourceClient.Get.
type ResourceClientGetResponse struct {
	Description
}

// ResourceClientGetStatsResponse contains the response from method ResourceClient.GetStats.
type ResourceClientGetStatsResponse struct {
	RegistryStatistics
}

// ResourceClientGetValidSKUsResponse contains the response from method ResourceClient.NewGetValidSKUsPager.
type ResourceClientGetValidSKUsResponse struct {
	SKUDescriptionListResult
}

// ResourceClientImportDevicesResponse contains the response from method ResourceClient.ImportDevices.
type ResourceClientImportDevicesResponse struct {
	JobResponse
}

// ResourceClientListByResourceGroupResponse contains the response from method ResourceClient.NewListByResourceGroupPager.
type ResourceClientListByResourceGroupResponse struct {
	DescriptionListResult
}

// ResourceClientListBySubscriptionResponse contains the response from method ResourceClient.NewListBySubscriptionPager.
type ResourceClientListBySubscriptionResponse struct {
	DescriptionListResult
}

// ResourceClientListEventHubConsumerGroupsResponse contains the response from method ResourceClient.NewListEventHubConsumerGroupsPager.
type ResourceClientListEventHubConsumerGroupsResponse struct {
	EventHubConsumerGroupsListResult
}

// ResourceClientListJobsResponse contains the response from method ResourceClient.NewListJobsPager.
type ResourceClientListJobsResponse struct {
	JobResponseListResult
}

// ResourceClientListKeysResponse contains the response from method ResourceClient.NewListKeysPager.
type ResourceClientListKeysResponse struct {
	SharedAccessSignatureAuthorizationRuleListResult
}

// ResourceClientTestAllRoutesResponse contains the response from method ResourceClient.TestAllRoutes.
type ResourceClientTestAllRoutesResponse struct {
	TestAllRoutesResult
}

// ResourceClientTestRouteResponse contains the response from method ResourceClient.TestRoute.
type ResourceClientTestRouteResponse struct {
	TestRouteResult
}

// ResourceClientUpdateResponse contains the response from method ResourceClient.BeginUpdate.
type ResourceClientUpdateResponse struct {
	Description
}

// ResourceProviderCommonClientGetSubscriptionQuotaResponse contains the response from method ResourceProviderCommonClient.GetSubscriptionQuota.
type ResourceProviderCommonClientGetSubscriptionQuotaResponse struct {
	UserSubscriptionQuotaListResult
}
