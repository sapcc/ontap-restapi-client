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

// LdapIPStatus ldap ip status
//
// swagger:model ldap_ip_status
type LdapIPStatus struct {

	// Code corresponding to the error message. If there is no error, it is 0 to indicate success.
	//
	// Example: 65537300
	Code *int64 `json:"code,omitempty"`

	// ldap ip status inline dn messages
	LdapIPStatusInlineDnMessages []*string `json:"dn_messages,omitempty"`

	// Provides additional details on the error.
	//
	Message *string `json:"message,omitempty"`

	// Status of the LDAP service.
	//
	// Enum: ["up","down"]
	State *string `json:"state,omitempty"`
}

// Validate validates this ldap ip status
func (m *LdapIPStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ldapIpStatusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["up","down"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapIpStatusTypeStatePropEnum = append(ldapIpStatusTypeStatePropEnum, v)
	}
}

const (

	// LdapIPStatusStateUp captures enum value "up"
	LdapIPStatusStateUp string = "up"

	// LdapIPStatusStateDown captures enum value "down"
	LdapIPStatusStateDown string = "down"
)

// prop value enum
func (m *LdapIPStatus) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapIpStatusTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapIPStatus) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ldap ip status based on context it is used
func (m *LdapIPStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LdapIPStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapIPStatus) UnmarshalBinary(b []byte) error {
	var res LdapIPStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
