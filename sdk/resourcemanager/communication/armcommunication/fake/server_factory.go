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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armcommunication.ClientFactory type.
type ServerFactory struct {
	DomainsServer                  DomainsServer
	EmailServicesServer            EmailServicesServer
	OperationsServer               OperationsServer
	SenderUsernamesServer          SenderUsernamesServer
	ServicesServer                 ServicesServer
	SuppressionListAddressesServer SuppressionListAddressesServer
	SuppressionListsServer         SuppressionListsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armcommunication.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armcommunication.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                              *ServerFactory
	trMu                             sync.Mutex
	trDomainsServer                  *DomainsServerTransport
	trEmailServicesServer            *EmailServicesServerTransport
	trOperationsServer               *OperationsServerTransport
	trSenderUsernamesServer          *SenderUsernamesServerTransport
	trServicesServer                 *ServicesServerTransport
	trSuppressionListAddressesServer *SuppressionListAddressesServerTransport
	trSuppressionListsServer         *SuppressionListsServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "DomainsClient":
		initServer(s, &s.trDomainsServer, func() *DomainsServerTransport { return NewDomainsServerTransport(&s.srv.DomainsServer) })
		resp, err = s.trDomainsServer.Do(req)
	case "EmailServicesClient":
		initServer(s, &s.trEmailServicesServer, func() *EmailServicesServerTransport {
			return NewEmailServicesServerTransport(&s.srv.EmailServicesServer)
		})
		resp, err = s.trEmailServicesServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "SenderUsernamesClient":
		initServer(s, &s.trSenderUsernamesServer, func() *SenderUsernamesServerTransport {
			return NewSenderUsernamesServerTransport(&s.srv.SenderUsernamesServer)
		})
		resp, err = s.trSenderUsernamesServer.Do(req)
	case "ServicesClient":
		initServer(s, &s.trServicesServer, func() *ServicesServerTransport { return NewServicesServerTransport(&s.srv.ServicesServer) })
		resp, err = s.trServicesServer.Do(req)
	case "SuppressionListAddressesClient":
		initServer(s, &s.trSuppressionListAddressesServer, func() *SuppressionListAddressesServerTransport {
			return NewSuppressionListAddressesServerTransport(&s.srv.SuppressionListAddressesServer)
		})
		resp, err = s.trSuppressionListAddressesServer.Do(req)
	case "SuppressionListsClient":
		initServer(s, &s.trSuppressionListsServer, func() *SuppressionListsServerTransport {
			return NewSuppressionListsServerTransport(&s.srv.SuppressionListsServer)
		})
		resp, err = s.trSuppressionListsServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}