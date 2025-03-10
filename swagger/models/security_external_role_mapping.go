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

// SecurityExternalRoleMapping security external role mapping
//
// swagger:model security_external_role_mapping
type SecurityExternalRoleMapping struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// Any comment regarding this external-role-mapping entry.
	Comment *string `json:"comment,omitempty"`

	// External Identity provider role.
	ExternalRole *string `json:"external_role,omitempty"`

	// ontap role
	OntapRole *SecurityExternalRoleMappingInlineOntapRole `json:"ontap_role,omitempty"`

	// Type of the external identity provider.
	// Enum: ["adfs","auth0","entra","keycloak","basic"]
	Provider *string `json:"provider,omitempty"`

	// Date and time indicating when this external-role-mapping entry was created.
	// Read Only: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this security external role mapping
func (m *SecurityExternalRoleMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOntapRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityExternalRoleMapping) validateLinks(formats strfmt.Registry) error {
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

func (m *SecurityExternalRoleMapping) validateOntapRole(formats strfmt.Registry) error {
	if swag.IsZero(m.OntapRole) { // not required
		return nil
	}

	if m.OntapRole != nil {
		if err := m.OntapRole.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ontap_role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ontap_role")
			}
			return err
		}
	}

	return nil
}

var securityExternalRoleMappingTypeProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["adfs","auth0","entra","keycloak","basic"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityExternalRoleMappingTypeProviderPropEnum = append(securityExternalRoleMappingTypeProviderPropEnum, v)
	}
}

const (

	// SecurityExternalRoleMappingProviderAdfs captures enum value "adfs"
	SecurityExternalRoleMappingProviderAdfs string = "adfs"

	// SecurityExternalRoleMappingProviderAuth0 captures enum value "auth0"
	SecurityExternalRoleMappingProviderAuth0 string = "auth0"

	// SecurityExternalRoleMappingProviderEntra captures enum value "entra"
	SecurityExternalRoleMappingProviderEntra string = "entra"

	// SecurityExternalRoleMappingProviderKeycloak captures enum value "keycloak"
	SecurityExternalRoleMappingProviderKeycloak string = "keycloak"

	// SecurityExternalRoleMappingProviderBasic captures enum value "basic"
	SecurityExternalRoleMappingProviderBasic string = "basic"
)

// prop value enum
func (m *SecurityExternalRoleMapping) validateProviderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, securityExternalRoleMappingTypeProviderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SecurityExternalRoleMapping) validateProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	// value enum
	if err := m.validateProviderEnum("provider", "body", *m.Provider); err != nil {
		return err
	}

	return nil
}

func (m *SecurityExternalRoleMapping) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this security external role mapping based on the context it is used
func (m *SecurityExternalRoleMapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOntapRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityExternalRoleMapping) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SecurityExternalRoleMapping) contextValidateOntapRole(ctx context.Context, formats strfmt.Registry) error {

	if m.OntapRole != nil {

		if swag.IsZero(m.OntapRole) { // not required
			return nil
		}

		if err := m.OntapRole.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ontap_role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ontap_role")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityExternalRoleMapping) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityExternalRoleMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityExternalRoleMapping) UnmarshalBinary(b []byte) error {
	var res SecurityExternalRoleMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityExternalRoleMappingInlineOntapRole ONTAP role to map with the external identity role.
//
// swagger:model security_external_role_mapping_inline_ontap_role
type SecurityExternalRoleMappingInlineOntapRole struct {

	// links
	Links *SecurityExternalRoleMappingInlineOntapRoleInlineLinks `json:"_links,omitempty"`

	// Role name
	// Example: admin
	Name *string `json:"name,omitempty"`
}

// Validate validates this security external role mapping inline ontap role
func (m *SecurityExternalRoleMappingInlineOntapRole) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityExternalRoleMappingInlineOntapRole) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ontap_role" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ontap_role" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security external role mapping inline ontap role based on the context it is used
func (m *SecurityExternalRoleMappingInlineOntapRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityExternalRoleMappingInlineOntapRole) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ontap_role" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ontap_role" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityExternalRoleMappingInlineOntapRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityExternalRoleMappingInlineOntapRole) UnmarshalBinary(b []byte) error {
	var res SecurityExternalRoleMappingInlineOntapRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityExternalRoleMappingInlineOntapRoleInlineLinks security external role mapping inline ontap role inline links
//
// swagger:model security_external_role_mapping_inline_ontap_role_inline__links
type SecurityExternalRoleMappingInlineOntapRoleInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this security external role mapping inline ontap role inline links
func (m *SecurityExternalRoleMappingInlineOntapRoleInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityExternalRoleMappingInlineOntapRoleInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ontap_role" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ontap_role" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security external role mapping inline ontap role inline links based on the context it is used
func (m *SecurityExternalRoleMappingInlineOntapRoleInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityExternalRoleMappingInlineOntapRoleInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ontap_role" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ontap_role" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityExternalRoleMappingInlineOntapRoleInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityExternalRoleMappingInlineOntapRoleInlineLinks) UnmarshalBinary(b []byte) error {
	var res SecurityExternalRoleMappingInlineOntapRoleInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
