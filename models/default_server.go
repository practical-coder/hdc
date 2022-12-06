// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DefaultServer Default Server
//
// swagger:model default_server
type DefaultServer struct {
	ServerParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DefaultServer) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ServerParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ServerParams = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DefaultServer) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ServerParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this default server
func (m *DefaultServer) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ServerParams
	if err := m.ServerParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this default server based on the context it is used
func (m *DefaultServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ServerParams
	if err := m.ServerParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DefaultServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DefaultServer) UnmarshalBinary(b []byte) error {
	var res DefaultServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
