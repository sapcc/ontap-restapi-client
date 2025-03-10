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

// SnaplockLogFile snaplock log file
//
// swagger:model snaplock_log_file
type SnaplockLogFile struct {

	// links
	Links *SnaplockLogFileInlineLinks `json:"_links,omitempty"`

	// Base name of log file
	// Read Only: true
	// Enum: ["legal_hold","privileged_delete","system"]
	BaseName *string `json:"base_name,omitempty"`

	// Expiry time of the log file in date-time format. Value '9999-12-31T00:00:00Z' indicates infinite expiry time.
	// Example: 2058-06-04 19:00:00
	// Read Only: true
	// Format: date-time
	ExpiryTime *strfmt.DateTime `json:"expiry_time,omitempty"`

	// Absolute path of the log file in the volume
	// Example: /snaplock_log/system_logs/20180822_005947_GMT-present
	// Read Only: true
	Path *string `json:"path,omitempty"`

	// Size of the log file in bytes
	// Example: 20000
	// Read Only: true
	Size *int64 `json:"size,omitempty"`
}

// Validate validates this snaplock log file
func (m *SnaplockLogFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaseName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiryTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogFile) validateLinks(formats strfmt.Registry) error {
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

var snaplockLogFileTypeBaseNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["legal_hold","privileged_delete","system"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snaplockLogFileTypeBaseNamePropEnum = append(snaplockLogFileTypeBaseNamePropEnum, v)
	}
}

const (

	// SnaplockLogFileBaseNameLegalHold captures enum value "legal_hold"
	SnaplockLogFileBaseNameLegalHold string = "legal_hold"

	// SnaplockLogFileBaseNamePrivilegedDelete captures enum value "privileged_delete"
	SnaplockLogFileBaseNamePrivilegedDelete string = "privileged_delete"

	// SnaplockLogFileBaseNameSystem captures enum value "system"
	SnaplockLogFileBaseNameSystem string = "system"
)

// prop value enum
func (m *SnaplockLogFile) validateBaseNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snaplockLogFileTypeBaseNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnaplockLogFile) validateBaseName(formats strfmt.Registry) error {
	if swag.IsZero(m.BaseName) { // not required
		return nil
	}

	// value enum
	if err := m.validateBaseNameEnum("base_name", "body", *m.BaseName); err != nil {
		return err
	}

	return nil
}

func (m *SnaplockLogFile) validateExpiryTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiryTime) { // not required
		return nil
	}

	if err := validate.FormatOf("expiry_time", "body", "date-time", m.ExpiryTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this snaplock log file based on the context it is used
func (m *SnaplockLogFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBaseName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpiryTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogFile) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SnaplockLogFile) contextValidateBaseName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "base_name", "body", m.BaseName); err != nil {
		return err
	}

	return nil
}

func (m *SnaplockLogFile) contextValidateExpiryTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "expiry_time", "body", m.ExpiryTime); err != nil {
		return err
	}

	return nil
}

func (m *SnaplockLogFile) contextValidatePath(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *SnaplockLogFile) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnaplockLogFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnaplockLogFile) UnmarshalBinary(b []byte) error {
	var res SnaplockLogFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnaplockLogFileInlineLinks snaplock log file inline links
//
// swagger:model snaplock_log_file_inline__links
type SnaplockLogFileInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snaplock log file inline links
func (m *SnaplockLogFileInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogFileInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this snaplock log file inline links based on the context it is used
func (m *SnaplockLogFileInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogFileInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SnaplockLogFileInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnaplockLogFileInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnaplockLogFileInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
