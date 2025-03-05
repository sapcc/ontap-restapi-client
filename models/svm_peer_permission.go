// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SvmPeerPermission Manage SVM peer permissions.
//
// swagger:model svm_peer_permission
type SvmPeerPermission struct {

	// links
	Links *SvmPeerPermissionInlineLinks `json:"_links,omitempty"`

	// cluster peer
	ClusterPeer *SvmPeerPermissionInlineClusterPeer `json:"cluster_peer,omitempty"`

	// svm
	Svm *SvmPeerPermissionInlineSvm `json:"svm,omitempty"`

	// A list of applications for an SVM peer relationship.
	// Example: ["snapmirror","flexcache"]
	SvmPeerPermissionInlineApplications []*SvmPeerPermissionApplications `json:"applications,omitempty"`
}

// Validate validates this svm peer permission
func (m *SvmPeerPermission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterPeer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvmPeerPermissionInlineApplications(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermission) validateLinks(formats strfmt.Registry) error {
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

func (m *SvmPeerPermission) validateClusterPeer(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterPeer) { // not required
		return nil
	}

	if m.ClusterPeer != nil {
		if err := m.ClusterPeer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_peer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_peer")
			}
			return err
		}
	}

	return nil
}

func (m *SvmPeerPermission) validateSvm(formats strfmt.Registry) error {
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

func (m *SvmPeerPermission) validateSvmPeerPermissionInlineApplications(formats strfmt.Registry) error {
	if swag.IsZero(m.SvmPeerPermissionInlineApplications) { // not required
		return nil
	}

	for i := 0; i < len(m.SvmPeerPermissionInlineApplications); i++ {
		if swag.IsZero(m.SvmPeerPermissionInlineApplications[i]) { // not required
			continue
		}

		if m.SvmPeerPermissionInlineApplications[i] != nil {
			if err := m.SvmPeerPermissionInlineApplications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("applications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this svm peer permission based on the context it is used
func (m *SvmPeerPermission) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterPeer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvmPeerPermissionInlineApplications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermission) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SvmPeerPermission) contextValidateClusterPeer(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterPeer != nil {

		if swag.IsZero(m.ClusterPeer) { // not required
			return nil
		}

		if err := m.ClusterPeer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_peer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_peer")
			}
			return err
		}
	}

	return nil
}

func (m *SvmPeerPermission) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SvmPeerPermission) contextValidateSvmPeerPermissionInlineApplications(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SvmPeerPermissionInlineApplications); i++ {

		if m.SvmPeerPermissionInlineApplications[i] != nil {

			if swag.IsZero(m.SvmPeerPermissionInlineApplications[i]) { // not required
				return nil
			}

			if err := m.SvmPeerPermissionInlineApplications[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("applications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmPeerPermission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerPermission) UnmarshalBinary(b []byte) error {
	var res SvmPeerPermission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerPermissionInlineClusterPeer Peer cluster details
//
// swagger:model svm_peer_permission_inline_cluster_peer
type SvmPeerPermissionInlineClusterPeer struct {

	// links
	Links *SvmPeerPermissionInlineClusterPeerInlineLinks `json:"_links,omitempty"`

	// name
	// Example: cluster2
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: ebe27c49-1adf-4496-8335-ab862aebebf2
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this svm peer permission inline cluster peer
func (m *SvmPeerPermissionInlineClusterPeer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionInlineClusterPeer) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_peer" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_peer" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm peer permission inline cluster peer based on the context it is used
func (m *SvmPeerPermissionInlineClusterPeer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionInlineClusterPeer) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_peer" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_peer" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmPeerPermissionInlineClusterPeer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerPermissionInlineClusterPeer) UnmarshalBinary(b []byte) error {
	var res SvmPeerPermissionInlineClusterPeer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerPermissionInlineClusterPeerInlineLinks svm peer permission inline cluster peer inline links
//
// swagger:model svm_peer_permission_inline_cluster_peer_inline__links
type SvmPeerPermissionInlineClusterPeerInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this svm peer permission inline cluster peer inline links
func (m *SvmPeerPermissionInlineClusterPeerInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionInlineClusterPeerInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_peer" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_peer" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm peer permission inline cluster peer inline links based on the context it is used
func (m *SvmPeerPermissionInlineClusterPeerInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionInlineClusterPeerInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_peer" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_peer" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmPeerPermissionInlineClusterPeerInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerPermissionInlineClusterPeerInlineLinks) UnmarshalBinary(b []byte) error {
	var res SvmPeerPermissionInlineClusterPeerInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerPermissionInlineLinks svm peer permission inline links
//
// swagger:model svm_peer_permission_inline__links
type SvmPeerPermissionInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this svm peer permission inline links
func (m *SvmPeerPermissionInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this svm peer permission inline links based on the context it is used
func (m *SvmPeerPermissionInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SvmPeerPermissionInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerPermissionInlineLinks) UnmarshalBinary(b []byte) error {
	var res SvmPeerPermissionInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerPermissionInlineSvm Local SVM permitted for peer relationship. To create peer permissions for all SVMs, specify the SVM name as "*".
//
// swagger:model svm_peer_permission_inline_svm
type SvmPeerPermissionInlineSvm struct {

	// links
	Links *SvmPeerPermissionInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this svm peer permission inline svm
func (m *SvmPeerPermissionInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this svm peer permission inline svm based on the context it is used
func (m *SvmPeerPermissionInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SvmPeerPermissionInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerPermissionInlineSvm) UnmarshalBinary(b []byte) error {
	var res SvmPeerPermissionInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmPeerPermissionInlineSvmInlineLinks svm peer permission inline svm inline links
//
// swagger:model svm_peer_permission_inline_svm_inline__links
type SvmPeerPermissionInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this svm peer permission inline svm inline links
func (m *SvmPeerPermissionInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this svm peer permission inline svm inline links based on the context it is used
func (m *SvmPeerPermissionInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmPeerPermissionInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SvmPeerPermissionInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmPeerPermissionInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res SvmPeerPermissionInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
