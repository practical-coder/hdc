// Code generated by go-swagger; DO NOT EDIT.

package service_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service discovery API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service discovery API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAWSRegion(params *CreateAWSRegionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAWSRegionCreated, error)

	CreateConsul(params *CreateConsulParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateConsulCreated, error)

	DeleteAWSRegion(params *DeleteAWSRegionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAWSRegionNoContent, error)

	DeleteConsul(params *DeleteConsulParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteConsulNoContent, error)

	GetAWSRegion(params *GetAWSRegionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAWSRegionOK, error)

	GetAWSRegions(params *GetAWSRegionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAWSRegionsOK, error)

	GetConsul(params *GetConsulParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConsulOK, error)

	GetConsuls(params *GetConsulsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConsulsOK, error)

	ReplaceAWSRegion(params *ReplaceAWSRegionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceAWSRegionOK, error)

	ReplaceConsul(params *ReplaceConsulParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceConsulOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	CreateAWSRegion adds a new a w s region

	Add a new AWS region.

Credentials are not required in case Dataplane API is running in an EC2 instance with proper IAM role attached.
*/
func (a *Client) CreateAWSRegion(params *CreateAWSRegionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAWSRegionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAWSRegionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAWSRegion",
		Method:             "POST",
		PathPattern:        "/service_discovery/aws",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAWSRegionReader{formats: a.formats},
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
	success, ok := result.(*CreateAWSRegionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateAWSRegionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateConsul adds a new consul server

Adds a new Consul server.
*/
func (a *Client) CreateConsul(params *CreateConsulParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateConsulCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateConsulParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createConsul",
		Method:             "POST",
		PathPattern:        "/service_discovery/consul",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateConsulReader{formats: a.formats},
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
	success, ok := result.(*CreateConsulCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateConsulDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteAWSRegion deletes an a w s region

Delete an AWS region configuration by it's id.
*/
func (a *Client) DeleteAWSRegion(params *DeleteAWSRegionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAWSRegionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAWSRegionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAWSRegion",
		Method:             "DELETE",
		PathPattern:        "/service_discovery/aws/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteAWSRegionReader{formats: a.formats},
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
	success, ok := result.(*DeleteAWSRegionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteAWSRegionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteConsul deletes a consul server

Deletes a Consul server configuration by it's id.
*/
func (a *Client) DeleteConsul(params *DeleteConsulParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteConsulNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteConsulParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteConsul",
		Method:             "DELETE",
		PathPattern:        "/service_discovery/consul/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteConsulReader{formats: a.formats},
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
	success, ok := result.(*DeleteConsulNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteConsulDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetAWSRegion returns an a w s region

Return one AWS Region configuration by it's id.
*/
func (a *Client) GetAWSRegion(params *GetAWSRegionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAWSRegionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAWSRegionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAWSRegion",
		Method:             "GET",
		PathPattern:        "/service_discovery/aws/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAWSRegionReader{formats: a.formats},
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
	success, ok := result.(*GetAWSRegionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAWSRegionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetAWSRegions returns an array of all configured a w s regions

Return all configured AWS regions.
*/
func (a *Client) GetAWSRegions(params *GetAWSRegionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAWSRegionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAWSRegionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAWSRegions",
		Method:             "GET",
		PathPattern:        "/service_discovery/aws",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAWSRegionsReader{formats: a.formats},
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
	success, ok := result.(*GetAWSRegionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAWSRegionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetConsul returns one consul server

Returns one Consul server configuration by it's id.
*/
func (a *Client) GetConsul(params *GetConsulParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConsulOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConsulParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getConsul",
		Method:             "GET",
		PathPattern:        "/service_discovery/consul/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetConsulReader{formats: a.formats},
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
	success, ok := result.(*GetConsulOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetConsulDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetConsuls returns an array of all configured consul servers

Returns all configured Consul servers.
*/
func (a *Client) GetConsuls(params *GetConsulsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConsulsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConsulsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getConsuls",
		Method:             "GET",
		PathPattern:        "/service_discovery/consul",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetConsulsReader{formats: a.formats},
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
	success, ok := result.(*GetConsulsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetConsulsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReplaceAWSRegion replaces an a w s region

Replace an AWS region configuration by its id.
*/
func (a *Client) ReplaceAWSRegion(params *ReplaceAWSRegionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceAWSRegionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceAWSRegionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceAWSRegion",
		Method:             "PUT",
		PathPattern:        "/service_discovery/aws/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReplaceAWSRegionReader{formats: a.formats},
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
	success, ok := result.(*ReplaceAWSRegionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceAWSRegionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReplaceConsul replaces a consul server

Replaces a Consul server configuration by it's id.
*/
func (a *Client) ReplaceConsul(params *ReplaceConsulParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceConsulOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceConsulParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceConsul",
		Method:             "PUT",
		PathPattern:        "/service_discovery/consul/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReplaceConsulReader{formats: a.formats},
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
	success, ok := result.(*ReplaceConsulOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceConsulDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
