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

// ErrorResponse error response
//
// swagger:model error_response
type ErrorResponse struct {

	// error
	Error *ReturnedError `json:"error,omitempty"`
}

// Validate validates this error response
func (m *ErrorResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorResponse) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this error response based on the context it is used
func (m *ErrorResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorResponse) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {

		if swag.IsZero(m.Error) { // not required
			return nil
		}

		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ErrorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorResponse) UnmarshalBinary(b []byte) error {
	var res ErrorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
