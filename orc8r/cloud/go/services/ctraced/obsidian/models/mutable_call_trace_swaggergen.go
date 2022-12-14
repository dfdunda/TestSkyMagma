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

// MutableCallTrace Subset of call trace fields which are mutable
//
// swagger:model mutable_call_trace
type MutableCallTrace struct {

	// True if requesting call trace to end
	// Required: true
	RequestedEnd *bool `json:"requested_end"`
}

// Validate validates this mutable call trace
func (m *MutableCallTrace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequestedEnd(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MutableCallTrace) validateRequestedEnd(formats strfmt.Registry) error {

	if err := validate.Required("requested_end", "body", m.RequestedEnd); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mutable call trace based on context it is used
func (m *MutableCallTrace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MutableCallTrace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MutableCallTrace) UnmarshalBinary(b []byte) error {
	var res MutableCallTrace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
