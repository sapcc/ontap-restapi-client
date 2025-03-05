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

// ConsistencyGroupQosReference consistency group qos reference
//
// swagger:model consistency_group_qos_reference
type ConsistencyGroupQosReference struct {

	// policy
	Policy *ConsistencyGroupQosReferenceInlinePolicy `json:"policy,omitempty"`
}

// Validate validates this consistency group qos reference
func (m *ConsistencyGroupQosReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupQosReference) validatePolicy(formats strfmt.Registry) error {
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

// ContextValidate validate this consistency group qos reference based on the context it is used
func (m *ConsistencyGroupQosReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupQosReference) contextValidatePolicy(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ConsistencyGroupQosReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupQosReference) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupQosReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsistencyGroupQosReferenceInlinePolicy The QoS policy
//
// swagger:model consistency_group_qos_reference_inline_policy
type ConsistencyGroupQosReferenceInlinePolicy struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// Specifies the maximum throughput in IOPS, 0 means none. This is mutually exclusive with name and UUID during POST and PATCH.
	// Example: 10000
	// Read Only: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	MaxThroughputIops *int64 `json:"max_throughput_iops,omitempty"`

	// Specifies the maximum throughput in Megabytes per sec, 0 means none. This is mutually exclusive with name and UUID during POST and PATCH.
	// Example: 500
	// Read Only: true
	// Maximum: 4.194303e+06
	// Minimum: 0
	MaxThroughputMbps *int64 `json:"max_throughput_mbps,omitempty"`

	// Specifies the minimum throughput in IOPS, 0 means none. Setting "min_throughput" is supported on AFF platforms only, unless FabricPool tiering policies are set. This is mutually exclusive with name and UUID during POST and PATCH.
	// Example: 2000
	// Read Only: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	MinThroughputIops *int64 `json:"min_throughput_iops,omitempty"`

	// Specifies the minimum throughput in Megabytes per sec, 0 means none. This is mutually exclusive with name and UUID during POST and PATCH.
	// Example: 500
	// Read Only: true
	// Maximum: 4.194303e+06
	// Minimum: 0
	MinThroughputMbps *int64 `json:"min_throughput_mbps,omitempty"`

	// The QoS policy group name. This is mutually exclusive with UUID and other QoS attributes during POST and PATCH.
	// Example: performance
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// The QoS policy group UUID. This is mutually exclusive with name and other QoS attributes during POST and PATCH.
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this consistency group qos reference inline policy
func (m *ConsistencyGroupQosReferenceInlinePolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxThroughputIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxThroughputMbps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinThroughputIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinThroughputMbps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupQosReferenceInlinePolicy) validateLinks(formats strfmt.Registry) error {
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

func (m *ConsistencyGroupQosReferenceInlinePolicy) validateMaxThroughputIops(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxThroughputIops) { // not required
		return nil
	}

	if err := validate.MinimumInt("policy"+"."+"max_throughput_iops", "body", *m.MaxThroughputIops, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("policy"+"."+"max_throughput_iops", "body", *m.MaxThroughputIops, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupQosReferenceInlinePolicy) validateMaxThroughputMbps(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxThroughputMbps) { // not required
		return nil
	}

	if err := validate.MinimumInt("policy"+"."+"max_throughput_mbps", "body", *m.MaxThroughputMbps, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("policy"+"."+"max_throughput_mbps", "body", *m.MaxThroughputMbps, 4.194303e+06, false); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupQosReferenceInlinePolicy) validateMinThroughputIops(formats strfmt.Registry) error {
	if swag.IsZero(m.MinThroughputIops) { // not required
		return nil
	}

	if err := validate.MinimumInt("policy"+"."+"min_throughput_iops", "body", *m.MinThroughputIops, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("policy"+"."+"min_throughput_iops", "body", *m.MinThroughputIops, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupQosReferenceInlinePolicy) validateMinThroughputMbps(formats strfmt.Registry) error {
	if swag.IsZero(m.MinThroughputMbps) { // not required
		return nil
	}

	if err := validate.MinimumInt("policy"+"."+"min_throughput_mbps", "body", *m.MinThroughputMbps, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("policy"+"."+"min_throughput_mbps", "body", *m.MinThroughputMbps, 4.194303e+06, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this consistency group qos reference inline policy based on the context it is used
func (m *ConsistencyGroupQosReferenceInlinePolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxThroughputIops(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxThroughputMbps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMinThroughputIops(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMinThroughputMbps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupQosReferenceInlinePolicy) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsistencyGroupQosReferenceInlinePolicy) contextValidateMaxThroughputIops(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "policy"+"."+"max_throughput_iops", "body", m.MaxThroughputIops); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupQosReferenceInlinePolicy) contextValidateMaxThroughputMbps(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "policy"+"."+"max_throughput_mbps", "body", m.MaxThroughputMbps); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupQosReferenceInlinePolicy) contextValidateMinThroughputIops(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "policy"+"."+"min_throughput_iops", "body", m.MinThroughputIops); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupQosReferenceInlinePolicy) contextValidateMinThroughputMbps(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "policy"+"."+"min_throughput_mbps", "body", m.MinThroughputMbps); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupQosReferenceInlinePolicy) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "policy"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupQosReferenceInlinePolicy) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "policy"+"."+"uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupQosReferenceInlinePolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupQosReferenceInlinePolicy) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupQosReferenceInlinePolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
