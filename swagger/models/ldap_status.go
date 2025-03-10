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

// LdapStatus ldap status
//
// swagger:model ldap_status
type LdapStatus struct {

	// This field is no longer supported. Use ipv4.code or ipv6.code instead.
	//
	// Example: 65537300
	Code *int64 `json:"code,omitempty"`

	// ipv4
	IPV4 *LdapStatusInlineIPV4 `json:"ipv4,omitempty"`

	// This field is no longer supported. Use ipv4.state instead.
	//
	// Enum: ["up","down"]
	IPV4State *string `json:"ipv4_state,omitempty"`

	// ipv6
	IPV6 *LdapStatusInlineIPV6 `json:"ipv6,omitempty"`

	// This field is no longer supported. Use ipv6.state instead.
	//
	// Enum: ["up","down"]
	IPV6State *string `json:"ipv6_state,omitempty"`

	// ldap status inline dn message
	LdapStatusInlineDnMessage []*string `json:"dn_message,omitempty"`

	// This field is no longer supported. Use ipv4.message or ipv6.message instead.
	//
	Message *string `json:"message,omitempty"`

	// The status of the LDAP service for the SVM. The LDAP service is up if either `ipv4_state` or `ipv6_state` is up.
	// The LDAP service is down if both `ipv4_state` and `ipv6_state` are down.
	//
	// Enum: ["up","down"]
	State *string `json:"state,omitempty"`
}

// Validate validates this ldap status
func (m *LdapStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPV4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV4State(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV6(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV6State(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapStatus) validateIPV4(formats strfmt.Registry) error {
	if swag.IsZero(m.IPV4) { // not required
		return nil
	}

	if m.IPV4 != nil {
		if err := m.IPV4.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv4")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipv4")
			}
			return err
		}
	}

	return nil
}

var ldapStatusTypeIPV4StatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["up","down"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapStatusTypeIPV4StatePropEnum = append(ldapStatusTypeIPV4StatePropEnum, v)
	}
}

const (

	// LdapStatusIPV4StateUp captures enum value "up"
	LdapStatusIPV4StateUp string = "up"

	// LdapStatusIPV4StateDown captures enum value "down"
	LdapStatusIPV4StateDown string = "down"
)

// prop value enum
func (m *LdapStatus) validateIPV4StateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapStatusTypeIPV4StatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapStatus) validateIPV4State(formats strfmt.Registry) error {
	if swag.IsZero(m.IPV4State) { // not required
		return nil
	}

	// value enum
	if err := m.validateIPV4StateEnum("ipv4_state", "body", *m.IPV4State); err != nil {
		return err
	}

	return nil
}

func (m *LdapStatus) validateIPV6(formats strfmt.Registry) error {
	if swag.IsZero(m.IPV6) { // not required
		return nil
	}

	if m.IPV6 != nil {
		if err := m.IPV6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipv6")
			}
			return err
		}
	}

	return nil
}

var ldapStatusTypeIPV6StatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["up","down"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapStatusTypeIPV6StatePropEnum = append(ldapStatusTypeIPV6StatePropEnum, v)
	}
}

const (

	// LdapStatusIPV6StateUp captures enum value "up"
	LdapStatusIPV6StateUp string = "up"

	// LdapStatusIPV6StateDown captures enum value "down"
	LdapStatusIPV6StateDown string = "down"
)

// prop value enum
func (m *LdapStatus) validateIPV6StateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapStatusTypeIPV6StatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapStatus) validateIPV6State(formats strfmt.Registry) error {
	if swag.IsZero(m.IPV6State) { // not required
		return nil
	}

	// value enum
	if err := m.validateIPV6StateEnum("ipv6_state", "body", *m.IPV6State); err != nil {
		return err
	}

	return nil
}

var ldapStatusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["up","down"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapStatusTypeStatePropEnum = append(ldapStatusTypeStatePropEnum, v)
	}
}

const (

	// LdapStatusStateUp captures enum value "up"
	LdapStatusStateUp string = "up"

	// LdapStatusStateDown captures enum value "down"
	LdapStatusStateDown string = "down"
)

