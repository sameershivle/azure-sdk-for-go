//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdnsresolver

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

// ForwardingRulesClient contains the methods for the ForwardingRules group.
// Don't use this type directly, use NewForwardingRulesClient() instead.
type ForwardingRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewForwardingRulesClient creates a new instance of ForwardingRulesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewForwardingRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ForwardingRulesClient, error) {
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
	client := &ForwardingRulesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a forwarding rule in a DNS forwarding ruleset.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-04-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsForwardingRulesetName - The name of the DNS forwarding ruleset.
// forwardingRuleName - The name of the forwarding rule.
// parameters - Parameters supplied to the CreateOrUpdate operation.
// options - ForwardingRulesClientCreateOrUpdateOptions contains the optional parameters for the ForwardingRulesClient.CreateOrUpdate
// method.
func (client *ForwardingRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, forwardingRuleName string, parameters ForwardingRule, options *ForwardingRulesClientCreateOrUpdateOptions) (ForwardingRulesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, dnsForwardingRulesetName, forwardingRuleName, parameters, options)
	if err != nil {
		return ForwardingRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ForwardingRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ForwardingRulesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ForwardingRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, forwardingRuleName string, parameters ForwardingRule, options *ForwardingRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}/forwardingRules/{forwardingRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsForwardingRulesetName == "" {
		return nil, errors.New("parameter dnsForwardingRulesetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsForwardingRulesetName}", url.PathEscape(dnsForwardingRulesetName))
	if forwardingRuleName == "" {
		return nil, errors.New("parameter forwardingRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{forwardingRuleName}", url.PathEscape(forwardingRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ForwardingRulesClient) createOrUpdateHandleResponse(resp *http.Response) (ForwardingRulesClientCreateOrUpdateResponse, error) {
	result := ForwardingRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ForwardingRule); err != nil {
		return ForwardingRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a forwarding rule in a DNS forwarding ruleset. WARNING: This operation cannot be undone.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-04-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsForwardingRulesetName - The name of the DNS forwarding ruleset.
// forwardingRuleName - The name of the forwarding rule.
// options - ForwardingRulesClientDeleteOptions contains the optional parameters for the ForwardingRulesClient.Delete method.
func (client *ForwardingRulesClient) Delete(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, forwardingRuleName string, options *ForwardingRulesClientDeleteOptions) (ForwardingRulesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dnsForwardingRulesetName, forwardingRuleName, options)
	if err != nil {
		return ForwardingRulesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ForwardingRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ForwardingRulesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ForwardingRulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ForwardingRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, forwardingRuleName string, options *ForwardingRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}/forwardingRules/{forwardingRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsForwardingRulesetName == "" {
		return nil, errors.New("parameter dnsForwardingRulesetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsForwardingRulesetName}", url.PathEscape(dnsForwardingRulesetName))
	if forwardingRuleName == "" {
		return nil, errors.New("parameter forwardingRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{forwardingRuleName}", url.PathEscape(forwardingRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets properties of a forwarding rule in a DNS forwarding ruleset.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-04-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsForwardingRulesetName - The name of the DNS forwarding ruleset.
// forwardingRuleName - The name of the forwarding rule.
// options - ForwardingRulesClientGetOptions contains the optional parameters for the ForwardingRulesClient.Get method.
func (client *ForwardingRulesClient) Get(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, forwardingRuleName string, options *ForwardingRulesClientGetOptions) (ForwardingRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, dnsForwardingRulesetName, forwardingRuleName, options)
	if err != nil {
		return ForwardingRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ForwardingRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ForwardingRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ForwardingRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, forwardingRuleName string, options *ForwardingRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}/forwardingRules/{forwardingRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsForwardingRulesetName == "" {
		return nil, errors.New("parameter dnsForwardingRulesetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsForwardingRulesetName}", url.PathEscape(dnsForwardingRulesetName))
	if forwardingRuleName == "" {
		return nil, errors.New("parameter forwardingRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{forwardingRuleName}", url.PathEscape(forwardingRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ForwardingRulesClient) getHandleResponse(resp *http.Response) (ForwardingRulesClientGetResponse, error) {
	result := ForwardingRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ForwardingRule); err != nil {
		return ForwardingRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists forwarding rules in a DNS forwarding ruleset.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-04-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsForwardingRulesetName - The name of the DNS forwarding ruleset.
// options - ForwardingRulesClientListOptions contains the optional parameters for the ForwardingRulesClient.List method.
func (client *ForwardingRulesClient) NewListPager(resourceGroupName string, dnsForwardingRulesetName string, options *ForwardingRulesClientListOptions) *runtime.Pager[ForwardingRulesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ForwardingRulesClientListResponse]{
		More: func(page ForwardingRulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ForwardingRulesClientListResponse) (ForwardingRulesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, dnsForwardingRulesetName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ForwardingRulesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ForwardingRulesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ForwardingRulesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ForwardingRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, options *ForwardingRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}/forwardingRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsForwardingRulesetName == "" {
		return nil, errors.New("parameter dnsForwardingRulesetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsForwardingRulesetName}", url.PathEscape(dnsForwardingRulesetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ForwardingRulesClient) listHandleResponse(resp *http.Response) (ForwardingRulesClientListResponse, error) {
	result := ForwardingRulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ForwardingRuleListResult); err != nil {
		return ForwardingRulesClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates a forwarding rule in a DNS forwarding ruleset.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-04-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsForwardingRulesetName - The name of the DNS forwarding ruleset.
// forwardingRuleName - The name of the forwarding rule.
// parameters - Parameters supplied to the Update operation.
// options - ForwardingRulesClientUpdateOptions contains the optional parameters for the ForwardingRulesClient.Update method.
func (client *ForwardingRulesClient) Update(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, forwardingRuleName string, parameters ForwardingRulePatch, options *ForwardingRulesClientUpdateOptions) (ForwardingRulesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, dnsForwardingRulesetName, forwardingRuleName, parameters, options)
	if err != nil {
		return ForwardingRulesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ForwardingRulesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ForwardingRulesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ForwardingRulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, dnsForwardingRulesetName string, forwardingRuleName string, parameters ForwardingRulePatch, options *ForwardingRulesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}/forwardingRules/{forwardingRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsForwardingRulesetName == "" {
		return nil, errors.New("parameter dnsForwardingRulesetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsForwardingRulesetName}", url.PathEscape(dnsForwardingRulesetName))
	if forwardingRuleName == "" {
		return nil, errors.New("parameter forwardingRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{forwardingRuleName}", url.PathEscape(forwardingRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *ForwardingRulesClient) updateHandleResponse(resp *http.Response) (ForwardingRulesClientUpdateResponse, error) {
	result := ForwardingRulesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ForwardingRule); err != nil {
		return ForwardingRulesClientUpdateResponse{}, err
	}
	return result, nil
}
