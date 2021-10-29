// Code generated by go-swagger; DO NOT EDIT.

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new maps API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for maps API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddMapEntry(params *AddMapEntryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddMapEntryCreated, error)

	ClearRuntimeMap(params *ClearRuntimeMapParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ClearRuntimeMapNoContent, error)

	DeleteRuntimeMapEntry(params *DeleteRuntimeMapEntryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRuntimeMapEntryNoContent, error)

	GetAllRuntimeMapFiles(params *GetAllRuntimeMapFilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllRuntimeMapFilesOK, error)

	GetOneRuntimeMap(params *GetOneRuntimeMapParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOneRuntimeMapOK, error)

	GetRuntimeMapEntry(params *GetRuntimeMapEntryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRuntimeMapEntryOK, error)

	ReplaceRuntimeMapEntry(params *ReplaceRuntimeMapEntryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceRuntimeMapEntryOK, error)

	ShowRuntimeMap(params *ShowRuntimeMapParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShowRuntimeMapOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddMapEntry adds an entry into the map file

  Adds an entry into the map file.
*/
func (a *Client) AddMapEntry(params *AddMapEntryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddMapEntryCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddMapEntryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addMapEntry",
		Method:             "POST",
		PathPattern:        "/services/haproxy/runtime/maps_entries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddMapEntryReader{formats: a.formats},
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
	success, ok := result.(*AddMapEntryCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddMapEntryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ClearRuntimeMap removes all map entries from the map file

  Remove all map entries from the map file.
*/
func (a *Client) ClearRuntimeMap(params *ClearRuntimeMapParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ClearRuntimeMapNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClearRuntimeMapParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "clearRuntimeMap",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/runtime/maps/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ClearRuntimeMapReader{formats: a.formats},
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
	success, ok := result.(*ClearRuntimeMapNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ClearRuntimeMapDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteRuntimeMapEntry deletes all the map entries from the map by its id

  Delete all the map entries from the map by its id.
*/
func (a *Client) DeleteRuntimeMapEntry(params *DeleteRuntimeMapEntryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRuntimeMapEntryNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRuntimeMapEntryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteRuntimeMapEntry",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/runtime/maps_entries/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRuntimeMapEntryReader{formats: a.formats},
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
	success, ok := result.(*DeleteRuntimeMapEntryNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteRuntimeMapEntryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetAllRuntimeMapFiles returns runtime map files

  Returns runtime map files.
*/
func (a *Client) GetAllRuntimeMapFiles(params *GetAllRuntimeMapFilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllRuntimeMapFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllRuntimeMapFilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllRuntimeMapFiles",
		Method:             "GET",
		PathPattern:        "/services/haproxy/runtime/maps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllRuntimeMapFilesReader{formats: a.formats},
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
	success, ok := result.(*GetAllRuntimeMapFilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAllRuntimeMapFilesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetOneRuntimeMap returns one runtime map file

  Returns one runtime map file.
*/
func (a *Client) GetOneRuntimeMap(params *GetOneRuntimeMapParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOneRuntimeMapOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOneRuntimeMapParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOneRuntimeMap",
		Method:             "GET",
		PathPattern:        "/services/haproxy/runtime/maps/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOneRuntimeMapReader{formats: a.formats},
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
	success, ok := result.(*GetOneRuntimeMapOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetOneRuntimeMapDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetRuntimeMapEntry returns one map runtime setting

  Returns one map runtime setting by it's id.
*/
func (a *Client) GetRuntimeMapEntry(params *GetRuntimeMapEntryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRuntimeMapEntryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuntimeMapEntryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRuntimeMapEntry",
		Method:             "GET",
		PathPattern:        "/services/haproxy/runtime/maps_entries/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRuntimeMapEntryReader{formats: a.formats},
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
	success, ok := result.(*GetRuntimeMapEntryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRuntimeMapEntryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReplaceRuntimeMapEntry replaces the value corresponding to each id in a map

  Replaces the value corresponding to each id in a map.
*/
func (a *Client) ReplaceRuntimeMapEntry(params *ReplaceRuntimeMapEntryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceRuntimeMapEntryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceRuntimeMapEntryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceRuntimeMapEntry",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/runtime/maps_entries/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceRuntimeMapEntryReader{formats: a.formats},
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
	success, ok := result.(*ReplaceRuntimeMapEntryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceRuntimeMapEntryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ShowRuntimeMap returns one map runtime entries

  Returns an array of all entries in a given runtime map file.
*/
func (a *Client) ShowRuntimeMap(params *ShowRuntimeMapParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShowRuntimeMapOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShowRuntimeMapParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "showRuntimeMap",
		Method:             "GET",
		PathPattern:        "/services/haproxy/runtime/maps_entries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ShowRuntimeMapReader{formats: a.formats},
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
	success, ok := result.(*ShowRuntimeMapOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ShowRuntimeMapDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
