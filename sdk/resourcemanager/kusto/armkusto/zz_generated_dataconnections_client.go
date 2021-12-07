//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkusto

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DataConnectionsClient contains the methods for the DataConnections group.
// Don't use this type directly, use NewDataConnectionsClient() instead.
type DataConnectionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDataConnectionsClient creates a new instance of DataConnectionsClient with the specified values.
func NewDataConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DataConnectionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &DataConnectionsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CheckNameAvailability - Checks that the data connection name is valid and is not already in use.
// If the operation fails it returns the *CloudError error type.
func (client *DataConnectionsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName DataConnectionCheckNameRequest, options *DataConnectionsCheckNameAvailabilityOptions) (DataConnectionsCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName, options)
	if err != nil {
		return DataConnectionsCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataConnectionsCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataConnectionsCheckNameAvailabilityResponse{}, client.checkNameAvailabilityHandleError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *DataConnectionsClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName DataConnectionCheckNameRequest, options *DataConnectionsCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/checkNameAvailability"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, dataConnectionName)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *DataConnectionsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (DataConnectionsCheckNameAvailabilityResponse, error) {
	result := DataConnectionsCheckNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameResult); err != nil {
		return DataConnectionsCheckNameAvailabilityResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// checkNameAvailabilityHandleError handles the CheckNameAvailability error response.
func (client *DataConnectionsClient) checkNameAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginCreateOrUpdate - Creates or updates a data connection.
// If the operation fails it returns the *CloudError error type.
func (client *DataConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *DataConnectionsBeginCreateOrUpdateOptions) (DataConnectionsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName, parameters, options)
	if err != nil {
		return DataConnectionsCreateOrUpdatePollerResponse{}, err
	}
	result := DataConnectionsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DataConnectionsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return DataConnectionsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &DataConnectionsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a data connection.
// If the operation fails it returns the *CloudError error type.
func (client *DataConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *DataConnectionsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DataConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *DataConnectionsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if dataConnectionName == "" {
		return nil, errors.New("parameter dataConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectionName}", url.PathEscape(dataConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DataConnectionsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDataConnectionValidation - Checks that the data connection parameters are valid.
// If the operation fails it returns the *CloudError error type.
func (client *DataConnectionsClient) BeginDataConnectionValidation(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DataConnectionValidation, options *DataConnectionsBeginDataConnectionValidationOptions) (DataConnectionsDataConnectionValidationPollerResponse, error) {
	resp, err := client.dataConnectionValidation(ctx, resourceGroupName, clusterName, databaseName, parameters, options)
	if err != nil {
		return DataConnectionsDataConnectionValidationPollerResponse{}, err
	}
	result := DataConnectionsDataConnectionValidationPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DataConnectionsClient.DataConnectionValidation", "location", resp, client.pl, client.dataConnectionValidationHandleError)
	if err != nil {
		return DataConnectionsDataConnectionValidationPollerResponse{}, err
	}
	result.Poller = &DataConnectionsDataConnectionValidationPoller{
		pt: pt,
	}
	return result, nil
}

// DataConnectionValidation - Checks that the data connection parameters are valid.
// If the operation fails it returns the *CloudError error type.
func (client *DataConnectionsClient) dataConnectionValidation(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DataConnectionValidation, options *DataConnectionsBeginDataConnectionValidationOptions) (*http.Response, error) {
	req, err := client.dataConnectionValidationCreateRequest(ctx, resourceGroupName, clusterName, databaseName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.dataConnectionValidationHandleError(resp)
	}
	return resp, nil
}

// dataConnectionValidationCreateRequest creates the DataConnectionValidation request.
func (client *DataConnectionsClient) dataConnectionValidationCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DataConnectionValidation, options *DataConnectionsBeginDataConnectionValidationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnectionValidation"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// dataConnectionValidationHandleError handles the DataConnectionValidation error response.
func (client *DataConnectionsClient) dataConnectionValidationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the data connection with the given name.
// If the operation fails it returns the *CloudError error type.
func (client *DataConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, options *DataConnectionsBeginDeleteOptions) (DataConnectionsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName, options)
	if err != nil {
		return DataConnectionsDeletePollerResponse{}, err
	}
	result := DataConnectionsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DataConnectionsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return DataConnectionsDeletePollerResponse{}, err
	}
	result.Poller = &DataConnectionsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the data connection with the given name.
// If the operation fails it returns the *CloudError error type.
func (client *DataConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, options *DataConnectionsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, options *DataConnectionsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if dataConnectionName == "" {
		return nil, errors.New("parameter dataConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectionName}", url.PathEscape(dataConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DataConnectionsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Returns a data connection.
// If the operation fails it returns the *CloudError error type.
func (client *DataConnectionsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, options *DataConnectionsGetOptions) (DataConnectionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName, options)
	if err != nil {
		return DataConnectionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataConnectionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataConnectionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, options *DataConnectionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if dataConnectionName == "" {
		return nil, errors.New("parameter dataConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectionName}", url.PathEscape(dataConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataConnectionsClient) getHandleResponse(resp *http.Response) (DataConnectionsGetResponse, error) {
	result := DataConnectionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return DataConnectionsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DataConnectionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByDatabase - Returns the list of data connections of the given Kusto database.
// If the operation fails it returns the *CloudError error type.
func (client *DataConnectionsClient) ListByDatabase(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DataConnectionsListByDatabaseOptions) (DataConnectionsListByDatabaseResponse, error) {
	req, err := client.listByDatabaseCreateRequest(ctx, resourceGroupName, clusterName, databaseName, options)
	if err != nil {
		return DataConnectionsListByDatabaseResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataConnectionsListByDatabaseResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataConnectionsListByDatabaseResponse{}, client.listByDatabaseHandleError(resp)
	}
	return client.listByDatabaseHandleResponse(resp)
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *DataConnectionsClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DataConnectionsListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *DataConnectionsClient) listByDatabaseHandleResponse(resp *http.Response) (DataConnectionsListByDatabaseResponse, error) {
	result := DataConnectionsListByDatabaseResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataConnectionListResult); err != nil {
		return DataConnectionsListByDatabaseResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByDatabaseHandleError handles the ListByDatabase error response.
func (client *DataConnectionsClient) listByDatabaseHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Updates a data connection.
// If the operation fails it returns the *CloudError error type.
func (client *DataConnectionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *DataConnectionsBeginUpdateOptions) (DataConnectionsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName, parameters, options)
	if err != nil {
		return DataConnectionsUpdatePollerResponse{}, err
	}
	result := DataConnectionsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DataConnectionsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return DataConnectionsUpdatePollerResponse{}, err
	}
	result.Poller = &DataConnectionsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates a data connection.
// If the operation fails it returns the *CloudError error type.
func (client *DataConnectionsClient) update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *DataConnectionsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DataConnectionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters DataConnectionClassification, options *DataConnectionsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if dataConnectionName == "" {
		return nil, errors.New("parameter dataConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectionName}", url.PathEscape(dataConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleError handles the Update error response.
func (client *DataConnectionsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}