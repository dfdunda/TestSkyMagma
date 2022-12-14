// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// NhRoutes neutral host PLMN ID to serving FeG network routes
//
// swagger:model nh_routes
type NhRoutes map[string]string

// Validate validates this nh routes
func (m NhRoutes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this nh routes based on context it is used
func (m NhRoutes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
