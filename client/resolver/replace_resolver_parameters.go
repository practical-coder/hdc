// Code generated by go-swagger; DO NOT EDIT.

package resolver

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

// NewReplaceResolverParams creates a new ReplaceResolverParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceResolverParams() *ReplaceResolverParams {
	return &ReplaceResolverParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceResolverParamsWithTimeout creates a new ReplaceResolverParams object
// with the ability to set a timeout on a request.
func NewReplaceResolverParamsWithTimeout(timeout time.Duration) *ReplaceResolverParams {
	return &ReplaceResolverParams{
		timeout: timeout,
	}
}

// NewReplaceResolverParamsWithContext creates a new ReplaceResolverParams object
// with the ability to set a context for a request.
func NewReplaceResolverParamsWithContext(ctx context.Context) *ReplaceResolverParams {
	return &ReplaceResolverParams{
		Context: ctx,
	}
}

// NewReplaceResolverParamsWithHTTPClient creates a new ReplaceResolverParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceResolverParamsWithHTTPClient(client *http.Client) *ReplaceResolverParams {
	return &ReplaceResolverParams{
		HTTPClient: client,
	}
}

/* ReplaceResolverParams contains all the parameters to send to the API endpoint
   for the replace resolver operation.

   Typically these are written to a http.Request.
*/
type ReplaceResolverParams struct {

	// Data.
	Data *models.Resolver

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Name.

	   Resolver name
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

// WithDefaults hydrates default values in the replace resolver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceResolverParams) WithDefaults() *ReplaceResolverParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace resolver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceResolverParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := ReplaceResolverParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the replace resolver params
func (o *ReplaceResolverParams) WithTimeout(timeout time.Duration) *ReplaceResolverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace resolver params
func (o *ReplaceResolverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace resolver params
func (o *ReplaceResolverParams) WithContext(ctx context.Context) *ReplaceResolverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace resolver params
func (o *ReplaceResolverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace resolver params
func (o *ReplaceResolverParams) WithHTTPClient(client *http.Client) *ReplaceResolverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace resolver params
func (o *ReplaceResolverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace resolver params
func (o *ReplaceResolverParams) WithData(data *models.Resolver) *ReplaceResolverParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace resolver params
func (o *ReplaceResolverParams) SetData(data *models.Resolver) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace resolver params
func (o *ReplaceResolverParams) WithForceReload(forceReload *bool) *ReplaceResolverParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace resolver params
func (o *ReplaceResolverParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithName adds the name to the replace resolver params
func (o *ReplaceResolverParams) WithName(name string) *ReplaceResolverParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace resolver params
func (o *ReplaceResolverParams) SetName(name string) {
	o.Name = name
}

// WithTransactionID adds the transactionID to the replace resolver params
func (o *ReplaceResolverParams) WithTransactionID(transactionID *string) *ReplaceResolverParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace resolver params
func (o *ReplaceResolverParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace resolver params
func (o *ReplaceResolverParams) WithVersion(version *int64) *ReplaceResolverParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace resolver params
func (o *ReplaceResolverParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceResolverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
