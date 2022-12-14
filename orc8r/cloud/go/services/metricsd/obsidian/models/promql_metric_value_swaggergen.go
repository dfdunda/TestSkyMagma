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

// PromqlMetricValue If resultType is 'vector' value is in the result, if 'matrix', values is in the result
//
// swagger:model promql_metric_value
type PromqlMetricValue struct {

	// metric
	// Required: true
	Metric *PromqlMetric `json:"metric"`

	// value
	Value MetricDatapoint `json:"value,omitempty"`

	// values
	Values MetricDatapoints `json:"values,omitempty"`
}

// Validate validates this promql metric value
func (m *PromqlMetricValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PromqlMetricValue) validateMetric(formats strfmt.Registry) error {

	if err := validate.Required("metric", "body", m.Metric); err != nil {
		return err
	}

	if m.Metric != nil {
		if err := m.Metric.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metric")
			}
			return err
		}
	}

	return nil
}

func (m *PromqlMetricValue) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if err := m.Value.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("value")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("value")
		}
		return err
	}

	return nil
}

func (m *PromqlMetricValue) validateValues(formats strfmt.Registry) error {
	if swag.IsZero(m.Values) { // not required
		return nil
	}

	if err := m.Values.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("values")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("values")
		}
		return err
	}

	return nil
}

// ContextValidate validate this promql metric value based on the context it is used
func (m *PromqlMetricValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PromqlMetricValue) contextValidateMetric(ctx context.Context, formats strfmt.Registry) error {

	if m.Metric != nil {
		if err := m.Metric.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metric")
			}
			return err
		}
	}

	return nil
}

func (m *PromqlMetricValue) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Value.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("value")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("value")
		}
		return err
	}

	return nil
}

func (m *PromqlMetricValue) contextValidateValues(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Values.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("values")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("values")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PromqlMetricValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PromqlMetricValue) UnmarshalBinary(b []byte) error {
	var res PromqlMetricValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
