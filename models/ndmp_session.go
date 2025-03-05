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

// NdmpSession ndmp session
//
// swagger:model ndmp_session
type NdmpSession struct {

	// links
	Links *NdmpSessionInlineLinks `json:"_links,omitempty"`

	// Indicates the NDMP backup engine.
	// Enum: ["dump","smtape"]
	BackupEngine *string `json:"backup_engine,omitempty"`

	// Indicates the NDMP client address.
	ClientAddress *string `json:"client_address,omitempty"`

	// Indicates the NDMP client port.
	ClientPort *int64 `json:"client_port,omitempty"`

	// Information about the NDMP data server.
	Data *NdmpData `json:"data,omitempty"`

	// Indicates the NDMP backup or restore path.
	// Example: /vserver1/vol1
	DataPath *string `json:"data_path,omitempty"`

	// NDMP session identifier.
	ID *string `json:"id,omitempty"`

	// Information about the NDMP mover.
	Mover *NdmpMover `json:"mover,omitempty"`

	// node
	Node *NdmpSessionInlineNode `json:"node,omitempty"`

	// Information about the NDMP SCSI server.
	Scsi *NdmpScsi `json:"scsi,omitempty"`

	// Indicates the NDMP local address on which connection was established.
	SourceAddress *string `json:"source_address,omitempty"`

	// svm
	Svm *NdmpSessionInlineSvm `json:"svm,omitempty"`

	// Indicates the NDMP tape device.
	// Example: nrst0a
	TapeDevice *string `json:"tape_device,omitempty"`

	// Indicates the NDMP tape device mode of operation.
	// Example: write
	TapeMode *NdmpMoverMode `json:"tape_mode,omitempty"`

	// The NDMP node or SVM UUID based on whether NDMP is operating in node-scope or SVM-scope mode.
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this ndmp session
func (m *NdmpSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupEngine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMover(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScsi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTapeMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NdmpSession) validateLinks(formats strfmt.Registry) error {
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

var ndmpSessionTypeBackupEnginePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dump","smtape"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ndmpSessionTypeBackupEnginePropEnum = append(ndmpSessionTypeBackupEnginePropEnum, v)
	}
}

const (

	// NdmpSessionBackupEngineDump captures enum value "dump"
	NdmpSessionBackupEngineDump string = "dump"

	// NdmpSessionBackupEngineSmtape captures enum value "smtape"
	NdmpSessionBackupEngineSmtape string = "smtape"
)

// prop value enum
func (m *NdmpSession) validateBackupEngineEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ndmpSessionTypeBackupEnginePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NdmpSession) validateBackupEngine(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupEngine) { // not required
		return nil
	}

	// value enum
	if err := m.validateBackupEngineEnum("backup_engine", "body", *m.BackupEngine); err != nil {
		return err
	}

	return nil
}

func (m *NdmpSession) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *NdmpSession) validateMover(formats strfmt.Registry) error {
	if swag.IsZero(m.Mover) { // not required
		return nil
	}

	if m.Mover != nil {
		if err := m.Mover.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mover")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mover")
			}
			return err
		}
	}

	return nil
}

func (m *NdmpSession) validateNode(formats strfmt.Registry) error {
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

func (m *NdmpSession) validateScsi(formats strfmt.Registry) error {
	if swag.IsZero(m.Scsi) { // not required
		return nil
	}

	if m.Scsi != nil {
		if err := m.Scsi.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scsi")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scsi")
			}
			return err
		}
	}

	return nil
}

func (m *NdmpSession) validateSvm(formats strfmt.Registry) error {
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

func (m *NdmpSession) validateTapeMode(formats strfmt.Registry) error {
	if swag.IsZero(m.TapeMode) { // not required
		return nil
	}

	if m.TapeMode != nil {
		if err := m.TapeMode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tape_mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tape_mode")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ndmp session based on the context it is used
func (m *NdmpSession) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMover(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScsi(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTapeMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NdmpSession) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NdmpSession) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {

		if swag.IsZero(m.Data) { // not required
			return nil
		}

		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *NdmpSession) contextValidateMover(ctx context.Context, formats strfmt.Registry) error {

	if m.Mover != nil {

		if swag.IsZero(m.Mover) { // not required
			return nil
		}

		if err := m.Mover.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mover")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mover")
			}
			return err
		}
	}

	return nil
}

func (m *NdmpSession) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NdmpSession) contextValidateScsi(ctx context.Context, formats strfmt.Registry) error {

	if m.Scsi != nil {

		if swag.IsZero(m.Scsi) { // not required
			return nil
		}

		if err := m.Scsi.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scsi")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scsi")
			}
			return err
		}
	}

	return nil
}

func (m *NdmpSession) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NdmpSession) contextValidateTapeMode(ctx context.Context, formats strfmt.Registry) error {

	if m.TapeMode != nil {

		if swag.IsZero(m.TapeMode) { // not required
			return nil
		}

		if err := m.TapeMode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tape_mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tape_mode")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NdmpSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NdmpSession) UnmarshalBinary(b []byte) error {
	var res NdmpSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NdmpSessionInlineLinks ndmp session inline links
//
// swagger:model ndmp_session_inline__links
type NdmpSessionInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ndmp session inline links
func (m *NdmpSessionInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NdmpSessionInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this ndmp session inline links based on the context it is used
func (m *NdmpSessionInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NdmpSessionInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NdmpSessionInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NdmpSessionInlineLinks) UnmarshalBinary(b []byte) error {
	var res NdmpSessionInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NdmpSessionInlineNode ndmp session inline node
//
// swagger:model ndmp_session_inline_node
type NdmpSessionInlineNode struct {

	// links
	Links *NdmpSessionInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this ndmp session inline node
func (m *NdmpSessionInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NdmpSessionInlineNode) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this ndmp session inline node based on the context it is used
func (m *NdmpSessionInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NdmpSessionInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NdmpSessionInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NdmpSessionInlineNode) UnmarshalBinary(b []byte) error {
	var res NdmpSessionInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NdmpSessionInlineNodeInlineLinks ndmp session inline node inline links
//
// swagger:model ndmp_session_inline_node_inline__links
type NdmpSessionInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ndmp session inline node inline links
func (m *NdmpSessionInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NdmpSessionInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this ndmp session inline node inline links based on the context it is used
func (m *NdmpSessionInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NdmpSessionInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NdmpSessionInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NdmpSessionInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res NdmpSessionInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NdmpSessionInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model ndmp_session_inline_svm
type NdmpSessionInlineSvm struct {

	// links
	Links *NdmpSessionInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this ndmp session inline svm
func (m *NdmpSessionInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NdmpSessionInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this ndmp session inline svm based on the context it is used
func (m *NdmpSessionInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NdmpSessionInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NdmpSessionInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NdmpSessionInlineSvm) UnmarshalBinary(b []byte) error {
	var res NdmpSessionInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NdmpSessionInlineSvmInlineLinks ndmp session inline svm inline links
//
// swagger:model ndmp_session_inline_svm_inline__links
type NdmpSessionInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ndmp session inline svm inline links
func (m *NdmpSessionInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NdmpSessionInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this ndmp session inline svm inline links based on the context it is used
func (m *NdmpSessionInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NdmpSessionInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NdmpSessionInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NdmpSessionInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res NdmpSessionInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
