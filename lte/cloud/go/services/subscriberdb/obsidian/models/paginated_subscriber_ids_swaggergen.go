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

// PaginatedSubscriberIds Page of subscriber IDs
//
// swagger:model paginated_subscriber_ids
type PaginatedSubscriberIds struct {

	// next page token
	// Required: true
	NextPageToken *PageToken `json:"next_page_token"`

	// subscribers
	// Required: true
	Subscribers []string `json:"subscribers"`

	// Estimated total number of subscriber entries
	// Example: 10
	// Required: true
	TotalCount int64 `json:"total_count"`
}

// Validate validates this paginated subscriber ids
func (m *PaginatedSubscriberIds) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNextPageToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscribers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaginatedSubscriberIds) validateNextPageToken(formats strfmt.Registry) error {

	if err := validate.Required("next_page_token", "body", m.NextPageToken); err != nil {
		return err
	}

	if err := validate.Required("next_page_token", "body", m.NextPageToken); err != nil {
		return err
	}

	if m.NextPageToken != nil {
		if err := m.NextPageToken.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next_page_token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next_page_token")
			}
			return err
		}
	}

	return nil
}

func (m *PaginatedSubscriberIds) validateSubscribers(formats strfmt.Registry) error {

	if err := validate.Required("subscribers", "body", m.Subscribers); err != nil {
		return err
	}

	return nil
}

func (m *PaginatedSubscriberIds) validateTotalCount(formats strfmt.Registry) error {

	if err := validate.Required("total_count", "body", int64(m.TotalCount)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this paginated subscriber ids based on the context it is used
func (m *PaginatedSubscriberIds) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNextPageToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaginatedSubscriberIds) contextValidateNextPageToken(ctx context.Context, formats strfmt.Registry) error {

	if m.NextPageToken != nil {
		if err := m.NextPageToken.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next_page_token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next_page_token")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaginatedSubscriberIds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginatedSubscriberIds) UnmarshalBinary(b []byte) error {
	var res PaginatedSubscriberIds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
