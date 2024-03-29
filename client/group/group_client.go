// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateGroup(params *CreateGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateGroupCreated, *CreateGroupAccepted, error)

	DeleteGroup(params *DeleteGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteGroupAccepted, *DeleteGroupNoContent, error)

	GetGroup(params *GetGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGroupOK, error)

	GetGroups(params *GetGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGroupsOK, error)

	ReplaceGroup(params *ReplaceGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceGroupOK, *ReplaceGroupAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateGroup adds a new userlist group
*/
func (a *Client) CreateGroup(params *CreateGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateGroupCreated, *CreateGroupAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createGroup",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateGroupReader{formats: a.formats},
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
	case *CreateGroupCreated:
		return value, nil, nil
	case *CreateGroupAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateGroupDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteGroup deletes a group
*/
func (a *Client) DeleteGroup(params *DeleteGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteGroupAccepted, *DeleteGroupNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteGroup",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/groups/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteGroupReader{formats: a.formats},
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
	case *DeleteGroupAccepted:
		return value, nil, nil
	case *DeleteGroupNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteGroupDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetGroup returns one userlist group
*/
func (a *Client) GetGroup(params *GetGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGroup",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/groups/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGroupReader{formats: a.formats},
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
	success, ok := result.(*GetGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetGroups returns an array of userlist groups
*/
func (a *Client) GetGroups(params *GetGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGroups",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGroupsReader{formats: a.formats},
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
	success, ok := result.(*GetGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReplaceGroup replaces a group
*/
func (a *Client) ReplaceGroup(params *ReplaceGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceGroupOK, *ReplaceGroupAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceGroup",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/groups/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReplaceGroupReader{formats: a.formats},
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
	case *ReplaceGroupOK:
		return value, nil, nil
	case *ReplaceGroupAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceGroupDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
