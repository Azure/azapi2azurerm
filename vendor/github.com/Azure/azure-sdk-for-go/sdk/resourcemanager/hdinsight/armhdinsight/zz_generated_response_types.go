//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsight

// ApplicationsClientCreateResponse contains the response from method ApplicationsClient.Create.
type ApplicationsClientCreateResponse struct {
	Application
}

// ApplicationsClientDeleteResponse contains the response from method ApplicationsClient.Delete.
type ApplicationsClientDeleteResponse struct {
	// placeholder for future response values
}

// ApplicationsClientGetAzureAsyncOperationStatusResponse contains the response from method ApplicationsClient.GetAzureAsyncOperationStatus.
type ApplicationsClientGetAzureAsyncOperationStatusResponse struct {
	AsyncOperationResult
}

// ApplicationsClientGetResponse contains the response from method ApplicationsClient.Get.
type ApplicationsClientGetResponse struct {
	Application
}

// ApplicationsClientListByClusterResponse contains the response from method ApplicationsClient.ListByCluster.
type ApplicationsClientListByClusterResponse struct {
	ApplicationListResult
}

// ClustersClientCreateResponse contains the response from method ClustersClient.Create.
type ClustersClientCreateResponse struct {
	Cluster
}

// ClustersClientDeleteResponse contains the response from method ClustersClient.Delete.
type ClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// ClustersClientExecuteScriptActionsResponse contains the response from method ClustersClient.ExecuteScriptActions.
type ClustersClientExecuteScriptActionsResponse struct {
	// placeholder for future response values
}

// ClustersClientGetAzureAsyncOperationStatusResponse contains the response from method ClustersClient.GetAzureAsyncOperationStatus.
type ClustersClientGetAzureAsyncOperationStatusResponse struct {
	AsyncOperationResult
}

// ClustersClientGetGatewaySettingsResponse contains the response from method ClustersClient.GetGatewaySettings.
type ClustersClientGetGatewaySettingsResponse struct {
	GatewaySettings
}

// ClustersClientGetResponse contains the response from method ClustersClient.Get.
type ClustersClientGetResponse struct {
	Cluster
}

// ClustersClientListByResourceGroupResponse contains the response from method ClustersClient.ListByResourceGroup.
type ClustersClientListByResourceGroupResponse struct {
	ClusterListResult
}

// ClustersClientListResponse contains the response from method ClustersClient.List.
type ClustersClientListResponse struct {
	ClusterListResult
}

// ClustersClientResizeResponse contains the response from method ClustersClient.Resize.
type ClustersClientResizeResponse struct {
	// placeholder for future response values
}

// ClustersClientRotateDiskEncryptionKeyResponse contains the response from method ClustersClient.RotateDiskEncryptionKey.
type ClustersClientRotateDiskEncryptionKeyResponse struct {
	// placeholder for future response values
}

// ClustersClientUpdateAutoScaleConfigurationResponse contains the response from method ClustersClient.UpdateAutoScaleConfiguration.
type ClustersClientUpdateAutoScaleConfigurationResponse struct {
	// placeholder for future response values
}

// ClustersClientUpdateGatewaySettingsResponse contains the response from method ClustersClient.UpdateGatewaySettings.
type ClustersClientUpdateGatewaySettingsResponse struct {
	// placeholder for future response values
}

// ClustersClientUpdateIdentityCertificateResponse contains the response from method ClustersClient.UpdateIdentityCertificate.
type ClustersClientUpdateIdentityCertificateResponse struct {
	// placeholder for future response values
}

// ClustersClientUpdateResponse contains the response from method ClustersClient.Update.
type ClustersClientUpdateResponse struct {
	Cluster
}

// ConfigurationsClientGetResponse contains the response from method ConfigurationsClient.Get.
type ConfigurationsClientGetResponse struct {
	// The configuration object for the specified configuration for the specified cluster.
	Value map[string]*string
}

// ConfigurationsClientListResponse contains the response from method ConfigurationsClient.List.
type ConfigurationsClientListResponse struct {
	ClusterConfigurations
}

// ConfigurationsClientUpdateResponse contains the response from method ConfigurationsClient.Update.
type ConfigurationsClientUpdateResponse struct {
	// placeholder for future response values
}

// ExtensionsClientCreateResponse contains the response from method ExtensionsClient.Create.
type ExtensionsClientCreateResponse struct {
	// placeholder for future response values
}

// ExtensionsClientDeleteResponse contains the response from method ExtensionsClient.Delete.
type ExtensionsClientDeleteResponse struct {
	// placeholder for future response values
}

// ExtensionsClientDisableAzureMonitorResponse contains the response from method ExtensionsClient.DisableAzureMonitor.
type ExtensionsClientDisableAzureMonitorResponse struct {
	// placeholder for future response values
}

// ExtensionsClientDisableMonitoringResponse contains the response from method ExtensionsClient.DisableMonitoring.
type ExtensionsClientDisableMonitoringResponse struct {
	// placeholder for future response values
}

