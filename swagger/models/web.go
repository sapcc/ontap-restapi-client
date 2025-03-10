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

// Web web
//
// swagger:model web
type Web struct {

	// links
	Links *WebInlineLinks `json:"_links,omitempty"`

	// certificate
	Certificate *WebInlineCertificate `json:"certificate,omitempty"`

	// Indicates whether client authentication is enabled.
	ClientEnabled *bool `json:"client_enabled,omitempty"`

	// csrf
	Csrf *WebInlineCsrf `json:"csrf,omitempty"`

	// Indicates whether remote clients can connect to the web services.
	Enabled *bool `json:"enabled,omitempty"`

	// Indicates whether HTTP is enabled.
	HTTPEnabled *bool `json:"http_enabled,omitempty"`

	// HTTP port for cluster-level web services.
	HTTPPort *int64 `json:"http_port,omitempty"`

	// HTTPS port for cluster-level web services.
	HTTPSPort *int64 `json:"https_port,omitempty"`

	// Indicates whether online certificate status protocol verification is enabled.
	OcspEnabled *bool `json:"ocsp_enabled,omitempty"`

	// The number of connections that can be processed concurrently from the same remote address.
	// Example: 42
	// Maximum: 999
	// Minimum: 24
	PerAddressLimit *int64 `json:"per_address_limit,omitempty"`

	// State of the cluster-level web services.
	// Read Only: true
	// Enum: ["offline","partial","mixed","online","unclustered"]
	State *string `json:"state,omitempty"`

	// The maximum size of the wait queue for connections exceeding the per-address-limit.
	WaitQueueCapacity *int64 `json:"wait_queue_capacity,omitempty"`
}

// Validate validates this web
func (m *Web) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCsrf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerAddressLimit(formats); err != nil {
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

func (m *Web) validateLinks(formats strfmt.Registry) error {
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

func (m *Web) validateCertificate(formats strfmt.Registry) error {
	if swag.IsZero(m.Certificate) { // not required
		return nil
	}

	if m.Certificate != nil {
		if err := m.Certificate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificate")
			}
			return err
		}
	}

	return nil
}

func (m *Web) validateCsrf(formats strfmt.Registry) error {
	if swag.IsZero(m.Csrf) { // not required
		return nil
	}

	if m.Csrf != nil {
		if err := m.Csrf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csrf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("csrf")
			}
			return err
		}
	}

	return nil
}

func (m *Web) validatePerAddressLimit(formats strfmt.Registry) error {
	if swag.IsZero(m.PerAddressLimit) { // not required
		return nil
	}

	if err := validate.MinimumInt("per_address_limit", "body", *m.PerAddressLimit, 24, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("per_address_limit", "body", *m.PerAddressLimit, 999, false); err != nil {
		return err
	}

	return nil
}

var webTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["offline","partial","mixed","online","unclustered"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webTypeStatePropEnum = append(webTypeStatePropEnum, v)
	}
}

const (

	// WebStateOffline captures enum value "offline"
	WebStateOffline string = "offline"

	// WebStatePartial captures enum value "partial"
	WebStatePartial string = "partial"

	// WebStateMixed captures enum value "mixed"
	WebStateMixed string = "mixed"

	// WebStateOnline captures enum value "online"
	WebStateOnline string = "online"

	// WebStateUnclustered captures enum value "unclustered"
	WebStateUnclustered string = "unclustered"
)

// prop value enum
func (m *Web) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, webTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Web) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this web based on the context it is used
func (m *Web) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCertificate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCsrf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Web) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Web) contextValidateCertificate(ctx context.Context, formats strfmt.Registry) error {

	if m.Certificate != nil {

		if swag.IsZero(m.Certificate) { // not required
			return nil
		}

		if err := m.Certificate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificate")
			}
			return err
		}
	}

	return nil
}

func (m *Web) contextValidateCsrf(ctx context.Context, formats strfmt.Registry) error {

	if m.Csrf != nil {

		if swag.IsZero(m.Csrf) { // not required
			return nil
		}

		if err := m.Csrf.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csrf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("csrf")
			}
			return err
		}
	}

	return nil
}

