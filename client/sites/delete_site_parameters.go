// Code generated by go-swagger; DO NOT EDIT.

package sites

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
)

// NewDeleteSiteParams creates a new DeleteSiteParams object
// with the default values initialized.
func NewDeleteSiteParams() *DeleteSiteParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &DeleteSiteParams{
		ForceReload: &forceReloadDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSiteParamsWithTimeout creates a new DeleteSiteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSiteParamsWithTimeout(timeout time.Duration) *DeleteSiteParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &DeleteSiteParams{
		ForceReload: &forceReloadDefault,

		timeout: timeout,
	}
}

// NewDeleteSiteParamsWithContext creates a new DeleteSiteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSiteParamsWithContext(ctx context.Context) *DeleteSiteParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &DeleteSiteParams{
		ForceReload: &forceReloadDefault,

		Context: ctx,
	}
}

// NewDeleteSiteParamsWithHTTPClient creates a new DeleteSiteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSiteParamsWithHTTPClient(client *http.Client) *DeleteSiteParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &DeleteSiteParams{
		ForceReload: &forceReloadDefault,
		HTTPClient:  client,
	}
}

/*DeleteSiteParams contains all the parameters to send to the API endpoint
for the delete site operation typically these are written to a http.Request
*/
type DeleteSiteParams struct {

	/*ForceReload
	  If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.

	*/
	ForceReload *bool
	/*Name
	  Site frontend name

	*/
	Name string
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

// WithTimeout adds the timeout to the delete site params
func (o *DeleteSiteParams) WithTimeout(timeout time.Duration) *DeleteSiteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete site params
func (o *DeleteSiteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete site params
func (o *DeleteSiteParams) WithContext(ctx context.Context) *DeleteSiteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete site params
func (o *DeleteSiteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete site params
func (o *DeleteSiteParams) WithHTTPClient(client *http.Client) *DeleteSiteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete site params
func (o *DeleteSiteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceReload adds the forceReload to the delete site params
func (o *DeleteSiteParams) WithForceReload(forceReload *bool) *DeleteSiteParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the delete site params
func (o *DeleteSiteParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithName adds the name to the delete site params
func (o *DeleteSiteParams) WithName(name string) *DeleteSiteParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete site params
func (o *DeleteSiteParams) SetName(name string) {
	o.Name = name
}

// WithTransactionID adds the transactionID to the delete site params
func (o *DeleteSiteParams) WithTransactionID(transactionID *string) *DeleteSiteParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the delete site params
func (o *DeleteSiteParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the delete site params
func (o *DeleteSiteParams) WithVersion(version *int64) *DeleteSiteParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete site params
func (o *DeleteSiteParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSiteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
