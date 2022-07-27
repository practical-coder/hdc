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

// NewGetSpoeAgentParams creates a new GetSpoeAgentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSpoeAgentParams() *GetSpoeAgentParams {
	return &GetSpoeAgentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpoeAgentParamsWithTimeout creates a new GetSpoeAgentParams object
// with the ability to set a timeout on a request.
func NewGetSpoeAgentParamsWithTimeout(timeout time.Duration) *GetSpoeAgentParams {
	return &GetSpoeAgentParams{
		timeout: timeout,
	}
}

// NewGetSpoeAgentParamsWithContext creates a new GetSpoeAgentParams object
// with the ability to set a context for a request.
func NewGetSpoeAgentParamsWithContext(ctx context.Context) *GetSpoeAgentParams {
	return &GetSpoeAgentParams{
		Context: ctx,
	}
}

// NewGetSpoeAgentParamsWithHTTPClient creates a new GetSpoeAgentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSpoeAgentParamsWithHTTPClient(client *http.Client) *GetSpoeAgentParams {
	return &GetSpoeAgentParams{
		HTTPClient: client,
	}
}

/* GetSpoeAgentParams contains all the parameters to send to the API endpoint
   for the get spoe agent operation.

   Typically these are written to a http.Request.
*/
type GetSpoeAgentParams struct {

	/* Name.

	   Spoe agent name
	*/
	Name string

	/* Scope.

	   Spoe scope
	*/
	Scope string

	/* Spoe.

	   Spoe file name
	*/
	Spoe string

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get spoe agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSpoeAgentParams) WithDefaults() *GetSpoeAgentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get spoe agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSpoeAgentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get spoe agent params
func (o *GetSpoeAgentParams) WithTimeout(timeout time.Duration) *GetSpoeAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get spoe agent params
func (o *GetSpoeAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get spoe agent params
func (o *GetSpoeAgentParams) WithContext(ctx context.Context) *GetSpoeAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get spoe agent params
func (o *GetSpoeAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get spoe agent params
func (o *GetSpoeAgentParams) WithHTTPClient(client *http.Client) *GetSpoeAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get spoe agent params
func (o *GetSpoeAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get spoe agent params
func (o *GetSpoeAgentParams) WithName(name string) *GetSpoeAgentParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get spoe agent params
func (o *GetSpoeAgentParams) SetName(name string) {
	o.Name = name
}

// WithScope adds the scope to the get spoe agent params
func (o *GetSpoeAgentParams) WithScope(scope string) *GetSpoeAgentParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get spoe agent params
func (o *GetSpoeAgentParams) SetScope(scope string) {
	o.Scope = scope
}

// WithSpoe adds the spoe to the get spoe agent params
func (o *GetSpoeAgentParams) WithSpoe(spoe string) *GetSpoeAgentParams {
	o.SetSpoe(spoe)
	return o
}

// SetSpoe adds the spoe to the get spoe agent params
func (o *GetSpoeAgentParams) SetSpoe(spoe string) {
	o.Spoe = spoe
}

// WithTransactionID adds the transactionID to the get spoe agent params
func (o *GetSpoeAgentParams) WithTransactionID(transactionID *string) *GetSpoeAgentParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get spoe agent params
func (o *GetSpoeAgentParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpoeAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
