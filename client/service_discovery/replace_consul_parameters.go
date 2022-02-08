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

	"github.com/haproxytech/client-native/v3/models"
)

// NewReplaceConsulParams creates a new ReplaceConsulParams object
// with the default values initialized.
func NewReplaceConsulParams() *ReplaceConsulParams {
	var ()
	return &ReplaceConsulParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceConsulParamsWithTimeout creates a new ReplaceConsulParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceConsulParamsWithTimeout(timeout time.Duration) *ReplaceConsulParams {
	var ()
	return &ReplaceConsulParams{

		timeout: timeout,
	}
}

// NewReplaceConsulParamsWithContext creates a new ReplaceConsulParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceConsulParamsWithContext(ctx context.Context) *ReplaceConsulParams {
	var ()
	return &ReplaceConsulParams{

		Context: ctx,
	}
}

// NewReplaceConsulParamsWithHTTPClient creates a new ReplaceConsulParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceConsulParamsWithHTTPClient(client *http.Client) *ReplaceConsulParams {
	var ()
	return &ReplaceConsulParams{
		HTTPClient: client,
	}
}

/*ReplaceConsulParams contains all the parameters to send to the API endpoint
for the replace consul operation typically these are written to a http.Request
*/
type ReplaceConsulParams struct {

	/*Data*/
	Data *models.Consul
	/*ID
	  Consul Index

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace consul params
func (o *ReplaceConsulParams) WithTimeout(timeout time.Duration) *ReplaceConsulParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace consul params
func (o *ReplaceConsulParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace consul params
func (o *ReplaceConsulParams) WithContext(ctx context.Context) *ReplaceConsulParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace consul params
func (o *ReplaceConsulParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace consul params
func (o *ReplaceConsulParams) WithHTTPClient(client *http.Client) *ReplaceConsulParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace consul params
func (o *ReplaceConsulParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace consul params
func (o *ReplaceConsulParams) WithData(data *models.Consul) *ReplaceConsulParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace consul params
func (o *ReplaceConsulParams) SetData(data *models.Consul) {
	o.Data = data
}

// WithID adds the id to the replace consul params
func (o *ReplaceConsulParams) WithID(id string) *ReplaceConsulParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace consul params
func (o *ReplaceConsulParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceConsulParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
