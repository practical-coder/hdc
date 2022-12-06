// Code generated by go-swagger; DO NOT EDIT.

package http_check

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

// NewCreateHTTPCheckParams creates a new CreateHTTPCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateHTTPCheckParams() *CreateHTTPCheckParams {
	return &CreateHTTPCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateHTTPCheckParamsWithTimeout creates a new CreateHTTPCheckParams object
// with the ability to set a timeout on a request.
func NewCreateHTTPCheckParamsWithTimeout(timeout time.Duration) *CreateHTTPCheckParams {
	return &CreateHTTPCheckParams{
		timeout: timeout,
	}
}

// NewCreateHTTPCheckParamsWithContext creates a new CreateHTTPCheckParams object
// with the ability to set a context for a request.
func NewCreateHTTPCheckParamsWithContext(ctx context.Context) *CreateHTTPCheckParams {
	return &CreateHTTPCheckParams{
		Context: ctx,
	}
}

// NewCreateHTTPCheckParamsWithHTTPClient creates a new CreateHTTPCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateHTTPCheckParamsWithHTTPClient(client *http.Client) *CreateHTTPCheckParams {
	return &CreateHTTPCheckParams{
		HTTPClient: client,
	}
}

/*
CreateHTTPCheckParams contains all the parameters to send to the API endpoint

	for the create HTTP check operation.

	Typically these are written to a http.Request.
*/
type CreateHTTPCheckParams struct {

	// Data.
	Data *models.HTTPCheck

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

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

// WithDefaults hydrates default values in the create HTTP check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateHTTPCheckParams) WithDefaults() *CreateHTTPCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create HTTP check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateHTTPCheckParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := CreateHTTPCheckParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create HTTP check params
func (o *CreateHTTPCheckParams) WithTimeout(timeout time.Duration) *CreateHTTPCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create HTTP check params
func (o *CreateHTTPCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create HTTP check params
func (o *CreateHTTPCheckParams) WithContext(ctx context.Context) *CreateHTTPCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create HTTP check params
func (o *CreateHTTPCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create HTTP check params
func (o *CreateHTTPCheckParams) WithHTTPClient(client *http.Client) *CreateHTTPCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create HTTP check params
func (o *CreateHTTPCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create HTTP check params
func (o *CreateHTTPCheckParams) WithData(data *models.HTTPCheck) *CreateHTTPCheckParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create HTTP check params
func (o *CreateHTTPCheckParams) SetData(data *models.HTTPCheck) {
	o.Data = data
}

// WithForceReload adds the forceReload to the create HTTP check params
func (o *CreateHTTPCheckParams) WithForceReload(forceReload *bool) *CreateHTTPCheckParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create HTTP check params
func (o *CreateHTTPCheckParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithParentName adds the parentName to the create HTTP check params
func (o *CreateHTTPCheckParams) WithParentName(parentName *string) *CreateHTTPCheckParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the create HTTP check params
func (o *CreateHTTPCheckParams) SetParentName(parentName *string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the create HTTP check params
func (o *CreateHTTPCheckParams) WithParentType(parentType string) *CreateHTTPCheckParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the create HTTP check params
func (o *CreateHTTPCheckParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the create HTTP check params
func (o *CreateHTTPCheckParams) WithTransactionID(transactionID *string) *CreateHTTPCheckParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create HTTP check params
func (o *CreateHTTPCheckParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create HTTP check params
func (o *CreateHTTPCheckParams) WithVersion(version *int64) *CreateHTTPCheckParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create HTTP check params
func (o *CreateHTTPCheckParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateHTTPCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
