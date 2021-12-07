//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybriddatamanager

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// JobsClient contains the methods for the Jobs group.
// Don't use this type directly, use NewJobsClient() instead.
type JobsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewJobsClient creates a new instance of JobsClient with the specified values.
func NewJobsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *JobsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &JobsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCancel - Cancels the given job.
// If the operation fails it returns a generic error.
func (client *JobsClient) BeginCancel(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, options *JobsBeginCancelOptions) (JobsCancelPollerResponse, error) {
	resp, err := client.cancel(ctx, dataServiceName, jobDefinitionName, jobID, resourceGroupName, dataManagerName, options)
	if err != nil {
		return JobsCancelPollerResponse{}, err
	}
	result := JobsCancelPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("JobsClient.Cancel", "", resp, client.pl, client.cancelHandleError)
	if err != nil {
		return JobsCancelPollerResponse{}, err
	}
	result.Poller = &JobsCancelPoller{
		pt: pt,
	}
	return result, nil
}

// Cancel - Cancels the given job.
// If the operation fails it returns a generic error.
func (client *JobsClient) cancel(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, options *JobsBeginCancelOptions) (*http.Response, error) {
	req, err := client.cancelCreateRequest(ctx, dataServiceName, jobDefinitionName, jobID, resourceGroupName, dataManagerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.cancelHandleError(resp)
	}
	return resp, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *JobsClient) cancelCreateRequest(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, options *JobsBeginCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs/{jobId}/cancel"
	if dataServiceName == "" {
		return nil, errors.New("parameter dataServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataServiceName}", url.PathEscape(dataServiceName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	if jobID == "" {
		return nil, errors.New("parameter jobID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobId}", url.PathEscape(jobID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// cancelHandleError handles the Cancel error response.
func (client *JobsClient) cancelHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - This method gets a data manager job given the jobId.
// If the operation fails it returns a generic error.
func (client *JobsClient) Get(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, options *JobsGetOptions) (JobsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, dataServiceName, jobDefinitionName, jobID, resourceGroupName, dataManagerName, options)
	if err != nil {
		return JobsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JobsClient) getCreateRequest(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, options *JobsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs/{jobId}"
	if dataServiceName == "" {
		return nil, errors.New("parameter dataServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataServiceName}", url.PathEscape(dataServiceName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	if jobID == "" {
		return nil, errors.New("parameter jobID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobId}", url.PathEscape(jobID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobsClient) getHandleResponse(resp *http.Response) (JobsGetResponse, error) {
	result := JobsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Job); err != nil {
		return JobsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *JobsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByDataManager - This method gets all the jobs at the data manager resource level.
// If the operation fails it returns a generic error.
func (client *JobsClient) ListByDataManager(resourceGroupName string, dataManagerName string, options *JobsListByDataManagerOptions) *JobsListByDataManagerPager {
	return &JobsListByDataManagerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDataManagerCreateRequest(ctx, resourceGroupName, dataManagerName, options)
		},
		advancer: func(ctx context.Context, resp JobsListByDataManagerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.JobList.NextLink)
		},
	}
}

// listByDataManagerCreateRequest creates the ListByDataManager request.
func (client *JobsClient) listByDataManagerCreateRequest(ctx context.Context, resourceGroupName string, dataManagerName string, options *JobsListByDataManagerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/jobs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDataManagerHandleResponse handles the ListByDataManager response.
func (client *JobsClient) listByDataManagerHandleResponse(resp *http.Response) (JobsListByDataManagerResponse, error) {
	result := JobsListByDataManagerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobList); err != nil {
		return JobsListByDataManagerResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByDataManagerHandleError handles the ListByDataManager error response.
func (client *JobsClient) listByDataManagerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByDataService - This method gets all the jobs of a data service type in a given resource.
// If the operation fails it returns a generic error.
func (client *JobsClient) ListByDataService(dataServiceName string, resourceGroupName string, dataManagerName string, options *JobsListByDataServiceOptions) *JobsListByDataServicePager {
	return &JobsListByDataServicePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDataServiceCreateRequest(ctx, dataServiceName, resourceGroupName, dataManagerName, options)
		},
		advancer: func(ctx context.Context, resp JobsListByDataServiceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.JobList.NextLink)
		},
	}
}

// listByDataServiceCreateRequest creates the ListByDataService request.
func (client *JobsClient) listByDataServiceCreateRequest(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string, options *JobsListByDataServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobs"
	if dataServiceName == "" {
		return nil, errors.New("parameter dataServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataServiceName}", url.PathEscape(dataServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDataServiceHandleResponse handles the ListByDataService response.
func (client *JobsClient) listByDataServiceHandleResponse(resp *http.Response) (JobsListByDataServiceResponse, error) {
	result := JobsListByDataServiceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobList); err != nil {
		return JobsListByDataServiceResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByDataServiceHandleError handles the ListByDataService error response.
func (client *JobsClient) listByDataServiceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByJobDefinition - This method gets all the jobs of a given job definition.
// If the operation fails it returns a generic error.
func (client *JobsClient) ListByJobDefinition(dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, options *JobsListByJobDefinitionOptions) *JobsListByJobDefinitionPager {
	return &JobsListByJobDefinitionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByJobDefinitionCreateRequest(ctx, dataServiceName, jobDefinitionName, resourceGroupName, dataManagerName, options)
		},
		advancer: func(ctx context.Context, resp JobsListByJobDefinitionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.JobList.NextLink)
		},
	}
}

// listByJobDefinitionCreateRequest creates the ListByJobDefinition request.
func (client *JobsClient) listByJobDefinitionCreateRequest(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, options *JobsListByJobDefinitionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs"
	if dataServiceName == "" {
		return nil, errors.New("parameter dataServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataServiceName}", url.PathEscape(dataServiceName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByJobDefinitionHandleResponse handles the ListByJobDefinition response.
func (client *JobsClient) listByJobDefinitionHandleResponse(resp *http.Response) (JobsListByJobDefinitionResponse, error) {
	result := JobsListByJobDefinitionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobList); err != nil {
		return JobsListByJobDefinitionResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByJobDefinitionHandleError handles the ListByJobDefinition error response.
func (client *JobsClient) listByJobDefinitionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginResume - Resumes the given job.
// If the operation fails it returns a generic error.
func (client *JobsClient) BeginResume(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, options *JobsBeginResumeOptions) (JobsResumePollerResponse, error) {
	resp, err := client.resume(ctx, dataServiceName, jobDefinitionName, jobID, resourceGroupName, dataManagerName, options)
	if err != nil {
		return JobsResumePollerResponse{}, err
	}
	result := JobsResumePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("JobsClient.Resume", "", resp, client.pl, client.resumeHandleError)
	if err != nil {
		return JobsResumePollerResponse{}, err
	}
	result.Poller = &JobsResumePoller{
		pt: pt,
	}
	return result, nil
}

// Resume - Resumes the given job.
// If the operation fails it returns a generic error.
func (client *JobsClient) resume(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, options *JobsBeginResumeOptions) (*http.Response, error) {
	req, err := client.resumeCreateRequest(ctx, dataServiceName, jobDefinitionName, jobID, resourceGroupName, dataManagerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.resumeHandleError(resp)
	}
	return resp, nil
}

// resumeCreateRequest creates the Resume request.
func (client *JobsClient) resumeCreateRequest(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, options *JobsBeginResumeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs/{jobId}/resume"
	if dataServiceName == "" {
		return nil, errors.New("parameter dataServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataServiceName}", url.PathEscape(dataServiceName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	if jobID == "" {
		return nil, errors.New("parameter jobID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobId}", url.PathEscape(jobID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// resumeHandleError handles the Resume error response.
func (client *JobsClient) resumeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}