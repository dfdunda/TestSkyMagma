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
)

// MachineInfo machine info
//
// swagger:model machine_info
type MachineInfo struct {

	// cpu info
	CPUInfo *MachineInfoCPUInfo `json:"cpu_info,omitempty"`

	// network info
	NetworkInfo *MachineInfoNetworkInfo `json:"network_info,omitempty"`
}

// Validate validates this machine info
func (m *MachineInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MachineInfo) validateCPUInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.CPUInfo) { // not required
		return nil
	}

	if m.CPUInfo != nil {
		if err := m.CPUInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cpu_info")
			}
			return err
		}
	}

	return nil
}

func (m *MachineInfo) validateNetworkInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkInfo) { // not required
		return nil
	}

	if m.NetworkInfo != nil {
		if err := m.NetworkInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network_info")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this machine info based on the context it is used
func (m *MachineInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCPUInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MachineInfo) contextValidateCPUInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.CPUInfo != nil {
		if err := m.CPUInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cpu_info")
			}
			return err
		}
	}

	return nil
}

func (m *MachineInfo) contextValidateNetworkInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkInfo != nil {
		if err := m.NetworkInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MachineInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineInfo) UnmarshalBinary(b []byte) error {
	var res MachineInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MachineInfoCPUInfo machine info CPU info
//
// swagger:model MachineInfoCPUInfo
type MachineInfoCPUInfo struct {

	// architecture
	// Example: x86_64
	Architecture string `json:"architecture,omitempty"`

	// core count
	// Example: 4
	CoreCount uint64 `json:"core_count,omitempty"`

	// model name
	// Example: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
	ModelName string `json:"model_name,omitempty"`

	// threads per core
	// Example: 1
	ThreadsPerCore uint64 `json:"threads_per_core,omitempty"`
}

// Validate validates this machine info CPU info
func (m *MachineInfoCPUInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this machine info CPU info based on context it is used
func (m *MachineInfoCPUInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MachineInfoCPUInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineInfoCPUInfo) UnmarshalBinary(b []byte) error {
	var res MachineInfoCPUInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MachineInfoNetworkInfo machine info network info
//
// swagger:model MachineInfoNetworkInfo
type MachineInfoNetworkInfo struct {

	// network interfaces
	NetworkInterfaces []*NetworkInterface `json:"network_interfaces,omitempty"`

	// routing table
	RoutingTable []*Route `json:"routing_table,omitempty"`
}

// Validate validates this machine info network info
func (m *MachineInfoNetworkInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutingTable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MachineInfoNetworkInfo) validateNetworkInterfaces(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkInterfaces) { // not required
		return nil
	}

	for i := 0; i < len(m.NetworkInterfaces); i++ {
		if swag.IsZero(m.NetworkInterfaces[i]) { // not required
			continue
		}

		if m.NetworkInterfaces[i] != nil {
			if err := m.NetworkInterfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("network_info" + "." + "network_interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("network_info" + "." + "network_interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MachineInfoNetworkInfo) validateRoutingTable(formats strfmt.Registry) error {
	if swag.IsZero(m.RoutingTable) { // not required
		return nil
	}

	for i := 0; i < len(m.RoutingTable); i++ {
		if swag.IsZero(m.RoutingTable[i]) { // not required
			continue
		}

		if m.RoutingTable[i] != nil {
			if err := m.RoutingTable[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("network_info" + "." + "routing_table" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("network_info" + "." + "routing_table" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this machine info network info based on the context it is used
func (m *MachineInfoNetworkInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworkInterfaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoutingTable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MachineInfoNetworkInfo) contextValidateNetworkInterfaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NetworkInterfaces); i++ {

		if m.NetworkInterfaces[i] != nil {
			if err := m.NetworkInterfaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("network_info" + "." + "network_interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("network_info" + "." + "network_interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MachineInfoNetworkInfo) contextValidateRoutingTable(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RoutingTable); i++ {

		if m.RoutingTable[i] != nil {
			if err := m.RoutingTable[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("network_info" + "." + "routing_table" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("network_info" + "." + "routing_table" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MachineInfoNetworkInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineInfoNetworkInfo) UnmarshalBinary(b []byte) error {
	var res MachineInfoNetworkInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
