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

// Sensors Environment Sensors
//
// swagger:model sensors
type Sensors struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// Value above which the sensor goes into a critically high state.
	// Read Only: true
	CriticalHighThreshold *int64 `json:"critical_high_threshold,omitempty"`

	// Value below which the sensor goes into a critically low state.
	// Read Only: true
	CriticalLowThreshold *int64 `json:"critical_low_threshold,omitempty"`

	// Used to determine whether the sensor is in a normal state or any other failed state based on the value of "discrete_value" field. This field is only applicable for discrete sensors.
	// Example: normal
	// Read Only: true
	// Enum: ["bad","crit_high","crit_low","disabled","failed","fault","ignored","init_failed","invalid","normal","not_available","not_present","retry","uninitialized","unknown","warn_high","warn_low"]
	DiscreteState *string `json:"discrete_state,omitempty"`

	// Applies to discrete sensors which do not have an integer value. It can have values like on, off, good, bad, ok.
	// Example: ok
	// Read Only: true
	DiscreteValue *string `json:"discrete_value,omitempty"`

	// Provides the sensor ID.
	// Read Only: true
	Index *int64 `json:"index,omitempty"`

	// Name of the sensor.
	// Example: PVCCSA CPU FD
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// node
	Node *SensorsInlineNode `json:"node,omitempty"`

	// Used to determine whether the sensor is in a normal state or any other failed state.
	// Example: normal
	// Read Only: true
	// Enum: ["bad","crit_high","crit_low","disabled","failed","fault","ignored","init_failed","invalid","normal","not_available","not_present","retry","uninitialized","unknown","warn_high","warn_low"]
	ThresholdState *string `json:"threshold_state,omitempty"`

	// Used to determine the type of the sensor.
	// Read Only: true
	// Enum: ["agent","battery_life","counter","current","discrete","fan","fru","minutes","nvmem","percent","thermal","unknown","voltage"]
	Type *string `json:"type,omitempty"`

	// Provides the sensor reading.
	// Example: 831
	// Read Only: true
	Value *int32 `json:"value,omitempty"`

	// Units in which the "value" is measured. Some examples of units are mV, mW*hr, C, RPM.
	// Example: mV
	// Read Only: true
	ValueUnits *string `json:"value_units,omitempty"`

	// Value above which the sensor goes into a warning high state.
	// Read Only: true
	WarningHighThreshold *int64 `json:"warning_high_threshold,omitempty"`

	// Value below which the sensor goes into a warning low state.
	// Read Only: true
	WarningLowThreshold *int64 `json:"warning_low_threshold,omitempty"`
}

// Validate validates this sensors
func (m *Sensors) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscreteState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThresholdState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Sensors) validateLinks(formats strfmt.Registry) error {
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

var sensorsTypeDiscreteStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bad","crit_high","crit_low","disabled","failed","fault","ignored","init_failed","invalid","normal","not_available","not_present","retry","uninitialized","unknown","warn_high","warn_low"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sensorsTypeDiscreteStatePropEnum = append(sensorsTypeDiscreteStatePropEnum, v)
	}
}

const (

	// SensorsDiscreteStateBad captures enum value "bad"
	SensorsDiscreteStateBad string = "bad"

	// SensorsDiscreteStateCritHigh captures enum value "crit_high"
	SensorsDiscreteStateCritHigh string = "crit_high"

	// SensorsDiscreteStateCritLow captures enum value "crit_low"
	SensorsDiscreteStateCritLow string = "crit_low"

	// SensorsDiscreteStateDisabled captures enum value "disabled"
	SensorsDiscreteStateDisabled string = "disabled"

	// SensorsDiscreteStateFailed captures enum value "failed"
	SensorsDiscreteStateFailed string = "failed"

	// SensorsDiscreteStateFault captures enum value "fault"
	SensorsDiscreteStateFault string = "fault"

	// SensorsDiscreteStateIgnored captures enum value "ignored"
	SensorsDiscreteStateIgnored string = "ignored"

	// SensorsDiscreteStateInitFailed captures enum value "init_failed"
	SensorsDiscreteStateInitFailed string = "init_failed"

	// SensorsDiscreteStateInvalid captures enum value "invalid"
	SensorsDiscreteStateInvalid string = "invalid"

	// SensorsDiscreteStateNormal captures enum value "normal"
	SensorsDiscreteStateNormal string = "normal"

	// SensorsDiscreteStateNotAvailable captures enum value "not_available"
	SensorsDiscreteStateNotAvailable string = "not_available"

	// SensorsDiscreteStateNotPresent captures enum value "not_present"
	SensorsDiscreteStateNotPresent string = "not_present"

	// SensorsDiscreteStateRetry captures enum value "retry"
	SensorsDiscreteStateRetry string = "retry"

	// SensorsDiscreteStateUninitialized captures enum value "uninitialized"
	SensorsDiscreteStateUninitialized string = "uninitialized"

	// SensorsDiscreteStateUnknown captures enum value "unknown"
	SensorsDiscreteStateUnknown string = "unknown"

	// SensorsDiscreteStateWarnHigh captures enum value "warn_high"
	SensorsDiscreteStateWarnHigh string = "warn_high"

	// SensorsDiscreteStateWarnLow captures enum value "warn_low"
	SensorsDiscreteStateWarnLow string = "warn_low"
)

