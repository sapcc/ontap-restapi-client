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

// HwAssist The hardware assist information.
//
// swagger:model hw_assist
type HwAssist struct {

	// status
	Status *HwAssistInlineStatus `json:"status,omitempty"`
}

// Validate validates this hw assist
func (m *HwAssist) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HwAssist) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hw assist based on the context it is used
func (m *HwAssist) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HwAssist) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HwAssist) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HwAssist) UnmarshalBinary(b []byte) error {
	var res HwAssist
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HwAssistInlineStatus hw assist inline status
//
// swagger:model hw_assist_inline_status
type HwAssistInlineStatus struct {

	// Indicates whether hardware assist is enabled on the node.
	Enabled *bool `json:"enabled,omitempty"`

	// local
	Local *HwAssistInlineStatusInlineLocal `json:"local,omitempty"`

	// partner
	Partner *HwAssistInlineStatusInlinePartner `json:"partner,omitempty"`
}

// Validate validates this hw assist inline status
func (m *HwAssistInlineStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HwAssistInlineStatus) validateLocal(formats strfmt.Registry) error {
	if swag.IsZero(m.Local) { // not required
		return nil
	}

	if m.Local != nil {
		if err := m.Local.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "local")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status" + "." + "local")
			}
			return err
		}
	}

	return nil
}

func (m *HwAssistInlineStatus) validatePartner(formats strfmt.Registry) error {
	if swag.IsZero(m.Partner) { // not required
		return nil
	}

	if m.Partner != nil {
		if err := m.Partner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "partner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status" + "." + "partner")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hw assist inline status based on the context it is used
func (m *HwAssistInlineStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePartner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HwAssistInlineStatus) contextValidateLocal(ctx context.Context, formats strfmt.Registry) error {

	if m.Local != nil {

		if swag.IsZero(m.Local) { // not required
			return nil
		}

		if err := m.Local.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "local")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status" + "." + "local")
			}
			return err
		}
	}

	return nil
}

func (m *HwAssistInlineStatus) contextValidatePartner(ctx context.Context, formats strfmt.Registry) error {

	if m.Partner != nil {

		if swag.IsZero(m.Partner) { // not required
			return nil
		}

		if err := m.Partner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "partner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status" + "." + "partner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HwAssistInlineStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HwAssistInlineStatus) UnmarshalBinary(b []byte) error {
	var res HwAssistInlineStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HwAssistInlineStatusInlineLocal hw assist inline status inline local
//
// swagger:model hw_assist_inline_status_inline_local
type HwAssistInlineStatusInlineLocal struct {

	// The hardware assist IP address.
	IP *string `json:"ip,omitempty"`

	// The hardware assist port.
	Port *int64 `json:"port,omitempty"`

	// The hardware assist monitor status.
	// Enum: ["active","inactive"]
	State *string `json:"state,omitempty"`
}

// Validate validates this hw assist inline status inline local
func (m *HwAssistInlineStatusInlineLocal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var hwAssistInlineStatusInlineLocalTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hwAssistInlineStatusInlineLocalTypeStatePropEnum = append(hwAssistInlineStatusInlineLocalTypeStatePropEnum, v)
	}
}

const (

	// HwAssistInlineStatusInlineLocalStateActive captures enum value "active"
	HwAssistInlineStatusInlineLocalStateActive string = "active"

	// HwAssistInlineStatusInlineLocalStateInactive captures enum value "inactive"
	HwAssistInlineStatusInlineLocalStateInactive string = "inactive"
)

// prop value enum
func (m *HwAssistInlineStatusInlineLocal) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, hwAssistInlineStatusInlineLocalTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HwAssistInlineStatusInlineLocal) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("status"+"."+"local"+"."+"state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this hw assist inline status inline local based on context it is used
func (m *HwAssistInlineStatusInlineLocal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HwAssistInlineStatusInlineLocal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HwAssistInlineStatusInlineLocal) UnmarshalBinary(b []byte) error {
	var res HwAssistInlineStatusInlineLocal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HwAssistInlineStatusInlinePartner hw assist inline status inline partner
//
// swagger:model hw_assist_inline_status_inline_partner
type HwAssistInlineStatusInlinePartner struct {

	// The hardware assist IP address.
	IP *string `json:"ip,omitempty"`

	// The hardware assist port.
	Port *int64 `json:"port,omitempty"`

	// The hardware assist monitor status.
	// Enum: ["active","inactive"]
	State *string `json:"state,omitempty"`
}

// Validate validates this hw assist inline status inline partner
func (m *HwAssistInlineStatusInlinePartner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var hwAssistInlineStatusInlinePartnerTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hwAssistInlineStatusInlinePartnerTypeStatePropEnum = append(hwAssistInlineStatusInlinePartnerTypeStatePropEnum, v)
	}
}

const (

	// HwAssistInlineStatusInlinePartnerStateActive captures enum value "active"
	HwAssistInlineStatusInlinePartnerStateActive string = "active"

	// HwAssistInlineStatusInlinePartnerStateInactive captures enum value "inactive"
	HwAssistInlineStatusInlinePartnerStateInactive string = "inactive"
)

// prop value enum
func (m *HwAssistInlineStatusInlinePartner) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, hwAssistInlineStatusInlinePartnerTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HwAssistInlineStatusInlinePartner) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("status"+"."+"partner"+"."+"state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this hw assist inline status inline partner based on context it is used
func (m *HwAssistInlineStatusInlinePartner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HwAssistInlineStatusInlinePartner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HwAssistInlineStatusInlinePartner) UnmarshalBinary(b []byte) error {
	var res HwAssistInlineStatusInlinePartner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
