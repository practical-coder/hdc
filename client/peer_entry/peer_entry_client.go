// Code generated by go-swagger; DO NOT EDIT.

package peer_entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new peer entry API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for peer entry API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePeerEntry(params *CreatePeerEntryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePeerEntryCreated, *CreatePeerEntryAccepted, error)

	DeletePeerEntry(params *DeletePeerEntryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePeerEntryAccepted, *DeletePeerEntryNoContent, error)

	GetPeerEntries(params *GetPeerEntriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPeerEntriesOK, error)

	GetPeerEntry(params *GetPeerEntryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPeerEntryOK, error)

	ReplacePeerEntry(params *ReplacePeerEntryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplacePeerEntryOK, *ReplacePeerEntryAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreatePeerEntry adds a new peer entry

  Adds a new peer entry in the specified peer section in the configuration file.
*/
func (a *Client) CreatePeerEntry(params *CreatePeerEntryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePeerEntryCreated, *CreatePeerEntryAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePeerEntryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPeerEntry",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/peer_entries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePeerEntryReader{formats: a.formats},
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
	case *CreatePeerEntryCreated:
		return value, nil, nil
	case *CreatePeerEntryAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreatePeerEntryDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeletePeerEntry deletes a peer entry

  Deletes a peer entry configuration by it's name in the specified peer section.
*/
func (a *Client) DeletePeerEntry(params *DeletePeerEntryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePeerEntryAccepted, *DeletePeerEntryNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePeerEntryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePeerEntry",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/peer_entries/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePeerEntryReader{formats: a.formats},
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
	case *DeletePeerEntryAccepted:
		return value, nil, nil
	case *DeletePeerEntryNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeletePeerEntryDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetPeerEntries returns an array of peer entries

  Returns an array of all peer_entries that are configured in specified peer section.
*/
func (a *Client) GetPeerEntries(params *GetPeerEntriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPeerEntriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPeerEntriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPeerEntries",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/peer_entries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPeerEntriesReader{formats: a.formats},
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
	success, ok := result.(*GetPeerEntriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPeerEntriesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetPeerEntry returns one peer entry

  Returns one peer_entry configuration by it's name in the specified peer section.
*/
func (a *Client) GetPeerEntry(params *GetPeerEntryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPeerEntryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPeerEntryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPeerEntry",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/peer_entries/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPeerEntryReader{formats: a.formats},
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
	success, ok := result.(*GetPeerEntryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPeerEntryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReplacePeerEntry replaces a peer entry

  Replaces a peer entry configuration by it's name in the specified peer section.
*/
func (a *Client) ReplacePeerEntry(params *ReplacePeerEntryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplacePeerEntryOK, *ReplacePeerEntryAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplacePeerEntryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replacePeerEntry",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/peer_entries/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplacePeerEntryReader{formats: a.formats},
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
	case *ReplacePeerEntryOK:
		return value, nil, nil
	case *ReplacePeerEntryAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplacePeerEntryDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