// prop value enum
func (m *Sensors) validateDiscreteStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sensorsTypeDiscreteStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Sensors) validateDiscreteState(formats strfmt.Registry) error {
	if swag.IsZero(m.DiscreteState) { // not required
		return nil
	}

	// value enum
	if err := m.validateDiscreteStateEnum("discrete_state", "body", *m.DiscreteState); err != nil {
		return err
	}

	return nil
}

func (m *Sensors) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

var sensorsTypeThresholdStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bad","crit_high","crit_low","disabled","failed","fault","ignored","init_failed","invalid","normal","not_available","not_present","retry","uninitialized","unknown","warn_high","warn_low"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sensorsTypeThresholdStatePropEnum = append(sensorsTypeThresholdStatePropEnum, v)
	}
}

const (

	// SensorsThresholdStateBad captures enum value "bad"
	SensorsThresholdStateBad string = "bad"

	// SensorsThresholdStateCritHigh captures enum value "crit_high"
	SensorsThresholdStateCritHigh string = "crit_high"

	// SensorsThresholdStateCritLow captures enum value "crit_low"
	SensorsThresholdStateCritLow string = "crit_low"

	// SensorsThresholdStateDisabled captures enum value "disabled"
	SensorsThresholdStateDisabled string = "disabled"

	// SensorsThresholdStateFailed captures enum value "failed"
	SensorsThresholdStateFailed string = "failed"

	// SensorsThresholdStateFault captures enum value "fault"
	SensorsThresholdStateFault string = "fault"

	// SensorsThresholdStateIgnored captures enum value "ignored"
	SensorsThresholdStateIgnored string = "ignored"

	// SensorsThresholdStateInitFailed captures enum value "init_failed"
	SensorsThresholdStateInitFailed string = "init_failed"

	// SensorsThresholdStateInvalid captures enum value "invalid"
	SensorsThresholdStateInvalid string = "invalid"

	// SensorsThresholdStateNormal captures enum value "normal"
	SensorsThresholdStateNormal string = "normal"

	// SensorsThresholdStateNotAvailable captures enum value "not_available"
	SensorsThresholdStateNotAvailable string = "not_available"

	// SensorsThresholdStateNotPresent captures enum value "not_present"
	SensorsThresholdStateNotPresent string = "not_present"

	// SensorsThresholdStateRetry captures enum value "retry"
	SensorsThresholdStateRetry string = "retry"

	// SensorsThresholdStateUninitialized captures enum value "uninitialized"
	SensorsThresholdStateUninitialized string = "uninitialized"

	// SensorsThresholdStateUnknown captures enum value "unknown"
	SensorsThresholdStateUnknown string = "unknown"

	// SensorsThresholdStateWarnHigh captures enum value "warn_high"
	SensorsThresholdStateWarnHigh string = "warn_high"

	// SensorsThresholdStateWarnLow captures enum value "warn_low"
	SensorsThresholdStateWarnLow string = "warn_low"
)

// prop value enum
func (m *Sensors) validateThresholdStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sensorsTypeThresholdStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Sensors) validateThresholdState(formats strfmt.Registry) error {
	if swag.IsZero(m.ThresholdState) { // not required
		return nil
	}

	// value enum
	if err := m.validateThresholdStateEnum("threshold_state", "body", *m.ThresholdState); err != nil {
		return err
	}

	return nil
}

var sensorsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["agent","battery_life","counter","current","discrete","fan","fru","minutes","nvmem","percent","thermal","unknown","voltage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sensorsTypeTypePropEnum = append(sensorsTypeTypePropEnum, v)
	}
}

const (

	// SensorsTypeAgent captures enum value "agent"
	SensorsTypeAgent string = "agent"

	// SensorsTypeBatteryLife captures enum value "battery_life"
	SensorsTypeBatteryLife string = "battery_life"

	// SensorsTypeCounter captures enum value "counter"
	SensorsTypeCounter string = "counter"

	// SensorsTypeCurrent captures enum value "current"
	SensorsTypeCurrent string = "current"

	// SensorsTypeDiscrete captures enum value "discrete"
	SensorsTypeDiscrete string = "discrete"

	// SensorsTypeFan captures enum value "fan"
	SensorsTypeFan string = "fan"

	// SensorsTypeFru captures enum value "fru"
	SensorsTypeFru string = "fru"

	// SensorsTypeMinutes captures enum value "minutes"
	SensorsTypeMinutes string = "minutes"

	// SensorsTypeNvmem captures enum value "nvmem"
	SensorsTypeNvmem string = "nvmem"

	// SensorsTypePercent captures enum value "percent"
	SensorsTypePercent string = "percent"

	// SensorsTypeThermal captures enum value "thermal"
	SensorsTypeThermal string = "thermal"

	// SensorsTypeUnknown captures enum value "unknown"
	SensorsTypeUnknown string = "unknown"

	// SensorsTypeVoltage captures enum value "voltage"
	SensorsTypeVoltage string = "voltage"
)

// prop value enum
func (m *Sensors) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sensorsTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Sensors) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this sensors based on the context it is used
func (m *Sensors) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCriticalHighThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCriticalLowThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDiscreteState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDiscreteValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIndex(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThresholdState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValueUnits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWarningHighThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWarningLowThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Sensors) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Sensors) contextValidateCriticalHighThreshold(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "critical_high_threshold", "body", m.CriticalHighThreshold); err != nil {
		return err
	}

	return nil
}

func (m *Sensors) contextValidateCriticalLowThreshold(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "critical_low_threshold", "body", m.CriticalLowThreshold); err != nil {
		return err
	}

	return nil
}

func (m *Sensors) contextValidateDiscreteState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "discrete_state", "body", m.DiscreteState); err != nil {
		return err
	}

	return nil
}

func (m *Sensors) contextValidateDiscreteValue(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "discrete_value", "body", m.DiscreteValue); err != nil {
		return err
	}

	return nil
}

func (m *Sensors) contextValidateIndex(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *Sensors) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Sensors) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {

		if swag.IsZero(m.Node) { // not required
			return nil
		}

		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *Sensors) contextValidateThresholdState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "threshold_state", "body", m.ThresholdState); err != nil {
		return err
	}

	return nil
}

func (m *Sensors) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Sensors) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

func (m *Sensors) contextValidateValueUnits(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "value_units", "body", m.ValueUnits); err != nil {
		return err
	}

	return nil
}

func (m *Sensors) contextValidateWarningHighThreshold(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "warning_high_threshold", "body", m.WarningHighThreshold); err != nil {
		return err
	}

	return nil
}

func (m *Sensors) contextValidateWarningLowThreshold(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "warning_low_threshold", "body", m.WarningLowThreshold); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Sensors) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Sensors) UnmarshalBinary(b []byte) error {
	var res Sensors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SensorsInlineNode sensors inline node
//
// swagger:model sensors_inline_node
type SensorsInlineNode struct {

	// links
	Links *SensorsInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this sensors inline node
func (m *SensorsInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SensorsInlineNode) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sensors inline node based on the context it is used
func (m *SensorsInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SensorsInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SensorsInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SensorsInlineNode) UnmarshalBinary(b []byte) error {
	var res SensorsInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SensorsInlineNodeInlineLinks sensors inline node inline links
//
// swagger:model sensors_inline_node_inline__links
type SensorsInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this sensors inline node inline links
func (m *SensorsInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SensorsInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sensors inline node inline links based on the context it is used
func (m *SensorsInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SensorsInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SensorsInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SensorsInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res SensorsInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
