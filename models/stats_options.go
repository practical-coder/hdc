// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StatsOptions stats options
//
// swagger:model stats_options
type StatsOptions struct {

	// stats admin
	StatsAdmin bool `json:"stats_admin,omitempty"`

	// stats admin cond
	// Enum: [if unless]
	StatsAdminCond string `json:"stats_admin_cond,omitempty"`

	// stats admin cond test
	StatsAdminCondTest string `json:"stats_admin_cond_test,omitempty"`

	// stats enable
	StatsEnable bool `json:"stats_enable,omitempty"`

	// stats hide version
	StatsHideVersion bool `json:"stats_hide_version,omitempty"`

	// stats maxconn
	// Minimum: 1
	StatsMaxconn int64 `json:"stats_maxconn,omitempty"`

	// stats refresh delay
	StatsRefreshDelay *int64 `json:"stats_refresh_delay,omitempty"`

	// stats show desc
	StatsShowDesc *string `json:"stats_show_desc,omitempty"`

	// stats show legends
	StatsShowLegends bool `json:"stats_show_legends,omitempty"`

	// stats show node name
	// Pattern: ^[^\s]+$
	StatsShowNodeName *string `json:"stats_show_node_name,omitempty"`

	// stats uri prefix
	// Pattern: ^[^\s]+$
	StatsURIPrefix string `json:"stats_uri_prefix,omitempty"`
}

// Validate validates this stats options
func (m *StatsOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatsAdminCond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatsMaxconn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatsShowNodeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatsURIPrefix(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var statsOptionsTypeStatsAdminCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statsOptionsTypeStatsAdminCondPropEnum = append(statsOptionsTypeStatsAdminCondPropEnum, v)
	}
}

const (

	// StatsOptionsStatsAdminCondIf captures enum value "if"
	StatsOptionsStatsAdminCondIf string = "if"

	// StatsOptionsStatsAdminCondUnless captures enum value "unless"
	StatsOptionsStatsAdminCondUnless string = "unless"
)

// prop value enum
func (m *StatsOptions) validateStatsAdminCondEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, statsOptionsTypeStatsAdminCondPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StatsOptions) validateStatsAdminCond(formats strfmt.Registry) error {

	if swag.IsZero(m.StatsAdminCond) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatsAdminCondEnum("stats_admin_cond", "body", m.StatsAdminCond); err != nil {
		return err
	}

	return nil
}

func (m *StatsOptions) validateStatsMaxconn(formats strfmt.Registry) error {

	if swag.IsZero(m.StatsMaxconn) { // not required
		return nil
	}

	if err := validate.MinimumInt("stats_maxconn", "body", int64(m.StatsMaxconn), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *StatsOptions) validateStatsShowNodeName(formats strfmt.Registry) error {

	if swag.IsZero(m.StatsShowNodeName) { // not required
		return nil
	}

	if err := validate.Pattern("stats_show_node_name", "body", string(*m.StatsShowNodeName), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *StatsOptions) validateStatsURIPrefix(formats strfmt.Registry) error {

	if swag.IsZero(m.StatsURIPrefix) { // not required
		return nil
	}

	if err := validate.Pattern("stats_uri_prefix", "body", string(m.StatsURIPrefix), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatsOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatsOptions) UnmarshalBinary(b []byte) error {
	var res StatsOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
