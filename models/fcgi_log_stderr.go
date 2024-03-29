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

// FcgiLogStderr Enables logging of STDERR messages that the FastCGI application reports.
// It is an optional setting. By default, HAProxy Enterprise ignores STDERR messages.
//
// swagger:model fcgiLogStderr
type FcgiLogStderr struct {

	// address
	Address string `json:"address,omitempty"`

	// Facility
	Facility string `json:"facility,omitempty"`

	// Format
	Format string `json:"format,omitempty"`

	// Global
	Global bool `json:"global,omitempty"`

	// Length
	Len int64 `json:"len,omitempty"`

	// Level
	Level string `json:"level,omitempty"`

	// Minimum level
	Minlevel string `json:"minlevel,omitempty"`

	// sample
	Sample *FcgiLogStderrSample `json:"sample,omitempty"`
}

// Validate validates this fcgi log stderr
func (m *FcgiLogStderr) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSample(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcgiLogStderr) validateSample(formats strfmt.Registry) error {
	if swag.IsZero(m.Sample) { // not required
		return nil
	}

	if m.Sample != nil {
		if err := m.Sample.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sample")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sample")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fcgi log stderr based on the context it is used
func (m *FcgiLogStderr) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSample(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcgiLogStderr) contextValidateSample(ctx context.Context, formats strfmt.Registry) error {

	if m.Sample != nil {

		if swag.IsZero(m.Sample) { // not required
			return nil
		}

		if err := m.Sample.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sample")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sample")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcgiLogStderr) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcgiLogStderr) UnmarshalBinary(b []byte) error {
	var res FcgiLogStderr
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcgiLogStderrSample Sample
//
// swagger:model FcgiLogStderrSample
type FcgiLogStderrSample struct {

	// Range
	// Required: true
	Ranges *string `json:"ranges"`

	// Size
	// Required: true
	Size *int64 `json:"size"`
}

// Validate validates this fcgi log stderr sample
func (m *FcgiLogStderrSample) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcgiLogStderrSample) validateRanges(formats strfmt.Registry) error {

	if err := validate.Required("sample"+"."+"ranges", "body", m.Ranges); err != nil {
		return err
	}

	return nil
}

func (m *FcgiLogStderrSample) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("sample"+"."+"size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this fcgi log stderr sample based on context it is used
func (m *FcgiLogStderrSample) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FcgiLogStderrSample) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcgiLogStderrSample) UnmarshalBinary(b []byte) error {
	var res FcgiLogStderrSample
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