// ExtensionsClientEnableAzureMonitorResponse contains the response from method ExtensionsClient.EnableAzureMonitor.
type ExtensionsClientEnableAzureMonitorResponse struct {
	// placeholder for future response values
}

// ExtensionsClientEnableMonitoringResponse contains the response from method ExtensionsClient.EnableMonitoring.
type ExtensionsClientEnableMonitoringResponse struct {
	// placeholder for future response values
}

// ExtensionsClientGetAzureAsyncOperationStatusResponse contains the response from method ExtensionsClient.GetAzureAsyncOperationStatus.
type ExtensionsClientGetAzureAsyncOperationStatusResponse struct {
	AsyncOperationResult
}

// ExtensionsClientGetAzureMonitorStatusResponse contains the response from method ExtensionsClient.GetAzureMonitorStatus.
type ExtensionsClientGetAzureMonitorStatusResponse struct {
	AzureMonitorResponse
}

// ExtensionsClientGetMonitoringStatusResponse contains the response from method ExtensionsClient.GetMonitoringStatus.
type ExtensionsClientGetMonitoringStatusResponse struct {
	ClusterMonitoringResponse
}

// ExtensionsClientGetResponse contains the response from method ExtensionsClient.Get.
type ExtensionsClientGetResponse struct {
	ClusterMonitoringResponse
}

// LocationsClientCheckNameAvailabilityResponse contains the response from method LocationsClient.CheckNameAvailability.
type LocationsClientCheckNameAvailabilityResponse struct {
	NameAvailabilityCheckResult
}

// LocationsClientGetAzureAsyncOperationStatusResponse contains the response from method LocationsClient.GetAzureAsyncOperationStatus.
type LocationsClientGetAzureAsyncOperationStatusResponse struct {
	AsyncOperationResult
}

// LocationsClientGetCapabilitiesResponse contains the response from method LocationsClient.GetCapabilities.
type LocationsClientGetCapabilitiesResponse struct {
	CapabilitiesResult
}

// LocationsClientListBillingSpecsResponse contains the response from method LocationsClient.ListBillingSpecs.
type LocationsClientListBillingSpecsResponse struct {
	BillingResponseListResult
}

// LocationsClientListUsagesResponse contains the response from method LocationsClient.ListUsages.
type LocationsClientListUsagesResponse struct {
	UsagesListResult
}

// LocationsClientValidateClusterCreateRequestResponse contains the response from method LocationsClient.ValidateClusterCreateRequest.
type LocationsClientValidateClusterCreateRequestResponse struct {
	ClusterCreateValidationResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByClusterResponse contains the response from method PrivateEndpointConnectionsClient.ListByCluster.
type PrivateEndpointConnectionsClientListByClusterResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByClusterResponse contains the response from method PrivateLinkResourcesClient.ListByCluster.
type PrivateLinkResourcesClientListByClusterResponse struct {
	PrivateLinkResourceListResult
}

// ScriptActionsClientDeleteResponse contains the response from method ScriptActionsClient.Delete.
type ScriptActionsClientDeleteResponse struct {
	// placeholder for future response values
}

// ScriptActionsClientGetExecutionAsyncOperationStatusResponse contains the response from method ScriptActionsClient.GetExecutionAsyncOperationStatus.
type ScriptActionsClientGetExecutionAsyncOperationStatusResponse struct {
	AsyncOperationResult
}

// ScriptActionsClientGetExecutionDetailResponse contains the response from method ScriptActionsClient.GetExecutionDetail.
type ScriptActionsClientGetExecutionDetailResponse struct {
	RuntimeScriptActionDetail
}

// ScriptActionsClientListByClusterResponse contains the response from method ScriptActionsClient.ListByCluster.
type ScriptActionsClientListByClusterResponse struct {
	ScriptActionsList
}

// ScriptExecutionHistoryClientListByClusterResponse contains the response from method ScriptExecutionHistoryClient.ListByCluster.
type ScriptExecutionHistoryClientListByClusterResponse struct {
	ScriptActionExecutionHistoryList
}

// ScriptExecutionHistoryClientPromoteResponse contains the response from method ScriptExecutionHistoryClient.Promote.
type ScriptExecutionHistoryClientPromoteResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientGetAsyncOperationStatusResponse contains the response from method VirtualMachinesClient.GetAsyncOperationStatus.
type VirtualMachinesClientGetAsyncOperationStatusResponse struct {
	AsyncOperationResult
}

// VirtualMachinesClientListHostsResponse contains the response from method VirtualMachinesClient.ListHosts.
type VirtualMachinesClientListHostsResponse struct {
	// Result of the request to list cluster hosts
	HostInfoArray []*HostInfo
}

// VirtualMachinesClientRestartHostsResponse contains the response from method VirtualMachinesClient.RestartHosts.
type VirtualMachinesClientRestartHostsResponse struct {
	// placeholder for future response values
}
