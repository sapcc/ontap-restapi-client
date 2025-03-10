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

// S3Group This is a container for S3 user groups.
//
// swagger:model s3_group
type S3Group struct {

	// Can contain any additional information about the group being created or modified.
	// Example: Admin group
	// Max Length: 256
	// Min Length: 0
	Comment *string `json:"comment,omitempty"`

	// Specifies a unique Group ID used to identify a particular group. This parameter should not be specified in the POST or PATCH method. A Group ID is automatically generated and it is retrieved using the GET method. Group ID is SVM-scoped.
	// Example: 5
	// Read Only: true
	ID *int64 `json:"id,omitempty"`

	// Specifies the name of the group. A group name length can range from 1 to 128 characters and can only contain the following combination of characters 0-9, A-Z, a-z, "_", "+", "=", ",", ".","@", and "-".
	// Example: Admin-Group
	// Max Length: 128
	// Min Length: 1
	// Pattern: ^[0-9A-Za-z_+=,.@-]{1,128}$
	Name *string `json:"name,omitempty"`

	// Specifies a list of policies that are attached to the group. The wildcard character "*" is a valid value for specifying all policies.
	S3GroupInlinePolicies []*S3GroupInlinePoliciesInlineArrayItem `json:"policies,omitempty"`

	// Specifies the list of users who belong to the group.
	S3GroupInlineUsers []*S3GroupInlineUsersInlineArrayItem `json:"users,omitempty"`

	// svm
	Svm *S3GroupInlineSvm `json:"svm,omitempty"`
}

// Validate validates this s3 group
func (m *S3Group) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3GroupInlinePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3GroupInlineUsers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3Group) validateComment(formats strfmt.Registry) error {
	if swag.IsZero(m.Comment) { // not required
		return nil
	}

	if err := validate.MinLength("comment", "body", *m.Comment, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("comment", "body", *m.Comment, 256); err != nil {
		return err
	}

	return nil
}

func (m *S3Group) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 128); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `^[0-9A-Za-z_+=,.@-]{1,128}$`); err != nil {
		return err
	}

	return nil
}

