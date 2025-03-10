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

// SnaplockLogArchive snaplock log archive
//
// swagger:model snaplock_log_archive
type SnaplockLogArchive struct {

	// links
	Links *SnaplockLogArchiveInlineLinks `json:"_links,omitempty"`

	// Archive the specified SnapLock log file for the given base_name, and create a new log file. If base_name is not mentioned, archive all log files.
	Archive *bool `json:"archive,omitempty"`

	// Base name of log archive
	// Enum: ["legal_hold","privileged_delete","system"]
	BaseName *string `json:"base_name,omitempty"`
}

// Validate validates this snaplock log archive
func (m *SnaplockLogArchive) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaseName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogArchive) validateLinks(formats strfmt.Registry) error {
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

var snaplockLogArchiveTypeBaseNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["legal_hold","privileged_delete","system"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snaplockLogArchiveTypeBaseNamePropEnum = append(snaplockLogArchiveTypeBaseNamePropEnum, v)
	}
}

const (

	// SnaplockLogArchiveBaseNameLegalHold captures enum value "legal_hold"
	SnaplockLogArchiveBaseNameLegalHold string = "legal_hold"

	// SnaplockLogArchiveBaseNamePrivilegedDelete captures enum value "privileged_delete"
	SnaplockLogArchiveBaseNamePrivilegedDelete string = "privileged_delete"

	// SnaplockLogArchiveBaseNameSystem captures enum value "system"
	SnaplockLogArchiveBaseNameSystem string = "system"
)

// prop value enum
func (m *SnaplockLogArchive) validateBaseNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snaplockLogArchiveTypeBaseNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnaplockLogArchive) validateBaseName(formats strfmt.Registry) error {
	if swag.IsZero(m.BaseName) { // not required
		return nil
	}

	// value enum
	if err := m.validateBaseNameEnum("base_name", "body", *m.BaseName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this snaplock log archive based on the context it is used
func (m *SnaplockLogArchive) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogArchive) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *SnaplockLogArchive) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnaplockLogArchive) UnmarshalBinary(b []byte) error {
	var res SnaplockLogArchive
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnaplockLogArchiveInlineLinks snaplock log archive inline links
//
// swagger:model snaplock_log_archive_inline__links
type SnaplockLogArchiveInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snaplock log archive inline links
func (m *SnaplockLogArchiveInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogArchiveInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this snaplock log archive inline links based on the context it is used
func (m *SnaplockLogArchiveInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogArchiveInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SnaplockLogArchiveInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnaplockLogArchiveInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnaplockLogArchiveInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
