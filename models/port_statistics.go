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

// PortStatistics These are raw performance and error counters for the Ethernet port.
//
// swagger:model port_statistics
type PortStatistics struct {

	// device
	Device *PortStatisticsInlineDevice `json:"device,omitempty"`

	// Errors associated with the sample. For example, if the aggregation of data over multiple nodes fails, then any partial errors might return "ok" on success or "error" on an internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "inconsistent_delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Enum: ["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]
	Status *string `json:"status,omitempty"`

	// throughput raw
	ThroughputRaw *PortStatisticsInlineThroughputRaw `json:"throughput_raw,omitempty"`

	// The timestamp of the throughput_raw performance data.
	// Example: 2017-01-25 11:20:13
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this port statistics
func (m *PortStatistics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThroughputRaw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortStatistics) validateDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.Device) { // not required
		return nil
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

var portStatisticsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		portStatisticsTypeStatusPropEnum = append(portStatisticsTypeStatusPropEnum, v)
	}
}

const (

	// PortStatisticsStatusOk captures enum value "ok"
	PortStatisticsStatusOk string = "ok"

	// PortStatisticsStatusError captures enum value "error"
	PortStatisticsStatusError string = "error"

	// PortStatisticsStatusPartialNoData captures enum value "partial_no_data"
	PortStatisticsStatusPartialNoData string = "partial_no_data"

	// PortStatisticsStatusPartialNoUUID captures enum value "partial_no_uuid"
	PortStatisticsStatusPartialNoUUID string = "partial_no_uuid"

	// PortStatisticsStatusPartialNoResponse captures enum value "partial_no_response"
	PortStatisticsStatusPartialNoResponse string = "partial_no_response"

	// PortStatisticsStatusPartialOtherError captures enum value "partial_other_error"
	PortStatisticsStatusPartialOtherError string = "partial_other_error"

	// PortStatisticsStatusNegativeDelta captures enum value "negative_delta"
	PortStatisticsStatusNegativeDelta string = "negative_delta"

	// PortStatisticsStatusBackfilledData captures enum value "backfilled_data"
	PortStatisticsStatusBackfilledData string = "backfilled_data"

	// PortStatisticsStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PortStatisticsStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// PortStatisticsStatusInconsistentOldData captures enum value "inconsistent_old_data"
	PortStatisticsStatusInconsistentOldData string = "inconsistent_old_data"
)

// prop value enum
func (m *PortStatistics) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, portStatisticsTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PortStatistics) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PortStatistics) validateThroughputRaw(formats strfmt.Registry) error {
	if swag.IsZero(m.ThroughputRaw) { // not required
		return nil
	}

	if m.ThroughputRaw != nil {
		if err := m.ThroughputRaw.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput_raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throughput_raw")
			}
			return err
		}
	}

	return nil
}

