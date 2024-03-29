// Code generated by go-swagger; DO NOT EDIT.

package log_target

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

// NewGetLogTargetsParams creates a new GetLogTargetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLogTargetsParams() *GetLogTargetsParams {
	return &GetLogTargetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogTargetsParamsWithTimeout creates a new GetLogTargetsParams object
// with the ability to set a timeout on a request.
func NewGetLogTargetsParamsWithTimeout(timeout time.Duration) *GetLogTargetsParams {
	return &GetLogTargetsParams{
		timeout: timeout,
	}
}

// NewGetLogTargetsParamsWithContext creates a new GetLogTargetsParams object
// with the ability to set a context for a request.
func NewGetLogTargetsParamsWithContext(ctx context.Context) *GetLogTargetsParams {
	return &GetLogTargetsParams{
		Context: ctx,
	}
}

// NewGetLogTargetsParamsWithHTTPClient creates a new GetLogTargetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLogTargetsParamsWithHTTPClient(client *http.Client) *GetLogTargetsParams {
	return &GetLogTargetsParams{
		HTTPClient: client,
	}
}

/*
GetLogTargetsParams contains all the parameters to send to the API endpoint

	for the get log targets operation.

	Typically these are written to a http.Request.
*/
type GetLogTargetsParams struct {

	/* ParentName.

	   Parent name
	*/
	ParentName *string

	/* ParentType.

	   Parent type
	*/
	ParentType string

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get log targets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogTargetsParams) WithDefaults() *GetLogTargetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get log targets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogTargetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get log targets params
func (o *GetLogTargetsParams) WithTimeout(timeout time.Duration) *GetLogTargetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get log targets params
func (o *GetLogTargetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get log targets params
func (o *GetLogTargetsParams) WithContext(ctx context.Context) *GetLogTargetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get log targets params
func (o *GetLogTargetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get log targets params
func (o *GetLogTargetsParams) WithHTTPClient(client *http.Client) *GetLogTargetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get log targets params
func (o *GetLogTargetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParentName adds the parentName to the get log targets params
func (o *GetLogTargetsParams) WithParentName(parentName *string) *GetLogTargetsParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the get log targets params
func (o *GetLogTargetsParams) SetParentName(parentName *string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the get log targets params
func (o *GetLogTargetsParams) WithParentType(parentType string) *GetLogTargetsParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the get log targets params
func (o *GetLogTargetsParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the get log targets params
func (o *GetLogTargetsParams) WithTransactionID(transactionID *string) *GetLogTargetsParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get log targets params
func (o *GetLogTargetsParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogTargetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ParentName != nil {

		// query param parent_name
		var qrParentName string

		if o.ParentName != nil {
			qrParentName = *o.ParentName
		}
		qParentName := qrParentName
		if qParentName != "" {

			if err := r.SetQueryParam("parent_name", qParentName); err != nil {
				return err
			}
		}
	}

	// query param parent_type
	qrParentType := o.ParentType
	qParentType := qrParentType
	if qParentType != "" {

		if err := r.SetQueryParam("parent_type", qParentType); err != nil {
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
