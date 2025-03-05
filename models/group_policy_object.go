// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GroupPolicyObject group policy object
//
// swagger:model group_policy_object
type GroupPolicyObject struct {

	// Will not be populated for objects that are yet to be applied.
	CentralAccessPolicySettings GroupPolicyObjectCentralAccessPolicySettingsArrayInline `json:"central_access_policy_settings,omitempty"`

	// Types of events to be audited.
	// Example: none
	// Enum: ["none","success","failure","both"]
	CentralAccessPolicyStagingAuditType *string `json:"central_access_policy_staging_audit_type,omitempty"`

	// Specifies whether group policies are enabled for the SVM.
	Enabled *bool `json:"enabled,omitempty"`

	// List of extensions.
	// Example: ["audit","security"]
	Extensions []*string `json:"extensions,omitempty"`

	// File system path.
	// Example: \\test.com\\SysVol\\test.com\\policies\\{42474212-3f9d-4489-ae01-6fcf4f805d4c}
	FileSystemPath *string `json:"file_system_path,omitempty"`

	// Group policy object index.
	// Example: 1
	Index *int64 `json:"index,omitempty"`

	// LDAP path to the GPO.
	// Example: cn={42474212-3f9d-4489-ae01-6fcf4f805d4c},cn=policies,cn=system,DC=TEST,DC=COM
	LdapPath *string `json:"ldap_path,omitempty"`

	// Link info.
	// Example: domain
	// Enum: ["local","site","domain","organizational_unit","rsop"]
	Link *string `json:"link,omitempty"`

	// name
	// Example: test_policy
	// Min Length: 1
	Name *string `json:"name,omitempty"`

	// registry settings
	RegistrySettings *GroupPolicyObjectRegistrySetting `json:"registry_settings,omitempty"`

	// security settings
	SecuritySettings *GroupPolicyObjectSecuritySetting `json:"security_settings,omitempty"`

	// svm
	Svm *GroupPolicyObjectInlineSvm `json:"svm,omitempty"`

	// Policy UUID.
	// Example: 42474212-3f9d-4489-ae01-6fcf4f805d4c
	UUID *string `json:"uuid,omitempty"`

	// Group policy object version.
	// Example: 7
	Version *int64 `json:"version,omitempty"`
}

// Validate validates this group policy object
func (m *GroupPolicyObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCentralAccessPolicySettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCentralAccessPolicyStagingAuditType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistrySettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecuritySettings(formats); err != nil {
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

func (m *GroupPolicyObject) validateCentralAccessPolicySettings(formats strfmt.Registry) error {
	if swag.IsZero(m.CentralAccessPolicySettings) { // not required
		return nil
	}

	if err := m.CentralAccessPolicySettings.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("central_access_policy_settings")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("central_access_policy_settings")
		}
		return err
	}

	return nil
}

var groupPolicyObjectTypeCentralAccessPolicyStagingAuditTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","success","failure","both"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupPolicyObjectTypeCentralAccessPolicyStagingAuditTypePropEnum = append(groupPolicyObjectTypeCentralAccessPolicyStagingAuditTypePropEnum, v)
	}
}

const (

	// GroupPolicyObjectCentralAccessPolicyStagingAuditTypeNone captures enum value "none"
	GroupPolicyObjectCentralAccessPolicyStagingAuditTypeNone string = "none"

	// GroupPolicyObjectCentralAccessPolicyStagingAuditTypeSuccess captures enum value "success"
	GroupPolicyObjectCentralAccessPolicyStagingAuditTypeSuccess string = "success"

	// GroupPolicyObjectCentralAccessPolicyStagingAuditTypeFailure captures enum value "failure"
	GroupPolicyObjectCentralAccessPolicyStagingAuditTypeFailure string = "failure"

	// GroupPolicyObjectCentralAccessPolicyStagingAuditTypeBoth captures enum value "both"
	GroupPolicyObjectCentralAccessPolicyStagingAuditTypeBoth string = "both"
)

// prop value enum
func (m *GroupPolicyObject) validateCentralAccessPolicyStagingAuditTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, groupPolicyObjectTypeCentralAccessPolicyStagingAuditTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GroupPolicyObject) validateCentralAccessPolicyStagingAuditType(formats strfmt.Registry) error {
	if swag.IsZero(m.CentralAccessPolicyStagingAuditType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCentralAccessPolicyStagingAuditTypeEnum("central_access_policy_staging_audit_type", "body", *m.CentralAccessPolicyStagingAuditType); err != nil {
		return err
	}

	return nil
}

var groupPolicyObjectExtensionsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["registry","disk_quota","scripts","security","efs_recovery","ip_security","unsupported","cap","audit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupPolicyObjectExtensionsItemsEnum = append(groupPolicyObjectExtensionsItemsEnum, v)
	}
}

