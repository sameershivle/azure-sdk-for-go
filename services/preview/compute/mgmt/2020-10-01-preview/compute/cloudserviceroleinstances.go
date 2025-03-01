package compute

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// CloudServiceRoleInstancesClient is the compute Client
type CloudServiceRoleInstancesClient struct {
	BaseClient
}

// NewCloudServiceRoleInstancesClient creates an instance of the CloudServiceRoleInstancesClient client.
func NewCloudServiceRoleInstancesClient(subscriptionID string) CloudServiceRoleInstancesClient {
	return NewCloudServiceRoleInstancesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCloudServiceRoleInstancesClientWithBaseURI creates an instance of the CloudServiceRoleInstancesClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewCloudServiceRoleInstancesClientWithBaseURI(baseURI string, subscriptionID string) CloudServiceRoleInstancesClient {
	return CloudServiceRoleInstancesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete deletes a role instance from a cloud service.
// Parameters:
// roleInstanceName - name of the role instance.
func (client CloudServiceRoleInstancesClient) Delete(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string) (result CloudServiceRoleInstancesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudServiceRoleInstancesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, roleInstanceName, resourceGroupName, cloudServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client CloudServiceRoleInstancesClient) DeletePreparer(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudServiceName":  autorest.Encode("path", cloudServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"roleInstanceName":  autorest.Encode("path", roleInstanceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CloudServiceRoleInstancesClient) DeleteSender(req *http.Request) (future CloudServiceRoleInstancesDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client CloudServiceRoleInstancesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a role instance from a cloud service.
// Parameters:
// roleInstanceName - name of the role instance.
// expand - the expand expression to apply to the operation.
func (client CloudServiceRoleInstancesClient) Get(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, expand InstanceViewTypes) (result RoleInstance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudServiceRoleInstancesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, roleInstanceName, resourceGroupName, cloudServiceName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client CloudServiceRoleInstancesClient) GetPreparer(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string, expand InstanceViewTypes) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudServiceName":  autorest.Encode("path", cloudServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"roleInstanceName":  autorest.Encode("path", roleInstanceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CloudServiceRoleInstancesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CloudServiceRoleInstancesClient) GetResponder(resp *http.Response) (result RoleInstance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetInstanceView retrieves information about the run-time state of a role instance in a cloud service.
// Parameters:
// roleInstanceName - name of the role instance.
func (client CloudServiceRoleInstancesClient) GetInstanceView(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string) (result RoleInstanceInstanceView, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudServiceRoleInstancesClient.GetInstanceView")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetInstanceViewPreparer(ctx, roleInstanceName, resourceGroupName, cloudServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "GetInstanceView", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetInstanceViewSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "GetInstanceView", resp, "Failure sending request")
		return
	}

	result, err = client.GetInstanceViewResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "GetInstanceView", resp, "Failure responding to request")
		return
	}

	return
}

// GetInstanceViewPreparer prepares the GetInstanceView request.
func (client CloudServiceRoleInstancesClient) GetInstanceViewPreparer(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudServiceName":  autorest.Encode("path", cloudServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"roleInstanceName":  autorest.Encode("path", roleInstanceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}/instanceView", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetInstanceViewSender sends the GetInstanceView request. The method will close the
// http.Response Body if it receives an error.
func (client CloudServiceRoleInstancesClient) GetInstanceViewSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetInstanceViewResponder handles the response to the GetInstanceView request. The method always
// closes the http.Response Body.
func (client CloudServiceRoleInstancesClient) GetInstanceViewResponder(resp *http.Response) (result RoleInstanceInstanceView, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetRemoteDesktopFile gets a remote desktop file for a role instance in a cloud service.
// Parameters:
// roleInstanceName - name of the role instance.
func (client CloudServiceRoleInstancesClient) GetRemoteDesktopFile(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string) (result ReadCloser, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudServiceRoleInstancesClient.GetRemoteDesktopFile")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetRemoteDesktopFilePreparer(ctx, roleInstanceName, resourceGroupName, cloudServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "GetRemoteDesktopFile", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRemoteDesktopFileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "GetRemoteDesktopFile", resp, "Failure sending request")
		return
	}

	result, err = client.GetRemoteDesktopFileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "GetRemoteDesktopFile", resp, "Failure responding to request")
		return
	}

	return
}

