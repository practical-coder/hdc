// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewDeleteServerParams creates a new DeleteServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteServerParams() *DeleteServerParams {
	return &DeleteServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteServerParamsWithTimeout creates a new DeleteServerParams object
// with the ability to set a timeout on a request.
func NewDeleteServerParamsWithTimeout(timeout time.Duration) *DeleteServerParams {
	return &DeleteServerParams{
		timeout: timeout,
	}
}

// NewDeleteServerParamsWithContext creates a new DeleteServerParams object
// with the ability to set a context for a request.
func NewDeleteServerParamsWithContext(ctx context.Context) *DeleteServerParams {
	return &DeleteServerParams{
		Context: ctx,
	}
}

// NewDeleteServerParamsWithHTTPClient creates a new DeleteServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteServerParamsWithHTTPClient(client *http.Client) *DeleteServerParams {
	return &DeleteServerParams{
		HTTPClient: client,
	}
}

/*
DeleteServerParams contains all the parameters to send to the API endpoint

	for the delete server operation.

	Typically these are written to a http.Request.
*/
type DeleteServerParams struct {

	/* Backend.

	   Parent backend name
	*/
	Backend *string

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Name.

	   Server name
	*/
	Name string

	/* ParentName.

	   Parent name
	*/
	ParentName *string

	/* ParentType.

	   Parent type
	*/
	ParentType *string

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

// WithDefaults hydrates default values in the delete server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteServerParams) WithDefaults() *DeleteServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteServerParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := DeleteServerParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete server params
func (o *DeleteServerParams) WithTimeout(timeout time.Duration) *DeleteServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete server params
func (o *DeleteServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete server params
func (o *DeleteServerParams) WithContext(ctx context.Context) *DeleteServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete server params
func (o *DeleteServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete server params
func (o *DeleteServerParams) WithHTTPClient(client *http.Client) *DeleteServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete server params
func (o *DeleteServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the delete server params
func (o *DeleteServerParams) WithBackend(backend *string) *DeleteServerParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the delete server params
func (o *DeleteServerParams) SetBackend(backend *string) {
	o.Backend = backend
}

// WithForceReload adds the forceReload to the delete server params
func (o *DeleteServerParams) WithForceReload(forceReload *bool) *DeleteServerParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the delete server params
func (o *DeleteServerParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithName adds the name to the delete server params
func (o *DeleteServerParams) WithName(name string) *DeleteServerParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete server params
func (o *DeleteServerParams) SetName(name string) {
	o.Name = name
}

// WithParentName adds the parentName to the delete server params
func (o *DeleteServerParams) WithParentName(parentName *string) *DeleteServerParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the delete server params
func (o *DeleteServerParams) SetParentName(parentName *string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the delete server params
func (o *DeleteServerParams) WithParentType(parentType *string) *DeleteServerParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the delete server params
func (o *DeleteServerParams) SetParentType(parentType *string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the delete server params
func (o *DeleteServerParams) WithTransactionID(transactionID *string) *DeleteServerParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the delete server params
func (o *DeleteServerParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the delete server params
func (o *DeleteServerParams) WithVersion(version *int64) *DeleteServerParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete server params
func (o *DeleteServerParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Backend != nil {

		// query param backend
		var qrBackend string

		if o.Backend != nil {
			qrBackend = *o.Backend
		}
		qBackend := qrBackend
		if qBackend != "" {

			if err := r.SetQueryParam("backend", qBackend); err != nil {
				return err
			}
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

	if o.ParentName != nil {

		// query param parent_name
		var qrParentName string

		if o.ParentName != nil {
			qrParentName = *o.ParentName
		}
		qParentName := qrParentName
		if qParentName != "" {

			if err := r.SetQueryParam("parent_name", qParentName); err != nil {
				return err
			}
		}
	}

	if o.ParentType != nil {

		// query param parent_type
		var qrParentType string

		if o.ParentType != nil {
			qrParentType = *o.ParentType
		}
		qParentType := qrParentType
		if qParentType != "" {

			if err := r.SetQueryParam("parent_type", qParentType); err != nil {
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
