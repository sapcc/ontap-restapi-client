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

// CloudStore cloud store
//
// swagger:model cloud_store
type CloudStore struct {

	// links
	Links *CloudStoreInlineLinks `json:"_links,omitempty"`

	// aggregate
	Aggregate *CloudStoreInlineAggregate `json:"aggregate,omitempty"`

	// Availability of the object store.
	// Read Only: true
	// Enum: ["available","unavailable"]
	Availability *string `json:"availability,omitempty"`

	// Availability of the object store at the HA partner.
	// Read Only: true
	// Enum: ["available","unavailable"]
	AvailabilityAtPartner *string `json:"availability_at_partner,omitempty"`

	// This field identifies if the mirror cloud store is in sync with the primary cloud store of a FabricPool.
	// Read Only: true
	MirrorDegraded *bool `json:"mirror_degraded,omitempty"`

	// This field indicates whether the cloud store is the primary cloud store of a mirrored FabricPool.
	Primary *bool `json:"primary,omitempty"`

	// Resync progress of the mirror object store in percentage.
	// Read Only: true
	ResyncProgress *int64 `json:"resync-progress,omitempty"`

	// target
	Target *CloudStoreInlineTarget `json:"target,omitempty"`

	// unavailable reason
	UnavailableReason *CloudStoreInlineUnavailableReason `json:"unavailable_reason,omitempty"`

	// Usage threshold for reclaiming unused space in the cloud store. Valid values are 0 to 99. The default value depends on the provider type. This can be specified in PATCH but not POST.
	// Example: 20
	UnreclaimedSpaceThreshold *int64 `json:"unreclaimed_space_threshold,omitempty"`

	// The amount of object space used. Calculated every 5 minutes and cached.
	// Read Only: true
	Used *int64 `json:"used,omitempty"`
}

// Validate validates this cloud store
func (m *CloudStore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAggregate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailabilityAtPartner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnavailableReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStore) validateLinks(formats strfmt.Registry) error {
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

func (m *CloudStore) validateAggregate(formats strfmt.Registry) error {
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

var cloudStoreTypeAvailabilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["available","unavailable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudStoreTypeAvailabilityPropEnum = append(cloudStoreTypeAvailabilityPropEnum, v)
	}
}

const (

	// CloudStoreAvailabilityAvailable captures enum value "available"
	CloudStoreAvailabilityAvailable string = "available"

	// CloudStoreAvailabilityUnavailable captures enum value "unavailable"
	CloudStoreAvailabilityUnavailable string = "unavailable"
)

// prop value enum
func (m *CloudStore) validateAvailabilityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cloudStoreTypeAvailabilityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CloudStore) validateAvailability(formats strfmt.Registry) error {
	if swag.IsZero(m.Availability) { // not required
		return nil
	}

	// value enum
	if err := m.validateAvailabilityEnum("availability", "body", *m.Availability); err != nil {
		return err
	}

	return nil
}

var cloudStoreTypeAvailabilityAtPartnerPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["available","unavailable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudStoreTypeAvailabilityAtPartnerPropEnum = append(cloudStoreTypeAvailabilityAtPartnerPropEnum, v)
	}
}

const (

	// CloudStoreAvailabilityAtPartnerAvailable captures enum value "available"
	CloudStoreAvailabilityAtPartnerAvailable string = "available"

	// CloudStoreAvailabilityAtPartnerUnavailable captures enum value "unavailable"
	CloudStoreAvailabilityAtPartnerUnavailable string = "unavailable"
)

// prop value enum
func (m *CloudStore) validateAvailabilityAtPartnerEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cloudStoreTypeAvailabilityAtPartnerPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CloudStore) validateAvailabilityAtPartner(formats strfmt.Registry) error {
	if swag.IsZero(m.AvailabilityAtPartner) { // not required
		return nil
	}

	// value enum
	if err := m.validateAvailabilityAtPartnerEnum("availability_at_partner", "body", *m.AvailabilityAtPartner); err != nil {
		return err
	}

	return nil
}

