// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubscriptionProfile HSS Subscription Profile
//
// swagger:model subscription_profile
type SubscriptionProfile struct {

	// max dl bit rate
	// Example: 200000000
	MaxDlBitRate uint64 `json:"max_dl_bit_rate,omitempty"`

	// max ul bit rate
	// Example: 100000000
	MaxUlBitRate uint64 `json:"max_ul_bit_rate,omitempty"`
}

// Validate validates this subscription profile
func (m *SubscriptionProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this subscription profile based on context it is used
func (m *SubscriptionProfile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionProfile) UnmarshalBinary(b []byte) error {
	var res SubscriptionProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
