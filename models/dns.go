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

// DNS dns
//
// swagger:model dns
type DNS struct {

	// links
	Links *DNSInlineLinks `json:"_links,omitempty"`

	// Number of attempts allowed when querying the DNS name servers.
	//
	// Example: 1
	// Maximum: 4
	// Minimum: 1
	Attempts *int64 `json:"attempts,omitempty"`

	// List of IP addresses for a DNS service. Addresses can be IPv4, IPv6 or both.
	//
	// Example: ["10.224.65.20","2001:db08:a0b:12f0::1"]
	// Read Only: true
	DNSInlineServiceIps []*string `json:"service_ips,omitempty"`

	// Status of all the DNS name servers configured for the specified SVM.
	//
	// Read Only: true
	DNSInlineStatus []*Status `json:"status,omitempty"`

	// domains
	Domains DNSDomainsArrayInline `json:"domains,omitempty"`

	// dynamic dns
	DynamicDNS *DNSInlineDynamicDNS `json:"dynamic_dns,omitempty"`

	// Indicates whether or not the query section of the reply packet is equal to that of the query packet.
	//
	PacketQueryMatch *bool `json:"packet_query_match,omitempty"`

	// Set to "svm" for DNS owned by an SVM, otherwise set to "cluster".
	//
	// Enum: ["svm","cluster"]
	Scope *string `json:"scope,omitempty"`

	// servers
	Servers NameServersArrayInline `json:"servers,omitempty"`

	// Indicates whether or not the validation for the specified DNS configuration is disabled.
	//
	SkipConfigValidation *bool `json:"skip_config_validation,omitempty"`

	// Indicates whether or not the DNS responses are from a different IP address to the IP address the request was sent to.
	//
	SourceAddressMatch *bool `json:"source_address_match,omitempty"`

	// svm
	Svm *DNSInlineSvm `json:"svm,omitempty"`

	// Timeout values for queries to the name servers, in seconds.
	//
	// Example: 2
	// Maximum: 5
	// Minimum: 1
	Timeout *int64 `json:"timeout,omitempty"`

	// Enable or disable top-level domain (TLD) queries.
	//
	TldQueryEnabled *bool `json:"tld_query_enabled,omitempty"`

	// UUID of the DNS object.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this dns
func (m *DNS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttempts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSInlineStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomains(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDynamicDNS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DNS) validateLinks(formats strfmt.Registry) error {
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

func (m *DNS) validateAttempts(formats strfmt.Registry) error {
	if swag.IsZero(m.Attempts) { // not required
		return nil
	}

	if err := validate.MinimumInt("attempts", "body", *m.Attempts, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("attempts", "body", *m.Attempts, 4, false); err != nil {
		return err
	}

	return nil
}

func (m *DNS) validateDNSInlineStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSInlineStatus) { // not required
		return nil
	}

	for i := 0; i < len(m.DNSInlineStatus); i++ {
		if swag.IsZero(m.DNSInlineStatus[i]) { // not required
			continue
		}

		if m.DNSInlineStatus[i] != nil {
			if err := m.DNSInlineStatus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DNS) validateDomains(formats strfmt.Registry) error {
	if swag.IsZero(m.Domains) { // not required
		return nil
	}

	if err := m.Domains.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("domains")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("domains")
		}
		return err
	}

	return nil
}

func (m *DNS) validateDynamicDNS(formats strfmt.Registry) error {
	if swag.IsZero(m.DynamicDNS) { // not required
		return nil
	}

	if m.DynamicDNS != nil {
		if err := m.DynamicDNS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dynamic_dns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dynamic_dns")
			}
			return err
		}
	}

	return nil
}

var dnsTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["svm","cluster"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dnsTypeScopePropEnum = append(dnsTypeScopePropEnum, v)
	}
}

const (

	// DNSScopeSvm captures enum value "svm"
	DNSScopeSvm string = "svm"

	// DNSScopeCluster captures enum value "cluster"
	DNSScopeCluster string = "cluster"
)

// prop value enum
func (m *DNS) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dnsTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DNS) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", *m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *DNS) validateServers(formats strfmt.Registry) error {
	if swag.IsZero(m.Servers) { // not required
		return nil
	}

	if err := m.Servers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("servers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("servers")
		}
		return err
	}

	return nil
}

