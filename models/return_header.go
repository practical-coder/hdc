// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReturnHeader return header
//
// swagger:model ReturnHeader
type ReturnHeader struct {

	// fmt
	// Required: true
	Fmt *string `json:"fmt"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this return header
func (m *ReturnHeader) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnHeader) validateFmt(formats strfmt.Registry) error {

	if err := validate.Required("fmt", "body", m.Fmt); err != nil {
		return err
	}

	return nil
}

func (m *ReturnHeader) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnHeader) UnmarshalBinary(b []byte) error {
	var res ReturnHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
