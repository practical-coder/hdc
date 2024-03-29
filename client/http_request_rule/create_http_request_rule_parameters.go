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

	"github.com/haproxytech/client-native/v5/models"
)

// NewCreateHTTPRequestRuleParams creates a new CreateHTTPRequestRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateHTTPRequestRuleParams() *CreateHTTPRequestRuleParams {
	return &CreateHTTPRequestRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateHTTPRequestRuleParamsWithTimeout creates a new CreateHTTPRequestRuleParams object
// with the ability to set a timeout on a request.
func NewCreateHTTPRequestRuleParamsWithTimeout(timeout time.Duration) *CreateHTTPRequestRuleParams {
	return &CreateHTTPRequestRuleParams{
		timeout: timeout,
	}
}

// NewCreateHTTPRequestRuleParamsWithContext creates a new CreateHTTPRequestRuleParams object
// with the ability to set a context for a request.
func NewCreateHTTPRequestRuleParamsWithContext(ctx context.Context) *CreateHTTPRequestRuleParams {
	return &CreateHTTPRequestRuleParams{
		Context: ctx,
	}
}

// NewCreateHTTPRequestRuleParamsWithHTTPClient creates a new CreateHTTPRequestRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateHTTPRequestRuleParamsWithHTTPClient(client *http.Client) *CreateHTTPRequestRuleParams {
	return &CreateHTTPRequestRuleParams{
		HTTPClient: client,
	}
}

/*
CreateHTTPRequestRuleParams contains all the parameters to send to the API endpoint

	for the create HTTP request rule operation.

	Typically these are written to a http.Request.
*/
type CreateHTTPRequestRuleParams struct {

	// Data.
	Data *models.HTTPRequestRule

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

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

// WithDefaults hydrates default values in the create HTTP request rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateHTTPRequestRuleParams) WithDefaults() *CreateHTTPRequestRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create HTTP request rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateHTTPRequestRuleParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := CreateHTTPRequestRuleParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) WithTimeout(timeout time.Duration) *CreateHTTPRequestRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) WithContext(ctx context.Context) *CreateHTTPRequestRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) WithHTTPClient(client *http.Client) *CreateHTTPRequestRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) WithData(data *models.HTTPRequestRule) *CreateHTTPRequestRuleParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) SetData(data *models.HTTPRequestRule) {
	o.Data = data
}

// WithForceReload adds the forceReload to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) WithForceReload(forceReload *bool) *CreateHTTPRequestRuleParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithParentName adds the parentName to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) WithParentName(parentName string) *CreateHTTPRequestRuleParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) WithParentType(parentType string) *CreateHTTPRequestRuleParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) WithTransactionID(transactionID *string) *CreateHTTPRequestRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) WithVersion(version *int64) *CreateHTTPRequestRuleParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create HTTP request rule params
func (o *CreateHTTPRequestRuleParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateHTTPRequestRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
