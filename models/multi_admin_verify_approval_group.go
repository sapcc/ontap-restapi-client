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

// MultiAdminVerifyApprovalGroup multi admin verify approval group
//
// swagger:model multi_admin_verify_approval_group
type MultiAdminVerifyApprovalGroup struct {

	// List of users that can approve a request.
	MultiAdminVerifyApprovalGroupInlineApprovers []*string `json:"approvers,omitempty"`

	// Email addresses that are notified when a request is created, approved, vetoed, or executed.
	MultiAdminVerifyApprovalGroupInlineEmail []*strfmt.Email `json:"email,omitempty"`

	// Name of the approval group.
	Name *string `json:"name,omitempty"`

	// owner
	Owner *MultiAdminVerifyApprovalGroupInlineOwner `json:"owner,omitempty"`
}

// Validate validates this multi admin verify approval group
func (m *MultiAdminVerifyApprovalGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMultiAdminVerifyApprovalGroupInlineEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MultiAdminVerifyApprovalGroup) validateMultiAdminVerifyApprovalGroupInlineEmail(formats strfmt.Registry) error {
	if swag.IsZero(m.MultiAdminVerifyApprovalGroupInlineEmail) { // not required
		return nil
	}

	for i := 0; i < len(m.MultiAdminVerifyApprovalGroupInlineEmail); i++ {
		if swag.IsZero(m.MultiAdminVerifyApprovalGroupInlineEmail[i]) { // not required
			continue
		}

		if err := validate.FormatOf("email"+"."+strconv.Itoa(i), "body", "email", m.MultiAdminVerifyApprovalGroupInlineEmail[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *MultiAdminVerifyApprovalGroup) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this multi admin verify approval group based on the context it is used
func (m *MultiAdminVerifyApprovalGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MultiAdminVerifyApprovalGroup) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {

		if swag.IsZero(m.Owner) { // not required
			return nil
		}

		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MultiAdminVerifyApprovalGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MultiAdminVerifyApprovalGroup) UnmarshalBinary(b []byte) error {
	var res MultiAdminVerifyApprovalGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MultiAdminVerifyApprovalGroupInlineOwner The owner of the approval group. The only valid owner is currently the cluster.
//
// swagger:model multi_admin_verify_approval_group_inline_owner
type MultiAdminVerifyApprovalGroupInlineOwner struct {

	// links
	Links *MultiAdminVerifyApprovalGroupInlineOwnerInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a POST or PATCH method.
	//
	// Example: svm1
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this multi admin verify approval group inline owner
func (m *MultiAdminVerifyApprovalGroupInlineOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MultiAdminVerifyApprovalGroupInlineOwner) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this multi admin verify approval group inline owner based on the context it is used
func (m *MultiAdminVerifyApprovalGroupInlineOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MultiAdminVerifyApprovalGroupInlineOwner) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *MultiAdminVerifyApprovalGroupInlineOwner) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "owner"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MultiAdminVerifyApprovalGroupInlineOwner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MultiAdminVerifyApprovalGroupInlineOwner) UnmarshalBinary(b []byte) error {
	var res MultiAdminVerifyApprovalGroupInlineOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MultiAdminVerifyApprovalGroupInlineOwnerInlineLinks multi admin verify approval group inline owner inline links
//
// swagger:model multi_admin_verify_approval_group_inline_owner_inline__links
type MultiAdminVerifyApprovalGroupInlineOwnerInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this multi admin verify approval group inline owner inline links
func (m *MultiAdminVerifyApprovalGroupInlineOwnerInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MultiAdminVerifyApprovalGroupInlineOwnerInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this multi admin verify approval group inline owner inline links based on the context it is used
func (m *MultiAdminVerifyApprovalGroupInlineOwnerInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MultiAdminVerifyApprovalGroupInlineOwnerInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MultiAdminVerifyApprovalGroupInlineOwnerInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MultiAdminVerifyApprovalGroupInlineOwnerInlineLinks) UnmarshalBinary(b []byte) error {
	var res MultiAdminVerifyApprovalGroupInlineOwnerInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
