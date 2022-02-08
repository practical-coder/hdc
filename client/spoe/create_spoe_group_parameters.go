// Code generated by go-swagger; DO NOT EDIT.

package spoe

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

// NewCreateSpoeGroupParams creates a new CreateSpoeGroupParams object
// with the default values initialized.
func NewCreateSpoeGroupParams() *CreateSpoeGroupParams {
	var ()
	return &CreateSpoeGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSpoeGroupParamsWithTimeout creates a new CreateSpoeGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSpoeGroupParamsWithTimeout(timeout time.Duration) *CreateSpoeGroupParams {
	var ()
	return &CreateSpoeGroupParams{

		timeout: timeout,
	}
}

// NewCreateSpoeGroupParamsWithContext creates a new CreateSpoeGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSpoeGroupParamsWithContext(ctx context.Context) *CreateSpoeGroupParams {
	var ()
	return &CreateSpoeGroupParams{

		Context: ctx,
	}
}

// NewCreateSpoeGroupParamsWithHTTPClient creates a new CreateSpoeGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSpoeGroupParamsWithHTTPClient(client *http.Client) *CreateSpoeGroupParams {
	var ()
	return &CreateSpoeGroupParams{
		HTTPClient: client,
	}
}

/*CreateSpoeGroupParams contains all the parameters to send to the API endpoint
for the create spoe group operation typically these are written to a http.Request
*/
type CreateSpoeGroupParams struct {

	/*Data*/
	Data *models.SpoeGroup
	/*Scope
	  Spoe scope

	*/
	Scope string
	/*Spoe
	  Spoe file name

	*/
	Spoe string
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

// WithTimeout adds the timeout to the create spoe group params
func (o *CreateSpoeGroupParams) WithTimeout(timeout time.Duration) *CreateSpoeGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create spoe group params
func (o *CreateSpoeGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create spoe group params
func (o *CreateSpoeGroupParams) WithContext(ctx context.Context) *CreateSpoeGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create spoe group params
func (o *CreateSpoeGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create spoe group params
func (o *CreateSpoeGroupParams) WithHTTPClient(client *http.Client) *CreateSpoeGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create spoe group params
func (o *CreateSpoeGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create spoe group params
func (o *CreateSpoeGroupParams) WithData(data *models.SpoeGroup) *CreateSpoeGroupParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create spoe group params
func (o *CreateSpoeGroupParams) SetData(data *models.SpoeGroup) {
	o.Data = data
}

// WithScope adds the scope to the create spoe group params
func (o *CreateSpoeGroupParams) WithScope(scope string) *CreateSpoeGroupParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the create spoe group params
func (o *CreateSpoeGroupParams) SetScope(scope string) {
	o.Scope = scope
}

// WithSpoe adds the spoe to the create spoe group params
func (o *CreateSpoeGroupParams) WithSpoe(spoe string) *CreateSpoeGroupParams {
	o.SetSpoe(spoe)
	return o
}

// SetSpoe adds the spoe to the create spoe group params
func (o *CreateSpoeGroupParams) SetSpoe(spoe string) {
	o.Spoe = spoe
}

// WithTransactionID adds the transactionID to the create spoe group params
func (o *CreateSpoeGroupParams) WithTransactionID(transactionID *string) *CreateSpoeGroupParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create spoe group params
func (o *CreateSpoeGroupParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create spoe group params
func (o *CreateSpoeGroupParams) WithVersion(version *int64) *CreateSpoeGroupParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create spoe group params
func (o *CreateSpoeGroupParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSpoeGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// query param scope
	qrScope := o.Scope
	qScope := qrScope
	if qScope != "" {
		if err := r.SetQueryParam("scope", qScope); err != nil {
			return err
		}
	}

	// query param spoe
	qrSpoe := o.Spoe
	qSpoe := qrSpoe
	if qSpoe != "" {
		if err := r.SetQueryParam("spoe", qSpoe); err != nil {
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
