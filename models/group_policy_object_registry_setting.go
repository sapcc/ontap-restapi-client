// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GroupPolicyObjectRegistrySetting group policy object registry setting
//
// swagger:model group_policy_object_registry_setting
type GroupPolicyObjectRegistrySetting struct {

	// branchcache
	Branchcache *GroupPolicyObjectBranchcache `json:"branchcache,omitempty"`

	// Refresh time interval in ISO-8601 format.
	// Example: P15M
	RefreshTimeInterval *string `json:"refresh_time_interval,omitempty"`

	// Random offset in ISO-8601 format.
	// Example: P1D
	RefreshTimeRandomOffset *string `json:"refresh_time_random_offset,omitempty"`
}

// Validate validates this group policy object registry setting
func (m *GroupPolicyObjectRegistrySetting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBranchcache(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectRegistrySetting) validateBranchcache(formats strfmt.Registry) error {
	if swag.IsZero(m.Branchcache) { // not required
		return nil
	}

	if m.Branchcache != nil {
		if err := m.Branchcache.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("branchcache")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("branchcache")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this group policy object registry setting based on the context it is used
func (m *GroupPolicyObjectRegistrySetting) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBranchcache(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectRegistrySetting) contextValidateBranchcache(ctx context.Context, formats strfmt.Registry) error {

	if m.Branchcache != nil {

		if swag.IsZero(m.Branchcache) { // not required
			return nil
		}

		if err := m.Branchcache.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("branchcache")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("branchcache")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupPolicyObjectRegistrySetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupPolicyObjectRegistrySetting) UnmarshalBinary(b []byte) error {
	var res GroupPolicyObjectRegistrySetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
