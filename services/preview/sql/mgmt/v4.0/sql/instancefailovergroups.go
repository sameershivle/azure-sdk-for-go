package sql

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// InstanceFailoverGroupsClient is the the Azure SQL Database management API provides a RESTful set of web services
// that interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve,
// update, and delete databases.
type InstanceFailoverGroupsClient struct {
	BaseClient
}

// NewInstanceFailoverGroupsClient creates an instance of the InstanceFailoverGroupsClient client.
func NewInstanceFailoverGroupsClient(subscriptionID string) InstanceFailoverGroupsClient {
	return NewInstanceFailoverGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInstanceFailoverGroupsClientWithBaseURI creates an instance of the InstanceFailoverGroupsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewInstanceFailoverGroupsClientWithBaseURI(baseURI string, subscriptionID string) InstanceFailoverGroupsClient {
	return InstanceFailoverGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a failover group.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// locationName - the name of the region where the resource is located.
// failoverGroupName - the name of the failover group.
// parameters - the failover group parameters.
func (client InstanceFailoverGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, parameters InstanceFailoverGroup) (result InstanceFailoverGroupsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InstanceFailoverGroupsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.InstanceFailoverGroupProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.InstanceFailoverGroupProperties.ReadWriteEndpoint", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.InstanceFailoverGroupProperties.PartnerRegions", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.InstanceFailoverGroupProperties.ManagedInstancePairs", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("sql.InstanceFailoverGroupsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, locationName, failoverGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.InstanceFailoverGroupsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.InstanceFailoverGroupsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client InstanceFailoverGroupsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, parameters InstanceFailoverGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"failoverGroupName": autorest.Encode("path", failoverGroupName),
		"locationName":      autorest.Encode("path", locationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups/{failoverGroupName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client InstanceFailoverGroupsClient) CreateOrUpdateSender(req *http.Request) (future InstanceFailoverGroupsCreateOrUpdateFuture, err error) {
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

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client InstanceFailoverGroupsClient) CreateOrUpdateResponder(resp *http.Response) (result InstanceFailoverGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a failover group.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// locationName - the name of the region where the resource is located.
// failoverGroupName - the name of the failover group.
func (client InstanceFailoverGroupsClient) Delete(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string) (result InstanceFailoverGroupsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InstanceFailoverGroupsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, locationName, failoverGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.InstanceFailoverGroupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.InstanceFailoverGroupsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client InstanceFailoverGroupsClient) DeletePreparer(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"failoverGroupName": autorest.Encode("path", failoverGroupName),
		"locationName":      autorest.Encode("path", locationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups/{failoverGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client InstanceFailoverGroupsClient) DeleteSender(req *http.Request) (future InstanceFailoverGroupsDeleteFuture, err error) {
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
func (client InstanceFailoverGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Failover fails over from the current primary managed instance to this managed instance.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// locationName - the name of the region where the resource is located.
// failoverGroupName - the name of the failover group.
func (client InstanceFailoverGroupsClient) Failover(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string) (result InstanceFailoverGroupsFailoverFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InstanceFailoverGroupsClient.Failover")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.FailoverPreparer(ctx, resourceGroupName, locationName, failoverGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.InstanceFailoverGroupsClient", "Failover", nil, "Failure preparing request")
		return
	}

	result, err = client.FailoverSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.InstanceFailoverGroupsClient", "Failover", result.Response(), "Failure sending request")
		return
	}

	return
}

// FailoverPreparer prepares the Failover request.
func (client InstanceFailoverGroupsClient) FailoverPreparer(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"failoverGroupName": autorest.Encode("path", failoverGroupName),
		"locationName":      autorest.Encode("path", locationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups/{failoverGroupName}/failover", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// FailoverSender sends the Failover request. The method will close the
// http.Response Body if it receives an error.
func (client InstanceFailoverGroupsClient) FailoverSender(req *http.Request) (future InstanceFailoverGroupsFailoverFuture, err error) {
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

// FailoverResponder handles the response to the Failover request. The method always
// closes the http.Response Body.
func (client InstanceFailoverGroupsClient) FailoverResponder(resp *http.Response) (result InstanceFailoverGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ForceFailoverAllowDataLoss fails over from the current primary managed instance to this managed instance. This
// operation might result in data loss.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// locationName - the name of the region where the resource is located.
// failoverGroupName - the name of the failover group.
func (client InstanceFailoverGroupsClient) ForceFailoverAllowDataLoss(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string) (result InstanceFailoverGroupsForceFailoverAllowDataLossFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InstanceFailoverGroupsClient.ForceFailoverAllowDataLoss")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ForceFailoverAllowDataLossPreparer(ctx, resourceGroupName, locationName, failoverGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.InstanceFailoverGroupsClient", "ForceFailoverAllowDataLoss", nil, "Failure preparing request")
		return
	}

	result, err = client.ForceFailoverAllowDataLossSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.InstanceFailoverGroupsClient", "ForceFailoverAllowDataLoss", result.Response(), "Failure sending request")
		return
	}

	return
}

// ForceFailoverAllowDataLossPreparer prepares the ForceFailoverAllowDataLoss request.
func (client InstanceFailoverGroupsClient) ForceFailoverAllowDataLossPreparer(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"failoverGroupName": autorest.Encode("path", failoverGroupName),
		"locationName":      autorest.Encode("path", locationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups/{failoverGroupName}/forceFailoverAllowDataLoss", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ForceFailoverAllowDataLossSender sends the ForceFailoverAllowDataLoss request. The method will close the
// http.Response Body if it receives an error.
func (client InstanceFailoverGroupsClient) ForceFailoverAllowDataLossSender(req *http.Request) (future InstanceFailoverGroupsForceFailoverAllowDataLossFuture, err error) {
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

// ForceFailoverAllowDataLossResponder handles the response to the ForceFailoverAllowDataLoss request. The method always
// closes the http.Response Body.
func (client InstanceFailoverGroupsClient) ForceFailoverAllowDataLossResponder(resp *http.Response) (result InstanceFailoverGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a failover group.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// locationName - the name of the region where the resource is located.
// failoverGroupName - the name of the failover group.
func (client InstanceFailoverGroupsClient) Get(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string) (result InstanceFailoverGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InstanceFailoverGroupsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, locationName, failoverGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.InstanceFailoverGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.InstanceFailoverGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.InstanceFailoverGroupsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client InstanceFailoverGroupsClient) GetPreparer(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"failoverGroupName": autorest.Encode("path", failoverGroupName),
		"locationName":      autorest.Encode("path", locationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups/{failoverGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client InstanceFailoverGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InstanceFailoverGroupsClient) GetResponder(resp *http.Response) (result InstanceFailoverGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByLocation lists the failover groups in a location.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// locationName - the name of the region where the resource is located.
func (client InstanceFailoverGroupsClient) ListByLocation(ctx context.Context, resourceGroupName string, locationName string) (result InstanceFailoverGroupListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InstanceFailoverGroupsClient.ListByLocation")
		defer func() {
			sc := -1
			if result.ifglr.Response.Response != nil {
				sc = result.ifglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByLocationNextResults
	req, err := client.ListByLocationPreparer(ctx, resourceGroupName, locationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.InstanceFailoverGroupsClient", "ListByLocation", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByLocationSender(req)
	if err != nil {
		result.ifglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.InstanceFailoverGroupsClient", "ListByLocation", resp, "Failure sending request")
		return
	}

	result.ifglr, err = client.ListByLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.InstanceFailoverGroupsClient", "ListByLocation", resp, "Failure responding to request")
		return
	}
	if result.ifglr.hasNextLink() && result.ifglr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByLocationPreparer prepares the ListByLocation request.
func (client InstanceFailoverGroupsClient) ListByLocationPreparer(ctx context.Context, resourceGroupName string, locationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":      autorest.Encode("path", locationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByLocationSender sends the ListByLocation request. The method will close the
// http.Response Body if it receives an error.
func (client InstanceFailoverGroupsClient) ListByLocationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByLocationResponder handles the response to the ListByLocation request. The method always
// closes the http.Response Body.
func (client InstanceFailoverGroupsClient) ListByLocationResponder(resp *http.Response) (result InstanceFailoverGroupListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByLocationNextResults retrieves the next set of results, if any.
func (client InstanceFailoverGroupsClient) listByLocationNextResults(ctx context.Context, lastResults InstanceFailoverGroupListResult) (result InstanceFailoverGroupListResult, err error) {
	req, err := lastResults.instanceFailoverGroupListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.InstanceFailoverGroupsClient", "listByLocationNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByLocationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.InstanceFailoverGroupsClient", "listByLocationNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.InstanceFailoverGroupsClient", "listByLocationNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByLocationComplete enumerates all values, automatically crossing page boundaries as required.
func (client InstanceFailoverGroupsClient) ListByLocationComplete(ctx context.Context, resourceGroupName string, locationName string) (result InstanceFailoverGroupListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InstanceFailoverGroupsClient.ListByLocation")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByLocation(ctx, resourceGroupName, locationName)
	return
}
