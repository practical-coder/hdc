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

	"github.com/haproxytech/client-native/v5/models"
)

// NewReplaceServerTemplateParams creates a new ReplaceServerTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceServerTemplateParams() *ReplaceServerTemplateParams {
	return &ReplaceServerTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceServerTemplateParamsWithTimeout creates a new ReplaceServerTemplateParams object
// with the ability to set a timeout on a request.
func NewReplaceServerTemplateParamsWithTimeout(timeout time.Duration) *ReplaceServerTemplateParams {
	return &ReplaceServerTemplateParams{
		timeout: timeout,
	}
}

// NewReplaceServerTemplateParamsWithContext creates a new ReplaceServerTemplateParams object
// with the ability to set a context for a request.
func NewReplaceServerTemplateParamsWithContext(ctx context.Context) *ReplaceServerTemplateParams {
	return &ReplaceServerTemplateParams{
		Context: ctx,
	}
}

// NewReplaceServerTemplateParamsWithHTTPClient creates a new ReplaceServerTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceServerTemplateParamsWithHTTPClient(client *http.Client) *ReplaceServerTemplateParams {
	return &ReplaceServerTemplateParams{
		HTTPClient: client,
	}
}

/*
ReplaceServerTemplateParams contains all the parameters to send to the API endpoint

	for the replace server template operation.

	Typically these are written to a http.Request.
*/
type ReplaceServerTemplateParams struct {

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

	/* Prefix.

	   Server template prefix
	*/
	Prefix string

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

// WithDefaults hydrates default values in the replace server template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceServerTemplateParams) WithDefaults() *ReplaceServerTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace server template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceServerTemplateParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := ReplaceServerTemplateParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the replace server template params
func (o *ReplaceServerTemplateParams) WithTimeout(timeout time.Duration) *ReplaceServerTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace server template params
func (o *ReplaceServerTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace server template params
func (o *ReplaceServerTemplateParams) WithContext(ctx context.Context) *ReplaceServerTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace server template params
func (o *ReplaceServerTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace server template params
func (o *ReplaceServerTemplateParams) WithHTTPClient(client *http.Client) *ReplaceServerTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace server template params
func (o *ReplaceServerTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the replace server template params
func (o *ReplaceServerTemplateParams) WithBackend(backend string) *ReplaceServerTemplateParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the replace server template params
func (o *ReplaceServerTemplateParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithData adds the data to the replace server template params
func (o *ReplaceServerTemplateParams) WithData(data *models.ServerTemplate) *ReplaceServerTemplateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace server template params
func (o *ReplaceServerTemplateParams) SetData(data *models.ServerTemplate) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace server template params
func (o *ReplaceServerTemplateParams) WithForceReload(forceReload *bool) *ReplaceServerTemplateParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace server template params
func (o *ReplaceServerTemplateParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithPrefix adds the prefix to the replace server template params
func (o *ReplaceServerTemplateParams) WithPrefix(prefix string) *ReplaceServerTemplateParams {
	o.SetPrefix(prefix)
	return o
}

// SetPrefix adds the prefix to the replace server template params
func (o *ReplaceServerTemplateParams) SetPrefix(prefix string) {
	o.Prefix = prefix
}

// WithTransactionID adds the transactionID to the replace server template params
func (o *ReplaceServerTemplateParams) WithTransactionID(transactionID *string) *ReplaceServerTemplateParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace server template params
func (o *ReplaceServerTemplateParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace server template params
func (o *ReplaceServerTemplateParams) WithVersion(version *int64) *ReplaceServerTemplateParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace server template params
func (o *ReplaceServerTemplateParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceServerTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param prefix
	if err := r.SetPathParam("prefix", o.Prefix); err != nil {
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
