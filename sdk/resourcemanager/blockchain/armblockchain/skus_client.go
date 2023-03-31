//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armblockchain

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

// SKUsClient contains the methods for the SKUs group.
// Don't use this type directly, use NewSKUsClient() instead.
type SKUsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSKUsClient creates a new instance of SKUsClient with the specified values.
//   - subscriptionID - Gets the subscription Id which uniquely identifies the Microsoft Azure subscription. The subscription
//     ID is part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SKUsClient, error) {
	cl, err := arm.NewClient(moduleName+".SKUsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SKUsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// List - Lists the Skus of the resource type.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01-preview
//   - options - SKUsClientListOptions contains the optional parameters for the SKUsClient.List method.
func (client *SKUsClient) List(ctx context.Context, options *SKUsClientListOptions) (SKUsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return SKUsClientListResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SKUsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SKUsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *SKUsClient) listCreateRequest(ctx context.Context, options *SKUsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Blockchain/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SKUsClient) listHandleResponse(resp *http.Response) (SKUsClientListResponse, error) {
	result := SKUsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceTypeSKUCollection); err != nil {
		return SKUsClientListResponse{}, err
	}
	return result, nil
}