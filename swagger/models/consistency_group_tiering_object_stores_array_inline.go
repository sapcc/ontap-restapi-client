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

// ConsistencyGroupTieringObjectStoresArrayInline Object stores to use. Used for placement.
//
// swagger:model consistency_group_tiering_object_stores_array_inline
type ConsistencyGroupTieringObjectStoresArrayInline []*ConsistencyGroupTieringObjectStoresArrayInlineInlineArrayItem

// Validate validates this consistency group tiering object stores array inline
func (m ConsistencyGroupTieringObjectStoresArrayInline) Validate(formats strfmt.Registry) error {
	var res []error

	iConsistencyGroupTieringObjectStoresArrayInlineSize := int64(len(m))

	if err := validate.MinItems("", "body", iConsistencyGroupTieringObjectStoresArrayInlineSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("", "body", iConsistencyGroupTieringObjectStoresArrayInlineSize, 2); err != nil {
		return err
	}

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

// ContextValidate validate this consistency group tiering object stores array inline based on the context it is used
func (m ConsistencyGroupTieringObjectStoresArrayInline) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

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

// ConsistencyGroupTieringObjectStoresArrayInlineInlineArrayItem consistency group tiering object stores array inline inline array item
//
// swagger:model consistency_group_tiering_object_stores_array_inline_inline_array_item
type ConsistencyGroupTieringObjectStoresArrayInlineInlineArrayItem struct {

	// The name of the object store to use. Used for placement.
	Name *string `json:"name,omitempty"`
}

// Validate validates this consistency group tiering object stores array inline inline array item
func (m *ConsistencyGroupTieringObjectStoresArrayInlineInlineArrayItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this consistency group tiering object stores array inline inline array item based on context it is used
func (m *ConsistencyGroupTieringObjectStoresArrayInlineInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupTieringObjectStoresArrayInlineInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupTieringObjectStoresArrayInlineInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupTieringObjectStoresArrayInlineInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
