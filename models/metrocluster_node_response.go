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
)

// MetroclusterNodeResponse metrocluster node response
//
// swagger:model metrocluster_node_response
type MetroclusterNodeResponse struct {

	// links
	Links *CollectionLinks `json:"_links,omitempty"`

	// metrocluster node response inline records
	MetroclusterNodeResponseInlineRecords []*MetroclusterNode `json:"records,omitempty"`

	// Number of Records
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty"`
}

// Validate validates this metrocluster node response
func (m *MetroclusterNodeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetroclusterNodeResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterNodeResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *MetroclusterNodeResponse) validateMetroclusterNodeResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.MetroclusterNodeResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.MetroclusterNodeResponseInlineRecords); i++ {
		if swag.IsZero(m.MetroclusterNodeResponseInlineRecords[i]) { // not required
			continue
		}

		if m.MetroclusterNodeResponseInlineRecords[i] != nil {
			if err := m.MetroclusterNodeResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this metrocluster node response based on the context it is used
func (m *MetroclusterNodeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetroclusterNodeResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterNodeResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MetroclusterNodeResponse) contextValidateMetroclusterNodeResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MetroclusterNodeResponseInlineRecords); i++ {

		if m.MetroclusterNodeResponseInlineRecords[i] != nil {

			if swag.IsZero(m.MetroclusterNodeResponseInlineRecords[i]) { // not required
				return nil
			}

			if err := m.MetroclusterNodeResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterNodeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterNodeResponse) UnmarshalBinary(b []byte) error {
	var res MetroclusterNodeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
