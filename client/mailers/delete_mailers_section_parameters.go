// Code generated by go-swagger; DO NOT EDIT.

package mailers

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

// NewDeleteMailersSectionParams creates a new DeleteMailersSectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMailersSectionParams() *DeleteMailersSectionParams {
	return &DeleteMailersSectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMailersSectionParamsWithTimeout creates a new DeleteMailersSectionParams object
// with the ability to set a timeout on a request.
func NewDeleteMailersSectionParamsWithTimeout(timeout time.Duration) *DeleteMailersSectionParams {
	return &DeleteMailersSectionParams{
		timeout: timeout,
	}
}

// NewDeleteMailersSectionParamsWithContext creates a new DeleteMailersSectionParams object
// with the ability to set a context for a request.
func NewDeleteMailersSectionParamsWithContext(ctx context.Context) *DeleteMailersSectionParams {
	return &DeleteMailersSectionParams{
		Context: ctx,
	}
}

// NewDeleteMailersSectionParamsWithHTTPClient creates a new DeleteMailersSectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMailersSectionParamsWithHTTPClient(client *http.Client) *DeleteMailersSectionParams {
	return &DeleteMailersSectionParams{
		HTTPClient: client,
	}
}

/*
DeleteMailersSectionParams contains all the parameters to send to the API endpoint

	for the delete mailers section operation.

	Typically these are written to a http.Request.
*/
type DeleteMailersSectionParams struct {

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Name.

	   Mailers name
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

// WithDefaults hydrates default values in the delete mailers section params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMailersSectionParams) WithDefaults() *DeleteMailersSectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete mailers section params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMailersSectionParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := DeleteMailersSectionParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete mailers section params
func (o *DeleteMailersSectionParams) WithTimeout(timeout time.Duration) *DeleteMailersSectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete mailers section params
func (o *DeleteMailersSectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete mailers section params
func (o *DeleteMailersSectionParams) WithContext(ctx context.Context) *DeleteMailersSectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete mailers section params
func (o *DeleteMailersSectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete mailers section params
func (o *DeleteMailersSectionParams) WithHTTPClient(client *http.Client) *DeleteMailersSectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete mailers section params
func (o *DeleteMailersSectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceReload adds the forceReload to the delete mailers section params
func (o *DeleteMailersSectionParams) WithForceReload(forceReload *bool) *DeleteMailersSectionParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the delete mailers section params
func (o *DeleteMailersSectionParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithName adds the name to the delete mailers section params
func (o *DeleteMailersSectionParams) WithName(name string) *DeleteMailersSectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete mailers section params
func (o *DeleteMailersSectionParams) SetName(name string) {
	o.Name = name
}

// WithTransactionID adds the transactionID to the delete mailers section params
func (o *DeleteMailersSectionParams) WithTransactionID(transactionID *string) *DeleteMailersSectionParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the delete mailers section params
func (o *DeleteMailersSectionParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the delete mailers section params
func (o *DeleteMailersSectionParams) WithVersion(version *int64) *DeleteMailersSectionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete mailers section params
func (o *DeleteMailersSectionParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMailersSectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
