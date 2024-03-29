// Code generated by go-swagger; DO NOT EDIT.

package mailer_entry

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

// NewReplaceMailerEntryParams creates a new ReplaceMailerEntryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceMailerEntryParams() *ReplaceMailerEntryParams {
	return &ReplaceMailerEntryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceMailerEntryParamsWithTimeout creates a new ReplaceMailerEntryParams object
// with the ability to set a timeout on a request.
func NewReplaceMailerEntryParamsWithTimeout(timeout time.Duration) *ReplaceMailerEntryParams {
	return &ReplaceMailerEntryParams{
		timeout: timeout,
	}
}

// NewReplaceMailerEntryParamsWithContext creates a new ReplaceMailerEntryParams object
// with the ability to set a context for a request.
func NewReplaceMailerEntryParamsWithContext(ctx context.Context) *ReplaceMailerEntryParams {
	return &ReplaceMailerEntryParams{
		Context: ctx,
	}
}

// NewReplaceMailerEntryParamsWithHTTPClient creates a new ReplaceMailerEntryParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceMailerEntryParamsWithHTTPClient(client *http.Client) *ReplaceMailerEntryParams {
	return &ReplaceMailerEntryParams{
		HTTPClient: client,
	}
}

/*
ReplaceMailerEntryParams contains all the parameters to send to the API endpoint

	for the replace mailer entry operation.

	Typically these are written to a http.Request.
*/
type ReplaceMailerEntryParams struct {

	// Data.
	Data *models.MailerEntry

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* MailersSection.

	   Parent mailers section name
	*/
	MailersSection string

	/* Name.

	   MailerEntry name
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

// WithDefaults hydrates default values in the replace mailer entry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceMailerEntryParams) WithDefaults() *ReplaceMailerEntryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace mailer entry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceMailerEntryParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := ReplaceMailerEntryParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the replace mailer entry params
func (o *ReplaceMailerEntryParams) WithTimeout(timeout time.Duration) *ReplaceMailerEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace mailer entry params
func (o *ReplaceMailerEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace mailer entry params
func (o *ReplaceMailerEntryParams) WithContext(ctx context.Context) *ReplaceMailerEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace mailer entry params
func (o *ReplaceMailerEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace mailer entry params
func (o *ReplaceMailerEntryParams) WithHTTPClient(client *http.Client) *ReplaceMailerEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace mailer entry params
func (o *ReplaceMailerEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace mailer entry params
func (o *ReplaceMailerEntryParams) WithData(data *models.MailerEntry) *ReplaceMailerEntryParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace mailer entry params
func (o *ReplaceMailerEntryParams) SetData(data *models.MailerEntry) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace mailer entry params
func (o *ReplaceMailerEntryParams) WithForceReload(forceReload *bool) *ReplaceMailerEntryParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace mailer entry params
func (o *ReplaceMailerEntryParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithMailersSection adds the mailersSection to the replace mailer entry params
func (o *ReplaceMailerEntryParams) WithMailersSection(mailersSection string) *ReplaceMailerEntryParams {
	o.SetMailersSection(mailersSection)
	return o
}

// SetMailersSection adds the mailersSection to the replace mailer entry params
func (o *ReplaceMailerEntryParams) SetMailersSection(mailersSection string) {
	o.MailersSection = mailersSection
}

// WithName adds the name to the replace mailer entry params
func (o *ReplaceMailerEntryParams) WithName(name string) *ReplaceMailerEntryParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace mailer entry params
func (o *ReplaceMailerEntryParams) SetName(name string) {
	o.Name = name
}

// WithTransactionID adds the transactionID to the replace mailer entry params
func (o *ReplaceMailerEntryParams) WithTransactionID(transactionID *string) *ReplaceMailerEntryParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace mailer entry params
func (o *ReplaceMailerEntryParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace mailer entry params
func (o *ReplaceMailerEntryParams) WithVersion(version *int64) *ReplaceMailerEntryParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace mailer entry params
func (o *ReplaceMailerEntryParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceMailerEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param mailers_section
	qrMailersSection := o.MailersSection
	qMailersSection := qrMailersSection
	if qMailersSection != "" {

		if err := r.SetQueryParam("mailers_section", qMailersSection); err != nil {
			return err
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
