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

// GroupPolicyObjectSecuritySetting group policy object security setting
//
// swagger:model group_policy_object_security_setting
type GroupPolicyObjectSecuritySetting struct {

	// event audit settings
	EventAuditSettings *GroupPolicyObjectEventAudit `json:"event_audit_settings,omitempty"`

	// event log settings
	EventLogSettings *GroupPolicyObjectEventLog `json:"event_log_settings,omitempty"`

	// Files/Directories for file security.
	// Example: ["/vol1/home","/vol1/dir1"]
	GroupPolicyObjectSecuritySettingInlineFilesOrFolders []*string `json:"files_or_folders,omitempty"`

	// kerberos
	Kerberos *GroupPolicyObjectKerberos `json:"kerberos,omitempty"`

	// privilege rights
	PrivilegeRights *GroupPolicyObjectPrivilegeRight `json:"privilege_rights,omitempty"`

	// registry values
	RegistryValues *GroupPolicyObjectRegistryValue `json:"registry_values,omitempty"`

	// restrict anonymous
	RestrictAnonymous *GroupPolicyObjectRestrictAnonymous `json:"restrict_anonymous,omitempty"`

	// restricted groups
	RestrictedGroups GroupPolicyObjectRestrictedGroupsArrayInline `json:"restricted_groups,omitempty"`
}

// Validate validates this group policy object security setting
func (m *GroupPolicyObjectSecuritySetting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventAuditSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventLogSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKerberos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivilegeRights(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistryValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestrictAnonymous(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestrictedGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectSecuritySetting) validateEventAuditSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.EventAuditSettings) { // not required
		return nil
	}

	if m.EventAuditSettings != nil {
		if err := m.EventAuditSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_audit_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_audit_settings")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObjectSecuritySetting) validateEventLogSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.EventLogSettings) { // not required
		return nil
	}

	if m.EventLogSettings != nil {
		if err := m.EventLogSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_log_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_log_settings")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObjectSecuritySetting) validateKerberos(formats strfmt.Registry) error {
	if swag.IsZero(m.Kerberos) { // not required
		return nil
	}

	if m.Kerberos != nil {
		if err := m.Kerberos.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kerberos")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kerberos")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObjectSecuritySetting) validatePrivilegeRights(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivilegeRights) { // not required
		return nil
	}

	if m.PrivilegeRights != nil {
		if err := m.PrivilegeRights.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privilege_rights")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("privilege_rights")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObjectSecuritySetting) validateRegistryValues(formats strfmt.Registry) error {
	if swag.IsZero(m.RegistryValues) { // not required
		return nil
	}

	if m.RegistryValues != nil {
		if err := m.RegistryValues.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry_values")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry_values")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObjectSecuritySetting) validateRestrictAnonymous(formats strfmt.Registry) error {
	if swag.IsZero(m.RestrictAnonymous) { // not required
		return nil
	}

	if m.RestrictAnonymous != nil {
		if err := m.RestrictAnonymous.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restrict_anonymous")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restrict_anonymous")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObjectSecuritySetting) validateRestrictedGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.RestrictedGroups) { // not required
		return nil
	}

	if err := m.RestrictedGroups.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("restricted_groups")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("restricted_groups")
		}
		return err
	}

	return nil
}

// ContextValidate validate this group policy object security setting based on the context it is used
func (m *GroupPolicyObjectSecuritySetting) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventAuditSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventLogSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKerberos(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrivilegeRights(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegistryValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRestrictAnonymous(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRestrictedGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectSecuritySetting) contextValidateEventAuditSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.EventAuditSettings != nil {

		if swag.IsZero(m.EventAuditSettings) { // not required
			return nil
		}

		if err := m.EventAuditSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_audit_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_audit_settings")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObjectSecuritySetting) contextValidateEventLogSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.EventLogSettings != nil {

		if swag.IsZero(m.EventLogSettings) { // not required
			return nil
		}

		if err := m.EventLogSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_log_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_log_settings")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObjectSecuritySetting) contextValidateKerberos(ctx context.Context, formats strfmt.Registry) error {

	if m.Kerberos != nil {

		if swag.IsZero(m.Kerberos) { // not required
			return nil
		}

		if err := m.Kerberos.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kerberos")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kerberos")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObjectSecuritySetting) contextValidatePrivilegeRights(ctx context.Context, formats strfmt.Registry) error {

	if m.PrivilegeRights != nil {

		if swag.IsZero(m.PrivilegeRights) { // not required
			return nil
		}

		if err := m.PrivilegeRights.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privilege_rights")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("privilege_rights")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObjectSecuritySetting) contextValidateRegistryValues(ctx context.Context, formats strfmt.Registry) error {

	if m.RegistryValues != nil {

		if swag.IsZero(m.RegistryValues) { // not required
			return nil
		}

		if err := m.RegistryValues.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry_values")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry_values")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObjectSecuritySetting) contextValidateRestrictAnonymous(ctx context.Context, formats strfmt.Registry) error {

	if m.RestrictAnonymous != nil {

		if swag.IsZero(m.RestrictAnonymous) { // not required
			return nil
		}

		if err := m.RestrictAnonymous.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restrict_anonymous")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restrict_anonymous")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObjectSecuritySetting) contextValidateRestrictedGroups(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RestrictedGroups.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("restricted_groups")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("restricted_groups")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupPolicyObjectSecuritySetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupPolicyObjectSecuritySetting) UnmarshalBinary(b []byte) error {
	var res GroupPolicyObjectSecuritySetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
