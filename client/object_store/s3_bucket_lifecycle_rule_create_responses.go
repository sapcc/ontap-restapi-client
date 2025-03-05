// Code generated by go-swagger; DO NOT EDIT.

package object_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sapcc/ontap-restapi/models"
)

// S3BucketLifecycleRuleCreateReader is a Reader for the S3BucketLifecycleRuleCreate structure.
type S3BucketLifecycleRuleCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3BucketLifecycleRuleCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewS3BucketLifecycleRuleCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewS3BucketLifecycleRuleCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3BucketLifecycleRuleCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3BucketLifecycleRuleCreateCreated creates a S3BucketLifecycleRuleCreateCreated with default headers values
func NewS3BucketLifecycleRuleCreateCreated() *S3BucketLifecycleRuleCreateCreated {
	return &S3BucketLifecycleRuleCreateCreated{}
}

/*
S3BucketLifecycleRuleCreateCreated describes a response with status code 201, with default header values.

Created
*/
type S3BucketLifecycleRuleCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string
}

// IsSuccess returns true when this s3 bucket lifecycle rule create created response has a 2xx status code
func (o *S3BucketLifecycleRuleCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 bucket lifecycle rule create created response has a 3xx status code
func (o *S3BucketLifecycleRuleCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 bucket lifecycle rule create created response has a 4xx status code
func (o *S3BucketLifecycleRuleCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 bucket lifecycle rule create created response has a 5xx status code
func (o *S3BucketLifecycleRuleCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 bucket lifecycle rule create created response a status code equal to that given
func (o *S3BucketLifecycleRuleCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the s3 bucket lifecycle rule create created response
func (o *S3BucketLifecycleRuleCreateCreated) Code() int {
	return 201
}

func (o *S3BucketLifecycleRuleCreateCreated) Error() string {
	return fmt.Sprintf("[POST /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/rules][%d] s3BucketLifecycleRuleCreateCreated", 201)
}

func (o *S3BucketLifecycleRuleCreateCreated) String() string {
	return fmt.Sprintf("[POST /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/rules][%d] s3BucketLifecycleRuleCreateCreated", 201)
}

func (o *S3BucketLifecycleRuleCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewS3BucketLifecycleRuleCreateAccepted creates a S3BucketLifecycleRuleCreateAccepted with default headers values
func NewS3BucketLifecycleRuleCreateAccepted() *S3BucketLifecycleRuleCreateAccepted {
	return &S3BucketLifecycleRuleCreateAccepted{}
}

/*
S3BucketLifecycleRuleCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type S3BucketLifecycleRuleCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this s3 bucket lifecycle rule create accepted response has a 2xx status code
func (o *S3BucketLifecycleRuleCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 bucket lifecycle rule create accepted response has a 3xx status code
func (o *S3BucketLifecycleRuleCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 bucket lifecycle rule create accepted response has a 4xx status code
func (o *S3BucketLifecycleRuleCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 bucket lifecycle rule create accepted response has a 5xx status code
func (o *S3BucketLifecycleRuleCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 bucket lifecycle rule create accepted response a status code equal to that given
func (o *S3BucketLifecycleRuleCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the s3 bucket lifecycle rule create accepted response
func (o *S3BucketLifecycleRuleCreateAccepted) Code() int {
	return 202
}

func (o *S3BucketLifecycleRuleCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/rules][%d] s3BucketLifecycleRuleCreateAccepted %s", 202, payload)
}

func (o *S3BucketLifecycleRuleCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/rules][%d] s3BucketLifecycleRuleCreateAccepted %s", 202, payload)
}

func (o *S3BucketLifecycleRuleCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *S3BucketLifecycleRuleCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3BucketLifecycleRuleCreateDefault creates a S3BucketLifecycleRuleCreateDefault with default headers values
func NewS3BucketLifecycleRuleCreateDefault(code int) *S3BucketLifecycleRuleCreateDefault {
	return &S3BucketLifecycleRuleCreateDefault{
		_statusCode: code,
	}
}

/*
	S3BucketLifecycleRuleCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error code | Message |
| ---------- | ------- |
| 92405861   | "The specified SVM UUID or bucket UUID does not exist.";
| 92406111   | "Lifecycle Management rule \"rule1\" for bucket \"testbuck1\" in SVM \"vs0\" is invalid because of the number of tags specified. The maximum number of tags supported is four.";
| 92406113   | "Lifecycle Management rule for bucket in Vserver is invalid because there is a mismatch in one of the filter-field. Filters must be the same for a particular rule identifier.";
| 92406114   | "The Expiration action requires specifying either object_expiry_date, object_age_days or expired_object_delete_marker.";
| 92406115   | "The NonCurrentVersionExpiration action requires either new_non_current_versions or non_current_days.";
| 92406116   | "The AbortIncompleteMultipartUpload action requires after_initiation_days.";
| 92406117   | "The \"Expiration\" action cannot have both an expiry date and an age.";
| 92406118   | "Using Lifecycle Management rules requires an effective cluster version of 9.13.1 .";
| 92406120   | "The \"AbortIncompleteMultipartUpload\" action cannot be specified with tags.";
| 92406121   | "\"expired_object_delete_marker\" cannot be specified with \"tags\"."
| 92406122   | "\"expired_object_delete_marker\" cannot be specified with \"object_age_days\" or \"object_expiry_date\".";
| 92406123   | "Expiration is supported on object store volumes only, bucket \"testbuck-nas\" on SVM \"vs0\" is not an object store volume.";
| 92406126   | "Lifecycle Management rule \"rule1\" for bucket \"testbuck1\" in SVM \"vs0\" is invalid because specified tags contain one or more invalid characters. Valid characters for a tag are 0-9, A-Z, a-z, \"+\", \"-\", \"=\", \".\", \"_\", \":\", \"/\", \"@\", and \" \". Each \"=\" character present in a tag key or value must be prefixed with the \"\\\" escape character.";
| 92406127   | "Lifecycle Management rule \"<rule>\" for bucket \"<bucket>\" in SVM \"<SVM>\" is invalid because specified tag \"key\" length is greater than the maximum allowed length: 128.";
| 92406128   | "Lifecycle Management rule \"rule1\" for bucket \"testbuck1\" in SVM \"vs0\" has tags with duplicate keys. Verify that each tag has a unique key and then try the operation again.";
| 92406129   | "Lifecycle Management rule \"rule1\" for bucket \"testbuck1\" in SVM \"vs0\" has a prefix that is too long. The maximum length of the prefix is 1024 characters.";
| 92406130   | "Lifecycle Management rule \"rule1\" for bucket \"testbuck1\" in SVM \"vs0\" is invalid because the minimum object size of 10485760 is larger than or equal to the maximum object size of 10240.";
| 92406131   | "Lifecycle Management rule \"testcheck2\" for bucket \"testbuck1\" in SVM \"vs0\" cannot be created because \"non_current_days\" must be specified along with \"new_non_current_versions\".";
| 92406132   | "Lifecycle Management rule \"<rule>\" for bucket \"<bucket>\" in SVM \"<SVM>\" requires \"object_age_days\" to be greater than zero.";
| 92406132   | "Lifecycle Management rule \"<rule>\" for bucket \"<bucket>\" in SVM \"<SVM>\" requires \"new_non_current_versions\" to be greater than zero.";
| 92406132   | "Lifecycle Management rule \"<rule>\" for bucket \"<bucket>\" in SVM \"<SVM>\" requires \"after_initiation_days\" to be greater than zero.";
| 92406133   | "Lifecycle Management rule for bucket in Vserver is invalid. The object_expiry_date must be later than January 1, 1970.";
| 92406134   | "Cannot exceed the max limit of 1000 Lifecycle Management rules per bucket.";
| 92406135   | "MetroCluster is configured on cluster. Object Expiration is not supported in a MetroCluster configuration.";
| 92406136   | "Lifecycle Management rule \"rule1\" for bucket \"buck1\" in SVM \"vs0\" is invalid. The \"object_expiry_date\" must be at midnight GMT.";
| 92406139   | "Lifecycle Management rule for bucket in Vserver with action is a stale entry. Contact technical support for assistance.";
| 92406141   | "Lifecycle Management rule \"rule1\" for bucket \"buck1\" in SVM \"vs0\" with action \"Expiration\" cannot have \"expired_object_delete_marker\" disabled. To disable \"expired_object_delete_marker\", run the \"vserver object-store-server bucket lifecycle-management-rule delete\" command.";
| 92406143   | "As part of Bucket Lifecycle, cannot create fabriclink relationship on bucket \"<bucket name>\" in SVM \"<SVM name>\". Reason : <error>.";
| 92406144   | "The \"AbortIncompleteMultipartUpload\" action cannot be specified with object size.";
| 92406150   | "\"expired_object_delete_marker\" cannot be specified with \"size_less_than\".";
| 92406148   | "Lifecycle Management rule \"rule1\" for bucket \"buck1\" in SVM \"vs0\" cannot have \"new_non_current_versions\" more than 100.";
| 92406149   | "Lifecycle Management rule \"rule1\" for bucket \"buck1\" in SVM \"vs0\" requires an action to be specified. Retry the operation after adding an action.";
*/
type S3BucketLifecycleRuleCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this s3 bucket lifecycle rule create default response has a 2xx status code
func (o *S3BucketLifecycleRuleCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 bucket lifecycle rule create default response has a 3xx status code
func (o *S3BucketLifecycleRuleCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 bucket lifecycle rule create default response has a 4xx status code
func (o *S3BucketLifecycleRuleCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 bucket lifecycle rule create default response has a 5xx status code
func (o *S3BucketLifecycleRuleCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 bucket lifecycle rule create default response a status code equal to that given
func (o *S3BucketLifecycleRuleCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s3 bucket lifecycle rule create default response
func (o *S3BucketLifecycleRuleCreateDefault) Code() int {
	return o._statusCode
}

func (o *S3BucketLifecycleRuleCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/rules][%d] s3_bucket_lifecycle_rule_create default %s", o._statusCode, payload)
}

func (o *S3BucketLifecycleRuleCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/rules][%d] s3_bucket_lifecycle_rule_create default %s", o._statusCode, payload)
}

func (o *S3BucketLifecycleRuleCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3BucketLifecycleRuleCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
