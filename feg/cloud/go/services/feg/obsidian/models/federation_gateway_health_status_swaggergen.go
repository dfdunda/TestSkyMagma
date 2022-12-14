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

// FederationGatewayHealthStatus Health status of a Federation Gateway
//
// swagger:model federation_gateway_health_status
type FederationGatewayHealthStatus struct {

	// description
	// Required: true
	Description string `json:"description"`

	// service status
	ServiceStatus map[string]ServiceStatusHealth `json:"service_status,omitempty"`

	// status
	// Required: true
	// Enum: [HEALTHY UNHEALTHY]
	Status string `json:"status"`
}

// Validate validates this federation gateway health status
func (m *FederationGatewayHealthStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FederationGatewayHealthStatus) validateDescription(formats strfmt.Registry) error {

	if err := validate.RequiredString("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *FederationGatewayHealthStatus) validateServiceStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceStatus) { // not required
		return nil
	}

	for k := range m.ServiceStatus {

		if err := validate.Required("service_status"+"."+k, "body", m.ServiceStatus[k]); err != nil {
			return err
		}
		if val, ok := m.ServiceStatus[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service_status" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("service_status" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

var federationGatewayHealthStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HEALTHY","UNHEALTHY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		federationGatewayHealthStatusTypeStatusPropEnum = append(federationGatewayHealthStatusTypeStatusPropEnum, v)
	}
}

const (

	// FederationGatewayHealthStatusStatusHEALTHY captures enum value "HEALTHY"
	FederationGatewayHealthStatusStatusHEALTHY string = "HEALTHY"

	// FederationGatewayHealthStatusStatusUNHEALTHY captures enum value "UNHEALTHY"
	FederationGatewayHealthStatusStatusUNHEALTHY string = "UNHEALTHY"
)

// prop value enum
func (m *FederationGatewayHealthStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, federationGatewayHealthStatusTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FederationGatewayHealthStatus) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this federation gateway health status based on the context it is used
func (m *FederationGatewayHealthStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FederationGatewayHealthStatus) contextValidateServiceStatus(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.ServiceStatus {

		if val, ok := m.ServiceStatus[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FederationGatewayHealthStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FederationGatewayHealthStatus) UnmarshalBinary(b []byte) error {
	var res FederationGatewayHealthStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
