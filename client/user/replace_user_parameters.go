// Code generated by go-swagger; DO NOT EDIT.

package user

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

	"github.com/haproxytech/client-native/v5/models"
)

// NewReplaceUserParams creates a new ReplaceUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceUserParams() *ReplaceUserParams {
	return &ReplaceUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceUserParamsWithTimeout creates a new ReplaceUserParams object
// with the ability to set a timeout on a request.
func NewReplaceUserParamsWithTimeout(timeout time.Duration) *ReplaceUserParams {
	return &ReplaceUserParams{
		timeout: timeout,
	}
}

// NewReplaceUserParamsWithContext creates a new ReplaceUserParams object
// with the ability to set a context for a request.
func NewReplaceUserParamsWithContext(ctx context.Context) *ReplaceUserParams {
	return &ReplaceUserParams{
		Context: ctx,
	}
}

// NewReplaceUserParamsWithHTTPClient creates a new ReplaceUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceUserParamsWithHTTPClient(client *http.Client) *ReplaceUserParams {
	return &ReplaceUserParams{
		HTTPClient: client,
	}
}

/*
ReplaceUserParams contains all the parameters to send to the API endpoint

	for the replace user operation.

	Typically these are written to a http.Request.
*/
type ReplaceUserParams struct {

	// Data.
	Data *models.User

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	/* Userlist.

	   Parent userlist name
	*/
	Userlist string

	/* Username.

	   User username
	*/
	Username string

	/* Version.

	   Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.
	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the replace user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceUserParams) WithDefaults() *ReplaceUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceUserParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := ReplaceUserParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the replace user params
func (o *ReplaceUserParams) WithTimeout(timeout time.Duration) *ReplaceUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace user params
func (o *ReplaceUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace user params
func (o *ReplaceUserParams) WithContext(ctx context.Context) *ReplaceUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace user params
func (o *ReplaceUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace user params
func (o *ReplaceUserParams) WithHTTPClient(client *http.Client) *ReplaceUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace user params
func (o *ReplaceUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace user params
func (o *ReplaceUserParams) WithData(data *models.User) *ReplaceUserParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace user params
func (o *ReplaceUserParams) SetData(data *models.User) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace user params
func (o *ReplaceUserParams) WithForceReload(forceReload *bool) *ReplaceUserParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace user params
func (o *ReplaceUserParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithTransactionID adds the transactionID to the replace user params
func (o *ReplaceUserParams) WithTransactionID(transactionID *string) *ReplaceUserParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace user params
func (o *ReplaceUserParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithUserlist adds the userlist to the replace user params
func (o *ReplaceUserParams) WithUserlist(userlist string) *ReplaceUserParams {
	o.SetUserlist(userlist)
	return o
}

// SetUserlist adds the userlist to the replace user params
func (o *ReplaceUserParams) SetUserlist(userlist string) {
	o.Userlist = userlist
}

// WithUsername adds the username to the replace user params
func (o *ReplaceUserParams) WithUsername(username string) *ReplaceUserParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the replace user params
func (o *ReplaceUserParams) SetUsername(username string) {
	o.Username = username
}

// WithVersion adds the version to the replace user params
func (o *ReplaceUserParams) WithVersion(version *int64) *ReplaceUserParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace user params
func (o *ReplaceUserParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
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

	// query param userlist
	qrUserlist := o.Userlist
	qUserlist := qrUserlist
	if qUserlist != "" {

		if err := r.SetQueryParam("userlist", qUserlist); err != nil {
			return err
		}
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
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
