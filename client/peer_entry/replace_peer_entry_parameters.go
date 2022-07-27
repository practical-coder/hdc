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

	"github.com/haproxytech/client-native/v4/models"
)

// NewReplacePeerEntryParams creates a new ReplacePeerEntryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplacePeerEntryParams() *ReplacePeerEntryParams {
	return &ReplacePeerEntryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplacePeerEntryParamsWithTimeout creates a new ReplacePeerEntryParams object
// with the ability to set a timeout on a request.
func NewReplacePeerEntryParamsWithTimeout(timeout time.Duration) *ReplacePeerEntryParams {
	return &ReplacePeerEntryParams{
		timeout: timeout,
	}
}

// NewReplacePeerEntryParamsWithContext creates a new ReplacePeerEntryParams object
// with the ability to set a context for a request.
func NewReplacePeerEntryParamsWithContext(ctx context.Context) *ReplacePeerEntryParams {
	return &ReplacePeerEntryParams{
		Context: ctx,
	}
}

// NewReplacePeerEntryParamsWithHTTPClient creates a new ReplacePeerEntryParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplacePeerEntryParamsWithHTTPClient(client *http.Client) *ReplacePeerEntryParams {
	return &ReplacePeerEntryParams{
		HTTPClient: client,
	}
}

/* ReplacePeerEntryParams contains all the parameters to send to the API endpoint
   for the replace peer entry operation.

   Typically these are written to a http.Request.
*/
type ReplacePeerEntryParams struct {

	// Data.
	Data *models.PeerEntry

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Name.

	   PeerEntry name
	*/
	Name string

	/* PeerSection.

	   Parent peers name
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

// WithDefaults hydrates default values in the replace peer entry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplacePeerEntryParams) WithDefaults() *ReplacePeerEntryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace peer entry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplacePeerEntryParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := ReplacePeerEntryParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the replace peer entry params
func (o *ReplacePeerEntryParams) WithTimeout(timeout time.Duration) *ReplacePeerEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace peer entry params
func (o *ReplacePeerEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace peer entry params
func (o *ReplacePeerEntryParams) WithContext(ctx context.Context) *ReplacePeerEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace peer entry params
func (o *ReplacePeerEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace peer entry params
func (o *ReplacePeerEntryParams) WithHTTPClient(client *http.Client) *ReplacePeerEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace peer entry params
func (o *ReplacePeerEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace peer entry params
func (o *ReplacePeerEntryParams) WithData(data *models.PeerEntry) *ReplacePeerEntryParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace peer entry params
func (o *ReplacePeerEntryParams) SetData(data *models.PeerEntry) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace peer entry params
func (o *ReplacePeerEntryParams) WithForceReload(forceReload *bool) *ReplacePeerEntryParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace peer entry params
func (o *ReplacePeerEntryParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithName adds the name to the replace peer entry params
func (o *ReplacePeerEntryParams) WithName(name string) *ReplacePeerEntryParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace peer entry params
func (o *ReplacePeerEntryParams) SetName(name string) {
	o.Name = name
}

// WithPeerSection adds the peerSection to the replace peer entry params
func (o *ReplacePeerEntryParams) WithPeerSection(peerSection string) *ReplacePeerEntryParams {
	o.SetPeerSection(peerSection)
	return o
}

// SetPeerSection adds the peerSection to the replace peer entry params
func (o *ReplacePeerEntryParams) SetPeerSection(peerSection string) {
	o.PeerSection = peerSection
}

// WithTransactionID adds the transactionID to the replace peer entry params
func (o *ReplacePeerEntryParams) WithTransactionID(transactionID *string) *ReplacePeerEntryParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace peer entry params
func (o *ReplacePeerEntryParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace peer entry params
func (o *ReplacePeerEntryParams) WithVersion(version *int64) *ReplacePeerEntryParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace peer entry params
func (o *ReplacePeerEntryParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplacePeerEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
