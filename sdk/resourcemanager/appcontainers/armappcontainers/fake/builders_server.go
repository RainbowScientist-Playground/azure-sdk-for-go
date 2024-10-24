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
	"net/http"
	"net/url"
	"regexp"

	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// BuildersServer is a fake server for instances of the armappcontainers.BuildersClient type.
type BuildersServer struct {
	// BeginCreateOrUpdate is the fake for method BuildersClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, builderName string, builderEnvelope armappcontainers.BuilderResource, options *armappcontainers.BuildersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armappcontainers.BuildersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method BuildersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, builderName string, options *armappcontainers.BuildersClientBeginDeleteOptions) (resp azfake.PollerResponder[armappcontainers.BuildersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method BuildersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, builderName string, options *armappcontainers.BuildersClientGetOptions) (resp azfake.Responder[armappcontainers.BuildersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method BuildersClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armappcontainers.BuildersClientListByResourceGroupOptions) (resp azfake.PagerResponder[armappcontainers.BuildersClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method BuildersClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armappcontainers.BuildersClientListBySubscriptionOptions) (resp azfake.PagerResponder[armappcontainers.BuildersClientListBySubscriptionResponse])

	// BeginUpdate is the fake for method BuildersClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, builderName string, builderEnvelope armappcontainers.BuilderResourceUpdate, options *armappcontainers.BuildersClientBeginUpdateOptions) (resp azfake.PollerResponder[armappcontainers.BuildersClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewBuildersServerTransport creates a new instance of BuildersServerTransport with the provided implementation.
// The returned BuildersServerTransport instance is connected to an instance of armappcontainers.BuildersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBuildersServerTransport(srv *BuildersServer) *BuildersServerTransport {
	return &BuildersServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armappcontainers.BuildersClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armappcontainers.BuildersClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armappcontainers.BuildersClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armappcontainers.BuildersClientListBySubscriptionResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armappcontainers.BuildersClientUpdateResponse]](),
	}
}

// BuildersServerTransport connects instances of armappcontainers.BuildersClient to instances of BuildersServer.
// Don't use this type directly, use NewBuildersServerTransport instead.
type BuildersServerTransport struct {
	srv                         *BuildersServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armappcontainers.BuildersClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armappcontainers.BuildersClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armappcontainers.BuildersClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armappcontainers.BuildersClientListBySubscriptionResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armappcontainers.BuildersClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for BuildersServerTransport.
func (b *BuildersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BuildersClient.BeginCreateOrUpdate":
		resp, err = b.dispatchBeginCreateOrUpdate(req)
	case "BuildersClient.BeginDelete":
		resp, err = b.dispatchBeginDelete(req)
	case "BuildersClient.Get":
		resp, err = b.dispatchGet(req)
	case "BuildersClient.NewListByResourceGroupPager":
		resp, err = b.dispatchNewListByResourceGroupPager(req)
	case "BuildersClient.NewListBySubscriptionPager":
		resp, err = b.dispatchNewListBySubscriptionPager(req)
	case "BuildersClient.BeginUpdate":
		resp, err = b.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BuildersServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := b.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/builders/(?P<builderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappcontainers.BuilderResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		builderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("builderName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, builderNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		b.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		b.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		b.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (b *BuildersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if b.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := b.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/builders/(?P<builderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		builderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("builderName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginDelete(req.Context(), resourceGroupNameParam, builderNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		b.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		b.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		b.beginDelete.remove(req)
	}

	return resp, nil
}

func (b *BuildersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/builders/(?P<builderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	builderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("builderName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), resourceGroupNameParam, builderNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BuilderResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BuildersServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := b.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/builders`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := b.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		b.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armappcontainers.BuildersClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		b.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (b *BuildersServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := b.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/builders`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := b.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		b.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armappcontainers.BuildersClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		b.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (b *BuildersServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := b.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/builders/(?P<builderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappcontainers.BuilderResourceUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		builderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("builderName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginUpdate(req.Context(), resourceGroupNameParam, builderNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		b.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		b.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		b.beginUpdate.remove(req)
	}

	return resp, nil
}
