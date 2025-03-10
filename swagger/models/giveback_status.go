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

// GivebackStatus giveback status
//
// swagger:model giveback_status
type GivebackStatus struct {

	// aggregate
	Aggregate *GivebackStatusInlineAggregate `json:"aggregate,omitempty"`

	// error
	Error *GivebackStatusInlineError `json:"error,omitempty"`

	// Giveback state of the aggregate. <br/>
	// Possible values include no aggregates to giveback(nothing_to_giveback), failed to disable background disk firmware update(BDFU) on source node(failed_bdfu_source), <br/>
	// giveback delayed as disk firmware update is in progress on source node(delayed_bdfu_source), performing veto checks(running_checks). <br/>
	//
	// Enum: ["done","failed","in_progress","not_started","nothing_to_giveback","failed_bdfu_source","failed_bdfu_dest","delayed_bdfu_source","delayed_bdfu_dest","running_checks"]
	State *string `json:"state,omitempty"`
}

// Validate validates this giveback status
func (m *GivebackStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GivebackStatus) validateAggregate(formats strfmt.Registry) error {
	if swag.IsZero(m.Aggregate) { // not required
		return nil
	}

	if m.Aggregate != nil {
		if err := m.Aggregate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aggregate")
			}
			return err
		}
	}

	return nil
}

func (m *GivebackStatus) validateError(formats strfmt.Registry) error {
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

var givebackStatusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["done","failed","in_progress","not_started","nothing_to_giveback","failed_bdfu_source","failed_bdfu_dest","delayed_bdfu_source","delayed_bdfu_dest","running_checks"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		givebackStatusTypeStatePropEnum = append(givebackStatusTypeStatePropEnum, v)
	}
}

const (

	// GivebackStatusStateDone captures enum value "done"
	GivebackStatusStateDone string = "done"

	// GivebackStatusStateFailed captures enum value "failed"
	GivebackStatusStateFailed string = "failed"

	// GivebackStatusStateInProgress captures enum value "in_progress"
	GivebackStatusStateInProgress string = "in_progress"

	// GivebackStatusStateNotStarted captures enum value "not_started"
	GivebackStatusStateNotStarted string = "not_started"

	// GivebackStatusStateNothingToGiveback captures enum value "nothing_to_giveback"
	GivebackStatusStateNothingToGiveback string = "nothing_to_giveback"

	// GivebackStatusStateFailedBdfuSource captures enum value "failed_bdfu_source"
	GivebackStatusStateFailedBdfuSource string = "failed_bdfu_source"

	// GivebackStatusStateFailedBdfuDest captures enum value "failed_bdfu_dest"
	GivebackStatusStateFailedBdfuDest string = "failed_bdfu_dest"

	// GivebackStatusStateDelayedBdfuSource captures enum value "delayed_bdfu_source"
	GivebackStatusStateDelayedBdfuSource string = "delayed_bdfu_source"

	// GivebackStatusStateDelayedBdfuDest captures enum value "delayed_bdfu_dest"
	GivebackStatusStateDelayedBdfuDest string = "delayed_bdfu_dest"

	// GivebackStatusStateRunningChecks captures enum value "running_checks"
	GivebackStatusStateRunningChecks string = "running_checks"
)

// prop value enum
func (m *GivebackStatus) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, givebackStatusTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GivebackStatus) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this giveback status based on the context it is used
func (m *GivebackStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GivebackStatus) contextValidateAggregate(ctx context.Context, formats strfmt.Registry) error {

	if m.Aggregate != nil {

		if swag.IsZero(m.Aggregate) { // not required
			return nil
		}

		if err := m.Aggregate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aggregate")
			}
			return err
		}
	}

	return nil
}

