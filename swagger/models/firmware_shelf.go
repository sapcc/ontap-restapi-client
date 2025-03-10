// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FirmwareShelf firmware shelf
//
// swagger:model firmware_shelf
type FirmwareShelf struct {

	// in progress count
	// Example: 2
	// Read Only: true
	InProgressCount *int64 `json:"in_progress_count,omitempty"`

	// Status of the shelf firmware update.
	// Read Only: true
	// Enum: ["running","idle"]
	UpdateStatus interface{} `json:"update_status,omitempty"`
}

// Validate validates this firmware shelf
func (m *FirmwareShelf) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this firmware shelf based on the context it is used
func (m *FirmwareShelf) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInProgressCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareShelf) contextValidateInProgressCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "in_progress_count", "body", m.InProgressCount); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FirmwareShelf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FirmwareShelf) UnmarshalBinary(b []byte) error {
	var res FirmwareShelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
