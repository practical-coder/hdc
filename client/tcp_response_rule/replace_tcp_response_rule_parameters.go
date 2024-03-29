// Code generated by go-swagger; DO NOT EDIT.

package tcp_response_rule

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

// NewReplaceTCPResponseRuleParams creates a new ReplaceTCPResponseRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceTCPResponseRuleParams() *ReplaceTCPResponseRuleParams {
	return &ReplaceTCPResponseRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceTCPResponseRuleParamsWithTimeout creates a new ReplaceTCPResponseRuleParams object
// with the ability to set a timeout on a request.
func NewReplaceTCPResponseRuleParamsWithTimeout(timeout time.Duration) *ReplaceTCPResponseRuleParams {
	return &ReplaceTCPResponseRuleParams{
		timeout: timeout,
	}
}

// NewReplaceTCPResponseRuleParamsWithContext creates a new ReplaceTCPResponseRuleParams object
// with the ability to set a context for a request.
func NewReplaceTCPResponseRuleParamsWithContext(ctx context.Context) *ReplaceTCPResponseRuleParams {
	return &ReplaceTCPResponseRuleParams{
		Context: ctx,
	}
}

// NewReplaceTCPResponseRuleParamsWithHTTPClient creates a new ReplaceTCPResponseRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceTCPResponseRuleParamsWithHTTPClient(client *http.Client) *ReplaceTCPResponseRuleParams {
	return &ReplaceTCPResponseRuleParams{
		HTTPClient: client,
	}
}

/*
ReplaceTCPResponseRuleParams contains all the parameters to send to the API endpoint

	for the replace TCP response rule operation.

	Typically these are written to a http.Request.
*/
type ReplaceTCPResponseRuleParams struct {

	/* Backend.

	   Parent backend name
	*/
	Backend string

	// Data.
	Data *models.TCPResponseRule

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Index.

	   TCP Response Rule Index
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

// WithDefaults hydrates default values in the replace TCP response rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceTCPResponseRuleParams) WithDefaults() *ReplaceTCPResponseRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace TCP response rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceTCPResponseRuleParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := ReplaceTCPResponseRuleParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) WithTimeout(timeout time.Duration) *ReplaceTCPResponseRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) WithContext(ctx context.Context) *ReplaceTCPResponseRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) WithHTTPClient(client *http.Client) *ReplaceTCPResponseRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) WithBackend(backend string) *ReplaceTCPResponseRuleParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithData adds the data to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) WithData(data *models.TCPResponseRule) *ReplaceTCPResponseRuleParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) SetData(data *models.TCPResponseRule) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) WithForceReload(forceReload *bool) *ReplaceTCPResponseRuleParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithIndex adds the index to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) WithIndex(index int64) *ReplaceTCPResponseRuleParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) SetIndex(index int64) {
	o.Index = index
}

// WithTransactionID adds the transactionID to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) WithTransactionID(transactionID *string) *ReplaceTCPResponseRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) WithVersion(version *int64) *ReplaceTCPResponseRuleParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace TCP response rule params
func (o *ReplaceTCPResponseRuleParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceTCPResponseRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
