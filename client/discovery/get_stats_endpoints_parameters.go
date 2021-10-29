// Code generated by go-swagger; DO NOT EDIT.

package discovery

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

// NewGetStatsEndpointsParams creates a new GetStatsEndpointsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStatsEndpointsParams() *GetStatsEndpointsParams {
	return &GetStatsEndpointsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStatsEndpointsParamsWithTimeout creates a new GetStatsEndpointsParams object
// with the ability to set a timeout on a request.
func NewGetStatsEndpointsParamsWithTimeout(timeout time.Duration) *GetStatsEndpointsParams {
	return &GetStatsEndpointsParams{
		timeout: timeout,
	}
}

// NewGetStatsEndpointsParamsWithContext creates a new GetStatsEndpointsParams object
// with the ability to set a context for a request.
func NewGetStatsEndpointsParamsWithContext(ctx context.Context) *GetStatsEndpointsParams {
	return &GetStatsEndpointsParams{
		Context: ctx,
	}
}

// NewGetStatsEndpointsParamsWithHTTPClient creates a new GetStatsEndpointsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStatsEndpointsParamsWithHTTPClient(client *http.Client) *GetStatsEndpointsParams {
	return &GetStatsEndpointsParams{
		HTTPClient: client,
	}
}

/* GetStatsEndpointsParams contains all the parameters to send to the API endpoint
   for the get stats endpoints operation.

   Typically these are written to a http.Request.
*/
type GetStatsEndpointsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get stats endpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStatsEndpointsParams) WithDefaults() *GetStatsEndpointsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get stats endpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStatsEndpointsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get stats endpoints params
func (o *GetStatsEndpointsParams) WithTimeout(timeout time.Duration) *GetStatsEndpointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stats endpoints params
func (o *GetStatsEndpointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stats endpoints params
func (o *GetStatsEndpointsParams) WithContext(ctx context.Context) *GetStatsEndpointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stats endpoints params
func (o *GetStatsEndpointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stats endpoints params
func (o *GetStatsEndpointsParams) WithHTTPClient(client *http.Client) *GetStatsEndpointsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stats endpoints params
func (o *GetStatsEndpointsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetStatsEndpointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
