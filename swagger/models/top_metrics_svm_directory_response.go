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

// TopMetricsSvmDirectoryResponse top metrics svm directory response
//
// swagger:model top_metrics_svm_directory_response
type TopMetricsSvmDirectoryResponse struct {

	// links
	Links *TopMetricsSvmDirectoryResponseInlineLinks `json:"_links,omitempty"`

	// notice
	Notice *TopMetricsSvmDirectoryResponseInlineNotice `json:"notice,omitempty"`

	// Number of records.
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty"`

	// List of volumes that are not included in the SVM activity tracking REST API.
	TopMetricsSvmDirectoryResponseInlineExcludedVolumes []*TopMetricsSvmDirectoryExcludedVolume `json:"excluded_volumes,omitempty"`

	// top metrics svm directory response inline records
	TopMetricsSvmDirectoryResponseInlineRecords []*TopMetricsSvmDirectory `json:"records,omitempty"`
}

// Validate validates this top metrics svm directory response
func (m *TopMetricsSvmDirectoryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopMetricsSvmDirectoryResponseInlineExcludedVolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopMetricsSvmDirectoryResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmDirectoryResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *TopMetricsSvmDirectoryResponse) validateNotice(formats strfmt.Registry) error {
	if swag.IsZero(m.Notice) { // not required
		return nil
	}

	if m.Notice != nil {
		if err := m.Notice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notice")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmDirectoryResponse) validateTopMetricsSvmDirectoryResponseInlineExcludedVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.TopMetricsSvmDirectoryResponseInlineExcludedVolumes) { // not required
		return nil
	}

	for i := 0; i < len(m.TopMetricsSvmDirectoryResponseInlineExcludedVolumes); i++ {
		if swag.IsZero(m.TopMetricsSvmDirectoryResponseInlineExcludedVolumes[i]) { // not required
			continue
		}

		if m.TopMetricsSvmDirectoryResponseInlineExcludedVolumes[i] != nil {
			if err := m.TopMetricsSvmDirectoryResponseInlineExcludedVolumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("excluded_volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("excluded_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TopMetricsSvmDirectoryResponse) validateTopMetricsSvmDirectoryResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.TopMetricsSvmDirectoryResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.TopMetricsSvmDirectoryResponseInlineRecords); i++ {
		if swag.IsZero(m.TopMetricsSvmDirectoryResponseInlineRecords[i]) { // not required
			continue
		}

		if m.TopMetricsSvmDirectoryResponseInlineRecords[i] != nil {
			if err := m.TopMetricsSvmDirectoryResponseInlineRecords[i].Validate(formats); err != nil {
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

// ContextValidate validate this top metrics svm directory response based on the context it is used
func (m *TopMetricsSvmDirectoryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTopMetricsSvmDirectoryResponseInlineExcludedVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTopMetricsSvmDirectoryResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmDirectoryResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TopMetricsSvmDirectoryResponse) contextValidateNotice(ctx context.Context, formats strfmt.Registry) error {

	if m.Notice != nil {

		if swag.IsZero(m.Notice) { // not required
			return nil
		}

		if err := m.Notice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notice")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmDirectoryResponse) contextValidateTopMetricsSvmDirectoryResponseInlineExcludedVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TopMetricsSvmDirectoryResponseInlineExcludedVolumes); i++ {

		if m.TopMetricsSvmDirectoryResponseInlineExcludedVolumes[i] != nil {

			if swag.IsZero(m.TopMetricsSvmDirectoryResponseInlineExcludedVolumes[i]) { // not required
				return nil
			}

			if err := m.TopMetricsSvmDirectoryResponseInlineExcludedVolumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("excluded_volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("excluded_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TopMetricsSvmDirectoryResponse) contextValidateTopMetricsSvmDirectoryResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TopMetricsSvmDirectoryResponseInlineRecords); i++ {

		if m.TopMetricsSvmDirectoryResponseInlineRecords[i] != nil {

			if swag.IsZero(m.TopMetricsSvmDirectoryResponseInlineRecords[i]) { // not required
				return nil
			}

			if err := m.TopMetricsSvmDirectoryResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (m *TopMetricsSvmDirectoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmDirectoryResponse) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmDirectoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmDirectoryResponseInlineLinks top metrics svm directory response inline links
//
// swagger:model top_metrics_svm_directory_response_inline__links
type TopMetricsSvmDirectoryResponseInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this top metrics svm directory response inline links
func (m *TopMetricsSvmDirectoryResponseInlineLinks) Validate(formats strfmt.Registry) error {
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

func (m *TopMetricsSvmDirectoryResponseInlineLinks) validateNext(formats strfmt.Registry) error {
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

func (m *TopMetricsSvmDirectoryResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this top metrics svm directory response inline links based on the context it is used
func (m *TopMetricsSvmDirectoryResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *TopMetricsSvmDirectoryResponseInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TopMetricsSvmDirectoryResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TopMetricsSvmDirectoryResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmDirectoryResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmDirectoryResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmDirectoryResponseInlineNotice Optional field that indicates why no records are returned by the SVM activity tracking REST API.
//
// swagger:model top_metrics_svm_directory_response_inline_notice
type TopMetricsSvmDirectoryResponseInlineNotice struct {

	// Warning code indicating why no records are returned.
	// Example: 111411207
	// Read Only: true
	Code *string `json:"code,omitempty"`

	// Details why no records are returned.
	// Example: No read/write traffic on svm.
	// Read Only: true
	Message *string `json:"message,omitempty"`
}

// Validate validates this top metrics svm directory response inline notice
func (m *TopMetricsSvmDirectoryResponseInlineNotice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this top metrics svm directory response inline notice based on the context it is used
func (m *TopMetricsSvmDirectoryResponseInlineNotice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmDirectoryResponseInlineNotice) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "notice"+"."+"code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsSvmDirectoryResponseInlineNotice) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "notice"+"."+"message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsSvmDirectoryResponseInlineNotice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmDirectoryResponseInlineNotice) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmDirectoryResponseInlineNotice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
