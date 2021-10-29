// Code generated by go-swagger; DO NOT EDIT.

package http_request_rule

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

	"github.com/haproxytech/cloud-native/models"
)

// NewReplaceHTTPRequestRuleParams creates a new ReplaceHTTPRequestRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceHTTPRequestRuleParams() *ReplaceHTTPRequestRuleParams {
	return &ReplaceHTTPRequestRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceHTTPRequestRuleParamsWithTimeout creates a new ReplaceHTTPRequestRuleParams object
// with the ability to set a timeout on a request.
func NewReplaceHTTPRequestRuleParamsWithTimeout(timeout time.Duration) *ReplaceHTTPRequestRuleParams {
	return &ReplaceHTTPRequestRuleParams{
		timeout: timeout,
	}
}

// NewReplaceHTTPRequestRuleParamsWithContext creates a new ReplaceHTTPRequestRuleParams object
// with the ability to set a context for a request.
func NewReplaceHTTPRequestRuleParamsWithContext(ctx context.Context) *ReplaceHTTPRequestRuleParams {
	return &ReplaceHTTPRequestRuleParams{
		Context: ctx,
	}
}

// NewReplaceHTTPRequestRuleParamsWithHTTPClient creates a new ReplaceHTTPRequestRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceHTTPRequestRuleParamsWithHTTPClient(client *http.Client) *ReplaceHTTPRequestRuleParams {
	return &ReplaceHTTPRequestRuleParams{
		HTTPClient: client,
	}
}

/* ReplaceHTTPRequestRuleParams contains all the parameters to send to the API endpoint
   for the replace HTTP request rule operation.

   Typically these are written to a http.Request.
*/
type ReplaceHTTPRequestRuleParams struct {

	// Data.
	Data *models.HTTPRequestRule

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Index.

	   HTTP Request Rule Index
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

// WithDefaults hydrates default values in the replace HTTP request rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceHTTPRequestRuleParams) WithDefaults() *ReplaceHTTPRequestRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace HTTP request rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceHTTPRequestRuleParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := ReplaceHTTPRequestRuleParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) WithTimeout(timeout time.Duration) *ReplaceHTTPRequestRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) WithContext(ctx context.Context) *ReplaceHTTPRequestRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) WithHTTPClient(client *http.Client) *ReplaceHTTPRequestRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) WithData(data *models.HTTPRequestRule) *ReplaceHTTPRequestRuleParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) SetData(data *models.HTTPRequestRule) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) WithForceReload(forceReload *bool) *ReplaceHTTPRequestRuleParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithIndex adds the index to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) WithIndex(index int64) *ReplaceHTTPRequestRuleParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) SetIndex(index int64) {
	o.Index = index
}

// WithParentName adds the parentName to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) WithParentName(parentName string) *ReplaceHTTPRequestRuleParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) WithParentType(parentType string) *ReplaceHTTPRequestRuleParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) WithTransactionID(transactionID *string) *ReplaceHTTPRequestRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) WithVersion(version *int64) *ReplaceHTTPRequestRuleParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace HTTP request rule params
func (o *ReplaceHTTPRequestRuleParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceHTTPRequestRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
