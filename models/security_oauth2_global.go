// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SecurityOauth2Global security oauth2 global
//
// swagger:model security_oauth2_global
type SecurityOauth2Global struct {

	// Indicates whether OAuth 2.0 is enabled or disabled globally. Default value is false.
	Enabled *bool `json:"enabled,omitempty"`
}

// Validate validates this security oauth2 global
func (m *SecurityOauth2Global) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this security oauth2 global based on context it is used
func (m *SecurityOauth2Global) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecurityOauth2Global) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityOauth2Global) UnmarshalBinary(b []byte) error {
	var res SecurityOauth2Global
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
