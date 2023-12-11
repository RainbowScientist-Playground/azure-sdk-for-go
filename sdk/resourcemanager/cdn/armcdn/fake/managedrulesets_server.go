//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
	"net/http"
	"regexp"
)

// ManagedRuleSetsServer is a fake server for instances of the armcdn.ManagedRuleSetsClient type.
type ManagedRuleSetsServer struct {
	// NewListPager is the fake for method ManagedRuleSetsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armcdn.ManagedRuleSetsClientListOptions) (resp azfake.PagerResponder[armcdn.ManagedRuleSetsClientListResponse])
}

// NewManagedRuleSetsServerTransport creates a new instance of ManagedRuleSetsServerTransport with the provided implementation.
// The returned ManagedRuleSetsServerTransport instance is connected to an instance of armcdn.ManagedRuleSetsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedRuleSetsServerTransport(srv *ManagedRuleSetsServer) *ManagedRuleSetsServerTransport {
	return &ManagedRuleSetsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armcdn.ManagedRuleSetsClientListResponse]](),
	}
}

// ManagedRuleSetsServerTransport connects instances of armcdn.ManagedRuleSetsClient to instances of ManagedRuleSetsServer.
// Don't use this type directly, use NewManagedRuleSetsServerTransport instead.
type ManagedRuleSetsServerTransport struct {
	srv          *ManagedRuleSetsServer
	newListPager *tracker[azfake.PagerResponder[armcdn.ManagedRuleSetsClientListResponse]]
}

// Do implements the policy.Transporter interface for ManagedRuleSetsServerTransport.
func (m *ManagedRuleSetsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedRuleSetsClient.NewListPager":
		resp, err = m.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedRuleSetsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := m.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/cdnWebApplicationFirewallManagedRuleSets`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := m.srv.NewListPager(nil)
		newListPager = &resp
		m.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcdn.ManagedRuleSetsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		m.newListPager.remove(req)
	}
	return resp, nil
}