// Code generated by go-swagger; DO NOT EDIT.

package tcp_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tcp check API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tcp check API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateTCPCheck(params *CreateTCPCheckParams, authInfo runtime.ClientAuthInfoWriter) (*CreateTCPCheckCreated, *CreateTCPCheckAccepted, error)

	DeleteTCPCheck(params *DeleteTCPCheckParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTCPCheckAccepted, *DeleteTCPCheckNoContent, error)

	GetTCPCheck(params *GetTCPCheckParams, authInfo runtime.ClientAuthInfoWriter) (*GetTCPCheckOK, error)

	GetTCPChecks(params *GetTCPChecksParams, authInfo runtime.ClientAuthInfoWriter) (*GetTCPChecksOK, error)

	ReplaceTCPCheck(params *ReplaceTCPCheckParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceTCPCheckOK, *ReplaceTCPCheckAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateTCPCheck adds a new TCP check

  Adds a new TCP check of the specified type in the specified parent.
*/
func (a *Client) CreateTCPCheck(params *CreateTCPCheckParams, authInfo runtime.ClientAuthInfoWriter) (*CreateTCPCheckCreated, *CreateTCPCheckAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTCPCheckParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createTCPCheck",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/tcp_checks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateTCPCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateTCPCheckCreated:
		return value, nil, nil
	case *CreateTCPCheckAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateTCPCheckDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteTCPCheck deletes a TCP check

  Deletes a TCP check configuration by it's index from the specified parent.
*/
func (a *Client) DeleteTCPCheck(params *DeleteTCPCheckParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTCPCheckAccepted, *DeleteTCPCheckNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTCPCheckParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteTCPCheck",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/tcp_checks/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTCPCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteTCPCheckAccepted:
		return value, nil, nil
	case *DeleteTCPCheckNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteTCPCheckDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetTCPCheck returns one TCP check

  Returns one TCP check configuration by it's index in the specified parent.
*/
func (a *Client) GetTCPCheck(params *GetTCPCheckParams, authInfo runtime.ClientAuthInfoWriter) (*GetTCPCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTCPCheckParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTCPCheck",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/tcp_checks/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTCPCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTCPCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetTCPCheckDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetTCPChecks returns an array of TCP checks

  Returns all TCP checks that are configured in specified parent.
*/
func (a *Client) GetTCPChecks(params *GetTCPChecksParams, authInfo runtime.ClientAuthInfoWriter) (*GetTCPChecksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTCPChecksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTCPChecks",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/tcp_checks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTCPChecksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTCPChecksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetTCPChecksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReplaceTCPCheck replaces a TCP check

  Replaces a TCP Check configuration by it's index in the specified parent.
*/
func (a *Client) ReplaceTCPCheck(params *ReplaceTCPCheckParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceTCPCheckOK, *ReplaceTCPCheckAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceTCPCheckParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceTCPCheck",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/tcp_checks/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceTCPCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceTCPCheckOK:
		return value, nil, nil
	case *ReplaceTCPCheckAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceTCPCheckDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
