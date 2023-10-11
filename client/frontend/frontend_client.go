// Code generated by go-swagger; DO NOT EDIT.

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new frontend API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for frontend API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateFrontend(params *CreateFrontendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateFrontendCreated, *CreateFrontendAccepted, error)

	DeleteFrontend(params *DeleteFrontendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFrontendAccepted, *DeleteFrontendNoContent, error)

	GetFrontend(params *GetFrontendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFrontendOK, error)

	GetFrontends(params *GetFrontendsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFrontendsOK, error)

	ReplaceFrontend(params *ReplaceFrontendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceFrontendOK, *ReplaceFrontendAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateFrontend adds a frontend

Adds a new frontend to the configuration file.
*/
func (a *Client) CreateFrontend(params *CreateFrontendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateFrontendCreated, *CreateFrontendAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFrontendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createFrontend",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/frontends",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateFrontendReader{formats: a.formats},
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
	case *CreateFrontendCreated:
		return value, nil, nil
	case *CreateFrontendAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateFrontendDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteFrontend deletes a frontend

Deletes a frontend from the configuration by it's name.
*/
func (a *Client) DeleteFrontend(params *DeleteFrontendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFrontendAccepted, *DeleteFrontendNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFrontendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteFrontend",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/frontends/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteFrontendReader{formats: a.formats},
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
	case *DeleteFrontendAccepted:
		return value, nil, nil
	case *DeleteFrontendNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteFrontendDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetFrontend returns a frontend

Returns one frontend configuration by it's name.
*/
func (a *Client) GetFrontend(params *GetFrontendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFrontendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFrontendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFrontend",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/frontends/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFrontendReader{formats: a.formats},
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
	success, ok := result.(*GetFrontendOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFrontendDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetFrontends returns an array of frontends

Returns an array of all configured frontends.
*/
func (a *Client) GetFrontends(params *GetFrontendsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFrontendsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFrontendsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFrontends",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/frontends",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFrontendsReader{formats: a.formats},
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
	success, ok := result.(*GetFrontendsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFrontendsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReplaceFrontend replaces a frontend

Replaces a frontend configuration by it's name.
*/
func (a *Client) ReplaceFrontend(params *ReplaceFrontendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceFrontendOK, *ReplaceFrontendAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceFrontendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceFrontend",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/frontends/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReplaceFrontendReader{formats: a.formats},
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
	case *ReplaceFrontendOK:
		return value, nil, nil
	case *ReplaceFrontendAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceFrontendDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
