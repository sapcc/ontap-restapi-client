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

// VsiOnNas A VSI application using NAS.
//
// swagger:model vsi_on_nas
type VsiOnNas struct {

	// datastore
	// Required: true
	Datastore *VsiOnNasInlineDatastore `json:"datastore"`

	// hyper v access
	HypervAccess *VsiOnNasInlineHypervAccess `json:"hyper_v_access,omitempty"`

	// protection type
	ProtectionType *VsiOnNasInlineProtectionType `json:"protection_type,omitempty"`

	// The list of NFS access controls. You must provide either 'host' or 'access' to enable NFS access.
	VsiOnNasInlineNfsAccess []*AppNfsAccess `json:"nfs_access,omitempty"`
}

// Validate validates this vsi on nas
func (m *VsiOnNas) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatastore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHypervAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsiOnNasInlineNfsAccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VsiOnNas) validateDatastore(formats strfmt.Registry) error {

	if err := validate.Required("datastore", "body", m.Datastore); err != nil {
		return err
	}

	if m.Datastore != nil {
		if err := m.Datastore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datastore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datastore")
			}
			return err
		}
	}

	return nil
}

func (m *VsiOnNas) validateHypervAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.HypervAccess) { // not required
		return nil
	}

	if m.HypervAccess != nil {
		if err := m.HypervAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hyper_v_access")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hyper_v_access")
			}
			return err
		}
	}

	return nil
}

func (m *VsiOnNas) validateProtectionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtectionType) { // not required
		return nil
	}

	if m.ProtectionType != nil {
		if err := m.ProtectionType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protection_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protection_type")
			}
			return err
		}
	}

	return nil
}

func (m *VsiOnNas) validateVsiOnNasInlineNfsAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.VsiOnNasInlineNfsAccess) { // not required
		return nil
	}

	for i := 0; i < len(m.VsiOnNasInlineNfsAccess); i++ {
		if swag.IsZero(m.VsiOnNasInlineNfsAccess[i]) { // not required
			continue
		}

		if m.VsiOnNasInlineNfsAccess[i] != nil {
			if err := m.VsiOnNasInlineNfsAccess[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nfs_access" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nfs_access" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vsi on nas based on the context it is used
func (m *VsiOnNas) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatastore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHypervAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtectionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVsiOnNasInlineNfsAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VsiOnNas) contextValidateDatastore(ctx context.Context, formats strfmt.Registry) error {

	if m.Datastore != nil {

		if err := m.Datastore.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datastore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datastore")
			}
			return err
		}
	}

	return nil
}

func (m *VsiOnNas) contextValidateHypervAccess(ctx context.Context, formats strfmt.Registry) error {

	if m.HypervAccess != nil {

		if swag.IsZero(m.HypervAccess) { // not required
			return nil
		}

		if err := m.HypervAccess.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hyper_v_access")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hyper_v_access")
			}
			return err
		}
	}

	return nil
}

func (m *VsiOnNas) contextValidateProtectionType(ctx context.Context, formats strfmt.Registry) error {

	if m.ProtectionType != nil {

		if swag.IsZero(m.ProtectionType) { // not required
			return nil
		}

		if err := m.ProtectionType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protection_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protection_type")
			}
			return err
		}
	}

	return nil
}

