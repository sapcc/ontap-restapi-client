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

// PerformanceMetricReducedThroughput Performance numbers, such as IOPS latency and throughput
//
// swagger:model performance_metric_reduced_throughput
type PerformanceMetricReducedThroughput struct {

	// links
	Links *PerformanceMetricReducedThroughputInlineLinks `json:"_links,omitempty"`

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Read Only: true
	// Enum: ["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]
	Duration *string `json:"duration,omitempty"`

	// iops
	Iops *PerformanceMetricReducedThroughputInlineIops `json:"iops,omitempty"`

	// latency
	Latency *PerformanceMetricReducedThroughputInlineLatency `json:"latency,omitempty"`

	// Any errors associated with the sample. For example, if the aggregation of data over multiple nodes fails then any of the partial errors might be returned, "ok" on success, or "error" on any internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "Inconsistent_ delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "Inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Read Only: true
	// Enum: ["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]
	Status *string `json:"status,omitempty"`

	// throughput
	Throughput *PerformanceMetricReducedThroughputInlineThroughput `json:"throughput,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25 11:20:13
	// Read Only: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this performance metric reduced throughput
func (m *PerformanceMetricReducedThroughput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThroughput(formats); err != nil {
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

func (m *PerformanceMetricReducedThroughput) validateLinks(formats strfmt.Registry) error {
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

var performanceMetricReducedThroughputTypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceMetricReducedThroughputTypeDurationPropEnum = append(performanceMetricReducedThroughputTypeDurationPropEnum, v)
	}
}

const (

	// PerformanceMetricReducedThroughputDurationPT15S captures enum value "PT15S"
	PerformanceMetricReducedThroughputDurationPT15S string = "PT15S"

	// PerformanceMetricReducedThroughputDurationPT4M captures enum value "PT4M"
	PerformanceMetricReducedThroughputDurationPT4M string = "PT4M"

	// PerformanceMetricReducedThroughputDurationPT30M captures enum value "PT30M"
	PerformanceMetricReducedThroughputDurationPT30M string = "PT30M"

	// PerformanceMetricReducedThroughputDurationPT2H captures enum value "PT2H"
	PerformanceMetricReducedThroughputDurationPT2H string = "PT2H"

	// PerformanceMetricReducedThroughputDurationP1D captures enum value "P1D"
	PerformanceMetricReducedThroughputDurationP1D string = "P1D"

	// PerformanceMetricReducedThroughputDurationPT5M captures enum value "PT5M"
	PerformanceMetricReducedThroughputDurationPT5M string = "PT5M"
)

// prop value enum
func (m *PerformanceMetricReducedThroughput) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceMetricReducedThroughputTypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceMetricReducedThroughput) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", *m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceMetricReducedThroughput) validateIops(formats strfmt.Registry) error {
	if swag.IsZero(m.Iops) { // not required
		return nil
	}

	if m.Iops != nil {
		if err := m.Iops.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricReducedThroughput) validateLatency(formats strfmt.Registry) error {
	if swag.IsZero(m.Latency) { // not required
		return nil
	}

	if m.Latency != nil {
		if err := m.Latency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latency")
			}
			return err
		}
	}

	return nil
}

var performanceMetricReducedThroughputTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceMetricReducedThroughputTypeStatusPropEnum = append(performanceMetricReducedThroughputTypeStatusPropEnum, v)
	}
}

const (

	// PerformanceMetricReducedThroughputStatusOk captures enum value "ok"
	PerformanceMetricReducedThroughputStatusOk string = "ok"

	// PerformanceMetricReducedThroughputStatusError captures enum value "error"
	PerformanceMetricReducedThroughputStatusError string = "error"

	// PerformanceMetricReducedThroughputStatusPartialNoData captures enum value "partial_no_data"
	PerformanceMetricReducedThroughputStatusPartialNoData string = "partial_no_data"

	// PerformanceMetricReducedThroughputStatusPartialNoResponse captures enum value "partial_no_response"
	PerformanceMetricReducedThroughputStatusPartialNoResponse string = "partial_no_response"

	// PerformanceMetricReducedThroughputStatusPartialOtherError captures enum value "partial_other_error"
	PerformanceMetricReducedThroughputStatusPartialOtherError string = "partial_other_error"

	// PerformanceMetricReducedThroughputStatusNegativeDelta captures enum value "negative_delta"
	PerformanceMetricReducedThroughputStatusNegativeDelta string = "negative_delta"

	// PerformanceMetricReducedThroughputStatusNotFound captures enum value "not_found"
	PerformanceMetricReducedThroughputStatusNotFound string = "not_found"

	// PerformanceMetricReducedThroughputStatusBackfilledData captures enum value "backfilled_data"
	PerformanceMetricReducedThroughputStatusBackfilledData string = "backfilled_data"

	// PerformanceMetricReducedThroughputStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PerformanceMetricReducedThroughputStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// PerformanceMetricReducedThroughputStatusInconsistentOldData captures enum value "inconsistent_old_data"
	PerformanceMetricReducedThroughputStatusInconsistentOldData string = "inconsistent_old_data"

	// PerformanceMetricReducedThroughputStatusPartialNoUUID captures enum value "partial_no_uuid"
	PerformanceMetricReducedThroughputStatusPartialNoUUID string = "partial_no_uuid"
)

// prop value enum
func (m *PerformanceMetricReducedThroughput) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceMetricReducedThroughputTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceMetricReducedThroughput) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceMetricReducedThroughput) validateThroughput(formats strfmt.Registry) error {
	if swag.IsZero(m.Throughput) { // not required
		return nil
	}

	if m.Throughput != nil {
		if err := m.Throughput.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricReducedThroughput) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this performance metric reduced throughput based on the context it is used
func (m *PerformanceMetricReducedThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIops(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceMetricReducedThroughput) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceMetricReducedThroughput) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceMetricReducedThroughput) contextValidateIops(ctx context.Context, formats strfmt.Registry) error {

	if m.Iops != nil {

		if swag.IsZero(m.Iops) { // not required
			return nil
		}

		if err := m.Iops.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricReducedThroughput) contextValidateLatency(ctx context.Context, formats strfmt.Registry) error {

	if m.Latency != nil {

		if swag.IsZero(m.Latency) { // not required
			return nil
		}

		if err := m.Latency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latency")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricReducedThroughput) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceMetricReducedThroughput) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

	if m.Throughput != nil {

		if swag.IsZero(m.Throughput) { // not required
			return nil
		}

		if err := m.Throughput.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricReducedThroughput) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricReducedThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricReducedThroughput) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricReducedThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceMetricReducedThroughputInlineIops The rate of I/O operations observed at the storage object.
//
// swagger:model performance_metric_reduced_throughput_inline_iops
type PerformanceMetricReducedThroughputInlineIops struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other *int64 `json:"other,omitempty"`

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

// Validate validates this performance metric reduced throughput inline iops
func (m *PerformanceMetricReducedThroughputInlineIops) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance metric reduced throughput inline iops based on the context it is used
func (m *PerformanceMetricReducedThroughputInlineIops) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricReducedThroughputInlineIops) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricReducedThroughputInlineIops) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricReducedThroughputInlineIops
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceMetricReducedThroughputInlineLatency The round trip latency in microseconds observed at the storage object.
//
// swagger:model performance_metric_reduced_throughput_inline_latency
type PerformanceMetricReducedThroughputInlineLatency struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other *int64 `json:"other,omitempty"`

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

// Validate validates this performance metric reduced throughput inline latency
func (m *PerformanceMetricReducedThroughputInlineLatency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance metric reduced throughput inline latency based on the context it is used
func (m *PerformanceMetricReducedThroughputInlineLatency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricReducedThroughputInlineLatency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricReducedThroughputInlineLatency) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricReducedThroughputInlineLatency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceMetricReducedThroughputInlineLinks performance metric reduced throughput inline links
//
// swagger:model performance_metric_reduced_throughput_inline__links
type PerformanceMetricReducedThroughputInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this performance metric reduced throughput inline links
func (m *PerformanceMetricReducedThroughputInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceMetricReducedThroughputInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this performance metric reduced throughput inline links based on the context it is used
func (m *PerformanceMetricReducedThroughputInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceMetricReducedThroughputInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricReducedThroughputInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricReducedThroughputInlineLinks) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricReducedThroughputInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceMetricReducedThroughputInlineThroughput The rate of throughput bytes per second observed at the storage object.
//
// swagger:model performance_metric_reduced_throughput_inline_throughput
type PerformanceMetricReducedThroughputInlineThroughput struct {

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

// Validate validates this performance metric reduced throughput inline throughput
func (m *PerformanceMetricReducedThroughputInlineThroughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance metric reduced throughput inline throughput based on the context it is used
func (m *PerformanceMetricReducedThroughputInlineThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricReducedThroughputInlineThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricReducedThroughputInlineThroughput) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricReducedThroughputInlineThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
