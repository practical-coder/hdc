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

	"github.com/haproxytech/client-native/v4/models"
)

// NewReplaceRuntimeServerParams creates a new ReplaceRuntimeServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceRuntimeServerParams() *ReplaceRuntimeServerParams {
	return &ReplaceRuntimeServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceRuntimeServerParamsWithTimeout creates a new ReplaceRuntimeServerParams object
// with the ability to set a timeout on a request.
func NewReplaceRuntimeServerParamsWithTimeout(timeout time.Duration) *ReplaceRuntimeServerParams {
	return &ReplaceRuntimeServerParams{
		timeout: timeout,
	}
}

// NewReplaceRuntimeServerParamsWithContext creates a new ReplaceRuntimeServerParams object
// with the ability to set a context for a request.
func NewReplaceRuntimeServerParamsWithContext(ctx context.Context) *ReplaceRuntimeServerParams {
	return &ReplaceRuntimeServerParams{
		Context: ctx,
	}
}

// NewReplaceRuntimeServerParamsWithHTTPClient creates a new ReplaceRuntimeServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceRuntimeServerParamsWithHTTPClient(client *http.Client) *ReplaceRuntimeServerParams {
	return &ReplaceRuntimeServerParams{
		HTTPClient: client,
	}
}

/*
ReplaceRuntimeServerParams contains all the parameters to send to the API endpoint

	for the replace runtime server operation.

	Typically these are written to a http.Request.
*/
type ReplaceRuntimeServerParams struct {

	/* Backend.

	   Parent backend name
	*/
	Backend string

	// Data.
	Data *models.RuntimeServer

	/* Name.

	   Server name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the replace runtime server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceRuntimeServerParams) WithDefaults() *ReplaceRuntimeServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace runtime server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceRuntimeServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the replace runtime server params
func (o *ReplaceRuntimeServerParams) WithTimeout(timeout time.Duration) *ReplaceRuntimeServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace runtime server params
func (o *ReplaceRuntimeServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace runtime server params
func (o *ReplaceRuntimeServerParams) WithContext(ctx context.Context) *ReplaceRuntimeServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace runtime server params
func (o *ReplaceRuntimeServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace runtime server params
func (o *ReplaceRuntimeServerParams) WithHTTPClient(client *http.Client) *ReplaceRuntimeServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace runtime server params
func (o *ReplaceRuntimeServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the replace runtime server params
func (o *ReplaceRuntimeServerParams) WithBackend(backend string) *ReplaceRuntimeServerParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the replace runtime server params
func (o *ReplaceRuntimeServerParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithData adds the data to the replace runtime server params
func (o *ReplaceRuntimeServerParams) WithData(data *models.RuntimeServer) *ReplaceRuntimeServerParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace runtime server params
func (o *ReplaceRuntimeServerParams) SetData(data *models.RuntimeServer) {
	o.Data = data
}

// WithName adds the name to the replace runtime server params
func (o *ReplaceRuntimeServerParams) WithName(name string) *ReplaceRuntimeServerParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace runtime server params
func (o *ReplaceRuntimeServerParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceRuntimeServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
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
