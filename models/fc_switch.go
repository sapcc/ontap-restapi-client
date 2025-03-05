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

// FcSwitch A Fibre Channel switch.
//
// swagger:model fc_switch
type FcSwitch struct {

	// links
	Links *FcSwitchInlineLinks `json:"_links,omitempty"`

	// cache
	Cache *FcSwitchInlineCache `json:"cache,omitempty"`

	// The domain identifier (ID) of the Fibre Channel (FC) switch. The domain ID is a unique identifier for the FC switch in the FC fabric.
	//
	// Example: 1
	// Read Only: true
	// Maximum: 239
	// Minimum: 1
	DomainID *int64 `json:"domain_id,omitempty"`

	// fabric
	Fabric *FcSwitchInlineFabric `json:"fabric,omitempty"`

	// An array of the Fibre Channel (FC) switch's ports and their attached FC devices.
	//
	// Read Only: true
	FcSwitchInlinePorts []*FcSwitchInlinePortsInlineArrayItem `json:"ports,omitempty"`

	// The logical name of the Fibre Channel switch.
	//
	// Example: switch1
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// The firmware release of the Fibre Channel switch.
	//
	// Example: 1.0.
	// Read Only: true
	Release *string `json:"release,omitempty"`

	// The vendor of the Fibre Channel switch.
	//
	// Example: vendor1
	// Read Only: true
	Vendor *string `json:"vendor,omitempty"`

	// The world-wide name (WWN) for the Fibre Channel switch.
	//
	// Example: 10:00:e1:e2:e3:e4:e5:e6
	// Read Only: true
	Wwn *string `json:"wwn,omitempty"`
}

// Validate validates this fc switch
func (m *FcSwitch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCache(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomainID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFabric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFcSwitchInlinePorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitch) validateLinks(formats strfmt.Registry) error {
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

func (m *FcSwitch) validateCache(formats strfmt.Registry) error {
	if swag.IsZero(m.Cache) { // not required
		return nil
	}

	if m.Cache != nil {
		if err := m.Cache.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cache")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cache")
			}
			return err
		}
	}

	return nil
}

func (m *FcSwitch) validateDomainID(formats strfmt.Registry) error {
	if swag.IsZero(m.DomainID) { // not required
		return nil
	}

	if err := validate.MinimumInt("domain_id", "body", *m.DomainID, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("domain_id", "body", *m.DomainID, 239, false); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitch) validateFabric(formats strfmt.Registry) error {
	if swag.IsZero(m.Fabric) { // not required
		return nil
	}

	if m.Fabric != nil {
		if err := m.Fabric.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fabric")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fabric")
			}
			return err
		}
	}

	return nil
}

func (m *FcSwitch) validateFcSwitchInlinePorts(formats strfmt.Registry) error {
	if swag.IsZero(m.FcSwitchInlinePorts) { // not required
		return nil
	}

	for i := 0; i < len(m.FcSwitchInlinePorts); i++ {
		if swag.IsZero(m.FcSwitchInlinePorts[i]) { // not required
			continue
		}

		if m.FcSwitchInlinePorts[i] != nil {
			if err := m.FcSwitchInlinePorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this fc switch based on the context it is used
func (m *FcSwitch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCache(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDomainID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFabric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFcSwitchInlinePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelease(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVendor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitch) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *FcSwitch) contextValidateCache(ctx context.Context, formats strfmt.Registry) error {

	if m.Cache != nil {

		if swag.IsZero(m.Cache) { // not required
			return nil
		}

		if err := m.Cache.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cache")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cache")
			}
			return err
		}
	}

	return nil
}

func (m *FcSwitch) contextValidateDomainID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "domain_id", "body", m.DomainID); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitch) contextValidateFabric(ctx context.Context, formats strfmt.Registry) error {

	if m.Fabric != nil {

		if swag.IsZero(m.Fabric) { // not required
			return nil
		}

		if err := m.Fabric.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fabric")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fabric")
			}
			return err
		}
	}

	return nil
}