func (m *S3Group) validateS3GroupInlinePolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.S3GroupInlinePolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.S3GroupInlinePolicies); i++ {
		if swag.IsZero(m.S3GroupInlinePolicies[i]) { // not required
			continue
		}

		if m.S3GroupInlinePolicies[i] != nil {
			if err := m.S3GroupInlinePolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *S3Group) validateS3GroupInlineUsers(formats strfmt.Registry) error {
	if swag.IsZero(m.S3GroupInlineUsers) { // not required
		return nil
	}

	for i := 0; i < len(m.S3GroupInlineUsers); i++ {
		if swag.IsZero(m.S3GroupInlineUsers[i]) { // not required
			continue
		}

		if m.S3GroupInlineUsers[i] != nil {
			if err := m.S3GroupInlineUsers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *S3Group) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 group based on the context it is used
func (m *S3Group) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateS3GroupInlinePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateS3GroupInlineUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3Group) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *S3Group) contextValidateS3GroupInlinePolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.S3GroupInlinePolicies); i++ {

		if m.S3GroupInlinePolicies[i] != nil {

			if swag.IsZero(m.S3GroupInlinePolicies[i]) { // not required
				return nil
			}

			if err := m.S3GroupInlinePolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *S3Group) contextValidateS3GroupInlineUsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.S3GroupInlineUsers); i++ {

		if m.S3GroupInlineUsers[i] != nil {

			if swag.IsZero(m.S3GroupInlineUsers[i]) { // not required
				return nil
			}

			if err := m.S3GroupInlineUsers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *S3Group) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {

		if swag.IsZero(m.Svm) { // not required
			return nil
		}

		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3Group) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3Group) UnmarshalBinary(b []byte) error {
	var res S3Group
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3GroupInlinePoliciesInlineArrayItem s3 group inline policies inline array item
//
// swagger:model s3_group_inline_policies_inline_array_item
type S3GroupInlinePoliciesInlineArrayItem struct {

	// links
	Links *S3GroupInlinePoliciesInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// Specifies the name of the policy. A policy name length can range from 1 to 128 characters and can only contain the following combination of characters 0-9, A-Z, a-z, "_", "+", "=", ",", ".","@", and "-".
	// Example: Policy1
	// Max Length: 128
	// Min Length: 1
	// Pattern: ^[0-9A-Za-z_+=,.@-]{1,128}$
	Name *string `json:"name,omitempty"`
}

// Validate validates this s3 group inline policies inline array item
func (m *S3GroupInlinePoliciesInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3GroupInlinePoliciesInlineArrayItem) validateLinks(formats strfmt.Registry) error {
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

func (m *S3GroupInlinePoliciesInlineArrayItem) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 128); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `^[0-9A-Za-z_+=,.@-]{1,128}$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this s3 group inline policies inline array item based on the context it is used
func (m *S3GroupInlinePoliciesInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3GroupInlinePoliciesInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *S3GroupInlinePoliciesInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3GroupInlinePoliciesInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res S3GroupInlinePoliciesInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3GroupInlinePoliciesInlineArrayItemInlineLinks s3 group inline policies inline array item inline links
//
// swagger:model s3_group_inline_policies_inline_array_item_inline__links
type S3GroupInlinePoliciesInlineArrayItemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this s3 group inline policies inline array item inline links
func (m *S3GroupInlinePoliciesInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3GroupInlinePoliciesInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 group inline policies inline array item inline links based on the context it is used
func (m *S3GroupInlinePoliciesInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3GroupInlinePoliciesInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3GroupInlinePoliciesInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3GroupInlinePoliciesInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res S3GroupInlinePoliciesInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3GroupInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model s3_group_inline_svm
type S3GroupInlineSvm struct {

	// links
	Links *S3GroupInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this s3 group inline svm
func (m *S3GroupInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3GroupInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 group inline svm based on the context it is used
func (m *S3GroupInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3GroupInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3GroupInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3GroupInlineSvm) UnmarshalBinary(b []byte) error {
	var res S3GroupInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3GroupInlineSvmInlineLinks s3 group inline svm inline links
//
// swagger:model s3_group_inline_svm_inline__links
type S3GroupInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this s3 group inline svm inline links
func (m *S3GroupInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3GroupInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 group inline svm inline links based on the context it is used
func (m *S3GroupInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3GroupInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3GroupInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3GroupInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res S3GroupInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3GroupInlineUsersInlineArrayItem s3 group inline users inline array item
//
// swagger:model s3_group_inline_users_inline_array_item
type S3GroupInlineUsersInlineArrayItem struct {

	// links
	Links *S3GroupInlineUsersInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// Specifies the name of the user. A user name length can range from 1 to 64 characters and can only contain the following combination of characters 0-9, A-Z, a-z, "_", "+", "=", ",", ".","@", and "-".
	// Example: user-1
	// Max Length: 64
	// Min Length: 1
	// Pattern: ^[0-9A-Za-z_+=,.@-]{1,64}$
	Name *string `json:"name,omitempty"`
}

// Validate validates this s3 group inline users inline array item
func (m *S3GroupInlineUsersInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3GroupInlineUsersInlineArrayItem) validateLinks(formats strfmt.Registry) error {
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

func (m *S3GroupInlineUsersInlineArrayItem) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `^[0-9A-Za-z_+=,.@-]{1,64}$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this s3 group inline users inline array item based on the context it is used
func (m *S3GroupInlineUsersInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3GroupInlineUsersInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *S3GroupInlineUsersInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3GroupInlineUsersInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res S3GroupInlineUsersInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3GroupInlineUsersInlineArrayItemInlineLinks s3 group inline users inline array item inline links
//
// swagger:model s3_group_inline_users_inline_array_item_inline__links
type S3GroupInlineUsersInlineArrayItemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this s3 group inline users inline array item inline links
func (m *S3GroupInlineUsersInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3GroupInlineUsersInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 group inline users inline array item inline links based on the context it is used
func (m *S3GroupInlineUsersInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3GroupInlineUsersInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3GroupInlineUsersInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3GroupInlineUsersInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res S3GroupInlineUsersInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
