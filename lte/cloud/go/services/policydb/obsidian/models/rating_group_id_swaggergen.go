// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// RatingGroupID rating group id
// Example: 1
//
// swagger:model rating_group_id
type RatingGroupID uint32

// Validate validates this rating group id
func (m RatingGroupID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rating group id based on context it is used
func (m RatingGroupID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
