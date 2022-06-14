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

// NewGetGroupsParams creates a new GetGroupsParams object
// with the default values initialized.
func NewGetGroupsParams() *GetGroupsParams {
	var ()
	return &GetGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupsParamsWithTimeout creates a new GetGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGroupsParamsWithTimeout(timeout time.Duration) *GetGroupsParams {
	var ()
	return &GetGroupsParams{

		timeout: timeout,
	}
}

// NewGetGroupsParamsWithContext creates a new GetGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGroupsParamsWithContext(ctx context.Context) *GetGroupsParams {
	var ()
	return &GetGroupsParams{

		Context: ctx,
	}
}

// NewGetGroupsParamsWithHTTPClient creates a new GetGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGroupsParamsWithHTTPClient(client *http.Client) *GetGroupsParams {
	var ()
	return &GetGroupsParams{
		HTTPClient: client,
	}
}

/*GetGroupsParams contains all the parameters to send to the API endpoint
for the get groups operation typically these are written to a http.Request
*/
type GetGroupsParams struct {

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

// WithTimeout adds the timeout to the get groups params
func (o *GetGroupsParams) WithTimeout(timeout time.Duration) *GetGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get groups params
func (o *GetGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get groups params
func (o *GetGroupsParams) WithContext(ctx context.Context) *GetGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get groups params
func (o *GetGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get groups params
func (o *GetGroupsParams) WithHTTPClient(client *http.Client) *GetGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get groups params
func (o *GetGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTransactionID adds the transactionID to the get groups params
func (o *GetGroupsParams) WithTransactionID(transactionID *string) *GetGroupsParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get groups params
func (o *GetGroupsParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithUserlist adds the userlist to the get groups params
func (o *GetGroupsParams) WithUserlist(userlist string) *GetGroupsParams {
	o.SetUserlist(userlist)
	return o
}

// SetUserlist adds the userlist to the get groups params
func (o *GetGroupsParams) SetUserlist(userlist string) {
	o.Userlist = userlist
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
