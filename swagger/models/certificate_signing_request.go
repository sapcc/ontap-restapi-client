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

// CertificateSigningRequest certificate signing request
//
// swagger:model certificate_signing_request
type CertificateSigningRequest struct {

	// links
	Links *CertificateSigningRequestInlineLinks `json:"_links,omitempty"`

	// Asymmetric Encryption Algorithm.
	// Enum: ["rsa","ec"]
	Algorithm *string `json:"algorithm,omitempty"`

	// A Certificate Signing Request (CSR) provided to a CA for obtaining a CA-signed certificate.
	// Read Only: true
	Csr *string `json:"csr,omitempty"`

	// A list of extended key usage extensions.
	ExtendedKeyUsages []*string `json:"extended_key_usages,omitempty"`

	// Private key generated for the CSR.
	// Read Only: true
	GeneratedPrivateKey *string `json:"generated_private_key,omitempty"`

	// Hashing function.
	// Enum: ["sha256","sha224","sha384","sha512"]
	HashFunction *string `json:"hash_function,omitempty"`

	// A list of key usage extensions.
	KeyUsages []*string `json:"key_usages,omitempty"`

	// Security strength of the certificate in bits.
	// Enum: [112,128,192]
	SecurityStrength *int64 `json:"security_strength,omitempty"`

	// subject alternatives
	SubjectAlternatives *CertificateSigningRequestInlineSubjectAlternatives `json:"subject_alternatives,omitempty"`

	// Subject name details of the certificate. The format is a list of comma separated key=value pairs.
	// Example: C=US,O=NTAP,CN=test.domain.com
	SubjectName *string `json:"subject_name,omitempty"`
}

// Validate validates this certificate signing request
func (m *CertificateSigningRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtendedKeyUsages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHashFunction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyUsages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityStrength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubjectAlternatives(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateSigningRequest) validateLinks(formats strfmt.Registry) error {
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

var certificateSigningRequestTypeAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["rsa","ec"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificateSigningRequestTypeAlgorithmPropEnum = append(certificateSigningRequestTypeAlgorithmPropEnum, v)
	}
}

const (

	// CertificateSigningRequestAlgorithmRsa captures enum value "rsa"
	CertificateSigningRequestAlgorithmRsa string = "rsa"

	// CertificateSigningRequestAlgorithmEc captures enum value "ec"
	CertificateSigningRequestAlgorithmEc string = "ec"
)

// prop value enum
func (m *CertificateSigningRequest) validateAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificateSigningRequestTypeAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificateSigningRequest) validateAlgorithm(formats strfmt.Registry) error {
	if swag.IsZero(m.Algorithm) { // not required
		return nil
	}

	// value enum
	if err := m.validateAlgorithmEnum("algorithm", "body", *m.Algorithm); err != nil {
		return err
	}

	return nil
}

var certificateSigningRequestExtendedKeyUsagesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["serverauth","clientauth","timestamping","dvcs","ocspsigning","codesigning","emailprotection","anyextendedkeyusage","critical"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificateSigningRequestExtendedKeyUsagesItemsEnum = append(certificateSigningRequestExtendedKeyUsagesItemsEnum, v)
	}
}

func (m *CertificateSigningRequest) validateExtendedKeyUsagesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificateSigningRequestExtendedKeyUsagesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificateSigningRequest) validateExtendedKeyUsages(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtendedKeyUsages) { // not required
		return nil
	}

	for i := 0; i < len(m.ExtendedKeyUsages); i++ {
		if swag.IsZero(m.ExtendedKeyUsages[i]) { // not required
			continue
		}

		// value enum
		if err := m.validateExtendedKeyUsagesItemsEnum("extended_key_usages"+"."+strconv.Itoa(i), "body", *m.ExtendedKeyUsages[i]); err != nil {
			return err
		}

	}

	return nil
}

var certificateSigningRequestTypeHashFunctionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha256","sha224","sha384","sha512"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificateSigningRequestTypeHashFunctionPropEnum = append(certificateSigningRequestTypeHashFunctionPropEnum, v)
	}
}

const (

	// CertificateSigningRequestHashFunctionSha256 captures enum value "sha256"
	CertificateSigningRequestHashFunctionSha256 string = "sha256"

	// CertificateSigningRequestHashFunctionSha224 captures enum value "sha224"
	CertificateSigningRequestHashFunctionSha224 string = "sha224"

	// CertificateSigningRequestHashFunctionSha384 captures enum value "sha384"
	CertificateSigningRequestHashFunctionSha384 string = "sha384"

	// CertificateSigningRequestHashFunctionSha512 captures enum value "sha512"
	CertificateSigningRequestHashFunctionSha512 string = "sha512"
)

