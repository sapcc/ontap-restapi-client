// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SoftwareImagesArrayInline List of software image information.
//
// swagger:model software_images_array_inline
type SoftwareImagesArrayInline []*SoftwareImagesArrayInlineInlineArrayItem

// Validate validates this software images array inline
func (m SoftwareImagesArrayInline) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this software images array inline based on the context it is used
func (m SoftwareImagesArrayInline) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := validate.ReadOnly(ctx, "", "body", []*SoftwareImagesArrayInlineInlineArrayItem(m)); err != nil {
		return err
	}

	for i := 0; i < len(m); i++ {

		if m[i] != nil {

			if swag.IsZero(m[i]) { // not required
				return nil
			}

			if err := m[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// SoftwareImagesArrayInlineInlineArrayItem software images array inline inline array item
//
// swagger:model software_images_array_inline_inline_array_item
type SoftwareImagesArrayInlineInlineArrayItem struct {

	// Package file name.
	// Example: image.tgz
	// Read Only: true
	Package *string `json:"package,omitempty"`
}

// Validate validates this software images array inline inline array item
func (m *SoftwareImagesArrayInlineInlineArrayItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this software images array inline inline array item based on the context it is used
func (m *SoftwareImagesArrayInlineInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePackage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareImagesArrayInlineInlineArrayItem) contextValidatePackage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "package", "body", m.Package); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareImagesArrayInlineInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareImagesArrayInlineInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res SoftwareImagesArrayInlineInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
