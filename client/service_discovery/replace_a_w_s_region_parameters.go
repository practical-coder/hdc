// Code generated by go-swagger; DO NOT EDIT.

package service_discovery

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

	"github.com/haproxytech/client-native/v5/models"
)

// NewReplaceAWSRegionParams creates a new ReplaceAWSRegionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceAWSRegionParams() *ReplaceAWSRegionParams {
	return &ReplaceAWSRegionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceAWSRegionParamsWithTimeout creates a new ReplaceAWSRegionParams object
// with the ability to set a timeout on a request.
func NewReplaceAWSRegionParamsWithTimeout(timeout time.Duration) *ReplaceAWSRegionParams {
	return &ReplaceAWSRegionParams{
		timeout: timeout,
	}
}

// NewReplaceAWSRegionParamsWithContext creates a new ReplaceAWSRegionParams object
// with the ability to set a context for a request.
func NewReplaceAWSRegionParamsWithContext(ctx context.Context) *ReplaceAWSRegionParams {
	return &ReplaceAWSRegionParams{
		Context: ctx,
	}
}

// NewReplaceAWSRegionParamsWithHTTPClient creates a new ReplaceAWSRegionParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceAWSRegionParamsWithHTTPClient(client *http.Client) *ReplaceAWSRegionParams {
	return &ReplaceAWSRegionParams{
		HTTPClient: client,
	}
}

/*
ReplaceAWSRegionParams contains all the parameters to send to the API endpoint

	for the replace a w s region operation.

	Typically these are written to a http.Request.
*/
type ReplaceAWSRegionParams struct {

	// Data.
	Data *models.AwsRegion

	/* ID.

	   AWS Region ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the replace a w s region params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceAWSRegionParams) WithDefaults() *ReplaceAWSRegionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace a w s region params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceAWSRegionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the replace a w s region params
func (o *ReplaceAWSRegionParams) WithTimeout(timeout time.Duration) *ReplaceAWSRegionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace a w s region params
func (o *ReplaceAWSRegionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace a w s region params
func (o *ReplaceAWSRegionParams) WithContext(ctx context.Context) *ReplaceAWSRegionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace a w s region params
func (o *ReplaceAWSRegionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace a w s region params
func (o *ReplaceAWSRegionParams) WithHTTPClient(client *http.Client) *ReplaceAWSRegionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace a w s region params
func (o *ReplaceAWSRegionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace a w s region params
func (o *ReplaceAWSRegionParams) WithData(data *models.AwsRegion) *ReplaceAWSRegionParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace a w s region params
func (o *ReplaceAWSRegionParams) SetData(data *models.AwsRegion) {
	o.Data = data
}

// WithID adds the id to the replace a w s region params
func (o *ReplaceAWSRegionParams) WithID(id string) *ReplaceAWSRegionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace a w s region params
func (o *ReplaceAWSRegionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceAWSRegionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