// GetRemoteDesktopFilePreparer prepares the GetRemoteDesktopFile request.
func (client CloudServiceRoleInstancesClient) GetRemoteDesktopFilePreparer(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudServiceName":  autorest.Encode("path", cloudServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"roleInstanceName":  autorest.Encode("path", roleInstanceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}/remoteDesktopFile", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetRemoteDesktopFileSender sends the GetRemoteDesktopFile request. The method will close the
// http.Response Body if it receives an error.
func (client CloudServiceRoleInstancesClient) GetRemoteDesktopFileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetRemoteDesktopFileResponder handles the response to the GetRemoteDesktopFile request. The method always
// closes the http.Response Body.
func (client CloudServiceRoleInstancesClient) GetRemoteDesktopFileResponder(resp *http.Response) (result ReadCloser, err error) {
	result.Value = &resp.Body
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the list of all role instances in a cloud service. Use nextLink property in the response to get the next
// page of role instances. Do this till nextLink is null to fetch all the role instances.
// Parameters:
// expand - the expand expression to apply to the operation.
func (client CloudServiceRoleInstancesClient) List(ctx context.Context, resourceGroupName string, cloudServiceName string, expand InstanceViewTypes) (result RoleInstanceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudServiceRoleInstancesClient.List")
		defer func() {
			sc := -1
			if result.rilr.Response.Response != nil {
				sc = result.rilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, cloudServiceName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "List", resp, "Failure sending request")
		return
	}

	result.rilr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rilr.hasNextLink() && result.rilr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client CloudServiceRoleInstancesClient) ListPreparer(ctx context.Context, resourceGroupName string, cloudServiceName string, expand InstanceViewTypes) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudServiceName":  autorest.Encode("path", cloudServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CloudServiceRoleInstancesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CloudServiceRoleInstancesClient) ListResponder(resp *http.Response) (result RoleInstanceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client CloudServiceRoleInstancesClient) listNextResults(ctx context.Context, lastResults RoleInstanceListResult) (result RoleInstanceListResult, err error) {
	req, err := lastResults.roleInstanceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client CloudServiceRoleInstancesClient) ListComplete(ctx context.Context, resourceGroupName string, cloudServiceName string, expand InstanceViewTypes) (result RoleInstanceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudServiceRoleInstancesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, cloudServiceName, expand)
	return
}

// Rebuild the Rebuild Role Instance asynchronous operation reinstalls the operating system on instances of web roles
// or worker roles and initializes the storage resources that are used by them. If you do not want to initialize
// storage resources, you can use Reimage Role Instance.
// Parameters:
// roleInstanceName - name of the role instance.
func (client CloudServiceRoleInstancesClient) Rebuild(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string) (result CloudServiceRoleInstancesRebuildFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudServiceRoleInstancesClient.Rebuild")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RebuildPreparer(ctx, roleInstanceName, resourceGroupName, cloudServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "Rebuild", nil, "Failure preparing request")
		return
	}

	result, err = client.RebuildSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "Rebuild", result.Response(), "Failure sending request")
		return
	}

	return
}

// RebuildPreparer prepares the Rebuild request.
func (client CloudServiceRoleInstancesClient) RebuildPreparer(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudServiceName":  autorest.Encode("path", cloudServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"roleInstanceName":  autorest.Encode("path", roleInstanceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}/rebuild", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RebuildSender sends the Rebuild request. The method will close the
// http.Response Body if it receives an error.
func (client CloudServiceRoleInstancesClient) RebuildSender(req *http.Request) (future CloudServiceRoleInstancesRebuildFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// RebuildResponder handles the response to the Rebuild request. The method always
// closes the http.Response Body.
func (client CloudServiceRoleInstancesClient) RebuildResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Reimage the Reimage Role Instance asynchronous operation reinstalls the operating system on instances of web roles
// or worker roles.
// Parameters:
// roleInstanceName - name of the role instance.
func (client CloudServiceRoleInstancesClient) Reimage(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string) (result CloudServiceRoleInstancesReimageFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudServiceRoleInstancesClient.Reimage")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ReimagePreparer(ctx, roleInstanceName, resourceGroupName, cloudServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "Reimage", nil, "Failure preparing request")
		return
	}

	result, err = client.ReimageSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "Reimage", result.Response(), "Failure sending request")
		return
	}

	return
}

// ReimagePreparer prepares the Reimage request.
func (client CloudServiceRoleInstancesClient) ReimagePreparer(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudServiceName":  autorest.Encode("path", cloudServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"roleInstanceName":  autorest.Encode("path", roleInstanceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}/reimage", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ReimageSender sends the Reimage request. The method will close the
// http.Response Body if it receives an error.
func (client CloudServiceRoleInstancesClient) ReimageSender(req *http.Request) (future CloudServiceRoleInstancesReimageFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// ReimageResponder handles the response to the Reimage request. The method always
// closes the http.Response Body.
func (client CloudServiceRoleInstancesClient) ReimageResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Restart the Reboot Role Instance asynchronous operation requests a reboot of a role instance in the cloud service.
// Parameters:
// roleInstanceName - name of the role instance.
func (client CloudServiceRoleInstancesClient) Restart(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string) (result CloudServiceRoleInstancesRestartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudServiceRoleInstancesClient.Restart")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RestartPreparer(ctx, roleInstanceName, resourceGroupName, cloudServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "Restart", nil, "Failure preparing request")
		return
	}

	result, err = client.RestartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServiceRoleInstancesClient", "Restart", result.Response(), "Failure sending request")
		return
	}

	return
}

// RestartPreparer prepares the Restart request.
func (client CloudServiceRoleInstancesClient) RestartPreparer(ctx context.Context, roleInstanceName string, resourceGroupName string, cloudServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudServiceName":  autorest.Encode("path", cloudServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"roleInstanceName":  autorest.Encode("path", roleInstanceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roleInstances/{roleInstanceName}/restart", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RestartSender sends the Restart request. The method will close the
// http.Response Body if it receives an error.
func (client CloudServiceRoleInstancesClient) RestartSender(req *http.Request) (future CloudServiceRoleInstancesRestartFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// RestartResponder handles the response to the Restart request. The method always
// closes the http.Response Body.
func (client CloudServiceRoleInstancesClient) RestartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
