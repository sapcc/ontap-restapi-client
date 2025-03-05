// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountPassword The password object
//
// swagger:model account_password
type AccountPassword struct {

	// The user account name whose password is being modified.
	Name *string `json:"name,omitempty"`

	// owner
	Owner *AccountPasswordInlineOwner `json:"owner,omitempty"`

	// The password string
	// Max Length: 128
	// Min Length: 8
	// Format: password
	Password *strfmt.Password `json:"password,omitempty"`

	// Optional property that specifies the password hash algorithm used to generate a hash of the user's password for password matching.
	// Example: sha512
	// Enum: ["sha512","sha256","md5"]
	PasswordHashAlgorithm *string `json:"password_hash_algorithm,omitempty"`
}

// Validate validates this account password
func (m *AccountPassword) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePasswordHashAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountPassword) validateOwner(formats strfmt.Registry) error {
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

func (m *AccountPassword) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.MinLength("password", "body", m.Password.String(), 8); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", m.Password.String(), 128); err != nil {
		return err
	}

	if err := validate.FormatOf("password", "body", "password", m.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

var accountPasswordTypePasswordHashAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha512","sha256","md5"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountPasswordTypePasswordHashAlgorithmPropEnum = append(accountPasswordTypePasswordHashAlgorithmPropEnum, v)
	}
}

const (

	// AccountPasswordPasswordHashAlgorithmSha512 captures enum value "sha512"
	AccountPasswordPasswordHashAlgorithmSha512 string = "sha512"

	// AccountPasswordPasswordHashAlgorithmSha256 captures enum value "sha256"
	AccountPasswordPasswordHashAlgorithmSha256 string = "sha256"

	// AccountPasswordPasswordHashAlgorithmMd5 captures enum value "md5"
	AccountPasswordPasswordHashAlgorithmMd5 string = "md5"
)

// prop value enum
func (m *AccountPassword) validatePasswordHashAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountPasswordTypePasswordHashAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccountPassword) validatePasswordHashAlgorithm(formats strfmt.Registry) error {
	if swag.IsZero(m.PasswordHashAlgorithm) { // not required
		return nil
	}

	// value enum
	if err := m.validatePasswordHashAlgorithmEnum("password_hash_algorithm", "body", *m.PasswordHashAlgorithm); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this account password based on the context it is used
func (m *AccountPassword) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountPassword) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AccountPassword) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountPassword) UnmarshalBinary(b []byte) error {
	var res AccountPassword
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AccountPasswordInlineOwner Owner name and UUID that uniquely identifies the user account. This field is optional and valid only when a cluster administrator is executing the API to uniquely identify the account whose password is being modified. The "owner" field is not required to be specified for SVM user accounts trying to modify their password.
//
// swagger:model account_password_inline_owner
type AccountPasswordInlineOwner struct {

	// links
	Links *AccountPasswordInlineOwnerInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this account password inline owner
func (m *AccountPasswordInlineOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountPasswordInlineOwner) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this account password inline owner based on the context it is used
func (m *AccountPasswordInlineOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountPasswordInlineOwner) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *AccountPasswordInlineOwner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountPasswordInlineOwner) UnmarshalBinary(b []byte) error {
	var res AccountPasswordInlineOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AccountPasswordInlineOwnerInlineLinks account password inline owner inline links
//
// swagger:model account_password_inline_owner_inline__links
type AccountPasswordInlineOwnerInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this account password inline owner inline links
func (m *AccountPasswordInlineOwnerInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountPasswordInlineOwnerInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this account password inline owner inline links based on the context it is used
func (m *AccountPasswordInlineOwnerInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountPasswordInlineOwnerInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AccountPasswordInlineOwnerInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountPasswordInlineOwnerInlineLinks) UnmarshalBinary(b []byte) error {
	var res AccountPasswordInlineOwnerInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
