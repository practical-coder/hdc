// Code generated by go-swagger; DO NOT EDIT.

package bind

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

// NewGetBindParams creates a new GetBindParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBindParams() *GetBindParams {
	return &GetBindParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBindParamsWithTimeout creates a new GetBindParams object
// with the ability to set a timeout on a request.
func NewGetBindParamsWithTimeout(timeout time.Duration) *GetBindParams {
	return &GetBindParams{
		timeout: timeout,
	}
}

// NewGetBindParamsWithContext creates a new GetBindParams object
// with the ability to set a context for a request.
func NewGetBindParamsWithContext(ctx context.Context) *GetBindParams {
	return &GetBindParams{
		Context: ctx,
	}
}

// NewGetBindParamsWithHTTPClient creates a new GetBindParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBindParamsWithHTTPClient(client *http.Client) *GetBindParams {
	return &GetBindParams{
		HTTPClient: client,
	}
}

/*
GetBindParams contains all the parameters to send to the API endpoint

	for the get bind operation.

	Typically these are written to a http.Request.
*/
type GetBindParams struct {

	/* Frontend.

	   Parent frontend name
	*/
	Frontend *string

	/* Name.

	   Bind name
	*/
	Name string

	/* ParentName.

	   Parent name
	*/
	ParentName *string

	/* ParentType.

	   Parent type
	*/
	ParentType *string

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get bind params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBindParams) WithDefaults() *GetBindParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bind params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBindParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get bind params
func (o *GetBindParams) WithTimeout(timeout time.Duration) *GetBindParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bind params
func (o *GetBindParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bind params
func (o *GetBindParams) WithContext(ctx context.Context) *GetBindParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bind params
func (o *GetBindParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bind params
func (o *GetBindParams) WithHTTPClient(client *http.Client) *GetBindParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bind params
func (o *GetBindParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrontend adds the frontend to the get bind params
func (o *GetBindParams) WithFrontend(frontend *string) *GetBindParams {
	o.SetFrontend(frontend)
	return o
}

// SetFrontend adds the frontend to the get bind params
func (o *GetBindParams) SetFrontend(frontend *string) {
	o.Frontend = frontend
}

// WithName adds the name to the get bind params
func (o *GetBindParams) WithName(name string) *GetBindParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get bind params
func (o *GetBindParams) SetName(name string) {
	o.Name = name
}

// WithParentName adds the parentName to the get bind params
func (o *GetBindParams) WithParentName(parentName *string) *GetBindParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the get bind params
func (o *GetBindParams) SetParentName(parentName *string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the get bind params
func (o *GetBindParams) WithParentType(parentType *string) *GetBindParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the get bind params
func (o *GetBindParams) SetParentType(parentType *string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the get bind params
func (o *GetBindParams) WithTransactionID(transactionID *string) *GetBindParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get bind params
func (o *GetBindParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Frontend != nil {

		// query param frontend
		var qrFrontend string

		if o.Frontend != nil {
			qrFrontend = *o.Frontend
		}
		qFrontend := qrFrontend
		if qFrontend != "" {

			if err := r.SetQueryParam("frontend", qFrontend); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

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

	if o.ParentType != nil {

		// query param parent_type
		var qrParentType string

		if o.ParentType != nil {
			qrParentType = *o.ParentType
		}
		qParentType := qrParentType
		if qParentType != "" {

			if err := r.SetQueryParam("parent_type", qParentType); err != nil {
				return err
			}
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
