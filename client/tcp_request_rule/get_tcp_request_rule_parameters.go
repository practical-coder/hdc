// Code generated by go-swagger; DO NOT EDIT.

package tcp_request_rule

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

// NewGetTCPRequestRuleParams creates a new GetTCPRequestRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTCPRequestRuleParams() *GetTCPRequestRuleParams {
	return &GetTCPRequestRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTCPRequestRuleParamsWithTimeout creates a new GetTCPRequestRuleParams object
// with the ability to set a timeout on a request.
func NewGetTCPRequestRuleParamsWithTimeout(timeout time.Duration) *GetTCPRequestRuleParams {
	return &GetTCPRequestRuleParams{
		timeout: timeout,
	}
}

// NewGetTCPRequestRuleParamsWithContext creates a new GetTCPRequestRuleParams object
// with the ability to set a context for a request.
func NewGetTCPRequestRuleParamsWithContext(ctx context.Context) *GetTCPRequestRuleParams {
	return &GetTCPRequestRuleParams{
		Context: ctx,
	}
}

// NewGetTCPRequestRuleParamsWithHTTPClient creates a new GetTCPRequestRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTCPRequestRuleParamsWithHTTPClient(client *http.Client) *GetTCPRequestRuleParams {
	return &GetTCPRequestRuleParams{
		HTTPClient: client,
	}
}

/*
GetTCPRequestRuleParams contains all the parameters to send to the API endpoint

	for the get TCP request rule operation.

	Typically these are written to a http.Request.
*/
type GetTCPRequestRuleParams struct {

	/* Index.

	   TCP Request Rule Index
	*/
	Index int64

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

// WithDefaults hydrates default values in the get TCP request rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTCPRequestRuleParams) WithDefaults() *GetTCPRequestRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get TCP request rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTCPRequestRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get TCP request rule params
func (o *GetTCPRequestRuleParams) WithTimeout(timeout time.Duration) *GetTCPRequestRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get TCP request rule params
func (o *GetTCPRequestRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get TCP request rule params
func (o *GetTCPRequestRuleParams) WithContext(ctx context.Context) *GetTCPRequestRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get TCP request rule params
func (o *GetTCPRequestRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get TCP request rule params
func (o *GetTCPRequestRuleParams) WithHTTPClient(client *http.Client) *GetTCPRequestRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get TCP request rule params
func (o *GetTCPRequestRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the get TCP request rule params
func (o *GetTCPRequestRuleParams) WithIndex(index int64) *GetTCPRequestRuleParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the get TCP request rule params
func (o *GetTCPRequestRuleParams) SetIndex(index int64) {
	o.Index = index
}

// WithParentName adds the parentName to the get TCP request rule params
func (o *GetTCPRequestRuleParams) WithParentName(parentName string) *GetTCPRequestRuleParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the get TCP request rule params
func (o *GetTCPRequestRuleParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the get TCP request rule params
func (o *GetTCPRequestRuleParams) WithParentType(parentType string) *GetTCPRequestRuleParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the get TCP request rule params
func (o *GetTCPRequestRuleParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the get TCP request rule params
func (o *GetTCPRequestRuleParams) WithTransactionID(transactionID *string) *GetTCPRequestRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get TCP request rule params
func (o *GetTCPRequestRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTCPRequestRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}

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
