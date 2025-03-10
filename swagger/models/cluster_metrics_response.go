// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterMetricsResponse cluster metrics response
//
// swagger:model cluster_metrics_response
type ClusterMetricsResponse struct {

	// links
	Links *ClusterMetricsResponseInlineLinks `json:"_links,omitempty"`

	// cluster metrics response inline records
	ClusterMetricsResponseInlineRecords []*ClusterMetricsResponseInlineRecordsInlineArrayItem `json:"records,omitempty"`

	// Number of records
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty"`
}

// Validate validates this cluster metrics response
func (m *ClusterMetricsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterMetricsResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterMetricsResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *ClusterMetricsResponse) validateClusterMetricsResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterMetricsResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterMetricsResponseInlineRecords); i++ {
		if swag.IsZero(m.ClusterMetricsResponseInlineRecords[i]) { // not required
			continue
		}

		if m.ClusterMetricsResponseInlineRecords[i] != nil {
			if err := m.ClusterMetricsResponseInlineRecords[i].Validate(formats); err != nil {
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

// ContextValidate validate this cluster metrics response based on the context it is used
func (m *ClusterMetricsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterMetricsResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterMetricsResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterMetricsResponse) contextValidateClusterMetricsResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClusterMetricsResponseInlineRecords); i++ {

		if m.ClusterMetricsResponseInlineRecords[i] != nil {

			if swag.IsZero(m.ClusterMetricsResponseInlineRecords[i]) { // not required
				return nil
			}

			if err := m.ClusterMetricsResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (m *ClusterMetricsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterMetricsResponse) UnmarshalBinary(b []byte) error {
	var res ClusterMetricsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterMetricsResponseInlineLinks cluster metrics response inline links
//
// swagger:model cluster_metrics_response_inline__links
type ClusterMetricsResponseInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cluster metrics response inline links
func (m *ClusterMetricsResponseInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterMetricsResponseInlineLinks) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterMetricsResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this cluster metrics response inline links based on the context it is used
func (m *ClusterMetricsResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterMetricsResponseInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {

		if swag.IsZero(m.Next) { // not required
			return nil
		}

		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterMetricsResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ClusterMetricsResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterMetricsResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res ClusterMetricsResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterMetricsResponseInlineRecordsInlineArrayItem Performance numbers, such as IOPS latency and throughput.
//
// swagger:model cluster_metrics_response_inline_records_inline_array_item
type ClusterMetricsResponseInlineRecordsInlineArrayItem struct {

	// links
	Links *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Read Only: true
	// Enum: ["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]
	Duration *string `json:"duration,omitempty"`

	// iops
	Iops *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineIops `json:"iops,omitempty"`

	// latency
	Latency *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLatency `json:"latency,omitempty"`

	// Errors associated with the sample. For example, if the aggregation of data over multiple nodes fails, then any partial errors might return "ok" on success or "error" on an internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "Inconsistent_ delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "Inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Read Only: true
	// Enum: ["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]
	Status *string `json:"status,omitempty"`

	// throughput
	Throughput *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineThroughput `json:"throughput,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25 11:20:13
	// Read Only: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this cluster metrics response inline records inline array item
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) Validate(formats strfmt.Registry) error {
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

func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) validateLinks(formats strfmt.Registry) error {
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

var clusterMetricsResponseInlineRecordsInlineArrayItemTypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterMetricsResponseInlineRecordsInlineArrayItemTypeDurationPropEnum = append(clusterMetricsResponseInlineRecordsInlineArrayItemTypeDurationPropEnum, v)
	}
}

const (

	// ClusterMetricsResponseInlineRecordsInlineArrayItemDurationPT15S captures enum value "PT15S"
	ClusterMetricsResponseInlineRecordsInlineArrayItemDurationPT15S string = "PT15S"

	// ClusterMetricsResponseInlineRecordsInlineArrayItemDurationPT4M captures enum value "PT4M"
	ClusterMetricsResponseInlineRecordsInlineArrayItemDurationPT4M string = "PT4M"

	// ClusterMetricsResponseInlineRecordsInlineArrayItemDurationPT30M captures enum value "PT30M"
	ClusterMetricsResponseInlineRecordsInlineArrayItemDurationPT30M string = "PT30M"

	// ClusterMetricsResponseInlineRecordsInlineArrayItemDurationPT2H captures enum value "PT2H"
	ClusterMetricsResponseInlineRecordsInlineArrayItemDurationPT2H string = "PT2H"

	// ClusterMetricsResponseInlineRecordsInlineArrayItemDurationP1D captures enum value "P1D"
	ClusterMetricsResponseInlineRecordsInlineArrayItemDurationP1D string = "P1D"

	// ClusterMetricsResponseInlineRecordsInlineArrayItemDurationPT5M captures enum value "PT5M"
	ClusterMetricsResponseInlineRecordsInlineArrayItemDurationPT5M string = "PT5M"
)

// prop value enum
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterMetricsResponseInlineRecordsInlineArrayItemTypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", *m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) validateIops(formats strfmt.Registry) error {
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

func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) validateLatency(formats strfmt.Registry) error {
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

var clusterMetricsResponseInlineRecordsInlineArrayItemTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterMetricsResponseInlineRecordsInlineArrayItemTypeStatusPropEnum = append(clusterMetricsResponseInlineRecordsInlineArrayItemTypeStatusPropEnum, v)
	}
}

const (

	// ClusterMetricsResponseInlineRecordsInlineArrayItemStatusOk captures enum value "ok"
	ClusterMetricsResponseInlineRecordsInlineArrayItemStatusOk string = "ok"

	// ClusterMetricsResponseInlineRecordsInlineArrayItemStatusError captures enum value "error"
	ClusterMetricsResponseInlineRecordsInlineArrayItemStatusError string = "error"

	// ClusterMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoData captures enum value "partial_no_data"
	ClusterMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoData string = "partial_no_data"

	// ClusterMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoResponse captures enum value "partial_no_response"
	ClusterMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoResponse string = "partial_no_response"

	// ClusterMetricsResponseInlineRecordsInlineArrayItemStatusPartialOtherError captures enum value "partial_other_error"
	ClusterMetricsResponseInlineRecordsInlineArrayItemStatusPartialOtherError string = "partial_other_error"

	// ClusterMetricsResponseInlineRecordsInlineArrayItemStatusNegativeDelta captures enum value "negative_delta"
	ClusterMetricsResponseInlineRecordsInlineArrayItemStatusNegativeDelta string = "negative_delta"

	// ClusterMetricsResponseInlineRecordsInlineArrayItemStatusNotFound captures enum value "not_found"
	ClusterMetricsResponseInlineRecordsInlineArrayItemStatusNotFound string = "not_found"

	// ClusterMetricsResponseInlineRecordsInlineArrayItemStatusBackfilledData captures enum value "backfilled_data"
	ClusterMetricsResponseInlineRecordsInlineArrayItemStatusBackfilledData string = "backfilled_data"

	// ClusterMetricsResponseInlineRecordsInlineArrayItemStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	ClusterMetricsResponseInlineRecordsInlineArrayItemStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// ClusterMetricsResponseInlineRecordsInlineArrayItemStatusInconsistentOldData captures enum value "inconsistent_old_data"
	ClusterMetricsResponseInlineRecordsInlineArrayItemStatusInconsistentOldData string = "inconsistent_old_data"

	// ClusterMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoUUID captures enum value "partial_no_uuid"
	ClusterMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoUUID string = "partial_no_uuid"
)

// prop value enum
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterMetricsResponseInlineRecordsInlineArrayItemTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) validateThroughput(formats strfmt.Registry) error {
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

func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cluster metrics response inline records inline array item based on the context it is used
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) contextValidateIops(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) contextValidateLatency(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res ClusterMetricsResponseInlineRecordsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterMetricsResponseInlineRecordsInlineArrayItemInlineIops The rate of I/O operations observed at the storage object.
//
// swagger:model cluster_metrics_response_inline_records_inline_array_item_inline_iops
type ClusterMetricsResponseInlineRecordsInlineArrayItemInlineIops struct {

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

// Validate validates this cluster metrics response inline records inline array item inline iops
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineIops) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this cluster metrics response inline records inline array item inline iops based on the context it is used
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineIops) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineIops) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineIops) UnmarshalBinary(b []byte) error {
	var res ClusterMetricsResponseInlineRecordsInlineArrayItemInlineIops
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLatency The round trip latency in microseconds observed at the storage object.
//
// swagger:model cluster_metrics_response_inline_records_inline_array_item_inline_latency
type ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLatency struct {

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

// Validate validates this cluster metrics response inline records inline array item inline latency
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLatency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this cluster metrics response inline records inline array item inline latency based on the context it is used
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLatency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLatency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLatency) UnmarshalBinary(b []byte) error {
	var res ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLatency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLinks cluster metrics response inline records inline array item inline links
//
// swagger:model cluster_metrics_response_inline_records_inline_array_item_inline__links
type ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cluster metrics response inline records inline array item inline links
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this cluster metrics response inline records inline array item inline links based on the context it is used
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res ClusterMetricsResponseInlineRecordsInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterMetricsResponseInlineRecordsInlineArrayItemInlineThroughput The rate of throughput bytes per second observed at the storage object.
//
// swagger:model cluster_metrics_response_inline_records_inline_array_item_inline_throughput
type ClusterMetricsResponseInlineRecordsInlineArrayItemInlineThroughput struct {

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

// Validate validates this cluster metrics response inline records inline array item inline throughput
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineThroughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this cluster metrics response inline records inline array item inline throughput based on the context it is used
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterMetricsResponseInlineRecordsInlineArrayItemInlineThroughput) UnmarshalBinary(b []byte) error {
	var res ClusterMetricsResponseInlineRecordsInlineArrayItemInlineThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
