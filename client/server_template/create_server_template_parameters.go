// Code generated by go-swagger; DO NOT EDIT.

package server_template

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

	"github.com/haproxytech/client-native/v4/models"
)

// NewCreateServerTemplateParams creates a new CreateServerTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateServerTemplateParams() *CreateServerTemplateParams {
	return &CreateServerTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateServerTemplateParamsWithTimeout creates a new CreateServerTemplateParams object
// with the ability to set a timeout on a request.
func NewCreateServerTemplateParamsWithTimeout(timeout time.Duration) *CreateServerTemplateParams {
	return &CreateServerTemplateParams{
		timeout: timeout,
	}
}

// NewCreateServerTemplateParamsWithContext creates a new CreateServerTemplateParams object
// with the ability to set a context for a request.
func NewCreateServerTemplateParamsWithContext(ctx context.Context) *CreateServerTemplateParams {
	return &CreateServerTemplateParams{
		Context: ctx,
	}
}

// NewCreateServerTemplateParamsWithHTTPClient creates a new CreateServerTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateServerTemplateParamsWithHTTPClient(client *http.Client) *CreateServerTemplateParams {
	return &CreateServerTemplateParams{
		HTTPClient: client,
	}
}

/*
CreateServerTemplateParams contains all the parameters to send to the API endpoint

	for the create server template operation.

	Typically these are written to a http.Request.
*/
type CreateServerTemplateParams struct {

	/* Backend.

	   Parent backend name
	*/
	Backend string

	// Data.
	Data *models.ServerTemplate

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

// WithDefaults hydrates default values in the create server template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServerTemplateParams) WithDefaults() *CreateServerTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create server template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServerTemplateParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := CreateServerTemplateParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create server template params
func (o *CreateServerTemplateParams) WithTimeout(timeout time.Duration) *CreateServerTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create server template params
func (o *CreateServerTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create server template params
func (o *CreateServerTemplateParams) WithContext(ctx context.Context) *CreateServerTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create server template params
func (o *CreateServerTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create server template params
func (o *CreateServerTemplateParams) WithHTTPClient(client *http.Client) *CreateServerTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create server template params
func (o *CreateServerTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the create server template params
func (o *CreateServerTemplateParams) WithBackend(backend string) *CreateServerTemplateParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the create server template params
func (o *CreateServerTemplateParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithData adds the data to the create server template params
func (o *CreateServerTemplateParams) WithData(data *models.ServerTemplate) *CreateServerTemplateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create server template params
func (o *CreateServerTemplateParams) SetData(data *models.ServerTemplate) {
	o.Data = data
}

// WithForceReload adds the forceReload to the create server template params
func (o *CreateServerTemplateParams) WithForceReload(forceReload *bool) *CreateServerTemplateParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create server template params
func (o *CreateServerTemplateParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithTransactionID adds the transactionID to the create server template params
func (o *CreateServerTemplateParams) WithTransactionID(transactionID *string) *CreateServerTemplateParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create server template params
func (o *CreateServerTemplateParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create server template params
func (o *CreateServerTemplateParams) WithVersion(version *int64) *CreateServerTemplateParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create server template params
func (o *CreateServerTemplateParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateServerTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param backend
	qrBackend := o.Backend
	qBackend := qrBackend
	if qBackend != "" {

		if err := r.SetQueryParam("backend", qBackend); err != nil {
			return err
		}
	}
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
