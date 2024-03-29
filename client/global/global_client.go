// Code generated by go-swagger; DO NOT EDIT.

package global

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new global API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for global API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetGlobal(params *GetGlobalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGlobalOK, error)

	ReplaceGlobal(params *ReplaceGlobalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceGlobalOK, *ReplaceGlobalAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetGlobal returns a global part of configuration

Returns global part of configuration.
*/
func (a *Client) GetGlobal(params *GetGlobalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGlobalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGlobalParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGlobal",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/global",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGlobalReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGlobalOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetGlobalDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReplaceGlobal replaces global

Replace global part of config
*/
func (a *Client) ReplaceGlobal(params *ReplaceGlobalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceGlobalOK, *ReplaceGlobalAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceGlobalParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceGlobal",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/global",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReplaceGlobalReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceGlobalOK:
		return value, nil, nil
	case *ReplaceGlobalAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceGlobalDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
