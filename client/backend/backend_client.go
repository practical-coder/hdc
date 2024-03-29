// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new backend API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for backend API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBackend(params *CreateBackendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBackendCreated, *CreateBackendAccepted, error)

	DeleteBackend(params *DeleteBackendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBackendAccepted, *DeleteBackendNoContent, error)

	GetBackend(params *GetBackendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackendOK, error)

	GetBackends(params *GetBackendsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackendsOK, error)

	ReplaceBackend(params *ReplaceBackendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceBackendOK, *ReplaceBackendAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateBackend adds a backend

Adds a new backend to the configuration file.
*/
func (a *Client) CreateBackend(params *CreateBackendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBackendCreated, *CreateBackendAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBackendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createBackend",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/backends",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateBackendReader{formats: a.formats},
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
	case *CreateBackendCreated:
		return value, nil, nil
	case *CreateBackendAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateBackendDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteBackend deletes a backend

Deletes a backend from the configuration by it's name.
*/
func (a *Client) DeleteBackend(params *DeleteBackendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBackendAccepted, *DeleteBackendNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBackendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteBackend",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/backends/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteBackendReader{formats: a.formats},
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
	case *DeleteBackendAccepted:
		return value, nil, nil
	case *DeleteBackendNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteBackendDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBackend returns a backend

Returns one backend configuration by it's name.
*/
func (a *Client) GetBackend(params *GetBackendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBackend",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/backends/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBackendReader{formats: a.formats},
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
	success, ok := result.(*GetBackendOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBackendDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBackends returns an array of backends

Returns an array of all configured backends.
*/
func (a *Client) GetBackends(params *GetBackendsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackendsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackendsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBackends",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/backends",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBackendsReader{formats: a.formats},
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
	success, ok := result.(*GetBackendsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBackendsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReplaceBackend replaces a backend

Replaces a backend configuration by it's name.
*/
func (a *Client) ReplaceBackend(params *ReplaceBackendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceBackendOK, *ReplaceBackendAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceBackendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceBackend",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/backends/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReplaceBackendReader{formats: a.formats},
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
	case *ReplaceBackendOK:
		return value, nil, nil
	case *ReplaceBackendAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceBackendDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
