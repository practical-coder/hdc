// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cluster API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cluster API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteCluster(params *DeleteClusterParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteClusterNoContent, error)

	EditCluster(params *EditClusterParams, authInfo runtime.ClientAuthInfoWriter) (*EditClusterOK, error)

	GetCluster(params *GetClusterParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterOK, error)

	InitiateCertificateRefresh(params *InitiateCertificateRefreshParams, authInfo runtime.ClientAuthInfoWriter) (*InitiateCertificateRefreshOK, error)

	PostCluster(params *PostClusterParams, authInfo runtime.ClientAuthInfoWriter) (*PostClusterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteCluster deletes cluster settings

  Delete cluster settings and move the node back to single mode
*/
func (a *Client) DeleteCluster(params *DeleteClusterParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteClusterNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCluster",
		Method:             "DELETE",
		PathPattern:        "/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteClusterNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EditCluster edits cluster settings

  Edit cluster settings
*/
func (a *Client) EditCluster(params *EditClusterParams, authInfo runtime.ClientAuthInfoWriter) (*EditClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "editCluster",
		Method:             "PUT",
		PathPattern:        "/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EditClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EditClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EditClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCluster returns cluster data

  Returns cluster data
*/
func (a *Client) GetCluster(params *GetClusterParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCluster",
		Method:             "GET",
		PathPattern:        "/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  InitiateCertificateRefresh initiates a certificate refresh

  Initiates a certificate refresh
*/
func (a *Client) InitiateCertificateRefresh(params *InitiateCertificateRefreshParams, authInfo runtime.ClientAuthInfoWriter) (*InitiateCertificateRefreshOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitiateCertificateRefreshParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "initiateCertificateRefresh",
		Method:             "POST",
		PathPattern:        "/cluster/certificate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InitiateCertificateRefreshReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InitiateCertificateRefreshOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*InitiateCertificateRefreshDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostCluster posts cluster settings

  Post cluster settings
*/
func (a *Client) PostCluster(params *PostClusterParams, authInfo runtime.ClientAuthInfoWriter) (*PostClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postCluster",
		Method:             "POST",
		PathPattern:        "/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
