// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewGetUsersParams creates a new GetUsersParams object
// with the default values initialized.
func NewGetUsersParams() *GetUsersParams {
	var ()
	return &GetUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersParamsWithTimeout creates a new GetUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersParamsWithTimeout(timeout time.Duration) *GetUsersParams {
	var ()
	return &GetUsersParams{

		timeout: timeout,
	}
}

// NewGetUsersParamsWithContext creates a new GetUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersParamsWithContext(ctx context.Context) *GetUsersParams {
	var ()
	return &GetUsersParams{

		Context: ctx,
	}
}

// NewGetUsersParamsWithHTTPClient creates a new GetUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersParamsWithHTTPClient(client *http.Client) *GetUsersParams {
	var ()
	return &GetUsersParams{
		HTTPClient: client,
	}
}

/*GetUsersParams contains all the parameters to send to the API endpoint
for the get users operation typically these are written to a http.Request
*/
type GetUsersParams struct {

	/*TransactionID
	  ID of the transaction where we want to add the operation. Cannot be used when version is specified.

	*/
	TransactionID *string
	/*Userlist
	  Parent userlist name

	*/
	Userlist string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users params
func (o *GetUsersParams) WithTimeout(timeout time.Duration) *GetUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users params
func (o *GetUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users params
func (o *GetUsersParams) WithContext(ctx context.Context) *GetUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users params
func (o *GetUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users params
func (o *GetUsersParams) WithHTTPClient(client *http.Client) *GetUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users params
func (o *GetUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTransactionID adds the transactionID to the get users params
func (o *GetUsersParams) WithTransactionID(transactionID *string) *GetUsersParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get users params
func (o *GetUsersParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithUserlist adds the userlist to the get users params
func (o *GetUsersParams) WithUserlist(userlist string) *GetUsersParams {
	o.SetUserlist(userlist)
	return o
}

// SetUserlist adds the userlist to the get users params
func (o *GetUsersParams) SetUserlist(userlist string) {
	o.Userlist = userlist
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
