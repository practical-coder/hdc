// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Frontend Frontend
//
// HAProxy frontend configuration
// Example: {"default_backend":"test_backend","http_connection_mode":"http-keep-alive","maxconn":2000,"mode":"http","name":"test_frontend"}
//
// swagger:model frontend
type Frontend struct {

	// accept invalid http request
	// Enum: [enabled disabled]
	AcceptInvalidHTTPRequest string `json:"accept_invalid_http_request,omitempty"`

	// bind process
	// Pattern: ^[^\s]+$
	BindProcess string `json:"bind_process,omitempty"`

	// clflog
	Clflog bool `json:"clflog,omitempty"`

	// client timeout
	ClientTimeout *int64 `json:"client_timeout,omitempty"`

	// clitcpka
	// Enum: [enabled disabled]
	Clitcpka string `json:"clitcpka,omitempty"`

	// contstats
	// Enum: [enabled]
	Contstats string `json:"contstats,omitempty"`

	// default backend
	// Pattern: ^[A-Za-z0-9-_.:]+$
	DefaultBackend string `json:"default_backend,omitempty"`

	// dontlognull
	// Enum: [enabled disabled]
	Dontlognull string `json:"dontlognull,omitempty"`

	// forwardfor
	Forwardfor *Forwardfor `json:"forwardfor,omitempty"`

	// h1 case adjust bogus client
	// Enum: [enabled disabled]
	H1CaseAdjustBogusClient string `json:"h1_case_adjust_bogus_client,omitempty"`

	// http buffer request
	// Enum: [enabled disabled]
	HTTPBufferRequest string `json:"http-buffer-request,omitempty"`

	// http use htx
	// Enum: [enabled disabled]
	HTTPUseHtx string `json:"http-use-htx,omitempty"`

	// http connection mode
	// Enum: [httpclose http-server-close http-keep-alive]
	HTTPConnectionMode string `json:"http_connection_mode,omitempty"`

	// http keep alive timeout
	HTTPKeepAliveTimeout *int64 `json:"http_keep_alive_timeout,omitempty"`

	// http request timeout
	HTTPRequestTimeout *int64 `json:"http_request_timeout,omitempty"`

	// httplog
	Httplog bool `json:"httplog,omitempty"`

	// log format
	LogFormat string `json:"log_format,omitempty"`

	// log format sd
	LogFormatSd string `json:"log_format_sd,omitempty"`

	// log separate errors
	// Enum: [enabled disabled]
	LogSeparateErrors string `json:"log_separate_errors,omitempty"`

	// log tag
	// Pattern: ^[A-Za-z0-9-_.:]+$
	LogTag string `json:"log_tag,omitempty"`

	// logasap
	// Enum: [enabled disabled]
	Logasap string `json:"logasap,omitempty"`

	// maxconn
	Maxconn *int64 `json:"maxconn,omitempty"`

	// mode
	// Enum: [http tcp]
	Mode string `json:"mode,omitempty"`

	// monitor fail
	MonitorFail *MonitorFail `json:"monitor_fail,omitempty"`

	// monitor uri
	MonitorURI MonitorURI `json:"monitor_uri,omitempty"`

	// name
	// Required: true
	// Pattern: ^[A-Za-z0-9-_.:]+$
	Name string `json:"name"`

	// stats options
	StatsOptions *StatsOptions `json:"stats_options,omitempty"`

	// tcplog
	Tcplog bool `json:"tcplog,omitempty"`

	// unique id format
	UniqueIDFormat string `json:"unique_id_format,omitempty"`

	// unique id header
	UniqueIDHeader string `json:"unique_id_header,omitempty"`
}

