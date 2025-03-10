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

// SvmMigrationIPInterfacePlacement Optional property used to specify the list of source SVM's IP interface and network port pairs in the destination for migrating the source SVM IP interfaces. Note that the SVM migration does not perform any reachability checks on the IP interfaces provided.
//
// swagger:model svm_migration_ip_interface_placement
type SvmMigrationIPInterfacePlacement struct {

	// List of source SVM's IP interface and port pairs on the destination for migrating the source SVM's IP interfaces.
	SvmMigrationIPInterfacePlacementInlineIPInterfaces []*SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItem `json:"ip_interfaces,omitempty"`
}

// Validate validates this svm migration ip interface placement
func (m *SvmMigrationIPInterfacePlacement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSvmMigrationIPInterfacePlacementInlineIPInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationIPInterfacePlacement) validateSvmMigrationIPInterfacePlacementInlineIPInterfaces(formats strfmt.Registry) error {
	if swag.IsZero(m.SvmMigrationIPInterfacePlacementInlineIPInterfaces) { // not required
		return nil
	}

	for i := 0; i < len(m.SvmMigrationIPInterfacePlacementInlineIPInterfaces); i++ {
		if swag.IsZero(m.SvmMigrationIPInterfacePlacementInlineIPInterfaces[i]) { // not required
			continue
		}

		if m.SvmMigrationIPInterfacePlacementInlineIPInterfaces[i] != nil {
			if err := m.SvmMigrationIPInterfacePlacementInlineIPInterfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ip_interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ip_interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this svm migration ip interface placement based on the context it is used
func (m *SvmMigrationIPInterfacePlacement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSvmMigrationIPInterfacePlacementInlineIPInterfaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationIPInterfacePlacement) contextValidateSvmMigrationIPInterfacePlacementInlineIPInterfaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SvmMigrationIPInterfacePlacementInlineIPInterfaces); i++ {

		if m.SvmMigrationIPInterfacePlacementInlineIPInterfaces[i] != nil {

			if swag.IsZero(m.SvmMigrationIPInterfacePlacementInlineIPInterfaces[i]) { // not required
				return nil
			}

			if err := m.SvmMigrationIPInterfacePlacementInlineIPInterfaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ip_interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ip_interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationIPInterfacePlacement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationIPInterfacePlacement) UnmarshalBinary(b []byte) error {
	var res SvmMigrationIPInterfacePlacement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItem IP interface and network port pair information.
//
// swagger:model svm_migration_ip_interface_placement_inline_ip_interfaces_inline_array_item
type SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItem struct {

	// interface
	Interface *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterface `json:"interface,omitempty"`

	// port
	Port *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePort `json:"port,omitempty"`
}

// Validate validates this svm migration ip interface placement inline ip interfaces inline array item
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItem) validateInterface(formats strfmt.Registry) error {
	if swag.IsZero(m.Interface) { // not required
		return nil
	}

	if m.Interface != nil {
		if err := m.Interface.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interface")
			}
			return err
		}
	}

	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItem) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if m.Port != nil {
		if err := m.Port.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("port")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("port")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm migration ip interface placement inline ip interfaces inline array item based on the context it is used
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInterface(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItem) contextValidateInterface(ctx context.Context, formats strfmt.Registry) error {

	if m.Interface != nil {

		if swag.IsZero(m.Interface) { // not required
			return nil
		}

		if err := m.Interface.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interface")
			}
			return err
		}
	}

	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItem) contextValidatePort(ctx context.Context, formats strfmt.Registry) error {

	if m.Port != nil {

		if swag.IsZero(m.Port) { // not required
			return nil
		}

		if err := m.Port.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("port")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("port")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterface Network interface on the source SVM.
//
// swagger:model svm_migration_ip_interface_placement_inline_ip_interfaces_inline_array_item_inline_interface
type SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterface struct {

	// links
	Links *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineLinks `json:"_links,omitempty"`

	// ip
	IP *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineIP `json:"ip,omitempty"`

	// The name of the interface. If only the name is provided, the SVM scope
	// must be provided by the object this object is embedded in.
	//
	// Example: lif1
	Name *string `json:"name,omitempty"`

	// The UUID that uniquely identifies the interface.
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this svm migration ip interface placement inline ip interfaces inline array item inline interface
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterface) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interface" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterface) validateIP(formats strfmt.Registry) error {
	if swag.IsZero(m.IP) { // not required
		return nil
	}

	if m.IP != nil {
		if err := m.IP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "ip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interface" + "." + "ip")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm migration ip interface placement inline ip interfaces inline array item inline interface based on the context it is used
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterface) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interface" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterface) contextValidateIP(ctx context.Context, formats strfmt.Registry) error {

	if m.IP != nil {

		if swag.IsZero(m.IP) { // not required
			return nil
		}

		if err := m.IP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "ip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interface" + "." + "ip")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterface) UnmarshalBinary(b []byte) error {
	var res SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineIP IP information
//
// swagger:model svm_migration_ip_interface_placement_inline_ip_interfaces_inline_array_item_inline_interface_inline_ip
type SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineIP struct {

	// address
	Address *IPAddressReadonly `json:"address,omitempty"`
}

// Validate validates this svm migration ip interface placement inline ip interfaces inline array item inline interface inline ip
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineIP) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "ip" + "." + "address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interface" + "." + "ip" + "." + "address")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm migration ip interface placement inline ip interfaces inline array item inline interface inline ip based on the context it is used
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineIP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineIP) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.Address != nil {

		if swag.IsZero(m.Address) { // not required
			return nil
		}

		if err := m.Address.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "ip" + "." + "address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interface" + "." + "ip" + "." + "address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineIP) UnmarshalBinary(b []byte) error {
	var res SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineLinks svm migration ip interface placement inline ip interfaces inline array item inline interface inline links
//
// swagger:model svm_migration_ip_interface_placement_inline_ip_interfaces_inline_array_item_inline_interface_inline__links
type SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this svm migration ip interface placement inline ip interfaces inline array item inline interface inline links
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interface" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm migration ip interface placement inline ip interfaces inline array item inline interface inline links based on the context it is used
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interface" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineLinks) UnmarshalBinary(b []byte) error {
	var res SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlineInterfaceInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePort Port to use for IP interface placement on the destination SVM.
//
// swagger:model svm_migration_ip_interface_placement_inline_ip_interfaces_inline_array_item_inline_port
type SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePort struct {

	// links
	Links *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineLinks `json:"_links,omitempty"`

	// name
	// Example: e1b
	Name *string `json:"name,omitempty"`

	// node
	Node *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineNode `json:"node,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this svm migration ip interface placement inline ip interfaces inline array item inline port
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePort) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("port" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("port" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePort) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("port" + "." + "node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("port" + "." + "node")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm migration ip interface placement inline ip interfaces inline array item inline port based on the context it is used
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePort) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("port" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("port" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePort) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {

		if swag.IsZero(m.Node) { // not required
			return nil
		}

		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("port" + "." + "node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("port" + "." + "node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePort) UnmarshalBinary(b []byte) error {
	var res SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineLinks svm migration ip interface placement inline ip interfaces inline array item inline port inline links
//
// swagger:model svm_migration_ip_interface_placement_inline_ip_interfaces_inline_array_item_inline_port_inline__links
type SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this svm migration ip interface placement inline ip interfaces inline array item inline port inline links
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("port" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("port" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm migration ip interface placement inline ip interfaces inline array item inline port inline links based on the context it is used
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("port" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("port" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineLinks) UnmarshalBinary(b []byte) error {
	var res SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineNode svm migration ip interface placement inline ip interfaces inline array item inline port inline node
//
// swagger:model svm_migration_ip_interface_placement_inline_ip_interfaces_inline_array_item_inline_port_inline_node
type SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineNode struct {

	// Name of node on which the port is located.
	// Example: node1
	Name *string `json:"name,omitempty"`
}

// Validate validates this svm migration ip interface placement inline ip interfaces inline array item inline port inline node
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineNode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this svm migration ip interface placement inline ip interfaces inline array item inline port inline node based on context it is used
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineNode) UnmarshalBinary(b []byte) error {
	var res SvmMigrationIPInterfacePlacementInlineIPInterfacesInlineArrayItemInlinePortInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