// prop value enum
func (m *LdapStatus) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapStatusTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapStatus) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ldap status based on the context it is used
func (m *LdapStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPV4(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPV6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapStatus) contextValidateIPV4(ctx context.Context, formats strfmt.Registry) error {

	if m.IPV4 != nil {

		if swag.IsZero(m.IPV4) { // not required
			return nil
		}

		if err := m.IPV4.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv4")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipv4")
			}
			return err
		}
	}

	return nil
}

func (m *LdapStatus) contextValidateIPV6(ctx context.Context, formats strfmt.Registry) error {

	if m.IPV6 != nil {

		if swag.IsZero(m.IPV6) { // not required
			return nil
		}

		if err := m.IPV6.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipv6")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LdapStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapStatus) UnmarshalBinary(b []byte) error {
	var res LdapStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LdapStatusInlineIPV4 ldap status inline ipv4
//
// swagger:model ldap_status_inline_ipv4
type LdapStatusInlineIPV4 struct {

	// Code corresponding to the error message. If there is no error, it is 0 to indicate success.
	//
	// Example: 65537300
	Code *int64 `json:"code,omitempty"`

	// dn messages
	DnMessages []*string `json:"dn_messages,omitempty"`

	// Provides additional details on the error.
	//
	Message *string `json:"message,omitempty"`

	// Status of the LDAP service.
	//
	// Enum: ["up","down"]
	State *string `json:"state,omitempty"`
}

// Validate validates this ldap status inline ipv4
func (m *LdapStatusInlineIPV4) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ldapStatusInlineIpv4TypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["up","down"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapStatusInlineIpv4TypeStatePropEnum = append(ldapStatusInlineIpv4TypeStatePropEnum, v)
	}
}

const (

	// LdapStatusInlineIPV4StateUp captures enum value "up"
	LdapStatusInlineIPV4StateUp string = "up"

	// LdapStatusInlineIPV4StateDown captures enum value "down"
	LdapStatusInlineIPV4StateDown string = "down"
)

// prop value enum
func (m *LdapStatusInlineIPV4) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapStatusInlineIpv4TypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapStatusInlineIPV4) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("ipv4"+"."+"state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ldap status inline ipv4 based on context it is used
func (m *LdapStatusInlineIPV4) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LdapStatusInlineIPV4) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapStatusInlineIPV4) UnmarshalBinary(b []byte) error {
	var res LdapStatusInlineIPV4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LdapStatusInlineIPV6 ldap status inline ipv6
//
// swagger:model ldap_status_inline_ipv6
type LdapStatusInlineIPV6 struct {

	// Code corresponding to the error message. If there is no error, it is 0 to indicate success.
	//
	// Example: 65537300
	Code *int64 `json:"code,omitempty"`

	// dn messages
	DnMessages []*string `json:"dn_messages,omitempty"`

	// Provides additional details on the error.
	//
	Message *string `json:"message,omitempty"`

	// Status of the LDAP service.
	//
	// Enum: ["up","down"]
	State *string `json:"state,omitempty"`
}

// Validate validates this ldap status inline ipv6
func (m *LdapStatusInlineIPV6) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ldapStatusInlineIpv6TypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["up","down"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapStatusInlineIpv6TypeStatePropEnum = append(ldapStatusInlineIpv6TypeStatePropEnum, v)
	}
}

const (

	// LdapStatusInlineIPV6StateUp captures enum value "up"
	LdapStatusInlineIPV6StateUp string = "up"

	// LdapStatusInlineIPV6StateDown captures enum value "down"
	LdapStatusInlineIPV6StateDown string = "down"
)

// prop value enum
func (m *LdapStatusInlineIPV6) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapStatusInlineIpv6TypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapStatusInlineIPV6) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("ipv6"+"."+"state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ldap status inline ipv6 based on context it is used
func (m *LdapStatusInlineIPV6) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LdapStatusInlineIPV6) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapStatusInlineIPV6) UnmarshalBinary(b []byte) error {
	var res LdapStatusInlineIPV6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