func (m *FcSwitch) contextValidateFcSwitchInlinePorts(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ports", "body", []*FcSwitchInlinePortsInlineArrayItem(m.FcSwitchInlinePorts)); err != nil {
		return err
	}

	for i := 0; i < len(m.FcSwitchInlinePorts); i++ {

		if m.FcSwitchInlinePorts[i] != nil {

			if swag.IsZero(m.FcSwitchInlinePorts[i]) { // not required
				return nil
			}

			if err := m.FcSwitchInlinePorts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FcSwitch) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitch) contextValidateRelease(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "release", "body", m.Release); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitch) contextValidateVendor(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "vendor", "body", m.Vendor); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitch) contextValidateWwn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "wwn", "body", m.Wwn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcSwitch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcSwitch) UnmarshalBinary(b []byte) error {
	var res FcSwitch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcSwitchInlineCache Properties of Fibre Chanel fabric cache.
//
// swagger:model fc_switch_inline_cache
type FcSwitchInlineCache struct {

	// The age of the Fibre Channel fabric data cache retrieved. If the FC fabric data cache has not been fully updated for a newly discovered fabric, or a fabric that has been re-discovered after being purged, a value for this property will not be retrieved. The value is in ISO 8601 duration format.
	//
	// Example: PT3M30S
	// Read Only: true
	Age *string `json:"age,omitempty"`

	// A boolean that indicates if the retrieved data is current relative to the `cache.maximum_age` value of the request. A value of `true` indicates that the data is no older than the requested maximum age. A value of `false` indicates that the data is older than the requested maximum age; if more current data is required, the caller should wait for some time for the cache update to complete and query the data again.
	//
	// Read Only: true
	IsCurrent *bool `json:"is_current,omitempty"`

	// The date and time at which the Fibre Channel fabric data cache retrieved was last updated. If the FC fabric data cache has not been fully updated for a newly discovered fabric, or a fabric that has been re-discovered after being purged, a value for this property will not be retrieved.
	//
	// Read Only: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this fc switch inline cache
func (m *FcSwitchInlineCache) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchInlineCache) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("cache"+"."+"update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this fc switch inline cache based on the context it is used
func (m *FcSwitchInlineCache) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsCurrent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchInlineCache) contextValidateAge(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cache"+"."+"age", "body", m.Age); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitchInlineCache) contextValidateIsCurrent(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cache"+"."+"is_current", "body", m.IsCurrent); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitchInlineCache) contextValidateUpdateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cache"+"."+"update_time", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcSwitchInlineCache) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcSwitchInlineCache) UnmarshalBinary(b []byte) error {
	var res FcSwitchInlineCache
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcSwitchInlineFabric A reference to a Fibre Channel fabric.
//
// swagger:model fc_switch_inline_fabric
type FcSwitchInlineFabric struct {

	// links
	Links *FcSwitchInlineFabricInlineLinks `json:"_links,omitempty"`

	// The world wide name (WWN) of the primary switch of the Fibre Channel (FC) fabric. This is used as a unique identifier for the FC fabric.
	//
	// Example: 10:00:d1:d2:d3:d4:d5:d6
	Name *string `json:"name,omitempty"`
}

// Validate validates this fc switch inline fabric
func (m *FcSwitchInlineFabric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchInlineFabric) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fabric" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fabric" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc switch inline fabric based on the context it is used
func (m *FcSwitchInlineFabric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchInlineFabric) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fabric" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fabric" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcSwitchInlineFabric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcSwitchInlineFabric) UnmarshalBinary(b []byte) error {
	var res FcSwitchInlineFabric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcSwitchInlineFabricInlineLinks fc switch inline fabric inline links
//
// swagger:model fc_switch_inline_fabric_inline__links
type FcSwitchInlineFabricInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this fc switch inline fabric inline links
func (m *FcSwitchInlineFabricInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchInlineFabricInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fabric" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fabric" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc switch inline fabric inline links based on the context it is used
func (m *FcSwitchInlineFabricInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchInlineFabricInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fabric" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fabric" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcSwitchInlineFabricInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcSwitchInlineFabricInlineLinks) UnmarshalBinary(b []byte) error {
	var res FcSwitchInlineFabricInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcSwitchInlineLinks fc switch inline links
//
// swagger:model fc_switch_inline__links
type FcSwitchInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this fc switch inline links
func (m *FcSwitchInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this fc switch inline links based on the context it is used
func (m *FcSwitchInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *FcSwitchInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcSwitchInlineLinks) UnmarshalBinary(b []byte) error {
	var res FcSwitchInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcSwitchInlinePortsInlineArrayItem A Fibre Channel switch port.
//
// swagger:model fc_switch_inline_ports_inline_array_item
type FcSwitchInlinePortsInlineArrayItem struct {

	// attached device
	AttachedDevice *FcSwitchInlinePortsInlineArrayItemInlineAttachedDevice `json:"attached_device,omitempty"`

	// The slot of the Fibre Channel switch port.
	//
	// Example: 1
	// Read Only: true
	Slot *string `json:"slot,omitempty"`

	// The state of the Fibre Channel switch port.
	//
	// Example: online
	// Read Only: true
	// Enum: ["unknown","online","offline","testing","fault"]
	State *string `json:"state,omitempty"`

	// The type of the Fibre Channel switch port.
	//
	// Read Only: true
	// Enum: ["b_port","e_port","f_port","fl_port","fnl_port","fv_port","n_port","nl_port","nv_port","nx_port","sd_port","te_port","tf_port","tl_port","tnp_port","none"]
	Type *string `json:"type,omitempty"`

	// The world wide port name (WWPN) of the Fibre Channel switch port.
	//
	// Example: 50:0a:31:32:33:34:35:36
	// Read Only: true
	Wwpn *string `json:"wwpn,omitempty"`
}

// Validate validates this fc switch inline ports inline array item
func (m *FcSwitchInlinePortsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachedDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
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

func (m *FcSwitchInlinePortsInlineArrayItem) validateAttachedDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.AttachedDevice) { // not required
		return nil
	}

	if m.AttachedDevice != nil {
		if err := m.AttachedDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attached_device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attached_device")
			}
			return err
		}
	}

	return nil
}

var fcSwitchInlinePortsInlineArrayItemTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","online","offline","testing","fault"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fcSwitchInlinePortsInlineArrayItemTypeStatePropEnum = append(fcSwitchInlinePortsInlineArrayItemTypeStatePropEnum, v)
	}
}

