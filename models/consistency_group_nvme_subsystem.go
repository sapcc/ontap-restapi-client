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

// ConsistencyGroupNvmeSubsystem An NVMe subsystem maintains configuration state and namespace access control for a set of NVMe-connected hosts.
//
// swagger:model consistency_group_nvme_subsystem
type ConsistencyGroupNvmeSubsystem struct {

	// A configurable comment for the NVMe subsystem. Optional in POST and PATCH.
	//
	// Max Length: 255
	// Min Length: 0
	Comment *string `json:"comment,omitempty"`

	// The NVMe hosts configured for access to the NVMe subsystem.
	// Optional in POST.
	//
	ConsistencyGroupNvmeSubsystemInlineHosts []*ConsistencyGroupNvmeHost `json:"hosts,omitempty"`

	// The name of the NVMe subsystem. Once created, an NVMe subsystem cannot be renamed. Required in POST.
	//
	// Example: subsystem1
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name,omitempty"`

	// The host operating system of the NVMe subsystem's hosts. Required in POST.
	//
	// Enum: ["aix","linux","vmware","windows"]
	OsType *string `json:"os_type,omitempty"`

	// The unique identifier of the NVMe subsystem.
	//
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this consistency group nvme subsystem
func (m *ConsistencyGroupNvmeSubsystem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsistencyGroupNvmeSubsystemInlineHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupNvmeSubsystem) validateComment(formats strfmt.Registry) error {
	if swag.IsZero(m.Comment) { // not required
		return nil
	}

	if err := validate.MinLength("comment", "body", *m.Comment, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("comment", "body", *m.Comment, 255); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupNvmeSubsystem) validateConsistencyGroupNvmeSubsystemInlineHosts(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsistencyGroupNvmeSubsystemInlineHosts) { // not required
		return nil
	}

	for i := 0; i < len(m.ConsistencyGroupNvmeSubsystemInlineHosts); i++ {
		if swag.IsZero(m.ConsistencyGroupNvmeSubsystemInlineHosts[i]) { // not required
			continue
		}

		if m.ConsistencyGroupNvmeSubsystemInlineHosts[i] != nil {
			if err := m.ConsistencyGroupNvmeSubsystemInlineHosts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hosts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsistencyGroupNvmeSubsystem) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

var consistencyGroupNvmeSubsystemTypeOsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aix","linux","vmware","windows"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consistencyGroupNvmeSubsystemTypeOsTypePropEnum = append(consistencyGroupNvmeSubsystemTypeOsTypePropEnum, v)
	}
}

const (

	// ConsistencyGroupNvmeSubsystemOsTypeAix captures enum value "aix"
	ConsistencyGroupNvmeSubsystemOsTypeAix string = "aix"

	// ConsistencyGroupNvmeSubsystemOsTypeLinux captures enum value "linux"
	ConsistencyGroupNvmeSubsystemOsTypeLinux string = "linux"

	// ConsistencyGroupNvmeSubsystemOsTypeVmware captures enum value "vmware"
	ConsistencyGroupNvmeSubsystemOsTypeVmware string = "vmware"

	// ConsistencyGroupNvmeSubsystemOsTypeWindows captures enum value "windows"
	ConsistencyGroupNvmeSubsystemOsTypeWindows string = "windows"
)

// prop value enum
func (m *ConsistencyGroupNvmeSubsystem) validateOsTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consistencyGroupNvmeSubsystemTypeOsTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsistencyGroupNvmeSubsystem) validateOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.OsType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOsTypeEnum("os_type", "body", *m.OsType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this consistency group nvme subsystem based on the context it is used
func (m *ConsistencyGroupNvmeSubsystem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConsistencyGroupNvmeSubsystemInlineHosts(ctx, formats); err != nil {
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

func (m *ConsistencyGroupNvmeSubsystem) contextValidateConsistencyGroupNvmeSubsystemInlineHosts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConsistencyGroupNvmeSubsystemInlineHosts); i++ {

		if m.ConsistencyGroupNvmeSubsystemInlineHosts[i] != nil {

			if swag.IsZero(m.ConsistencyGroupNvmeSubsystemInlineHosts[i]) { // not required
				return nil
			}

			if err := m.ConsistencyGroupNvmeSubsystemInlineHosts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hosts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsistencyGroupNvmeSubsystem) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupNvmeSubsystem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupNvmeSubsystem) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupNvmeSubsystem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
