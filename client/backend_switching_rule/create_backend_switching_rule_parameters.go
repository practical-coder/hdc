// Code generated by go-swagger; DO NOT EDIT.

package backend_switching_rule

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

	"github.com/haproxytech/client-native/v2/models"
)

// NewCreateBackendSwitchingRuleParams creates a new CreateBackendSwitchingRuleParams object
// with the default values initialized.
func NewCreateBackendSwitchingRuleParams() *CreateBackendSwitchingRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateBackendSwitchingRuleParams{
		ForceReload: &forceReloadDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBackendSwitchingRuleParamsWithTimeout creates a new CreateBackendSwitchingRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateBackendSwitchingRuleParamsWithTimeout(timeout time.Duration) *CreateBackendSwitchingRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateBackendSwitchingRuleParams{
		ForceReload: &forceReloadDefault,

		timeout: timeout,
	}
}

// NewCreateBackendSwitchingRuleParamsWithContext creates a new CreateBackendSwitchingRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateBackendSwitchingRuleParamsWithContext(ctx context.Context) *CreateBackendSwitchingRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateBackendSwitchingRuleParams{
		ForceReload: &forceReloadDefault,

		Context: ctx,
	}
}

// NewCreateBackendSwitchingRuleParamsWithHTTPClient creates a new CreateBackendSwitchingRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateBackendSwitchingRuleParamsWithHTTPClient(client *http.Client) *CreateBackendSwitchingRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateBackendSwitchingRuleParams{
		ForceReload: &forceReloadDefault,
		HTTPClient:  client,
	}
}

/*CreateBackendSwitchingRuleParams contains all the parameters to send to the API endpoint
for the create backend switching rule operation typically these are written to a http.Request
*/
type CreateBackendSwitchingRuleParams struct {

	/*Data*/
	Data *models.BackendSwitchingRule
	/*ForceReload
	  If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.

	*/
	ForceReload *bool
	/*Frontend
	  Frontend name

	*/
	Frontend string
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

// WithTimeout adds the timeout to the create backend switching rule params
func (o *CreateBackendSwitchingRuleParams) WithTimeout(timeout time.Duration) *CreateBackendSwitchingRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create backend switching rule params
func (o *CreateBackendSwitchingRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create backend switching rule params
func (o *CreateBackendSwitchingRuleParams) WithContext(ctx context.Context) *CreateBackendSwitchingRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create backend switching rule params
func (o *CreateBackendSwitchingRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create backend switching rule params
func (o *CreateBackendSwitchingRuleParams) WithHTTPClient(client *http.Client) *CreateBackendSwitchingRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create backend switching rule params
func (o *CreateBackendSwitchingRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create backend switching rule params
func (o *CreateBackendSwitchingRuleParams) WithData(data *models.BackendSwitchingRule) *CreateBackendSwitchingRuleParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create backend switching rule params
func (o *CreateBackendSwitchingRuleParams) SetData(data *models.BackendSwitchingRule) {
	o.Data = data
}

// WithForceReload adds the forceReload to the create backend switching rule params
func (o *CreateBackendSwitchingRuleParams) WithForceReload(forceReload *bool) *CreateBackendSwitchingRuleParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create backend switching rule params
func (o *CreateBackendSwitchingRuleParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithFrontend adds the frontend to the create backend switching rule params
func (o *CreateBackendSwitchingRuleParams) WithFrontend(frontend string) *CreateBackendSwitchingRuleParams {
	o.SetFrontend(frontend)
	return o
}

// SetFrontend adds the frontend to the create backend switching rule params
func (o *CreateBackendSwitchingRuleParams) SetFrontend(frontend string) {
	o.Frontend = frontend
}

// WithTransactionID adds the transactionID to the create backend switching rule params
func (o *CreateBackendSwitchingRuleParams) WithTransactionID(transactionID *string) *CreateBackendSwitchingRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create backend switching rule params
func (o *CreateBackendSwitchingRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create backend switching rule params
func (o *CreateBackendSwitchingRuleParams) WithVersion(version *int64) *CreateBackendSwitchingRuleParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create backend switching rule params
func (o *CreateBackendSwitchingRuleParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBackendSwitchingRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param frontend
	qrFrontend := o.Frontend
	qFrontend := qrFrontend
	if qFrontend != "" {
		if err := r.SetQueryParam("frontend", qFrontend); err != nil {
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
