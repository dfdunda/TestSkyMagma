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

// DNSConfigRecord Mapping used for DNS resolving from a domain
//
// swagger:model dns_config_record
type DNSConfigRecord struct {

	// a record
	ARecord []strfmt.IPv4 `json:"a_record"`

	// aaaa record
	AaaaRecord []strfmt.IPv6 `json:"aaaa_record"`

	// cname record
	CnameRecord []string `json:"cname_record"`

	// domain
	// Example: example.com
	// Required: true
	// Min Length: 1
	Domain string `json:"domain"`
}

// Validate validates this dns config record
func (m *DNSConfigRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateARecord(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAaaaRecord(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCnameRecord(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DNSConfigRecord) validateARecord(formats strfmt.Registry) error {
	if swag.IsZero(m.ARecord) { // not required
		return nil
	}

	for i := 0; i < len(m.ARecord); i++ {

		if err := validate.MinLength("a_record"+"."+strconv.Itoa(i), "body", m.ARecord[i].String(), 1); err != nil {
			return err
		}

		if err := validate.FormatOf("a_record"+"."+strconv.Itoa(i), "body", "ipv4", m.ARecord[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *DNSConfigRecord) validateAaaaRecord(formats strfmt.Registry) error {
	if swag.IsZero(m.AaaaRecord) { // not required
		return nil
	}

	for i := 0; i < len(m.AaaaRecord); i++ {

		if err := validate.MinLength("aaaa_record"+"."+strconv.Itoa(i), "body", m.AaaaRecord[i].String(), 1); err != nil {
			return err
		}

		if err := validate.FormatOf("aaaa_record"+"."+strconv.Itoa(i), "body", "ipv6", m.AaaaRecord[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *DNSConfigRecord) validateCnameRecord(formats strfmt.Registry) error {
	if swag.IsZero(m.CnameRecord) { // not required
		return nil
	}

	for i := 0; i < len(m.CnameRecord); i++ {

		if err := validate.MinLength("cname_record"+"."+strconv.Itoa(i), "body", m.CnameRecord[i], 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *DNSConfigRecord) validateDomain(formats strfmt.Registry) error {

	if err := validate.RequiredString("domain", "body", m.Domain); err != nil {
		return err
	}

	if err := validate.MinLength("domain", "body", m.Domain, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this dns config record based on context it is used
func (m *DNSConfigRecord) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DNSConfigRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNSConfigRecord) UnmarshalBinary(b []byte) error {
	var res DNSConfigRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
