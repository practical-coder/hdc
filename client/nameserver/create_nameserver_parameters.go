// Code generated by go-swagger; DO NOT EDIT.

package nameserver

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

	"client-native/models"
)

// NewCreateNameserverParams creates a new CreateNameserverParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNameserverParams() *CreateNameserverParams {
	return &CreateNameserverParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNameserverParamsWithTimeout creates a new CreateNameserverParams object
// with the ability to set a timeout on a request.
func NewCreateNameserverParamsWithTimeout(timeout time.Duration) *CreateNameserverParams {
	return &CreateNameserverParams{
		timeout: timeout,
	}
}

// NewCreateNameserverParamsWithContext creates a new CreateNameserverParams object
// with the ability to set a context for a request.
func NewCreateNameserverParamsWithContext(ctx context.Context) *CreateNameserverParams {
	return &CreateNameserverParams{
		Context: ctx,
	}
}

// NewCreateNameserverParamsWithHTTPClient creates a new CreateNameserverParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNameserverParamsWithHTTPClient(client *http.Client) *CreateNameserverParams {
	return &CreateNameserverParams{
		HTTPClient: client,
	}
}

/* CreateNameserverParams contains all the parameters to send to the API endpoint
   for the create nameserver operation.

   Typically these are written to a http.Request.
*/
type CreateNameserverParams struct {

	// Data.
	Data *models.Nameserver

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Resolver.

	   Parent resolver name
	*/
	Resolver string

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

// WithDefaults hydrates default values in the create nameserver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNameserverParams) WithDefaults() *CreateNameserverParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create nameserver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNameserverParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := CreateNameserverParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create nameserver params
func (o *CreateNameserverParams) WithTimeout(timeout time.Duration) *CreateNameserverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create nameserver params
func (o *CreateNameserverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create nameserver params
func (o *CreateNameserverParams) WithContext(ctx context.Context) *CreateNameserverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create nameserver params
func (o *CreateNameserverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create nameserver params
func (o *CreateNameserverParams) WithHTTPClient(client *http.Client) *CreateNameserverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create nameserver params
func (o *CreateNameserverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create nameserver params
func (o *CreateNameserverParams) WithData(data *models.Nameserver) *CreateNameserverParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create nameserver params
func (o *CreateNameserverParams) SetData(data *models.Nameserver) {
	o.Data = data
}

// WithForceReload adds the forceReload to the create nameserver params
func (o *CreateNameserverParams) WithForceReload(forceReload *bool) *CreateNameserverParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create nameserver params
func (o *CreateNameserverParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithResolver adds the resolver to the create nameserver params
func (o *CreateNameserverParams) WithResolver(resolver string) *CreateNameserverParams {
	o.SetResolver(resolver)
	return o
}

// SetResolver adds the resolver to the create nameserver params
func (o *CreateNameserverParams) SetResolver(resolver string) {
	o.Resolver = resolver
}

// WithTransactionID adds the transactionID to the create nameserver params
func (o *CreateNameserverParams) WithTransactionID(transactionID *string) *CreateNameserverParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create nameserver params
func (o *CreateNameserverParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create nameserver params
func (o *CreateNameserverParams) WithVersion(version *int64) *CreateNameserverParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create nameserver params
func (o *CreateNameserverParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNameserverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param resolver
	qrResolver := o.Resolver
	qResolver := qrResolver
	if qResolver != "" {

		if err := r.SetQueryParam("resolver", qResolver); err != nil {
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
