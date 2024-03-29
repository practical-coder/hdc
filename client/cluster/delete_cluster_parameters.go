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
)

// NewDeleteClusterParams creates a new DeleteClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteClusterParams() *DeleteClusterParams {
	return &DeleteClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterParamsWithTimeout creates a new DeleteClusterParams object
// with the ability to set a timeout on a request.
func NewDeleteClusterParamsWithTimeout(timeout time.Duration) *DeleteClusterParams {
	return &DeleteClusterParams{
		timeout: timeout,
	}
}

// NewDeleteClusterParamsWithContext creates a new DeleteClusterParams object
// with the ability to set a context for a request.
func NewDeleteClusterParamsWithContext(ctx context.Context) *DeleteClusterParams {
	return &DeleteClusterParams{
		Context: ctx,
	}
}

// NewDeleteClusterParamsWithHTTPClient creates a new DeleteClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteClusterParamsWithHTTPClient(client *http.Client) *DeleteClusterParams {
	return &DeleteClusterParams{
		HTTPClient: client,
	}
}

/*
DeleteClusterParams contains all the parameters to send to the API endpoint

	for the delete cluster operation.

	Typically these are written to a http.Request.
*/
type DeleteClusterParams struct {

	/* Configuration.

	   In case of moving to single mode do we keep or clean configuration
	*/
	Configuration *string

	/* Version.

	   Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.
	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteClusterParams) WithDefaults() *DeleteClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete cluster params
func (o *DeleteClusterParams) WithTimeout(timeout time.Duration) *DeleteClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster params
func (o *DeleteClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster params
func (o *DeleteClusterParams) WithContext(ctx context.Context) *DeleteClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster params
func (o *DeleteClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster params
func (o *DeleteClusterParams) WithHTTPClient(client *http.Client) *DeleteClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster params
func (o *DeleteClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfiguration adds the configuration to the delete cluster params
func (o *DeleteClusterParams) WithConfiguration(configuration *string) *DeleteClusterParams {
	o.SetConfiguration(configuration)
	return o
}

// SetConfiguration adds the configuration to the delete cluster params
func (o *DeleteClusterParams) SetConfiguration(configuration *string) {
	o.Configuration = configuration
}

// WithVersion adds the version to the delete cluster params
func (o *DeleteClusterParams) WithVersion(version *int64) *DeleteClusterParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete cluster params
func (o *DeleteClusterParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Configuration != nil {

		// query param configuration
		var qrConfiguration string

		if o.Configuration != nil {
			qrConfiguration = *o.Configuration
		}
		qConfiguration := qrConfiguration
		if qConfiguration != "" {

			if err := r.SetQueryParam("configuration", qConfiguration); err != nil {
				return err
			}
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
