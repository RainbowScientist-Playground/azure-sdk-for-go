// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armiotoperations

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DataflowEndpointClient contains the methods for the DataflowEndpoint group.
// Don't use this type directly, use NewDataflowEndpointClient() instead.
type DataflowEndpointClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDataflowEndpointClient creates a new instance of DataflowEndpointClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDataflowEndpointClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataflowEndpointClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DataflowEndpointClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a DataflowEndpointResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - instanceName - Name of instance.
//   - dataflowEndpointName - Name of Instance dataflowEndpoint resource
//   - resource - Resource create parameters.
//   - options - DataflowEndpointClientBeginCreateOrUpdateOptions contains the optional parameters for the DataflowEndpointClient.BeginCreateOrUpdate
//     method.
func (client *DataflowEndpointClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, instanceName string, dataflowEndpointName string, resource DataflowEndpointResource, options *DataflowEndpointClientBeginCreateOrUpdateOptions) (*runtime.Poller[DataflowEndpointClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, instanceName, dataflowEndpointName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DataflowEndpointClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DataflowEndpointClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a DataflowEndpointResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
func (client *DataflowEndpointClient) createOrUpdate(ctx context.Context, resourceGroupName string, instanceName string, dataflowEndpointName string, resource DataflowEndpointResource, options *DataflowEndpointClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DataflowEndpointClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, instanceName, dataflowEndpointName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DataflowEndpointClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, instanceName string, dataflowEndpointName string, resource DataflowEndpointResource, _ *DataflowEndpointClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTOperations/instances/{instanceName}/dataflowEndpoints/{dataflowEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	if dataflowEndpointName == "" {
		return nil, errors.New("parameter dataflowEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataflowEndpointName}", url.PathEscape(dataflowEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a DataflowEndpointResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - instanceName - Name of instance.
//   - dataflowEndpointName - Name of Instance dataflowEndpoint resource
//   - options - DataflowEndpointClientBeginDeleteOptions contains the optional parameters for the DataflowEndpointClient.BeginDelete
//     method.
func (client *DataflowEndpointClient) BeginDelete(ctx context.Context, resourceGroupName string, instanceName string, dataflowEndpointName string, options *DataflowEndpointClientBeginDeleteOptions) (*runtime.Poller[DataflowEndpointClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, instanceName, dataflowEndpointName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DataflowEndpointClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DataflowEndpointClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a DataflowEndpointResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
func (client *DataflowEndpointClient) deleteOperation(ctx context.Context, resourceGroupName string, instanceName string, dataflowEndpointName string, options *DataflowEndpointClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DataflowEndpointClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, instanceName, dataflowEndpointName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataflowEndpointClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, instanceName string, dataflowEndpointName string, _ *DataflowEndpointClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTOperations/instances/{instanceName}/dataflowEndpoints/{dataflowEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	if dataflowEndpointName == "" {
		return nil, errors.New("parameter dataflowEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataflowEndpointName}", url.PathEscape(dataflowEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a DataflowEndpointResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - instanceName - Name of instance.
//   - dataflowEndpointName - Name of Instance dataflowEndpoint resource
//   - options - DataflowEndpointClientGetOptions contains the optional parameters for the DataflowEndpointClient.Get method.
func (client *DataflowEndpointClient) Get(ctx context.Context, resourceGroupName string, instanceName string, dataflowEndpointName string, options *DataflowEndpointClientGetOptions) (DataflowEndpointClientGetResponse, error) {
	var err error
	const operationName = "DataflowEndpointClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, instanceName, dataflowEndpointName, options)
	if err != nil {
		return DataflowEndpointClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataflowEndpointClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DataflowEndpointClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DataflowEndpointClient) getCreateRequest(ctx context.Context, resourceGroupName string, instanceName string, dataflowEndpointName string, _ *DataflowEndpointClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTOperations/instances/{instanceName}/dataflowEndpoints/{dataflowEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	if dataflowEndpointName == "" {
		return nil, errors.New("parameter dataflowEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataflowEndpointName}", url.PathEscape(dataflowEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataflowEndpointClient) getHandleResponse(resp *http.Response) (DataflowEndpointClientGetResponse, error) {
	result := DataflowEndpointClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataflowEndpointResource); err != nil {
		return DataflowEndpointClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List DataflowEndpointResource resources by InstanceResource
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - instanceName - Name of instance.
//   - options - DataflowEndpointClientListByResourceGroupOptions contains the optional parameters for the DataflowEndpointClient.NewListByResourceGroupPager
//     method.
func (client *DataflowEndpointClient) NewListByResourceGroupPager(resourceGroupName string, instanceName string, options *DataflowEndpointClientListByResourceGroupOptions) *runtime.Pager[DataflowEndpointClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataflowEndpointClientListByResourceGroupResponse]{
		More: func(page DataflowEndpointClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataflowEndpointClientListByResourceGroupResponse) (DataflowEndpointClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DataflowEndpointClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, instanceName, options)
			}, nil)
			if err != nil {
				return DataflowEndpointClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DataflowEndpointClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, instanceName string, _ *DataflowEndpointClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTOperations/instances/{instanceName}/dataflowEndpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DataflowEndpointClient) listByResourceGroupHandleResponse(resp *http.Response) (DataflowEndpointClientListByResourceGroupResponse, error) {
	result := DataflowEndpointClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataflowEndpointResourceListResult); err != nil {
		return DataflowEndpointClientListByResourceGroupResponse{}, err
	}
	return result, nil
}
