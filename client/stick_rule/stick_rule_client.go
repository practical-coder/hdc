// Code generated by go-swagger; DO NOT EDIT.

package stick_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new stick rule API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for stick rule API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateStickRule(params *CreateStickRuleParams, authInfo runtime.ClientAuthInfoWriter) (*CreateStickRuleCreated, *CreateStickRuleAccepted, error)

	DeleteStickRule(params *DeleteStickRuleParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteStickRuleAccepted, *DeleteStickRuleNoContent, error)

	GetStickRule(params *GetStickRuleParams, authInfo runtime.ClientAuthInfoWriter) (*GetStickRuleOK, error)

	GetStickRules(params *GetStickRulesParams, authInfo runtime.ClientAuthInfoWriter) (*GetStickRulesOK, error)

	ReplaceStickRule(params *ReplaceStickRuleParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceStickRuleOK, *ReplaceStickRuleAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateStickRule adds a new stick rule

  Adds a new Stick Rule of the specified type in the specified backend.
*/
func (a *Client) CreateStickRule(params *CreateStickRuleParams, authInfo runtime.ClientAuthInfoWriter) (*CreateStickRuleCreated, *CreateStickRuleAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateStickRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createStickRule",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/stick_rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateStickRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateStickRuleCreated:
		return value, nil, nil
	case *CreateStickRuleAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateStickRuleDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteStickRule deletes a stick rule

  Deletes a Stick Rule configuration by it's index from the specified backend.
*/
func (a *Client) DeleteStickRule(params *DeleteStickRuleParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteStickRuleAccepted, *DeleteStickRuleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStickRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteStickRule",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/stick_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteStickRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteStickRuleAccepted:
		return value, nil, nil
	case *DeleteStickRuleNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteStickRuleDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetStickRule returns one stick rule

  Returns one Stick Rule configuration by it's index in the specified backend.
*/
func (a *Client) GetStickRule(params *GetStickRuleParams, authInfo runtime.ClientAuthInfoWriter) (*GetStickRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStickRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStickRule",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/stick_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStickRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStickRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStickRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetStickRules returns an array of all stick rules

  Returns all Stick Rules that are configured in specified backend.
*/
func (a *Client) GetStickRules(params *GetStickRulesParams, authInfo runtime.ClientAuthInfoWriter) (*GetStickRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStickRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStickRules",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/stick_rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStickRulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStickRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStickRulesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReplaceStickRule replaces a stick rule

  Replaces a Stick Rule configuration by it's index in the specified backend.
*/
func (a *Client) ReplaceStickRule(params *ReplaceStickRuleParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceStickRuleOK, *ReplaceStickRuleAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceStickRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceStickRule",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/stick_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceStickRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceStickRuleOK:
		return value, nil, nil
	case *ReplaceStickRuleAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceStickRuleDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
