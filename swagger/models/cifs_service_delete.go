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

// CifsServiceDelete cifs service delete
//
// swagger:model cifs_service_delete
type CifsServiceDelete struct {

	// ad domain
	AdDomain *AdDomainDelete `json:"ad_domain,omitempty"`

	// Specifies the type of user who can access the SMB Volume. The default is domain_user. In the case of a hybrid-user, ONTAP won't contact on-premise ADDS.
	//
	// Enum: ["domain_user","hybrid_user"]
	AuthUserType *string `json:"auth_user_type,omitempty"`

	// Specifies the authentication method.
	// The available values are:
	//   * client_secret
	//   * certificate
	//
	// Enum: ["client_secret","certificate"]
	AuthenticationMethod *string `json:"authentication_method,omitempty"`

	// PKCS12 certificate used by the application to prove its identity to AKV.
	// Example: PEM Cert
	// Format: password
	ClientCertificate *strfmt.Password `json:"client_certificate,omitempty"`

	// Application client ID of the deployed Azure application with appropriate access to an AKV or EntraId.
	// Example: e959d1b5-5a63-4284-9268-851e30e3eceb
	ClientID *string `json:"client_id,omitempty"`

	// Secret used by the application to prove its identity to AKV.
	// Example: \u003cid_value\u003e
	// Format: password
	ClientSecret *strfmt.Password `json:"client_secret,omitempty"`

	// URI of the deployed AKV that is used by ONTAP for storing keys.
	// Example: https://kmip-akv-keyvault.vault.azure.net/
	// Format: uri
	KeyVaultURI *strfmt.URI `json:"key_vault_uri,omitempty"`

	// Open authorization server host name.
	// Example: login.microsoftonline.com
	OauthHost *string `json:"oauth_host,omitempty"`

	// Proxy host.
	// Example: proxy.eng.com
	ProxyHost *string `json:"proxy_host,omitempty"`

	// Proxy password. Password is not audited.
	// Example: proxypassword
	ProxyPassword *string `json:"proxy_password,omitempty"`

	// Proxy port.
	// Example: 1234
	ProxyPort *int64 `json:"proxy_port,omitempty"`

	// Proxy type.
	// Enum: ["http","https"]
	ProxyType *string `json:"proxy_type,omitempty"`

	// Proxy username.
	// Example: proxyuser
	ProxyUsername *string `json:"proxy_username,omitempty"`

	// Directory (tenant) ID of the deployed Azure application with appropriate access to an AKV or EntraId.
	// Example: c9f32fcb-4ab7-40fe-af1b-1850d46cfbbe
	TenantID *string `json:"tenant_id,omitempty"`

	// AKV connection timeout, in seconds. The allowed range is between 0 to 30 seconds.
	// Example: 25
	Timeout *int64 `json:"timeout,omitempty"`

	// Verify the identity of the AKV host name. By default, verify_host is set to true.
	VerifyHost *bool `json:"verify_host,omitempty"`

	// The workgroup name.
	// Example: workgrp1
	// Max Length: 15
	// Min Length: 1
	Workgroup *string `json:"workgroup,omitempty"`
}

// Validate validates this cifs service delete
func (m *CifsServiceDelete) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthUserType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyVaultURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkgroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsServiceDelete) validateAdDomain(formats strfmt.Registry) error {
	if swag.IsZero(m.AdDomain) { // not required
		return nil
	}

	if m.AdDomain != nil {
		if err := m.AdDomain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ad_domain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ad_domain")
			}
			return err
		}
	}

	return nil
}

var cifsServiceDeleteTypeAuthUserTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["domain_user","hybrid_user"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cifsServiceDeleteTypeAuthUserTypePropEnum = append(cifsServiceDeleteTypeAuthUserTypePropEnum, v)
	}
}

const (

	// CifsServiceDeleteAuthUserTypeDomainUser captures enum value "domain_user"
	CifsServiceDeleteAuthUserTypeDomainUser string = "domain_user"

	// CifsServiceDeleteAuthUserTypeHybridUser captures enum value "hybrid_user"
	CifsServiceDeleteAuthUserTypeHybridUser string = "hybrid_user"
)

// prop value enum
func (m *CifsServiceDelete) validateAuthUserTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cifsServiceDeleteTypeAuthUserTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CifsServiceDelete) validateAuthUserType(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthUserType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthUserTypeEnum("auth_user_type", "body", *m.AuthUserType); err != nil {
		return err
	}

	return nil
}

var cifsServiceDeleteTypeAuthenticationMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["client_secret","certificate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cifsServiceDeleteTypeAuthenticationMethodPropEnum = append(cifsServiceDeleteTypeAuthenticationMethodPropEnum, v)
	}
}

const (

	// CifsServiceDeleteAuthenticationMethodClientSecret captures enum value "client_secret"
	CifsServiceDeleteAuthenticationMethodClientSecret string = "client_secret"

	// CifsServiceDeleteAuthenticationMethodCertificate captures enum value "certificate"
	CifsServiceDeleteAuthenticationMethodCertificate string = "certificate"
)

// prop value enum
func (m *CifsServiceDelete) validateAuthenticationMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cifsServiceDeleteTypeAuthenticationMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CifsServiceDelete) validateAuthenticationMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationMethodEnum("authentication_method", "body", *m.AuthenticationMethod); err != nil {
		return err
	}

	return nil
}

func (m *CifsServiceDelete) validateClientCertificate(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientCertificate) { // not required
		return nil
	}

	if err := validate.FormatOf("client_certificate", "body", "password", m.ClientCertificate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CifsServiceDelete) validateClientSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientSecret) { // not required
		return nil
	}

	if err := validate.FormatOf("client_secret", "body", "password", m.ClientSecret.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CifsServiceDelete) validateKeyVaultURI(formats strfmt.Registry) error {
	if swag.IsZero(m.KeyVaultURI) { // not required
		return nil
	}

	if err := validate.FormatOf("key_vault_uri", "body", "uri", m.KeyVaultURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var cifsServiceDeleteTypeProxyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http","https"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cifsServiceDeleteTypeProxyTypePropEnum = append(cifsServiceDeleteTypeProxyTypePropEnum, v)
	}
}

const (

	// CifsServiceDeleteProxyTypeHTTP captures enum value "http"
	CifsServiceDeleteProxyTypeHTTP string = "http"

	// CifsServiceDeleteProxyTypeHTTPS captures enum value "https"
	CifsServiceDeleteProxyTypeHTTPS string = "https"
)

// prop value enum
func (m *CifsServiceDelete) validateProxyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cifsServiceDeleteTypeProxyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CifsServiceDelete) validateProxyType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProxyType) { // not required
		return nil
	}

	// value enum
	if err := m.validateProxyTypeEnum("proxy_type", "body", *m.ProxyType); err != nil {
		return err
	}

	return nil
}

func (m *CifsServiceDelete) validateWorkgroup(formats strfmt.Registry) error {
	if swag.IsZero(m.Workgroup) { // not required
		return nil
	}

	if err := validate.MinLength("workgroup", "body", *m.Workgroup, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("workgroup", "body", *m.Workgroup, 15); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cifs service delete based on the context it is used
func (m *CifsServiceDelete) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdDomain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsServiceDelete) contextValidateAdDomain(ctx context.Context, formats strfmt.Registry) error {

	if m.AdDomain != nil {

		if swag.IsZero(m.AdDomain) { // not required
			return nil
		}

		if err := m.AdDomain.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ad_domain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ad_domain")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CifsServiceDelete) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsServiceDelete) UnmarshalBinary(b []byte) error {
	var res CifsServiceDelete
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
