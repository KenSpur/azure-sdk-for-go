//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

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
	"strconv"
	"strings"
)

// ServiceFabricSchedulesClient contains the methods for the ServiceFabricSchedules group.
// Don't use this type directly, use NewServiceFabricSchedulesClient() instead.
type ServiceFabricSchedulesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewServiceFabricSchedulesClient creates a new instance of ServiceFabricSchedulesClient with the specified values.
func NewServiceFabricSchedulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ServiceFabricSchedulesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ServiceFabricSchedulesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Create or replace an existing schedule.
// If the operation fails it returns the *CloudError error type.
func (client *ServiceFabricSchedulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, schedule Schedule, options *ServiceFabricSchedulesCreateOrUpdateOptions) (ServiceFabricSchedulesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, userName, serviceFabricName, name, schedule, options)
	if err != nil {
		return ServiceFabricSchedulesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceFabricSchedulesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ServiceFabricSchedulesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServiceFabricSchedulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, schedule Schedule, options *ServiceFabricSchedulesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}"
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
	if serviceFabricName == "" {
		return nil, errors.New("parameter serviceFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceFabricName}", url.PathEscape(serviceFabricName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, schedule)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ServiceFabricSchedulesClient) createOrUpdateHandleResponse(resp *http.Response) (ServiceFabricSchedulesCreateOrUpdateResponse, error) {
	result := ServiceFabricSchedulesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return ServiceFabricSchedulesCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ServiceFabricSchedulesClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - Delete schedule.
// If the operation fails it returns the *CloudError error type.
func (client *ServiceFabricSchedulesClient) Delete(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, options *ServiceFabricSchedulesDeleteOptions) (ServiceFabricSchedulesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, userName, serviceFabricName, name, options)
	if err != nil {
		return ServiceFabricSchedulesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceFabricSchedulesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ServiceFabricSchedulesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return ServiceFabricSchedulesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServiceFabricSchedulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, options *ServiceFabricSchedulesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}"
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
	if serviceFabricName == "" {
		return nil, errors.New("parameter serviceFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceFabricName}", url.PathEscape(serviceFabricName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ServiceFabricSchedulesClient) deleteHandleError(resp *http.Response) error {
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

// BeginExecute - Execute a schedule. This operation can take a while to complete.
// If the operation fails it returns the *CloudError error type.
func (client *ServiceFabricSchedulesClient) BeginExecute(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, options *ServiceFabricSchedulesBeginExecuteOptions) (ServiceFabricSchedulesExecutePollerResponse, error) {
	resp, err := client.execute(ctx, resourceGroupName, labName, userName, serviceFabricName, name, options)
	if err != nil {
		return ServiceFabricSchedulesExecutePollerResponse{}, err
	}
	result := ServiceFabricSchedulesExecutePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServiceFabricSchedulesClient.Execute", "", resp, client.pl, client.executeHandleError)
	if err != nil {
		return ServiceFabricSchedulesExecutePollerResponse{}, err
	}
	result.Poller = &ServiceFabricSchedulesExecutePoller{
		pt: pt,
	}
	return result, nil
}

// Execute - Execute a schedule. This operation can take a while to complete.
// If the operation fails it returns the *CloudError error type.
func (client *ServiceFabricSchedulesClient) execute(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, options *ServiceFabricSchedulesBeginExecuteOptions) (*http.Response, error) {
	req, err := client.executeCreateRequest(ctx, resourceGroupName, labName, userName, serviceFabricName, name, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.executeHandleError(resp)
	}
	return resp, nil
}

// executeCreateRequest creates the Execute request.
func (client *ServiceFabricSchedulesClient) executeCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, options *ServiceFabricSchedulesBeginExecuteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}/execute"
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
	if serviceFabricName == "" {
		return nil, errors.New("parameter serviceFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceFabricName}", url.PathEscape(serviceFabricName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// executeHandleError handles the Execute error response.
func (client *ServiceFabricSchedulesClient) executeHandleError(resp *http.Response) error {
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

// Get - Get schedule.
// If the operation fails it returns the *CloudError error type.
func (client *ServiceFabricSchedulesClient) Get(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, options *ServiceFabricSchedulesGetOptions) (ServiceFabricSchedulesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, userName, serviceFabricName, name, options)
	if err != nil {
		return ServiceFabricSchedulesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceFabricSchedulesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceFabricSchedulesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServiceFabricSchedulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, options *ServiceFabricSchedulesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}"
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
	if serviceFabricName == "" {
		return nil, errors.New("parameter serviceFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceFabricName}", url.PathEscape(serviceFabricName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceFabricSchedulesClient) getHandleResponse(resp *http.Response) (ServiceFabricSchedulesGetResponse, error) {
	result := ServiceFabricSchedulesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return ServiceFabricSchedulesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ServiceFabricSchedulesClient) getHandleError(resp *http.Response) error {
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

// List - List schedules in a given service fabric.
// If the operation fails it returns the *CloudError error type.
func (client *ServiceFabricSchedulesClient) List(resourceGroupName string, labName string, userName string, serviceFabricName string, options *ServiceFabricSchedulesListOptions) *ServiceFabricSchedulesListPager {
	return &ServiceFabricSchedulesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, labName, userName, serviceFabricName, options)
		},
		advancer: func(ctx context.Context, resp ServiceFabricSchedulesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ScheduleList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ServiceFabricSchedulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, options *ServiceFabricSchedulesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules"
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
	if serviceFabricName == "" {
		return nil, errors.New("parameter serviceFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceFabricName}", url.PathEscape(serviceFabricName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
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
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServiceFabricSchedulesClient) listHandleResponse(resp *http.Response) (ServiceFabricSchedulesListResponse, error) {
	result := ServiceFabricSchedulesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScheduleList); err != nil {
		return ServiceFabricSchedulesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ServiceFabricSchedulesClient) listHandleError(resp *http.Response) error {
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

// Update - Allows modifying tags of schedules. All other properties will be ignored.
// If the operation fails it returns the *CloudError error type.
func (client *ServiceFabricSchedulesClient) Update(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, schedule ScheduleFragment, options *ServiceFabricSchedulesUpdateOptions) (ServiceFabricSchedulesUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labName, userName, serviceFabricName, name, schedule, options)
	if err != nil {
		return ServiceFabricSchedulesUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceFabricSchedulesUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceFabricSchedulesUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ServiceFabricSchedulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, schedule ScheduleFragment, options *ServiceFabricSchedulesUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}"
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
	if serviceFabricName == "" {
		return nil, errors.New("parameter serviceFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceFabricName}", url.PathEscape(serviceFabricName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, schedule)
}

// updateHandleResponse handles the Update response.
func (client *ServiceFabricSchedulesClient) updateHandleResponse(resp *http.Response) (ServiceFabricSchedulesUpdateResponse, error) {
	result := ServiceFabricSchedulesUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return ServiceFabricSchedulesUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *ServiceFabricSchedulesClient) updateHandleError(resp *http.Response) error {
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