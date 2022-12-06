// Code generated by go-swagger; DO NOT EDIT.

package userlist

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new userlist API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for userlist API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateUserlist(params *CreateUserlistParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateUserlistCreated, *CreateUserlistAccepted, error)

	DeleteUserlist(params *DeleteUserlistParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserlistAccepted, *DeleteUserlistNoContent, error)

	GetUserlist(params *GetUserlistParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserlistOK, error)

	GetUserlists(params *GetUserlistsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserlistsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateUserlist adds a new userlist

Adds a new userlist to the configuration file.
*/
func (a *Client) CreateUserlist(params *CreateUserlistParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateUserlistCreated, *CreateUserlistAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserlistParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createUserlist",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/userlists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateUserlistReader{formats: a.formats},
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
	case *CreateUserlistCreated:
		return value, nil, nil
	case *CreateUserlistAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateUserlistDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteUserlist deletes a userlist

Deletes a userlist configuration by it's name.
*/
func (a *Client) DeleteUserlist(params *DeleteUserlistParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserlistAccepted, *DeleteUserlistNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserlistParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteUserlist",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/userlists/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteUserlistReader{formats: a.formats},
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
	case *DeleteUserlistAccepted:
		return value, nil, nil
	case *DeleteUserlistNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteUserlistDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetUserlist returns one userlist

Returns one userlist configuration by it's name.
*/
func (a *Client) GetUserlist(params *GetUserlistParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserlistOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserlistParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUserlist",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/userlists/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserlistReader{formats: a.formats},
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
	success, ok := result.(*GetUserlistOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetUserlistDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetUserlists returns an array of userlists

Returns an array of all configured userlists.
*/
func (a *Client) GetUserlists(params *GetUserlistsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserlistsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserlistsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUserlists",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/userlists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserlistsReader{formats: a.formats},
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
	success, ok := result.(*GetUserlistsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetUserlistsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
