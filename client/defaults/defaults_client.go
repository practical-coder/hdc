// Code generated by go-swagger; DO NOT EDIT.

package defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new defaults API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for defaults API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDefaultsSection(params *CreateDefaultsSectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDefaultsSectionCreated, *CreateDefaultsSectionAccepted, error)

	DeleteDefaultsSection(params *DeleteDefaultsSectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDefaultsSectionAccepted, *DeleteDefaultsSectionNoContent, error)

	GetDefaults(params *GetDefaultsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDefaultsOK, error)

	GetDefaultsSection(params *GetDefaultsSectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDefaultsSectionOK, error)

	GetDefaultsSections(params *GetDefaultsSectionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDefaultsSectionsOK, error)

	ReplaceDefaults(params *ReplaceDefaultsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceDefaultsOK, *ReplaceDefaultsAccepted, error)

	ReplaceDefaultsSection(params *ReplaceDefaultsSectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceDefaultsSectionOK, *ReplaceDefaultsSectionAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateDefaultsSection adds a defaults section

Adds a new defaults section to the configuration file.
*/
func (a *Client) CreateDefaultsSection(params *CreateDefaultsSectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDefaultsSectionCreated, *CreateDefaultsSectionAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDefaultsSectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createDefaultsSection",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/named_defaults",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateDefaultsSectionReader{formats: a.formats},
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
	case *CreateDefaultsSectionCreated:
		return value, nil, nil
	case *CreateDefaultsSectionAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateDefaultsSectionDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteDefaultsSection deletes a defaults section

Deletes a defaults section from the configuration by it's name.
*/
func (a *Client) DeleteDefaultsSection(params *DeleteDefaultsSectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDefaultsSectionAccepted, *DeleteDefaultsSectionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDefaultsSectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteDefaultsSection",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/named_defaults/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteDefaultsSectionReader{formats: a.formats},
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
	case *DeleteDefaultsSectionAccepted:
		return value, nil, nil
	case *DeleteDefaultsSectionNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteDefaultsSectionDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetDefaults returns defaults part of configuration

Returns defaults part of configuration, this has been deprecated, use named_defaults instead.
*/
func (a *Client) GetDefaults(params *GetDefaultsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDefaultsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDefaultsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDefaults",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/defaults",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDefaultsReader{formats: a.formats},
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
	success, ok := result.(*GetDefaultsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDefaultsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetDefaultsSection returns a defaults section

Returns one defaults section configuration by it's name.
*/
func (a *Client) GetDefaultsSection(params *GetDefaultsSectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDefaultsSectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDefaultsSectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDefaultsSection",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/named_defaults/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDefaultsSectionReader{formats: a.formats},
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
	success, ok := result.(*GetDefaultsSectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDefaultsSectionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetDefaultsSections returns an array of defaults

Returns an array of all configured defaults.
*/
func (a *Client) GetDefaultsSections(params *GetDefaultsSectionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDefaultsSectionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDefaultsSectionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDefaultsSections",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/named_defaults",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDefaultsSectionsReader{formats: a.formats},
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
	success, ok := result.(*GetDefaultsSectionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDefaultsSectionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReplaceDefaults replaces defaults

Replace defaults part of config, this has been deprecated, use named_defaults instead.
*/
func (a *Client) ReplaceDefaults(params *ReplaceDefaultsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceDefaultsOK, *ReplaceDefaultsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceDefaultsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceDefaults",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/defaults",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReplaceDefaultsReader{formats: a.formats},
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
	case *ReplaceDefaultsOK:
		return value, nil, nil
	case *ReplaceDefaultsAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceDefaultsDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReplaceDefaultsSection replaces a defatults section

Replaces a defatults section configuration by it's name.
*/
func (a *Client) ReplaceDefaultsSection(params *ReplaceDefaultsSectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceDefaultsSectionOK, *ReplaceDefaultsSectionAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceDefaultsSectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceDefaultsSection",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/named_defaults/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReplaceDefaultsSectionReader{formats: a.formats},
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
	case *ReplaceDefaultsSectionOK:
		return value, nil, nil
	case *ReplaceDefaultsSectionAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceDefaultsSectionDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
