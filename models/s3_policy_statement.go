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

// S3PolicyStatement Specifies information about a single access policy statement.
//
// swagger:model s3_policy_statement
type S3PolicyStatement struct {

	// Specifies whether access is allowed or denied. If access (to allow) is not granted explicitly to a resource, access is implicitly denied. Access can also be denied explicitly to a resource, in order to make sure that a user cannot access it, even if a different policy grants access.
	// Example: allow
	// Enum: ["allow","deny"]
	Effect *string `json:"effect,omitempty"`

	// Specifies a unique statement index used to identify a particular statement. This parameter should not be specified in the POST method. A statement index is automatically generated. It is not retrieved in the GET method.
	// Read Only: true
	Index *int64 `json:"index,omitempty"`

	// For each resource, S3 supports a set of operations. The resource operations allowed or denied are identified by an action list:
	// * GetObject - retrieves objects from a bucket.
	// * PutObject - puts objects in a bucket.
	// * DeleteObject - deletes objects from a bucket.
	// * ListBucket - lists the objects in a bucket.
	// * GetBucketAcl - retrieves the access control list (ACL) of a bucket.
	// * GetObjectAcl - retrieves the access control list (ACL) of an object.
	// * ListAllMyBuckets - lists all of the buckets in a server.
	// * ListBucketMultipartUploads - lists the multipart uploads in progress for a bucket.
	// * ListMultipartUploadParts - lists the parts in a multipart upload.
	// * CreateBucket - creates a new bucket.
	// * DeleteBucket - deletes an existing bucket.
	// * GetObjectTagging - retrieves the tag set of an object.
	// * PutObjecttagging - sets the tag set for an object.
	// * DeleteObjectTagging - deletes the tag set of an object.
	// * GetBucketLocation - retrieves the location of a bucket.
	// * GetBucketVersioning - retrieves the versioning configuration of a bucket.
	// * PutBucketVersioning - modifies the versioning configuration of a bucket.
	// * ListBucketVersions - lists the object versions in a bucket.
	// * PutBucketPolicy - puts bucket policy on the bucket specified.
	// * GetBucketPolicy - retrieves the bucket policy of a bucket.
	// * DeleteBucketPolicy - deletes the policy created for a bucket.
	// The wildcard character "*" can be used to form a regular expression for specifying actions.
	//
	// Example: ["*"]
	S3PolicyStatementInlineActions []*string `json:"actions,omitempty"`

	// s3 policy statement inline resources
	// Example: ["bucket1","bucket1/*"]
	S3PolicyStatementInlineResources []*string `json:"resources,omitempty"`

	// Specifies the statement identifier which contains additional information about the statement.
	// Example: FullAccessToBucket1
	// Max Length: 256
	// Min Length: 0
	Sid *string `json:"sid,omitempty"`
}

// Validate validates this s3 policy statement
func (m *S3PolicyStatement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var s3PolicyStatementTypeEffectPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["allow","deny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		s3PolicyStatementTypeEffectPropEnum = append(s3PolicyStatementTypeEffectPropEnum, v)
	}
}

const (

	// S3PolicyStatementEffectAllow captures enum value "allow"
	S3PolicyStatementEffectAllow string = "allow"

	// S3PolicyStatementEffectDeny captures enum value "deny"
	S3PolicyStatementEffectDeny string = "deny"
)

// prop value enum
func (m *S3PolicyStatement) validateEffectEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, s3PolicyStatementTypeEffectPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *S3PolicyStatement) validateEffect(formats strfmt.Registry) error {
	if swag.IsZero(m.Effect) { // not required
		return nil
	}

	// value enum
	if err := m.validateEffectEnum("effect", "body", *m.Effect); err != nil {
		return err
	}

	return nil
}

func (m *S3PolicyStatement) validateSid(formats strfmt.Registry) error {
	if swag.IsZero(m.Sid) { // not required
		return nil
	}

	if err := validate.MinLength("sid", "body", *m.Sid, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("sid", "body", *m.Sid, 256); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this s3 policy statement based on the context it is used
func (m *S3PolicyStatement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIndex(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3PolicyStatement) contextValidateIndex(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3PolicyStatement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3PolicyStatement) UnmarshalBinary(b []byte) error {
	var res S3PolicyStatement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
