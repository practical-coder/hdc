// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HTTPResponseRule HTTP Response Rule
//
// HAProxy HTTP response rule configuration (corresponds to http-response directives)
//
// swagger:model http_response_rule
type HTTPResponseRule struct {

	// acl file
	// Pattern: ^[^\s]+$
	ACLFile string `json:"acl_file,omitempty"`

	// acl keyfmt
	// Pattern: ^[^\s]+$
	ACLKeyfmt string `json:"acl_keyfmt,omitempty"`

	// capture id
	CaptureID *int64 `json:"capture_id,omitempty"`

	// capture sample
	// Pattern: ^[^\s]+$
	CaptureSample string `json:"capture_sample,omitempty"`

	// cond
	// Enum: [if unless]
	Cond string `json:"cond,omitempty"`

	// cond test
	CondTest string `json:"cond_test,omitempty"`

	// hdr format
	HdrFormat string `json:"hdr_format,omitempty"`

	// hdr match
	HdrMatch string `json:"hdr_match,omitempty"`

	// hdr name
	HdrName string `json:"hdr_name,omitempty"`

	// index
	// Required: true
	Index *int64 `json:"index"`

	// log level
	// Enum: [emerg alert crit err warning notice info debug silent]
	LogLevel string `json:"log_level,omitempty"`

	// lua action
	// Pattern: ^[^\s]+$
	LuaAction string `json:"lua_action,omitempty"`

	// lua params
	LuaParams string `json:"lua_params,omitempty"`

	// map file
	// Pattern: ^[^\s]+$
	MapFile string `json:"map_file,omitempty"`

	// map keyfmt
	// Pattern: ^[^\s]+$
	MapKeyfmt string `json:"map_keyfmt,omitempty"`

	// map valuefmt
	// Pattern: ^[^\s]+$
	MapValuefmt string `json:"map_valuefmt,omitempty"`

	// mark value
	// Pattern: ^(0x[0-9A-Fa-f]+|[0-9]+)$
	MarkValue string `json:"mark_value,omitempty"`

	// nice value
	// Maximum: 1024
	// Minimum: -1024
	NiceValue int64 `json:"nice_value,omitempty"`

	// redir code
	// Enum: [301 302 303 307 308]
	RedirCode *int64 `json:"redir_code,omitempty"`

	// redir option
	RedirOption string `json:"redir_option,omitempty"`

	// redir type
	// Enum: [location prefix scheme]
	RedirType string `json:"redir_type,omitempty"`

	// redir value
	// Pattern: ^[^\s]+$
	RedirValue string `json:"redir_value,omitempty"`

	// sc expr
	ScExpr string `json:"sc_expr,omitempty"`

	// sc id
	ScID int64 `json:"sc_id,omitempty"`

	// sc int
	ScInt *int64 `json:"sc_int,omitempty"`

	// spoe engine
	// Pattern: ^[^\s]+$
	SpoeEngine string `json:"spoe_engine,omitempty"`

	// spoe group
	// Pattern: ^[^\s]+$
	SpoeGroup string `json:"spoe_group,omitempty"`

	// status
	// Maximum: 999
	// Minimum: 100
	Status int64 `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"status_reason,omitempty"`

	// strict mode
	// Enum: [on off]
	StrictMode string `json:"strict_mode,omitempty"`

	// tos value
	// Pattern: ^(0x[0-9A-Fa-f]+|[0-9]+)$
	TosValue string `json:"tos_value,omitempty"`

	// track sc0 key
	// Pattern: ^[^\s]+$
	TrackSc0Key string `json:"track-sc0-key,omitempty"`

	// track sc0 table
	// Pattern: ^[^\s]+$
	TrackSc0Table string `json:"track-sc0-table,omitempty"`

	// track sc1 key
	// Pattern: ^[^\s]+$
	TrackSc1Key string `json:"track-sc1-key,omitempty"`

	// track sc1 table
	// Pattern: ^[^\s]+$
	TrackSc1Table string `json:"track-sc1-table,omitempty"`

	// track sc2 key
	// Pattern: ^[^\s]+$
	TrackSc2Key string `json:"track-sc2-key,omitempty"`

	// track sc2 table
	// Pattern: ^[^\s]+$
	TrackSc2Table string `json:"track-sc2-table,omitempty"`

	// type
	// Required: true
	// Enum: [allow deny redirect add-header set-header del-header set-log-level set-var set-status send-spoe-group replace-header replace-value add-acl del-acl capture set-map del-map sc-inc-gpc0 sc-inc-gpc1 sc-set-gpt0 set-mark set-nice set-tos silent-drop unset-var track-sc0 track-sc1 track-sc2 strict-mode lua]
	Type string `json:"type"`

	// var expr
	VarExpr string `json:"var_expr,omitempty"`

	// var name
	// Pattern: ^[^\s]+$
	VarName string `json:"var_name,omitempty"`

	// var scope
	// Pattern: ^[^\s]+$
	VarScope string `json:"var_scope,omitempty"`
}

