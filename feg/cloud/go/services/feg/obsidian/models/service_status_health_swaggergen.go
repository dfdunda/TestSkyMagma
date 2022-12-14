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

// ServiceStatusHealth health status of a specific
//
// swagger:model service_status_health
type ServiceStatusHealth struct {

	// health status
	// Enum: [HEALTHY UNHEALTHY]
	HealthStatus string `json:"health_status,omitempty"`

	// service state
	// Enum: [AVAILABLE UNAVAILABLE]
	ServiceState string `json:"service_state,omitempty"`
}

// Validate validates this service status health
func (m *ServiceStatusHealth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealthStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serviceStatusHealthTypeHealthStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HEALTHY","UNHEALTHY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceStatusHealthTypeHealthStatusPropEnum = append(serviceStatusHealthTypeHealthStatusPropEnum, v)
	}
}

const (

	// ServiceStatusHealthHealthStatusHEALTHY captures enum value "HEALTHY"
	ServiceStatusHealthHealthStatusHEALTHY string = "HEALTHY"

	// ServiceStatusHealthHealthStatusUNHEALTHY captures enum value "UNHEALTHY"
	ServiceStatusHealthHealthStatusUNHEALTHY string = "UNHEALTHY"
)

// prop value enum
func (m *ServiceStatusHealth) validateHealthStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serviceStatusHealthTypeHealthStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServiceStatusHealth) validateHealthStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.HealthStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateHealthStatusEnum("health_status", "body", m.HealthStatus); err != nil {
		return err
	}

	return nil
}

var serviceStatusHealthTypeServiceStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AVAILABLE","UNAVAILABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceStatusHealthTypeServiceStatePropEnum = append(serviceStatusHealthTypeServiceStatePropEnum, v)
	}
}

const (

	// ServiceStatusHealthServiceStateAVAILABLE captures enum value "AVAILABLE"
	ServiceStatusHealthServiceStateAVAILABLE string = "AVAILABLE"

	// ServiceStatusHealthServiceStateUNAVAILABLE captures enum value "UNAVAILABLE"
	ServiceStatusHealthServiceStateUNAVAILABLE string = "UNAVAILABLE"
)

// prop value enum
func (m *ServiceStatusHealth) validateServiceStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serviceStatusHealthTypeServiceStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServiceStatusHealth) validateServiceState(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceState) { // not required
		return nil
	}

	// value enum
	if err := m.validateServiceStateEnum("service_state", "body", m.ServiceState); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service status health based on context it is used
func (m *ServiceStatusHealth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceStatusHealth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceStatusHealth) UnmarshalBinary(b []byte) error {
	var res ServiceStatusHealth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
