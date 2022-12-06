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

// FcgiPassHeader Specifies the name of a request header to pass to the FastCGI application.
// Optionally, you can follow it with an ACL-based condition, in which case the FastCGI application evaluates it only if the condition is true.
// Most request headers are already available to the FastCGI application with the prefix "HTTP".
// Thus, you only need this directive to pass headers that are purposefully omitted.
// Currently, the headers "Authorization", "Proxy-Authorization", and hop-by-hop headers are omitted.
// Note that the headers "Content-type" and "Content-length" never pass to the FastCGI application because they are already converted into parameters.
//
// swagger:model fcgiPassHeader
type FcgiPassHeader struct {

	// cond
	// Enum: [if unless]
	Cond string `json:"cond,omitempty"`

	// cond test
	CondTest string `json:"cond_test,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this fcgi pass header
func (m *FcgiPassHeader) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCond(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var fcgiPassHeaderTypeCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fcgiPassHeaderTypeCondPropEnum = append(fcgiPassHeaderTypeCondPropEnum, v)
	}
}

const (

	// FcgiPassHeaderCondIf captures enum value "if"
	FcgiPassHeaderCondIf string = "if"

	// FcgiPassHeaderCondUnless captures enum value "unless"
	FcgiPassHeaderCondUnless string = "unless"
)

// prop value enum
func (m *FcgiPassHeader) validateCondEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fcgiPassHeaderTypeCondPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FcgiPassHeader) validateCond(formats strfmt.Registry) error {
	if swag.IsZero(m.Cond) { // not required
		return nil
	}

	// value enum
	if err := m.validateCondEnum("cond", "body", m.Cond); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this fcgi pass header based on context it is used
func (m *FcgiPassHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FcgiPassHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcgiPassHeader) UnmarshalBinary(b []byte) error {
	var res FcgiPassHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
