// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RelatedLink related link
//
// swagger:model related_link
type RelatedLink struct {

	// related
	Related *Href `json:"related,omitempty"`
}

// Validate validates this related link
func (m *RelatedLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRelated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelatedLink) validateRelated(formats strfmt.Registry) error {
	if swag.IsZero(m.Related) { // not required
		return nil
	}

	if m.Related != nil {
		if err := m.Related.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("related")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("related")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this related link based on the context it is used
func (m *RelatedLink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRelated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelatedLink) contextValidateRelated(ctx context.Context, formats strfmt.Registry) error {

	if m.Related != nil {

		if swag.IsZero(m.Related) { // not required
			return nil
		}

		if err := m.Related.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("related")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("related")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RelatedLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelatedLink) UnmarshalBinary(b []byte) error {
	var res RelatedLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
