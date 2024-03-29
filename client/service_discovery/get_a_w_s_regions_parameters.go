// Code generated by go-swagger; DO NOT EDIT.

package service_discovery

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

// NewGetAWSRegionsParams creates a new GetAWSRegionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAWSRegionsParams() *GetAWSRegionsParams {
	return &GetAWSRegionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAWSRegionsParamsWithTimeout creates a new GetAWSRegionsParams object
// with the ability to set a timeout on a request.
func NewGetAWSRegionsParamsWithTimeout(timeout time.Duration) *GetAWSRegionsParams {
	return &GetAWSRegionsParams{
		timeout: timeout,
	}
}

// NewGetAWSRegionsParamsWithContext creates a new GetAWSRegionsParams object
// with the ability to set a context for a request.
func NewGetAWSRegionsParamsWithContext(ctx context.Context) *GetAWSRegionsParams {
	return &GetAWSRegionsParams{
		Context: ctx,
	}
}

// NewGetAWSRegionsParamsWithHTTPClient creates a new GetAWSRegionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAWSRegionsParamsWithHTTPClient(client *http.Client) *GetAWSRegionsParams {
	return &GetAWSRegionsParams{
		HTTPClient: client,
	}
}

/*
GetAWSRegionsParams contains all the parameters to send to the API endpoint

	for the get a w s regions operation.

	Typically these are written to a http.Request.
*/
type GetAWSRegionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get a w s regions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAWSRegionsParams) WithDefaults() *GetAWSRegionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get a w s regions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAWSRegionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get a w s regions params
func (o *GetAWSRegionsParams) WithTimeout(timeout time.Duration) *GetAWSRegionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get a w s regions params
func (o *GetAWSRegionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get a w s regions params
func (o *GetAWSRegionsParams) WithContext(ctx context.Context) *GetAWSRegionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get a w s regions params
func (o *GetAWSRegionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get a w s regions params
func (o *GetAWSRegionsParams) WithHTTPClient(client *http.Client) *GetAWSRegionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get a w s regions params
func (o *GetAWSRegionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAWSRegionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