// Validate validates this frontend
func (m *Frontend) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptInvalidHTTPRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBindProcess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClitcpka(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContstats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultBackend(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDontlognull(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForwardfor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateH1CaseAdjustBogusClient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPBufferRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPUseHtx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPConnectionMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogSeparateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogasap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitorFail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitorURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatsOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var frontendTypeAcceptInvalidHTTPRequestPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeAcceptInvalidHTTPRequestPropEnum = append(frontendTypeAcceptInvalidHTTPRequestPropEnum, v)
	}
}

const (

	// FrontendAcceptInvalidHTTPRequestEnabled captures enum value "enabled"
	FrontendAcceptInvalidHTTPRequestEnabled string = "enabled"

	// FrontendAcceptInvalidHTTPRequestDisabled captures enum value "disabled"
	FrontendAcceptInvalidHTTPRequestDisabled string = "disabled"
)

// prop value enum
func (m *Frontend) validateAcceptInvalidHTTPRequestEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, frontendTypeAcceptInvalidHTTPRequestPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateAcceptInvalidHTTPRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.AcceptInvalidHTTPRequest) { // not required
		return nil
	}

	// value enum
	if err := m.validateAcceptInvalidHTTPRequestEnum("accept_invalid_http_request", "body", m.AcceptInvalidHTTPRequest); err != nil {
		return err
	}

	return nil
}

func (m *Frontend) validateBindProcess(formats strfmt.Registry) error {
	if swag.IsZero(m.BindProcess) { // not required
		return nil
	}

	if err := validate.Pattern("bind_process", "body", m.BindProcess, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var frontendTypeClitcpkaPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeClitcpkaPropEnum = append(frontendTypeClitcpkaPropEnum, v)
	}
}

const (

	// FrontendClitcpkaEnabled captures enum value "enabled"
	FrontendClitcpkaEnabled string = "enabled"

	// FrontendClitcpkaDisabled captures enum value "disabled"
	FrontendClitcpkaDisabled string = "disabled"
)

// prop value enum
func (m *Frontend) validateClitcpkaEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, frontendTypeClitcpkaPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateClitcpka(formats strfmt.Registry) error {
	if swag.IsZero(m.Clitcpka) { // not required
		return nil
	}

	// value enum
	if err := m.validateClitcpkaEnum("clitcpka", "body", m.Clitcpka); err != nil {
		return err
	}

	return nil
}

var frontendTypeContstatsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeContstatsPropEnum = append(frontendTypeContstatsPropEnum, v)
	}
}

const (

	// FrontendContstatsEnabled captures enum value "enabled"
	FrontendContstatsEnabled string = "enabled"
)

// prop value enum
func (m *Frontend) validateContstatsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, frontendTypeContstatsPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateContstats(formats strfmt.Registry) error {
	if swag.IsZero(m.Contstats) { // not required
		return nil
	}

	// value enum
	if err := m.validateContstatsEnum("contstats", "body", m.Contstats); err != nil {
		return err
	}

	return nil
}

