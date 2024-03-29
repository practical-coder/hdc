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

// Errorloc errorloc
//
// swagger:model errorloc
type Errorloc struct {

	// code
	// Required: true
	// Enum: [200 400 401 403 404 405 407 408 410 413 425 429 500 501 502 503 504]
	Code *int64 `json:"code"`

	// url
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this errorloc
func (m *Errorloc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var errorlocTypeCodePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[200,400,401,403,404,405,407,408,410,413,425,429,500,501,502,503,504]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		errorlocTypeCodePropEnum = append(errorlocTypeCodePropEnum, v)
	}
}

// prop value enum
func (m *Errorloc) validateCodeEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, errorlocTypeCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Errorloc) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	// value enum
	if err := m.validateCodeEnum("code", "body", *m.Code); err != nil {
		return err
	}

	return nil
}

func (m *Errorloc) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this errorloc based on context it is used
func (m *Errorloc) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Errorloc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Errorloc) UnmarshalBinary(b []byte) error {
	var res Errorloc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
