// Code generated by go-swagger; DO NOT EDIT.

package f_c_g_i_app

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

// NewGetFCGIAppsParams creates a new GetFCGIAppsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFCGIAppsParams() *GetFCGIAppsParams {
	return &GetFCGIAppsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFCGIAppsParamsWithTimeout creates a new GetFCGIAppsParams object
// with the ability to set a timeout on a request.
func NewGetFCGIAppsParamsWithTimeout(timeout time.Duration) *GetFCGIAppsParams {
	return &GetFCGIAppsParams{
		timeout: timeout,
	}
}

// NewGetFCGIAppsParamsWithContext creates a new GetFCGIAppsParams object
// with the ability to set a context for a request.
func NewGetFCGIAppsParamsWithContext(ctx context.Context) *GetFCGIAppsParams {
	return &GetFCGIAppsParams{
		Context: ctx,
	}
}

// NewGetFCGIAppsParamsWithHTTPClient creates a new GetFCGIAppsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFCGIAppsParamsWithHTTPClient(client *http.Client) *GetFCGIAppsParams {
	return &GetFCGIAppsParams{
		HTTPClient: client,
	}
}

/*
GetFCGIAppsParams contains all the parameters to send to the API endpoint

	for the get f c g i apps operation.

	Typically these are written to a http.Request.
*/
type GetFCGIAppsParams struct {

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get f c g i apps params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFCGIAppsParams) WithDefaults() *GetFCGIAppsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get f c g i apps params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFCGIAppsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get f c g i apps params
func (o *GetFCGIAppsParams) WithTimeout(timeout time.Duration) *GetFCGIAppsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get f c g i apps params
func (o *GetFCGIAppsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get f c g i apps params
func (o *GetFCGIAppsParams) WithContext(ctx context.Context) *GetFCGIAppsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get f c g i apps params
func (o *GetFCGIAppsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get f c g i apps params
func (o *GetFCGIAppsParams) WithHTTPClient(client *http.Client) *GetFCGIAppsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get f c g i apps params
func (o *GetFCGIAppsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTransactionID adds the transactionID to the get f c g i apps params
func (o *GetFCGIAppsParams) WithTransactionID(transactionID *string) *GetFCGIAppsParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get f c g i apps params
func (o *GetFCGIAppsParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFCGIAppsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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