//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DisksClient contains the methods for the Disks group.
// Don't use this type directly, use NewDisksClient() instead.
type DisksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDisksClient creates a new instance of DisksClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDisksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DisksClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DisksClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginAttach - Attach and create the lease of the disk to the virtual machine. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-15
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// userName - The name of the user profile.
// name - The name of the disk.
// attachDiskProperties - Properties of the disk to attach.
// options - DisksClientBeginAttachOptions contains the optional parameters for the DisksClient.BeginAttach method.
func (client *DisksClient) BeginAttach(ctx context.Context, resourceGroupName string, labName string, userName string, name string, attachDiskProperties AttachDiskProperties, options *DisksClientBeginAttachOptions) (*runtime.Poller[DisksClientAttachResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.attach(ctx, resourceGroupName, labName, userName, name, attachDiskProperties, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DisksClientAttachResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DisksClientAttachResponse](options.ResumeToken, client.pl, nil)
	}
}

// Attach - Attach and create the lease of the disk to the virtual machine. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-15
func (client *DisksClient) attach(ctx context.Context, resourceGroupName string, labName string, userName string, name string, attachDiskProperties AttachDiskProperties, options *DisksClientBeginAttachOptions) (*http.Response, error) {
	req, err := client.attachCreateRequest(ctx, resourceGroupName, labName, userName, name, attachDiskProperties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// attachCreateRequest creates the Attach request.
func (client *DisksClient) attachCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, attachDiskProperties AttachDiskProperties, options *DisksClientBeginAttachOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}/attach"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, attachDiskProperties)
}

// BeginCreateOrUpdate - Create or replace an existing disk. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-15
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// userName - The name of the user profile.
// name - The name of the disk.
// disk - A Disk.
// options - DisksClientBeginCreateOrUpdateOptions contains the optional parameters for the DisksClient.BeginCreateOrUpdate
// method.
func (client *DisksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk Disk, options *DisksClientBeginCreateOrUpdateOptions) (*runtime.Poller[DisksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, labName, userName, name, disk, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DisksClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DisksClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or replace an existing disk. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-15
func (client *DisksClient) createOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk Disk, options *DisksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, userName, name, disk, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DisksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk Disk, options *DisksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, disk)
}

// BeginDelete - Delete disk. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-15
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// userName - The name of the user profile.
// name - The name of the disk.
// options - DisksClientBeginDeleteOptions contains the optional parameters for the DisksClient.BeginDelete method.
func (client *DisksClient) BeginDelete(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *DisksClientBeginDeleteOptions) (*runtime.Poller[DisksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, labName, userName, name, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DisksClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DisksClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete disk. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-15
func (client *DisksClient) deleteOperation(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *DisksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, userName, name, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DisksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *DisksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDetach - Detach and break the lease of the disk attached to the virtual machine. This operation can take a while to
// complete.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-15
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// userName - The name of the user profile.
// name - The name of the disk.
// detachDiskProperties - Properties of the disk to detach.
// options - DisksClientBeginDetachOptions contains the optional parameters for the DisksClient.BeginDetach method.
func (client *DisksClient) BeginDetach(ctx context.Context, resourceGroupName string, labName string, userName string, name string, detachDiskProperties DetachDiskProperties, options *DisksClientBeginDetachOptions) (*runtime.Poller[DisksClientDetachResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.detach(ctx, resourceGroupName, labName, userName, name, detachDiskProperties, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DisksClientDetachResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DisksClientDetachResponse](options.ResumeToken, client.pl, nil)
	}
}

// Detach - Detach and break the lease of the disk attached to the virtual machine. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-15
func (client *DisksClient) detach(ctx context.Context, resourceGroupName string, labName string, userName string, name string, detachDiskProperties DetachDiskProperties, options *DisksClientBeginDetachOptions) (*http.Response, error) {
	req, err := client.detachCreateRequest(ctx, resourceGroupName, labName, userName, name, detachDiskProperties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// detachCreateRequest creates the Detach request.
func (client *DisksClient) detachCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, detachDiskProperties DetachDiskProperties, options *DisksClientBeginDetachOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}/detach"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, detachDiskProperties)
}

// Get - Get disk.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-15
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// userName - The name of the user profile.
// name - The name of the disk.
// options - DisksClientGetOptions contains the optional parameters for the DisksClient.Get method.
func (client *DisksClient) Get(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *DisksClientGetOptions) (DisksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, userName, name, options)
	if err != nil {
		return DisksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DisksClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *DisksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DisksClient) getHandleResponse(resp *http.Response) (DisksClientGetResponse, error) {
	result := DisksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Disk); err != nil {
		return DisksClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List disks in a given user profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-15
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// userName - The name of the user profile.
// options - DisksClientListOptions contains the optional parameters for the DisksClient.List method.
func (client *DisksClient) NewListPager(resourceGroupName string, labName string, userName string, options *DisksClientListOptions) *runtime.Pager[DisksClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DisksClientListResponse]{
		More: func(page DisksClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DisksClientListResponse) (DisksClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, labName, userName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DisksClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DisksClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DisksClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DisksClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, options *DisksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DisksClient) listHandleResponse(resp *http.Response) (DisksClientListResponse, error) {
	result := DisksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskList); err != nil {
		return DisksClientListResponse{}, err
	}
	return result, nil
}

// Update - Allows modifying tags of disks. All other properties will be ignored.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-15
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// userName - The name of the user profile.
// name - The name of the disk.
// disk - A Disk.
// options - DisksClientUpdateOptions contains the optional parameters for the DisksClient.Update method.
func (client *DisksClient) Update(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk DiskFragment, options *DisksClientUpdateOptions) (DisksClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labName, userName, name, disk, options)
	if err != nil {
		return DisksClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisksClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisksClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *DisksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk DiskFragment, options *DisksClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, disk)
}

// updateHandleResponse handles the Update response.
func (client *DisksClient) updateHandleResponse(resp *http.Response) (DisksClientUpdateResponse, error) {
	result := DisksClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Disk); err != nil {
		return DisksClientUpdateResponse{}, err
	}
	return result, nil
}
