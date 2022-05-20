//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// DataConnectorsClient contains the methods for the DataConnectors group.
// Don't use this type directly, use NewDataConnectorsClient() instead.
type DataConnectorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDataConnectorsClient creates a new instance of DataConnectorsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDataConnectorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataConnectorsClient, error) {
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
	client := &DataConnectorsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the data connector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// dataConnectorID - Connector ID
// dataConnector - The data connector
// options - DataConnectorsClientCreateOrUpdateOptions contains the optional parameters for the DataConnectorsClient.CreateOrUpdate
// method.
func (client *DataConnectorsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, dataConnector DataConnectorClassification, options *DataConnectorsClientCreateOrUpdateOptions) (DataConnectorsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, dataConnectorID, dataConnector, options)
	if err != nil {
		return DataConnectorsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataConnectorsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DataConnectorsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DataConnectorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, dataConnector DataConnectorClassification, options *DataConnectorsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectors/{dataConnectorId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataConnectorID == "" {
		return nil, errors.New("parameter dataConnectorID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectorId}", url.PathEscape(dataConnectorID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, dataConnector)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DataConnectorsClient) createOrUpdateHandleResponse(resp *http.Response) (DataConnectorsClientCreateOrUpdateResponse, error) {
	result := DataConnectorsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return DataConnectorsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the data connector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// dataConnectorID - Connector ID
// options - DataConnectorsClientDeleteOptions contains the optional parameters for the DataConnectorsClient.Delete method.
func (client *DataConnectorsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, options *DataConnectorsClientDeleteOptions) (DataConnectorsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, dataConnectorID, options)
	if err != nil {
		return DataConnectorsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataConnectorsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DataConnectorsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DataConnectorsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataConnectorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, options *DataConnectorsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectors/{dataConnectorId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataConnectorID == "" {
		return nil, errors.New("parameter dataConnectorID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectorId}", url.PathEscape(dataConnectorID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a data connector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// dataConnectorID - Connector ID
// options - DataConnectorsClientGetOptions contains the optional parameters for the DataConnectorsClient.Get method.
func (client *DataConnectorsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, options *DataConnectorsClientGetOptions) (DataConnectorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, dataConnectorID, options)
	if err != nil {
		return DataConnectorsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataConnectorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataConnectorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataConnectorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, options *DataConnectorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectors/{dataConnectorId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataConnectorID == "" {
		return nil, errors.New("parameter dataConnectorID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectorId}", url.PathEscape(dataConnectorID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataConnectorsClient) getHandleResponse(resp *http.Response) (DataConnectorsClientGetResponse, error) {
	result := DataConnectorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return DataConnectorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all data connectors.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - DataConnectorsClientListOptions contains the optional parameters for the DataConnectorsClient.List method.
func (client *DataConnectorsClient) NewListPager(resourceGroupName string, workspaceName string, options *DataConnectorsClientListOptions) *runtime.Pager[DataConnectorsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataConnectorsClientListResponse]{
		More: func(page DataConnectorsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataConnectorsClientListResponse) (DataConnectorsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataConnectorsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DataConnectorsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataConnectorsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DataConnectorsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *DataConnectorsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DataConnectorsClient) listHandleResponse(resp *http.Response) (DataConnectorsClientListResponse, error) {
	result := DataConnectorsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataConnectorList); err != nil {
		return DataConnectorsClientListResponse{}, err
	}
	return result, nil
}
