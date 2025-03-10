// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TopMetricsSvmFile Information about a file's IO activity.
// Example: {"_links":{"metadata":{"href":"/api/storage/volumes/02178914-5f67-11eb-b987-005056ac5da5/files/dir_1%2ffile_1?return_metadata=true"}},"iops":{"error":{"lower_bound":2,"upper_bound":8},"read":2,"write":7},"junction-path":"/fv","path":"/vol/fv/dir_1/file_1","svm":{"name":"vserver_1","uuid":"42ee3002-67dd-11ea-8508-005056a7b8ac"},"throughput":{"error":{"lower_bound":24,"upper_bound":25},"read":24,"write":11},"volume":{"name":"vol_6","uuid":"02178914-5f67-11eb-b987-005056ac5da5"}}
//
// swagger:model top_metrics_svm_file
type TopMetricsSvmFile struct {

	// links
	Links *TopMetricsSvmFileInlineLinks `json:"_links,omitempty"`

	// iops
	Iops *TopMetricsSvmFileInlineIops `json:"iops,omitempty"`

	// Junction path of the file.
	// Example: /fv
	// Read Only: true
	JunctionPath *string `json:"junction-path,omitempty"`

	// Path of the file.
	// Example: /vol/fv/dir_abc/dir_123/file_1
	// Read Only: true
	Path *string `json:"path,omitempty"`

	// svm
	Svm *TopMetricsSvmFileInlineSvm `json:"svm,omitempty"`

	// throughput
	Throughput *TopMetricsSvmFileInlineThroughput `json:"throughput,omitempty"`

	// volume
	Volume *TopMetricsSvmFileInlineVolume `json:"volume,omitempty"`
}