func (m *VsiOnNas) contextValidateVsiOnNasInlineNfsAccess(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VsiOnNasInlineNfsAccess); i++ {

		if m.VsiOnNasInlineNfsAccess[i] != nil {

			if swag.IsZero(m.VsiOnNasInlineNfsAccess[i]) { // not required
				return nil
			}

			if err := m.VsiOnNasInlineNfsAccess[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nfs_access" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nfs_access" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VsiOnNas) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VsiOnNas) UnmarshalBinary(b []byte) error {
	var res VsiOnNas
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VsiOnNasInlineDatastore vsi on nas inline datastore
//
// swagger:model vsi_on_nas_inline_datastore
type VsiOnNasInlineDatastore struct {

	// The number of datastores to support.
	// Maximum: 16
	// Minimum: 1
	Count *int64 `json:"count,omitempty"`

	// The size of the datastore. Usage: {&lt;integer&gt;[KB|MB|GB|TB|PB]}
	// Required: true
	Size *int64 `json:"size"`

	// storage service
	StorageService *VsiOnNasInlineDatastoreInlineStorageService `json:"storage_service,omitempty"`
}

// Validate validates this vsi on nas inline datastore
func (m *VsiOnNasInlineDatastore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VsiOnNasInlineDatastore) validateCount(formats strfmt.Registry) error {
	if swag.IsZero(m.Count) { // not required
		return nil
	}

	if err := validate.MinimumInt("datastore"+"."+"count", "body", *m.Count, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("datastore"+"."+"count", "body", *m.Count, 16, false); err != nil {
		return err
	}

	return nil
}

func (m *VsiOnNasInlineDatastore) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("datastore"+"."+"size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *VsiOnNasInlineDatastore) validateStorageService(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageService) { // not required
		return nil
	}

	if m.StorageService != nil {
		if err := m.StorageService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datastore" + "." + "storage_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datastore" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vsi on nas inline datastore based on the context it is used
func (m *VsiOnNasInlineDatastore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorageService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VsiOnNasInlineDatastore) contextValidateStorageService(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageService != nil {

		if swag.IsZero(m.StorageService) { // not required
			return nil
		}

		if err := m.StorageService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datastore" + "." + "storage_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datastore" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VsiOnNasInlineDatastore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VsiOnNasInlineDatastore) UnmarshalBinary(b []byte) error {
	var res VsiOnNasInlineDatastore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VsiOnNasInlineDatastoreInlineStorageService vsi on nas inline datastore inline storage service
//
// swagger:model vsi_on_nas_inline_datastore_inline_storage_service
type VsiOnNasInlineDatastoreInlineStorageService struct {

	// The storage service of the datastore.
	// Enum: ["extreme","performance","value"]
	Name *string `json:"name,omitempty"`
}

// Validate validates this vsi on nas inline datastore inline storage service
func (m *VsiOnNasInlineDatastoreInlineStorageService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vsiOnNasInlineDatastoreInlineStorageServiceTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["extreme","performance","value"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vsiOnNasInlineDatastoreInlineStorageServiceTypeNamePropEnum = append(vsiOnNasInlineDatastoreInlineStorageServiceTypeNamePropEnum, v)
	}
}

const (

	// VsiOnNasInlineDatastoreInlineStorageServiceNameExtreme captures enum value "extreme"
	VsiOnNasInlineDatastoreInlineStorageServiceNameExtreme string = "extreme"

	// VsiOnNasInlineDatastoreInlineStorageServiceNamePerformance captures enum value "performance"
	VsiOnNasInlineDatastoreInlineStorageServiceNamePerformance string = "performance"

	// VsiOnNasInlineDatastoreInlineStorageServiceNameValue captures enum value "value"
	VsiOnNasInlineDatastoreInlineStorageServiceNameValue string = "value"
)

// prop value enum
func (m *VsiOnNasInlineDatastoreInlineStorageService) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vsiOnNasInlineDatastoreInlineStorageServiceTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VsiOnNasInlineDatastoreInlineStorageService) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("datastore"+"."+"storage_service"+"."+"name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vsi on nas inline datastore inline storage service based on context it is used
func (m *VsiOnNasInlineDatastoreInlineStorageService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VsiOnNasInlineDatastoreInlineStorageService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VsiOnNasInlineDatastoreInlineStorageService) UnmarshalBinary(b []byte) error {
	var res VsiOnNasInlineDatastoreInlineStorageService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VsiOnNasInlineHypervAccess vsi on nas inline hyper v access
//
// swagger:model vsi_on_nas_inline_hyper_v_access
type VsiOnNasInlineHypervAccess struct {

	// Hyper-V service account.
	// Max Length: 256
	// Min Length: 1
	ServiceAccount *string `json:"service_account,omitempty"`
}

// Validate validates this vsi on nas inline hyper v access
func (m *VsiOnNasInlineHypervAccess) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VsiOnNasInlineHypervAccess) validateServiceAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceAccount) { // not required
		return nil
	}

	if err := validate.MinLength("hyper_v_access"+"."+"service_account", "body", *m.ServiceAccount, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("hyper_v_access"+"."+"service_account", "body", *m.ServiceAccount, 256); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vsi on nas inline hyper v access based on context it is used
func (m *VsiOnNasInlineHypervAccess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VsiOnNasInlineHypervAccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VsiOnNasInlineHypervAccess) UnmarshalBinary(b []byte) error {
	var res VsiOnNasInlineHypervAccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VsiOnNasInlineProtectionType vsi on nas inline protection type
//
// swagger:model vsi_on_nas_inline_protection_type
type VsiOnNasInlineProtectionType struct {

	// The local RPO of the application.
	// Enum: ["hourly","none"]
	LocalRpo *string `json:"local_rpo,omitempty"`

	// The remote RPO of the application.
	// Enum: ["none","zero"]
	RemoteRpo *string `json:"remote_rpo,omitempty"`
}

// Validate validates this vsi on nas inline protection type
func (m *VsiOnNasInlineProtectionType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocalRpo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteRpo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vsiOnNasInlineProtectionTypeTypeLocalRpoPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hourly","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vsiOnNasInlineProtectionTypeTypeLocalRpoPropEnum = append(vsiOnNasInlineProtectionTypeTypeLocalRpoPropEnum, v)
	}
}

const (

	// VsiOnNasInlineProtectionTypeLocalRpoHourly captures enum value "hourly"
	VsiOnNasInlineProtectionTypeLocalRpoHourly string = "hourly"

	// VsiOnNasInlineProtectionTypeLocalRpoNone captures enum value "none"
	VsiOnNasInlineProtectionTypeLocalRpoNone string = "none"
)

// prop value enum
func (m *VsiOnNasInlineProtectionType) validateLocalRpoEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vsiOnNasInlineProtectionTypeTypeLocalRpoPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VsiOnNasInlineProtectionType) validateLocalRpo(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalRpo) { // not required
		return nil
	}

	// value enum
	if err := m.validateLocalRpoEnum("protection_type"+"."+"local_rpo", "body", *m.LocalRpo); err != nil {
		return err
	}

	return nil
}

var vsiOnNasInlineProtectionTypeTypeRemoteRpoPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","zero"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vsiOnNasInlineProtectionTypeTypeRemoteRpoPropEnum = append(vsiOnNasInlineProtectionTypeTypeRemoteRpoPropEnum, v)
	}
}

const (

	// VsiOnNasInlineProtectionTypeRemoteRpoNone captures enum value "none"
	VsiOnNasInlineProtectionTypeRemoteRpoNone string = "none"

	// VsiOnNasInlineProtectionTypeRemoteRpoZero captures enum value "zero"
	VsiOnNasInlineProtectionTypeRemoteRpoZero string = "zero"
)

// prop value enum
func (m *VsiOnNasInlineProtectionType) validateRemoteRpoEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vsiOnNasInlineProtectionTypeTypeRemoteRpoPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VsiOnNasInlineProtectionType) validateRemoteRpo(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteRpo) { // not required
		return nil
	}

	// value enum
	if err := m.validateRemoteRpoEnum("protection_type"+"."+"remote_rpo", "body", *m.RemoteRpo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vsi on nas inline protection type based on context it is used
func (m *VsiOnNasInlineProtectionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VsiOnNasInlineProtectionType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VsiOnNasInlineProtectionType) UnmarshalBinary(b []byte) error {
	var res VsiOnNasInlineProtectionType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
