// Code generated by go-swagger; DO NOT EDIT.

package http_errors

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

// NewGetHTTPErrorsSectionsParams creates a new GetHTTPErrorsSectionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHTTPErrorsSectionsParams() *GetHTTPErrorsSectionsParams {
	return &GetHTTPErrorsSectionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHTTPErrorsSectionsParamsWithTimeout creates a new GetHTTPErrorsSectionsParams object
// with the ability to set a timeout on a request.
func NewGetHTTPErrorsSectionsParamsWithTimeout(timeout time.Duration) *GetHTTPErrorsSectionsParams {
	return &GetHTTPErrorsSectionsParams{
		timeout: timeout,
	}
}

// NewGetHTTPErrorsSectionsParamsWithContext creates a new GetHTTPErrorsSectionsParams object
// with the ability to set a context for a request.
func NewGetHTTPErrorsSectionsParamsWithContext(ctx context.Context) *GetHTTPErrorsSectionsParams {
	return &GetHTTPErrorsSectionsParams{
		Context: ctx,
	}
}

// NewGetHTTPErrorsSectionsParamsWithHTTPClient creates a new GetHTTPErrorsSectionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHTTPErrorsSectionsParamsWithHTTPClient(client *http.Client) *GetHTTPErrorsSectionsParams {
	return &GetHTTPErrorsSectionsParams{
		HTTPClient: client,
	}
}

/*
GetHTTPErrorsSectionsParams contains all the parameters to send to the API endpoint

	for the get HTTP errors sections operation.

	Typically these are written to a http.Request.
*/
type GetHTTPErrorsSectionsParams struct {

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get HTTP errors sections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHTTPErrorsSectionsParams) WithDefaults() *GetHTTPErrorsSectionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get HTTP errors sections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHTTPErrorsSectionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get HTTP errors sections params
func (o *GetHTTPErrorsSectionsParams) WithTimeout(timeout time.Duration) *GetHTTPErrorsSectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get HTTP errors sections params
func (o *GetHTTPErrorsSectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get HTTP errors sections params
func (o *GetHTTPErrorsSectionsParams) WithContext(ctx context.Context) *GetHTTPErrorsSectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get HTTP errors sections params
func (o *GetHTTPErrorsSectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get HTTP errors sections params
func (o *GetHTTPErrorsSectionsParams) WithHTTPClient(client *http.Client) *GetHTTPErrorsSectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get HTTP errors sections params
func (o *GetHTTPErrorsSectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTransactionID adds the transactionID to the get HTTP errors sections params
func (o *GetHTTPErrorsSectionsParams) WithTransactionID(transactionID *string) *GetHTTPErrorsSectionsParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get HTTP errors sections params
func (o *GetHTTPErrorsSectionsParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetHTTPErrorsSectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
