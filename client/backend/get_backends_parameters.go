// Code generated by go-swagger; DO NOT EDIT.

package backend

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

// NewGetBackendsParams creates a new GetBackendsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBackendsParams() *GetBackendsParams {
	return &GetBackendsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBackendsParamsWithTimeout creates a new GetBackendsParams object
// with the ability to set a timeout on a request.
func NewGetBackendsParamsWithTimeout(timeout time.Duration) *GetBackendsParams {
	return &GetBackendsParams{
		timeout: timeout,
	}
}

// NewGetBackendsParamsWithContext creates a new GetBackendsParams object
// with the ability to set a context for a request.
func NewGetBackendsParamsWithContext(ctx context.Context) *GetBackendsParams {
	return &GetBackendsParams{
		Context: ctx,
	}
}

// NewGetBackendsParamsWithHTTPClient creates a new GetBackendsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBackendsParamsWithHTTPClient(client *http.Client) *GetBackendsParams {
	return &GetBackendsParams{
		HTTPClient: client,
	}
}

/*
GetBackendsParams contains all the parameters to send to the API endpoint

	for the get backends operation.

	Typically these are written to a http.Request.
*/
type GetBackendsParams struct {

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get backends params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBackendsParams) WithDefaults() *GetBackendsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get backends params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBackendsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get backends params
func (o *GetBackendsParams) WithTimeout(timeout time.Duration) *GetBackendsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get backends params
func (o *GetBackendsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get backends params
func (o *GetBackendsParams) WithContext(ctx context.Context) *GetBackendsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get backends params
func (o *GetBackendsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get backends params
func (o *GetBackendsParams) WithHTTPClient(client *http.Client) *GetBackendsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get backends params
func (o *GetBackendsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTransactionID adds the transactionID to the get backends params
func (o *GetBackendsParams) WithTransactionID(transactionID *string) *GetBackendsParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get backends params
func (o *GetBackendsParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBackendsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
