// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new configuration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for configuration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetConfigurationVersion(params *GetConfigurationVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConfigurationVersionOK, error)

	GetHAProxyConfiguration(params *GetHAProxyConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHAProxyConfigurationOK, error)

	PostHAProxyConfiguration(params *PostHAProxyConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostHAProxyConfigurationCreated, *PostHAProxyConfigurationAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetConfigurationVersion returns a configuration version

Returns configuration version.
*/
func (a *Client) GetConfigurationVersion(params *GetConfigurationVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConfigurationVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigurationVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getConfigurationVersion",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetConfigurationVersionReader{formats: a.formats},
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
	success, ok := result.(*GetConfigurationVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetConfigurationVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetHAProxyConfiguration returns h a proxy configuration

Returns HAProxy configuration file in plain text
*/
func (a *Client) GetHAProxyConfiguration(params *GetHAProxyConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHAProxyConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHAProxyConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getHAProxyConfiguration",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/raw",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetHAProxyConfigurationReader{formats: a.formats},
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
	success, ok := result.(*GetHAProxyConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHAProxyConfigurationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostHAProxyConfiguration pushes new haproxy configuration

Push a new haproxy configuration file in plain text
*/
func (a *Client) PostHAProxyConfiguration(params *PostHAProxyConfigurationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostHAProxyConfigurationCreated, *PostHAProxyConfigurationAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHAProxyConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postHAProxyConfiguration",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/raw",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostHAProxyConfigurationReader{formats: a.formats},
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
	case *PostHAProxyConfigurationCreated:
		return value, nil, nil
	case *PostHAProxyConfigurationAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostHAProxyConfigurationDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
