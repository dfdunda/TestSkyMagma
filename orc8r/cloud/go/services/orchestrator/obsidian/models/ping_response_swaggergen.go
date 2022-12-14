// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PingResponse ping response
//
// swagger:model ping_response
type PingResponse struct {

	// pings
	// Required: true
	Pings []*PingResult `json:"pings"`
}

// Validate validates this ping response
func (m *PingResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PingResponse) validatePings(formats strfmt.Registry) error {

	if err := validate.Required("pings", "body", m.Pings); err != nil {
		return err
	}

	for i := 0; i < len(m.Pings); i++ {
		if swag.IsZero(m.Pings[i]) { // not required
			continue
		}

		if m.Pings[i] != nil {
			if err := m.Pings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ping response based on the context it is used
func (m *PingResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PingResponse) contextValidatePings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Pings); i++ {

		if m.Pings[i] != nil {
			if err := m.Pings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PingResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PingResponse) UnmarshalBinary(b []byte) error {
	var res PingResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
