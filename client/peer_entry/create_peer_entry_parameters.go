// Code generated by go-swagger; DO NOT EDIT.

package peer_entry

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

// NewCreatePeerEntryParams creates a new CreatePeerEntryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePeerEntryParams() *CreatePeerEntryParams {
	return &CreatePeerEntryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePeerEntryParamsWithTimeout creates a new CreatePeerEntryParams object
// with the ability to set a timeout on a request.
func NewCreatePeerEntryParamsWithTimeout(timeout time.Duration) *CreatePeerEntryParams {
	return &CreatePeerEntryParams{
		timeout: timeout,
	}
}

// NewCreatePeerEntryParamsWithContext creates a new CreatePeerEntryParams object
// with the ability to set a context for a request.
func NewCreatePeerEntryParamsWithContext(ctx context.Context) *CreatePeerEntryParams {
	return &CreatePeerEntryParams{
		Context: ctx,
	}
}

// NewCreatePeerEntryParamsWithHTTPClient creates a new CreatePeerEntryParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePeerEntryParamsWithHTTPClient(client *http.Client) *CreatePeerEntryParams {
	return &CreatePeerEntryParams{
		HTTPClient: client,
	}
}

/*
CreatePeerEntryParams contains all the parameters to send to the API endpoint

	for the create peer entry operation.

	Typically these are written to a http.Request.
*/
type CreatePeerEntryParams struct {

	// Data.
	Data *models.PeerEntry

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* PeerSection.

	   Parent peer section name
	*/
	PeerSection string

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

// WithDefaults hydrates default values in the create peer entry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePeerEntryParams) WithDefaults() *CreatePeerEntryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create peer entry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePeerEntryParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := CreatePeerEntryParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create peer entry params
func (o *CreatePeerEntryParams) WithTimeout(timeout time.Duration) *CreatePeerEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create peer entry params
func (o *CreatePeerEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create peer entry params
func (o *CreatePeerEntryParams) WithContext(ctx context.Context) *CreatePeerEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create peer entry params
func (o *CreatePeerEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create peer entry params
func (o *CreatePeerEntryParams) WithHTTPClient(client *http.Client) *CreatePeerEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create peer entry params
func (o *CreatePeerEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create peer entry params
func (o *CreatePeerEntryParams) WithData(data *models.PeerEntry) *CreatePeerEntryParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create peer entry params
func (o *CreatePeerEntryParams) SetData(data *models.PeerEntry) {
	o.Data = data
}

// WithForceReload adds the forceReload to the create peer entry params
func (o *CreatePeerEntryParams) WithForceReload(forceReload *bool) *CreatePeerEntryParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create peer entry params
func (o *CreatePeerEntryParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithPeerSection adds the peerSection to the create peer entry params
func (o *CreatePeerEntryParams) WithPeerSection(peerSection string) *CreatePeerEntryParams {
	o.SetPeerSection(peerSection)
	return o
}

// SetPeerSection adds the peerSection to the create peer entry params
func (o *CreatePeerEntryParams) SetPeerSection(peerSection string) {
	o.PeerSection = peerSection
}

// WithTransactionID adds the transactionID to the create peer entry params
func (o *CreatePeerEntryParams) WithTransactionID(transactionID *string) *CreatePeerEntryParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create peer entry params
func (o *CreatePeerEntryParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create peer entry params
func (o *CreatePeerEntryParams) WithVersion(version *int64) *CreatePeerEntryParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create peer entry params
func (o *CreatePeerEntryParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePeerEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param peer_section
	qrPeerSection := o.PeerSection
	qPeerSection := qrPeerSection
	if qPeerSection != "" {

		if err := r.SetQueryParam("peer_section", qPeerSection); err != nil {
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
