// Code generated by go-swagger; DO NOT EDIT.

package http_error_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new http error rule API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for http error rule API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateHTTPErrorRule(params *CreateHTTPErrorRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateHTTPErrorRuleCreated, *CreateHTTPErrorRuleAccepted, error)

	DeleteHTTPErrorRule(params *DeleteHTTPErrorRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteHTTPErrorRuleAccepted, *DeleteHTTPErrorRuleNoContent, error)

	GetHTTPErrorRule(params *GetHTTPErrorRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHTTPErrorRuleOK, error)

	GetHTTPErrorRules(params *GetHTTPErrorRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHTTPErrorRulesOK, error)

	ReplaceHTTPErrorRule(params *ReplaceHTTPErrorRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceHTTPErrorRuleOK, *ReplaceHTTPErrorRuleAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateHTTPErrorRule adds a new HTTP error rule

Adds a new HTTP Error Rule of the specified type in the specified parent.
*/
func (a *Client) CreateHTTPErrorRule(params *CreateHTTPErrorRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateHTTPErrorRuleCreated, *CreateHTTPErrorRuleAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateHTTPErrorRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createHTTPErrorRule",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/http_error_rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateHTTPErrorRuleReader{formats: a.formats},
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
	case *CreateHTTPErrorRuleCreated:
		return value, nil, nil
	case *CreateHTTPErrorRuleAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateHTTPErrorRuleDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteHTTPErrorRule deletes a HTTP error rule

Deletes a HTTP Error Rule configuration by its index from the specified parent.
*/
func (a *Client) DeleteHTTPErrorRule(params *DeleteHTTPErrorRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteHTTPErrorRuleAccepted, *DeleteHTTPErrorRuleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteHTTPErrorRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteHTTPErrorRule",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/http_error_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteHTTPErrorRuleReader{formats: a.formats},
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
	case *DeleteHTTPErrorRuleAccepted:
		return value, nil, nil
	case *DeleteHTTPErrorRuleNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteHTTPErrorRuleDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetHTTPErrorRule returns one HTTP error rule

Returns one HTTP Error Rule configuration by its index in the specified parent.
*/
func (a *Client) GetHTTPErrorRule(params *GetHTTPErrorRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHTTPErrorRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHTTPErrorRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getHTTPErrorRule",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/http_error_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetHTTPErrorRuleReader{formats: a.formats},
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
	success, ok := result.(*GetHTTPErrorRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHTTPErrorRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetHTTPErrorRules returns an array of all HTTP error rules

Returns all HTTP Error Rules that are configured in the specified parent.
*/
func (a *Client) GetHTTPErrorRules(params *GetHTTPErrorRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHTTPErrorRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHTTPErrorRulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getHTTPErrorRules",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/http_error_rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetHTTPErrorRulesReader{formats: a.formats},
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
	success, ok := result.(*GetHTTPErrorRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHTTPErrorRulesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReplaceHTTPErrorRule replaces a HTTP error rule

Replaces a HTTP Error Rule configuration by its index in the specified parent.
*/
func (a *Client) ReplaceHTTPErrorRule(params *ReplaceHTTPErrorRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceHTTPErrorRuleOK, *ReplaceHTTPErrorRuleAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceHTTPErrorRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceHTTPErrorRule",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/http_error_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReplaceHTTPErrorRuleReader{formats: a.formats},
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
	case *ReplaceHTTPErrorRuleOK:
		return value, nil, nil
	case *ReplaceHTTPErrorRuleAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceHTTPErrorRuleDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
