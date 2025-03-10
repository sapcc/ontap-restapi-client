// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// S3BucketLifecycleNonCurrentVersionExpiration Information about the non-current-version-expiration lifecycle management action.
//
// swagger:model s3_bucket_lifecycle_non_current_version_expiration
type S3BucketLifecycleNonCurrentVersionExpiration struct {

	// links
	Links *S3BucketLifecycleNonCurrentVersionExpirationInlineLinks `json:"_links,omitempty"`

	// Number of latest non-current versions to be retained.
	NewNonCurrentVersions *int64 `json:"new_non_current_versions,omitempty"`

	// Number of days after which non-current versions can be deleted.
	NonCurrentDays *int64 `json:"non_current_days,omitempty"`
}

// Validate validates this s3 bucket lifecycle non current version expiration
func (m *S3BucketLifecycleNonCurrentVersionExpiration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleNonCurrentVersionExpiration) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this s3 bucket lifecycle non current version expiration based on the context it is used
func (m *S3BucketLifecycleNonCurrentVersionExpiration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleNonCurrentVersionExpiration) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *S3BucketLifecycleNonCurrentVersionExpiration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketLifecycleNonCurrentVersionExpiration) UnmarshalBinary(b []byte) error {
	var res S3BucketLifecycleNonCurrentVersionExpiration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketLifecycleNonCurrentVersionExpirationInlineLinks s3 bucket lifecycle non current version expiration inline links
//
// swagger:model s3_bucket_lifecycle_non_current_version_expiration_inline__links
type S3BucketLifecycleNonCurrentVersionExpirationInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this s3 bucket lifecycle non current version expiration inline links
func (m *S3BucketLifecycleNonCurrentVersionExpirationInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleNonCurrentVersionExpirationInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this s3 bucket lifecycle non current version expiration inline links based on the context it is used
func (m *S3BucketLifecycleNonCurrentVersionExpirationInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleNonCurrentVersionExpirationInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *S3BucketLifecycleNonCurrentVersionExpirationInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketLifecycleNonCurrentVersionExpirationInlineLinks) UnmarshalBinary(b []byte) error {
	var res S3BucketLifecycleNonCurrentVersionExpirationInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
