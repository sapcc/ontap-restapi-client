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

// MetroclusterOperation Data for a MetroCluster operation. REST: /api/cluster/metrocluster/operations
//
// swagger:model metrocluster_operation
type MetroclusterOperation struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// Additional information for the auto heal.
	// Example: MetroCluster switchover with auto heal completed successfully.
	// Read Only: true
	AdditionalInfo *string `json:"additional_info,omitempty"`

	// Command line executed with the options specified.
	// Example: metrocluster switchover
	// Read Only: true
	CommandLine *string `json:"command_line,omitempty"`

	// End Time
	// Example: 2016-03-10 22:35:16
	// Read Only: true
	// Format: date-time
	EndTime *strfmt.DateTime `json:"end_time,omitempty"`

	// List of errors in the operation.
	// Example: ["siteB (warning): Unable to prepare the partner cluster for a pending switchback operation. Reason: entry doesn't exist. Reboot the nodes in the partner cluster before using the \"metrocluster switchback\" command."]
	// Read Only: true
	MetroclusterOperationInlineErrors []*string `json:"errors,omitempty"`

	// node
	Node *MetroclusterOperationInlineNode `json:"node,omitempty"`

	// Start Time
	// Example: 2016-03-10 22:33:16
	// Read Only: true
	// Format: date-time
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`

	// Indicates the state of the operation.
	// Example: completed_with_warnings
	// Read Only: true
	// Enum: ["completed_with_manual_recovery_needed","completed_with_warnings","failed","in_progress","partially_successful","successful","unknown"]
	State *string `json:"state,omitempty"`

	// Name of the operation.
	// Example: switchover
	// Read Only: true
	// Enum: ["check","configure","connect","disconnect","dr_group_create","dr_group_delete","heal_aggr_auto","heal_aggregates","heal_root_aggr_auto","heal_root_aggregates","interface_create","interface_delete","interface_modify","ip_setup","ip_teardown","modify","switchback","switchback_continuation_agent","switchback_simulate","switchover","switchover_simulate","unconfigure","unconfigure_continuation_agent","unknown"]
	Type *string `json:"type,omitempty"`

	// Identifier for the operation.
	// Example: 11111111-2222-3333-4444-abcdefabcdef
	// Required: true
	// Read Only: true
	UUID *string `json:"uuid"`
}

// Validate validates this metrocluster operation
func (m *MetroclusterOperation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterOperation) validateLinks(formats strfmt.Registry) error {
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

func (m *MetroclusterOperation) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterOperation) validateNode(formats strfmt.Registry) error {
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

func (m *MetroclusterOperation) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var metroclusterOperationTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["completed_with_manual_recovery_needed","completed_with_warnings","failed","in_progress","partially_successful","successful","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metroclusterOperationTypeStatePropEnum = append(metroclusterOperationTypeStatePropEnum, v)
	}
}

const (

	// MetroclusterOperationStateCompletedWithManualRecoveryNeeded captures enum value "completed_with_manual_recovery_needed"
	MetroclusterOperationStateCompletedWithManualRecoveryNeeded string = "completed_with_manual_recovery_needed"

	// MetroclusterOperationStateCompletedWithWarnings captures enum value "completed_with_warnings"
	MetroclusterOperationStateCompletedWithWarnings string = "completed_with_warnings"

	// MetroclusterOperationStateFailed captures enum value "failed"
	MetroclusterOperationStateFailed string = "failed"

	// MetroclusterOperationStateInProgress captures enum value "in_progress"
	MetroclusterOperationStateInProgress string = "in_progress"

	// MetroclusterOperationStatePartiallySuccessful captures enum value "partially_successful"
	MetroclusterOperationStatePartiallySuccessful string = "partially_successful"

	// MetroclusterOperationStateSuccessful captures enum value "successful"
	MetroclusterOperationStateSuccessful string = "successful"

	// MetroclusterOperationStateUnknown captures enum value "unknown"
	MetroclusterOperationStateUnknown string = "unknown"
)

// prop value enum
func (m *MetroclusterOperation) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metroclusterOperationTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MetroclusterOperation) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

var metroclusterOperationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["check","configure","connect","disconnect","dr_group_create","dr_group_delete","heal_aggr_auto","heal_aggregates","heal_root_aggr_auto","heal_root_aggregates","interface_create","interface_delete","interface_modify","ip_setup","ip_teardown","modify","switchback","switchback_continuation_agent","switchback_simulate","switchover","switchover_simulate","unconfigure","unconfigure_continuation_agent","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metroclusterOperationTypeTypePropEnum = append(metroclusterOperationTypeTypePropEnum, v)
	}
}

const (

	// MetroclusterOperationTypeCheck captures enum value "check"
	MetroclusterOperationTypeCheck string = "check"

	// MetroclusterOperationTypeConfigure captures enum value "configure"
	MetroclusterOperationTypeConfigure string = "configure"

	// MetroclusterOperationTypeConnect captures enum value "connect"
	MetroclusterOperationTypeConnect string = "connect"

	// MetroclusterOperationTypeDisconnect captures enum value "disconnect"
	MetroclusterOperationTypeDisconnect string = "disconnect"

	// MetroclusterOperationTypeDrGroupCreate captures enum value "dr_group_create"
	MetroclusterOperationTypeDrGroupCreate string = "dr_group_create"

	// MetroclusterOperationTypeDrGroupDelete captures enum value "dr_group_delete"
	MetroclusterOperationTypeDrGroupDelete string = "dr_group_delete"

	// MetroclusterOperationTypeHealAggrAuto captures enum value "heal_aggr_auto"
	MetroclusterOperationTypeHealAggrAuto string = "heal_aggr_auto"

	// MetroclusterOperationTypeHealAggregates captures enum value "heal_aggregates"
	MetroclusterOperationTypeHealAggregates string = "heal_aggregates"

	// MetroclusterOperationTypeHealRootAggrAuto captures enum value "heal_root_aggr_auto"
	MetroclusterOperationTypeHealRootAggrAuto string = "heal_root_aggr_auto"

	// MetroclusterOperationTypeHealRootAggregates captures enum value "heal_root_aggregates"
	MetroclusterOperationTypeHealRootAggregates string = "heal_root_aggregates"

	// MetroclusterOperationTypeInterfaceCreate captures enum value "interface_create"
	MetroclusterOperationTypeInterfaceCreate string = "interface_create"

	// MetroclusterOperationTypeInterfaceDelete captures enum value "interface_delete"
	MetroclusterOperationTypeInterfaceDelete string = "interface_delete"

	// MetroclusterOperationTypeInterfaceModify captures enum value "interface_modify"
	MetroclusterOperationTypeInterfaceModify string = "interface_modify"

	// MetroclusterOperationTypeIPSetup captures enum value "ip_setup"
	MetroclusterOperationTypeIPSetup string = "ip_setup"

	// MetroclusterOperationTypeIPTeardown captures enum value "ip_teardown"
	MetroclusterOperationTypeIPTeardown string = "ip_teardown"

	// MetroclusterOperationTypeModify captures enum value "modify"
	MetroclusterOperationTypeModify string = "modify"

	// MetroclusterOperationTypeSwitchback captures enum value "switchback"
	MetroclusterOperationTypeSwitchback string = "switchback"

	// MetroclusterOperationTypeSwitchbackContinuationAgent captures enum value "switchback_continuation_agent"
	MetroclusterOperationTypeSwitchbackContinuationAgent string = "switchback_continuation_agent"

	// MetroclusterOperationTypeSwitchbackSimulate captures enum value "switchback_simulate"
	MetroclusterOperationTypeSwitchbackSimulate string = "switchback_simulate"

	// MetroclusterOperationTypeSwitchover captures enum value "switchover"
	MetroclusterOperationTypeSwitchover string = "switchover"

	// MetroclusterOperationTypeSwitchoverSimulate captures enum value "switchover_simulate"
	MetroclusterOperationTypeSwitchoverSimulate string = "switchover_simulate"

	// MetroclusterOperationTypeUnconfigure captures enum value "unconfigure"
	MetroclusterOperationTypeUnconfigure string = "unconfigure"

	// MetroclusterOperationTypeUnconfigureContinuationAgent captures enum value "unconfigure_continuation_agent"
	MetroclusterOperationTypeUnconfigureContinuationAgent string = "unconfigure_continuation_agent"

	// MetroclusterOperationTypeUnknown captures enum value "unknown"
	MetroclusterOperationTypeUnknown string = "unknown"
)

// prop value enum
func (m *MetroclusterOperation) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metroclusterOperationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MetroclusterOperation) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterOperation) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this metrocluster operation based on the context it is used
func (m *MetroclusterOperation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAdditionalInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommandLine(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetroclusterOperationInlineErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterOperation) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MetroclusterOperation) contextValidateAdditionalInfo(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "additional_info", "body", m.AdditionalInfo); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterOperation) contextValidateCommandLine(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "command_line", "body", m.CommandLine); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterOperation) contextValidateEndTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "end_time", "body", m.EndTime); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterOperation) contextValidateMetroclusterOperationInlineErrors(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "errors", "body", []*string(m.MetroclusterOperationInlineErrors)); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterOperation) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MetroclusterOperation) contextValidateStartTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "start_time", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterOperation) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterOperation) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterOperation) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterOperation) UnmarshalBinary(b []byte) error {
	var res MetroclusterOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterOperationInlineNode Node from where the command is executed.
//
// swagger:model metrocluster_operation_inline_node
type MetroclusterOperationInlineNode struct {

	// links
	Links *MetroclusterOperationInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this metrocluster operation inline node
func (m *MetroclusterOperationInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterOperationInlineNode) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this metrocluster operation inline node based on the context it is used
func (m *MetroclusterOperationInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterOperationInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MetroclusterOperationInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterOperationInlineNode) UnmarshalBinary(b []byte) error {
	var res MetroclusterOperationInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterOperationInlineNodeInlineLinks metrocluster operation inline node inline links
//
// swagger:model metrocluster_operation_inline_node_inline__links
type MetroclusterOperationInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this metrocluster operation inline node inline links
func (m *MetroclusterOperationInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterOperationInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this metrocluster operation inline node inline links based on the context it is used
func (m *MetroclusterOperationInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterOperationInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MetroclusterOperationInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterOperationInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res MetroclusterOperationInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