// prop value enum
func (m *CertificateSigningRequest) validateHashFunctionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificateSigningRequestTypeHashFunctionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificateSigningRequest) validateHashFunction(formats strfmt.Registry) error {
	if swag.IsZero(m.HashFunction) { // not required
		return nil
	}

	// value enum
	if err := m.validateHashFunctionEnum("hash_function", "body", *m.HashFunction); err != nil {
		return err
	}

	return nil
}

var certificateSigningRequestKeyUsagesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["digitalsignature","nonrepudiation","keyencipherment","dataencipherment","keyagreement","keycertsign","crlsign","encipheronly","decipheronly","critical"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificateSigningRequestKeyUsagesItemsEnum = append(certificateSigningRequestKeyUsagesItemsEnum, v)
	}
}

func (m *CertificateSigningRequest) validateKeyUsagesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificateSigningRequestKeyUsagesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificateSigningRequest) validateKeyUsages(formats strfmt.Registry) error {
	if swag.IsZero(m.KeyUsages) { // not required
		return nil
	}

	for i := 0; i < len(m.KeyUsages); i++ {
		if swag.IsZero(m.KeyUsages[i]) { // not required
			continue
		}

		// value enum
		if err := m.validateKeyUsagesItemsEnum("key_usages"+"."+strconv.Itoa(i), "body", *m.KeyUsages[i]); err != nil {
			return err
		}

	}

	return nil
}

var certificateSigningRequestTypeSecurityStrengthPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[112,128,192]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificateSigningRequestTypeSecurityStrengthPropEnum = append(certificateSigningRequestTypeSecurityStrengthPropEnum, v)
	}
}

// prop value enum
func (m *CertificateSigningRequest) validateSecurityStrengthEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, certificateSigningRequestTypeSecurityStrengthPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificateSigningRequest) validateSecurityStrength(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityStrength) { // not required
		return nil
	}

	// value enum
	if err := m.validateSecurityStrengthEnum("security_strength", "body", *m.SecurityStrength); err != nil {
		return err
	}

	return nil
}

func (m *CertificateSigningRequest) validateSubjectAlternatives(formats strfmt.Registry) error {
	if swag.IsZero(m.SubjectAlternatives) { // not required
		return nil
	}

	if m.SubjectAlternatives != nil {
		if err := m.SubjectAlternatives.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subject_alternatives")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subject_alternatives")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this certificate signing request based on the context it is used
func (m *CertificateSigningRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCsr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGeneratedPrivateKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubjectAlternatives(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateSigningRequest) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CertificateSigningRequest) contextValidateCsr(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "csr", "body", m.Csr); err != nil {
		return err
	}

	return nil
}

func (m *CertificateSigningRequest) contextValidateGeneratedPrivateKey(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "generated_private_key", "body", m.GeneratedPrivateKey); err != nil {
		return err
	}

	return nil
}

func (m *CertificateSigningRequest) contextValidateSubjectAlternatives(ctx context.Context, formats strfmt.Registry) error {

	if m.SubjectAlternatives != nil {

		if swag.IsZero(m.SubjectAlternatives) { // not required
			return nil
		}

		if err := m.SubjectAlternatives.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subject_alternatives")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subject_alternatives")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificateSigningRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateSigningRequest) UnmarshalBinary(b []byte) error {
	var res CertificateSigningRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificateSigningRequestInlineLinks certificate signing request inline links
//
// swagger:model certificate_signing_request_inline__links
type CertificateSigningRequestInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this certificate signing request inline links
func (m *CertificateSigningRequestInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateSigningRequestInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this certificate signing request inline links based on the context it is used
func (m *CertificateSigningRequestInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateSigningRequestInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CertificateSigningRequestInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateSigningRequestInlineLinks) UnmarshalBinary(b []byte) error {
	var res CertificateSigningRequestInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificateSigningRequestInlineSubjectAlternatives certificate signing request inline subject alternatives
//
// swagger:model certificate_signing_request_inline_subject_alternatives
type CertificateSigningRequestInlineSubjectAlternatives struct {

	// A list of DNS names for Subject Alternate name extension.
	DNS []*string `json:"dns,omitempty"`

	// A list of email addresses for Subject Alternate name extension
	Email []*string `json:"email,omitempty"`

	// A list of IP addresses for Subject Alternate name extension.
	IP []*string `json:"ip,omitempty"`

	// A list of URIs for Subject Alternate name extension.
	URI []*string `json:"uri,omitempty"`
}

// Validate validates this certificate signing request inline subject alternatives
func (m *CertificateSigningRequestInlineSubjectAlternatives) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this certificate signing request inline subject alternatives based on context it is used
func (m *CertificateSigningRequestInlineSubjectAlternatives) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateSigningRequestInlineSubjectAlternatives) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateSigningRequestInlineSubjectAlternatives) UnmarshalBinary(b []byte) error {
	var res CertificateSigningRequestInlineSubjectAlternatives
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
