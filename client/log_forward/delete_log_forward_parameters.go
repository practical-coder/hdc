// Code generated by go-swagger; DO NOT EDIT.

package log_forward

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

// NewDeleteLogForwardParams creates a new DeleteLogForwardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteLogForwardParams() *DeleteLogForwardParams {
	return &DeleteLogForwardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLogForwardParamsWithTimeout creates a new DeleteLogForwardParams object
// with the ability to set a timeout on a request.
func NewDeleteLogForwardParamsWithTimeout(timeout time.Duration) *DeleteLogForwardParams {
	return &DeleteLogForwardParams{
		timeout: timeout,
	}
}

// NewDeleteLogForwardParamsWithContext creates a new DeleteLogForwardParams object
// with the ability to set a context for a request.
func NewDeleteLogForwardParamsWithContext(ctx context.Context) *DeleteLogForwardParams {
	return &DeleteLogForwardParams{
		Context: ctx,
	}
}

// NewDeleteLogForwardParamsWithHTTPClient creates a new DeleteLogForwardParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteLogForwardParamsWithHTTPClient(client *http.Client) *DeleteLogForwardParams {
	return &DeleteLogForwardParams{
		HTTPClient: client,
	}
}

/*
DeleteLogForwardParams contains all the parameters to send to the API endpoint

	for the delete log forward operation.

	Typically these are written to a http.Request.
*/
type DeleteLogForwardParams struct {

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Name.

	   Log Forward name
	*/
	Name string

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

// WithDefaults hydrates default values in the delete log forward params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLogForwardParams) WithDefaults() *DeleteLogForwardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete log forward params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLogForwardParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := DeleteLogForwardParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete log forward params
func (o *DeleteLogForwardParams) WithTimeout(timeout time.Duration) *DeleteLogForwardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete log forward params
func (o *DeleteLogForwardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete log forward params
func (o *DeleteLogForwardParams) WithContext(ctx context.Context) *DeleteLogForwardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete log forward params
func (o *DeleteLogForwardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete log forward params
func (o *DeleteLogForwardParams) WithHTTPClient(client *http.Client) *DeleteLogForwardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete log forward params
func (o *DeleteLogForwardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceReload adds the forceReload to the delete log forward params
func (o *DeleteLogForwardParams) WithForceReload(forceReload *bool) *DeleteLogForwardParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the delete log forward params
func (o *DeleteLogForwardParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithName adds the name to the delete log forward params
func (o *DeleteLogForwardParams) WithName(name string) *DeleteLogForwardParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete log forward params
func (o *DeleteLogForwardParams) SetName(name string) {
	o.Name = name
}

// WithTransactionID adds the transactionID to the delete log forward params
func (o *DeleteLogForwardParams) WithTransactionID(transactionID *string) *DeleteLogForwardParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the delete log forward params
func (o *DeleteLogForwardParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the delete log forward params
func (o *DeleteLogForwardParams) WithVersion(version *int64) *DeleteLogForwardParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete log forward params
func (o *DeleteLogForwardParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLogForwardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForceReload != nil {

		// query param force_reload
		var qrForceReload bool

		if o.ForceReload != nil {
			qrForceReload = *o.ForceReload
		}
		qForceReload := swag.FormatBool(qrForceReload)
		if qForceReload != "" {

			if err := r.SetQueryParam("force_reload", qForceReload); err != nil {
				return err
			}
		}
	}

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
