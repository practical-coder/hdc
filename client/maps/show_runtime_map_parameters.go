// Code generated by go-swagger; DO NOT EDIT.

package maps

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

// NewShowRuntimeMapParams creates a new ShowRuntimeMapParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShowRuntimeMapParams() *ShowRuntimeMapParams {
	return &ShowRuntimeMapParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShowRuntimeMapParamsWithTimeout creates a new ShowRuntimeMapParams object
// with the ability to set a timeout on a request.
func NewShowRuntimeMapParamsWithTimeout(timeout time.Duration) *ShowRuntimeMapParams {
	return &ShowRuntimeMapParams{
		timeout: timeout,
	}
}

// NewShowRuntimeMapParamsWithContext creates a new ShowRuntimeMapParams object
// with the ability to set a context for a request.
func NewShowRuntimeMapParamsWithContext(ctx context.Context) *ShowRuntimeMapParams {
	return &ShowRuntimeMapParams{
		Context: ctx,
	}
}

// NewShowRuntimeMapParamsWithHTTPClient creates a new ShowRuntimeMapParams object
// with the ability to set a custom HTTPClient for a request.
func NewShowRuntimeMapParamsWithHTTPClient(client *http.Client) *ShowRuntimeMapParams {
	return &ShowRuntimeMapParams{
		HTTPClient: client,
	}
}

/* ShowRuntimeMapParams contains all the parameters to send to the API endpoint
   for the show runtime map operation.

   Typically these are written to a http.Request.
*/
type ShowRuntimeMapParams struct {

	/* Map.

	   Mapfile attribute storage_name
	*/
	Map string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the show runtime map params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowRuntimeMapParams) WithDefaults() *ShowRuntimeMapParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the show runtime map params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowRuntimeMapParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the show runtime map params
func (o *ShowRuntimeMapParams) WithTimeout(timeout time.Duration) *ShowRuntimeMapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the show runtime map params
func (o *ShowRuntimeMapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the show runtime map params
func (o *ShowRuntimeMapParams) WithContext(ctx context.Context) *ShowRuntimeMapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the show runtime map params
func (o *ShowRuntimeMapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the show runtime map params
func (o *ShowRuntimeMapParams) WithHTTPClient(client *http.Client) *ShowRuntimeMapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the show runtime map params
func (o *ShowRuntimeMapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMap adds the mapVar to the show runtime map params
func (o *ShowRuntimeMapParams) WithMap(mapVar string) *ShowRuntimeMapParams {
	o.SetMap(mapVar)
	return o
}

// SetMap adds the map to the show runtime map params
func (o *ShowRuntimeMapParams) SetMap(mapVar string) {
	o.Map = mapVar
}

// WriteToRequest writes these params to a swagger request
func (o *ShowRuntimeMapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param map
	qrMap := o.Map
	qMap := qrMap
	if qMap != "" {

		if err := r.SetQueryParam("map", qMap); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
