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

	"github.com/practical-coder/hdc/models"
)

// NewPostServicesHaproxyRuntimeACLFileEntriesParams creates a new PostServicesHaproxyRuntimeACLFileEntriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostServicesHaproxyRuntimeACLFileEntriesParams() *PostServicesHaproxyRuntimeACLFileEntriesParams {
	return &PostServicesHaproxyRuntimeACLFileEntriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostServicesHaproxyRuntimeACLFileEntriesParamsWithTimeout creates a new PostServicesHaproxyRuntimeACLFileEntriesParams object
// with the ability to set a timeout on a request.
func NewPostServicesHaproxyRuntimeACLFileEntriesParamsWithTimeout(timeout time.Duration) *PostServicesHaproxyRuntimeACLFileEntriesParams {
	return &PostServicesHaproxyRuntimeACLFileEntriesParams{
		timeout: timeout,
	}
}

// NewPostServicesHaproxyRuntimeACLFileEntriesParamsWithContext creates a new PostServicesHaproxyRuntimeACLFileEntriesParams object
// with the ability to set a context for a request.
func NewPostServicesHaproxyRuntimeACLFileEntriesParamsWithContext(ctx context.Context) *PostServicesHaproxyRuntimeACLFileEntriesParams {
	return &PostServicesHaproxyRuntimeACLFileEntriesParams{
		Context: ctx,
	}
}

// NewPostServicesHaproxyRuntimeACLFileEntriesParamsWithHTTPClient creates a new PostServicesHaproxyRuntimeACLFileEntriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostServicesHaproxyRuntimeACLFileEntriesParamsWithHTTPClient(client *http.Client) *PostServicesHaproxyRuntimeACLFileEntriesParams {
	return &PostServicesHaproxyRuntimeACLFileEntriesParams{
		HTTPClient: client,
	}
}

/* PostServicesHaproxyRuntimeACLFileEntriesParams contains all the parameters to send to the API endpoint
   for the post services haproxy runtime ACL file entries operation.

   Typically these are written to a http.Request.
*/
type PostServicesHaproxyRuntimeACLFileEntriesParams struct {

	/* ACLID.

	   ACL ID
	*/
	ACLID string

	// Data.
	Data *models.ACLFileEntry

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post services haproxy runtime ACL file entries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostServicesHaproxyRuntimeACLFileEntriesParams) WithDefaults() *PostServicesHaproxyRuntimeACLFileEntriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post services haproxy runtime ACL file entries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostServicesHaproxyRuntimeACLFileEntriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post services haproxy runtime ACL file entries params
func (o *PostServicesHaproxyRuntimeACLFileEntriesParams) WithTimeout(timeout time.Duration) *PostServicesHaproxyRuntimeACLFileEntriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post services haproxy runtime ACL file entries params
func (o *PostServicesHaproxyRuntimeACLFileEntriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post services haproxy runtime ACL file entries params
func (o *PostServicesHaproxyRuntimeACLFileEntriesParams) WithContext(ctx context.Context) *PostServicesHaproxyRuntimeACLFileEntriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post services haproxy runtime ACL file entries params
func (o *PostServicesHaproxyRuntimeACLFileEntriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post services haproxy runtime ACL file entries params
func (o *PostServicesHaproxyRuntimeACLFileEntriesParams) WithHTTPClient(client *http.Client) *PostServicesHaproxyRuntimeACLFileEntriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post services haproxy runtime ACL file entries params
func (o *PostServicesHaproxyRuntimeACLFileEntriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithACLID adds the aCLID to the post services haproxy runtime ACL file entries params
func (o *PostServicesHaproxyRuntimeACLFileEntriesParams) WithACLID(aCLID string) *PostServicesHaproxyRuntimeACLFileEntriesParams {
	o.SetACLID(aCLID)
	return o
}

// SetACLID adds the aclId to the post services haproxy runtime ACL file entries params
func (o *PostServicesHaproxyRuntimeACLFileEntriesParams) SetACLID(aCLID string) {
	o.ACLID = aCLID
}

// WithData adds the data to the post services haproxy runtime ACL file entries params
func (o *PostServicesHaproxyRuntimeACLFileEntriesParams) WithData(data *models.ACLFileEntry) *PostServicesHaproxyRuntimeACLFileEntriesParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the post services haproxy runtime ACL file entries params
func (o *PostServicesHaproxyRuntimeACLFileEntriesParams) SetData(data *models.ACLFileEntry) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PostServicesHaproxyRuntimeACLFileEntriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param acl_id
	qrACLID := o.ACLID
	qACLID := qrACLID
	if qACLID != "" {

		if err := r.SetQueryParam("acl_id", qACLID); err != nil {
			return err
		}
	}
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
