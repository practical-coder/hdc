// Code generated by go-swagger; DO NOT EDIT.

package bind

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

// NewDeleteBindParams creates a new DeleteBindParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBindParams() *DeleteBindParams {
	return &DeleteBindParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBindParamsWithTimeout creates a new DeleteBindParams object
// with the ability to set a timeout on a request.
func NewDeleteBindParamsWithTimeout(timeout time.Duration) *DeleteBindParams {
	return &DeleteBindParams{
		timeout: timeout,
	}
}

// NewDeleteBindParamsWithContext creates a new DeleteBindParams object
// with the ability to set a context for a request.
func NewDeleteBindParamsWithContext(ctx context.Context) *DeleteBindParams {
	return &DeleteBindParams{
		Context: ctx,
	}
}

// NewDeleteBindParamsWithHTTPClient creates a new DeleteBindParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBindParamsWithHTTPClient(client *http.Client) *DeleteBindParams {
	return &DeleteBindParams{
		HTTPClient: client,
	}
}

/*
DeleteBindParams contains all the parameters to send to the API endpoint

	for the delete bind operation.

	Typically these are written to a http.Request.
*/
type DeleteBindParams struct {

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Frontend.

	   Parent frontend name
	*/
	Frontend *string

	/* Name.

	   Bind name
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

// WithDefaults hydrates default values in the delete bind params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBindParams) WithDefaults() *DeleteBindParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete bind params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBindParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := DeleteBindParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete bind params
func (o *DeleteBindParams) WithTimeout(timeout time.Duration) *DeleteBindParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete bind params
func (o *DeleteBindParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete bind params
func (o *DeleteBindParams) WithContext(ctx context.Context) *DeleteBindParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete bind params
func (o *DeleteBindParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete bind params
func (o *DeleteBindParams) WithHTTPClient(client *http.Client) *DeleteBindParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete bind params
func (o *DeleteBindParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceReload adds the forceReload to the delete bind params
func (o *DeleteBindParams) WithForceReload(forceReload *bool) *DeleteBindParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the delete bind params
func (o *DeleteBindParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithFrontend adds the frontend to the delete bind params
func (o *DeleteBindParams) WithFrontend(frontend *string) *DeleteBindParams {
	o.SetFrontend(frontend)
	return o
}

// SetFrontend adds the frontend to the delete bind params
func (o *DeleteBindParams) SetFrontend(frontend *string) {
	o.Frontend = frontend
}

// WithName adds the name to the delete bind params
func (o *DeleteBindParams) WithName(name string) *DeleteBindParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete bind params
func (o *DeleteBindParams) SetName(name string) {
	o.Name = name
}

// WithParentName adds the parentName to the delete bind params
func (o *DeleteBindParams) WithParentName(parentName *string) *DeleteBindParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the delete bind params
func (o *DeleteBindParams) SetParentName(parentName *string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the delete bind params
func (o *DeleteBindParams) WithParentType(parentType *string) *DeleteBindParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the delete bind params
func (o *DeleteBindParams) SetParentType(parentType *string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the delete bind params
func (o *DeleteBindParams) WithTransactionID(transactionID *string) *DeleteBindParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the delete bind params
func (o *DeleteBindParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the delete bind params
func (o *DeleteBindParams) WithVersion(version *int64) *DeleteBindParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete bind params
func (o *DeleteBindParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Frontend != nil {

		// query param frontend
		var qrFrontend string

		if o.Frontend != nil {
			qrFrontend = *o.Frontend
		}
		qFrontend := qrFrontend
		if qFrontend != "" {

			if err := r.SetQueryParam("frontend", qFrontend); err != nil {
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
