//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// APIIssueAttachmentClient contains the methods for the APIIssueAttachment group.
// Don't use this type directly, use NewAPIIssueAttachmentClient() instead.
type APIIssueAttachmentClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAPIIssueAttachmentClient creates a new instance of APIIssueAttachmentClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAPIIssueAttachmentClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*APIIssueAttachmentClient, error) {
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
	client := &APIIssueAttachmentClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a new Attachment for the Issue in an API or updates an existing one.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API identifier. Must be unique in the current API Management service instance.
// issueID - Issue identifier. Must be unique in the current API Management service instance.
// attachmentID - Attachment identifier within an Issue. Must be unique in the current Issue.
// parameters - Create parameters.
// options - APIIssueAttachmentClientCreateOrUpdateOptions contains the optional parameters for the APIIssueAttachmentClient.CreateOrUpdate
// method.
func (client *APIIssueAttachmentClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, attachmentID string, parameters IssueAttachmentContract, options *APIIssueAttachmentClientCreateOrUpdateOptions) (APIIssueAttachmentClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, apiID, issueID, attachmentID, parameters, options)
	if err != nil {
		return APIIssueAttachmentClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIIssueAttachmentClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return APIIssueAttachmentClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *APIIssueAttachmentClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, attachmentID string, parameters IssueAttachmentContract, options *APIIssueAttachmentClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/attachments/{attachmentId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if issueID == "" {
		return nil, errors.New("parameter issueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueId}", url.PathEscape(issueID))
	if attachmentID == "" {
		return nil, errors.New("parameter attachmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachmentId}", url.PathEscape(attachmentID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *APIIssueAttachmentClient) createOrUpdateHandleResponse(resp *http.Response) (APIIssueAttachmentClientCreateOrUpdateResponse, error) {
	result := APIIssueAttachmentClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.IssueAttachmentContract); err != nil {
		return APIIssueAttachmentClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified comment from an Issue.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API identifier. Must be unique in the current API Management service instance.
// issueID - Issue identifier. Must be unique in the current API Management service instance.
// attachmentID - Attachment identifier within an Issue. Must be unique in the current Issue.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// options - APIIssueAttachmentClientDeleteOptions contains the optional parameters for the APIIssueAttachmentClient.Delete
// method.
func (client *APIIssueAttachmentClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, attachmentID string, ifMatch string, options *APIIssueAttachmentClientDeleteOptions) (APIIssueAttachmentClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, apiID, issueID, attachmentID, ifMatch, options)
	if err != nil {
		return APIIssueAttachmentClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIIssueAttachmentClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return APIIssueAttachmentClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return APIIssueAttachmentClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *APIIssueAttachmentClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, attachmentID string, ifMatch string, options *APIIssueAttachmentClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/attachments/{attachmentId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if issueID == "" {
		return nil, errors.New("parameter issueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueId}", url.PathEscape(issueID))
	if attachmentID == "" {
		return nil, errors.New("parameter attachmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachmentId}", url.PathEscape(attachmentID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the details of the issue Attachment for an API specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API identifier. Must be unique in the current API Management service instance.
// issueID - Issue identifier. Must be unique in the current API Management service instance.
// attachmentID - Attachment identifier within an Issue. Must be unique in the current Issue.
// options - APIIssueAttachmentClientGetOptions contains the optional parameters for the APIIssueAttachmentClient.Get method.
func (client *APIIssueAttachmentClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, attachmentID string, options *APIIssueAttachmentClientGetOptions) (APIIssueAttachmentClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, apiID, issueID, attachmentID, options)
	if err != nil {
		return APIIssueAttachmentClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIIssueAttachmentClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIIssueAttachmentClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *APIIssueAttachmentClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, attachmentID string, options *APIIssueAttachmentClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/attachments/{attachmentId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if issueID == "" {
		return nil, errors.New("parameter issueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueId}", url.PathEscape(issueID))
	if attachmentID == "" {
		return nil, errors.New("parameter attachmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachmentId}", url.PathEscape(attachmentID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *APIIssueAttachmentClient) getHandleResponse(resp *http.Response) (APIIssueAttachmentClientGetResponse, error) {
	result := APIIssueAttachmentClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.IssueAttachmentContract); err != nil {
		return APIIssueAttachmentClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the issue Attachment for an API specified by its identifier.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API identifier. Must be unique in the current API Management service instance.
// issueID - Issue identifier. Must be unique in the current API Management service instance.
// attachmentID - Attachment identifier within an Issue. Must be unique in the current Issue.
// options - APIIssueAttachmentClientGetEntityTagOptions contains the optional parameters for the APIIssueAttachmentClient.GetEntityTag
// method.
func (client *APIIssueAttachmentClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, attachmentID string, options *APIIssueAttachmentClientGetEntityTagOptions) (APIIssueAttachmentClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, apiID, issueID, attachmentID, options)
	if err != nil {
		return APIIssueAttachmentClientGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIIssueAttachmentClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIIssueAttachmentClientGetEntityTagResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *APIIssueAttachmentClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, attachmentID string, options *APIIssueAttachmentClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/attachments/{attachmentId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if issueID == "" {
		return nil, errors.New("parameter issueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueId}", url.PathEscape(issueID))
	if attachmentID == "" {
		return nil, errors.New("parameter attachmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachmentId}", url.PathEscape(attachmentID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *APIIssueAttachmentClient) getEntityTagHandleResponse(resp *http.Response) (APIIssueAttachmentClientGetEntityTagResponse, error) {
	result := APIIssueAttachmentClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
	return result, nil
}

// NewListByServicePager - Lists all attachments for the Issue associated with the specified API.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API identifier. Must be unique in the current API Management service instance.
// issueID - Issue identifier. Must be unique in the current API Management service instance.
// options - APIIssueAttachmentClientListByServiceOptions contains the optional parameters for the APIIssueAttachmentClient.ListByService
// method.
func (client *APIIssueAttachmentClient) NewListByServicePager(resourceGroupName string, serviceName string, apiID string, issueID string, options *APIIssueAttachmentClientListByServiceOptions) *runtime.Pager[APIIssueAttachmentClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[APIIssueAttachmentClientListByServiceResponse]{
		More: func(page APIIssueAttachmentClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *APIIssueAttachmentClientListByServiceResponse) (APIIssueAttachmentClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, apiID, issueID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return APIIssueAttachmentClientListByServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return APIIssueAttachmentClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return APIIssueAttachmentClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *APIIssueAttachmentClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, options *APIIssueAttachmentClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/attachments"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if issueID == "" {
		return nil, errors.New("parameter issueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueId}", url.PathEscape(issueID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *APIIssueAttachmentClient) listByServiceHandleResponse(resp *http.Response) (APIIssueAttachmentClientListByServiceResponse, error) {
	result := APIIssueAttachmentClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IssueAttachmentCollection); err != nil {
		return APIIssueAttachmentClientListByServiceResponse{}, err
	}
	return result, nil
}
