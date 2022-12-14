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

// Event event
//
// swagger:model event
type Event struct {

	// event type
	// Example: fake_event
	// Required: true
	// Min Length: 1
	EventType string `json:"event_type"`

	// hardware id
	// Required: true
	// Min Length: 1
	HardwareID string `json:"hardware_id"`

	// stream name
	// Example: fake_stream
	// Required: true
	// Min Length: 1
	StreamName string `json:"stream_name"`

	// tag
	// Required: true
	// Min Length: 1
	Tag string `json:"tag"`

	// The timestamp in ISO 8601 format
	// Example: 2020-03-11T00:36:59.650093129+00:00
	// Required: true
	// Min Length: 1
	Timestamp string `json:"timestamp"`

	// value
	// Required: true
	Value interface{} `json:"value"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHardwareID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreamName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateEventType(formats strfmt.Registry) error {

	if err := validate.RequiredString("event_type", "body", m.EventType); err != nil {
		return err
	}

	if err := validate.MinLength("event_type", "body", m.EventType, 1); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateHardwareID(formats strfmt.Registry) error {

	if err := validate.RequiredString("hardware_id", "body", m.HardwareID); err != nil {
		return err
	}

	if err := validate.MinLength("hardware_id", "body", m.HardwareID, 1); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateStreamName(formats strfmt.Registry) error {

	if err := validate.RequiredString("stream_name", "body", m.StreamName); err != nil {
		return err
	}

	if err := validate.MinLength("stream_name", "body", m.StreamName, 1); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateTag(formats strfmt.Registry) error {

	if err := validate.RequiredString("tag", "body", m.Tag); err != nil {
		return err
	}

	if err := validate.MinLength("tag", "body", m.Tag, 1); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.RequiredString("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.MinLength("timestamp", "body", m.Timestamp, 1); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateValue(formats strfmt.Registry) error {

	if m.Value == nil {
		return errors.Required("value", "body", nil)
	}

	return nil
}

// ContextValidate validates this event based on context it is used
func (m *Event) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
