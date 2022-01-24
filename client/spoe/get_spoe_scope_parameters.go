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

// NewGetSpoeScopeParams creates a new GetSpoeScopeParams object
// with the default values initialized.
func NewGetSpoeScopeParams() *GetSpoeScopeParams {
	var ()
	return &GetSpoeScopeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpoeScopeParamsWithTimeout creates a new GetSpoeScopeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSpoeScopeParamsWithTimeout(timeout time.Duration) *GetSpoeScopeParams {
	var ()
	return &GetSpoeScopeParams{

		timeout: timeout,
	}
}

// NewGetSpoeScopeParamsWithContext creates a new GetSpoeScopeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSpoeScopeParamsWithContext(ctx context.Context) *GetSpoeScopeParams {
	var ()
	return &GetSpoeScopeParams{

		Context: ctx,
	}
}

// NewGetSpoeScopeParamsWithHTTPClient creates a new GetSpoeScopeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSpoeScopeParamsWithHTTPClient(client *http.Client) *GetSpoeScopeParams {
	var ()
	return &GetSpoeScopeParams{
		HTTPClient: client,
	}
}

/*GetSpoeScopeParams contains all the parameters to send to the API endpoint
for the get spoe scope operation typically these are written to a http.Request
*/
type GetSpoeScopeParams struct {

	/*Name
	  Spoe scope

	*/
	Name string
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

// WithTimeout adds the timeout to the get spoe scope params
func (o *GetSpoeScopeParams) WithTimeout(timeout time.Duration) *GetSpoeScopeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get spoe scope params
func (o *GetSpoeScopeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get spoe scope params
func (o *GetSpoeScopeParams) WithContext(ctx context.Context) *GetSpoeScopeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get spoe scope params
func (o *GetSpoeScopeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get spoe scope params
func (o *GetSpoeScopeParams) WithHTTPClient(client *http.Client) *GetSpoeScopeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get spoe scope params
func (o *GetSpoeScopeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get spoe scope params
func (o *GetSpoeScopeParams) WithName(name string) *GetSpoeScopeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get spoe scope params
func (o *GetSpoeScopeParams) SetName(name string) {
	o.Name = name
}

// WithSpoe adds the spoe to the get spoe scope params
func (o *GetSpoeScopeParams) WithSpoe(spoe string) *GetSpoeScopeParams {
	o.SetSpoe(spoe)
	return o
}

// SetSpoe adds the spoe to the get spoe scope params
func (o *GetSpoeScopeParams) SetSpoe(spoe string) {
	o.Spoe = spoe
}

// WithTransactionID adds the transactionID to the get spoe scope params
func (o *GetSpoeScopeParams) WithTransactionID(transactionID *string) *GetSpoeScopeParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get spoe scope params
func (o *GetSpoeScopeParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpoeScopeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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
