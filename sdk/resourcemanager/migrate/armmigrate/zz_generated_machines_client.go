//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

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

// MachinesClient contains the methods for the Machines group.
// Don't use this type directly, use NewMachinesClient() instead.
type MachinesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMachinesClient creates a new instance of MachinesClient with the specified values.
// subscriptionID - Azure Subscription Id in which project was created.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMachinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MachinesClient, error) {
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
	client := &MachinesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get the machine with the specified name. Returns a json object of type 'machine' defined in Models section.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// resourceGroupName - Name of the Azure Resource Group that project is part of.
// projectName - Name of the Azure Migrate project.
// machineName - Unique name of a machine in private datacenter.
// options - MachinesClientGetOptions contains the optional parameters for the MachinesClient.Get method.
func (client *MachinesClient) Get(ctx context.Context, resourceGroupName string, projectName string, machineName string, options *MachinesClientGetOptions) (MachinesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, machineName, options)
	if err != nil {
		return MachinesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MachinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MachinesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MachinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, machineName string, options *MachinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/machines/{machineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MachinesClient) getHandleResponse(resp *http.Response) (MachinesClientGetResponse, error) {
	result := MachinesClientGetResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Machine); err != nil {
		return MachinesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByProjectPager - Get data of all the machines available in the project. Returns a json array of objects of type
// 'machine' defined in Models section.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// resourceGroupName - Name of the Azure Resource Group that project is part of.
// projectName - Name of the Azure Migrate project.
// options - MachinesClientListByProjectOptions contains the optional parameters for the MachinesClient.ListByProject method.
func (client *MachinesClient) NewListByProjectPager(resourceGroupName string, projectName string, options *MachinesClientListByProjectOptions) *runtime.Pager[MachinesClientListByProjectResponse] {
	return runtime.NewPager(runtime.PagingHandler[MachinesClientListByProjectResponse]{
		More: func(page MachinesClientListByProjectResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MachinesClientListByProjectResponse) (MachinesClientListByProjectResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByProjectCreateRequest(ctx, resourceGroupName, projectName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MachinesClientListByProjectResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return MachinesClientListByProjectResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MachinesClientListByProjectResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByProjectHandleResponse(resp)
		},
	})
}

// listByProjectCreateRequest creates the ListByProject request.
func (client *MachinesClient) listByProjectCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *MachinesClientListByProjectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/machines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProjectHandleResponse handles the ListByProject response.
func (client *MachinesClient) listByProjectHandleResponse(resp *http.Response) (MachinesClientListByProjectResponse, error) {
	result := MachinesClientListByProjectResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.MachineResultList); err != nil {
		return MachinesClientListByProjectResponse{}, err
	}
	return result, nil
}