func (m *DNS) validateSvm(formats strfmt.Registry) error {
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

func (m *DNS) validateTimeout(formats strfmt.Registry) error {
	if swag.IsZero(m.Timeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("timeout", "body", *m.Timeout, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("timeout", "body", *m.Timeout, 5, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this dns based on the context it is used
func (m *DNS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDNSInlineServiceIps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDNSInlineStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDomains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDynamicDNS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServers(ctx, formats); err != nil {
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

func (m *DNS) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DNS) contextValidateDNSInlineServiceIps(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "service_ips", "body", []*string(m.DNSInlineServiceIps)); err != nil {
		return err
	}

	return nil
}

func (m *DNS) contextValidateDNSInlineStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", []*Status(m.DNSInlineStatus)); err != nil {
		return err
	}

	for i := 0; i < len(m.DNSInlineStatus); i++ {

		if m.DNSInlineStatus[i] != nil {

			if swag.IsZero(m.DNSInlineStatus[i]) { // not required
				return nil
			}

			if err := m.DNSInlineStatus[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DNS) contextValidateDomains(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Domains.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("domains")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("domains")
		}
		return err
	}

	return nil
}

func (m *DNS) contextValidateDynamicDNS(ctx context.Context, formats strfmt.Registry) error {

	if m.DynamicDNS != nil {

		if swag.IsZero(m.DynamicDNS) { // not required
			return nil
		}

		if err := m.DynamicDNS.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dynamic_dns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dynamic_dns")
			}
			return err
		}
	}

	return nil
}

func (m *DNS) contextValidateServers(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Servers.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("servers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("servers")
		}
		return err
	}

	return nil
}

func (m *DNS) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DNS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNS) UnmarshalBinary(b []byte) error {
	var res DNS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DNSInlineDynamicDNS dns inline dynamic dns
//
// swagger:model dns_inline_dynamic_dns
type DNSInlineDynamicDNS struct {

	// Enable or disable Dynamic DNS (DDNS) updates for the specified SVM.
	//
	Enabled *bool `json:"enabled,omitempty"`

	// Fully Qualified Domain Name (FQDN) to be used for dynamic DNS updates.
	//
	// Example: example.com
	Fqdn *string `json:"fqdn,omitempty"`

	// Enable or disable FQDN validation.
	//
	SkipFqdnValidation *bool `json:"skip_fqdn_validation,omitempty"`

	// Time to live value for the dynamic DNS updates, in an ISO-8601 duration formatted string.
	// Maximum Time To Live is 720 hours(P30D in ISO-8601 format) and the default is 24 hours(P1D in ISO-8601 format).
	//
	// Example: P2D
	TimeToLive *string `json:"time_to_live,omitempty"`

	// Enable or disable secure dynamic DNS updates for the specified SVM.
	//
	UseSecure *bool `json:"use_secure,omitempty"`
}

// Validate validates this dns inline dynamic dns
func (m *DNSInlineDynamicDNS) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dns inline dynamic dns based on context it is used
func (m *DNSInlineDynamicDNS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DNSInlineDynamicDNS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNSInlineDynamicDNS) UnmarshalBinary(b []byte) error {
	var res DNSInlineDynamicDNS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DNSInlineLinks dns inline links
//
// swagger:model dns_inline__links
type DNSInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this dns inline links
func (m *DNSInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DNSInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this dns inline links based on the context it is used
func (m *DNSInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DNSInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DNSInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNSInlineLinks) UnmarshalBinary(b []byte) error {
	var res DNSInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DNSInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model dns_inline_svm
type DNSInlineSvm struct {

	// links
	Links *DNSInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this dns inline svm
func (m *DNSInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DNSInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this dns inline svm based on the context it is used
func (m *DNSInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DNSInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DNSInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNSInlineSvm) UnmarshalBinary(b []byte) error {
	var res DNSInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DNSInlineSvmInlineLinks dns inline svm inline links
//
// swagger:model dns_inline_svm_inline__links
type DNSInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this dns inline svm inline links
func (m *DNSInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DNSInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this dns inline svm inline links based on the context it is used
func (m *DNSInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DNSInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DNSInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNSInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res DNSInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
