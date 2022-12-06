// Code generated by go-swagger; DO NOT EDIT.

package resolver

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

// NewGetResolverParams creates a new GetResolverParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResolverParams() *GetResolverParams {
	return &GetResolverParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResolverParamsWithTimeout creates a new GetResolverParams object
// with the ability to set a timeout on a request.
func NewGetResolverParamsWithTimeout(timeout time.Duration) *GetResolverParams {
	return &GetResolverParams{
		timeout: timeout,
	}
}

// NewGetResolverParamsWithContext creates a new GetResolverParams object
// with the ability to set a context for a request.
func NewGetResolverParamsWithContext(ctx context.Context) *GetResolverParams {
	return &GetResolverParams{
		Context: ctx,
	}
}

// NewGetResolverParamsWithHTTPClient creates a new GetResolverParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetResolverParamsWithHTTPClient(client *http.Client) *GetResolverParams {
	return &GetResolverParams{
		HTTPClient: client,
	}
}

/*
GetResolverParams contains all the parameters to send to the API endpoint

	for the get resolver operation.

	Typically these are written to a http.Request.
*/
type GetResolverParams struct {

	/* Name.

	   Resolver name
	*/
	Name string

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get resolver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResolverParams) WithDefaults() *GetResolverParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resolver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResolverParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resolver params
func (o *GetResolverParams) WithTimeout(timeout time.Duration) *GetResolverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resolver params
func (o *GetResolverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resolver params
func (o *GetResolverParams) WithContext(ctx context.Context) *GetResolverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resolver params
func (o *GetResolverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resolver params
func (o *GetResolverParams) WithHTTPClient(client *http.Client) *GetResolverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resolver params
func (o *GetResolverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get resolver params
func (o *GetResolverParams) WithName(name string) *GetResolverParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get resolver params
func (o *GetResolverParams) SetName(name string) {
	o.Name = name
}

// WithTransactionID adds the transactionID to the get resolver params
func (o *GetResolverParams) WithTransactionID(transactionID *string) *GetResolverParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get resolver params
func (o *GetResolverParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetResolverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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
