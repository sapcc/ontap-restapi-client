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

// MetroclusterDiagConnection metrocluster diag connection
//
// swagger:model metrocluster_diag_connection
type MetroclusterDiagConnection struct {

	// destination address
	// Read Only: true
	DestinationAddress *string `json:"destination_address,omitempty"`

	// partner
	Partner *MetroclusterDiagConnectionInlinePartner `json:"partner,omitempty"`

	// port
	// Read Only: true
	Port *string `json:"port,omitempty"`

	// Result of the diagnostic operation on this component.
	// Read Only: true
	// Enum: ["ok","warning","not_run","not_applicable"]
	Result *string `json:"result,omitempty"`

	// source address
	// Read Only: true
	SourceAddress *string `json:"source_address,omitempty"`

	// state
	// Read Only: true
	// Enum: ["disconnected","completed","pinging_partner_nodes","enabling_ip_ports","connecting_to_storage","disconnecting_from_storage","disabling_ip_ports","configuring_ip_addresses","updating_node_roles","connecting_to_mediator","disconnecting_from_mediator"]
	State *string `json:"state,omitempty"`
}

// Validate validates this metrocluster diag connection
func (m *MetroclusterDiagConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePartner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
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

func (m *MetroclusterDiagConnection) validatePartner(formats strfmt.Registry) error {
	if swag.IsZero(m.Partner) { // not required
		return nil
	}

	if m.Partner != nil {
		if err := m.Partner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partner")
			}
			return err
		}
	}

	return nil
}

var metroclusterDiagConnectionTypeResultPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","warning","not_run","not_applicable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metroclusterDiagConnectionTypeResultPropEnum = append(metroclusterDiagConnectionTypeResultPropEnum, v)
	}
}

const (

	// MetroclusterDiagConnectionResultOk captures enum value "ok"
	MetroclusterDiagConnectionResultOk string = "ok"

	// MetroclusterDiagConnectionResultWarning captures enum value "warning"
	MetroclusterDiagConnectionResultWarning string = "warning"

	// MetroclusterDiagConnectionResultNotRun captures enum value "not_run"
	MetroclusterDiagConnectionResultNotRun string = "not_run"

	// MetroclusterDiagConnectionResultNotApplicable captures enum value "not_applicable"
	MetroclusterDiagConnectionResultNotApplicable string = "not_applicable"
)

// prop value enum
func (m *MetroclusterDiagConnection) validateResultEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metroclusterDiagConnectionTypeResultPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MetroclusterDiagConnection) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(m.Result) { // not required
		return nil
	}

	// value enum
	if err := m.validateResultEnum("result", "body", *m.Result); err != nil {
		return err
	}

	return nil
}

var metroclusterDiagConnectionTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["disconnected","completed","pinging_partner_nodes","enabling_ip_ports","connecting_to_storage","disconnecting_from_storage","disabling_ip_ports","configuring_ip_addresses","updating_node_roles","connecting_to_mediator","disconnecting_from_mediator"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metroclusterDiagConnectionTypeStatePropEnum = append(metroclusterDiagConnectionTypeStatePropEnum, v)
	}
}

const (

	// MetroclusterDiagConnectionStateDisconnected captures enum value "disconnected"
	MetroclusterDiagConnectionStateDisconnected string = "disconnected"

	// MetroclusterDiagConnectionStateCompleted captures enum value "completed"
	MetroclusterDiagConnectionStateCompleted string = "completed"

	// MetroclusterDiagConnectionStatePingingPartnerNodes captures enum value "pinging_partner_nodes"
	MetroclusterDiagConnectionStatePingingPartnerNodes string = "pinging_partner_nodes"

	// MetroclusterDiagConnectionStateEnablingIPPorts captures enum value "enabling_ip_ports"
	MetroclusterDiagConnectionStateEnablingIPPorts string = "enabling_ip_ports"

	// MetroclusterDiagConnectionStateConnectingToStorage captures enum value "connecting_to_storage"
	MetroclusterDiagConnectionStateConnectingToStorage string = "connecting_to_storage"

	// MetroclusterDiagConnectionStateDisconnectingFromStorage captures enum value "disconnecting_from_storage"
	MetroclusterDiagConnectionStateDisconnectingFromStorage string = "disconnecting_from_storage"

	// MetroclusterDiagConnectionStateDisablingIPPorts captures enum value "disabling_ip_ports"
	MetroclusterDiagConnectionStateDisablingIPPorts string = "disabling_ip_ports"

	// MetroclusterDiagConnectionStateConfiguringIPAddresses captures enum value "configuring_ip_addresses"
	MetroclusterDiagConnectionStateConfiguringIPAddresses string = "configuring_ip_addresses"

	// MetroclusterDiagConnectionStateUpdatingNodeRoles captures enum value "updating_node_roles"
	MetroclusterDiagConnectionStateUpdatingNodeRoles string = "updating_node_roles"

	// MetroclusterDiagConnectionStateConnectingToMediator captures enum value "connecting_to_mediator"
	MetroclusterDiagConnectionStateConnectingToMediator string = "connecting_to_mediator"

	// MetroclusterDiagConnectionStateDisconnectingFromMediator captures enum value "disconnecting_from_mediator"
	MetroclusterDiagConnectionStateDisconnectingFromMediator string = "disconnecting_from_mediator"
)

