// Code generated by go-swagger; DO NOT EDIT.

package spoe

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

// NewGetAllSpoeFilesParams creates a new GetAllSpoeFilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllSpoeFilesParams() *GetAllSpoeFilesParams {
	return &GetAllSpoeFilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllSpoeFilesParamsWithTimeout creates a new GetAllSpoeFilesParams object
// with the ability to set a timeout on a request.
func NewGetAllSpoeFilesParamsWithTimeout(timeout time.Duration) *GetAllSpoeFilesParams {
	return &GetAllSpoeFilesParams{
		timeout: timeout,
	}
}

// NewGetAllSpoeFilesParamsWithContext creates a new GetAllSpoeFilesParams object
// with the ability to set a context for a request.
func NewGetAllSpoeFilesParamsWithContext(ctx context.Context) *GetAllSpoeFilesParams {
	return &GetAllSpoeFilesParams{
		Context: ctx,
	}
}

// NewGetAllSpoeFilesParamsWithHTTPClient creates a new GetAllSpoeFilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllSpoeFilesParamsWithHTTPClient(client *http.Client) *GetAllSpoeFilesParams {
	return &GetAllSpoeFilesParams{
		HTTPClient: client,
	}
}

/*
GetAllSpoeFilesParams contains all the parameters to send to the API endpoint

	for the get all spoe files operation.

	Typically these are written to a http.Request.
*/
type GetAllSpoeFilesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all spoe files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllSpoeFilesParams) WithDefaults() *GetAllSpoeFilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all spoe files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllSpoeFilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all spoe files params
func (o *GetAllSpoeFilesParams) WithTimeout(timeout time.Duration) *GetAllSpoeFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all spoe files params
func (o *GetAllSpoeFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all spoe files params
func (o *GetAllSpoeFilesParams) WithContext(ctx context.Context) *GetAllSpoeFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all spoe files params
func (o *GetAllSpoeFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all spoe files params
func (o *GetAllSpoeFilesParams) WithHTTPClient(client *http.Client) *GetAllSpoeFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all spoe files params
func (o *GetAllSpoeFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllSpoeFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
