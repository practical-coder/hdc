// Code generated by go-swagger; DO NOT EDIT.

package http_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new http check API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for http check API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateHTTPCheck(params *CreateHTTPCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateHTTPCheckCreated, *CreateHTTPCheckAccepted, error)

	DeleteHTTPCheck(params *DeleteHTTPCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteHTTPCheckAccepted, *DeleteHTTPCheckNoContent, error)

	GetHTTPCheck(params *GetHTTPCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHTTPCheckOK, error)

	GetHTTPChecks(params *GetHTTPChecksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHTTPChecksOK, error)

	ReplaceHTTPCheck(params *ReplaceHTTPCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceHTTPCheckOK, *ReplaceHTTPCheckAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateHTTPCheck adds a new HTTP check

Adds a new HTTP check of the specified type in the specified parent.
*/
func (a *Client) CreateHTTPCheck(params *CreateHTTPCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateHTTPCheckCreated, *CreateHTTPCheckAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateHTTPCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createHTTPCheck",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/http_checks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateHTTPCheckReader{formats: a.formats},
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
	case *CreateHTTPCheckCreated:
		return value, nil, nil
	case *CreateHTTPCheckAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateHTTPCheckDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteHTTPCheck deletes a HTTP check

Deletes a HTTP check configuration by it's index from the specified parent.
*/
func (a *Client) DeleteHTTPCheck(params *DeleteHTTPCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteHTTPCheckAccepted, *DeleteHTTPCheckNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteHTTPCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteHTTPCheck",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/http_checks/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteHTTPCheckReader{formats: a.formats},
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
	case *DeleteHTTPCheckAccepted:
		return value, nil, nil
	case *DeleteHTTPCheckNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteHTTPCheckDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetHTTPCheck returns one HTTP check

Returns one HTTP check configuration by it's index in the specified parent.
*/
func (a *Client) GetHTTPCheck(params *GetHTTPCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHTTPCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHTTPCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getHTTPCheck",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/http_checks/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHTTPCheckReader{formats: a.formats},
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
	success, ok := result.(*GetHTTPCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHTTPCheckDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetHTTPChecks returns an array of HTTP checks

Returns all HTTP checks that are configured in specified parent.
*/
func (a *Client) GetHTTPChecks(params *GetHTTPChecksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHTTPChecksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHTTPChecksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getHTTPChecks",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/http_checks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHTTPChecksReader{formats: a.formats},
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
	success, ok := result.(*GetHTTPChecksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHTTPChecksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReplaceHTTPCheck replaces a HTTP check

Replaces a HTTP Check configuration by it's index in the specified parent.
*/
func (a *Client) ReplaceHTTPCheck(params *ReplaceHTTPCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceHTTPCheckOK, *ReplaceHTTPCheckAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceHTTPCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceHTTPCheck",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/http_checks/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceHTTPCheckReader{formats: a.formats},
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
	case *ReplaceHTTPCheckOK:
		return value, nil, nil
	case *ReplaceHTTPCheckAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceHTTPCheckDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
