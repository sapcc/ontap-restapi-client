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
	"github.com/go-openapi/validate"
)

// MediatorResponse mediator response
//
// swagger:model mediator_response
type MediatorResponse struct {

	// links
	Links *MediatorResponseInlineLinks `json:"_links,omitempty"`

	// mediator response inline records
	MediatorResponseInlineRecords []*MediatorResponseInlineRecordsInlineArrayItem `json:"records,omitempty"`

	// Number of records
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty"`
}

// Validate validates this mediator response
func (m *MediatorResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediatorResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediatorResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *MediatorResponse) validateMediatorResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.MediatorResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.MediatorResponseInlineRecords); i++ {
		if swag.IsZero(m.MediatorResponseInlineRecords[i]) { // not required
			continue
		}

		if m.MediatorResponseInlineRecords[i] != nil {
			if err := m.MediatorResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this mediator response based on the context it is used
func (m *MediatorResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMediatorResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediatorResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MediatorResponse) contextValidateMediatorResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MediatorResponseInlineRecords); i++ {

		if m.MediatorResponseInlineRecords[i] != nil {

			if swag.IsZero(m.MediatorResponseInlineRecords[i]) { // not required
				return nil
			}

			if err := m.MediatorResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MediatorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediatorResponse) UnmarshalBinary(b []byte) error {
	var res MediatorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MediatorResponseInlineLinks mediator response inline links
//
// swagger:model mediator_response_inline__links
type MediatorResponseInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this mediator response inline links
func (m *MediatorResponseInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediatorResponseInlineLinks) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *MediatorResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this mediator response inline links based on the context it is used
func (m *MediatorResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediatorResponseInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {

		if swag.IsZero(m.Next) { // not required
			return nil
		}

		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *MediatorResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MediatorResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediatorResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res MediatorResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MediatorResponseInlineRecordsInlineArrayItem Mediator information
//
// swagger:model mediator_response_inline_records_inline_array_item
type MediatorResponseInlineRecordsInlineArrayItem struct {

	// CA certificate for ONTAP Mediator. This is optional if the certificate is already installed.
	CaCertificate *string `json:"ca_certificate,omitempty"`

	// dr group
	DrGroup *MediatorResponseInlineRecordsInlineArrayItemInlineDrGroup `json:"dr_group,omitempty"`

	// The IP address of the mediator.
	// Example: 10.10.10.7
	IPAddress *string `json:"ip_address,omitempty"`

	// The password used to connect to the REST server on the mediator.
	// Example: mypassword
	// Format: password
	Password *strfmt.Password `json:"password,omitempty"`

	// peer cluster
	PeerCluster *MediatorResponseInlineRecordsInlineArrayItemInlinePeerCluster `json:"peer_cluster,omitempty"`

	// Indicates the mediator connectivity status of the peer cluster. Possible values are connected, unreachable, unknown.
	// Example: connected
	// Read Only: true
	PeerMediatorConnectivity *string `json:"peer_mediator_connectivity,omitempty"`

	// The REST server's port number on the mediator.
	// Example: 31784
	Port *int64 `json:"port,omitempty"`

	// Indicates the connectivity status of the mediator.
	// Example: true
	// Read Only: true
	Reachable *bool `json:"reachable,omitempty"`

	// The username used to connect to the REST server on the mediator.
	// Example: myusername
	User *string `json:"user,omitempty"`

	// The unique identifier for the mediator service.
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this mediator response inline records inline array item
func (m *MediatorResponseInlineRecordsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDrGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeerCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediatorResponseInlineRecordsInlineArrayItem) validateDrGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.DrGroup) { // not required
		return nil
	}

	if m.DrGroup != nil {
		if err := m.DrGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dr_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dr_group")
			}
			return err
		}
	}

	return nil
}

func (m *MediatorResponseInlineRecordsInlineArrayItem) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.FormatOf("password", "body", "password", m.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MediatorResponseInlineRecordsInlineArrayItem) validatePeerCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.PeerCluster) { // not required
		return nil
	}

	if m.PeerCluster != nil {
		if err := m.PeerCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer_cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peer_cluster")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mediator response inline records inline array item based on the context it is used
func (m *MediatorResponseInlineRecordsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDrGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePeerCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePeerMediatorConnectivity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReachable(ctx, formats); err != nil {
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

func (m *MediatorResponseInlineRecordsInlineArrayItem) contextValidateDrGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.DrGroup != nil {

		if swag.IsZero(m.DrGroup) { // not required
			return nil
		}

		if err := m.DrGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dr_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dr_group")
			}
			return err
		}
	}

	return nil
}

func (m *MediatorResponseInlineRecordsInlineArrayItem) contextValidatePeerCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.PeerCluster != nil {

		if swag.IsZero(m.PeerCluster) { // not required
			return nil
		}

		if err := m.PeerCluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer_cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peer_cluster")
			}
			return err
		}
	}

	return nil
}

