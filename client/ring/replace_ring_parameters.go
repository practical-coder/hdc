// Code generated by go-swagger; DO NOT EDIT.

package ring

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

// NewReplaceRingParams creates a new ReplaceRingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceRingParams() *ReplaceRingParams {
	return &ReplaceRingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceRingParamsWithTimeout creates a new ReplaceRingParams object
// with the ability to set a timeout on a request.
func NewReplaceRingParamsWithTimeout(timeout time.Duration) *ReplaceRingParams {
	return &ReplaceRingParams{
		timeout: timeout,
	}
}

// NewReplaceRingParamsWithContext creates a new ReplaceRingParams object
// with the ability to set a context for a request.
func NewReplaceRingParamsWithContext(ctx context.Context) *ReplaceRingParams {
	return &ReplaceRingParams{
		Context: ctx,
	}
}

// NewReplaceRingParamsWithHTTPClient creates a new ReplaceRingParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceRingParamsWithHTTPClient(client *http.Client) *ReplaceRingParams {
	return &ReplaceRingParams{
		HTTPClient: client,
	}
}

/*
ReplaceRingParams contains all the parameters to send to the API endpoint

	for the replace ring operation.

	Typically these are written to a http.Request.
*/
type ReplaceRingParams struct {

	// Data.
	Data *models.Ring

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Name.

	   Ring name
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

// WithDefaults hydrates default values in the replace ring params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceRingParams) WithDefaults() *ReplaceRingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace ring params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceRingParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := ReplaceRingParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the replace ring params
func (o *ReplaceRingParams) WithTimeout(timeout time.Duration) *ReplaceRingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace ring params
func (o *ReplaceRingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace ring params
func (o *ReplaceRingParams) WithContext(ctx context.Context) *ReplaceRingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace ring params
func (o *ReplaceRingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace ring params
func (o *ReplaceRingParams) WithHTTPClient(client *http.Client) *ReplaceRingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace ring params
func (o *ReplaceRingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace ring params
func (o *ReplaceRingParams) WithData(data *models.Ring) *ReplaceRingParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace ring params
func (o *ReplaceRingParams) SetData(data *models.Ring) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace ring params
func (o *ReplaceRingParams) WithForceReload(forceReload *bool) *ReplaceRingParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace ring params
func (o *ReplaceRingParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithName adds the name to the replace ring params
func (o *ReplaceRingParams) WithName(name string) *ReplaceRingParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace ring params
func (o *ReplaceRingParams) SetName(name string) {
	o.Name = name
}

// WithTransactionID adds the transactionID to the replace ring params
func (o *ReplaceRingParams) WithTransactionID(transactionID *string) *ReplaceRingParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace ring params
func (o *ReplaceRingParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace ring params
func (o *ReplaceRingParams) WithVersion(version *int64) *ReplaceRingParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace ring params
func (o *ReplaceRingParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceRingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
