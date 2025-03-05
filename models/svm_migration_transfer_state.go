// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SvmMigrationTransferState svm migration transfer state
//
// swagger:model svm_migration_transfer_state
type SvmMigrationTransferState string

func NewSvmMigrationTransferState(value SvmMigrationTransferState) *SvmMigrationTransferState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SvmMigrationTransferState.
func (m SvmMigrationTransferState) Pointer() *SvmMigrationTransferState {
	return &m
}

const (

	// SvmMigrationTransferStateIdle captures enum value "Idle"
	SvmMigrationTransferStateIdle SvmMigrationTransferState = "Idle"

	// SvmMigrationTransferStateTransferring captures enum value "Transferring"
	SvmMigrationTransferStateTransferring SvmMigrationTransferState = "Transferring"

	// SvmMigrationTransferStateAborting captures enum value "Aborting"
	SvmMigrationTransferStateAborting SvmMigrationTransferState = "Aborting"

	// SvmMigrationTransferStateOutOfSync captures enum value "OutOfSync"
	SvmMigrationTransferStateOutOfSync SvmMigrationTransferState = "OutOfSync"

	// SvmMigrationTransferStateInSync captures enum value "InSync"
	SvmMigrationTransferStateInSync SvmMigrationTransferState = "InSync"

	// SvmMigrationTransferStateTransitioning captures enum value "Transitioning"
	SvmMigrationTransferStateTransitioning SvmMigrationTransferState = "Transitioning"

	// SvmMigrationTransferStateReadyForCutoverPreCommit captures enum value "ReadyForCutoverPreCommit"
	SvmMigrationTransferStateReadyForCutoverPreCommit SvmMigrationTransferState = "ReadyForCutoverPreCommit"

	// SvmMigrationTransferStateCutoverPreCommitting captures enum value "CutoverPreCommitting"
	SvmMigrationTransferStateCutoverPreCommitting SvmMigrationTransferState = "CutoverPreCommitting"

	// SvmMigrationTransferStateCuttingOver captures enum value "CuttingOver"
	SvmMigrationTransferStateCuttingOver SvmMigrationTransferState = "CuttingOver"
)

// for schema
var svmMigrationTransferStateEnum []interface{}

func init() {
	var res []SvmMigrationTransferState
	if err := json.Unmarshal([]byte(`["Idle","Transferring","Aborting","OutOfSync","InSync","Transitioning","ReadyForCutoverPreCommit","CutoverPreCommitting","CuttingOver"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		svmMigrationTransferStateEnum = append(svmMigrationTransferStateEnum, v)
	}
}

func (m SvmMigrationTransferState) validateSvmMigrationTransferStateEnum(path, location string, value SvmMigrationTransferState) error {
	if err := validate.EnumCase(path, location, value, svmMigrationTransferStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this svm migration transfer state
func (m SvmMigrationTransferState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSvmMigrationTransferStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this svm migration transfer state based on context it is used
func (m SvmMigrationTransferState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
