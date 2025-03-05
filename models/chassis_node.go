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

// ChassisNode List of nodes in chassis.
//
// swagger:model chassis_node
type ChassisNode struct {

	// links
	Links *ChassisNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// pcis
	Pcis *ChassisNodeInlinePcis `json:"pcis,omitempty"`

	// The position of the node in the chassis, when viewed from the rear of the system.
	// Example: top
	// Enum: ["top","bottom","left","right","centre","unknown"]
	Position *string `json:"position,omitempty"`

	// usbs
	Usbs *ChassisNodeInlineUsbs `json:"usbs,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this chassis node
func (m *ChassisNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePcis(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsbs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisNode) validateLinks(formats strfmt.Registry) error {
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

func (m *ChassisNode) validatePcis(formats strfmt.Registry) error {
	if swag.IsZero(m.Pcis) { // not required
		return nil
	}

	if m.Pcis != nil {
		if err := m.Pcis.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pcis")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pcis")
			}
			return err
		}
	}

	return nil
}

var chassisNodeTypePositionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["top","bottom","left","right","centre","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		chassisNodeTypePositionPropEnum = append(chassisNodeTypePositionPropEnum, v)
	}
}

const (

	// ChassisNodePositionTop captures enum value "top"
	ChassisNodePositionTop string = "top"

	// ChassisNodePositionBottom captures enum value "bottom"
	ChassisNodePositionBottom string = "bottom"

	// ChassisNodePositionLeft captures enum value "left"
	ChassisNodePositionLeft string = "left"

	// ChassisNodePositionRight captures enum value "right"
	ChassisNodePositionRight string = "right"

	// ChassisNodePositionCentre captures enum value "centre"
	ChassisNodePositionCentre string = "centre"

	// ChassisNodePositionUnknown captures enum value "unknown"
	ChassisNodePositionUnknown string = "unknown"
)

// prop value enum
func (m *ChassisNode) validatePositionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, chassisNodeTypePositionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ChassisNode) validatePosition(formats strfmt.Registry) error {
	if swag.IsZero(m.Position) { // not required
		return nil
	}

	// value enum
	if err := m.validatePositionEnum("position", "body", *m.Position); err != nil {
		return err
	}

	return nil
}

func (m *ChassisNode) validateUsbs(formats strfmt.Registry) error {
	if swag.IsZero(m.Usbs) { // not required
		return nil
	}

	if m.Usbs != nil {
		if err := m.Usbs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usbs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usbs")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this chassis node based on the context it is used
func (m *ChassisNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePcis(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsbs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ChassisNode) contextValidatePcis(ctx context.Context, formats strfmt.Registry) error {

	if m.Pcis != nil {

		if swag.IsZero(m.Pcis) { // not required
			return nil
		}

		if err := m.Pcis.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pcis")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pcis")
			}
			return err
		}
	}

	return nil
}

func (m *ChassisNode) contextValidateUsbs(ctx context.Context, formats strfmt.Registry) error {

	if m.Usbs != nil {

		if swag.IsZero(m.Usbs) { // not required
			return nil
		}

		if err := m.Usbs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usbs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usbs")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChassisNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChassisNode) UnmarshalBinary(b []byte) error {
	var res ChassisNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ChassisNodeInlineLinks chassis node inline links
//
// swagger:model chassis_node_inline__links
type ChassisNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this chassis node inline links
func (m *ChassisNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this chassis node inline links based on the context it is used
func (m *ChassisNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ChassisNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChassisNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res ChassisNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ChassisNodeInlinePcis chassis node inline pcis
//
// swagger:model chassis_node_inline_pcis
type ChassisNodeInlinePcis struct {

	// cards
	Cards []*ChassisNodePcisCardsItems0 `json:"cards,omitempty"`
}

// Validate validates this chassis node inline pcis
func (m *ChassisNodeInlinePcis) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCards(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisNodeInlinePcis) validateCards(formats strfmt.Registry) error {
	if swag.IsZero(m.Cards) { // not required
		return nil
	}

	for i := 0; i < len(m.Cards); i++ {
		if swag.IsZero(m.Cards[i]) { // not required
			continue
		}

		if m.Cards[i] != nil {
			if err := m.Cards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pcis" + "." + "cards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pcis" + "." + "cards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this chassis node inline pcis based on the context it is used
func (m *ChassisNodeInlinePcis) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCards(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisNodeInlinePcis) contextValidateCards(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Cards); i++ {

		if m.Cards[i] != nil {

			if swag.IsZero(m.Cards[i]) { // not required
				return nil
			}

			if err := m.Cards[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pcis" + "." + "cards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pcis" + "." + "cards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChassisNodeInlinePcis) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChassisNodeInlinePcis) UnmarshalBinary(b []byte) error {
	var res ChassisNodeInlinePcis
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ChassisNodePcisCardsItems0 chassis node pcis cards items0
//
// swagger:model ChassisNodePcisCardsItems0
type ChassisNodePcisCardsItems0 struct {

	// The description of the PCI card.
	// Example: Intel Lewisburg series chipset SATA Controller
	Device *string `json:"device,omitempty"`

	// The info string from the device driver of the PCI card.
	// Example: Additional Info: 0 (0xaaf00000)   SHM2S86Q120GLM22NP FW1146 114473MB 512B/sect (SPG190108GW)
	Info *string `json:"info,omitempty"`

	// The slot where the PCI card is placed. This can sometimes take the form of "6-1" to indicate slot and subslot.
	// Example: 0
	Slot *string `json:"slot,omitempty"`
}

// Validate validates this chassis node pcis cards items0
func (m *ChassisNodePcisCardsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this chassis node pcis cards items0 based on context it is used
func (m *ChassisNodePcisCardsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChassisNodePcisCardsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChassisNodePcisCardsItems0) UnmarshalBinary(b []byte) error {
	var res ChassisNodePcisCardsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ChassisNodeInlineUsbs The status of the USB ports on the controller.
//
// swagger:model chassis_node_inline_usbs
type ChassisNodeInlineUsbs struct {

	// Indicates whether or not the USB ports are enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// ports
	Ports []*ChassisNodeUsbsPortsItems0 `json:"ports,omitempty"`

	// Indicates whether or not USB ports are supported on the current platform.
	Supported *bool `json:"supported,omitempty"`
}

// Validate validates this chassis node inline usbs
func (m *ChassisNodeInlineUsbs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisNodeInlineUsbs) validatePorts(formats strfmt.Registry) error {
	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	for i := 0; i < len(m.Ports); i++ {
		if swag.IsZero(m.Ports[i]) { // not required
			continue
		}

		if m.Ports[i] != nil {
			if err := m.Ports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usbs" + "." + "ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usbs" + "." + "ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this chassis node inline usbs based on the context it is used
func (m *ChassisNodeInlineUsbs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisNodeInlineUsbs) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ports); i++ {

		if m.Ports[i] != nil {

			if swag.IsZero(m.Ports[i]) { // not required
				return nil
			}

			if err := m.Ports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usbs" + "." + "ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usbs" + "." + "ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChassisNodeInlineUsbs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChassisNodeInlineUsbs) UnmarshalBinary(b []byte) error {
	var res ChassisNodeInlineUsbs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ChassisNodeUsbsPortsItems0 chassis node usbs ports items0
//
// swagger:model ChassisNodeUsbsPortsItems0
type ChassisNodeUsbsPortsItems0 struct {

	// Indicates whether or not the USB port has a device connected to it.
	Connected *bool `json:"connected,omitempty"`
}

// Validate validates this chassis node usbs ports items0
func (m *ChassisNodeUsbsPortsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this chassis node usbs ports items0 based on context it is used
func (m *ChassisNodeUsbsPortsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChassisNodeUsbsPortsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChassisNodeUsbsPortsItems0) UnmarshalBinary(b []byte) error {
	var res ChassisNodeUsbsPortsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