func (m *Frontend) validateDefaultBackend(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultBackend) { // not required
		return nil
	}

	if err := validate.Pattern("default_backend", "body", m.DefaultBackend, `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

var frontendTypeDontlognullPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeDontlognullPropEnum = append(frontendTypeDontlognullPropEnum, v)
	}
}

const (

	// FrontendDontlognullEnabled captures enum value "enabled"
	FrontendDontlognullEnabled string = "enabled"

	// FrontendDontlognullDisabled captures enum value "disabled"
	FrontendDontlognullDisabled string = "disabled"
)

// prop value enum
func (m *Frontend) validateDontlognullEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, frontendTypeDontlognullPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateDontlognull(formats strfmt.Registry) error {
	if swag.IsZero(m.Dontlognull) { // not required
		return nil
	}

	// value enum
	if err := m.validateDontlognullEnum("dontlognull", "body", m.Dontlognull); err != nil {
		return err
	}

	return nil
}

func (m *Frontend) validateForwardfor(formats strfmt.Registry) error {
	if swag.IsZero(m.Forwardfor) { // not required
		return nil
	}

	if m.Forwardfor != nil {
		if err := m.Forwardfor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forwardfor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("forwardfor")
			}
			return err
		}
	}

	return nil
}

var frontendTypeH1CaseAdjustBogusClientPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeH1CaseAdjustBogusClientPropEnum = append(frontendTypeH1CaseAdjustBogusClientPropEnum, v)
	}
}

const (

	// FrontendH1CaseAdjustBogusClientEnabled captures enum value "enabled"
	FrontendH1CaseAdjustBogusClientEnabled string = "enabled"

	// FrontendH1CaseAdjustBogusClientDisabled captures enum value "disabled"
	FrontendH1CaseAdjustBogusClientDisabled string = "disabled"
)

// prop value enum
func (m *Frontend) validateH1CaseAdjustBogusClientEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, frontendTypeH1CaseAdjustBogusClientPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateH1CaseAdjustBogusClient(formats strfmt.Registry) error {
	if swag.IsZero(m.H1CaseAdjustBogusClient) { // not required
		return nil
	}

	// value enum
	if err := m.validateH1CaseAdjustBogusClientEnum("h1_case_adjust_bogus_client", "body", m.H1CaseAdjustBogusClient); err != nil {
		return err
	}

	return nil
}

var frontendTypeHTTPBufferRequestPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeHTTPBufferRequestPropEnum = append(frontendTypeHTTPBufferRequestPropEnum, v)
	}
}

const (

	// FrontendHTTPBufferRequestEnabled captures enum value "enabled"
	FrontendHTTPBufferRequestEnabled string = "enabled"

	// FrontendHTTPBufferRequestDisabled captures enum value "disabled"
	FrontendHTTPBufferRequestDisabled string = "disabled"
)

// prop value enum
func (m *Frontend) validateHTTPBufferRequestEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, frontendTypeHTTPBufferRequestPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateHTTPBufferRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPBufferRequest) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPBufferRequestEnum("http-buffer-request", "body", m.HTTPBufferRequest); err != nil {
		return err
	}

	return nil
}

var frontendTypeHTTPUseHtxPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeHTTPUseHtxPropEnum = append(frontendTypeHTTPUseHtxPropEnum, v)
	}
}

const (

	// FrontendHTTPUseHtxEnabled captures enum value "enabled"
	FrontendHTTPUseHtxEnabled string = "enabled"

	// FrontendHTTPUseHtxDisabled captures enum value "disabled"
	FrontendHTTPUseHtxDisabled string = "disabled"
)

// prop value enum
func (m *Frontend) validateHTTPUseHtxEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, frontendTypeHTTPUseHtxPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateHTTPUseHtx(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPUseHtx) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPUseHtxEnum("http-use-htx", "body", m.HTTPUseHtx); err != nil {
		return err
	}

	return nil
}

var frontendTypeHTTPConnectionModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["httpclose","http-server-close","http-keep-alive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeHTTPConnectionModePropEnum = append(frontendTypeHTTPConnectionModePropEnum, v)
	}
}

const (

	// FrontendHTTPConnectionModeHttpclose captures enum value "httpclose"
	FrontendHTTPConnectionModeHttpclose string = "httpclose"

	// FrontendHTTPConnectionModeHTTPDashServerDashClose captures enum value "http-server-close"
	FrontendHTTPConnectionModeHTTPDashServerDashClose string = "http-server-close"

	// FrontendHTTPConnectionModeHTTPDashKeepDashAlive captures enum value "http-keep-alive"
	FrontendHTTPConnectionModeHTTPDashKeepDashAlive string = "http-keep-alive"
)

// prop value enum
func (m *Frontend) validateHTTPConnectionModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, frontendTypeHTTPConnectionModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateHTTPConnectionMode(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPConnectionMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPConnectionModeEnum("http_connection_mode", "body", m.HTTPConnectionMode); err != nil {
		return err
	}

	return nil
}

var frontendTypeLogSeparateErrorsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeLogSeparateErrorsPropEnum = append(frontendTypeLogSeparateErrorsPropEnum, v)
	}
}

const (

	// FrontendLogSeparateErrorsEnabled captures enum value "enabled"
	FrontendLogSeparateErrorsEnabled string = "enabled"

	// FrontendLogSeparateErrorsDisabled captures enum value "disabled"
	FrontendLogSeparateErrorsDisabled string = "disabled"
)

// prop value enum
func (m *Frontend) validateLogSeparateErrorsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, frontendTypeLogSeparateErrorsPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateLogSeparateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.LogSeparateErrors) { // not required
		return nil
	}

	// value enum
	if err := m.validateLogSeparateErrorsEnum("log_separate_errors", "body", m.LogSeparateErrors); err != nil {
		return err
	}

	return nil
}

func (m *Frontend) validateLogTag(formats strfmt.Registry) error {
	if swag.IsZero(m.LogTag) { // not required
		return nil
	}

	if err := validate.Pattern("log_tag", "body", m.LogTag, `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

var frontendTypeLogasapPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeLogasapPropEnum = append(frontendTypeLogasapPropEnum, v)
	}
}

