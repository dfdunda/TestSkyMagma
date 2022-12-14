// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VirtualApnRule Virtual APN Rule configuration
//
// swagger:model virtual_apn_rule
type VirtualApnRule struct {

	// Regex to match APN values
	// Example: .*\\.magma.*
	ApnFilter string `json:"apn_filter,omitempty"`

	// New APN to overwrite when filter matches
	// Example: new.magma.apn
	ApnOverwrite string `json:"apn_overwrite,omitempty"`

	// Regex to match ChargingCharacteristics values
	// Example: 12
	ChargingCharacteristicsFilter string `json:"charging_characteristics_filter,omitempty"`
}

// Validate validates this virtual apn rule
func (m *VirtualApnRule) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this virtual apn rule based on context it is used
func (m *VirtualApnRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualApnRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualApnRule) UnmarshalBinary(b []byte) error {
	var res VirtualApnRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
