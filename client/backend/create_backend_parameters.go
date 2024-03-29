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
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v5/models"
)

// NewCreateBackendParams creates a new CreateBackendParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBackendParams() *CreateBackendParams {
	return &CreateBackendParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBackendParamsWithTimeout creates a new CreateBackendParams object
// with the ability to set a timeout on a request.
func NewCreateBackendParamsWithTimeout(timeout time.Duration) *CreateBackendParams {
	return &CreateBackendParams{
		timeout: timeout,
	}
}

// NewCreateBackendParamsWithContext creates a new CreateBackendParams object
// with the ability to set a context for a request.
func NewCreateBackendParamsWithContext(ctx context.Context) *CreateBackendParams {
	return &CreateBackendParams{
		Context: ctx,
	}
}

// NewCreateBackendParamsWithHTTPClient creates a new CreateBackendParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBackendParamsWithHTTPClient(client *http.Client) *CreateBackendParams {
	return &CreateBackendParams{
		HTTPClient: client,
	}
}

/*
CreateBackendParams contains all the parameters to send to the API endpoint

	for the create backend operation.

	Typically these are written to a http.Request.
*/
type CreateBackendParams struct {

	// Data.
	Data *models.Backend

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

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

// WithDefaults hydrates default values in the create backend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBackendParams) WithDefaults() *CreateBackendParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create backend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBackendParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := CreateBackendParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create backend params
func (o *CreateBackendParams) WithTimeout(timeout time.Duration) *CreateBackendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create backend params
func (o *CreateBackendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create backend params
func (o *CreateBackendParams) WithContext(ctx context.Context) *CreateBackendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create backend params
func (o *CreateBackendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create backend params
func (o *CreateBackendParams) WithHTTPClient(client *http.Client) *CreateBackendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create backend params
func (o *CreateBackendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create backend params
func (o *CreateBackendParams) WithData(data *models.Backend) *CreateBackendParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create backend params
func (o *CreateBackendParams) SetData(data *models.Backend) {
	o.Data = data
}

// WithForceReload adds the forceReload to the create backend params
func (o *CreateBackendParams) WithForceReload(forceReload *bool) *CreateBackendParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create backend params
func (o *CreateBackendParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithTransactionID adds the transactionID to the create backend params
func (o *CreateBackendParams) WithTransactionID(transactionID *string) *CreateBackendParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create backend params
func (o *CreateBackendParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create backend params
func (o *CreateBackendParams) WithVersion(version *int64) *CreateBackendParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create backend params
func (o *CreateBackendParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBackendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

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
