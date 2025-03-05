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

// ConsistencyGroupExportPolicy The policy associated with volumes to export them for protocol access.
//
// swagger:model consistency_group_export_policy
type ConsistencyGroupExportPolicy struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// The set of rules that govern the export policy.
	ConsistencyGroupExportPolicyInlineRules []*ExportRules `json:"rules,omitempty"`

	// Identifier for the export policy.
	// Read Only: true
	ID *int64 `json:"id,omitempty"`

	// Name of the export policy.
	Name *string `json:"name,omitempty"`
}

// Validate validates this consistency group export policy
func (m *ConsistencyGroupExportPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsistencyGroupExportPolicyInlineRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupExportPolicy) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupExportPolicy) validateConsistencyGroupExportPolicyInlineRules(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsistencyGroupExportPolicyInlineRules) { // not required
		return nil
	}

	for i := 0; i < len(m.ConsistencyGroupExportPolicyInlineRules); i++ {
		if swag.IsZero(m.ConsistencyGroupExportPolicyInlineRules[i]) { // not required
			continue
		}

		if m.ConsistencyGroupExportPolicyInlineRules[i] != nil {
			if err := m.ConsistencyGroupExportPolicyInlineRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this consistency group export policy based on the context it is used
func (m *ConsistencyGroupExportPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConsistencyGroupExportPolicyInlineRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupExportPolicy) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupExportPolicy) contextValidateConsistencyGroupExportPolicyInlineRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConsistencyGroupExportPolicyInlineRules); i++ {

		if m.ConsistencyGroupExportPolicyInlineRules[i] != nil {

			if swag.IsZero(m.ConsistencyGroupExportPolicyInlineRules[i]) { // not required
				return nil
			}

			if err := m.ConsistencyGroupExportPolicyInlineRules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsistencyGroupExportPolicy) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupExportPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupExportPolicy) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupExportPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