const (

	// FcSwitchInlinePortsInlineArrayItemStateUnknown captures enum value "unknown"
	FcSwitchInlinePortsInlineArrayItemStateUnknown string = "unknown"

	// FcSwitchInlinePortsInlineArrayItemStateOnline captures enum value "online"
	FcSwitchInlinePortsInlineArrayItemStateOnline string = "online"

	// FcSwitchInlinePortsInlineArrayItemStateOffline captures enum value "offline"
	FcSwitchInlinePortsInlineArrayItemStateOffline string = "offline"

	// FcSwitchInlinePortsInlineArrayItemStateTesting captures enum value "testing"
	FcSwitchInlinePortsInlineArrayItemStateTesting string = "testing"

	// FcSwitchInlinePortsInlineArrayItemStateFault captures enum value "fault"
	FcSwitchInlinePortsInlineArrayItemStateFault string = "fault"
)

// prop value enum
func (m *FcSwitchInlinePortsInlineArrayItem) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fcSwitchInlinePortsInlineArrayItemTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FcSwitchInlinePortsInlineArrayItem) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

var fcSwitchInlinePortsInlineArrayItemTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["b_port","e_port","f_port","fl_port","fnl_port","fv_port","n_port","nl_port","nv_port","nx_port","sd_port","te_port","tf_port","tl_port","tnp_port","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fcSwitchInlinePortsInlineArrayItemTypeTypePropEnum = append(fcSwitchInlinePortsInlineArrayItemTypeTypePropEnum, v)
	}
}

