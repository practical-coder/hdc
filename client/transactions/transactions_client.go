// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new transactions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for transactions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CommitTransaction(params *CommitTransactionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CommitTransactionOK, *CommitTransactionAccepted, error)

	DeleteTransaction(params *DeleteTransactionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTransactionNoContent, error)

	GetTransaction(params *GetTransactionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTransactionOK, error)

	GetTransactions(params *GetTransactionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTransactionsOK, error)

	StartTransaction(params *StartTransactionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StartTransactionCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CommitTransaction commits transaction

Commit transaction, execute all operations in transaction and return msg
*/
func (a *Client) CommitTransaction(params *CommitTransactionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CommitTransactionOK, *CommitTransactionAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCommitTransactionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "commitTransaction",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/transactions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CommitTransactionReader{formats: a.formats},
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
	case *CommitTransactionOK:
		return value, nil, nil
	case *CommitTransactionAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CommitTransactionDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteTransaction deletes a transaction

Deletes a transaction.
*/
func (a *Client) DeleteTransaction(params *DeleteTransactionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTransactionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTransactionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteTransaction",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/transactions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTransactionReader{formats: a.formats},
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
	success, ok := result.(*DeleteTransactionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteTransactionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetTransaction returns one h a proxy configuration transactions

Returns one HAProxy configuration transactions.
*/
func (a *Client) GetTransaction(params *GetTransactionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTransactionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTransactionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTransaction",
		Method:             "GET",
		PathPattern:        "/services/haproxy/transactions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTransactionReader{formats: a.formats},
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
	success, ok := result.(*GetTransactionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetTransactionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetTransactions returns list of h a proxy configuration transactions

Returns a list of HAProxy configuration transactions. Transactions can be filtered by their status.
*/
func (a *Client) GetTransactions(params *GetTransactionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTransactionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTransactionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTransactions",
		Method:             "GET",
		PathPattern:        "/services/haproxy/transactions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTransactionsReader{formats: a.formats},
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
	success, ok := result.(*GetTransactionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetTransactionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
StartTransaction starts a new transaction

Starts a new transaction and returns it's id
*/
func (a *Client) StartTransaction(params *StartTransactionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StartTransactionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartTransactionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "startTransaction",
		Method:             "POST",
		PathPattern:        "/services/haproxy/transactions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StartTransactionReader{formats: a.formats},
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
	success, ok := result.(*StartTransactionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartTransactionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
