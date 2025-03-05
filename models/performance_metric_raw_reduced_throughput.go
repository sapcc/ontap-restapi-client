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

// PerformanceMetricRawReducedThroughput These are raw performance numbers, such as IOPS latency and throughput. These numbers are aggregated across all nodes in the cluster and increase with the uptime of the cluster.
//
// swagger:model performance_metric_raw_reduced_throughput
type PerformanceMetricRawReducedThroughput struct {

	// iops raw
	IopsRaw *PerformanceMetricRawReducedThroughputInlineIopsRaw `json:"iops_raw,omitempty"`

	// latency raw
	LatencyRaw *PerformanceMetricRawReducedThroughputInlineLatencyRaw `json:"latency_raw,omitempty"`

	// Any errors associated with the sample. For example, if the aggregation of data over multiple nodes fails then any of the partial errors might be returned, "ok" on success, or "error" on any internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "Inconsistent_delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "Inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Read Only: true
	// Enum: ["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]
	Status *string `json:"status,omitempty"`

	// throughput raw
	ThroughputRaw *PerformanceMetricRawReducedThroughputInlineThroughputRaw `json:"throughput_raw,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25 11:20:13
	// Read Only: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this performance metric raw reduced throughput
func (m *PerformanceMetricRawReducedThroughput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIopsRaw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatencyRaw(formats); err != nil {
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

func (m *PerformanceMetricRawReducedThroughput) validateIopsRaw(formats strfmt.Registry) error {
	if swag.IsZero(m.IopsRaw) { // not required
		return nil
	}

	if m.IopsRaw != nil {
		if err := m.IopsRaw.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops_raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iops_raw")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricRawReducedThroughput) validateLatencyRaw(formats strfmt.Registry) error {
	if swag.IsZero(m.LatencyRaw) { // not required
		return nil
	}

	if m.LatencyRaw != nil {
		if err := m.LatencyRaw.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency_raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latency_raw")
			}
			return err
		}
	}

	return nil
}

var performanceMetricRawReducedThroughputTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceMetricRawReducedThroughputTypeStatusPropEnum = append(performanceMetricRawReducedThroughputTypeStatusPropEnum, v)
	}
}

const (

	// PerformanceMetricRawReducedThroughputStatusOk captures enum value "ok"
	PerformanceMetricRawReducedThroughputStatusOk string = "ok"

	// PerformanceMetricRawReducedThroughputStatusError captures enum value "error"
	PerformanceMetricRawReducedThroughputStatusError string = "error"

	// PerformanceMetricRawReducedThroughputStatusPartialNoData captures enum value "partial_no_data"
	PerformanceMetricRawReducedThroughputStatusPartialNoData string = "partial_no_data"

	// PerformanceMetricRawReducedThroughputStatusPartialNoResponse captures enum value "partial_no_response"
	PerformanceMetricRawReducedThroughputStatusPartialNoResponse string = "partial_no_response"

	// PerformanceMetricRawReducedThroughputStatusPartialOtherError captures enum value "partial_other_error"
	PerformanceMetricRawReducedThroughputStatusPartialOtherError string = "partial_other_error"

	// PerformanceMetricRawReducedThroughputStatusNegativeDelta captures enum value "negative_delta"
	PerformanceMetricRawReducedThroughputStatusNegativeDelta string = "negative_delta"

	// PerformanceMetricRawReducedThroughputStatusNotFound captures enum value "not_found"
	PerformanceMetricRawReducedThroughputStatusNotFound string = "not_found"

	// PerformanceMetricRawReducedThroughputStatusBackfilledData captures enum value "backfilled_data"
	PerformanceMetricRawReducedThroughputStatusBackfilledData string = "backfilled_data"

	// PerformanceMetricRawReducedThroughputStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PerformanceMetricRawReducedThroughputStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// PerformanceMetricRawReducedThroughputStatusInconsistentOldData captures enum value "inconsistent_old_data"
	PerformanceMetricRawReducedThroughputStatusInconsistentOldData string = "inconsistent_old_data"

	// PerformanceMetricRawReducedThroughputStatusPartialNoUUID captures enum value "partial_no_uuid"
	PerformanceMetricRawReducedThroughputStatusPartialNoUUID string = "partial_no_uuid"
)

// prop value enum
func (m *PerformanceMetricRawReducedThroughput) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceMetricRawReducedThroughputTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceMetricRawReducedThroughput) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceMetricRawReducedThroughput) validateThroughputRaw(formats strfmt.Registry) error {
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

func (m *PerformanceMetricRawReducedThroughput) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this performance metric raw reduced throughput based on the context it is used
func (m *PerformanceMetricRawReducedThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIopsRaw(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatencyRaw(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughputRaw(ctx, formats); err != nil {
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

func (m *PerformanceMetricRawReducedThroughput) contextValidateIopsRaw(ctx context.Context, formats strfmt.Registry) error {

	if m.IopsRaw != nil {

		if swag.IsZero(m.IopsRaw) { // not required
			return nil
		}

		if err := m.IopsRaw.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops_raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iops_raw")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricRawReducedThroughput) contextValidateLatencyRaw(ctx context.Context, formats strfmt.Registry) error {

	if m.LatencyRaw != nil {

		if swag.IsZero(m.LatencyRaw) { // not required
			return nil
		}

		if err := m.LatencyRaw.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency_raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latency_raw")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricRawReducedThroughput) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceMetricRawReducedThroughput) contextValidateThroughputRaw(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceMetricRawReducedThroughput) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricRawReducedThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricRawReducedThroughput) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricRawReducedThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceMetricRawReducedThroughputInlineIopsRaw The number of I/O operations observed at the storage object. This should be used along with delta time to calculate the rate of I/O operations per unit of time.
//
// swagger:model performance_metric_raw_reduced_throughput_inline_iops_raw
type PerformanceMetricRawReducedThroughputInlineIopsRaw struct {

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

// Validate validates this performance metric raw reduced throughput inline iops raw
func (m *PerformanceMetricRawReducedThroughputInlineIopsRaw) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance metric raw reduced throughput inline iops raw based on the context it is used
func (m *PerformanceMetricRawReducedThroughputInlineIopsRaw) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricRawReducedThroughputInlineIopsRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricRawReducedThroughputInlineIopsRaw) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricRawReducedThroughputInlineIopsRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceMetricRawReducedThroughputInlineLatencyRaw The raw latency in microseconds observed at the storage object. This should be divided by the raw IOPS value to calculate the average latency per I/O operation.
//
// swagger:model performance_metric_raw_reduced_throughput_inline_latency_raw
type PerformanceMetricRawReducedThroughputInlineLatencyRaw struct {

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

// Validate validates this performance metric raw reduced throughput inline latency raw
func (m *PerformanceMetricRawReducedThroughputInlineLatencyRaw) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance metric raw reduced throughput inline latency raw based on the context it is used
func (m *PerformanceMetricRawReducedThroughputInlineLatencyRaw) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricRawReducedThroughputInlineLatencyRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricRawReducedThroughputInlineLatencyRaw) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricRawReducedThroughputInlineLatencyRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceMetricRawReducedThroughputInlineThroughputRaw Throughput bytes observed at the storage object. This should be used along with delta time to calculate the rate of throughput bytes per unit of time.
//
// swagger:model performance_metric_raw_reduced_throughput_inline_throughput_raw
type PerformanceMetricRawReducedThroughputInlineThroughputRaw struct {

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

// Validate validates this performance metric raw reduced throughput inline throughput raw
func (m *PerformanceMetricRawReducedThroughputInlineThroughputRaw) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance metric raw reduced throughput inline throughput raw based on the context it is used
func (m *PerformanceMetricRawReducedThroughputInlineThroughputRaw) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricRawReducedThroughputInlineThroughputRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricRawReducedThroughputInlineThroughputRaw) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricRawReducedThroughputInlineThroughputRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