func (m *Web) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Web) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Web) UnmarshalBinary(b []byte) error {
	var res Web
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WebInlineCertificate Certificate used by cluster and node management interfaces for TLS connection requests.
//
// swagger:model web_inline_certificate
type WebInlineCertificate struct {

	// links
	Links *WebInlineCertificateInlineLinks `json:"_links,omitempty"`

	// Certificate name
	// Example: cert1
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// Certificate UUID
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this web inline certificate
func (m *WebInlineCertificate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebInlineCertificate) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificate" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this web inline certificate based on the context it is used
func (m *WebInlineCertificate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebInlineCertificate) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificate" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *WebInlineCertificate) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "certificate"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebInlineCertificate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebInlineCertificate) UnmarshalBinary(b []byte) error {
	var res WebInlineCertificate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WebInlineCertificateInlineLinks web inline certificate inline links
//
// swagger:model web_inline_certificate_inline__links
type WebInlineCertificateInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this web inline certificate inline links
func (m *WebInlineCertificateInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebInlineCertificateInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificate" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this web inline certificate inline links based on the context it is used
func (m *WebInlineCertificateInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebInlineCertificateInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificate" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebInlineCertificateInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebInlineCertificateInlineLinks) UnmarshalBinary(b []byte) error {
	var res WebInlineCertificateInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WebInlineCsrf web inline csrf
//
// swagger:model web_inline_csrf
type WebInlineCsrf struct {

	// Indicates whether CSRF protection is enabled.
	ProtectionEnabled *bool `json:"protection_enabled,omitempty"`

	// token
	Token *WebInlineCsrfInlineToken `json:"token,omitempty"`
}

// Validate validates this web inline csrf
func (m *WebInlineCsrf) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebInlineCsrf) validateToken(formats strfmt.Registry) error {
	if swag.IsZero(m.Token) { // not required
		return nil
	}

	if m.Token != nil {
		if err := m.Token.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csrf" + "." + "token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("csrf" + "." + "token")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this web inline csrf based on the context it is used
func (m *WebInlineCsrf) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebInlineCsrf) contextValidateToken(ctx context.Context, formats strfmt.Registry) error {

	if m.Token != nil {

		if swag.IsZero(m.Token) { // not required
			return nil
		}

		if err := m.Token.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csrf" + "." + "token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("csrf" + "." + "token")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebInlineCsrf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebInlineCsrf) UnmarshalBinary(b []byte) error {
	var res WebInlineCsrf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WebInlineCsrfInlineToken web inline csrf inline token
//
// swagger:model web_inline_csrf_inline_token
type WebInlineCsrfInlineToken struct {

	// Maximum number of concurrent CSRF tokens.
	// Example: 120
	// Maximum: 9999
	// Minimum: 100
	ConcurrentLimit *int64 `json:"concurrent_limit,omitempty"`

	// Time for which an unused CSRF token is retained, in seconds.
	// Example: 200
	// Minimum: 180
	IdleTimeout *int64 `json:"idle_timeout,omitempty"`

	// Time for which an unused CSRF token, regardless of usage is retained, in seconds.
	// Example: 1200
	// Minimum: 600
	MaxTimeout *int64 `json:"max_timeout,omitempty"`
}

// Validate validates this web inline csrf inline token
func (m *WebInlineCsrfInlineToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConcurrentLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdleTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxTimeout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebInlineCsrfInlineToken) validateConcurrentLimit(formats strfmt.Registry) error {
	if swag.IsZero(m.ConcurrentLimit) { // not required
		return nil
	}

	if err := validate.MinimumInt("csrf"+"."+"token"+"."+"concurrent_limit", "body", *m.ConcurrentLimit, 100, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("csrf"+"."+"token"+"."+"concurrent_limit", "body", *m.ConcurrentLimit, 9999, false); err != nil {
		return err
	}

	return nil
}

func (m *WebInlineCsrfInlineToken) validateIdleTimeout(formats strfmt.Registry) error {
	if swag.IsZero(m.IdleTimeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("csrf"+"."+"token"+"."+"idle_timeout", "body", *m.IdleTimeout, 180, false); err != nil {
		return err
	}

	return nil
}

func (m *WebInlineCsrfInlineToken) validateMaxTimeout(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxTimeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("csrf"+"."+"token"+"."+"max_timeout", "body", *m.MaxTimeout, 600, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this web inline csrf inline token based on context it is used
func (m *WebInlineCsrfInlineToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebInlineCsrfInlineToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebInlineCsrfInlineToken) UnmarshalBinary(b []byte) error {
	var res WebInlineCsrfInlineToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WebInlineLinks web inline links
//
// swagger:model web_inline__links
type WebInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this web inline links
func (m *WebInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this web inline links based on the context it is used
func (m *WebInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *WebInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebInlineLinks) UnmarshalBinary(b []byte) error {
	var res WebInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
