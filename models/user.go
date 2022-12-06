// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// User User
//
// # HAProxy userlist user
//
// swagger:model user
type User struct {

	// groups
	Groups string `json:"groups,omitempty"`

	// password
	// Required: true
	Password string `json:"password"`

	// secure password
	// Required: true
	SecurePassword *bool `json:"secure_password"`

	// username
	// Required: true
	// Pattern: ^[A-Za-z0-9-_.:]+$
	Username string `json:"username"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validatePassword(formats strfmt.Registry) error {

	if err := validate.RequiredString("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *User) validateSecurePassword(formats strfmt.Registry) error {

	if err := validate.Required("secure_password", "body", m.SecurePassword); err != nil {
		return err
	}

	return nil
}

func (m *User) validateUsername(formats strfmt.Registry) error {

	if err := validate.RequiredString("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.Pattern("username", "body", m.Username, `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user based on context it is used
func (m *User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