func (m *GivebackStatus) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GivebackStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GivebackStatus) UnmarshalBinary(b []byte) error {
	var res GivebackStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GivebackStatusInlineAggregate Aggregate name and UUID.
//
// swagger:model giveback_status_inline_aggregate
type GivebackStatusInlineAggregate struct {

	// links
	Links *GivebackStatusInlineAggregateInlineLinks `json:"_links,omitempty"`

	// name
	// Example: aggr1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this giveback status inline aggregate
func (m *GivebackStatusInlineAggregate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GivebackStatusInlineAggregate) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aggregate" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this giveback status inline aggregate based on the context it is used
func (m *GivebackStatusInlineAggregate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GivebackStatusInlineAggregate) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aggregate" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GivebackStatusInlineAggregate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GivebackStatusInlineAggregate) UnmarshalBinary(b []byte) error {
	var res GivebackStatusInlineAggregate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GivebackStatusInlineAggregateInlineLinks giveback status inline aggregate inline links
//
// swagger:model giveback_status_inline_aggregate_inline__links
type GivebackStatusInlineAggregateInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this giveback status inline aggregate inline links
func (m *GivebackStatusInlineAggregateInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GivebackStatusInlineAggregateInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aggregate" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this giveback status inline aggregate inline links based on the context it is used
func (m *GivebackStatusInlineAggregateInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GivebackStatusInlineAggregateInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aggregate" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GivebackStatusInlineAggregateInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GivebackStatusInlineAggregateInlineLinks) UnmarshalBinary(b []byte) error {
	var res GivebackStatusInlineAggregateInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GivebackStatusInlineError Indicates the failed aggregate giveback code and message.
//
// swagger:model giveback_status_inline_error
type GivebackStatusInlineError struct {

	// Message code.
	// Example: 852126
	// Read Only: true
	Code *string `json:"code,omitempty"`

	// Detailed message based on the state.
	// Read Only: true
	// Enum: ["shutdown","not_homes_partner","not_sfo","failed_limbo","offline_failed","migrating","veto","communication_err","online_timeout","online_failed","hdd_to_aff_dest"]
	Message *string `json:"message,omitempty"`
}

// Validate validates this giveback status inline error
func (m *GivebackStatusInlineError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var givebackStatusInlineErrorTypeMessagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["shutdown","not_homes_partner","not_sfo","failed_limbo","offline_failed","migrating","veto","communication_err","online_timeout","online_failed","hdd_to_aff_dest"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		givebackStatusInlineErrorTypeMessagePropEnum = append(givebackStatusInlineErrorTypeMessagePropEnum, v)
	}
}

const (

	// GivebackStatusInlineErrorMessageShutdown captures enum value "shutdown"
	GivebackStatusInlineErrorMessageShutdown string = "shutdown"

	// GivebackStatusInlineErrorMessageNotHomesPartner captures enum value "not_homes_partner"
	GivebackStatusInlineErrorMessageNotHomesPartner string = "not_homes_partner"

	// GivebackStatusInlineErrorMessageNotSfo captures enum value "not_sfo"
	GivebackStatusInlineErrorMessageNotSfo string = "not_sfo"

	// GivebackStatusInlineErrorMessageFailedLimbo captures enum value "failed_limbo"
	GivebackStatusInlineErrorMessageFailedLimbo string = "failed_limbo"

	// GivebackStatusInlineErrorMessageOfflineFailed captures enum value "offline_failed"
	GivebackStatusInlineErrorMessageOfflineFailed string = "offline_failed"

	// GivebackStatusInlineErrorMessageMigrating captures enum value "migrating"
	GivebackStatusInlineErrorMessageMigrating string = "migrating"

	// GivebackStatusInlineErrorMessageVeto captures enum value "veto"
	GivebackStatusInlineErrorMessageVeto string = "veto"

	// GivebackStatusInlineErrorMessageCommunicationErr captures enum value "communication_err"
	GivebackStatusInlineErrorMessageCommunicationErr string = "communication_err"

	// GivebackStatusInlineErrorMessageOnlineTimeout captures enum value "online_timeout"
	GivebackStatusInlineErrorMessageOnlineTimeout string = "online_timeout"

	// GivebackStatusInlineErrorMessageOnlineFailed captures enum value "online_failed"
	GivebackStatusInlineErrorMessageOnlineFailed string = "online_failed"

	// GivebackStatusInlineErrorMessageHddToAffDest captures enum value "hdd_to_aff_dest"
	GivebackStatusInlineErrorMessageHddToAffDest string = "hdd_to_aff_dest"
)

// prop value enum
func (m *GivebackStatusInlineError) validateMessageEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, givebackStatusInlineErrorTypeMessagePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GivebackStatusInlineError) validateMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.Message) { // not required
		return nil
	}

	// value enum
	if err := m.validateMessageEnum("error"+"."+"message", "body", *m.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this giveback status inline error based on the context it is used
func (m *GivebackStatusInlineError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GivebackStatusInlineError) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "error"+"."+"code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *GivebackStatusInlineError) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "error"+"."+"message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GivebackStatusInlineError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GivebackStatusInlineError) UnmarshalBinary(b []byte) error {
	var res GivebackStatusInlineError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
