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

// Server Server
//
// HAProxy backend server configuration
// Example: {"address":"10.1.1.1","check":"enabled","name":"www","port":8080,"weight":80}
//
// swagger:model server
type Server struct {
	ServerParams

	// address
	// Required: true
	// Pattern: ^[^\s]+$
	Address string `json:"address"`

	// id
	ID *int64 `json:"id,omitempty"`

	// name
	// Required: true
	// Pattern: ^[^\s]+$
	Name string `json:"name"`

	// port
	// Maximum: 65535
	// Minimum: 1
	Port *int64 `json:"port,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Server) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ServerParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ServerParams = aO0

	// now for regular properties
	var propsServer struct {
		Address string `json:"address"`

		ID *int64 `json:"id,omitempty"`

		Name string `json:"name"`

		Port *int64 `json:"port,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsServer); err != nil {
		return err
	}
	m.Address = propsServer.Address

	m.ID = propsServer.ID

	m.Name = propsServer.Name

	m.Port = propsServer.Port

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Server) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ServerParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsServer struct {
		Address string `json:"address"`

		ID *int64 `json:"id,omitempty"`

		Name string `json:"name"`

		Port *int64 `json:"port,omitempty"`
	}
	propsServer.Address = m.Address

	propsServer.ID = m.ID

	propsServer.Name = m.Name

	propsServer.Port = m.Port

	jsonDataPropsServer, errServer := swag.WriteJSON(propsServer)
	if errServer != nil {
		return nil, errServer
	}
	_parts = append(_parts, jsonDataPropsServer)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this server
func (m *Server) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ServerParams
	if err := m.ServerParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Server) validateAddress(formats strfmt.Registry) error {

	if err := validate.RequiredString("address", "body", m.Address); err != nil {
		return err
	}

	if err := validate.Pattern("address", "body", m.Address, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Server) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Server) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", *m.Port, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", *m.Port, 65535, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this server based on the context it is used
func (m *Server) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
func (m *Server) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Server) UnmarshalBinary(b []byte) error {
	var res Server
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
