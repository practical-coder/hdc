// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetServerParams creates a new GetServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServerParams() *GetServerParams {
	return &GetServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServerParamsWithTimeout creates a new GetServerParams object
// with the ability to set a timeout on a request.
func NewGetServerParamsWithTimeout(timeout time.Duration) *GetServerParams {
	return &GetServerParams{
		timeout: timeout,
	}
}

// NewGetServerParamsWithContext creates a new GetServerParams object
// with the ability to set a context for a request.
func NewGetServerParamsWithContext(ctx context.Context) *GetServerParams {
	return &GetServerParams{
		Context: ctx,
	}
}

// NewGetServerParamsWithHTTPClient creates a new GetServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServerParamsWithHTTPClient(client *http.Client) *GetServerParams {
	return &GetServerParams{
		HTTPClient: client,
	}
}

/*
GetServerParams contains all the parameters to send to the API endpoint

	for the get server operation.

	Typically these are written to a http.Request.
*/
type GetServerParams struct {

	/* Backend.

	   Parent backend name
	*/
	Backend *string

	/* Name.

	   Server name
	*/
	Name string

	/* ParentName.

	   Parent name
	*/
	ParentName *string

	/* ParentType.

	   Parent type
	*/
	ParentType *string

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerParams) WithDefaults() *GetServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get server params
func (o *GetServerParams) WithTimeout(timeout time.Duration) *GetServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get server params
func (o *GetServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get server params
func (o *GetServerParams) WithContext(ctx context.Context) *GetServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get server params
func (o *GetServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get server params
func (o *GetServerParams) WithHTTPClient(client *http.Client) *GetServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get server params
func (o *GetServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the get server params
func (o *GetServerParams) WithBackend(backend *string) *GetServerParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the get server params
func (o *GetServerParams) SetBackend(backend *string) {
	o.Backend = backend
}

// WithName adds the name to the get server params
func (o *GetServerParams) WithName(name string) *GetServerParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get server params
func (o *GetServerParams) SetName(name string) {
	o.Name = name
}

// WithParentName adds the parentName to the get server params
func (o *GetServerParams) WithParentName(parentName *string) *GetServerParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the get server params
func (o *GetServerParams) SetParentName(parentName *string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the get server params
func (o *GetServerParams) WithParentType(parentType *string) *GetServerParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the get server params
func (o *GetServerParams) SetParentType(parentType *string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the get server params
func (o *GetServerParams) WithTransactionID(transactionID *string) *GetServerParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get server params
func (o *GetServerParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Backend != nil {

		// query param backend
		var qrBackend string

		if o.Backend != nil {
			qrBackend = *o.Backend
		}
		qBackend := qrBackend
		if qBackend != "" {

			if err := r.SetQueryParam("backend", qBackend); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.ParentName != nil {

		// query param parent_name
		var qrParentName string

		if o.ParentName != nil {
			qrParentName = *o.ParentName
		}
		qParentName := qrParentName
		if qParentName != "" {

			if err := r.SetQueryParam("parent_name", qParentName); err != nil {
				return err
			}
		}
	}

	if o.ParentType != nil {

		// query param parent_type
		var qrParentType string

		if o.ParentType != nil {
			qrParentType = *o.ParentType
		}
		qParentType := qrParentType
		if qParentType != "" {

			if err := r.SetQueryParam("parent_type", qParentType); err != nil {
				return err
			}
		}
	}

	if o.TransactionID != nil {

		// query param transaction_id
		var qrTransactionID string

		if o.TransactionID != nil {
			qrTransactionID = *o.TransactionID
		}
		qTransactionID := qrTransactionID
		if qTransactionID != "" {

			if err := r.SetQueryParam("transaction_id", qTransactionID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
