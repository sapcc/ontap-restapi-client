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

// NetworkRouteForSvm network route for svm
//
// swagger:model network_route_for_svm
type NetworkRouteForSvm struct {

	// destination
	Destination *IPInfo `json:"destination,omitempty"`

	// The IP address of the gateway router leading to the destination.
	// Example: 10.1.1.1
	Gateway *string `json:"gateway,omitempty"`
}

// Validate validates this network route for svm
func (m *NetworkRouteForSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkRouteForSvm) validateDestination(formats strfmt.Registry) error {
	if swag.IsZero(m.Destination) { // not required
		return nil
	}

	if m.Destination != nil {
		if err := m.Destination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network route for svm based on the context it is used
func (m *NetworkRouteForSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkRouteForSvm) contextValidateDestination(ctx context.Context, formats strfmt.Registry) error {

	if m.Destination != nil {

		if swag.IsZero(m.Destination) { // not required
			return nil
		}

		if err := m.Destination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkRouteForSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkRouteForSvm) UnmarshalBinary(b []byte) error {
	var res NetworkRouteForSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
