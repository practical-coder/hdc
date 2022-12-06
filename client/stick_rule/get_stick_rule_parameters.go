// Code generated by go-swagger; DO NOT EDIT.

package stick_rule

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

// NewGetStickRuleParams creates a new GetStickRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStickRuleParams() *GetStickRuleParams {
	return &GetStickRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStickRuleParamsWithTimeout creates a new GetStickRuleParams object
// with the ability to set a timeout on a request.
func NewGetStickRuleParamsWithTimeout(timeout time.Duration) *GetStickRuleParams {
	return &GetStickRuleParams{
		timeout: timeout,
	}
}

// NewGetStickRuleParamsWithContext creates a new GetStickRuleParams object
// with the ability to set a context for a request.
func NewGetStickRuleParamsWithContext(ctx context.Context) *GetStickRuleParams {
	return &GetStickRuleParams{
		Context: ctx,
	}
}

// NewGetStickRuleParamsWithHTTPClient creates a new GetStickRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStickRuleParamsWithHTTPClient(client *http.Client) *GetStickRuleParams {
	return &GetStickRuleParams{
		HTTPClient: client,
	}
}

/*
GetStickRuleParams contains all the parameters to send to the API endpoint

	for the get stick rule operation.

	Typically these are written to a http.Request.
*/
type GetStickRuleParams struct {

	/* Backend.

	   Backend name
	*/
	Backend string

	/* Index.

	   Stick Rule Index
	*/
	Index int64

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get stick rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStickRuleParams) WithDefaults() *GetStickRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get stick rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStickRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get stick rule params
func (o *GetStickRuleParams) WithTimeout(timeout time.Duration) *GetStickRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stick rule params
func (o *GetStickRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stick rule params
func (o *GetStickRuleParams) WithContext(ctx context.Context) *GetStickRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stick rule params
func (o *GetStickRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stick rule params
func (o *GetStickRuleParams) WithHTTPClient(client *http.Client) *GetStickRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stick rule params
func (o *GetStickRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the get stick rule params
func (o *GetStickRuleParams) WithBackend(backend string) *GetStickRuleParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the get stick rule params
func (o *GetStickRuleParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithIndex adds the index to the get stick rule params
func (o *GetStickRuleParams) WithIndex(index int64) *GetStickRuleParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the get stick rule params
func (o *GetStickRuleParams) SetIndex(index int64) {
	o.Index = index
}

// WithTransactionID adds the transactionID to the get stick rule params
func (o *GetStickRuleParams) WithTransactionID(transactionID *string) *GetStickRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get stick rule params
func (o *GetStickRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStickRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param backend
	qrBackend := o.Backend
	qBackend := qrBackend
	if qBackend != "" {

		if err := r.SetQueryParam("backend", qBackend); err != nil {
			return err
		}
	}

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
