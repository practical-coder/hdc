// Code generated by go-swagger; DO NOT EDIT.

package log_forward

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new log forward API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for log forward API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateLogForward(params *CreateLogForwardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateLogForwardCreated, *CreateLogForwardAccepted, error)

	DeleteLogForward(params *DeleteLogForwardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteLogForwardAccepted, *DeleteLogForwardNoContent, error)

	GetLogForward(params *GetLogForwardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLogForwardOK, error)

	GetLogForwards(params *GetLogForwardsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLogForwardsOK, error)

	ReplaceLogForward(params *ReplaceLogForwardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceLogForwardOK, *ReplaceLogForwardAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateLogForward adds a log forward

Adds a new log_forward to the configuration file.
*/
func (a *Client) CreateLogForward(params *CreateLogForwardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateLogForwardCreated, *CreateLogForwardAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateLogForwardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createLogForward",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/log_forwards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateLogForwardReader{formats: a.formats},
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
	case *CreateLogForwardCreated:
		return value, nil, nil
	case *CreateLogForwardAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateLogForwardDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteLogForward deletes a log forward

Deletes a log forward from the configuration by it's name.
*/
func (a *Client) DeleteLogForward(params *DeleteLogForwardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteLogForwardAccepted, *DeleteLogForwardNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLogForwardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteLogForward",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/log_forwards/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteLogForwardReader{formats: a.formats},
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
	case *DeleteLogForwardAccepted:
		return value, nil, nil
	case *DeleteLogForwardNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteLogForwardDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetLogForward returns a log forward

Returns one log forward configuration by it's name.
*/
func (a *Client) GetLogForward(params *GetLogForwardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLogForwardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLogForwardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLogForward",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/log_forwards/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLogForwardReader{formats: a.formats},
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
	success, ok := result.(*GetLogForwardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLogForwardDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetLogForwards returns an array of log forwards

Returns an array of all configured log forwards.
*/
func (a *Client) GetLogForwards(params *GetLogForwardsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLogForwardsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLogForwardsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLogForwards",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/log_forwards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLogForwardsReader{formats: a.formats},
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
	success, ok := result.(*GetLogForwardsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLogForwardsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReplaceLogForward replaces a log forward

Replaces a log forward configuration by it's name.
*/
func (a *Client) ReplaceLogForward(params *ReplaceLogForwardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceLogForwardOK, *ReplaceLogForwardAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceLogForwardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceLogForward",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/log_forwards/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReplaceLogForwardReader{formats: a.formats},
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
	case *ReplaceLogForwardOK:
		return value, nil, nil
	case *ReplaceLogForwardAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceLogForwardDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
