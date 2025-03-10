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

// ArwVserver arw vserver
//
// swagger:model arw_vserver
type ArwVserver struct {

	// event log
	EventLog *ArwVserverInlineEventLog `json:"event_log,omitempty"`
}

// Validate validates this arw vserver
func (m *ArwVserver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventLog(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArwVserver) validateEventLog(formats strfmt.Registry) error {
	if swag.IsZero(m.EventLog) { // not required
		return nil
	}

	if m.EventLog != nil {
		if err := m.EventLog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_log")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_log")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this arw vserver based on the context it is used
func (m *ArwVserver) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventLog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArwVserver) contextValidateEventLog(ctx context.Context, formats strfmt.Registry) error {

	if m.EventLog != nil {

		if swag.IsZero(m.EventLog) { // not required
			return nil
		}

		if err := m.EventLog.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_log")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_log")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ArwVserver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArwVserver) UnmarshalBinary(b []byte) error {
	var res ArwVserver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ArwVserverInlineEventLog arw vserver inline event log
//
// swagger:model arw_vserver_inline_event_log
type ArwVserverInlineEventLog struct {

	// Specifies whether to send an EMS when a new file extension is discovered.
	IsEnabledOnNewFileExtensionSeen *bool `json:"is_enabled_on_new_file_extension_seen,omitempty"`

	// Specifies whether to send an EMS when a snapshot is created.
	IsEnabledOnSnapshotCopyCreation *bool `json:"is_enabled_on_snapshot_copy_creation,omitempty"`
}

// Validate validates this arw vserver inline event log
func (m *ArwVserverInlineEventLog) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this arw vserver inline event log based on context it is used
func (m *ArwVserverInlineEventLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ArwVserverInlineEventLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArwVserverInlineEventLog) UnmarshalBinary(b []byte) error {
	var res ArwVserverInlineEventLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
