// Code generated by go-swagger; DO NOT EDIT.

package frontend

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

// NewGetFrontendParams creates a new GetFrontendParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFrontendParams() *GetFrontendParams {
	return &GetFrontendParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFrontendParamsWithTimeout creates a new GetFrontendParams object
// with the ability to set a timeout on a request.
func NewGetFrontendParamsWithTimeout(timeout time.Duration) *GetFrontendParams {
	return &GetFrontendParams{
		timeout: timeout,
	}
}

// NewGetFrontendParamsWithContext creates a new GetFrontendParams object
// with the ability to set a context for a request.
func NewGetFrontendParamsWithContext(ctx context.Context) *GetFrontendParams {
	return &GetFrontendParams{
		Context: ctx,
	}
}

// NewGetFrontendParamsWithHTTPClient creates a new GetFrontendParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFrontendParamsWithHTTPClient(client *http.Client) *GetFrontendParams {
	return &GetFrontendParams{
		HTTPClient: client,
	}
}

/*
GetFrontendParams contains all the parameters to send to the API endpoint

	for the get frontend operation.

	Typically these are written to a http.Request.
*/
type GetFrontendParams struct {

	/* Name.

	   Frontend name
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

// WithDefaults hydrates default values in the get frontend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFrontendParams) WithDefaults() *GetFrontendParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get frontend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFrontendParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get frontend params
func (o *GetFrontendParams) WithTimeout(timeout time.Duration) *GetFrontendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get frontend params
func (o *GetFrontendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get frontend params
func (o *GetFrontendParams) WithContext(ctx context.Context) *GetFrontendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get frontend params
func (o *GetFrontendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get frontend params
func (o *GetFrontendParams) WithHTTPClient(client *http.Client) *GetFrontendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get frontend params
func (o *GetFrontendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get frontend params
func (o *GetFrontendParams) WithName(name string) *GetFrontendParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get frontend params
func (o *GetFrontendParams) SetName(name string) {
	o.Name = name
}

// WithTransactionID adds the transactionID to the get frontend params
func (o *GetFrontendParams) WithTransactionID(transactionID *string) *GetFrontendParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get frontend params
func (o *GetFrontendParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFrontendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
