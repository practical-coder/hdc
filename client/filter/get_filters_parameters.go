// Code generated by go-swagger; DO NOT EDIT.

package filter

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

// NewGetFiltersParams creates a new GetFiltersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFiltersParams() *GetFiltersParams {
	return &GetFiltersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFiltersParamsWithTimeout creates a new GetFiltersParams object
// with the ability to set a timeout on a request.
func NewGetFiltersParamsWithTimeout(timeout time.Duration) *GetFiltersParams {
	return &GetFiltersParams{
		timeout: timeout,
	}
}

// NewGetFiltersParamsWithContext creates a new GetFiltersParams object
// with the ability to set a context for a request.
func NewGetFiltersParamsWithContext(ctx context.Context) *GetFiltersParams {
	return &GetFiltersParams{
		Context: ctx,
	}
}

// NewGetFiltersParamsWithHTTPClient creates a new GetFiltersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFiltersParamsWithHTTPClient(client *http.Client) *GetFiltersParams {
	return &GetFiltersParams{
		HTTPClient: client,
	}
}

/*
GetFiltersParams contains all the parameters to send to the API endpoint

	for the get filters operation.

	Typically these are written to a http.Request.
*/
type GetFiltersParams struct {

	/* ParentName.

	   Parent name
	*/
	ParentName string

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

// WithDefaults hydrates default values in the get filters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFiltersParams) WithDefaults() *GetFiltersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get filters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFiltersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get filters params
func (o *GetFiltersParams) WithTimeout(timeout time.Duration) *GetFiltersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get filters params
func (o *GetFiltersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get filters params
func (o *GetFiltersParams) WithContext(ctx context.Context) *GetFiltersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get filters params
func (o *GetFiltersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get filters params
func (o *GetFiltersParams) WithHTTPClient(client *http.Client) *GetFiltersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get filters params
func (o *GetFiltersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParentName adds the parentName to the get filters params
func (o *GetFiltersParams) WithParentName(parentName string) *GetFiltersParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the get filters params
func (o *GetFiltersParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the get filters params
func (o *GetFiltersParams) WithParentType(parentType string) *GetFiltersParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the get filters params
func (o *GetFiltersParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the get filters params
func (o *GetFiltersParams) WithTransactionID(transactionID *string) *GetFiltersParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get filters params
func (o *GetFiltersParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFiltersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param parent_name
	qrParentName := o.ParentName
	qParentName := qrParentName
	if qParentName != "" {

		if err := r.SetQueryParam("parent_name", qParentName); err != nil {
			return err
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
