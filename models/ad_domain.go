// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdDomain ad domain
//
// swagger:model ad_domain
type AdDomain struct {

	// The default site used by LIFs that do not have a site membership.
	DefaultSite *string `json:"default_site,omitempty"`

	// The fully qualified domain name of the Windows Active Directory to which this CIFS server belongs. A CIFS server appears as a member of Windows server object in the Active Directory store. POST and PATCH only.
	// Example: example.com
	Fqdn *string `json:"fqdn,omitempty"`

	// Specifies the organizational unit within the Active Directory domain to associate with the CIFS server. POST and PATCH only.
	OrganizationalUnit *string `json:"organizational_unit,omitempty"`

	// The account password used to add this CIFS server to the Active Directory. This is not audited.
	Password *string `json:"password,omitempty"`

	// The user account used to add this CIFS server to the Active Directory.
	User *string `json:"user,omitempty"`
}

// Validate validates this ad domain
func (m *AdDomain) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ad domain based on context it is used
func (m *AdDomain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AdDomain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdDomain) UnmarshalBinary(b []byte) error {
	var res AdDomain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
