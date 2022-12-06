// Code generated by go-swagger; DO NOT EDIT.

package dgram_bind

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

// NewDeleteDgramBindParams creates a new DeleteDgramBindParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDgramBindParams() *DeleteDgramBindParams {
	return &DeleteDgramBindParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDgramBindParamsWithTimeout creates a new DeleteDgramBindParams object
// with the ability to set a timeout on a request.
func NewDeleteDgramBindParamsWithTimeout(timeout time.Duration) *DeleteDgramBindParams {
	return &DeleteDgramBindParams{
		timeout: timeout,
	}
}

// NewDeleteDgramBindParamsWithContext creates a new DeleteDgramBindParams object
// with the ability to set a context for a request.
func NewDeleteDgramBindParamsWithContext(ctx context.Context) *DeleteDgramBindParams {
	return &DeleteDgramBindParams{
		Context: ctx,
	}
}

// NewDeleteDgramBindParamsWithHTTPClient creates a new DeleteDgramBindParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDgramBindParamsWithHTTPClient(client *http.Client) *DeleteDgramBindParams {
	return &DeleteDgramBindParams{
		HTTPClient: client,
	}
}

/*
DeleteDgramBindParams contains all the parameters to send to the API endpoint

	for the delete dgram bind operation.

	Typically these are written to a http.Request.
*/
type DeleteDgramBindParams struct {

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* LogForward.

	   Parent log forward name
	*/
	LogForward string

	/* Name.

	   Bind name
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

// WithDefaults hydrates default values in the delete dgram bind params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDgramBindParams) WithDefaults() *DeleteDgramBindParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete dgram bind params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDgramBindParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := DeleteDgramBindParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete dgram bind params
func (o *DeleteDgramBindParams) WithTimeout(timeout time.Duration) *DeleteDgramBindParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete dgram bind params
func (o *DeleteDgramBindParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete dgram bind params
func (o *DeleteDgramBindParams) WithContext(ctx context.Context) *DeleteDgramBindParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete dgram bind params
func (o *DeleteDgramBindParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete dgram bind params
func (o *DeleteDgramBindParams) WithHTTPClient(client *http.Client) *DeleteDgramBindParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete dgram bind params
func (o *DeleteDgramBindParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceReload adds the forceReload to the delete dgram bind params
func (o *DeleteDgramBindParams) WithForceReload(forceReload *bool) *DeleteDgramBindParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the delete dgram bind params
func (o *DeleteDgramBindParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithLogForward adds the logForward to the delete dgram bind params
func (o *DeleteDgramBindParams) WithLogForward(logForward string) *DeleteDgramBindParams {
	o.SetLogForward(logForward)
	return o
}

// SetLogForward adds the logForward to the delete dgram bind params
func (o *DeleteDgramBindParams) SetLogForward(logForward string) {
	o.LogForward = logForward
}

// WithName adds the name to the delete dgram bind params
func (o *DeleteDgramBindParams) WithName(name string) *DeleteDgramBindParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete dgram bind params
func (o *DeleteDgramBindParams) SetName(name string) {
	o.Name = name
}

// WithTransactionID adds the transactionID to the delete dgram bind params
func (o *DeleteDgramBindParams) WithTransactionID(transactionID *string) *DeleteDgramBindParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the delete dgram bind params
func (o *DeleteDgramBindParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the delete dgram bind params
func (o *DeleteDgramBindParams) WithVersion(version *int64) *DeleteDgramBindParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete dgram bind params
func (o *DeleteDgramBindParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDgramBindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param log_forward
	qrLogForward := o.LogForward
	qLogForward := qrLogForward
	if qLogForward != "" {

		if err := r.SetQueryParam("log_forward", qLogForward); err != nil {
			return err
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
