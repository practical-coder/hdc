// Code generated by go-swagger; DO NOT EDIT.

package stats

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

// NewGetStatsParams creates a new GetStatsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStatsParams() *GetStatsParams {
	return &GetStatsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStatsParamsWithTimeout creates a new GetStatsParams object
// with the ability to set a timeout on a request.
func NewGetStatsParamsWithTimeout(timeout time.Duration) *GetStatsParams {
	return &GetStatsParams{
		timeout: timeout,
	}
}

// NewGetStatsParamsWithContext creates a new GetStatsParams object
// with the ability to set a context for a request.
func NewGetStatsParamsWithContext(ctx context.Context) *GetStatsParams {
	return &GetStatsParams{
		Context: ctx,
	}
}

// NewGetStatsParamsWithHTTPClient creates a new GetStatsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStatsParamsWithHTTPClient(client *http.Client) *GetStatsParams {
	return &GetStatsParams{
		HTTPClient: client,
	}
}

/* GetStatsParams contains all the parameters to send to the API endpoint
   for the get stats operation.

   Typically these are written to a http.Request.
*/
type GetStatsParams struct {

	/* Name.

	   Object name to get stats for
	*/
	Name *string

	/* Parent.

	   Object parent name to get stats for, in case the object is a server
	*/
	Parent *string

	/* Type.

	   Object type to get stats for (one of frontend, backend, server)
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStatsParams) WithDefaults() *GetStatsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStatsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get stats params
func (o *GetStatsParams) WithTimeout(timeout time.Duration) *GetStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stats params
func (o *GetStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stats params
func (o *GetStatsParams) WithContext(ctx context.Context) *GetStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stats params
func (o *GetStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stats params
func (o *GetStatsParams) WithHTTPClient(client *http.Client) *GetStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stats params
func (o *GetStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get stats params
func (o *GetStatsParams) WithName(name *string) *GetStatsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get stats params
func (o *GetStatsParams) SetName(name *string) {
	o.Name = name
}

// WithParent adds the parent to the get stats params
func (o *GetStatsParams) WithParent(parent *string) *GetStatsParams {
	o.SetParent(parent)
	return o
}

// SetParent adds the parent to the get stats params
func (o *GetStatsParams) SetParent(parent *string) {
	o.Parent = parent
}

// WithType adds the typeVar to the get stats params
func (o *GetStatsParams) WithType(typeVar *string) *GetStatsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get stats params
func (o *GetStatsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Parent != nil {

		// query param parent
		var qrParent string

		if o.Parent != nil {
			qrParent = *o.Parent
		}
		qParent := qrParent
		if qParent != "" {

			if err := r.SetQueryParam("parent", qParent); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
