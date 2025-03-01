//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armworkloads

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
	"strings"
)

// SAPCentralInstancesClient contains the methods for the SAPCentralInstances group.
// Don't use this type directly, use NewSAPCentralInstancesClient() instead.
type SAPCentralInstancesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSAPCentralInstancesClient creates a new instance of SAPCentralInstancesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSAPCentralInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SAPCentralInstancesClient, error) {
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
	client := &SAPCentralInstancesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Puts the SAP Central Instance.
// This will be used by service only. PUT by end user will return a Bad Request error.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// sapVirtualInstanceName - The name of the Virtual Instances for SAP.
// centralInstanceName - Central Instance name string modeled as parameter for auto generation to work correctly.
// body - The SAP Central Server instance request body.
// options - SAPCentralInstancesClientBeginCreateOptions contains the optional parameters for the SAPCentralInstancesClient.BeginCreate
// method.
func (client *SAPCentralInstancesClient) BeginCreate(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, centralInstanceName string, body SAPCentralServerInstance, options *SAPCentralInstancesClientBeginCreateOptions) (*runtime.Poller[SAPCentralInstancesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, sapVirtualInstanceName, centralInstanceName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SAPCentralInstancesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SAPCentralInstancesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Puts the SAP Central Instance.
// This will be used by service only. PUT by end user will return a Bad Request error.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
func (client *SAPCentralInstancesClient) create(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, centralInstanceName string, body SAPCentralServerInstance, options *SAPCentralInstancesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, centralInstanceName, body, options)
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

// createCreateRequest creates the Create request.
func (client *SAPCentralInstancesClient) createCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, centralInstanceName string, body SAPCentralServerInstance, options *SAPCentralInstancesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/centralInstances/{centralInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if centralInstanceName == "" {
		return nil, errors.New("parameter centralInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{centralInstanceName}", url.PathEscape(centralInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Deletes the SAP Central Instance.
// This will be used by service only. Delete by end user will return a Bad Request error.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// sapVirtualInstanceName - The name of the Virtual Instances for SAP.
// centralInstanceName - Central Instance name string modeled as parameter for auto generation to work correctly.
// options - SAPCentralInstancesClientBeginDeleteOptions contains the optional parameters for the SAPCentralInstancesClient.BeginDelete
// method.
func (client *SAPCentralInstancesClient) BeginDelete(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, centralInstanceName string, options *SAPCentralInstancesClientBeginDeleteOptions) (*runtime.Poller[SAPCentralInstancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, sapVirtualInstanceName, centralInstanceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SAPCentralInstancesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SAPCentralInstancesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the SAP Central Instance.
// This will be used by service only. Delete by end user will return a Bad Request error.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
func (client *SAPCentralInstancesClient) deleteOperation(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, centralInstanceName string, options *SAPCentralInstancesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, centralInstanceName, options)
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
func (client *SAPCentralInstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, centralInstanceName string, options *SAPCentralInstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/centralInstances/{centralInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if centralInstanceName == "" {
		return nil, errors.New("parameter centralInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{centralInstanceName}", url.PathEscape(centralInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the SAP Central Instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// sapVirtualInstanceName - The name of the Virtual Instances for SAP.
// centralInstanceName - Central Instance name string modeled as parameter for auto generation to work correctly.
// options - SAPCentralInstancesClientGetOptions contains the optional parameters for the SAPCentralInstancesClient.Get method.
func (client *SAPCentralInstancesClient) Get(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, centralInstanceName string, options *SAPCentralInstancesClientGetOptions) (SAPCentralInstancesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, centralInstanceName, options)
	if err != nil {
		return SAPCentralInstancesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SAPCentralInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SAPCentralInstancesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SAPCentralInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, centralInstanceName string, options *SAPCentralInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/centralInstances/{centralInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if centralInstanceName == "" {
		return nil, errors.New("parameter centralInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{centralInstanceName}", url.PathEscape(centralInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SAPCentralInstancesClient) getHandleResponse(resp *http.Response) (SAPCentralInstancesClientGetResponse, error) {
	result := SAPCentralInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPCentralServerInstance); err != nil {
		return SAPCentralInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the SAP Central Instances in an SVI.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// sapVirtualInstanceName - The name of the Virtual Instances for SAP.
// options - SAPCentralInstancesClientListOptions contains the optional parameters for the SAPCentralInstancesClient.List
// method.
func (client *SAPCentralInstancesClient) NewListPager(resourceGroupName string, sapVirtualInstanceName string, options *SAPCentralInstancesClientListOptions) *runtime.Pager[SAPCentralInstancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SAPCentralInstancesClientListResponse]{
		More: func(page SAPCentralInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SAPCentralInstancesClientListResponse) (SAPCentralInstancesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SAPCentralInstancesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SAPCentralInstancesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SAPCentralInstancesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SAPCentralInstancesClient) listCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, options *SAPCentralInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/centralInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SAPCentralInstancesClient) listHandleResponse(resp *http.Response) (SAPCentralInstancesClientListResponse, error) {
	result := SAPCentralInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPCentralInstanceList); err != nil {
		return SAPCentralInstancesClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the SAP Central Instance.
// This can be used to update tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// sapVirtualInstanceName - The name of the Virtual Instances for SAP.
// centralInstanceName - Central Instance name string modeled as parameter for auto generation to work correctly.
// body - The SAP Central Server instance request body.
// options - SAPCentralInstancesClientBeginUpdateOptions contains the optional parameters for the SAPCentralInstancesClient.BeginUpdate
// method.
func (client *SAPCentralInstancesClient) BeginUpdate(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, centralInstanceName string, body UpdateSAPCentralInstanceRequest, options *SAPCentralInstancesClientBeginUpdateOptions) (*runtime.Poller[SAPCentralInstancesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, sapVirtualInstanceName, centralInstanceName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SAPCentralInstancesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SAPCentralInstancesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates the SAP Central Instance.
// This can be used to update tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
func (client *SAPCentralInstancesClient) update(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, centralInstanceName string, body UpdateSAPCentralInstanceRequest, options *SAPCentralInstancesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, centralInstanceName, body, options)
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

// updateCreateRequest creates the Update request.
func (client *SAPCentralInstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, centralInstanceName string, body UpdateSAPCentralInstanceRequest, options *SAPCentralInstancesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/centralInstances/{centralInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if centralInstanceName == "" {
		return nil, errors.New("parameter centralInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{centralInstanceName}", url.PathEscape(centralInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
