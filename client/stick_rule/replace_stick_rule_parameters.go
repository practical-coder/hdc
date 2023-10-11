// Code generated by go-swagger; DO NOT EDIT.

package stick_rule

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

// NewReplaceStickRuleParams creates a new ReplaceStickRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceStickRuleParams() *ReplaceStickRuleParams {
	return &ReplaceStickRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceStickRuleParamsWithTimeout creates a new ReplaceStickRuleParams object
// with the ability to set a timeout on a request.
func NewReplaceStickRuleParamsWithTimeout(timeout time.Duration) *ReplaceStickRuleParams {
	return &ReplaceStickRuleParams{
		timeout: timeout,
	}
}

// NewReplaceStickRuleParamsWithContext creates a new ReplaceStickRuleParams object
// with the ability to set a context for a request.
func NewReplaceStickRuleParamsWithContext(ctx context.Context) *ReplaceStickRuleParams {
	return &ReplaceStickRuleParams{
		Context: ctx,
	}
}

// NewReplaceStickRuleParamsWithHTTPClient creates a new ReplaceStickRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceStickRuleParamsWithHTTPClient(client *http.Client) *ReplaceStickRuleParams {
	return &ReplaceStickRuleParams{
		HTTPClient: client,
	}
}

/*
ReplaceStickRuleParams contains all the parameters to send to the API endpoint

	for the replace stick rule operation.

	Typically these are written to a http.Request.
*/
type ReplaceStickRuleParams struct {

	/* Backend.

	   Backend name
	*/
	Backend string

	// Data.
	Data *models.StickRule

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Index.

	   Stick Rule Index
	*/
	Index int64

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

// WithDefaults hydrates default values in the replace stick rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceStickRuleParams) WithDefaults() *ReplaceStickRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace stick rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceStickRuleParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := ReplaceStickRuleParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the replace stick rule params
func (o *ReplaceStickRuleParams) WithTimeout(timeout time.Duration) *ReplaceStickRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace stick rule params
func (o *ReplaceStickRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace stick rule params
func (o *ReplaceStickRuleParams) WithContext(ctx context.Context) *ReplaceStickRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace stick rule params
func (o *ReplaceStickRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace stick rule params
func (o *ReplaceStickRuleParams) WithHTTPClient(client *http.Client) *ReplaceStickRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace stick rule params
func (o *ReplaceStickRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the replace stick rule params
func (o *ReplaceStickRuleParams) WithBackend(backend string) *ReplaceStickRuleParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the replace stick rule params
func (o *ReplaceStickRuleParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithData adds the data to the replace stick rule params
func (o *ReplaceStickRuleParams) WithData(data *models.StickRule) *ReplaceStickRuleParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace stick rule params
func (o *ReplaceStickRuleParams) SetData(data *models.StickRule) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace stick rule params
func (o *ReplaceStickRuleParams) WithForceReload(forceReload *bool) *ReplaceStickRuleParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace stick rule params
func (o *ReplaceStickRuleParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithIndex adds the index to the replace stick rule params
func (o *ReplaceStickRuleParams) WithIndex(index int64) *ReplaceStickRuleParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the replace stick rule params
func (o *ReplaceStickRuleParams) SetIndex(index int64) {
	o.Index = index
}

// WithTransactionID adds the transactionID to the replace stick rule params
func (o *ReplaceStickRuleParams) WithTransactionID(transactionID *string) *ReplaceStickRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace stick rule params
func (o *ReplaceStickRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace stick rule params
func (o *ReplaceStickRuleParams) WithVersion(version *int64) *ReplaceStickRuleParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace stick rule params
func (o *ReplaceStickRuleParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceStickRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
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
