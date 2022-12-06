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
	"github.com/go-openapi/swag"
)

// NewStartSpoeTransactionParams creates a new StartSpoeTransactionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartSpoeTransactionParams() *StartSpoeTransactionParams {
	return &StartSpoeTransactionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartSpoeTransactionParamsWithTimeout creates a new StartSpoeTransactionParams object
// with the ability to set a timeout on a request.
func NewStartSpoeTransactionParamsWithTimeout(timeout time.Duration) *StartSpoeTransactionParams {
	return &StartSpoeTransactionParams{
		timeout: timeout,
	}
}

// NewStartSpoeTransactionParamsWithContext creates a new StartSpoeTransactionParams object
// with the ability to set a context for a request.
func NewStartSpoeTransactionParamsWithContext(ctx context.Context) *StartSpoeTransactionParams {
	return &StartSpoeTransactionParams{
		Context: ctx,
	}
}

// NewStartSpoeTransactionParamsWithHTTPClient creates a new StartSpoeTransactionParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartSpoeTransactionParamsWithHTTPClient(client *http.Client) *StartSpoeTransactionParams {
	return &StartSpoeTransactionParams{
		HTTPClient: client,
	}
}

/*
StartSpoeTransactionParams contains all the parameters to send to the API endpoint

	for the start spoe transaction operation.

	Typically these are written to a http.Request.
*/
type StartSpoeTransactionParams struct {

	/* Spoe.

	   Spoe file name
	*/
	Spoe string

	/* Version.

	   Configuration version on which to work on
	*/
	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start spoe transaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartSpoeTransactionParams) WithDefaults() *StartSpoeTransactionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start spoe transaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartSpoeTransactionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start spoe transaction params
func (o *StartSpoeTransactionParams) WithTimeout(timeout time.Duration) *StartSpoeTransactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start spoe transaction params
func (o *StartSpoeTransactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start spoe transaction params
func (o *StartSpoeTransactionParams) WithContext(ctx context.Context) *StartSpoeTransactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start spoe transaction params
func (o *StartSpoeTransactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start spoe transaction params
func (o *StartSpoeTransactionParams) WithHTTPClient(client *http.Client) *StartSpoeTransactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start spoe transaction params
func (o *StartSpoeTransactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSpoe adds the spoe to the start spoe transaction params
func (o *StartSpoeTransactionParams) WithSpoe(spoe string) *StartSpoeTransactionParams {
	o.SetSpoe(spoe)
	return o
}

// SetSpoe adds the spoe to the start spoe transaction params
func (o *StartSpoeTransactionParams) SetSpoe(spoe string) {
	o.Spoe = spoe
}

// WithVersion adds the version to the start spoe transaction params
func (o *StartSpoeTransactionParams) WithVersion(version int64) *StartSpoeTransactionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the start spoe transaction params
func (o *StartSpoeTransactionParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *StartSpoeTransactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param version
	qrVersion := o.Version
	qVersion := swag.FormatInt64(qrVersion)
	if qVersion != "" {

		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
