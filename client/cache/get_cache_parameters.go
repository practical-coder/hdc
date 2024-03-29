// Code generated by go-swagger; DO NOT EDIT.

package cache

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

// NewGetCacheParams creates a new GetCacheParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCacheParams() *GetCacheParams {
	return &GetCacheParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCacheParamsWithTimeout creates a new GetCacheParams object
// with the ability to set a timeout on a request.
func NewGetCacheParamsWithTimeout(timeout time.Duration) *GetCacheParams {
	return &GetCacheParams{
		timeout: timeout,
	}
}

// NewGetCacheParamsWithContext creates a new GetCacheParams object
// with the ability to set a context for a request.
func NewGetCacheParamsWithContext(ctx context.Context) *GetCacheParams {
	return &GetCacheParams{
		Context: ctx,
	}
}

// NewGetCacheParamsWithHTTPClient creates a new GetCacheParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCacheParamsWithHTTPClient(client *http.Client) *GetCacheParams {
	return &GetCacheParams{
		HTTPClient: client,
	}
}

/*
GetCacheParams contains all the parameters to send to the API endpoint

	for the get cache operation.

	Typically these are written to a http.Request.
*/
type GetCacheParams struct {

	/* Name.

	   Cache name
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

// WithDefaults hydrates default values in the get cache params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCacheParams) WithDefaults() *GetCacheParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cache params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCacheParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cache params
func (o *GetCacheParams) WithTimeout(timeout time.Duration) *GetCacheParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cache params
func (o *GetCacheParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cache params
func (o *GetCacheParams) WithContext(ctx context.Context) *GetCacheParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cache params
func (o *GetCacheParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cache params
func (o *GetCacheParams) WithHTTPClient(client *http.Client) *GetCacheParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cache params
func (o *GetCacheParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get cache params
func (o *GetCacheParams) WithName(name string) *GetCacheParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get cache params
func (o *GetCacheParams) SetName(name string) {
	o.Name = name
}

// WithTransactionID adds the transactionID to the get cache params
func (o *GetCacheParams) WithTransactionID(transactionID *string) *GetCacheParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get cache params
func (o *GetCacheParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCacheParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
