// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// ServedNetworkIds served network IDs
//
// swagger:model served_network_ids
type ServedNetworkIds []string

// Validate validates this served network ids
func (m ServedNetworkIds) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this served network ids based on context it is used
func (m ServedNetworkIds) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}