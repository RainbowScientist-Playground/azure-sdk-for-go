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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// TicketCommunicationsNoSubscriptionServer is a fake server for instances of the armsupport.TicketCommunicationsNoSubscriptionClient type.
type TicketCommunicationsNoSubscriptionServer struct {
	// NewListPager is the fake for method TicketCommunicationsNoSubscriptionClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(supportTicketName string, options *armsupport.TicketCommunicationsNoSubscriptionClientListOptions) (resp azfake.PagerResponder[armsupport.TicketCommunicationsNoSubscriptionClientListResponse])
}

// NewTicketCommunicationsNoSubscriptionServerTransport creates a new instance of TicketCommunicationsNoSubscriptionServerTransport with the provided implementation.
// The returned TicketCommunicationsNoSubscriptionServerTransport instance is connected to an instance of armsupport.TicketCommunicationsNoSubscriptionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTicketCommunicationsNoSubscriptionServerTransport(srv *TicketCommunicationsNoSubscriptionServer) *TicketCommunicationsNoSubscriptionServerTransport {
	return &TicketCommunicationsNoSubscriptionServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsupport.TicketCommunicationsNoSubscriptionClientListResponse]](),
	}
}

// TicketCommunicationsNoSubscriptionServerTransport connects instances of armsupport.TicketCommunicationsNoSubscriptionClient to instances of TicketCommunicationsNoSubscriptionServer.
// Don't use this type directly, use NewTicketCommunicationsNoSubscriptionServerTransport instead.
type TicketCommunicationsNoSubscriptionServerTransport struct {
	srv          *TicketCommunicationsNoSubscriptionServer
	newListPager *tracker[azfake.PagerResponder[armsupport.TicketCommunicationsNoSubscriptionClientListResponse]]
}

// Do implements the policy.Transporter interface for TicketCommunicationsNoSubscriptionServerTransport.
func (t *TicketCommunicationsNoSubscriptionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TicketCommunicationsNoSubscriptionClient.NewListPager":
		resp, err = t.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TicketCommunicationsNoSubscriptionServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := t.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/providers/Microsoft\.Support/supportTickets/(?P<supportTicketName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/communications`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		supportTicketNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("supportTicketName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armsupport.TicketCommunicationsNoSubscriptionClientListOptions
		if topParam != nil || filterParam != nil {
			options = &armsupport.TicketCommunicationsNoSubscriptionClientListOptions{
				Top:    topParam,
				Filter: filterParam,
			}
		}
		resp := t.srv.NewListPager(supportTicketNameParam, options)
		newListPager = &resp
		t.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsupport.TicketCommunicationsNoSubscriptionClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		t.newListPager.remove(req)
	}
	return resp, nil
}