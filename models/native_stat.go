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

// NativeStat Stats
//
// Current stats for one object.
//
// swagger:model native_stat
type NativeStat struct {

	// backend name
	BackendName string `json:"backend_name,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// stats
	Stats *NativeStatStats `json:"stats,omitempty"`

	// type
	// Enum: [backend server frontend]
	Type string `json:"type,omitempty"`
}

// Validate validates this native stat
func (m *NativeStat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStats(formats); err != nil {
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

func (m *NativeStat) validateStats(formats strfmt.Registry) error {
	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

var nativeStatTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["backend","server","frontend"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nativeStatTypeTypePropEnum = append(nativeStatTypeTypePropEnum, v)
	}
}

const (

	// NativeStatTypeBackend captures enum value "backend"
	NativeStatTypeBackend string = "backend"

	// NativeStatTypeServer captures enum value "server"
	NativeStatTypeServer string = "server"

	// NativeStatTypeFrontend captures enum value "frontend"
	NativeStatTypeFrontend string = "frontend"
)

// prop value enum
func (m *NativeStat) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nativeStatTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NativeStat) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this native stat based on the context it is used
func (m *NativeStat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NativeStat) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

	if m.Stats != nil {

		if swag.IsZero(m.Stats) { // not required
			return nil
		}

		if err := m.Stats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NativeStat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NativeStat) UnmarshalBinary(b []byte) error {
	var res NativeStat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
