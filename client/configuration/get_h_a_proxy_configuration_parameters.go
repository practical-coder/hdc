// Code generated by go-swagger; DO NOT EDIT.

package configuration

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

// NewGetHAProxyConfigurationParams creates a new GetHAProxyConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHAProxyConfigurationParams() *GetHAProxyConfigurationParams {
	return &GetHAProxyConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHAProxyConfigurationParamsWithTimeout creates a new GetHAProxyConfigurationParams object
// with the ability to set a timeout on a request.
func NewGetHAProxyConfigurationParamsWithTimeout(timeout time.Duration) *GetHAProxyConfigurationParams {
	return &GetHAProxyConfigurationParams{
		timeout: timeout,
	}
}

// NewGetHAProxyConfigurationParamsWithContext creates a new GetHAProxyConfigurationParams object
// with the ability to set a context for a request.
func NewGetHAProxyConfigurationParamsWithContext(ctx context.Context) *GetHAProxyConfigurationParams {
	return &GetHAProxyConfigurationParams{
		Context: ctx,
	}
}

// NewGetHAProxyConfigurationParamsWithHTTPClient creates a new GetHAProxyConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHAProxyConfigurationParamsWithHTTPClient(client *http.Client) *GetHAProxyConfigurationParams {
	return &GetHAProxyConfigurationParams{
		HTTPClient: client,
	}
}

/*
GetHAProxyConfigurationParams contains all the parameters to send to the API endpoint

	for the get h a proxy configuration operation.

	Typically these are written to a http.Request.
*/
type GetHAProxyConfigurationParams struct {

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	/* Version.

	   Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.
	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get h a proxy configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHAProxyConfigurationParams) WithDefaults() *GetHAProxyConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get h a proxy configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHAProxyConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get h a proxy configuration params
func (o *GetHAProxyConfigurationParams) WithTimeout(timeout time.Duration) *GetHAProxyConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get h a proxy configuration params
func (o *GetHAProxyConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get h a proxy configuration params
func (o *GetHAProxyConfigurationParams) WithContext(ctx context.Context) *GetHAProxyConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get h a proxy configuration params
func (o *GetHAProxyConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get h a proxy configuration params
func (o *GetHAProxyConfigurationParams) WithHTTPClient(client *http.Client) *GetHAProxyConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get h a proxy configuration params
func (o *GetHAProxyConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTransactionID adds the transactionID to the get h a proxy configuration params
func (o *GetHAProxyConfigurationParams) WithTransactionID(transactionID *string) *GetHAProxyConfigurationParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get h a proxy configuration params
func (o *GetHAProxyConfigurationParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the get h a proxy configuration params
func (o *GetHAProxyConfigurationParams) WithVersion(version *int64) *GetHAProxyConfigurationParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get h a proxy configuration params
func (o *GetHAProxyConfigurationParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetHAProxyConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Version != nil {

		// query param version
		var qrVersion int64

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
