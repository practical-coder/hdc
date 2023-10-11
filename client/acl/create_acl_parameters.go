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

// NewCreateACLParams creates a new CreateACLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateACLParams() *CreateACLParams {
	return &CreateACLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateACLParamsWithTimeout creates a new CreateACLParams object
// with the ability to set a timeout on a request.
func NewCreateACLParamsWithTimeout(timeout time.Duration) *CreateACLParams {
	return &CreateACLParams{
		timeout: timeout,
	}
}

// NewCreateACLParamsWithContext creates a new CreateACLParams object
// with the ability to set a context for a request.
func NewCreateACLParamsWithContext(ctx context.Context) *CreateACLParams {
	return &CreateACLParams{
		Context: ctx,
	}
}

// NewCreateACLParamsWithHTTPClient creates a new CreateACLParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateACLParamsWithHTTPClient(client *http.Client) *CreateACLParams {
	return &CreateACLParams{
		HTTPClient: client,
	}
}

/*
CreateACLParams contains all the parameters to send to the API endpoint

	for the create Acl operation.

	Typically these are written to a http.Request.
*/
type CreateACLParams struct {

	// Data.
	Data *models.ACL

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

// WithDefaults hydrates default values in the create Acl params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateACLParams) WithDefaults() *CreateACLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create Acl params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateACLParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := CreateACLParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create Acl params
func (o *CreateACLParams) WithTimeout(timeout time.Duration) *CreateACLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create Acl params
func (o *CreateACLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create Acl params
func (o *CreateACLParams) WithContext(ctx context.Context) *CreateACLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create Acl params
func (o *CreateACLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create Acl params
func (o *CreateACLParams) WithHTTPClient(client *http.Client) *CreateACLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create Acl params
func (o *CreateACLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create Acl params
func (o *CreateACLParams) WithData(data *models.ACL) *CreateACLParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create Acl params
func (o *CreateACLParams) SetData(data *models.ACL) {
	o.Data = data
}

// WithForceReload adds the forceReload to the create Acl params
func (o *CreateACLParams) WithForceReload(forceReload *bool) *CreateACLParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create Acl params
func (o *CreateACLParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithParentName adds the parentName to the create Acl params
func (o *CreateACLParams) WithParentName(parentName string) *CreateACLParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the create Acl params
func (o *CreateACLParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the create Acl params
func (o *CreateACLParams) WithParentType(parentType string) *CreateACLParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the create Acl params
func (o *CreateACLParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the create Acl params
func (o *CreateACLParams) WithTransactionID(transactionID *string) *CreateACLParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create Acl params
func (o *CreateACLParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create Acl params
func (o *CreateACLParams) WithVersion(version *int64) *CreateACLParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create Acl params
func (o *CreateACLParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateACLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