// prop value enum
func (m *MetroclusterDiagConnection) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metroclusterDiagConnectionTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MetroclusterDiagConnection) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this metrocluster diag connection based on the context it is used
func (m *MetroclusterDiagConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestinationAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePartner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterDiagConnection) contextValidateDestinationAddress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "destination_address", "body", m.DestinationAddress); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterDiagConnection) contextValidatePartner(ctx context.Context, formats strfmt.Registry) error {

	if m.Partner != nil {

		if swag.IsZero(m.Partner) { // not required
			return nil
		}

		if err := m.Partner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partner")
			}
			return err
		}
	}

	return nil
}

func (m *MetroclusterDiagConnection) contextValidatePort(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterDiagConnection) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "result", "body", m.Result); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterDiagConnection) contextValidateSourceAddress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "source_address", "body", m.SourceAddress); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterDiagConnection) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterDiagConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterDiagConnection) UnmarshalBinary(b []byte) error {
	var res MetroclusterDiagConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterDiagConnectionInlinePartner metrocluster diag connection inline partner
//
// swagger:model metrocluster_diag_connection_inline_partner
type MetroclusterDiagConnectionInlinePartner struct {

	// node
	Node *MetroclusterDiagConnectionInlinePartnerInlineNode `json:"node,omitempty"`

	// type
	// Enum: ["ha","dr","aux"]
	Type *string `json:"type,omitempty"`
}

// Validate validates this metrocluster diag connection inline partner
func (m *MetroclusterDiagConnectionInlinePartner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNode(formats); err != nil {
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

func (m *MetroclusterDiagConnectionInlinePartner) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partner" + "." + "node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partner" + "." + "node")
			}
			return err
		}
	}

	return nil
}

var metroclusterDiagConnectionInlinePartnerTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ha","dr","aux"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metroclusterDiagConnectionInlinePartnerTypeTypePropEnum = append(metroclusterDiagConnectionInlinePartnerTypeTypePropEnum, v)
	}
}

const (

	// MetroclusterDiagConnectionInlinePartnerTypeHa captures enum value "ha"
	MetroclusterDiagConnectionInlinePartnerTypeHa string = "ha"

	// MetroclusterDiagConnectionInlinePartnerTypeDr captures enum value "dr"
	MetroclusterDiagConnectionInlinePartnerTypeDr string = "dr"

	// MetroclusterDiagConnectionInlinePartnerTypeAux captures enum value "aux"
	MetroclusterDiagConnectionInlinePartnerTypeAux string = "aux"
)

// prop value enum
func (m *MetroclusterDiagConnectionInlinePartner) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metroclusterDiagConnectionInlinePartnerTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MetroclusterDiagConnectionInlinePartner) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("partner"+"."+"type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this metrocluster diag connection inline partner based on the context it is used
func (m *MetroclusterDiagConnectionInlinePartner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterDiagConnectionInlinePartner) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {

		if swag.IsZero(m.Node) { // not required
			return nil
		}

		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partner" + "." + "node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partner" + "." + "node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterDiagConnectionInlinePartner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterDiagConnectionInlinePartner) UnmarshalBinary(b []byte) error {
	var res MetroclusterDiagConnectionInlinePartner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterDiagConnectionInlinePartnerInlineNode metrocluster diag connection inline partner inline node
//
// swagger:model metrocluster_diag_connection_inline_partner_inline_node
type MetroclusterDiagConnectionInlinePartnerInlineNode struct {

	// links
	Links *MetroclusterDiagConnectionInlinePartnerInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this metrocluster diag connection inline partner inline node
func (m *MetroclusterDiagConnectionInlinePartnerInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterDiagConnectionInlinePartnerInlineNode) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partner" + "." + "node" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partner" + "." + "node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metrocluster diag connection inline partner inline node based on the context it is used
func (m *MetroclusterDiagConnectionInlinePartnerInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterDiagConnectionInlinePartnerInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partner" + "." + "node" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partner" + "." + "node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterDiagConnectionInlinePartnerInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterDiagConnectionInlinePartnerInlineNode) UnmarshalBinary(b []byte) error {
	var res MetroclusterDiagConnectionInlinePartnerInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterDiagConnectionInlinePartnerInlineNodeInlineLinks metrocluster diag connection inline partner inline node inline links
//
// swagger:model metrocluster_diag_connection_inline_partner_inline_node_inline__links
type MetroclusterDiagConnectionInlinePartnerInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this metrocluster diag connection inline partner inline node inline links
func (m *MetroclusterDiagConnectionInlinePartnerInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterDiagConnectionInlinePartnerInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partner" + "." + "node" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partner" + "." + "node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metrocluster diag connection inline partner inline node inline links based on the context it is used
func (m *MetroclusterDiagConnectionInlinePartnerInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterDiagConnectionInlinePartnerInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partner" + "." + "node" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partner" + "." + "node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterDiagConnectionInlinePartnerInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterDiagConnectionInlinePartnerInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res MetroclusterDiagConnectionInlinePartnerInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
