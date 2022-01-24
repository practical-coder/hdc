// Code generated by go-swagger; DO NOT EDIT.

package acl_runtime

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

// NewGetServicesHaproxyRuntimeAclsIDParams creates a new GetServicesHaproxyRuntimeAclsIDParams object
// with the default values initialized.
func NewGetServicesHaproxyRuntimeAclsIDParams() *GetServicesHaproxyRuntimeAclsIDParams {
	var ()
	return &GetServicesHaproxyRuntimeAclsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServicesHaproxyRuntimeAclsIDParamsWithTimeout creates a new GetServicesHaproxyRuntimeAclsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServicesHaproxyRuntimeAclsIDParamsWithTimeout(timeout time.Duration) *GetServicesHaproxyRuntimeAclsIDParams {
	var ()
	return &GetServicesHaproxyRuntimeAclsIDParams{

		timeout: timeout,
	}
}

// NewGetServicesHaproxyRuntimeAclsIDParamsWithContext creates a new GetServicesHaproxyRuntimeAclsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServicesHaproxyRuntimeAclsIDParamsWithContext(ctx context.Context) *GetServicesHaproxyRuntimeAclsIDParams {
	var ()
	return &GetServicesHaproxyRuntimeAclsIDParams{

		Context: ctx,
	}
}

// NewGetServicesHaproxyRuntimeAclsIDParamsWithHTTPClient creates a new GetServicesHaproxyRuntimeAclsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServicesHaproxyRuntimeAclsIDParamsWithHTTPClient(client *http.Client) *GetServicesHaproxyRuntimeAclsIDParams {
	var ()
	return &GetServicesHaproxyRuntimeAclsIDParams{
		HTTPClient: client,
	}
}

/*GetServicesHaproxyRuntimeAclsIDParams contains all the parameters to send to the API endpoint
for the get services haproxy runtime acls ID operation typically these are written to a http.Request
*/
type GetServicesHaproxyRuntimeAclsIDParams struct {

	/*ID
	  ACL file entry ID

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get services haproxy runtime acls ID params
func (o *GetServicesHaproxyRuntimeAclsIDParams) WithTimeout(timeout time.Duration) *GetServicesHaproxyRuntimeAclsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get services haproxy runtime acls ID params
func (o *GetServicesHaproxyRuntimeAclsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get services haproxy runtime acls ID params
func (o *GetServicesHaproxyRuntimeAclsIDParams) WithContext(ctx context.Context) *GetServicesHaproxyRuntimeAclsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get services haproxy runtime acls ID params
func (o *GetServicesHaproxyRuntimeAclsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get services haproxy runtime acls ID params
func (o *GetServicesHaproxyRuntimeAclsIDParams) WithHTTPClient(client *http.Client) *GetServicesHaproxyRuntimeAclsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get services haproxy runtime acls ID params
func (o *GetServicesHaproxyRuntimeAclsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get services haproxy runtime acls ID params
func (o *GetServicesHaproxyRuntimeAclsIDParams) WithID(id string) *GetServicesHaproxyRuntimeAclsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get services haproxy runtime acls ID params
func (o *GetServicesHaproxyRuntimeAclsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetServicesHaproxyRuntimeAclsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
