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

// CapacityPool Information on a capacity pool license and how it is associated with the cluster.
//
// swagger:model capacity_pool
type CapacityPool struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// Nodes in the cluster associated with this capacity pool.
	CapacityPoolInlineNodes []*CapacityPoolInlineNodesInlineArrayItem `json:"nodes,omitempty"`

	// license manager
	LicenseManager *CapacityPoolInlineLicenseManager `json:"license_manager,omitempty"`

	// Serial number of the capacity pool license.
	// Example: 390000100
	SerialNumber *string `json:"serial_number,omitempty"`
}

// Validate validates this capacity pool
func (m *CapacityPool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapacityPoolInlineNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicenseManager(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityPool) validateLinks(formats strfmt.Registry) error {
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

func (m *CapacityPool) validateCapacityPoolInlineNodes(formats strfmt.Registry) error {
	if swag.IsZero(m.CapacityPoolInlineNodes) { // not required
		return nil
	}

	for i := 0; i < len(m.CapacityPoolInlineNodes); i++ {
		if swag.IsZero(m.CapacityPoolInlineNodes[i]) { // not required
			continue
		}

		if m.CapacityPoolInlineNodes[i] != nil {
			if err := m.CapacityPoolInlineNodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CapacityPool) validateLicenseManager(formats strfmt.Registry) error {
	if swag.IsZero(m.LicenseManager) { // not required
		return nil
	}

	if m.LicenseManager != nil {
		if err := m.LicenseManager.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("license_manager")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("license_manager")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this capacity pool based on the context it is used
func (m *CapacityPool) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCapacityPoolInlineNodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLicenseManager(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityPool) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CapacityPool) contextValidateCapacityPoolInlineNodes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CapacityPoolInlineNodes); i++ {

		if m.CapacityPoolInlineNodes[i] != nil {

			if swag.IsZero(m.CapacityPoolInlineNodes[i]) { // not required
				return nil
			}

			if err := m.CapacityPoolInlineNodes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CapacityPool) contextValidateLicenseManager(ctx context.Context, formats strfmt.Registry) error {

	if m.LicenseManager != nil {

		if swag.IsZero(m.LicenseManager) { // not required
			return nil
		}

		if err := m.LicenseManager.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("license_manager")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("license_manager")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CapacityPool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapacityPool) UnmarshalBinary(b []byte) error {
	var res CapacityPool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CapacityPoolInlineLicenseManager License manager instance where this capacity pool license in installed.
//
// swagger:model capacity_pool_inline_license_manager
type CapacityPoolInlineLicenseManager struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// uuid
	// Example: 4ea7a442-86d1-11e0-ae1c-112233445566
	// Format: uuid
	UUID *strfmt.UUID `json:"uuid,omitempty"`
}

// Validate validates this capacity pool inline license manager
func (m *CapacityPoolInlineLicenseManager) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityPoolInlineLicenseManager) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("license_manager" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("license_manager" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPoolInlineLicenseManager) validateUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("license_manager"+"."+"uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this capacity pool inline license manager based on the context it is used
func (m *CapacityPoolInlineLicenseManager) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityPoolInlineLicenseManager) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("license_manager" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("license_manager" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CapacityPoolInlineLicenseManager) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapacityPoolInlineLicenseManager) UnmarshalBinary(b []byte) error {
	var res CapacityPoolInlineLicenseManager
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CapacityPoolInlineNodesInlineArrayItem Information on a node from the capacity licensing perspective.
//
// swagger:model capacity_pool_inline_nodes_inline_array_item
type CapacityPoolInlineNodesInlineArrayItem struct {

	// node
	Node *NodeReference `json:"node,omitempty"`

	// Capacity, in bytes, that is currently used by the node.
	// Read Only: true
	UsedSize *int64 `json:"used_size,omitempty"`
}

// Validate validates this capacity pool inline nodes inline array item
func (m *CapacityPoolInlineNodesInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityPoolInlineNodesInlineArrayItem) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this capacity pool inline nodes inline array item based on the context it is used
func (m *CapacityPoolInlineNodesInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsedSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityPoolInlineNodesInlineArrayItem) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {

		if swag.IsZero(m.Node) { // not required
			return nil
		}

		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPoolInlineNodesInlineArrayItem) contextValidateUsedSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "used_size", "body", m.UsedSize); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CapacityPoolInlineNodesInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapacityPoolInlineNodesInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res CapacityPoolInlineNodesInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
