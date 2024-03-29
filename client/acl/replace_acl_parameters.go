// Code generated by go-swagger; DO NOT EDIT.

package acl

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

// NewReplaceACLParams creates a new ReplaceACLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceACLParams() *ReplaceACLParams {
	return &ReplaceACLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceACLParamsWithTimeout creates a new ReplaceACLParams object
// with the ability to set a timeout on a request.
func NewReplaceACLParamsWithTimeout(timeout time.Duration) *ReplaceACLParams {
	return &ReplaceACLParams{
		timeout: timeout,
	}
}

// NewReplaceACLParamsWithContext creates a new ReplaceACLParams object
// with the ability to set a context for a request.
func NewReplaceACLParamsWithContext(ctx context.Context) *ReplaceACLParams {
	return &ReplaceACLParams{
		Context: ctx,
	}
}

// NewReplaceACLParamsWithHTTPClient creates a new ReplaceACLParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceACLParamsWithHTTPClient(client *http.Client) *ReplaceACLParams {
	return &ReplaceACLParams{
		HTTPClient: client,
	}
}

/*
ReplaceACLParams contains all the parameters to send to the API endpoint

	for the replace Acl operation.

	Typically these are written to a http.Request.
*/
type ReplaceACLParams struct {

	// Data.
	Data *models.ACL

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Index.

	   ACL line Index
	*/
	Index int64

	/* ParentName.

	   Parent name
	*/
	ParentName string

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

// WithDefaults hydrates default values in the replace Acl params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceACLParams) WithDefaults() *ReplaceACLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace Acl params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceACLParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := ReplaceACLParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the replace Acl params
func (o *ReplaceACLParams) WithTimeout(timeout time.Duration) *ReplaceACLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace Acl params
func (o *ReplaceACLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace Acl params
func (o *ReplaceACLParams) WithContext(ctx context.Context) *ReplaceACLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace Acl params
func (o *ReplaceACLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace Acl params
func (o *ReplaceACLParams) WithHTTPClient(client *http.Client) *ReplaceACLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace Acl params
func (o *ReplaceACLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace Acl params
func (o *ReplaceACLParams) WithData(data *models.ACL) *ReplaceACLParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace Acl params
func (o *ReplaceACLParams) SetData(data *models.ACL) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace Acl params
func (o *ReplaceACLParams) WithForceReload(forceReload *bool) *ReplaceACLParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace Acl params
func (o *ReplaceACLParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithIndex adds the index to the replace Acl params
func (o *ReplaceACLParams) WithIndex(index int64) *ReplaceACLParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the replace Acl params
func (o *ReplaceACLParams) SetIndex(index int64) {
	o.Index = index
}

// WithParentName adds the parentName to the replace Acl params
func (o *ReplaceACLParams) WithParentName(parentName string) *ReplaceACLParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the replace Acl params
func (o *ReplaceACLParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the replace Acl params
func (o *ReplaceACLParams) WithParentType(parentType string) *ReplaceACLParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the replace Acl params
func (o *ReplaceACLParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the replace Acl params
func (o *ReplaceACLParams) WithTransactionID(transactionID *string) *ReplaceACLParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace Acl params
func (o *ReplaceACLParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace Acl params
func (o *ReplaceACLParams) WithVersion(version *int64) *ReplaceACLParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace Acl params
func (o *ReplaceACLParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceACLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}

	// query param parent_name
	qrParentName := o.ParentName
	qParentName := qrParentName
	if qParentName != "" {

		if err := r.SetQueryParam("parent_name", qParentName); err != nil {
			return err
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
