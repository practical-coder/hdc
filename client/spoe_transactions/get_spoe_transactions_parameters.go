// Code generated by go-swagger; DO NOT EDIT.

package spoe_transactions

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

// NewGetSpoeTransactionsParams creates a new GetSpoeTransactionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSpoeTransactionsParams() *GetSpoeTransactionsParams {
	return &GetSpoeTransactionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpoeTransactionsParamsWithTimeout creates a new GetSpoeTransactionsParams object
// with the ability to set a timeout on a request.
func NewGetSpoeTransactionsParamsWithTimeout(timeout time.Duration) *GetSpoeTransactionsParams {
	return &GetSpoeTransactionsParams{
		timeout: timeout,
	}
}

// NewGetSpoeTransactionsParamsWithContext creates a new GetSpoeTransactionsParams object
// with the ability to set a context for a request.
func NewGetSpoeTransactionsParamsWithContext(ctx context.Context) *GetSpoeTransactionsParams {
	return &GetSpoeTransactionsParams{
		Context: ctx,
	}
}

// NewGetSpoeTransactionsParamsWithHTTPClient creates a new GetSpoeTransactionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSpoeTransactionsParamsWithHTTPClient(client *http.Client) *GetSpoeTransactionsParams {
	return &GetSpoeTransactionsParams{
		HTTPClient: client,
	}
}

/*
GetSpoeTransactionsParams contains all the parameters to send to the API endpoint

	for the get spoe transactions operation.

	Typically these are written to a http.Request.
*/
type GetSpoeTransactionsParams struct {

	/* Spoe.

	   Spoe file name
	*/
	Spoe string

	/* Status.

	   Filter by transaction status
	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get spoe transactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSpoeTransactionsParams) WithDefaults() *GetSpoeTransactionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get spoe transactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSpoeTransactionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get spoe transactions params
func (o *GetSpoeTransactionsParams) WithTimeout(timeout time.Duration) *GetSpoeTransactionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get spoe transactions params
func (o *GetSpoeTransactionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get spoe transactions params
func (o *GetSpoeTransactionsParams) WithContext(ctx context.Context) *GetSpoeTransactionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get spoe transactions params
func (o *GetSpoeTransactionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get spoe transactions params
func (o *GetSpoeTransactionsParams) WithHTTPClient(client *http.Client) *GetSpoeTransactionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get spoe transactions params
func (o *GetSpoeTransactionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSpoe adds the spoe to the get spoe transactions params
func (o *GetSpoeTransactionsParams) WithSpoe(spoe string) *GetSpoeTransactionsParams {
	o.SetSpoe(spoe)
	return o
}

// SetSpoe adds the spoe to the get spoe transactions params
func (o *GetSpoeTransactionsParams) SetSpoe(spoe string) {
	o.Spoe = spoe
}

// WithStatus adds the status to the get spoe transactions params
func (o *GetSpoeTransactionsParams) WithStatus(status *string) *GetSpoeTransactionsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get spoe transactions params
func (o *GetSpoeTransactionsParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpoeTransactionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param spoe
	qrSpoe := o.Spoe
	qSpoe := qrSpoe
	if qSpoe != "" {

		if err := r.SetQueryParam("spoe", qSpoe); err != nil {
			return err
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