func (m *MediatorResponseInlineRecordsInlineArrayItem) contextValidatePeerMediatorConnectivity(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "peer_mediator_connectivity", "body", m.PeerMediatorConnectivity); err != nil {
		return err
	}

	return nil
}

func (m *MediatorResponseInlineRecordsInlineArrayItem) contextValidateReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "reachable", "body", m.Reachable); err != nil {
		return err
	}

	return nil
}

func (m *MediatorResponseInlineRecordsInlineArrayItem) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MediatorResponseInlineRecordsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediatorResponseInlineRecordsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res MediatorResponseInlineRecordsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MediatorResponseInlineRecordsInlineArrayItemInlineDrGroup DR group reference.
//
// swagger:model mediator_response_inline_records_inline_array_item_inline_dr_group
type MediatorResponseInlineRecordsInlineArrayItemInlineDrGroup struct {

	// DR Group ID
	// Read Only: true
	ID *int64 `json:"id,omitempty"`
}

// Validate validates this mediator response inline records inline array item inline dr group
func (m *MediatorResponseInlineRecordsInlineArrayItemInlineDrGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this mediator response inline records inline array item inline dr group based on the context it is used
func (m *MediatorResponseInlineRecordsInlineArrayItemInlineDrGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediatorResponseInlineRecordsInlineArrayItemInlineDrGroup) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dr_group"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MediatorResponseInlineRecordsInlineArrayItemInlineDrGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediatorResponseInlineRecordsInlineArrayItemInlineDrGroup) UnmarshalBinary(b []byte) error {
	var res MediatorResponseInlineRecordsInlineArrayItemInlineDrGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MediatorResponseInlineRecordsInlineArrayItemInlinePeerCluster The peer cluster that the mediator service is used for.
//
// swagger:model mediator_response_inline_records_inline_array_item_inline_peer_cluster
type MediatorResponseInlineRecordsInlineArrayItemInlinePeerCluster struct {

	// links
	Links *MediatorResponseInlineRecordsInlineArrayItemInlinePeerClusterInlineLinks `json:"_links,omitempty"`

	// name
	// Example: cluster2
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: ebe27c49-1adf-4496-8335-ab862aebebf2
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this mediator response inline records inline array item inline peer cluster
func (m *MediatorResponseInlineRecordsInlineArrayItemInlinePeerCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediatorResponseInlineRecordsInlineArrayItemInlinePeerCluster) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer_cluster" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peer_cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mediator response inline records inline array item inline peer cluster based on the context it is used
func (m *MediatorResponseInlineRecordsInlineArrayItemInlinePeerCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediatorResponseInlineRecordsInlineArrayItemInlinePeerCluster) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer_cluster" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peer_cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MediatorResponseInlineRecordsInlineArrayItemInlinePeerCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediatorResponseInlineRecordsInlineArrayItemInlinePeerCluster) UnmarshalBinary(b []byte) error {
	var res MediatorResponseInlineRecordsInlineArrayItemInlinePeerCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MediatorResponseInlineRecordsInlineArrayItemInlinePeerClusterInlineLinks mediator response inline records inline array item inline peer cluster inline links
//
// swagger:model mediator_response_inline_records_inline_array_item_inline_peer_cluster_inline__links
type MediatorResponseInlineRecordsInlineArrayItemInlinePeerClusterInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this mediator response inline records inline array item inline peer cluster inline links
func (m *MediatorResponseInlineRecordsInlineArrayItemInlinePeerClusterInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediatorResponseInlineRecordsInlineArrayItemInlinePeerClusterInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer_cluster" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peer_cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mediator response inline records inline array item inline peer cluster inline links based on the context it is used
func (m *MediatorResponseInlineRecordsInlineArrayItemInlinePeerClusterInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediatorResponseInlineRecordsInlineArrayItemInlinePeerClusterInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer_cluster" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peer_cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MediatorResponseInlineRecordsInlineArrayItemInlinePeerClusterInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediatorResponseInlineRecordsInlineArrayItemInlinePeerClusterInlineLinks) UnmarshalBinary(b []byte) error {
	var res MediatorResponseInlineRecordsInlineArrayItemInlinePeerClusterInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
