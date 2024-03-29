// Code generated by go-swagger; DO NOT EDIT.

package group

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

// NewGetGroupParams creates a new GetGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGroupParams() *GetGroupParams {
	return &GetGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupParamsWithTimeout creates a new GetGroupParams object
// with the ability to set a timeout on a request.
func NewGetGroupParamsWithTimeout(timeout time.Duration) *GetGroupParams {
	return &GetGroupParams{
		timeout: timeout,
	}
}

// NewGetGroupParamsWithContext creates a new GetGroupParams object
// with the ability to set a context for a request.
func NewGetGroupParamsWithContext(ctx context.Context) *GetGroupParams {
	return &GetGroupParams{
		Context: ctx,
	}
}

// NewGetGroupParamsWithHTTPClient creates a new GetGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGroupParamsWithHTTPClient(client *http.Client) *GetGroupParams {
	return &GetGroupParams{
		HTTPClient: client,
	}
}

/*
GetGroupParams contains all the parameters to send to the API endpoint

	for the get group operation.

	Typically these are written to a http.Request.
*/
type GetGroupParams struct {

	/* Name.

	   Group name
	*/
	Name string

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	/* Userlist.

	   Parent userlist name
	*/
	Userlist string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGroupParams) WithDefaults() *GetGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get group params
func (o *GetGroupParams) WithTimeout(timeout time.Duration) *GetGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get group params
func (o *GetGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get group params
func (o *GetGroupParams) WithContext(ctx context.Context) *GetGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get group params
func (o *GetGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get group params
func (o *GetGroupParams) WithHTTPClient(client *http.Client) *GetGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get group params
func (o *GetGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get group params
func (o *GetGroupParams) WithName(name string) *GetGroupParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get group params
func (o *GetGroupParams) SetName(name string) {
	o.Name = name
}

// WithTransactionID adds the transactionID to the get group params
func (o *GetGroupParams) WithTransactionID(transactionID *string) *GetGroupParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get group params
func (o *GetGroupParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithUserlist adds the userlist to the get group params
func (o *GetGroupParams) WithUserlist(userlist string) *GetGroupParams {
	o.SetUserlist(userlist)
	return o
}

// SetUserlist adds the userlist to the get group params
func (o *GetGroupParams) SetUserlist(userlist string) {
	o.Userlist = userlist
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param userlist
	qrUserlist := o.Userlist
	qUserlist := qrUserlist
	if qUserlist != "" {

		if err := r.SetQueryParam("userlist", qUserlist); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
