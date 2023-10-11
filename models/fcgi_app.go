// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FcgiApp FCGI application
//
// # HAProxy FastCGI application configuration
//
// swagger:model fcgiApp
type FcgiApp struct {

	// Defines the document root on the remote host. The parameter serves to build the default value of FastCGI parameters SCRIPT_FILENAME and PATH_TRANSLATED. It is a mandatory setting.
	// Required: true
	Docroot *string `json:"docroot"`

	// Enables or disables the retrieval of variables related to connection management.
	// Enum: [enabled disabled]
	GetValues string `json:"get_values,omitempty"`

	// Defines the script name to append after a URI that ends with a slash ("/") to set the default value for the FastCGI parameter SCRIPT_NAME. It is an optional setting.
	Index string `json:"index,omitempty"`

	// Tells the FastCGI application whether or not to keep the connection open after it sends a response. If disabled, the FastCGI application closes the connection after responding to this request.
	// Enum: [enabled disabled]
	KeepConn string `json:"keep_conn,omitempty"`

	// log stderrs
	LogStderrs []*FcgiLogStderr `json:"log_stderrs,omitempty"`

	// Defines the maximum number of concurrent requests this application can accept. If the FastCGI application retrieves the variable FCGI_MAX_REQS during connection establishment, it can override this option. Furthermore, if the application does not do multiplexing, it will ignore this option.
	// Minimum: 1
	MaxReqs int64 `json:"max_reqs,omitempty"`

	// Enables or disables the support of connection multiplexing. If the FastCGI application retrieves the variable FCGI_MPXS_CONNS during connection establishment, it can override this option.
	// Enum: [enabled disabled]
	MpxsConns string `json:"mpxs_conns,omitempty"`

	// Declares a FastCGI application
	// Required: true
	// Pattern: ^[^\s]+$
	Name string `json:"name"`

	// pass headers
	PassHeaders []*FcgiPassHeader `json:"pass_headers,omitempty"`

	// Defines a regular expression to extract the script-name and the path-info from the URI.
	// Thus, <regex> must have two captures: the first to capture the script name, and the second to capture the path- info.
	// If not defined, it does not perform matching on the URI, and does not fill the FastCGI parameters PATH_INFO and PATH_TRANSLATED.
	PathInfo string `json:"path_info,omitempty"`

	// set params
	SetParams []*FcgiSetParam `json:"set_params,omitempty"`
}

// Validate validates this fcgi app
func (m *FcgiApp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocroot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGetValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeepConn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogStderrs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxReqs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMpxsConns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSetParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcgiApp) validateDocroot(formats strfmt.Registry) error {

	if err := validate.Required("docroot", "body", m.Docroot); err != nil {
		return err
	}

	return nil
}

var fcgiAppTypeGetValuesPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fcgiAppTypeGetValuesPropEnum = append(fcgiAppTypeGetValuesPropEnum, v)
	}
}

const (

	// FcgiAppGetValuesEnabled captures enum value "enabled"
	FcgiAppGetValuesEnabled string = "enabled"

	// FcgiAppGetValuesDisabled captures enum value "disabled"
	FcgiAppGetValuesDisabled string = "disabled"
)

// prop value enum
func (m *FcgiApp) validateGetValuesEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fcgiAppTypeGetValuesPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FcgiApp) validateGetValues(formats strfmt.Registry) error {
	if swag.IsZero(m.GetValues) { // not required
		return nil
	}

	// value enum
	if err := m.validateGetValuesEnum("get_values", "body", m.GetValues); err != nil {
		return err
	}

	return nil
}

var fcgiAppTypeKeepConnPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fcgiAppTypeKeepConnPropEnum = append(fcgiAppTypeKeepConnPropEnum, v)
	}
}

const (

	// FcgiAppKeepConnEnabled captures enum value "enabled"
	FcgiAppKeepConnEnabled string = "enabled"

	// FcgiAppKeepConnDisabled captures enum value "disabled"
	FcgiAppKeepConnDisabled string = "disabled"
)