func (m *CloudStore) validateTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *CloudStore) validateUnavailableReason(formats strfmt.Registry) error {
	if swag.IsZero(m.UnavailableReason) { // not required
		return nil
	}

	if m.UnavailableReason != nil {
		if err := m.UnavailableReason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unavailable_reason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unavailable_reason")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cloud store based on the context it is used
func (m *CloudStore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAggregate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAvailability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAvailabilityAtPartner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMirrorDegraded(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResyncProgress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnavailableReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStore) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CloudStore) contextValidateAggregate(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CloudStore) contextValidateAvailability(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "availability", "body", m.Availability); err != nil {
		return err
	}

	return nil
}

func (m *CloudStore) contextValidateAvailabilityAtPartner(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "availability_at_partner", "body", m.AvailabilityAtPartner); err != nil {
		return err
	}

	return nil
}

func (m *CloudStore) contextValidateMirrorDegraded(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "mirror_degraded", "body", m.MirrorDegraded); err != nil {
		return err
	}

	return nil
}

func (m *CloudStore) contextValidateResyncProgress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "resync-progress", "body", m.ResyncProgress); err != nil {
		return err
	}

	return nil
}

func (m *CloudStore) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {

		if swag.IsZero(m.Target) { // not required
			return nil
		}

		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *CloudStore) contextValidateUnavailableReason(ctx context.Context, formats strfmt.Registry) error {

	if m.UnavailableReason != nil {

		if swag.IsZero(m.UnavailableReason) { // not required
			return nil
		}

		if err := m.UnavailableReason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unavailable_reason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unavailable_reason")
			}
			return err
		}
	}

	return nil
}

func (m *CloudStore) contextValidateUsed(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "used", "body", m.Used); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudStore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudStore) UnmarshalBinary(b []byte) error {
	var res CloudStore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CloudStoreInlineAggregate Aggregate
//
// swagger:model cloud_store_inline_aggregate
type CloudStoreInlineAggregate struct {

	// name
	// Example: aggr1
	Name *string `json:"name,omitempty"`
}

// Validate validates this cloud store inline aggregate
func (m *CloudStoreInlineAggregate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cloud store inline aggregate based on context it is used
func (m *CloudStoreInlineAggregate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloudStoreInlineAggregate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudStoreInlineAggregate) UnmarshalBinary(b []byte) error {
	var res CloudStoreInlineAggregate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CloudStoreInlineLinks cloud store inline links
//
// swagger:model cloud_store_inline__links
type CloudStoreInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cloud store inline links
func (m *CloudStoreInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStoreInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this cloud store inline links based on the context it is used
func (m *CloudStoreInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStoreInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CloudStoreInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudStoreInlineLinks) UnmarshalBinary(b []byte) error {
	var res CloudStoreInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CloudStoreInlineTarget Cloud target
//
// swagger:model cloud_store_inline_target
type CloudStoreInlineTarget struct {

	// links
	Links *CloudStoreInlineTargetInlineLinks `json:"_links,omitempty"`

	// name
	// Example: target1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this cloud store inline target
func (m *CloudStoreInlineTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStoreInlineTarget) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cloud store inline target based on the context it is used
func (m *CloudStoreInlineTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStoreInlineTarget) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudStoreInlineTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudStoreInlineTarget) UnmarshalBinary(b []byte) error {
	var res CloudStoreInlineTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CloudStoreInlineTargetInlineLinks cloud store inline target inline links
//
// swagger:model cloud_store_inline_target_inline__links
type CloudStoreInlineTargetInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cloud store inline target inline links
func (m *CloudStoreInlineTargetInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStoreInlineTargetInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cloud store inline target inline links based on the context it is used
func (m *CloudStoreInlineTargetInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStoreInlineTargetInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudStoreInlineTargetInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudStoreInlineTargetInlineLinks) UnmarshalBinary(b []byte) error {
	var res CloudStoreInlineTargetInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CloudStoreInlineUnavailableReason cloud store inline unavailable reason
//
// swagger:model cloud_store_inline_unavailable_reason
type CloudStoreInlineUnavailableReason struct {

	// Indicates why the object store is unavailable.
	// Read Only: true
	Message *string `json:"message,omitempty"`
}

// Validate validates this cloud store inline unavailable reason
func (m *CloudStoreInlineUnavailableReason) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this cloud store inline unavailable reason based on the context it is used
func (m *CloudStoreInlineUnavailableReason) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStoreInlineUnavailableReason) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "unavailable_reason"+"."+"message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudStoreInlineUnavailableReason) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudStoreInlineUnavailableReason) UnmarshalBinary(b []byte) error {
	var res CloudStoreInlineUnavailableReason
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
