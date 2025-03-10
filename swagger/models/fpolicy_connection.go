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

// FpolicyConnection Displays the connection status information of the FPolicy server.
//
// swagger:model fpolicy_connection
type FpolicyConnection struct {

	// links
	Links *FpolicyConnectionInlineLinks `json:"_links,omitempty"`

	// disconnected reason
	DisconnectedReason *FpolicyConnectionInlineDisconnectedReason `json:"disconnected_reason,omitempty"`

	// node
	Node *FpolicyConnectionInlineNode `json:"node,omitempty"`

	// policy
	// Read Only: true
	Policy *FpolicyPolicyReference `json:"policy,omitempty"`

	// IP address of the FPolicy server.
	// Example: 10.132.145.20
	// Read Only: true
	Server *string `json:"server,omitempty"`

	// Unique session ID associated with each connection to the FPolicy server and it can be used to identify
	// the established connection.
	//
	// Example: 5224ec64-b336-11eb-841c-0050568e14c2
	// Read Only: true
	SessionUUID *string `json:"session_uuid,omitempty"`

	// Specifies the FPolicy server connection state indicating if it is in the connected or disconnected state.
	// The following is a list of the possible states:
	// * connected                 - Connected
	// * disconnected              - Disconnected
	//
	// Example: connected
	// Enum: ["connected","disconnected"]
	State *string `json:"state,omitempty"`

	// svm
	Svm *FpolicyConnectionInlineSvm `json:"svm,omitempty"`

	// FPolicy server type. The possible values are:
	//   * primary - Primary server
	//   * secondary  - Secondary server
	//
	// Example: primary
	// Read Only: true
	// Enum: ["primary","secondary"]
	Type *string `json:"type,omitempty"`

	// Specifies the time at which FPolicy server is connected or disconnected.
	// Example: 2019-06-12 15:00:16
	// Read Only: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this fpolicy connection
func (m *FpolicyConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisconnectedReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyConnection) validateLinks(formats strfmt.Registry) error {
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

func (m *FpolicyConnection) validateDisconnectedReason(formats strfmt.Registry) error {
	if swag.IsZero(m.DisconnectedReason) { // not required
		return nil
	}

	if m.DisconnectedReason != nil {
		if err := m.DisconnectedReason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disconnected_reason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disconnected_reason")
			}
			return err
		}
	}

	return nil
}

func (m *FpolicyConnection) validateNode(formats strfmt.Registry) error {
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

func (m *FpolicyConnection) validatePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	if m.Policy != nil {
		if err := m.Policy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

var fpolicyConnectionTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["connected","disconnected"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fpolicyConnectionTypeStatePropEnum = append(fpolicyConnectionTypeStatePropEnum, v)
	}
}

const (

	// FpolicyConnectionStateConnected captures enum value "connected"
	FpolicyConnectionStateConnected string = "connected"

	// FpolicyConnectionStateDisconnected captures enum value "disconnected"
	FpolicyConnectionStateDisconnected string = "disconnected"
)

// prop value enum
func (m *FpolicyConnection) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fpolicyConnectionTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FpolicyConnection) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

func (m *FpolicyConnection) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

var fpolicyConnectionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["primary","secondary"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fpolicyConnectionTypeTypePropEnum = append(fpolicyConnectionTypeTypePropEnum, v)
	}
}

const (

	// FpolicyConnectionTypePrimary captures enum value "primary"
	FpolicyConnectionTypePrimary string = "primary"

	// FpolicyConnectionTypeSecondary captures enum value "secondary"
	FpolicyConnectionTypeSecondary string = "secondary"
)

// prop value enum
func (m *FpolicyConnection) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fpolicyConnectionTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FpolicyConnection) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *FpolicyConnection) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this fpolicy connection based on the context it is used
func (m *FpolicyConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisconnectedReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSessionUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
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

func (m *FpolicyConnection) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *FpolicyConnection) contextValidateDisconnectedReason(ctx context.Context, formats strfmt.Registry) error {

	if m.DisconnectedReason != nil {

		if swag.IsZero(m.DisconnectedReason) { // not required
			return nil
		}

		if err := m.DisconnectedReason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disconnected_reason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disconnected_reason")
			}
			return err
		}
	}

	return nil
}