func (m *PortStatistics) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this port statistics based on the context it is used
func (m *PortStatistics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughputRaw(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortStatistics) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {

		if swag.IsZero(m.Device) { // not required
			return nil
		}

		if err := m.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *PortStatistics) contextValidateThroughputRaw(ctx context.Context, formats strfmt.Registry) error {

	if m.ThroughputRaw != nil {

		if swag.IsZero(m.ThroughputRaw) { // not required
			return nil
		}

		if err := m.ThroughputRaw.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput_raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throughput_raw")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortStatistics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortStatistics) UnmarshalBinary(b []byte) error {
	var res PortStatistics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortStatisticsInlineDevice Device-related counters for the port object. These counters are applicable at the lowest layer of the networking stack. These values can be used to calculate both transmit and receive packet and error rates by comparing two samples taken at different times and calculating the increase in counter value divided by the elapsed time between the two samples.
//
// swagger:model port_statistics_inline_device
type PortStatisticsInlineDevice struct {

	// The number of link state changes from up to down seen on the device.
	// Example: 3
	LinkDownCountRaw *int64 `json:"link_down_count_raw,omitempty"`

	// receive raw
	ReceiveRaw *PortStatisticsInlineDeviceInlineReceiveRaw `json:"receive_raw,omitempty"`

	// The timestamp when the device specific counters were collected.
	// Example: 2017-01-25 11:20:13
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// transmit raw
	TransmitRaw *PortStatisticsInlineDeviceInlineTransmitRaw `json:"transmit_raw,omitempty"`
}

// Validate validates this port statistics inline device
func (m *PortStatisticsInlineDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReceiveRaw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransmitRaw(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortStatisticsInlineDevice) validateReceiveRaw(formats strfmt.Registry) error {
	if swag.IsZero(m.ReceiveRaw) { // not required
		return nil
	}

	if m.ReceiveRaw != nil {
		if err := m.ReceiveRaw.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device" + "." + "receive_raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device" + "." + "receive_raw")
			}
			return err
		}
	}

	return nil
}

func (m *PortStatisticsInlineDevice) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("device"+"."+"timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PortStatisticsInlineDevice) validateTransmitRaw(formats strfmt.Registry) error {
	if swag.IsZero(m.TransmitRaw) { // not required
		return nil
	}

	if m.TransmitRaw != nil {
		if err := m.TransmitRaw.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device" + "." + "transmit_raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device" + "." + "transmit_raw")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this port statistics inline device based on the context it is used
func (m *PortStatisticsInlineDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReceiveRaw(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransmitRaw(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortStatisticsInlineDevice) contextValidateReceiveRaw(ctx context.Context, formats strfmt.Registry) error {

	if m.ReceiveRaw != nil {

		if swag.IsZero(m.ReceiveRaw) { // not required
			return nil
		}

		if err := m.ReceiveRaw.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device" + "." + "receive_raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device" + "." + "receive_raw")
			}
			return err
		}
	}

	return nil
}

func (m *PortStatisticsInlineDevice) contextValidateTransmitRaw(ctx context.Context, formats strfmt.Registry) error {

	if m.TransmitRaw != nil {

		if swag.IsZero(m.TransmitRaw) { // not required
			return nil
		}

		if err := m.TransmitRaw.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device" + "." + "transmit_raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device" + "." + "transmit_raw")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortStatisticsInlineDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortStatisticsInlineDevice) UnmarshalBinary(b []byte) error {
	var res PortStatisticsInlineDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortStatisticsInlineDeviceInlineReceiveRaw Packet receive counters for the Ethernet port.
//
// swagger:model port_statistics_inline_device_inline_receive_raw
type PortStatisticsInlineDeviceInlineReceiveRaw struct {

	// Total number of discarded packets.
	// Example: 100
	Discards *int64 `json:"discards,omitempty"`

	// Number of packet errors.
	// Example: 200
	Errors *int64 `json:"errors,omitempty"`

	// Total packet count.
	// Example: 500
	Packets *int64 `json:"packets,omitempty"`
}

// Validate validates this port statistics inline device inline receive raw
func (m *PortStatisticsInlineDeviceInlineReceiveRaw) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this port statistics inline device inline receive raw based on context it is used
func (m *PortStatisticsInlineDeviceInlineReceiveRaw) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortStatisticsInlineDeviceInlineReceiveRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortStatisticsInlineDeviceInlineReceiveRaw) UnmarshalBinary(b []byte) error {
	var res PortStatisticsInlineDeviceInlineReceiveRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortStatisticsInlineDeviceInlineTransmitRaw Packet transmit counters for the Ethernet port.
//
// swagger:model port_statistics_inline_device_inline_transmit_raw
type PortStatisticsInlineDeviceInlineTransmitRaw struct {

	// Total number of discarded packets.
	// Example: 100
	Discards *int64 `json:"discards,omitempty"`

	// Number of packet errors.
	// Example: 200
	Errors *int64 `json:"errors,omitempty"`

	// Total packet count.
	// Example: 500
	Packets *int64 `json:"packets,omitempty"`
}

// Validate validates this port statistics inline device inline transmit raw
func (m *PortStatisticsInlineDeviceInlineTransmitRaw) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this port statistics inline device inline transmit raw based on context it is used
func (m *PortStatisticsInlineDeviceInlineTransmitRaw) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortStatisticsInlineDeviceInlineTransmitRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortStatisticsInlineDeviceInlineTransmitRaw) UnmarshalBinary(b []byte) error {
	var res PortStatisticsInlineDeviceInlineTransmitRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortStatisticsInlineThroughputRaw Throughput bytes observed at the port object. This can be used along with delta time to calculate the rate of throughput bytes per unit of time.
//
// swagger:model port_statistics_inline_throughput_raw
type PortStatisticsInlineThroughputRaw struct {

	// Performance metric for read I/O operations.
	// Example: 200
	Read *int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total *int64 `json:"total,omitempty"`

	// Performance metric for write I/O operations.
	// Example: 100
	Write *int64 `json:"write,omitempty"`
}

// Validate validates this port statistics inline throughput raw
func (m *PortStatisticsInlineThroughputRaw) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this port statistics inline throughput raw based on context it is used
func (m *PortStatisticsInlineThroughputRaw) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortStatisticsInlineThroughputRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortStatisticsInlineThroughputRaw) UnmarshalBinary(b []byte) error {
	var res PortStatisticsInlineThroughputRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
