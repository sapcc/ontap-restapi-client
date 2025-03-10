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

// ConfigurationBackupFile The configuration backup file.
//
// swagger:model configuration_backup_file
type ConfigurationBackupFile struct {

	// Indicates if the backup was created automatically.
	// Read Only: true
	Auto *bool `json:"auto,omitempty"`

	// The list of nodes included in the backup.
	// Read Only: true
	ConfigurationBackupFileInlineBackupNodes []*BackupNode `json:"backup_nodes,omitempty"`

	// The link to download the backup file.
	// Example: https://10.224.65.198/backups/backup_file.7z
	// Read Only: true
	DownloadLink *string `json:"download_link,omitempty"`

	// The backup name.
	// Example: backup_file.7z
	Name *string `json:"name,omitempty"`

	// node
	Node *ConfigurationBackupFileInlineNode `json:"node,omitempty"`

	// The size of the backup in bytes.
	// Example: 4787563
	// Read Only: true
	Size *int64 `json:"size,omitempty"`

	// The backup creation time.
	// Example: 2019-02-04 18:33:48
	// Read Only: true
	// Format: date-time
	Time *strfmt.DateTime `json:"time,omitempty"`

	// The backup type.
	// Read Only: true
	// Enum: ["node","cluster"]
	Type *string `json:"type,omitempty"`

	// The software version.
	// Example: 9.7.0
	// Read Only: true
	Version *string `json:"version,omitempty"`
}

// Validate validates this configuration backup file
func (m *ConfigurationBackupFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigurationBackupFileInlineBackupNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
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

func (m *ConfigurationBackupFile) validateConfigurationBackupFileInlineBackupNodes(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigurationBackupFileInlineBackupNodes) { // not required
		return nil
	}

	for i := 0; i < len(m.ConfigurationBackupFileInlineBackupNodes); i++ {
		if swag.IsZero(m.ConfigurationBackupFileInlineBackupNodes[i]) { // not required
			continue
		}

		if m.ConfigurationBackupFileInlineBackupNodes[i] != nil {
			if err := m.ConfigurationBackupFileInlineBackupNodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backup_nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backup_nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigurationBackupFile) validateNode(formats strfmt.Registry) error {
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

func (m *ConfigurationBackupFile) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

var configurationBackupFileTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["node","cluster"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configurationBackupFileTypeTypePropEnum = append(configurationBackupFileTypeTypePropEnum, v)
	}
}

const (

	// ConfigurationBackupFileTypeNode captures enum value "node"
	ConfigurationBackupFileTypeNode string = "node"

	// ConfigurationBackupFileTypeCluster captures enum value "cluster"
	ConfigurationBackupFileTypeCluster string = "cluster"
)

// prop value enum
func (m *ConfigurationBackupFile) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, configurationBackupFileTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConfigurationBackupFile) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this configuration backup file based on the context it is used
func (m *ConfigurationBackupFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuto(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfigurationBackupFileInlineBackupNodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDownloadLink(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigurationBackupFile) contextValidateAuto(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "auto", "body", m.Auto); err != nil {
		return err
	}

	return nil
}

func (m *ConfigurationBackupFile) contextValidateConfigurationBackupFileInlineBackupNodes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "backup_nodes", "body", []*BackupNode(m.ConfigurationBackupFileInlineBackupNodes)); err != nil {
		return err
	}

	for i := 0; i < len(m.ConfigurationBackupFileInlineBackupNodes); i++ {

		if m.ConfigurationBackupFileInlineBackupNodes[i] != nil {

			if swag.IsZero(m.ConfigurationBackupFileInlineBackupNodes[i]) { // not required
				return nil
			}

			if err := m.ConfigurationBackupFileInlineBackupNodes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backup_nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backup_nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigurationBackupFile) contextValidateDownloadLink(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "download_link", "body", m.DownloadLink); err != nil {
		return err
	}

	return nil
}

func (m *ConfigurationBackupFile) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConfigurationBackupFile) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *ConfigurationBackupFile) contextValidateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "time", "body", m.Time); err != nil {
		return err
	}

	return nil
}

func (m *ConfigurationBackupFile) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ConfigurationBackupFile) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigurationBackupFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigurationBackupFile) UnmarshalBinary(b []byte) error {
	var res ConfigurationBackupFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConfigurationBackupFileInlineNode The node that owns the configuration backup.
//
// swagger:model configuration_backup_file_inline_node
type ConfigurationBackupFileInlineNode struct {

	// links
	Links *ConfigurationBackupFileInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this configuration backup file inline node
func (m *ConfigurationBackupFileInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigurationBackupFileInlineNode) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this configuration backup file inline node based on the context it is used
func (m *ConfigurationBackupFileInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigurationBackupFileInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ConfigurationBackupFileInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigurationBackupFileInlineNode) UnmarshalBinary(b []byte) error {
	var res ConfigurationBackupFileInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConfigurationBackupFileInlineNodeInlineLinks configuration backup file inline node inline links
//
// swagger:model configuration_backup_file_inline_node_inline__links
type ConfigurationBackupFileInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this configuration backup file inline node inline links
func (m *ConfigurationBackupFileInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigurationBackupFileInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this configuration backup file inline node inline links based on the context it is used
func (m *ConfigurationBackupFileInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigurationBackupFileInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ConfigurationBackupFileInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigurationBackupFileInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res ConfigurationBackupFileInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
