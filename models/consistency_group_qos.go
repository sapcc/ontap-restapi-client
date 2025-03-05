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

// ConsistencyGroupQos consistency group qos
//
// swagger:model consistency_group_qos
type ConsistencyGroupQos struct {

	// policy
	Policy *ConsistencyGroupQosInlinePolicy `json:"policy,omitempty"`
}

// Validate validates this consistency group qos
func (m *ConsistencyGroupQos) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupQos) validatePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	if m.Policy != nil {
		if err := m.Policy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this consistency group qos based on the context it is used
func (m *ConsistencyGroupQos) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupQos) contextValidatePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.Policy != nil {

		if swag.IsZero(m.Policy) { // not required
			return nil
		}

		if err := m.Policy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupQos) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupQos) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupQos
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsistencyGroupQosInlinePolicy The QoS policy
//
// swagger:model consistency_group_qos_inline_policy
type ConsistencyGroupQosInlinePolicy struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// The QoS policy group name. This is mutually exclusive with UUID and other QoS attributes during POST and PATCH.
	// Example: performance
	Name *string `json:"name,omitempty"`

	// The QoS policy group UUID. This is mutually exclusive with name and other QoS attributes during POST and PATCH.
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this consistency group qos inline policy
func (m *ConsistencyGroupQosInlinePolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupQosInlinePolicy) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this consistency group qos inline policy based on the context it is used
func (m *ConsistencyGroupQosInlinePolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupQosInlinePolicy) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupQosInlinePolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupQosInlinePolicy) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupQosInlinePolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
