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

// Enodeb Representation of an enodeB
//
// swagger:model enodeb
type Enodeb struct {

	// attached gateway id
	// Example: gw1
	// Read Only: true
	AttachedGatewayID string `json:"attached_gateway_id,omitempty"`

	// config
	// Required: true
	Config *EnodebConfiguration `json:"config"`

	// description
	// Example: Description of this radio
	Description string `json:"description,omitempty"`

	// enodeb config
	EnodebConfig *EnodebConfig `json:"enodeb_config,omitempty"`

	// name
	// Example: Foobar EnodeB
	// Required: true
	// Min Length: 1
	Name string `json:"name"`

	// serial
	// Example: 1202000038269KP0037
	// Required: true
	// Min Length: 1
	Serial string `json:"serial"`
}

// Validate validates this enodeb
func (m *Enodeb) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnodebConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Enodeb) validateConfig(formats strfmt.Registry) error {

	if err := validate.Required("config", "body", m.Config); err != nil {
		return err
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *Enodeb) validateEnodebConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.EnodebConfig) { // not required
		return nil
	}

	if m.EnodebConfig != nil {
		if err := m.EnodebConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enodeb_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enodeb_config")
			}
			return err
		}
	}

	return nil
}

func (m *Enodeb) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	return nil
}

func (m *Enodeb) validateSerial(formats strfmt.Registry) error {

	if err := validate.RequiredString("serial", "body", m.Serial); err != nil {
		return err
	}

	if err := validate.MinLength("serial", "body", m.Serial, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this enodeb based on the context it is used
func (m *Enodeb) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachedGatewayID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnodebConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Enodeb) contextValidateAttachedGatewayID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "attached_gateway_id", "body", string(m.AttachedGatewayID)); err != nil {
		return err
	}

	return nil
}

func (m *Enodeb) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {
		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *Enodeb) contextValidateEnodebConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.EnodebConfig != nil {
		if err := m.EnodebConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enodeb_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enodeb_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Enodeb) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Enodeb) UnmarshalBinary(b []byte) error {
	var res Enodeb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
