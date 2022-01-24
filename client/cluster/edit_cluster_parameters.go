// Code generated by go-swagger; DO NOT EDIT.

package cluster

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
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v2/models"
)

// NewEditClusterParams creates a new EditClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEditClusterParams() *EditClusterParams {
	return &EditClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEditClusterParamsWithTimeout creates a new EditClusterParams object
// with the ability to set a timeout on a request.
func NewEditClusterParamsWithTimeout(timeout time.Duration) *EditClusterParams {
	return &EditClusterParams{
		timeout: timeout,
	}
}

// NewEditClusterParamsWithContext creates a new EditClusterParams object
// with the ability to set a context for a request.
func NewEditClusterParamsWithContext(ctx context.Context) *EditClusterParams {
	return &EditClusterParams{
		Context: ctx,
	}
}

// NewEditClusterParamsWithHTTPClient creates a new EditClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewEditClusterParamsWithHTTPClient(client *http.Client) *EditClusterParams {
	return &EditClusterParams{
		HTTPClient: client,
	}
}

/* EditClusterParams contains all the parameters to send to the API endpoint
   for the edit cluster operation.

   Typically these are written to a http.Request.
*/
type EditClusterParams struct {

	// Data.
	Data *models.ClusterSettings

	/* Version.

	   Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.
	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edit cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EditClusterParams) WithDefaults() *EditClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edit cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EditClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edit cluster params
func (o *EditClusterParams) WithTimeout(timeout time.Duration) *EditClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit cluster params
func (o *EditClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit cluster params
func (o *EditClusterParams) WithContext(ctx context.Context) *EditClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit cluster params
func (o *EditClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit cluster params
func (o *EditClusterParams) WithHTTPClient(client *http.Client) *EditClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit cluster params
func (o *EditClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the edit cluster params
func (o *EditClusterParams) WithData(data *models.ClusterSettings) *EditClusterParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the edit cluster params
func (o *EditClusterParams) SetData(data *models.ClusterSettings) {
	o.Data = data
}

// WithVersion adds the version to the edit cluster params
func (o *EditClusterParams) WithVersion(version *int64) *EditClusterParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the edit cluster params
func (o *EditClusterParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *EditClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if o.Version != nil {

		// query param version
		var qrVersion int64

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}