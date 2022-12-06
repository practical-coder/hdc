// Code generated by go-swagger; DO NOT EDIT.

package peer

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

// NewGetPeerSectionParams creates a new GetPeerSectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPeerSectionParams() *GetPeerSectionParams {
	return &GetPeerSectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPeerSectionParamsWithTimeout creates a new GetPeerSectionParams object
// with the ability to set a timeout on a request.
func NewGetPeerSectionParamsWithTimeout(timeout time.Duration) *GetPeerSectionParams {
	return &GetPeerSectionParams{
		timeout: timeout,
	}
}

// NewGetPeerSectionParamsWithContext creates a new GetPeerSectionParams object
// with the ability to set a context for a request.
func NewGetPeerSectionParamsWithContext(ctx context.Context) *GetPeerSectionParams {
	return &GetPeerSectionParams{
		Context: ctx,
	}
}

// NewGetPeerSectionParamsWithHTTPClient creates a new GetPeerSectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPeerSectionParamsWithHTTPClient(client *http.Client) *GetPeerSectionParams {
	return &GetPeerSectionParams{
		HTTPClient: client,
	}
}

/*
GetPeerSectionParams contains all the parameters to send to the API endpoint

	for the get peer section operation.

	Typically these are written to a http.Request.
*/
type GetPeerSectionParams struct {

	/* Name.

	   Peer name
	*/
	Name string

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get peer section params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPeerSectionParams) WithDefaults() *GetPeerSectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get peer section params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPeerSectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get peer section params
func (o *GetPeerSectionParams) WithTimeout(timeout time.Duration) *GetPeerSectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get peer section params
func (o *GetPeerSectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get peer section params
func (o *GetPeerSectionParams) WithContext(ctx context.Context) *GetPeerSectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get peer section params
func (o *GetPeerSectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get peer section params
func (o *GetPeerSectionParams) WithHTTPClient(client *http.Client) *GetPeerSectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get peer section params
func (o *GetPeerSectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get peer section params
func (o *GetPeerSectionParams) WithName(name string) *GetPeerSectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get peer section params
func (o *GetPeerSectionParams) SetName(name string) {
	o.Name = name
}

// WithTransactionID adds the transactionID to the get peer section params
func (o *GetPeerSectionParams) WithTransactionID(transactionID *string) *GetPeerSectionParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get peer section params
func (o *GetPeerSectionParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPeerSectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
