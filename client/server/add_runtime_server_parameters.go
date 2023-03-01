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

// NewAddRuntimeServerParams creates a new AddRuntimeServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddRuntimeServerParams() *AddRuntimeServerParams {
	return &AddRuntimeServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddRuntimeServerParamsWithTimeout creates a new AddRuntimeServerParams object
// with the ability to set a timeout on a request.
func NewAddRuntimeServerParamsWithTimeout(timeout time.Duration) *AddRuntimeServerParams {
	return &AddRuntimeServerParams{
		timeout: timeout,
	}
}

// NewAddRuntimeServerParamsWithContext creates a new AddRuntimeServerParams object
// with the ability to set a context for a request.
func NewAddRuntimeServerParamsWithContext(ctx context.Context) *AddRuntimeServerParams {
	return &AddRuntimeServerParams{
		Context: ctx,
	}
}

// NewAddRuntimeServerParamsWithHTTPClient creates a new AddRuntimeServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddRuntimeServerParamsWithHTTPClient(client *http.Client) *AddRuntimeServerParams {
	return &AddRuntimeServerParams{
		HTTPClient: client,
	}
}

/*
AddRuntimeServerParams contains all the parameters to send to the API endpoint

	for the add runtime server operation.

	Typically these are written to a http.Request.
*/
type AddRuntimeServerParams struct {

	/* Backend.

	   Parent backend name
	*/
	Backend string

	// Data.
	Data *models.RuntimeAddServer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add runtime server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRuntimeServerParams) WithDefaults() *AddRuntimeServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add runtime server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRuntimeServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add runtime server params
func (o *AddRuntimeServerParams) WithTimeout(timeout time.Duration) *AddRuntimeServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add runtime server params
func (o *AddRuntimeServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add runtime server params
func (o *AddRuntimeServerParams) WithContext(ctx context.Context) *AddRuntimeServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add runtime server params
func (o *AddRuntimeServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add runtime server params
func (o *AddRuntimeServerParams) WithHTTPClient(client *http.Client) *AddRuntimeServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add runtime server params
func (o *AddRuntimeServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the add runtime server params
func (o *AddRuntimeServerParams) WithBackend(backend string) *AddRuntimeServerParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the add runtime server params
func (o *AddRuntimeServerParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithData adds the data to the add runtime server params
func (o *AddRuntimeServerParams) WithData(data *models.RuntimeAddServer) *AddRuntimeServerParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the add runtime server params
func (o *AddRuntimeServerParams) SetData(data *models.RuntimeAddServer) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *AddRuntimeServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}