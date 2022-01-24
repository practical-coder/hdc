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

// SslCertEntry One SSL Certificate Entry
//
// One SSL/TLS certificate
//
// swagger:model ssl_cert_entry
type SslCertEntry struct {

	// algorithm
	Algorithm string `json:"algorithm,omitempty"`

	// chain issuer
	ChainIssuer string `json:"chain_issuer,omitempty"`

	// chain subject
	ChainSubject string `json:"chain_subject,omitempty"`

	// issuer
	Issuer string `json:"issuer,omitempty"`

	// not after
	// Format: date
	NotAfter strfmt.Date `json:"not_after,omitempty"`

	// not before
	// Format: date
	NotBefore strfmt.Date `json:"not_before,omitempty"`

	// serial
	Serial string `json:"serial,omitempty"`

	// sha1 finger print
	Sha1FingerPrint string `json:"sha1_finger_print,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// storage name
	StorageName string `json:"storage_name,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`

	// subject alternative names
	SubjectAlternativeNames []string `json:"subject_alternative_names"`
}

// Validate validates this ssl cert entry
func (m *SslCertEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotBefore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SslCertEntry) validateNotAfter(formats strfmt.Registry) error {

	if swag.IsZero(m.NotAfter) { // not required
		return nil
	}

	if err := validate.FormatOf("not_after", "body", "date", m.NotAfter.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SslCertEntry) validateNotBefore(formats strfmt.Registry) error {

	if swag.IsZero(m.NotBefore) { // not required
		return nil
	}

	if err := validate.FormatOf("not_before", "body", "date", m.NotBefore.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SslCertEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SslCertEntry) UnmarshalBinary(b []byte) error {
	var res SslCertEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