// Validate validates this top metrics svm file
func (m *TopMetricsSvmFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThroughput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFile) validateLinks(formats strfmt.Registry) error {
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

func (m *TopMetricsSvmFile) validateIops(formats strfmt.Registry) error {
	if swag.IsZero(m.Iops) { // not required
		return nil
	}

	if m.Iops != nil {
		if err := m.Iops.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmFile) validateSvm(formats strfmt.Registry) error {
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

func (m *TopMetricsSvmFile) validateThroughput(formats strfmt.Registry) error {
	if swag.IsZero(m.Throughput) { // not required
		return nil
	}

	if m.Throughput != nil {
		if err := m.Throughput.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmFile) validateVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.Volume) { // not required
		return nil
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this top metrics svm file based on the context it is used
func (m *TopMetricsSvmFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIops(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJunctionPath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFile) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TopMetricsSvmFile) contextValidateIops(ctx context.Context, formats strfmt.Registry) error {

	if m.Iops != nil {

		if swag.IsZero(m.Iops) { // not required
			return nil
		}

		if err := m.Iops.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmFile) contextValidateJunctionPath(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "junction-path", "body", m.JunctionPath); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsSvmFile) contextValidatePath(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsSvmFile) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TopMetricsSvmFile) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

	if m.Throughput != nil {

		if swag.IsZero(m.Throughput) { // not required
			return nil
		}

		if err := m.Throughput.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmFile) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.Volume != nil {

		if swag.IsZero(m.Volume) { // not required
			return nil
		}

		if err := m.Volume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsSvmFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmFile) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmFileInlineIops top metrics svm file inline iops
//
// swagger:model top_metrics_svm_file_inline_iops
type TopMetricsSvmFileInlineIops struct {

	// error
	Error *TopMetricValueErrorBounds `json:"error,omitempty"`

	// Average number of read operations per second.
	// Example: 5
	// Read Only: true
	Read *int64 `json:"read,omitempty"`

	// Average number of write operations per second.
	// Example: 4
	// Read Only: true
	Write *int64 `json:"write,omitempty"`
}

// Validate validates this top metrics svm file inline iops
func (m *TopMetricsSvmFileInlineIops) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFileInlineIops) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iops" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this top metrics svm file inline iops based on the context it is used
func (m *TopMetricsSvmFileInlineIops) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRead(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWrite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFileInlineIops) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {

		if swag.IsZero(m.Error) { // not required
			return nil
		}

		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iops" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmFileInlineIops) contextValidateRead(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "iops"+"."+"read", "body", m.Read); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsSvmFileInlineIops) contextValidateWrite(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "iops"+"."+"write", "body", m.Write); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsSvmFileInlineIops) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmFileInlineIops) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmFileInlineIops
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmFileInlineLinks top metrics svm file inline links
//
// swagger:model top_metrics_svm_file_inline__links
type TopMetricsSvmFileInlineLinks struct {

	// metadata
	Metadata *Href `json:"metadata,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this top metrics svm file inline links
func (m *TopMetricsSvmFileInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
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

func (m *TopMetricsSvmFileInlineLinks) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "metadata")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmFileInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this top metrics svm file inline links based on the context it is used
func (m *TopMetricsSvmFileInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
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

func (m *TopMetricsSvmFileInlineLinks) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {

		if swag.IsZero(m.Metadata) { // not required
			return nil
		}

		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "metadata")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmFileInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TopMetricsSvmFileInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmFileInlineLinks) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmFileInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmFileInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model top_metrics_svm_file_inline_svm
type TopMetricsSvmFileInlineSvm struct {

	// links
	Links *TopMetricsSvmFileInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this top metrics svm file inline svm
func (m *TopMetricsSvmFileInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFileInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this top metrics svm file inline svm based on the context it is used
func (m *TopMetricsSvmFileInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFileInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TopMetricsSvmFileInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmFileInlineSvm) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmFileInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmFileInlineSvmInlineLinks top metrics svm file inline svm inline links
//
// swagger:model top_metrics_svm_file_inline_svm_inline__links
type TopMetricsSvmFileInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this top metrics svm file inline svm inline links
func (m *TopMetricsSvmFileInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFileInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this top metrics svm file inline svm inline links based on the context it is used
func (m *TopMetricsSvmFileInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFileInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TopMetricsSvmFileInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmFileInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmFileInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmFileInlineThroughput top metrics svm file inline throughput
//
// swagger:model top_metrics_svm_file_inline_throughput
type TopMetricsSvmFileInlineThroughput struct {

	// error
	Error *TopMetricValueErrorBounds `json:"error,omitempty"`

	// Average number of read bytes received per second.
	// Example: 2
	// Read Only: true
	Read *int64 `json:"read,omitempty"`

	// Average number of write bytes received per second.
	// Example: 20
	// Read Only: true
	Write *int64 `json:"write,omitempty"`
}

// Validate validates this top metrics svm file inline throughput
func (m *TopMetricsSvmFileInlineThroughput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFileInlineThroughput) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throughput" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this top metrics svm file inline throughput based on the context it is used
func (m *TopMetricsSvmFileInlineThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRead(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWrite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFileInlineThroughput) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {

		if swag.IsZero(m.Error) { // not required
			return nil
		}

		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throughput" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmFileInlineThroughput) contextValidateRead(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "throughput"+"."+"read", "body", m.Read); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsSvmFileInlineThroughput) contextValidateWrite(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "throughput"+"."+"write", "body", m.Write); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsSvmFileInlineThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmFileInlineThroughput) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmFileInlineThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmFileInlineVolume top metrics svm file inline volume
//
// swagger:model top_metrics_svm_file_inline_volume
type TopMetricsSvmFileInlineVolume struct {

	// links
	Links *TopMetricsSvmFileInlineVolumeInlineLinks `json:"_links,omitempty"`

	// The name of the volume. This field cannot be specified in a PATCH method.
	// Example: volume1
	Name *string `json:"name,omitempty"`

	// Unique identifier for the volume. This corresponds to the instance-uuid that is exposed in the CLI and ONTAPI. It does not change due to a volume move.
	// Example: 028baa66-41bd-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this top metrics svm file inline volume
func (m *TopMetricsSvmFileInlineVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFileInlineVolume) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this top metrics svm file inline volume based on the context it is used
func (m *TopMetricsSvmFileInlineVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFileInlineVolume) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsSvmFileInlineVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmFileInlineVolume) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmFileInlineVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmFileInlineVolumeInlineLinks top metrics svm file inline volume inline links
//
// swagger:model top_metrics_svm_file_inline_volume_inline__links
type TopMetricsSvmFileInlineVolumeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this top metrics svm file inline volume inline links
func (m *TopMetricsSvmFileInlineVolumeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFileInlineVolumeInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this top metrics svm file inline volume inline links based on the context it is used
func (m *TopMetricsSvmFileInlineVolumeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFileInlineVolumeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsSvmFileInlineVolumeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmFileInlineVolumeInlineLinks) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmFileInlineVolumeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