func (m *FpolicyConnection) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *FpolicyConnection) contextValidatePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.Policy != nil {

		if swag.IsZero(m.Policy) { // not required
			return nil
		}

		if err := m.Policy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

func (m *FpolicyConnection) contextValidateServer(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "server", "body", m.Server); err != nil {
		return err
	}

	return nil
}

func (m *FpolicyConnection) contextValidateSessionUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "session_uuid", "body", m.SessionUUID); err != nil {
		return err
	}

	return nil
}

func (m *FpolicyConnection) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {

		if swag.IsZero(m.Svm) { // not required
			return nil
		}

		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

func (m *FpolicyConnection) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *FpolicyConnection) contextValidateUpdateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "update_time", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyConnection) UnmarshalBinary(b []byte) error {
	var res FpolicyConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FpolicyConnectionInlineDisconnectedReason Indicates the reason for FPolicy server disconnection.
//
// swagger:model fpolicy_connection_inline_disconnected_reason
type FpolicyConnectionInlineDisconnectedReason struct {

	// Reason ID for the FPolicy Server disconnection.
	// Example: 9307
	Code *int64 `json:"code,omitempty"`

	// FPolicy server reason for disconnection message.
	// Example: TCP Connection to FPolicy server failed.
	Message *string `json:"message,omitempty"`
}

// Validate validates this fpolicy connection inline disconnected reason
func (m *FpolicyConnectionInlineDisconnectedReason) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this fpolicy connection inline disconnected reason based on the context it is used
func (m *FpolicyConnectionInlineDisconnectedReason) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyConnectionInlineDisconnectedReason) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyConnectionInlineDisconnectedReason) UnmarshalBinary(b []byte) error {
	var res FpolicyConnectionInlineDisconnectedReason
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FpolicyConnectionInlineLinks fpolicy connection inline links
//
// swagger:model fpolicy_connection_inline__links
type FpolicyConnectionInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this fpolicy connection inline links
func (m *FpolicyConnectionInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyConnectionInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this fpolicy connection inline links based on the context it is used
func (m *FpolicyConnectionInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyConnectionInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *FpolicyConnectionInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyConnectionInlineLinks) UnmarshalBinary(b []byte) error {
	var res FpolicyConnectionInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FpolicyConnectionInlineNode fpolicy connection inline node
//
// swagger:model fpolicy_connection_inline_node
type FpolicyConnectionInlineNode struct {

	// links
	Links *FpolicyConnectionInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this fpolicy connection inline node
func (m *FpolicyConnectionInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyConnectionInlineNode) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this fpolicy connection inline node based on the context it is used
func (m *FpolicyConnectionInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyConnectionInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *FpolicyConnectionInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyConnectionInlineNode) UnmarshalBinary(b []byte) error {
	var res FpolicyConnectionInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FpolicyConnectionInlineNodeInlineLinks fpolicy connection inline node inline links
//
// swagger:model fpolicy_connection_inline_node_inline__links
type FpolicyConnectionInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this fpolicy connection inline node inline links
func (m *FpolicyConnectionInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyConnectionInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this fpolicy connection inline node inline links based on the context it is used
func (m *FpolicyConnectionInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyConnectionInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *FpolicyConnectionInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyConnectionInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res FpolicyConnectionInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FpolicyConnectionInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model fpolicy_connection_inline_svm
type FpolicyConnectionInlineSvm struct {

	// links
	Links *FpolicyConnectionInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this fpolicy connection inline svm
func (m *FpolicyConnectionInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyConnectionInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fpolicy connection inline svm based on the context it is used
func (m *FpolicyConnectionInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyConnectionInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyConnectionInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyConnectionInlineSvm) UnmarshalBinary(b []byte) error {
	var res FpolicyConnectionInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FpolicyConnectionInlineSvmInlineLinks fpolicy connection inline svm inline links
//
// swagger:model fpolicy_connection_inline_svm_inline__links
type FpolicyConnectionInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this fpolicy connection inline svm inline links
func (m *FpolicyConnectionInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyConnectionInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fpolicy connection inline svm inline links based on the context it is used
func (m *FpolicyConnectionInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyConnectionInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyConnectionInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyConnectionInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res FpolicyConnectionInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
