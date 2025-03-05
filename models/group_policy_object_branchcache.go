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

// GroupPolicyObjectBranchcache group policy object branchcache
//
// swagger:model group_policy_object_branchcache
type GroupPolicyObjectBranchcache struct {

	// Hash publication mode.
	// Example: disabled
	// Enum: ["per_share","disabled","all_shares"]
	HashPublicationMode *string `json:"hash_publication_mode,omitempty"`

	// Hash version.
	// Example: version1
	// Enum: ["version1","version2","all_versions"]
	SupportedHashVersion *string `json:"supported_hash_version,omitempty"`
}

// Validate validates this group policy object branchcache
func (m *GroupPolicyObjectBranchcache) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHashPublicationMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportedHashVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var groupPolicyObjectBranchcacheTypeHashPublicationModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["per_share","disabled","all_shares"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupPolicyObjectBranchcacheTypeHashPublicationModePropEnum = append(groupPolicyObjectBranchcacheTypeHashPublicationModePropEnum, v)
	}
}

const (

	// GroupPolicyObjectBranchcacheHashPublicationModePerShare captures enum value "per_share"
	GroupPolicyObjectBranchcacheHashPublicationModePerShare string = "per_share"

	// GroupPolicyObjectBranchcacheHashPublicationModeDisabled captures enum value "disabled"
	GroupPolicyObjectBranchcacheHashPublicationModeDisabled string = "disabled"

	// GroupPolicyObjectBranchcacheHashPublicationModeAllShares captures enum value "all_shares"
	GroupPolicyObjectBranchcacheHashPublicationModeAllShares string = "all_shares"
)

// prop value enum
func (m *GroupPolicyObjectBranchcache) validateHashPublicationModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, groupPolicyObjectBranchcacheTypeHashPublicationModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GroupPolicyObjectBranchcache) validateHashPublicationMode(formats strfmt.Registry) error {
	if swag.IsZero(m.HashPublicationMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateHashPublicationModeEnum("hash_publication_mode", "body", *m.HashPublicationMode); err != nil {
		return err
	}

	return nil
}

var groupPolicyObjectBranchcacheTypeSupportedHashVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["version1","version2","all_versions"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupPolicyObjectBranchcacheTypeSupportedHashVersionPropEnum = append(groupPolicyObjectBranchcacheTypeSupportedHashVersionPropEnum, v)
	}
}

const (

	// GroupPolicyObjectBranchcacheSupportedHashVersionVersion1 captures enum value "version1"
	GroupPolicyObjectBranchcacheSupportedHashVersionVersion1 string = "version1"

	// GroupPolicyObjectBranchcacheSupportedHashVersionVersion2 captures enum value "version2"
	GroupPolicyObjectBranchcacheSupportedHashVersionVersion2 string = "version2"

	// GroupPolicyObjectBranchcacheSupportedHashVersionAllVersions captures enum value "all_versions"
	GroupPolicyObjectBranchcacheSupportedHashVersionAllVersions string = "all_versions"
)

// prop value enum
func (m *GroupPolicyObjectBranchcache) validateSupportedHashVersionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, groupPolicyObjectBranchcacheTypeSupportedHashVersionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GroupPolicyObjectBranchcache) validateSupportedHashVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.SupportedHashVersion) { // not required
		return nil
	}

	// value enum
	if err := m.validateSupportedHashVersionEnum("supported_hash_version", "body", *m.SupportedHashVersion); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this group policy object branchcache based on context it is used
func (m *GroupPolicyObjectBranchcache) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GroupPolicyObjectBranchcache) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupPolicyObjectBranchcache) UnmarshalBinary(b []byte) error {
	var res GroupPolicyObjectBranchcache
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
