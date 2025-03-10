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

// IscsiSession An iSCSI session is one or more TCP connections that link an iSCSI initiator with an iSCSI target. TCP connections can be added and removed from an iSCSI session by the iSCSI initiator. Across all TCP connections within an iSCSI session, an initiator sees one and the same target. After the connection is established, iSCSI control, data, and status messages are communicated over the session.
//
// swagger:model iscsi_session
type IscsiSession struct {

	// links
	Links *IscsiSessionInlineLinks `json:"_links,omitempty"`

	// initiator
	Initiator *IscsiSessionInlineInitiator `json:"initiator,omitempty"`

	// The iSCSI connections that make up the iSCSI session.
	//
	// Read Only: true
	IscsiSessionInlineConnections []*IscsiConnection `json:"connections,omitempty"`

	// The initiator groups in which the initiator is a member.
	//
	// Read Only: true
	IscsiSessionInlineIgroups []*IscsiSessionInlineIgroupsInlineArrayItem `json:"igroups,omitempty"`

	// The initiator portion of the session identifier specified by the initiator during login.
	//
	// Example: 61:62:63:64:65:00
	// Read Only: true
	Isid *string `json:"isid,omitempty"`

	// svm
	Svm *IscsiSessionInlineSvm `json:"svm,omitempty"`

	// The target portal group to which the session belongs.
	//
	// Example: tpgroup1
	// Read Only: true
	TargetPortalGroup *string `json:"target_portal_group,omitempty"`

	// The target portal group tag of the session.
	//
	// Read Only: true
	TargetPortalGroupTag *int64 `json:"target_portal_group_tag,omitempty"`

	// The target session identifier handle (TSIH) of the session.
	//
	// Read Only: true
	Tsih *int64 `json:"tsih,omitempty"`
}

