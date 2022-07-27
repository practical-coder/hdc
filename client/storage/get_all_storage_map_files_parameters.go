// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewGetAllStorageMapFilesParams creates a new GetAllStorageMapFilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllStorageMapFilesParams() *GetAllStorageMapFilesParams {
	return &GetAllStorageMapFilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllStorageMapFilesParamsWithTimeout creates a new GetAllStorageMapFilesParams object
// with the ability to set a timeout on a request.
func NewGetAllStorageMapFilesParamsWithTimeout(timeout time.Duration) *GetAllStorageMapFilesParams {
	return &GetAllStorageMapFilesParams{
		timeout: timeout,
	}
}

// NewGetAllStorageMapFilesParamsWithContext creates a new GetAllStorageMapFilesParams object
// with the ability to set a context for a request.
func NewGetAllStorageMapFilesParamsWithContext(ctx context.Context) *GetAllStorageMapFilesParams {
	return &GetAllStorageMapFilesParams{
		Context: ctx,
	}
}

// NewGetAllStorageMapFilesParamsWithHTTPClient creates a new GetAllStorageMapFilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllStorageMapFilesParamsWithHTTPClient(client *http.Client) *GetAllStorageMapFilesParams {
	return &GetAllStorageMapFilesParams{
		HTTPClient: client,
	}
}

/* GetAllStorageMapFilesParams contains all the parameters to send to the API endpoint
   for the get all storage map files operation.

   Typically these are written to a http.Request.
*/
type GetAllStorageMapFilesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all storage map files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllStorageMapFilesParams) WithDefaults() *GetAllStorageMapFilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all storage map files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllStorageMapFilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all storage map files params
func (o *GetAllStorageMapFilesParams) WithTimeout(timeout time.Duration) *GetAllStorageMapFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all storage map files params
func (o *GetAllStorageMapFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all storage map files params
func (o *GetAllStorageMapFilesParams) WithContext(ctx context.Context) *GetAllStorageMapFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all storage map files params
func (o *GetAllStorageMapFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all storage map files params
func (o *GetAllStorageMapFilesParams) WithHTTPClient(client *http.Client) *GetAllStorageMapFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all storage map files params
func (o *GetAllStorageMapFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllStorageMapFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
