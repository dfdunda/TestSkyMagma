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

// RatingGroup rating group
//
// swagger:model rating_group
type RatingGroup struct {

	// id
	// Required: true
	ID RatingGroupID `json:"id"`

	// limit type
	// Required: true
	// Enum: [FINITE INFINITE_UNMETERED INFINITE_METERED]
	LimitType *string `json:"limit_type"`
}

// Validate validates this rating group
func (m *RatingGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimitType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RatingGroup) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", RatingGroupID(m.ID)); err != nil {
		return err
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("id")
		}
		return err
	}

	return nil
}

var ratingGroupTypeLimitTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FINITE","INFINITE_UNMETERED","INFINITE_METERED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ratingGroupTypeLimitTypePropEnum = append(ratingGroupTypeLimitTypePropEnum, v)
	}
}

const (

	// RatingGroupLimitTypeFINITE captures enum value "FINITE"
	RatingGroupLimitTypeFINITE string = "FINITE"

	// RatingGroupLimitTypeINFINITEUNMETERED captures enum value "INFINITE_UNMETERED"
	RatingGroupLimitTypeINFINITEUNMETERED string = "INFINITE_UNMETERED"

	// RatingGroupLimitTypeINFINITEMETERED captures enum value "INFINITE_METERED"
	RatingGroupLimitTypeINFINITEMETERED string = "INFINITE_METERED"
)

// prop value enum
func (m *RatingGroup) validateLimitTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ratingGroupTypeLimitTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RatingGroup) validateLimitType(formats strfmt.Registry) error {

	if err := validate.Required("limit_type", "body", m.LimitType); err != nil {
		return err
	}

	// value enum
	if err := m.validateLimitTypeEnum("limit_type", "body", *m.LimitType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this rating group based on the context it is used
func (m *RatingGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RatingGroup) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("id")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RatingGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RatingGroup) UnmarshalBinary(b []byte) error {
	var res RatingGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
