// Code generated by go-swagger; DO NOT EDIT.

package http_error_rule

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

// NewDeleteHTTPErrorRuleParams creates a new DeleteHTTPErrorRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteHTTPErrorRuleParams() *DeleteHTTPErrorRuleParams {
	return &DeleteHTTPErrorRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteHTTPErrorRuleParamsWithTimeout creates a new DeleteHTTPErrorRuleParams object
// with the ability to set a timeout on a request.
func NewDeleteHTTPErrorRuleParamsWithTimeout(timeout time.Duration) *DeleteHTTPErrorRuleParams {
	return &DeleteHTTPErrorRuleParams{
		timeout: timeout,
	}
}

// NewDeleteHTTPErrorRuleParamsWithContext creates a new DeleteHTTPErrorRuleParams object
// with the ability to set a context for a request.
func NewDeleteHTTPErrorRuleParamsWithContext(ctx context.Context) *DeleteHTTPErrorRuleParams {
	return &DeleteHTTPErrorRuleParams{
		Context: ctx,
	}
}

// NewDeleteHTTPErrorRuleParamsWithHTTPClient creates a new DeleteHTTPErrorRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteHTTPErrorRuleParamsWithHTTPClient(client *http.Client) *DeleteHTTPErrorRuleParams {
	return &DeleteHTTPErrorRuleParams{
		HTTPClient: client,
	}
}

/*
DeleteHTTPErrorRuleParams contains all the parameters to send to the API endpoint

	for the delete HTTP error rule operation.

	Typically these are written to a http.Request.
*/
type DeleteHTTPErrorRuleParams struct {

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Index.

	   HTTP Error Rule Index
	*/
	Index int64

	/* ParentName.

	   Parent name
	*/
	ParentName *string

	/* ParentType.

	   Parent type
	*/
	ParentType string

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

// WithDefaults hydrates default values in the delete HTTP error rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteHTTPErrorRuleParams) WithDefaults() *DeleteHTTPErrorRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete HTTP error rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteHTTPErrorRuleParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := DeleteHTTPErrorRuleParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) WithTimeout(timeout time.Duration) *DeleteHTTPErrorRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) WithContext(ctx context.Context) *DeleteHTTPErrorRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) WithHTTPClient(client *http.Client) *DeleteHTTPErrorRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceReload adds the forceReload to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) WithForceReload(forceReload *bool) *DeleteHTTPErrorRuleParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithIndex adds the index to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) WithIndex(index int64) *DeleteHTTPErrorRuleParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) SetIndex(index int64) {
	o.Index = index
}

// WithParentName adds the parentName to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) WithParentName(parentName *string) *DeleteHTTPErrorRuleParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) SetParentName(parentName *string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) WithParentType(parentType string) *DeleteHTTPErrorRuleParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) WithTransactionID(transactionID *string) *DeleteHTTPErrorRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) WithVersion(version *int64) *DeleteHTTPErrorRuleParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete HTTP error rule params
func (o *DeleteHTTPErrorRuleParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteHTTPErrorRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
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

	// query param parent_type
	qrParentType := o.ParentType
	qParentType := qrParentType
	if qParentType != "" {

		if err := r.SetQueryParam("parent_type", qParentType); err != nil {
			return err
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
