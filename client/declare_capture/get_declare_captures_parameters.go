// Code generated by go-swagger; DO NOT EDIT.

package declare_capture

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

// NewGetDeclareCapturesParams creates a new GetDeclareCapturesParams object
// with the default values initialized.
func NewGetDeclareCapturesParams() *GetDeclareCapturesParams {
	var ()
	return &GetDeclareCapturesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeclareCapturesParamsWithTimeout creates a new GetDeclareCapturesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeclareCapturesParamsWithTimeout(timeout time.Duration) *GetDeclareCapturesParams {
	var ()
	return &GetDeclareCapturesParams{

		timeout: timeout,
	}
}

// NewGetDeclareCapturesParamsWithContext creates a new GetDeclareCapturesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeclareCapturesParamsWithContext(ctx context.Context) *GetDeclareCapturesParams {
	var ()
	return &GetDeclareCapturesParams{

		Context: ctx,
	}
}

// NewGetDeclareCapturesParamsWithHTTPClient creates a new GetDeclareCapturesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeclareCapturesParamsWithHTTPClient(client *http.Client) *GetDeclareCapturesParams {
	var ()
	return &GetDeclareCapturesParams{
		HTTPClient: client,
	}
}

/*GetDeclareCapturesParams contains all the parameters to send to the API endpoint
for the get declare captures operation typically these are written to a http.Request
*/
type GetDeclareCapturesParams struct {

	/*Frontend
	  Parent frontend name

	*/
	Frontend string
	/*TransactionID
	  ID of the transaction where we want to add the operation. Cannot be used when version is specified.

	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get declare captures params
func (o *GetDeclareCapturesParams) WithTimeout(timeout time.Duration) *GetDeclareCapturesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get declare captures params
func (o *GetDeclareCapturesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get declare captures params
func (o *GetDeclareCapturesParams) WithContext(ctx context.Context) *GetDeclareCapturesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get declare captures params
func (o *GetDeclareCapturesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get declare captures params
func (o *GetDeclareCapturesParams) WithHTTPClient(client *http.Client) *GetDeclareCapturesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get declare captures params
func (o *GetDeclareCapturesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrontend adds the frontend to the get declare captures params
func (o *GetDeclareCapturesParams) WithFrontend(frontend string) *GetDeclareCapturesParams {
	o.SetFrontend(frontend)
	return o
}

// SetFrontend adds the frontend to the get declare captures params
func (o *GetDeclareCapturesParams) SetFrontend(frontend string) {
	o.Frontend = frontend
}

// WithTransactionID adds the transactionID to the get declare captures params
func (o *GetDeclareCapturesParams) WithTransactionID(transactionID *string) *GetDeclareCapturesParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get declare captures params
func (o *GetDeclareCapturesParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeclareCapturesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param frontend
	qrFrontend := o.Frontend
	qFrontend := qrFrontend
	if qFrontend != "" {
		if err := r.SetQueryParam("frontend", qFrontend); err != nil {
			return err
		}
	}

	if o.TransactionID != nil {

		// query param transaction_id
		var qrTransactionID string
		if o.TransactionID != nil {
			qrTransactionID = *o.TransactionID
		}
		qTransactionID := qrTransactionID
		if qTransactionID != "" {
			if err := r.SetQueryParam("transaction_id", qTransactionID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