const (

	// FcSwitchInlinePortsInlineArrayItemTypeBPort captures enum value "b_port"
	FcSwitchInlinePortsInlineArrayItemTypeBPort string = "b_port"

	// FcSwitchInlinePortsInlineArrayItemTypeEPort captures enum value "e_port"
	FcSwitchInlinePortsInlineArrayItemTypeEPort string = "e_port"

	// FcSwitchInlinePortsInlineArrayItemTypeFPort captures enum value "f_port"
	FcSwitchInlinePortsInlineArrayItemTypeFPort string = "f_port"

	// FcSwitchInlinePortsInlineArrayItemTypeFlPort captures enum value "fl_port"
	FcSwitchInlinePortsInlineArrayItemTypeFlPort string = "fl_port"

	// FcSwitchInlinePortsInlineArrayItemTypeFnlPort captures enum value "fnl_port"
	FcSwitchInlinePortsInlineArrayItemTypeFnlPort string = "fnl_port"

	// FcSwitchInlinePortsInlineArrayItemTypeFvPort captures enum value "fv_port"
	FcSwitchInlinePortsInlineArrayItemTypeFvPort string = "fv_port"

	// FcSwitchInlinePortsInlineArrayItemTypeNPort captures enum value "n_port"
	FcSwitchInlinePortsInlineArrayItemTypeNPort string = "n_port"

	// FcSwitchInlinePortsInlineArrayItemTypeNlPort captures enum value "nl_port"
	FcSwitchInlinePortsInlineArrayItemTypeNlPort string = "nl_port"

	// FcSwitchInlinePortsInlineArrayItemTypeNvPort captures enum value "nv_port"
	FcSwitchInlinePortsInlineArrayItemTypeNvPort string = "nv_port"

	// FcSwitchInlinePortsInlineArrayItemTypeNxPort captures enum value "nx_port"
	FcSwitchInlinePortsInlineArrayItemTypeNxPort string = "nx_port"

	// FcSwitchInlinePortsInlineArrayItemTypeSdPort captures enum value "sd_port"
	FcSwitchInlinePortsInlineArrayItemTypeSdPort string = "sd_port"

	// FcSwitchInlinePortsInlineArrayItemTypeTePort captures enum value "te_port"
	FcSwitchInlinePortsInlineArrayItemTypeTePort string = "te_port"

	// FcSwitchInlinePortsInlineArrayItemTypeTfPort captures enum value "tf_port"
	FcSwitchInlinePortsInlineArrayItemTypeTfPort string = "tf_port"

	// FcSwitchInlinePortsInlineArrayItemTypeTlPort captures enum value "tl_port"
	FcSwitchInlinePortsInlineArrayItemTypeTlPort string = "tl_port"

	// FcSwitchInlinePortsInlineArrayItemTypeTnpPort captures enum value "tnp_port"
	FcSwitchInlinePortsInlineArrayItemTypeTnpPort string = "tnp_port"

	// FcSwitchInlinePortsInlineArrayItemTypeNone captures enum value "none"
	FcSwitchInlinePortsInlineArrayItemTypeNone string = "none"
)

// prop value enum
func (m *FcSwitchInlinePortsInlineArrayItem) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fcSwitchInlinePortsInlineArrayItemTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FcSwitchInlinePortsInlineArrayItem) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this fc switch inline ports inline array item based on the context it is used
func (m *FcSwitchInlinePortsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachedDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSlot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwpn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchInlinePortsInlineArrayItem) contextValidateAttachedDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.AttachedDevice != nil {

		if swag.IsZero(m.AttachedDevice) { // not required
			return nil
		}

		if err := m.AttachedDevice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attached_device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attached_device")
			}
			return err
		}
	}

	return nil
}

func (m *FcSwitchInlinePortsInlineArrayItem) contextValidateSlot(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "slot", "body", m.Slot); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitchInlinePortsInlineArrayItem) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitchInlinePortsInlineArrayItem) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitchInlinePortsInlineArrayItem) contextValidateWwpn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "wwpn", "body", m.Wwpn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcSwitchInlinePortsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcSwitchInlinePortsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res FcSwitchInlinePortsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcSwitchInlinePortsInlineArrayItemInlineAttachedDevice The Fibre Channel (FC) device attached to the FC switch port.
//
// swagger:model fc_switch_inline_ports_inline_array_item_inline_attached_device
type FcSwitchInlinePortsInlineArrayItemInlineAttachedDevice struct {

	// The Fibre Channel port identifier of the attach device.
	//
	// Example: 70400
	// Read Only: true
	PortID *string `json:"port_id,omitempty"`

	// The world-wide port name (WWPN) of the attached device.
	//
	// Example: 50:0a:21:22:23:24:25:26
	// Read Only: true
	Wwpn *string `json:"wwpn,omitempty"`
}

// Validate validates this fc switch inline ports inline array item inline attached device
func (m *FcSwitchInlinePortsInlineArrayItemInlineAttachedDevice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this fc switch inline ports inline array item inline attached device based on the context it is used
func (m *FcSwitchInlinePortsInlineArrayItemInlineAttachedDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePortID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwpn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchInlinePortsInlineArrayItemInlineAttachedDevice) contextValidatePortID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "attached_device"+"."+"port_id", "body", m.PortID); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitchInlinePortsInlineArrayItemInlineAttachedDevice) contextValidateWwpn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "attached_device"+"."+"wwpn", "body", m.Wwpn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcSwitchInlinePortsInlineArrayItemInlineAttachedDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcSwitchInlinePortsInlineArrayItemInlineAttachedDevice) UnmarshalBinary(b []byte) error {
	var res FcSwitchInlinePortsInlineArrayItemInlineAttachedDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
