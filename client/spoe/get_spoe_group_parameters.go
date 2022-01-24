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
)

// NewGetSpoeGroupParams creates a new GetSpoeGroupParams object
// with the default values initialized.
func NewGetSpoeGroupParams() *GetSpoeGroupParams {
	var ()
	return &GetSpoeGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpoeGroupParamsWithTimeout creates a new GetSpoeGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSpoeGroupParamsWithTimeout(timeout time.Duration) *GetSpoeGroupParams {
	var ()
	return &GetSpoeGroupParams{

		timeout: timeout,
	}
}

// NewGetSpoeGroupParamsWithContext creates a new GetSpoeGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSpoeGroupParamsWithContext(ctx context.Context) *GetSpoeGroupParams {
	var ()
	return &GetSpoeGroupParams{

		Context: ctx,
	}
}

// NewGetSpoeGroupParamsWithHTTPClient creates a new GetSpoeGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSpoeGroupParamsWithHTTPClient(client *http.Client) *GetSpoeGroupParams {
	var ()
	return &GetSpoeGroupParams{
		HTTPClient: client,
	}
}

/*GetSpoeGroupParams contains all the parameters to send to the API endpoint
for the get spoe group operation typically these are written to a http.Request
*/
type GetSpoeGroupParams struct {

	/*Name
	  Spoe group name

	*/
	Name string
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get spoe group params
func (o *GetSpoeGroupParams) WithTimeout(timeout time.Duration) *GetSpoeGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get spoe group params
func (o *GetSpoeGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get spoe group params
func (o *GetSpoeGroupParams) WithContext(ctx context.Context) *GetSpoeGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get spoe group params
func (o *GetSpoeGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get spoe group params
func (o *GetSpoeGroupParams) WithHTTPClient(client *http.Client) *GetSpoeGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get spoe group params
func (o *GetSpoeGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get spoe group params
func (o *GetSpoeGroupParams) WithName(name string) *GetSpoeGroupParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get spoe group params
func (o *GetSpoeGroupParams) SetName(name string) {
	o.Name = name
}

// WithScope adds the scope to the get spoe group params
func (o *GetSpoeGroupParams) WithScope(scope string) *GetSpoeGroupParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get spoe group params
func (o *GetSpoeGroupParams) SetScope(scope string) {
	o.Scope = scope
}

// WithSpoe adds the spoe to the get spoe group params
func (o *GetSpoeGroupParams) WithSpoe(spoe string) *GetSpoeGroupParams {
	o.SetSpoe(spoe)
	return o
}

// SetSpoe adds the spoe to the get spoe group params
func (o *GetSpoeGroupParams) SetSpoe(spoe string) {
	o.Spoe = spoe
}

// WithTransactionID adds the transactionID to the get spoe group params
func (o *GetSpoeGroupParams) WithTransactionID(transactionID *string) *GetSpoeGroupParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get spoe group params
func (o *GetSpoeGroupParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpoeGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