// prop value enum
func (m *FcgiApp) validateKeepConnEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fcgiAppTypeKeepConnPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FcgiApp) validateKeepConn(formats strfmt.Registry) error {
	if swag.IsZero(m.KeepConn) { // not required
		return nil
	}

	// value enum
	if err := m.validateKeepConnEnum("keep_conn", "body", m.KeepConn); err != nil {
		return err
	}

	return nil
}

func (m *FcgiApp) validateLogStderrs(formats strfmt.Registry) error {
	if swag.IsZero(m.LogStderrs) { // not required
		return nil
	}

	for i := 0; i < len(m.LogStderrs); i++ {
		if swag.IsZero(m.LogStderrs[i]) { // not required
			continue
		}

		if m.LogStderrs[i] != nil {
			if err := m.LogStderrs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("log_stderrs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("log_stderrs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FcgiApp) validateMaxReqs(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxReqs) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_reqs", "body", m.MaxReqs, 1, false); err != nil {
		return err
	}

	return nil
}

var fcgiAppTypeMpxsConnsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fcgiAppTypeMpxsConnsPropEnum = append(fcgiAppTypeMpxsConnsPropEnum, v)
	}
}

const (

	// FcgiAppMpxsConnsEnabled captures enum value "enabled"
	FcgiAppMpxsConnsEnabled string = "enabled"

	// FcgiAppMpxsConnsDisabled captures enum value "disabled"
	FcgiAppMpxsConnsDisabled string = "disabled"
)

// prop value enum
func (m *FcgiApp) validateMpxsConnsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fcgiAppTypeMpxsConnsPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FcgiApp) validateMpxsConns(formats strfmt.Registry) error {
	if swag.IsZero(m.MpxsConns) { // not required
		return nil
	}

	// value enum
	if err := m.validateMpxsConnsEnum("mpxs_conns", "body", m.MpxsConns); err != nil {
		return err
	}

	return nil
}

func (m *FcgiApp) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *FcgiApp) validatePassHeaders(formats strfmt.Registry) error {
	if swag.IsZero(m.PassHeaders) { // not required
		return nil
	}

	for i := 0; i < len(m.PassHeaders); i++ {
		if swag.IsZero(m.PassHeaders[i]) { // not required
			continue
		}

		if m.PassHeaders[i] != nil {
			if err := m.PassHeaders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pass_headers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pass_headers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FcgiApp) validateSetParams(formats strfmt.Registry) error {
	if swag.IsZero(m.SetParams) { // not required
		return nil
	}

	for i := 0; i < len(m.SetParams); i++ {
		if swag.IsZero(m.SetParams[i]) { // not required
			continue
		}

		if m.SetParams[i] != nil {
			if err := m.SetParams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("set_params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("set_params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this fcgi app based on the context it is used
func (m *FcgiApp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogStderrs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePassHeaders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSetParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcgiApp) contextValidateLogStderrs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LogStderrs); i++ {

		if m.LogStderrs[i] != nil {

			if swag.IsZero(m.LogStderrs[i]) { // not required
				return nil
			}

			if err := m.LogStderrs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("log_stderrs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("log_stderrs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FcgiApp) contextValidatePassHeaders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PassHeaders); i++ {

		if m.PassHeaders[i] != nil {

			if swag.IsZero(m.PassHeaders[i]) { // not required
				return nil
			}

			if err := m.PassHeaders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pass_headers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pass_headers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FcgiApp) contextValidateSetParams(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SetParams); i++ {

		if m.SetParams[i] != nil {

			if swag.IsZero(m.SetParams[i]) { // not required
				return nil
			}

			if err := m.SetParams[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("set_params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("set_params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcgiApp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcgiApp) UnmarshalBinary(b []byte) error {
	var res FcgiApp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