// Validate validates this http response rule
func (m *HTTPResponseRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateACLFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateACLKeyfmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCaptureSample(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLuaAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapKeyfmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapValuefmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarkValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNiceValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpoeEngine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpoeGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStrictMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTosValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc0Key(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc0Table(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc1Key(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc1Table(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc2Key(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc2Table(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVarName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVarScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPResponseRule) validateACLFile(formats strfmt.Registry) error {

	if swag.IsZero(m.ACLFile) { // not required
		return nil
	}

	if err := validate.Pattern("acl_file", "body", string(m.ACLFile), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateACLKeyfmt(formats strfmt.Registry) error {

	if swag.IsZero(m.ACLKeyfmt) { // not required
		return nil
	}

	if err := validate.Pattern("acl_keyfmt", "body", string(m.ACLKeyfmt), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateCaptureSample(formats strfmt.Registry) error {

	if swag.IsZero(m.CaptureSample) { // not required
		return nil
	}

	if err := validate.Pattern("capture_sample", "body", string(m.CaptureSample), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeCondPropEnum = append(httpResponseRuleTypeCondPropEnum, v)
	}
}

const (

	// HTTPResponseRuleCondIf captures enum value "if"
	HTTPResponseRuleCondIf string = "if"

	// HTTPResponseRuleCondUnless captures enum value "unless"
	HTTPResponseRuleCondUnless string = "unless"
)

// prop value enum
func (m *HTTPResponseRule) validateCondEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, httpResponseRuleTypeCondPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateCond(formats strfmt.Registry) error {

	if swag.IsZero(m.Cond) { // not required
		return nil
	}

	// value enum
	if err := m.validateCondEnum("cond", "body", m.Cond); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeLogLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["emerg","alert","crit","err","warning","notice","info","debug","silent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeLogLevelPropEnum = append(httpResponseRuleTypeLogLevelPropEnum, v)
	}
}

const (

	// HTTPResponseRuleLogLevelEmerg captures enum value "emerg"
	HTTPResponseRuleLogLevelEmerg string = "emerg"

	// HTTPResponseRuleLogLevelAlert captures enum value "alert"
	HTTPResponseRuleLogLevelAlert string = "alert"

	// HTTPResponseRuleLogLevelCrit captures enum value "crit"
	HTTPResponseRuleLogLevelCrit string = "crit"

	// HTTPResponseRuleLogLevelErr captures enum value "err"
	HTTPResponseRuleLogLevelErr string = "err"

	// HTTPResponseRuleLogLevelWarning captures enum value "warning"
	HTTPResponseRuleLogLevelWarning string = "warning"

	// HTTPResponseRuleLogLevelNotice captures enum value "notice"
	HTTPResponseRuleLogLevelNotice string = "notice"

	// HTTPResponseRuleLogLevelInfo captures enum value "info"
	HTTPResponseRuleLogLevelInfo string = "info"

	// HTTPResponseRuleLogLevelDebug captures enum value "debug"
	HTTPResponseRuleLogLevelDebug string = "debug"

	// HTTPResponseRuleLogLevelSilent captures enum value "silent"
	HTTPResponseRuleLogLevelSilent string = "silent"
)

// prop value enum
func (m *HTTPResponseRule) validateLogLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, httpResponseRuleTypeLogLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateLogLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.LogLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateLogLevelEnum("log_level", "body", m.LogLevel); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateLuaAction(formats strfmt.Registry) error {

	if swag.IsZero(m.LuaAction) { // not required
		return nil
	}

	if err := validate.Pattern("lua_action", "body", string(m.LuaAction), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateMapFile(formats strfmt.Registry) error {

	if swag.IsZero(m.MapFile) { // not required
		return nil
	}

	if err := validate.Pattern("map_file", "body", string(m.MapFile), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateMapKeyfmt(formats strfmt.Registry) error {

	if swag.IsZero(m.MapKeyfmt) { // not required
		return nil
	}

	if err := validate.Pattern("map_keyfmt", "body", string(m.MapKeyfmt), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateMapValuefmt(formats strfmt.Registry) error {

	if swag.IsZero(m.MapValuefmt) { // not required
		return nil
	}

	if err := validate.Pattern("map_valuefmt", "body", string(m.MapValuefmt), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateMarkValue(formats strfmt.Registry) error {

	if swag.IsZero(m.MarkValue) { // not required
		return nil
	}

	if err := validate.Pattern("mark_value", "body", string(m.MarkValue), `^(0x[0-9A-Fa-f]+|[0-9]+)$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateNiceValue(formats strfmt.Registry) error {

	if swag.IsZero(m.NiceValue) { // not required
		return nil
	}

	if err := validate.MinimumInt("nice_value", "body", int64(m.NiceValue), -1024, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("nice_value", "body", int64(m.NiceValue), 1024, false); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeRedirCodePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[301,302,303,307,308]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeRedirCodePropEnum = append(httpResponseRuleTypeRedirCodePropEnum, v)
	}
}

// prop value enum
func (m *HTTPResponseRule) validateRedirCodeEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, httpResponseRuleTypeRedirCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateRedirCode(formats strfmt.Registry) error {

	if swag.IsZero(m.RedirCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateRedirCodeEnum("redir_code", "body", *m.RedirCode); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeRedirTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["location","prefix","scheme"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeRedirTypePropEnum = append(httpResponseRuleTypeRedirTypePropEnum, v)
	}
}

const (

	// HTTPResponseRuleRedirTypeLocation captures enum value "location"
	HTTPResponseRuleRedirTypeLocation string = "location"

	// HTTPResponseRuleRedirTypePrefix captures enum value "prefix"
	HTTPResponseRuleRedirTypePrefix string = "prefix"

	// HTTPResponseRuleRedirTypeScheme captures enum value "scheme"
	HTTPResponseRuleRedirTypeScheme string = "scheme"
)

// prop value enum
func (m *HTTPResponseRule) validateRedirTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, httpResponseRuleTypeRedirTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateRedirType(formats strfmt.Registry) error {

	if swag.IsZero(m.RedirType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRedirTypeEnum("redir_type", "body", m.RedirType); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateRedirValue(formats strfmt.Registry) error {

	if swag.IsZero(m.RedirValue) { // not required
		return nil
	}

	if err := validate.Pattern("redir_value", "body", string(m.RedirValue), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateSpoeEngine(formats strfmt.Registry) error {

	if swag.IsZero(m.SpoeEngine) { // not required
		return nil
	}

	if err := validate.Pattern("spoe_engine", "body", string(m.SpoeEngine), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateSpoeGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.SpoeGroup) { // not required
		return nil
	}

	if err := validate.Pattern("spoe_group", "body", string(m.SpoeGroup), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := validate.MinimumInt("status", "body", int64(m.Status), 100, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("status", "body", int64(m.Status), 999, false); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeStrictModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on","off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeStrictModePropEnum = append(httpResponseRuleTypeStrictModePropEnum, v)
	}
}

const (

	// HTTPResponseRuleStrictModeOn captures enum value "on"
	HTTPResponseRuleStrictModeOn string = "on"

	// HTTPResponseRuleStrictModeOff captures enum value "off"
	HTTPResponseRuleStrictModeOff string = "off"
)

// prop value enum
func (m *HTTPResponseRule) validateStrictModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, httpResponseRuleTypeStrictModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateStrictMode(formats strfmt.Registry) error {

	if swag.IsZero(m.StrictMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateStrictModeEnum("strict_mode", "body", m.StrictMode); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTosValue(formats strfmt.Registry) error {

	if swag.IsZero(m.TosValue) { // not required
		return nil
	}

	if err := validate.Pattern("tos_value", "body", string(m.TosValue), `^(0x[0-9A-Fa-f]+|[0-9]+)$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc0Key(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc0Key) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc0-key", "body", string(m.TrackSc0Key), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc0Table(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc0Table) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc0-table", "body", string(m.TrackSc0Table), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc1Key(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc1Key) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc1-key", "body", string(m.TrackSc1Key), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc1Table(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc1Table) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc1-table", "body", string(m.TrackSc1Table), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc2Key(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc2Key) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc2-key", "body", string(m.TrackSc2Key), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateTrackSc2Table(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc2Table) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc2-table", "body", string(m.TrackSc2Table), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["allow","deny","redirect","add-header","set-header","del-header","set-log-level","set-var","set-status","send-spoe-group","replace-header","replace-value","add-acl","del-acl","capture","set-map","del-map","sc-inc-gpc0","sc-inc-gpc1","sc-set-gpt0","set-mark","set-nice","set-tos","silent-drop","unset-var","track-sc0","track-sc1","track-sc2","strict-mode","lua"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeTypePropEnum = append(httpResponseRuleTypeTypePropEnum, v)
	}
}

const (

	// HTTPResponseRuleTypeAllow captures enum value "allow"
	HTTPResponseRuleTypeAllow string = "allow"

	// HTTPResponseRuleTypeDeny captures enum value "deny"
	HTTPResponseRuleTypeDeny string = "deny"

	// HTTPResponseRuleTypeRedirect captures enum value "redirect"
	HTTPResponseRuleTypeRedirect string = "redirect"

	// HTTPResponseRuleTypeAddHeader captures enum value "add-header"
	HTTPResponseRuleTypeAddHeader string = "add-header"

	// HTTPResponseRuleTypeSetHeader captures enum value "set-header"
	HTTPResponseRuleTypeSetHeader string = "set-header"

	// HTTPResponseRuleTypeDelHeader captures enum value "del-header"
	HTTPResponseRuleTypeDelHeader string = "del-header"

	// HTTPResponseRuleTypeSetLogLevel captures enum value "set-log-level"
	HTTPResponseRuleTypeSetLogLevel string = "set-log-level"

	// HTTPResponseRuleTypeSetVar captures enum value "set-var"
	HTTPResponseRuleTypeSetVar string = "set-var"

	// HTTPResponseRuleTypeSetStatus captures enum value "set-status"
	HTTPResponseRuleTypeSetStatus string = "set-status"

	// HTTPResponseRuleTypeSendSpoeGroup captures enum value "send-spoe-group"
	HTTPResponseRuleTypeSendSpoeGroup string = "send-spoe-group"

	// HTTPResponseRuleTypeReplaceHeader captures enum value "replace-header"
	HTTPResponseRuleTypeReplaceHeader string = "replace-header"

	// HTTPResponseRuleTypeReplaceValue captures enum value "replace-value"
	HTTPResponseRuleTypeReplaceValue string = "replace-value"

	// HTTPResponseRuleTypeAddACL captures enum value "add-acl"
	HTTPResponseRuleTypeAddACL string = "add-acl"

	// HTTPResponseRuleTypeDelACL captures enum value "del-acl"
	HTTPResponseRuleTypeDelACL string = "del-acl"

	// HTTPResponseRuleTypeCapture captures enum value "capture"
	HTTPResponseRuleTypeCapture string = "capture"

	// HTTPResponseRuleTypeSetMap captures enum value "set-map"
	HTTPResponseRuleTypeSetMap string = "set-map"

	// HTTPResponseRuleTypeDelMap captures enum value "del-map"
	HTTPResponseRuleTypeDelMap string = "del-map"

	// HTTPResponseRuleTypeScIncGpc0 captures enum value "sc-inc-gpc0"
	HTTPResponseRuleTypeScIncGpc0 string = "sc-inc-gpc0"

	// HTTPResponseRuleTypeScIncGpc1 captures enum value "sc-inc-gpc1"
	HTTPResponseRuleTypeScIncGpc1 string = "sc-inc-gpc1"

	// HTTPResponseRuleTypeScSetGpt0 captures enum value "sc-set-gpt0"
	HTTPResponseRuleTypeScSetGpt0 string = "sc-set-gpt0"

	// HTTPResponseRuleTypeSetMark captures enum value "set-mark"
	HTTPResponseRuleTypeSetMark string = "set-mark"

	// HTTPResponseRuleTypeSetNice captures enum value "set-nice"
	HTTPResponseRuleTypeSetNice string = "set-nice"

	// HTTPResponseRuleTypeSetTos captures enum value "set-tos"
	HTTPResponseRuleTypeSetTos string = "set-tos"

	// HTTPResponseRuleTypeSilentDrop captures enum value "silent-drop"
	HTTPResponseRuleTypeSilentDrop string = "silent-drop"

	// HTTPResponseRuleTypeUnsetVar captures enum value "unset-var"
	HTTPResponseRuleTypeUnsetVar string = "unset-var"

	// HTTPResponseRuleTypeTrackSc0 captures enum value "track-sc0"
	HTTPResponseRuleTypeTrackSc0 string = "track-sc0"

	// HTTPResponseRuleTypeTrackSc1 captures enum value "track-sc1"
	HTTPResponseRuleTypeTrackSc1 string = "track-sc1"

	// HTTPResponseRuleTypeTrackSc2 captures enum value "track-sc2"
	HTTPResponseRuleTypeTrackSc2 string = "track-sc2"

	// HTTPResponseRuleTypeStrictMode captures enum value "strict-mode"
	HTTPResponseRuleTypeStrictMode string = "strict-mode"

	// HTTPResponseRuleTypeLua captures enum value "lua"
	HTTPResponseRuleTypeLua string = "lua"
)

// prop value enum
func (m *HTTPResponseRule) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, httpResponseRuleTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateVarName(formats strfmt.Registry) error {

	if swag.IsZero(m.VarName) { // not required
		return nil
	}

	if err := validate.Pattern("var_name", "body", string(m.VarName), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateVarScope(formats strfmt.Registry) error {

	if swag.IsZero(m.VarScope) { // not required
		return nil
	}

	if err := validate.Pattern("var_scope", "body", string(m.VarScope), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPResponseRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPResponseRule) UnmarshalBinary(b []byte) error {
	var res HTTPResponseRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
