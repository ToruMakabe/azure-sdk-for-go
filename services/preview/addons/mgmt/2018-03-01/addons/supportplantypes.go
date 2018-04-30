package addons

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// SupportPlanTypesClient is the the service for managing third party addons.
type SupportPlanTypesClient struct {
	BaseClient
}

// NewSupportPlanTypesClient creates an instance of the SupportPlanTypesClient client.
func NewSupportPlanTypesClient(subscriptionID string) SupportPlanTypesClient {
	return NewSupportPlanTypesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSupportPlanTypesClientWithBaseURI creates an instance of the SupportPlanTypesClient client.
func NewSupportPlanTypesClientWithBaseURI(baseURI string, subscriptionID string) SupportPlanTypesClient {
	return SupportPlanTypesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the Canonical support plan of type {type} for the subscription.
// Parameters:
// providerName - the support plan type. For now the only valid type is "canonical".
// planTypeName - the Canonical support plan type.
func (client SupportPlanTypesClient) CreateOrUpdate(ctx context.Context, providerName string, planTypeName PlanTypeName) (result SupportPlanTypesCreateOrUpdateFuture, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, providerName, planTypeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "addons.SupportPlanTypesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "addons.SupportPlanTypesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SupportPlanTypesClient) CreateOrUpdatePreparer(ctx context.Context, providerName string, planTypeName PlanTypeName) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"planTypeName":   autorest.Encode("path", planTypeName),
		"providerName":   autorest.Encode("path", providerName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Addons/supportProviders/{providerName}/supportPlanTypes/{planTypeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SupportPlanTypesClient) CreateOrUpdateSender(req *http.Request) (future SupportPlanTypesCreateOrUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusNotFound))
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SupportPlanTypesClient) CreateOrUpdateResponder(resp *http.Response) (result CanonicalSupportPlanResponseEnvelope, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete cancels the Canonical support plan of type {type} for the subscription.
// Parameters:
// providerName - the support plan type. For now the only valid type is "canonical".
// planTypeName - the Canonical support plan type.
func (client SupportPlanTypesClient) Delete(ctx context.Context, providerName string, planTypeName PlanTypeName) (result SupportPlanTypesDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, providerName, planTypeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "addons.SupportPlanTypesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "addons.SupportPlanTypesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SupportPlanTypesClient) DeletePreparer(ctx context.Context, providerName string, planTypeName PlanTypeName) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"planTypeName":   autorest.Encode("path", planTypeName),
		"providerName":   autorest.Encode("path", providerName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Addons/supportProviders/{providerName}/supportPlanTypes/{planTypeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SupportPlanTypesClient) DeleteSender(req *http.Request) (future SupportPlanTypesDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SupportPlanTypesClient) DeleteResponder(resp *http.Response) (result CanonicalSupportPlanResponseEnvelope, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get returns whether or not the canonical support plan of type {type} is enabled for the subscription.
// Parameters:
// providerName - the support plan type. For now the only valid type is "canonical".
// planTypeName - the Canonical support plan type.
func (client SupportPlanTypesClient) Get(ctx context.Context, providerName string, planTypeName PlanTypeName) (result CanonicalSupportPlanResponseEnvelope, err error) {
	req, err := client.GetPreparer(ctx, providerName, planTypeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "addons.SupportPlanTypesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "addons.SupportPlanTypesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "addons.SupportPlanTypesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SupportPlanTypesClient) GetPreparer(ctx context.Context, providerName string, planTypeName PlanTypeName) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"planTypeName":   autorest.Encode("path", planTypeName),
		"providerName":   autorest.Encode("path", providerName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Addons/supportProviders/{providerName}/supportPlanTypes/{planTypeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SupportPlanTypesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SupportPlanTypesClient) GetResponder(resp *http.Response) (result CanonicalSupportPlanResponseEnvelope, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListInfo returns the canonical support plan information for all types for the subscription.
func (client SupportPlanTypesClient) ListInfo(ctx context.Context) (result ListCanonicalSupportPlanInfoDefinition, err error) {
	req, err := client.ListInfoPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "addons.SupportPlanTypesClient", "ListInfo", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListInfoSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "addons.SupportPlanTypesClient", "ListInfo", resp, "Failure sending request")
		return
	}

	result, err = client.ListInfoResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "addons.SupportPlanTypesClient", "ListInfo", resp, "Failure responding to request")
	}

	return
}

// ListInfoPreparer prepares the ListInfo request.
func (client SupportPlanTypesClient) ListInfoPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Addons/supportProviders/canonical/listSupportPlanInfo", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListInfoSender sends the ListInfo request. The method will close the
// http.Response Body if it receives an error.
func (client SupportPlanTypesClient) ListInfoSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListInfoResponder handles the response to the ListInfo request. The method always
// closes the http.Response Body.
func (client SupportPlanTypesClient) ListInfoResponder(resp *http.Response) (result ListCanonicalSupportPlanInfoDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
