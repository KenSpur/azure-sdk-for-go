//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// DatabaseRecommendedActionsClient contains the methods for the DatabaseRecommendedActions group.
// Don't use this type directly, use NewDatabaseRecommendedActionsClient() instead.
type DatabaseRecommendedActionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDatabaseRecommendedActionsClient creates a new instance of DatabaseRecommendedActionsClient with the specified values.
func NewDatabaseRecommendedActionsClient(con *arm.Connection, subscriptionID string) *DatabaseRecommendedActionsClient {
	return &DatabaseRecommendedActionsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Gets a database recommended action.
// If the operation fails it returns a generic error.
func (client *DatabaseRecommendedActionsClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string, recommendedActionName string, options *DatabaseRecommendedActionsGetOptions) (DatabaseRecommendedActionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, advisorName, recommendedActionName, options)
	if err != nil {
		return DatabaseRecommendedActionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabaseRecommendedActionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabaseRecommendedActionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DatabaseRecommendedActionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string, recommendedActionName string, options *DatabaseRecommendedActionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advisors/{advisorName}/recommendedActions/{recommendedActionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if advisorName == "" {
		return nil, errors.New("parameter advisorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advisorName}", url.PathEscape(advisorName))
	if recommendedActionName == "" {
		return nil, errors.New("parameter recommendedActionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recommendedActionName}", url.PathEscape(recommendedActionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatabaseRecommendedActionsClient) getHandleResponse(resp *http.Response) (DatabaseRecommendedActionsGetResponse, error) {
	result := DatabaseRecommendedActionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecommendedAction); err != nil {
		return DatabaseRecommendedActionsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DatabaseRecommendedActionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByDatabaseAdvisor - Gets list of Database Recommended Actions.
// If the operation fails it returns a generic error.
func (client *DatabaseRecommendedActionsClient) ListByDatabaseAdvisor(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string, options *DatabaseRecommendedActionsListByDatabaseAdvisorOptions) (DatabaseRecommendedActionsListByDatabaseAdvisorResponse, error) {
	req, err := client.listByDatabaseAdvisorCreateRequest(ctx, resourceGroupName, serverName, databaseName, advisorName, options)
	if err != nil {
		return DatabaseRecommendedActionsListByDatabaseAdvisorResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabaseRecommendedActionsListByDatabaseAdvisorResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabaseRecommendedActionsListByDatabaseAdvisorResponse{}, client.listByDatabaseAdvisorHandleError(resp)
	}
	return client.listByDatabaseAdvisorHandleResponse(resp)
}

// listByDatabaseAdvisorCreateRequest creates the ListByDatabaseAdvisor request.
func (client *DatabaseRecommendedActionsClient) listByDatabaseAdvisorCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string, options *DatabaseRecommendedActionsListByDatabaseAdvisorOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advisors/{advisorName}/recommendedActions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if advisorName == "" {
		return nil, errors.New("parameter advisorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advisorName}", url.PathEscape(advisorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDatabaseAdvisorHandleResponse handles the ListByDatabaseAdvisor response.
func (client *DatabaseRecommendedActionsClient) listByDatabaseAdvisorHandleResponse(resp *http.Response) (DatabaseRecommendedActionsListByDatabaseAdvisorResponse, error) {
	result := DatabaseRecommendedActionsListByDatabaseAdvisorResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecommendedActionArray); err != nil {
		return DatabaseRecommendedActionsListByDatabaseAdvisorResponse{}, err
	}
	return result, nil
}

// listByDatabaseAdvisorHandleError handles the ListByDatabaseAdvisor error response.
func (client *DatabaseRecommendedActionsClient) listByDatabaseAdvisorHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Update - Updates a database recommended action.
// If the operation fails it returns a generic error.
func (client *DatabaseRecommendedActionsClient) Update(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string, recommendedActionName string, parameters RecommendedAction, options *DatabaseRecommendedActionsUpdateOptions) (DatabaseRecommendedActionsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, databaseName, advisorName, recommendedActionName, parameters, options)
	if err != nil {
		return DatabaseRecommendedActionsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabaseRecommendedActionsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabaseRecommendedActionsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *DatabaseRecommendedActionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string, recommendedActionName string, parameters RecommendedAction, options *DatabaseRecommendedActionsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advisors/{advisorName}/recommendedActions/{recommendedActionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if advisorName == "" {
		return nil, errors.New("parameter advisorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advisorName}", url.PathEscape(advisorName))
	if recommendedActionName == "" {
		return nil, errors.New("parameter recommendedActionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recommendedActionName}", url.PathEscape(recommendedActionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *DatabaseRecommendedActionsClient) updateHandleResponse(resp *http.Response) (DatabaseRecommendedActionsUpdateResponse, error) {
	result := DatabaseRecommendedActionsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecommendedAction); err != nil {
		return DatabaseRecommendedActionsUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *DatabaseRecommendedActionsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}