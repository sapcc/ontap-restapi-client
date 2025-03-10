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

// AccountApplication account application
//
// swagger:model account_application
type AccountApplication struct {

	// Applications
	// Enum: ["amqp","console","http","ontapi","service_processor","ssh"]
	Application *string `json:"application,omitempty"`

	// authentication methods
	AuthenticationMethods []*string `json:"authentication_methods,omitempty"`

	// Optional property that specifies the mode of authentication as LDAP Fastbind.
	IsLdapFastbind *bool `json:"is_ldap_fastbind,omitempty"`

	// Optional property that specifies whether the user is an LDAP or NIS group.
	IsNsSwitchGroup *bool `json:"is_ns_switch_group,omitempty"`

	// An optional additional authentication method for multi-factor authentication (MFA). This property is only supported for SSH (_ssh_) and Service Processor (_service_processor_) applications. It is ignored for all other applications. Time-based One-Time Passwords (TOTPs) are only supported with the authentication method password or public key. For the Service Processor (_service_processor_) application, _none_ and _publickey_ are the only supported enum values.
	// Enum: ["none","password","publickey","nsswitch","domain","totp"]
	SecondAuthenticationMethod *string `json:"second_authentication_method,omitempty"`
}

// Validate validates this account application
func (m *AccountApplication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationMethods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondAuthenticationMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var accountApplicationTypeApplicationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["amqp","console","http","ontapi","service_processor","ssh"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountApplicationTypeApplicationPropEnum = append(accountApplicationTypeApplicationPropEnum, v)
	}
}

const (

	// AccountApplicationApplicationAmqp captures enum value "amqp"
	AccountApplicationApplicationAmqp string = "amqp"

	// AccountApplicationApplicationConsole captures enum value "console"
	AccountApplicationApplicationConsole string = "console"

	// AccountApplicationApplicationHTTP captures enum value "http"
	AccountApplicationApplicationHTTP string = "http"

	// AccountApplicationApplicationOntapi captures enum value "ontapi"
	AccountApplicationApplicationOntapi string = "ontapi"

	// AccountApplicationApplicationServiceProcessor captures enum value "service_processor"
	AccountApplicationApplicationServiceProcessor string = "service_processor"

	// AccountApplicationApplicationSSH captures enum value "ssh"
	AccountApplicationApplicationSSH string = "ssh"
)

// prop value enum
func (m *AccountApplication) validateApplicationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountApplicationTypeApplicationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccountApplication) validateApplication(formats strfmt.Registry) error {
	if swag.IsZero(m.Application) { // not required
		return nil
	}

	// value enum
	if err := m.validateApplicationEnum("application", "body", *m.Application); err != nil {
		return err
	}

	return nil
}

var accountApplicationAuthenticationMethodsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["password","publickey","domain","nsswitch","certificate","saml"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountApplicationAuthenticationMethodsItemsEnum = append(accountApplicationAuthenticationMethodsItemsEnum, v)
	}
}

func (m *AccountApplication) validateAuthenticationMethodsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountApplicationAuthenticationMethodsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccountApplication) validateAuthenticationMethods(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationMethods) { // not required
		return nil
	}

	for i := 0; i < len(m.AuthenticationMethods); i++ {
		if swag.IsZero(m.AuthenticationMethods[i]) { // not required
			continue
		}

		// value enum
		if err := m.validateAuthenticationMethodsItemsEnum("authentication_methods"+"."+strconv.Itoa(i), "body", *m.AuthenticationMethods[i]); err != nil {
			return err
		}

	}

	return nil
}

var accountApplicationTypeSecondAuthenticationMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","password","publickey","nsswitch","domain","totp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountApplicationTypeSecondAuthenticationMethodPropEnum = append(accountApplicationTypeSecondAuthenticationMethodPropEnum, v)
	}
}

const (

	// AccountApplicationSecondAuthenticationMethodNone captures enum value "none"
	AccountApplicationSecondAuthenticationMethodNone string = "none"

	// AccountApplicationSecondAuthenticationMethodPassword captures enum value "password"
	AccountApplicationSecondAuthenticationMethodPassword string = "password"

	// AccountApplicationSecondAuthenticationMethodPublickey captures enum value "publickey"
	AccountApplicationSecondAuthenticationMethodPublickey string = "publickey"

	// AccountApplicationSecondAuthenticationMethodNsswitch captures enum value "nsswitch"
	AccountApplicationSecondAuthenticationMethodNsswitch string = "nsswitch"

	// AccountApplicationSecondAuthenticationMethodDomain captures enum value "domain"
	AccountApplicationSecondAuthenticationMethodDomain string = "domain"

	// AccountApplicationSecondAuthenticationMethodTotp captures enum value "totp"
	AccountApplicationSecondAuthenticationMethodTotp string = "totp"
)

// prop value enum
func (m *AccountApplication) validateSecondAuthenticationMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountApplicationTypeSecondAuthenticationMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccountApplication) validateSecondAuthenticationMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondAuthenticationMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateSecondAuthenticationMethodEnum("second_authentication_method", "body", *m.SecondAuthenticationMethod); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this account application based on context it is used
func (m *AccountApplication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountApplication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountApplication) UnmarshalBinary(b []byte) error {
	var res AccountApplication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
