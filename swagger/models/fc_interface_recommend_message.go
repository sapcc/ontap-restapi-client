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

// FcInterfaceRecommendMessage fc interface recommend message
//
// swagger:model fc_interface_recommend_message
type FcInterfaceRecommendMessage struct {

	// The message code. Possible messages:
	//   ONTAP Error Response Codes
	//   | Error Code | Description |
	//   | ---------- | ----------- |
	//   | 5375959 | Network ports are disabled. |
	//   | 5375960 | Network ports are enabled, but not reporting a connected FC fabric. |
	//   | 5375961 | The limit for the number of FC network interfaces on a cluster node has been reached. |
	//   | 5375962 | The limit for the number of FC network interfaces on a port has been reached. |
	//   | 5375963 | An HA pair of cluster nodes has a discrepancy in the presence of FC ports. |
	//   | 5375964 | An HA pair of cluster nodes has a discrepancy in support for an FC data protocol. |
	//   | 5375965 | An HA pair of cluster nodes cannot be reached from the same FC fabrics. |
	//   | 5375966 | A cluster node cannot be reached from all of the FC fabrics from which other cluster nodes with FC interfaces in the SVM can be reached. |
	//   | 5375967 | The limit for the number of FC network interfaces on a cluster node has been exceeded. |
	//   | 5375968 | The limit for the number of FC network interfaces on an FC port has been exceeded. |
	//   | 5375969 | The requested number of network interfaces per FC fabric per cluster node has not been achieved. |
	//   | 5375970 | The SVM cannot be reached from all of the FC fabrics to which the cluster is connected. |
	//   | 5375971 | The limit for the number of NVMe network interfaces on a cluster node has been exceeded. |
	//   | 5375972 | The limit for the number of cluster nodes containing NVMe network interfaces for the SVM has been exceeded. |
	//   | 5375973 | The SVM can be reached from a number of FC fabrics other than what is preferred. |
	//   Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
	//
	// Example: 5375959
	// Read Only: true
	Code *string `json:"code,omitempty"`

	// The message substitution arguments.
	// Read Only: true
	FcInterfaceRecommendMessageInlineArguments []*ErrorArguments `json:"arguments,omitempty"`

	// The message text.
	// Example: Network ports are disabled.
	// Read Only: true
	Message *string `json:"message,omitempty"`

	// The severity of the message. Message severities are as follows:
	// - `error` - Messages reporting problems that must be corrected before creating the FC network interfaces.
	// - `warning` - Messages indicating issues that need rectifying in order to achieve an optimal configuration.
	// - `informational` - Messages providing relevant information for consideration.
	//
	// Example: informational
	// Read Only: true
	// Enum: ["error","warning","informational"]
	Severity *string `json:"severity,omitempty"`
}

// Validate validates this fc interface recommend message
func (m *FcInterfaceRecommendMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFcInterfaceRecommendMessageInlineArguments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcInterfaceRecommendMessage) validateFcInterfaceRecommendMessageInlineArguments(formats strfmt.Registry) error {
	if swag.IsZero(m.FcInterfaceRecommendMessageInlineArguments) { // not required
		return nil
	}

	for i := 0; i < len(m.FcInterfaceRecommendMessageInlineArguments); i++ {
		if swag.IsZero(m.FcInterfaceRecommendMessageInlineArguments[i]) { // not required
			continue
		}

		if m.FcInterfaceRecommendMessageInlineArguments[i] != nil {
			if err := m.FcInterfaceRecommendMessageInlineArguments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("arguments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("arguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var fcInterfaceRecommendMessageTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["error","warning","informational"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fcInterfaceRecommendMessageTypeSeverityPropEnum = append(fcInterfaceRecommendMessageTypeSeverityPropEnum, v)
	}
}

const (

	// FcInterfaceRecommendMessageSeverityError captures enum value "error"
	FcInterfaceRecommendMessageSeverityError string = "error"

	// FcInterfaceRecommendMessageSeverityWarning captures enum value "warning"
	FcInterfaceRecommendMessageSeverityWarning string = "warning"

	// FcInterfaceRecommendMessageSeverityInformational captures enum value "informational"
	FcInterfaceRecommendMessageSeverityInformational string = "informational"
)

// prop value enum
func (m *FcInterfaceRecommendMessage) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fcInterfaceRecommendMessageTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FcInterfaceRecommendMessage) validateSeverity(formats strfmt.Registry) error {
	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeverityEnum("severity", "body", *m.Severity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this fc interface recommend message based on the context it is used
func (m *FcInterfaceRecommendMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFcInterfaceRecommendMessageInlineArguments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeverity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcInterfaceRecommendMessage) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *FcInterfaceRecommendMessage) contextValidateFcInterfaceRecommendMessageInlineArguments(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "arguments", "body", []*ErrorArguments(m.FcInterfaceRecommendMessageInlineArguments)); err != nil {
		return err
	}

	for i := 0; i < len(m.FcInterfaceRecommendMessageInlineArguments); i++ {

		if m.FcInterfaceRecommendMessageInlineArguments[i] != nil {

			if swag.IsZero(m.FcInterfaceRecommendMessageInlineArguments[i]) { // not required
				return nil
			}

			if err := m.FcInterfaceRecommendMessageInlineArguments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("arguments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("arguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FcInterfaceRecommendMessage) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *FcInterfaceRecommendMessage) contextValidateSeverity(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcInterfaceRecommendMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcInterfaceRecommendMessage) UnmarshalBinary(b []byte) error {
	var res FcInterfaceRecommendMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
