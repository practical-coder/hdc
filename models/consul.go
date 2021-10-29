// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Consul Consul server
//
// Consul server configuration
// Example: {"address":"127.0.0.1","enabled":true,"id":"0","port":90,"retry_timeout":10}
//
// swagger:model consul
type Consul struct {

	// address
	// Required: true
	// Pattern: ^[^\s]+$
	Address *string `json:"address"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// Auto generated ID.
	// Pattern: ^[^\s]+$
	ID *string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// port
	// Required: true
	// Maximum: 65535
	// Minimum: 1
	Port *int64 `json:"port"`

	// Duration in seconds in-between data pulling requests to the consul server
	// Required: true
	// Minimum: 1
	RetryTimeout *int64 `json:"retry_timeout"`

	// server slots base
	ServerSlotsBase *int64 `json:"server_slots_base,omitempty"`

	// server slots growth increment
	ServerSlotsGrowthIncrement int64 `json:"server_slots_growth_increment,omitempty"`

	// server slots growth type
	// Enum: [linear exponential]
	ServerSlotsGrowthType *string `json:"server_slots_growth_type,omitempty"`

	// service blacklist
	ServiceBlacklist []string `json:"service-blacklist"`

	// service whitelist
	ServiceWhitelist []string `json:"service-whitelist"`

	// token
	// Pattern: ^[^\s]+$
	Token string `json:"token,omitempty"`
}

// Validate validates this consul
func (m *Consul) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetryTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerSlotsGrowthType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceBlacklist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceWhitelist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Consul) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	if err := validate.Pattern("address", "body", *m.Address, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Consul) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *Consul) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", *m.ID, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Consul) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	if err := validate.MinimumInt("port", "body", *m.Port, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", *m.Port, 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *Consul) validateRetryTimeout(formats strfmt.Registry) error {

	if err := validate.Required("retry_timeout", "body", m.RetryTimeout); err != nil {
		return err
	}

	if err := validate.MinimumInt("retry_timeout", "body", *m.RetryTimeout, 1, false); err != nil {
		return err
	}

	return nil
}

var consulTypeServerSlotsGrowthTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["linear","exponential"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consulTypeServerSlotsGrowthTypePropEnum = append(consulTypeServerSlotsGrowthTypePropEnum, v)
	}
}

const (

	// ConsulServerSlotsGrowthTypeLinear captures enum value "linear"
	ConsulServerSlotsGrowthTypeLinear string = "linear"

	// ConsulServerSlotsGrowthTypeExponential captures enum value "exponential"
	ConsulServerSlotsGrowthTypeExponential string = "exponential"
)

// prop value enum
func (m *Consul) validateServerSlotsGrowthTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consulTypeServerSlotsGrowthTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Consul) validateServerSlotsGrowthType(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerSlotsGrowthType) { // not required
		return nil
	}

	// value enum
	if err := m.validateServerSlotsGrowthTypeEnum("server_slots_growth_type", "body", *m.ServerSlotsGrowthType); err != nil {
		return err
	}

	return nil
}

func (m *Consul) validateServiceBlacklist(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceBlacklist) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceBlacklist); i++ {

		if err := validate.Pattern("service-blacklist"+"."+strconv.Itoa(i), "body", m.ServiceBlacklist[i], `^[^\s]+$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *Consul) validateServiceWhitelist(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceWhitelist) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceWhitelist); i++ {

		if err := validate.Pattern("service-whitelist"+"."+strconv.Itoa(i), "body", m.ServiceWhitelist[i], `^[^\s]+$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *Consul) validateToken(formats strfmt.Registry) error {
	if swag.IsZero(m.Token) { // not required
		return nil
	}

	if err := validate.Pattern("token", "body", m.Token, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this consul based on context it is used
func (m *Consul) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Consul) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Consul) UnmarshalBinary(b []byte) error {
	var res Consul
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