// Validate validates this iscsi session
func (m *IscsiSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIscsiSessionInlineConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIscsiSessionInlineIgroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiSession) validateLinks(formats strfmt.Registry) error {
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

func (m *IscsiSession) validateInitiator(formats strfmt.Registry) error {
	if swag.IsZero(m.Initiator) { // not required
		return nil
	}

	if m.Initiator != nil {
		if err := m.Initiator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initiator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("initiator")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiSession) validateIscsiSessionInlineConnections(formats strfmt.Registry) error {
	if swag.IsZero(m.IscsiSessionInlineConnections) { // not required
		return nil
	}

	for i := 0; i < len(m.IscsiSessionInlineConnections); i++ {
		if swag.IsZero(m.IscsiSessionInlineConnections[i]) { // not required
			continue
		}

		if m.IscsiSessionInlineConnections[i] != nil {
			if err := m.IscsiSessionInlineConnections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IscsiSession) validateIscsiSessionInlineIgroups(formats strfmt.Registry) error {
	if swag.IsZero(m.IscsiSessionInlineIgroups) { // not required
		return nil
	}

	for i := 0; i < len(m.IscsiSessionInlineIgroups); i++ {
		if swag.IsZero(m.IscsiSessionInlineIgroups[i]) { // not required
			continue
		}

		if m.IscsiSessionInlineIgroups[i] != nil {
			if err := m.IscsiSessionInlineIgroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("igroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IscsiSession) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this iscsi session based on the context it is used
func (m *IscsiSession) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInitiator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIscsiSessionInlineConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIscsiSessionInlineIgroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetPortalGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetPortalGroupTag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTsih(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiSession) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IscsiSession) contextValidateInitiator(ctx context.Context, formats strfmt.Registry) error {

	if m.Initiator != nil {

		if swag.IsZero(m.Initiator) { // not required
			return nil
		}

		if err := m.Initiator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initiator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("initiator")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiSession) contextValidateIscsiSessionInlineConnections(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connections", "body", []*IscsiConnection(m.IscsiSessionInlineConnections)); err != nil {
		return err
	}

	for i := 0; i < len(m.IscsiSessionInlineConnections); i++ {

		if m.IscsiSessionInlineConnections[i] != nil {

			if swag.IsZero(m.IscsiSessionInlineConnections[i]) { // not required
				return nil
			}

			if err := m.IscsiSessionInlineConnections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IscsiSession) contextValidateIscsiSessionInlineIgroups(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "igroups", "body", []*IscsiSessionInlineIgroupsInlineArrayItem(m.IscsiSessionInlineIgroups)); err != nil {
		return err
	}

	for i := 0; i < len(m.IscsiSessionInlineIgroups); i++ {

		if m.IscsiSessionInlineIgroups[i] != nil {

			if swag.IsZero(m.IscsiSessionInlineIgroups[i]) { // not required
				return nil
			}

			if err := m.IscsiSessionInlineIgroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("igroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IscsiSession) contextValidateIsid(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "isid", "body", m.Isid); err != nil {
		return err
	}

	return nil
}

func (m *IscsiSession) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IscsiSession) contextValidateTargetPortalGroup(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "target_portal_group", "body", m.TargetPortalGroup); err != nil {
		return err
	}

	return nil
}

func (m *IscsiSession) contextValidateTargetPortalGroupTag(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "target_portal_group_tag", "body", m.TargetPortalGroupTag); err != nil {
		return err
	}

	return nil
}

func (m *IscsiSession) contextValidateTsih(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "tsih", "body", m.Tsih); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IscsiSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiSession) UnmarshalBinary(b []byte) error {
	var res IscsiSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiSessionInlineIgroupsInlineArrayItem iscsi session inline igroups inline array item
//
// swagger:model iscsi_session_inline_igroups_inline_array_item
type IscsiSessionInlineIgroupsInlineArrayItem struct {

	// links
	Links *IscsiSessionInlineIgroupsInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// The name of the initiator group.
	//
	// Example: igroup1
	// Max Length: 96
	// Min Length: 1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the initiator group.
	//
	// Example: 4ea7a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this iscsi session inline igroups inline array item
func (m *IscsiSessionInlineIgroupsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiSessionInlineIgroupsInlineArrayItem) validateLinks(formats strfmt.Registry) error {
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

func (m *IscsiSessionInlineIgroupsInlineArrayItem) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 96); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this iscsi session inline igroups inline array item based on the context it is used
func (m *IscsiSessionInlineIgroupsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiSessionInlineIgroupsInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *IscsiSessionInlineIgroupsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiSessionInlineIgroupsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res IscsiSessionInlineIgroupsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiSessionInlineIgroupsInlineArrayItemInlineLinks iscsi session inline igroups inline array item inline links
//
// swagger:model iscsi_session_inline_igroups_inline_array_item_inline__links
type IscsiSessionInlineIgroupsInlineArrayItemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this iscsi session inline igroups inline array item inline links
func (m *IscsiSessionInlineIgroupsInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiSessionInlineIgroupsInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this iscsi session inline igroups inline array item inline links based on the context it is used
func (m *IscsiSessionInlineIgroupsInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiSessionInlineIgroupsInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IscsiSessionInlineIgroupsInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiSessionInlineIgroupsInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res IscsiSessionInlineIgroupsInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiSessionInlineInitiator The initiator that created the session.
//
// swagger:model iscsi_session_inline_initiator
type IscsiSessionInlineInitiator struct {

	// The initiator alias.
	//
	// Example: initiator_alias1
	// Read Only: true
	Alias *string `json:"alias,omitempty"`

	// A comment available for use by the administrator. This is modifiable from the initiator REST endpoint directly. See [`PATCH /protocols/san/igroups/{igroup.uuid}/initiators/{name}`](#/SAN/igroup_initiator_modify).
	//
	// Example: This is an iSCSI initiator for host 5
	// Read Only: true
	Comment *string `json:"comment,omitempty"`

	// The world wide unique name of the initiator.
	//
	// Example: iqn.1992-01.example.com:string
	// Read Only: true
	Name *string `json:"name,omitempty"`
}

// Validate validates this iscsi session inline initiator
func (m *IscsiSessionInlineInitiator) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this iscsi session inline initiator based on the context it is used
func (m *IscsiSessionInlineInitiator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlias(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiSessionInlineInitiator) contextValidateAlias(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "initiator"+"."+"alias", "body", m.Alias); err != nil {
		return err
	}

	return nil
}

func (m *IscsiSessionInlineInitiator) contextValidateComment(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "initiator"+"."+"comment", "body", m.Comment); err != nil {
		return err
	}

	return nil
}

func (m *IscsiSessionInlineInitiator) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "initiator"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IscsiSessionInlineInitiator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiSessionInlineInitiator) UnmarshalBinary(b []byte) error {
	var res IscsiSessionInlineInitiator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiSessionInlineLinks iscsi session inline links
//
// swagger:model iscsi_session_inline__links
type IscsiSessionInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this iscsi session inline links
func (m *IscsiSessionInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiSessionInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this iscsi session inline links based on the context it is used
func (m *IscsiSessionInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiSessionInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IscsiSessionInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiSessionInlineLinks) UnmarshalBinary(b []byte) error {
	var res IscsiSessionInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiSessionInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model iscsi_session_inline_svm
type IscsiSessionInlineSvm struct {

	// links
	Links *IscsiSessionInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this iscsi session inline svm
func (m *IscsiSessionInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiSessionInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this iscsi session inline svm based on the context it is used
func (m *IscsiSessionInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiSessionInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IscsiSessionInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiSessionInlineSvm) UnmarshalBinary(b []byte) error {
	var res IscsiSessionInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiSessionInlineSvmInlineLinks iscsi session inline svm inline links
//
// swagger:model iscsi_session_inline_svm_inline__links
type IscsiSessionInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this iscsi session inline svm inline links
func (m *IscsiSessionInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiSessionInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this iscsi session inline svm inline links based on the context it is used
func (m *IscsiSessionInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiSessionInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IscsiSessionInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiSessionInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res IscsiSessionInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
