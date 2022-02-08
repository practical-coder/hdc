// Code generated by go-swagger; DO NOT EDIT.

package defaults

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

	"github.com/haproxytech/client-native/v3/models"
)

// NewReplaceDefaultsParams creates a new ReplaceDefaultsParams object
// with the default values initialized.
func NewReplaceDefaultsParams() *ReplaceDefaultsParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceDefaultsParams{
		ForceReload: &forceReloadDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceDefaultsParamsWithTimeout creates a new ReplaceDefaultsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceDefaultsParamsWithTimeout(timeout time.Duration) *ReplaceDefaultsParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceDefaultsParams{
		ForceReload: &forceReloadDefault,

		timeout: timeout,
	}
}

// NewReplaceDefaultsParamsWithContext creates a new ReplaceDefaultsParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceDefaultsParamsWithContext(ctx context.Context) *ReplaceDefaultsParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceDefaultsParams{
		ForceReload: &forceReloadDefault,

		Context: ctx,
	}
}

// NewReplaceDefaultsParamsWithHTTPClient creates a new ReplaceDefaultsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceDefaultsParamsWithHTTPClient(client *http.Client) *ReplaceDefaultsParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceDefaultsParams{
		ForceReload: &forceReloadDefault,
		HTTPClient:  client,
	}
}

/*ReplaceDefaultsParams contains all the parameters to send to the API endpoint
for the replace defaults operation typically these are written to a http.Request
*/
type ReplaceDefaultsParams struct {

	/*Data*/
	Data *models.Defaults
	/*ForceReload
	  If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.

	*/
	ForceReload *bool
	/*TransactionID
	  ID of the transaction where we want to add the operation. Cannot be used when version is specified.

	*/
	TransactionID *string
	/*Version
	  Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.

	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace defaults params
func (o *ReplaceDefaultsParams) WithTimeout(timeout time.Duration) *ReplaceDefaultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace defaults params
func (o *ReplaceDefaultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace defaults params
func (o *ReplaceDefaultsParams) WithContext(ctx context.Context) *ReplaceDefaultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace defaults params
func (o *ReplaceDefaultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace defaults params
func (o *ReplaceDefaultsParams) WithHTTPClient(client *http.Client) *ReplaceDefaultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace defaults params
func (o *ReplaceDefaultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace defaults params
func (o *ReplaceDefaultsParams) WithData(data *models.Defaults) *ReplaceDefaultsParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace defaults params
func (o *ReplaceDefaultsParams) SetData(data *models.Defaults) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace defaults params
func (o *ReplaceDefaultsParams) WithForceReload(forceReload *bool) *ReplaceDefaultsParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace defaults params
func (o *ReplaceDefaultsParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithTransactionID adds the transactionID to the replace defaults params
func (o *ReplaceDefaultsParams) WithTransactionID(transactionID *string) *ReplaceDefaultsParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace defaults params
func (o *ReplaceDefaultsParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace defaults params
func (o *ReplaceDefaultsParams) WithVersion(version *int64) *ReplaceDefaultsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace defaults params
func (o *ReplaceDefaultsParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceDefaultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