const (

	// FrontendLogasapEnabled captures enum value "enabled"
	FrontendLogasapEnabled string = "enabled"

	// FrontendLogasapDisabled captures enum value "disabled"
	FrontendLogasapDisabled string = "disabled"
)

// prop value enum
func (m *Frontend) validateLogasapEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, frontendTypeLogasapPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateLogasap(formats strfmt.Registry) error {
	if swag.IsZero(m.Logasap) { // not required
		return nil
	}

	// value enum
	if err := m.validateLogasapEnum("logasap", "body", m.Logasap); err != nil {
		return err
	}

	return nil
}

var frontendTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http","tcp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendTypeModePropEnum = append(frontendTypeModePropEnum, v)
	}
}

const (

	// FrontendModeHTTP captures enum value "http"
	FrontendModeHTTP string = "http"

	// FrontendModeTCP captures enum value "tcp"
	FrontendModeTCP string = "tcp"
)

// prop value enum
func (m *Frontend) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, frontendTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Frontend) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *Frontend) validateMonitorFail(formats strfmt.Registry) error {
	if swag.IsZero(m.MonitorFail) { // not required
		return nil
	}

	if m.MonitorFail != nil {
		if err := m.MonitorFail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monitor_fail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("monitor_fail")
			}
			return err
		}
	}

	return nil
}

func (m *Frontend) validateMonitorURI(formats strfmt.Registry) error {
	if swag.IsZero(m.MonitorURI) { // not required
		return nil
	}

	if err := m.MonitorURI.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("monitor_uri")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("monitor_uri")
		}
		return err
	}

	return nil
}

func (m *Frontend) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Frontend) validateStatsOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.StatsOptions) { // not required
		return nil
	}

	if m.StatsOptions != nil {
		if err := m.StatsOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats_options")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this frontend based on the context it is used
func (m *Frontend) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateForwardfor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMonitorFail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMonitorURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatsOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Frontend) contextValidateForwardfor(ctx context.Context, formats strfmt.Registry) error {

	if m.Forwardfor != nil {
		if err := m.Forwardfor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forwardfor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("forwardfor")
			}
			return err
		}
	}

	return nil
}

func (m *Frontend) contextValidateMonitorFail(ctx context.Context, formats strfmt.Registry) error {

	if m.MonitorFail != nil {
		if err := m.MonitorFail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monitor_fail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("monitor_fail")
			}
			return err
		}
	}

	return nil
}

func (m *Frontend) contextValidateMonitorURI(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MonitorURI.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("monitor_uri")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("monitor_uri")
		}
		return err
	}

	return nil
}

func (m *Frontend) contextValidateStatsOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.StatsOptions != nil {
		if err := m.StatsOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats_options")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Frontend) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Frontend) UnmarshalBinary(b []byte) error {
	var res Frontend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
