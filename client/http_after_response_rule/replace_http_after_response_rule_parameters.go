// Code generated by go-swagger; DO NOT EDIT.

package http_after_response_rule

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

// NewReplaceHTTPAfterResponseRuleParams creates a new ReplaceHTTPAfterResponseRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceHTTPAfterResponseRuleParams() *ReplaceHTTPAfterResponseRuleParams {
	return &ReplaceHTTPAfterResponseRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceHTTPAfterResponseRuleParamsWithTimeout creates a new ReplaceHTTPAfterResponseRuleParams object
// with the ability to set a timeout on a request.
func NewReplaceHTTPAfterResponseRuleParamsWithTimeout(timeout time.Duration) *ReplaceHTTPAfterResponseRuleParams {
	return &ReplaceHTTPAfterResponseRuleParams{
		timeout: timeout,
	}
}

// NewReplaceHTTPAfterResponseRuleParamsWithContext creates a new ReplaceHTTPAfterResponseRuleParams object
// with the ability to set a context for a request.
func NewReplaceHTTPAfterResponseRuleParamsWithContext(ctx context.Context) *ReplaceHTTPAfterResponseRuleParams {
	return &ReplaceHTTPAfterResponseRuleParams{
		Context: ctx,
	}
}

// NewReplaceHTTPAfterResponseRuleParamsWithHTTPClient creates a new ReplaceHTTPAfterResponseRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceHTTPAfterResponseRuleParamsWithHTTPClient(client *http.Client) *ReplaceHTTPAfterResponseRuleParams {
	return &ReplaceHTTPAfterResponseRuleParams{
		HTTPClient: client,
	}
}

/* ReplaceHTTPAfterResponseRuleParams contains all the parameters to send to the API endpoint
   for the replace HTTP after response rule operation.

   Typically these are written to a http.Request.
*/
type ReplaceHTTPAfterResponseRuleParams struct {

	// Data.
	Data *models.HTTPAfterResponseRule

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Index.

	   HTTP After Response Rule Index
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

// WithDefaults hydrates default values in the replace HTTP after response rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceHTTPAfterResponseRuleParams) WithDefaults() *ReplaceHTTPAfterResponseRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace HTTP after response rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceHTTPAfterResponseRuleParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := ReplaceHTTPAfterResponseRuleParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) WithTimeout(timeout time.Duration) *ReplaceHTTPAfterResponseRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) WithContext(ctx context.Context) *ReplaceHTTPAfterResponseRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) WithHTTPClient(client *http.Client) *ReplaceHTTPAfterResponseRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) WithData(data *models.HTTPAfterResponseRule) *ReplaceHTTPAfterResponseRuleParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) SetData(data *models.HTTPAfterResponseRule) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) WithForceReload(forceReload *bool) *ReplaceHTTPAfterResponseRuleParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithIndex adds the index to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) WithIndex(index int64) *ReplaceHTTPAfterResponseRuleParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) SetIndex(index int64) {
	o.Index = index
}

// WithParentName adds the parentName to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) WithParentName(parentName string) *ReplaceHTTPAfterResponseRuleParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) WithParentType(parentType string) *ReplaceHTTPAfterResponseRuleParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) WithTransactionID(transactionID *string) *ReplaceHTTPAfterResponseRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) WithVersion(version *int64) *ReplaceHTTPAfterResponseRuleParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace HTTP after response rule params
func (o *ReplaceHTTPAfterResponseRuleParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceHTTPAfterResponseRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
