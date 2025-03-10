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

// CifsShareACL The permissions that users and groups have on a CIFS share.
//
// swagger:model cifs_share_acl
type CifsShareACL struct {

	// links
	Links *CifsShareACLInlineLinks `json:"_links,omitempty"`

	// Specifies the access rights that a user or group has on the defined CIFS Share.
	// The following values are allowed:
	// * no_access    - User does not have CIFS share access
	// * read         - User has only read access
	// * change       - User has change access
	// * full_control - User has full_control access
	//
	// Enum: ["no_access","read","change","full_control"]
	Permission *string `json:"permission,omitempty"`

	// CIFS share name
	// Read Only: true
	Share *string `json:"share,omitempty"`

	// Specifies the user or group secure identifier (SID).
	// Example: S-1-5-21-256008430-3394229847-3930036330-1001
	// Read Only: true
	Sid *string `json:"sid,omitempty"`

	// svm
	Svm *CifsShareACLInlineSvm `json:"svm,omitempty"`

	// Specifies the type of the user or group to add to the access control
	// list of a CIFS share. The following values are allowed:
	// * windows    - Windows user or group
	// * unix_user  - UNIX user
	// * unix_group - UNIX group
	//
	// Enum: ["windows","unix_user","unix_group"]
	Type *string `json:"type,omitempty"`

	// Specifies the UNIX user or group identifier (UID/GID).
	// Example: 100
	// Read Only: true
	UnixID *int64 `json:"unix_id,omitempty"`

	// Specifies the user or group name to add to the access control list of a CIFS share.
	// Example: ENGDOMAIN\\ad_user
	UserOrGroup *string `json:"user_or_group,omitempty"`
}

// Validate validates this cifs share acl
func (m *CifsShareACL) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareACL) validateLinks(formats strfmt.Registry) error {
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

var cifsShareAclTypePermissionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["no_access","read","change","full_control"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cifsShareAclTypePermissionPropEnum = append(cifsShareAclTypePermissionPropEnum, v)
	}
}

const (

	// CifsShareACLPermissionNoAccess captures enum value "no_access"
	CifsShareACLPermissionNoAccess string = "no_access"

	// CifsShareACLPermissionRead captures enum value "read"
	CifsShareACLPermissionRead string = "read"

	// CifsShareACLPermissionChange captures enum value "change"
	CifsShareACLPermissionChange string = "change"

	// CifsShareACLPermissionFullControl captures enum value "full_control"
	CifsShareACLPermissionFullControl string = "full_control"
)

// prop value enum
func (m *CifsShareACL) validatePermissionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cifsShareAclTypePermissionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CifsShareACL) validatePermission(formats strfmt.Registry) error {
	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	// value enum
	if err := m.validatePermissionEnum("permission", "body", *m.Permission); err != nil {
		return err
	}

	return nil
}

func (m *CifsShareACL) validateSvm(formats strfmt.Registry) error {
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

var cifsShareAclTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["windows","unix_user","unix_group"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cifsShareAclTypeTypePropEnum = append(cifsShareAclTypeTypePropEnum, v)
	}
}

const (

	// CifsShareACLTypeWindows captures enum value "windows"
	CifsShareACLTypeWindows string = "windows"

	// CifsShareACLTypeUnixUser captures enum value "unix_user"
	CifsShareACLTypeUnixUser string = "unix_user"

	// CifsShareACLTypeUnixGroup captures enum value "unix_group"
	CifsShareACLTypeUnixGroup string = "unix_group"
)

// prop value enum
func (m *CifsShareACL) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cifsShareAclTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CifsShareACL) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cifs share acl based on the context it is used
func (m *CifsShareACL) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShare(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnixID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareACL) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CifsShareACL) contextValidateShare(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "share", "body", m.Share); err != nil {
		return err
	}

	return nil
}

func (m *CifsShareACL) contextValidateSid(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sid", "body", m.Sid); err != nil {
		return err
	}

	return nil
}

func (m *CifsShareACL) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CifsShareACL) contextValidateUnixID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "unix_id", "body", m.UnixID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CifsShareACL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsShareACL) UnmarshalBinary(b []byte) error {
	var res CifsShareACL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsShareACLInlineLinks cifs share acl inline links
//
// swagger:model cifs_share_acl_inline__links
type CifsShareACLInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cifs share acl inline links
func (m *CifsShareACLInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareACLInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this cifs share acl inline links based on the context it is used
func (m *CifsShareACLInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareACLInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CifsShareACLInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsShareACLInlineLinks) UnmarshalBinary(b []byte) error {
	var res CifsShareACLInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsShareACLInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model cifs_share_acl_inline_svm
type CifsShareACLInlineSvm struct {

	// links
	Links *CifsShareACLInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this cifs share acl inline svm
func (m *CifsShareACLInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareACLInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this cifs share acl inline svm based on the context it is used
func (m *CifsShareACLInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareACLInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CifsShareACLInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsShareACLInlineSvm) UnmarshalBinary(b []byte) error {
	var res CifsShareACLInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsShareACLInlineSvmInlineLinks cifs share acl inline svm inline links
//
// swagger:model cifs_share_acl_inline_svm_inline__links
type CifsShareACLInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cifs share acl inline svm inline links
func (m *CifsShareACLInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareACLInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this cifs share acl inline svm inline links based on the context it is used
func (m *CifsShareACLInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsShareACLInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CifsShareACLInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsShareACLInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res CifsShareACLInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
