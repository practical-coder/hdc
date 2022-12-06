// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Map Map File
//
// # Map File
//
// swagger:model map
type Map struct {

	// description
	Description string `json:"description,omitempty"`

	// file
	File string `json:"file,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// storage name
	StorageName string `json:"storage_name,omitempty"`
}

// Validate validates this map
func (m *Map) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this map based on context it is used
func (m *Map) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Map) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Map) UnmarshalBinary(b []byte) error {
	var res Map
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
