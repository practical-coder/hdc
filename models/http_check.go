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

// HTTPCheck http check
//
// swagger:model http-check
type HTTPCheck struct {

	// exclamation mark
	ExclamationMark bool `json:"exclamation_mark,omitempty"`

	// match
	// Pattern: ^[^\s]+$
	// Enum: [status rstatus string rstring]
	Match string `json:"match,omitempty"`

	// pattern
	Pattern string `json:"pattern,omitempty"`

	// type
	// Required: true
	// Enum: [disable-on-404 expect send-state]
	Type *string `json:"type"`
}

// Validate validates this http check
func (m *HTTPCheck) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var httpCheckTypeMatchPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["status","rstatus","string","rstring"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpCheckTypeMatchPropEnum = append(httpCheckTypeMatchPropEnum, v)
	}
}

const (

	// HTTPCheckMatchStatus captures enum value "status"
	HTTPCheckMatchStatus string = "status"

	// HTTPCheckMatchRstatus captures enum value "rstatus"
	HTTPCheckMatchRstatus string = "rstatus"

	// HTTPCheckMatchString captures enum value "string"
	HTTPCheckMatchString string = "string"

	// HTTPCheckMatchRstring captures enum value "rstring"
	HTTPCheckMatchRstring string = "rstring"
)

// prop value enum
func (m *HTTPCheck) validateMatchEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, httpCheckTypeMatchPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPCheck) validateMatch(formats strfmt.Registry) error {

	if swag.IsZero(m.Match) { // not required
		return nil
	}

	if err := validate.Pattern("match", "body", string(m.Match), `^[^\s]+$`); err != nil {
		return err
	}

	// value enum
	if err := m.validateMatchEnum("match", "body", m.Match); err != nil {
		return err
	}

	return nil
}

var httpCheckTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["disable-on-404","expect","send-state"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpCheckTypeTypePropEnum = append(httpCheckTypeTypePropEnum, v)
	}
}

const (

	// HTTPCheckTypeDisableOn404 captures enum value "disable-on-404"
	HTTPCheckTypeDisableOn404 string = "disable-on-404"

	// HTTPCheckTypeExpect captures enum value "expect"
	HTTPCheckTypeExpect string = "expect"

	// HTTPCheckTypeSendState captures enum value "send-state"
	HTTPCheckTypeSendState string = "send-state"
)

// prop value enum
func (m *HTTPCheck) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, httpCheckTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPCheck) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPCheck) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPCheck) UnmarshalBinary(b []byte) error {
	var res HTTPCheck
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
