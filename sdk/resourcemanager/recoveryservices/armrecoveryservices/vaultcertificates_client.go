//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservices

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

// VaultCertificatesClient contains the methods for the VaultCertificates group.
// Don't use this type directly, use NewVaultCertificatesClient() instead.
type VaultCertificatesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVaultCertificatesClient creates a new instance of VaultCertificatesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVaultCertificatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VaultCertificatesClient, error) {
	cl, err := arm.NewClient(moduleName+".VaultCertificatesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VaultCertificatesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Uploads a certificate for a resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the recovery services vault.
//   - certificateName - Certificate friendly name.
//   - certificateRequest - Input parameters for uploading the vault certificate.
//   - options - VaultCertificatesClientCreateOptions contains the optional parameters for the VaultCertificatesClient.Create
//     method.
func (client *VaultCertificatesClient) Create(ctx context.Context, resourceGroupName string, vaultName string, certificateName string, certificateRequest CertificateRequest, options *VaultCertificatesClientCreateOptions) (VaultCertificatesClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, vaultName, certificateName, certificateRequest, options)
	if err != nil {
		return VaultCertificatesClientCreateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VaultCertificatesClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VaultCertificatesClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *VaultCertificatesClient) createCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, certificateName string, certificateRequest CertificateRequest, options *VaultCertificatesClientCreateOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, certificateRequest)
}

// createHandleResponse handles the Create response.
func (client *VaultCertificatesClient) createHandleResponse(resp *http.Response) (VaultCertificatesClientCreateResponse, error) {
	result := VaultCertificatesClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VaultCertificateResponse); err != nil {
		return VaultCertificatesClientCreateResponse{}, err
	}
	return result, nil
}
