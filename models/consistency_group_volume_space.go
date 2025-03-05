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

// ConsistencyGroupVolumeSpace consistency group volume space
//
// swagger:model consistency_group_volume_space
type ConsistencyGroupVolumeSpace struct {

	// The available space, in bytes.
	// Read Only: true
	Available *int64 `json:"available,omitempty"`

	// Total provisioned size, in bytes.
	Size *int64 `json:"size,omitempty"`

	// The virtual space used (includes volume reserves) before storage efficiency, in bytes.
	// Read Only: true
	Used *int64 `json:"used,omitempty"`
}

// Validate validates this consistency group volume space
func (m *ConsistencyGroupVolumeSpace) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this consistency group volume space based on the context it is used
func (m *ConsistencyGroupVolumeSpace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAvailable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupVolumeSpace) contextValidateAvailable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "available", "body", m.Available); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupVolumeSpace) contextValidateUsed(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "used", "body", m.Used); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupVolumeSpace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupVolumeSpace) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupVolumeSpace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
