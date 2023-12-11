//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappconfiguration

// ConfigurationStoresClientBeginCreateOptions contains the optional parameters for the ConfigurationStoresClient.BeginCreate
// method.
type ConfigurationStoresClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ConfigurationStoresClientBeginDeleteOptions contains the optional parameters for the ConfigurationStoresClient.BeginDelete
// method.
type ConfigurationStoresClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ConfigurationStoresClientBeginPurgeDeletedOptions contains the optional parameters for the ConfigurationStoresClient.BeginPurgeDeleted
// method.
type ConfigurationStoresClientBeginPurgeDeletedOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ConfigurationStoresClientBeginUpdateOptions contains the optional parameters for the ConfigurationStoresClient.BeginUpdate
// method.
type ConfigurationStoresClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ConfigurationStoresClientGetDeletedOptions contains the optional parameters for the ConfigurationStoresClient.GetDeleted
// method.
type ConfigurationStoresClientGetDeletedOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationStoresClientGetOptions contains the optional parameters for the ConfigurationStoresClient.Get method.
type ConfigurationStoresClientGetOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationStoresClientListByResourceGroupOptions contains the optional parameters for the ConfigurationStoresClient.NewListByResourceGroupPager
// method.
type ConfigurationStoresClientListByResourceGroupOptions struct {
	// A skip token is used to continue retrieving items after an operation returns a partial result. If a previous response contains
	// a nextLink element, the value of the nextLink element will include a
	// skipToken parameter that specifies a starting point to use for subsequent calls.
	SkipToken *string
}

// ConfigurationStoresClientListDeletedOptions contains the optional parameters for the ConfigurationStoresClient.NewListDeletedPager
// method.
type ConfigurationStoresClientListDeletedOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationStoresClientListKeysOptions contains the optional parameters for the ConfigurationStoresClient.NewListKeysPager
// method.
type ConfigurationStoresClientListKeysOptions struct {
	// A skip token is used to continue retrieving items after an operation returns a partial result. If a previous response contains
	// a nextLink element, the value of the nextLink element will include a
	// skipToken parameter that specifies a starting point to use for subsequent calls.
	SkipToken *string
}

// ConfigurationStoresClientListOptions contains the optional parameters for the ConfigurationStoresClient.NewListPager method.
type ConfigurationStoresClientListOptions struct {
	// A skip token is used to continue retrieving items after an operation returns a partial result. If a previous response contains
	// a nextLink element, the value of the nextLink element will include a
	// skipToken parameter that specifies a starting point to use for subsequent calls.
	SkipToken *string
}

// ConfigurationStoresClientRegenerateKeyOptions contains the optional parameters for the ConfigurationStoresClient.RegenerateKey
// method.
type ConfigurationStoresClientRegenerateKeyOptions struct {
	// placeholder for future optional parameters
}

// KeyValuesClientBeginDeleteOptions contains the optional parameters for the KeyValuesClient.BeginDelete method.
type KeyValuesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// KeyValuesClientCreateOrUpdateOptions contains the optional parameters for the KeyValuesClient.CreateOrUpdate method.
type KeyValuesClientCreateOrUpdateOptions struct {
	// The parameters for creating a key-value.
	KeyValueParameters *KeyValue
}

// KeyValuesClientGetOptions contains the optional parameters for the KeyValuesClient.Get method.
type KeyValuesClientGetOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientCheckNameAvailabilityOptions contains the optional parameters for the OperationsClient.CheckNameAvailability
// method.
type OperationsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// A skip token is used to continue retrieving items after an operation returns a partial result. If a previous response contains
	// a nextLink element, the value of the nextLink element will include a
	// skipToken parameter that specifies a starting point to use for subsequent calls.
	SkipToken *string
}

// OperationsClientRegionalCheckNameAvailabilityOptions contains the optional parameters for the OperationsClient.RegionalCheckNameAvailability
// method.
type OperationsClientRegionalCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginCreateOrUpdate
// method.
type PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
// method.
type PrivateEndpointConnectionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
type PrivateEndpointConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientListByConfigurationStoreOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListByConfigurationStorePager
// method.
type PrivateEndpointConnectionsClientListByConfigurationStoreOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientGetOptions contains the optional parameters for the PrivateLinkResourcesClient.Get method.
type PrivateLinkResourcesClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientListByConfigurationStoreOptions contains the optional parameters for the PrivateLinkResourcesClient.NewListByConfigurationStorePager
// method.
type PrivateLinkResourcesClientListByConfigurationStoreOptions struct {
	// placeholder for future optional parameters
}

// ReplicasClientBeginCreateOptions contains the optional parameters for the ReplicasClient.BeginCreate method.
type ReplicasClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ReplicasClientBeginDeleteOptions contains the optional parameters for the ReplicasClient.BeginDelete method.
type ReplicasClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ReplicasClientGetOptions contains the optional parameters for the ReplicasClient.Get method.
type ReplicasClientGetOptions struct {
	// placeholder for future optional parameters
}

// ReplicasClientListByConfigurationStoreOptions contains the optional parameters for the ReplicasClient.NewListByConfigurationStorePager
// method.
type ReplicasClientListByConfigurationStoreOptions struct {
	// A skip token is used to continue retrieving items after an operation returns a partial result. If a previous response contains
	// a nextLink element, the value of the nextLink element will include a
	// skipToken parameter that specifies a starting point to use for subsequent calls.
	SkipToken *string
}