func (m *GroupPolicyObject) validateExtensionsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, groupPolicyObjectExtensionsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GroupPolicyObject) validateExtensions(formats strfmt.Registry) error {
	if swag.IsZero(m.Extensions) { // not required
		return nil
	}

	for i := 0; i < len(m.Extensions); i++ {
		if swag.IsZero(m.Extensions[i]) { // not required
			continue
		}

		// value enum
		if err := m.validateExtensionsItemsEnum("extensions"+"."+strconv.Itoa(i), "body", *m.Extensions[i]); err != nil {
			return err
		}

	}

	return nil
}

var groupPolicyObjectTypeLinkPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["local","site","domain","organizational_unit","rsop"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupPolicyObjectTypeLinkPropEnum = append(groupPolicyObjectTypeLinkPropEnum, v)
	}
}

const (

	// GroupPolicyObjectLinkLocal captures enum value "local"
	GroupPolicyObjectLinkLocal string = "local"

	// GroupPolicyObjectLinkSite captures enum value "site"
	GroupPolicyObjectLinkSite string = "site"

	// GroupPolicyObjectLinkDomain captures enum value "domain"
	GroupPolicyObjectLinkDomain string = "domain"

	// GroupPolicyObjectLinkOrganizationalUnit captures enum value "organizational_unit"
	GroupPolicyObjectLinkOrganizationalUnit string = "organizational_unit"

	// GroupPolicyObjectLinkRsop captures enum value "rsop"
	GroupPolicyObjectLinkRsop string = "rsop"
)

// prop value enum
func (m *GroupPolicyObject) validateLinkEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, groupPolicyObjectTypeLinkPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GroupPolicyObject) validateLink(formats strfmt.Registry) error {
	if swag.IsZero(m.Link) { // not required
		return nil
	}

	// value enum
	if err := m.validateLinkEnum("link", "body", *m.Link); err != nil {
		return err
	}

	return nil
}

func (m *GroupPolicyObject) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	return nil
}

func (m *GroupPolicyObject) validateRegistrySettings(formats strfmt.Registry) error {
	if swag.IsZero(m.RegistrySettings) { // not required
		return nil
	}

	if m.RegistrySettings != nil {
		if err := m.RegistrySettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry_settings")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObject) validateSecuritySettings(formats strfmt.Registry) error {
	if swag.IsZero(m.SecuritySettings) { // not required
		return nil
	}

	if m.SecuritySettings != nil {
		if err := m.SecuritySettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_settings")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObject) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this group policy object based on the context it is used
func (m *GroupPolicyObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCentralAccessPolicySettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegistrySettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecuritySettings(ctx, formats); err != nil {
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

func (m *GroupPolicyObject) contextValidateCentralAccessPolicySettings(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CentralAccessPolicySettings.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("central_access_policy_settings")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("central_access_policy_settings")
		}
		return err
	}

	return nil
}

func (m *GroupPolicyObject) contextValidateRegistrySettings(ctx context.Context, formats strfmt.Registry) error {

	if m.RegistrySettings != nil {

		if swag.IsZero(m.RegistrySettings) { // not required
			return nil
		}

		if err := m.RegistrySettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry_settings")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObject) contextValidateSecuritySettings(ctx context.Context, formats strfmt.Registry) error {

	if m.SecuritySettings != nil {

		if swag.IsZero(m.SecuritySettings) { // not required
			return nil
		}

		if err := m.SecuritySettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_settings")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObject) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GroupPolicyObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupPolicyObject) UnmarshalBinary(b []byte) error {
	var res GroupPolicyObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GroupPolicyObjectInlineSvm Will not be populated for objects that are yet to be applied.
//
// swagger:model group_policy_object_inline_svm
type GroupPolicyObjectInlineSvm struct {

	// links
	Links *GroupPolicyObjectInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this group policy object inline svm
func (m *GroupPolicyObjectInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this group policy object inline svm based on the context it is used
func (m *GroupPolicyObjectInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GroupPolicyObjectInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupPolicyObjectInlineSvm) UnmarshalBinary(b []byte) error {
	var res GroupPolicyObjectInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GroupPolicyObjectInlineSvmInlineLinks group policy object inline svm inline links
//
// swagger:model group_policy_object_inline_svm_inline__links
type GroupPolicyObjectInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this group policy object inline svm inline links
func (m *GroupPolicyObjectInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this group policy object inline svm inline links based on the context it is used
func (m *GroupPolicyObjectInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GroupPolicyObjectInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupPolicyObjectInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res GroupPolicyObjectInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
