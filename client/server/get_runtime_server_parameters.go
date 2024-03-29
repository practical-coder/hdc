// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewGetRuntimeServerParams creates a new GetRuntimeServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRuntimeServerParams() *GetRuntimeServerParams {
	return &GetRuntimeServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRuntimeServerParamsWithTimeout creates a new GetRuntimeServerParams object
// with the ability to set a timeout on a request.
func NewGetRuntimeServerParamsWithTimeout(timeout time.Duration) *GetRuntimeServerParams {
	return &GetRuntimeServerParams{
		timeout: timeout,
	}
}

// NewGetRuntimeServerParamsWithContext creates a new GetRuntimeServerParams object
// with the ability to set a context for a request.
func NewGetRuntimeServerParamsWithContext(ctx context.Context) *GetRuntimeServerParams {
	return &GetRuntimeServerParams{
		Context: ctx,
	}
}

// NewGetRuntimeServerParamsWithHTTPClient creates a new GetRuntimeServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRuntimeServerParamsWithHTTPClient(client *http.Client) *GetRuntimeServerParams {
	return &GetRuntimeServerParams{
		HTTPClient: client,
	}
}

/*
GetRuntimeServerParams contains all the parameters to send to the API endpoint

	for the get runtime server operation.

	Typically these are written to a http.Request.
*/
type GetRuntimeServerParams struct {

	/* Backend.

	   Parent backend name
	*/
	Backend string

	/* Name.

	   Server name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get runtime server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRuntimeServerParams) WithDefaults() *GetRuntimeServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get runtime server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRuntimeServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get runtime server params
func (o *GetRuntimeServerParams) WithTimeout(timeout time.Duration) *GetRuntimeServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runtime server params
func (o *GetRuntimeServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runtime server params
func (o *GetRuntimeServerParams) WithContext(ctx context.Context) *GetRuntimeServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runtime server params
func (o *GetRuntimeServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runtime server params
func (o *GetRuntimeServerParams) WithHTTPClient(client *http.Client) *GetRuntimeServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runtime server params
func (o *GetRuntimeServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the get runtime server params
func (o *GetRuntimeServerParams) WithBackend(backend string) *GetRuntimeServerParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the get runtime server params
func (o *GetRuntimeServerParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithName adds the name to the get runtime server params
func (o *GetRuntimeServerParams) WithName(name string) *GetRuntimeServerParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get runtime server params
func (o *GetRuntimeServerParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetRuntimeServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param backend
	qrBackend := o.Backend
	qBackend := qrBackend
	if qBackend != "" {

		if err := r.SetQueryParam("backend", qBackend); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
