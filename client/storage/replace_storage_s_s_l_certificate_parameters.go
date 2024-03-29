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
	"github.com/go-openapi/swag"
)

// NewReplaceStorageSSLCertificateParams creates a new ReplaceStorageSSLCertificateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceStorageSSLCertificateParams() *ReplaceStorageSSLCertificateParams {
	return &ReplaceStorageSSLCertificateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceStorageSSLCertificateParamsWithTimeout creates a new ReplaceStorageSSLCertificateParams object
// with the ability to set a timeout on a request.
func NewReplaceStorageSSLCertificateParamsWithTimeout(timeout time.Duration) *ReplaceStorageSSLCertificateParams {
	return &ReplaceStorageSSLCertificateParams{
		timeout: timeout,
	}
}

// NewReplaceStorageSSLCertificateParamsWithContext creates a new ReplaceStorageSSLCertificateParams object
// with the ability to set a context for a request.
func NewReplaceStorageSSLCertificateParamsWithContext(ctx context.Context) *ReplaceStorageSSLCertificateParams {
	return &ReplaceStorageSSLCertificateParams{
		Context: ctx,
	}
}

// NewReplaceStorageSSLCertificateParamsWithHTTPClient creates a new ReplaceStorageSSLCertificateParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceStorageSSLCertificateParamsWithHTTPClient(client *http.Client) *ReplaceStorageSSLCertificateParams {
	return &ReplaceStorageSSLCertificateParams{
		HTTPClient: client,
	}
}

/*
ReplaceStorageSSLCertificateParams contains all the parameters to send to the API endpoint

	for the replace storage s s l certificate operation.

	Typically these are written to a http.Request.
*/
type ReplaceStorageSSLCertificateParams struct {

	// Data.
	Data string

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Name.

	   SSL certificate name
	*/
	Name string

	/* SkipReload.

	   If set, no reload will be initiated after update
	*/
	SkipReload *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the replace storage s s l certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceStorageSSLCertificateParams) WithDefaults() *ReplaceStorageSSLCertificateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace storage s s l certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceStorageSSLCertificateParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)

		skipReloadDefault = bool(false)
	)

	val := ReplaceStorageSSLCertificateParams{
		ForceReload: &forceReloadDefault,
		SkipReload:  &skipReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the replace storage s s l certificate params
func (o *ReplaceStorageSSLCertificateParams) WithTimeout(timeout time.Duration) *ReplaceStorageSSLCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace storage s s l certificate params
func (o *ReplaceStorageSSLCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace storage s s l certificate params
func (o *ReplaceStorageSSLCertificateParams) WithContext(ctx context.Context) *ReplaceStorageSSLCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace storage s s l certificate params
func (o *ReplaceStorageSSLCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace storage s s l certificate params
func (o *ReplaceStorageSSLCertificateParams) WithHTTPClient(client *http.Client) *ReplaceStorageSSLCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace storage s s l certificate params
func (o *ReplaceStorageSSLCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace storage s s l certificate params
func (o *ReplaceStorageSSLCertificateParams) WithData(data string) *ReplaceStorageSSLCertificateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace storage s s l certificate params
func (o *ReplaceStorageSSLCertificateParams) SetData(data string) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace storage s s l certificate params
func (o *ReplaceStorageSSLCertificateParams) WithForceReload(forceReload *bool) *ReplaceStorageSSLCertificateParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace storage s s l certificate params
func (o *ReplaceStorageSSLCertificateParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithName adds the name to the replace storage s s l certificate params
func (o *ReplaceStorageSSLCertificateParams) WithName(name string) *ReplaceStorageSSLCertificateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace storage s s l certificate params
func (o *ReplaceStorageSSLCertificateParams) SetName(name string) {
	o.Name = name
}

// WithSkipReload adds the skipReload to the replace storage s s l certificate params
func (o *ReplaceStorageSSLCertificateParams) WithSkipReload(skipReload *bool) *ReplaceStorageSSLCertificateParams {
	o.SetSkipReload(skipReload)
	return o
}

// SetSkipReload adds the skipReload to the replace storage s s l certificate params
func (o *ReplaceStorageSSLCertificateParams) SetSkipReload(skipReload *bool) {
	o.SkipReload = skipReload
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceStorageSSLCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	if o.ForceReload != nil {

		// query param force_reload
		var qrForceReload bool

		if o.ForceReload != nil {
			qrForceReload = *o.ForceReload
		}
		qForceReload := swag.FormatBool(qrForceReload)
		if qForceReload != "" {

			if err := r.SetQueryParam("force_reload", qForceReload); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.SkipReload != nil {

		// query param skip_reload
		var qrSkipReload bool

		if o.SkipReload != nil {
			qrSkipReload = *o.SkipReload
		}
		qSkipReload := swag.FormatBool(qrSkipReload)
		if qSkipReload != "" {

			if err := r.SetQueryParam("skip_reload", qSkipReload); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
