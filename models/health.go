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

// Health health
//
// swagger:model health
type Health struct {

	// haproxy
	// Enum: [up down unknown]
	Haproxy string `json:"haproxy,omitempty"`
}

// Validate validates this health
func (m *Health) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHaproxy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var healthTypeHaproxyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["up","down","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		healthTypeHaproxyPropEnum = append(healthTypeHaproxyPropEnum, v)
	}
}

const (

	// HealthHaproxyUp captures enum value "up"
	HealthHaproxyUp string = "up"

	// HealthHaproxyDown captures enum value "down"
	HealthHaproxyDown string = "down"

	// HealthHaproxyUnknown captures enum value "unknown"
	HealthHaproxyUnknown string = "unknown"
)

// prop value enum
func (m *Health) validateHaproxyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, healthTypeHaproxyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Health) validateHaproxy(formats strfmt.Registry) error {
	if swag.IsZero(m.Haproxy) { // not required
		return nil
	}

	// value enum
	if err := m.validateHaproxyEnum("haproxy", "body", m.Haproxy); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this health based on context it is used
func (m *Health) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Health) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Health) UnmarshalBinary(b []byte) error {
	var res Health
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
