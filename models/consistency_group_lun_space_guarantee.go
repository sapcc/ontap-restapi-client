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

// ConsistencyGroupLunSpaceGuarantee Properties that request and report the space guarantee for the LUN.
//
// swagger:model consistency_group_lun_space_guarantee
type ConsistencyGroupLunSpaceGuarantee struct {

	// The requested space reservation policy for the LUN. If _true_, a space reservation is requested for the LUN; if _false_, the LUN is thin provisioned. Guaranteeing a space reservation request for a LUN requires that the volume in which the LUN resides is also space reserved and that the fractional reserve for the volume is 100%. Valid in POST and PATCH.
	//
	//
	Requested *bool `json:"requested,omitempty"`

	// Reports if the LUN is space guaranteed.<br/>
	// If _true_, a space guarantee is requested and the containing volume and aggregate support the request. If _false_, a space guarantee is not requested or a space guarantee is requested and either the containing volume or aggregate do not support the request.
	//
	// Read Only: true
	Reserved *bool `json:"reserved,omitempty"`
}

// Validate validates this consistency group lun space guarantee
func (m *ConsistencyGroupLunSpaceGuarantee) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this consistency group lun space guarantee based on the context it is used
func (m *ConsistencyGroupLunSpaceGuarantee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReserved(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupLunSpaceGuarantee) contextValidateReserved(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "reserved", "body", m.Reserved); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupLunSpaceGuarantee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupLunSpaceGuarantee) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupLunSpaceGuarantee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
