// Code generated by go-swagger; DO NOT EDIT.

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new discovery API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for discovery API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAPIEndpoints(params *GetAPIEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPIEndpointsOK, error)

	GetConfigurationEndpoints(params *GetConfigurationEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConfigurationEndpointsOK, error)

	GetHaproxyEndpoints(params *GetHaproxyEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHaproxyEndpointsOK, error)

	GetRuntimeEndpoints(params *GetRuntimeEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRuntimeEndpointsOK, error)

	GetServicesEndpoints(params *GetServicesEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServicesEndpointsOK, error)

	GetSpoeEndpoints(params *GetSpoeEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeEndpointsOK, error)

	GetStatsEndpoints(params *GetStatsEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStatsEndpointsOK, error)

	GetStorageEndpoints(params *GetStorageEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStorageEndpointsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetAPIEndpoints returns list of root endpoints

Returns a list of root endpoints.
*/
func (a *Client) GetAPIEndpoints(params *GetAPIEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPIEndpointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIEndpointsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAPIEndpoints",
		Method:             "GET",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAPIEndpointsReader{formats: a.formats},
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
	success, ok := result.(*GetAPIEndpointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAPIEndpointsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetConfigurationEndpoints returns list of h a proxy advanced configuration endpoints

Returns a list of endpoints to be used for advanced configuration of HAProxy objects.
*/
func (a *Client) GetConfigurationEndpoints(params *GetConfigurationEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConfigurationEndpointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigurationEndpointsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getConfigurationEndpoints",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetConfigurationEndpointsReader{formats: a.formats},
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
	success, ok := result.(*GetConfigurationEndpointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetConfigurationEndpointsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetHaproxyEndpoints returns list of h a proxy related endpoints

Returns a list of HAProxy related endpoints.
*/
func (a *Client) GetHaproxyEndpoints(params *GetHaproxyEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHaproxyEndpointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHaproxyEndpointsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getHaproxyEndpoints",
		Method:             "GET",
		PathPattern:        "/services/haproxy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetHaproxyEndpointsReader{formats: a.formats},
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
	success, ok := result.(*GetHaproxyEndpointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHaproxyEndpointsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetRuntimeEndpoints returns list of h a proxy advanced runtime endpoints

Returns a list of endpoints to be used for advanced runtime settings of HAProxy objects.
*/
func (a *Client) GetRuntimeEndpoints(params *GetRuntimeEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRuntimeEndpointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuntimeEndpointsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRuntimeEndpoints",
		Method:             "GET",
		PathPattern:        "/services/haproxy/runtime",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRuntimeEndpointsReader{formats: a.formats},
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
	success, ok := result.(*GetRuntimeEndpointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRuntimeEndpointsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetServicesEndpoints returns list of service endpoints

Returns a list of API managed services endpoints.
*/
func (a *Client) GetServicesEndpoints(params *GetServicesEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServicesEndpointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServicesEndpointsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getServicesEndpoints",
		Method:             "GET",
		PathPattern:        "/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServicesEndpointsReader{formats: a.formats},
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
	success, ok := result.(*GetServicesEndpointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetServicesEndpointsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetSpoeEndpoints returns list of h a proxy s p o e endpoints

Returns a list of endpoints to be used for SPOE settings of HAProxy.
*/
func (a *Client) GetSpoeEndpoints(params *GetSpoeEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSpoeEndpointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSpoeEndpointsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSpoeEndpoints",
		Method:             "GET",
		PathPattern:        "/services/haproxy/spoe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSpoeEndpointsReader{formats: a.formats},
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
	success, ok := result.(*GetSpoeEndpointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSpoeEndpointsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetStatsEndpoints returns list of h a proxy stats endpoints

Returns a list of HAProxy stats endpoints.
*/
func (a *Client) GetStatsEndpoints(params *GetStatsEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStatsEndpointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStatsEndpointsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getStatsEndpoints",
		Method:             "GET",
		PathPattern:        "/services/haproxy/stats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStatsEndpointsReader{formats: a.formats},
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
	success, ok := result.(*GetStatsEndpointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStatsEndpointsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetStorageEndpoints returns list of h a proxy storage endpoints

Returns a list of endpoints that use HAProxy storage for persistency, e.g. maps, ssl certificates...
*/
func (a *Client) GetStorageEndpoints(params *GetStorageEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStorageEndpointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStorageEndpointsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getStorageEndpoints",
		Method:             "GET",
		PathPattern:        "/services/haproxy/storage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStorageEndpointsReader{formats: a.formats},
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
	success, ok := result.(*GetStorageEndpointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStorageEndpointsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
