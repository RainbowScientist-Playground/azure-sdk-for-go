//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ManagedDatabaseSensitivityLabelsServer is a fake server for instances of the armsql.ManagedDatabaseSensitivityLabelsClient type.
type ManagedDatabaseSensitivityLabelsServer struct {
	// CreateOrUpdate is the fake for method ManagedDatabaseSensitivityLabelsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, parameters armsql.SensitivityLabel, options *armsql.ManagedDatabaseSensitivityLabelsClientCreateOrUpdateOptions) (resp azfake.Responder[armsql.ManagedDatabaseSensitivityLabelsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ManagedDatabaseSensitivityLabelsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK
	Delete func(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, options *armsql.ManagedDatabaseSensitivityLabelsClientDeleteOptions) (resp azfake.Responder[armsql.ManagedDatabaseSensitivityLabelsClientDeleteResponse], errResp azfake.ErrorResponder)

	// DisableRecommendation is the fake for method ManagedDatabaseSensitivityLabelsClient.DisableRecommendation
	// HTTP status codes to indicate success: http.StatusOK
	DisableRecommendation func(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, options *armsql.ManagedDatabaseSensitivityLabelsClientDisableRecommendationOptions) (resp azfake.Responder[armsql.ManagedDatabaseSensitivityLabelsClientDisableRecommendationResponse], errResp azfake.ErrorResponder)

	// EnableRecommendation is the fake for method ManagedDatabaseSensitivityLabelsClient.EnableRecommendation
	// HTTP status codes to indicate success: http.StatusOK
	EnableRecommendation func(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, options *armsql.ManagedDatabaseSensitivityLabelsClientEnableRecommendationOptions) (resp azfake.Responder[armsql.ManagedDatabaseSensitivityLabelsClientEnableRecommendationResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ManagedDatabaseSensitivityLabelsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, sensitivityLabelSource armsql.SensitivityLabelSource, options *armsql.ManagedDatabaseSensitivityLabelsClientGetOptions) (resp azfake.Responder[armsql.ManagedDatabaseSensitivityLabelsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListCurrentByDatabasePager is the fake for method ManagedDatabaseSensitivityLabelsClient.NewListCurrentByDatabasePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListCurrentByDatabasePager func(resourceGroupName string, managedInstanceName string, databaseName string, options *armsql.ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseOptions) (resp azfake.PagerResponder[armsql.ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseResponse])

	// NewListRecommendedByDatabasePager is the fake for method ManagedDatabaseSensitivityLabelsClient.NewListRecommendedByDatabasePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListRecommendedByDatabasePager func(resourceGroupName string, managedInstanceName string, databaseName string, options *armsql.ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseOptions) (resp azfake.PagerResponder[armsql.ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseResponse])

	// Update is the fake for method ManagedDatabaseSensitivityLabelsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters armsql.SensitivityLabelUpdateList, options *armsql.ManagedDatabaseSensitivityLabelsClientUpdateOptions) (resp azfake.Responder[armsql.ManagedDatabaseSensitivityLabelsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewManagedDatabaseSensitivityLabelsServerTransport creates a new instance of ManagedDatabaseSensitivityLabelsServerTransport with the provided implementation.
// The returned ManagedDatabaseSensitivityLabelsServerTransport instance is connected to an instance of armsql.ManagedDatabaseSensitivityLabelsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedDatabaseSensitivityLabelsServerTransport(srv *ManagedDatabaseSensitivityLabelsServer) *ManagedDatabaseSensitivityLabelsServerTransport {
	return &ManagedDatabaseSensitivityLabelsServerTransport{
		srv:                               srv,
		newListCurrentByDatabasePager:     newTracker[azfake.PagerResponder[armsql.ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseResponse]](),
		newListRecommendedByDatabasePager: newTracker[azfake.PagerResponder[armsql.ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseResponse]](),
	}
}

// ManagedDatabaseSensitivityLabelsServerTransport connects instances of armsql.ManagedDatabaseSensitivityLabelsClient to instances of ManagedDatabaseSensitivityLabelsServer.
// Don't use this type directly, use NewManagedDatabaseSensitivityLabelsServerTransport instead.
type ManagedDatabaseSensitivityLabelsServerTransport struct {
	srv                               *ManagedDatabaseSensitivityLabelsServer
	newListCurrentByDatabasePager     *tracker[azfake.PagerResponder[armsql.ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseResponse]]
	newListRecommendedByDatabasePager *tracker[azfake.PagerResponder[armsql.ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseResponse]]
}

// Do implements the policy.Transporter interface for ManagedDatabaseSensitivityLabelsServerTransport.
func (m *ManagedDatabaseSensitivityLabelsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedDatabaseSensitivityLabelsClient.CreateOrUpdate":
		resp, err = m.dispatchCreateOrUpdate(req)
	case "ManagedDatabaseSensitivityLabelsClient.Delete":
		resp, err = m.dispatchDelete(req)
	case "ManagedDatabaseSensitivityLabelsClient.DisableRecommendation":
		resp, err = m.dispatchDisableRecommendation(req)
	case "ManagedDatabaseSensitivityLabelsClient.EnableRecommendation":
		resp, err = m.dispatchEnableRecommendation(req)
	case "ManagedDatabaseSensitivityLabelsClient.Get":
		resp, err = m.dispatchGet(req)
	case "ManagedDatabaseSensitivityLabelsClient.NewListCurrentByDatabasePager":
		resp, err = m.dispatchNewListCurrentByDatabasePager(req)
	case "ManagedDatabaseSensitivityLabelsClient.NewListRecommendedByDatabasePager":
		resp, err = m.dispatchNewListRecommendedByDatabasePager(req)
	case "ManagedDatabaseSensitivityLabelsClient.Update":
		resp, err = m.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedDatabaseSensitivityLabelsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/columns/(?P<columnName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sensitivityLabels/(?P<sensitivityLabelSource>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsql.SensitivityLabel](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	schemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaName")])
	if err != nil {
		return nil, err
	}
	tableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("tableName")])
	if err != nil {
		return nil, err
	}
	columnNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("columnName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, schemaNameParam, tableNameParam, columnNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SensitivityLabel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedDatabaseSensitivityLabelsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if m.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/columns/(?P<columnName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sensitivityLabels/(?P<sensitivityLabelSource>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	schemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaName")])
	if err != nil {
		return nil, err
	}
	tableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("tableName")])
	if err != nil {
		return nil, err
	}
	columnNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("columnName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Delete(req.Context(), resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, schemaNameParam, tableNameParam, columnNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedDatabaseSensitivityLabelsServerTransport) dispatchDisableRecommendation(req *http.Request) (*http.Response, error) {
	if m.srv.DisableRecommendation == nil {
		return nil, &nonRetriableError{errors.New("fake for method DisableRecommendation not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/columns/(?P<columnName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sensitivityLabels/(?P<sensitivityLabelSource>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disable`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	schemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaName")])
	if err != nil {
		return nil, err
	}
	tableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("tableName")])
	if err != nil {
		return nil, err
	}
	columnNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("columnName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.DisableRecommendation(req.Context(), resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, schemaNameParam, tableNameParam, columnNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedDatabaseSensitivityLabelsServerTransport) dispatchEnableRecommendation(req *http.Request) (*http.Response, error) {
	if m.srv.EnableRecommendation == nil {
		return nil, &nonRetriableError{errors.New("fake for method EnableRecommendation not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/columns/(?P<columnName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sensitivityLabels/(?P<sensitivityLabelSource>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/enable`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	schemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaName")])
	if err != nil {
		return nil, err
	}
	tableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("tableName")])
	if err != nil {
		return nil, err
	}
	columnNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("columnName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.EnableRecommendation(req.Context(), resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, schemaNameParam, tableNameParam, columnNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedDatabaseSensitivityLabelsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/columns/(?P<columnName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sensitivityLabels/(?P<sensitivityLabelSource>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 8 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	schemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaName")])
	if err != nil {
		return nil, err
	}
	tableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("tableName")])
	if err != nil {
		return nil, err
	}
	columnNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("columnName")])
	if err != nil {
		return nil, err
	}
	sensitivityLabelSourceParam, err := parseWithCast(matches[regex.SubexpIndex("sensitivityLabelSource")], func(v string) (armsql.SensitivityLabelSource, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsql.SensitivityLabelSource(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, schemaNameParam, tableNameParam, columnNameParam, sensitivityLabelSourceParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SensitivityLabel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedDatabaseSensitivityLabelsServerTransport) dispatchNewListCurrentByDatabasePager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListCurrentByDatabasePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListCurrentByDatabasePager not implemented")}
	}
	newListCurrentByDatabasePager := m.newListCurrentByDatabasePager.get(req)
	if newListCurrentByDatabasePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/currentSensitivityLabels`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		countUnescaped, err := url.QueryUnescape(qp.Get("$count"))
		if err != nil {
			return nil, err
		}
		countParam, err := parseOptional(countUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armsql.ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseOptions
		if skipTokenParam != nil || countParam != nil || filterParam != nil {
			options = &armsql.ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseOptions{
				SkipToken: skipTokenParam,
				Count:     countParam,
				Filter:    filterParam,
			}
		}
		resp := m.srv.NewListCurrentByDatabasePager(resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, options)
		newListCurrentByDatabasePager = &resp
		m.newListCurrentByDatabasePager.add(req, newListCurrentByDatabasePager)
		server.PagerResponderInjectNextLinks(newListCurrentByDatabasePager, req, func(page *armsql.ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListCurrentByDatabasePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListCurrentByDatabasePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListCurrentByDatabasePager) {
		m.newListCurrentByDatabasePager.remove(req)
	}
	return resp, nil
}

func (m *ManagedDatabaseSensitivityLabelsServerTransport) dispatchNewListRecommendedByDatabasePager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListRecommendedByDatabasePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListRecommendedByDatabasePager not implemented")}
	}
	newListRecommendedByDatabasePager := m.newListRecommendedByDatabasePager.get(req)
	if newListRecommendedByDatabasePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/recommendedSensitivityLabels`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		includeDisabledRecommendationsUnescaped, err := url.QueryUnescape(qp.Get("includeDisabledRecommendations"))
		if err != nil {
			return nil, err
		}
		includeDisabledRecommendationsParam, err := parseOptional(includeDisabledRecommendationsUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armsql.ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseOptions
		if skipTokenParam != nil || includeDisabledRecommendationsParam != nil || filterParam != nil {
			options = &armsql.ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseOptions{
				SkipToken:                      skipTokenParam,
				IncludeDisabledRecommendations: includeDisabledRecommendationsParam,
				Filter:                         filterParam,
			}
		}
		resp := m.srv.NewListRecommendedByDatabasePager(resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, options)
		newListRecommendedByDatabasePager = &resp
		m.newListRecommendedByDatabasePager.add(req, newListRecommendedByDatabasePager)
		server.PagerResponderInjectNextLinks(newListRecommendedByDatabasePager, req, func(page *armsql.ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListRecommendedByDatabasePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListRecommendedByDatabasePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListRecommendedByDatabasePager) {
		m.newListRecommendedByDatabasePager.remove(req)
	}
	return resp, nil
}

func (m *ManagedDatabaseSensitivityLabelsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/currentSensitivityLabels`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsql.SensitivityLabelUpdateList](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Update(req.Context(), resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}