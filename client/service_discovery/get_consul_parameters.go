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

// NewGetConsulParams creates a new GetConsulParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConsulParams() *GetConsulParams {
	return &GetConsulParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConsulParamsWithTimeout creates a new GetConsulParams object
// with the ability to set a timeout on a request.
func NewGetConsulParamsWithTimeout(timeout time.Duration) *GetConsulParams {
	return &GetConsulParams{
		timeout: timeout,
	}
}

// NewGetConsulParamsWithContext creates a new GetConsulParams object
// with the ability to set a context for a request.
func NewGetConsulParamsWithContext(ctx context.Context) *GetConsulParams {
	return &GetConsulParams{
		Context: ctx,
	}
}

// NewGetConsulParamsWithHTTPClient creates a new GetConsulParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConsulParamsWithHTTPClient(client *http.Client) *GetConsulParams {
	return &GetConsulParams{
		HTTPClient: client,
	}
}

/*
GetConsulParams contains all the parameters to send to the API endpoint

	for the get consul operation.

	Typically these are written to a http.Request.
*/
type GetConsulParams struct {

	/* ID.

	   Consul server id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get consul params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConsulParams) WithDefaults() *GetConsulParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get consul params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConsulParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get consul params
func (o *GetConsulParams) WithTimeout(timeout time.Duration) *GetConsulParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get consul params
func (o *GetConsulParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get consul params
func (o *GetConsulParams) WithContext(ctx context.Context) *GetConsulParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get consul params
func (o *GetConsulParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get consul params
func (o *GetConsulParams) WithHTTPClient(client *http.Client) *GetConsulParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get consul params
func (o *GetConsulParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get consul params
func (o *GetConsulParams) WithID(id string) *GetConsulParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get consul params
func (o *GetConsulParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetConsulParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
