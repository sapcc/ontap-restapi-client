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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// S3ServiceCreateReader is a Reader for the S3ServiceCreate structure.
type S3ServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3ServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewS3ServiceCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3ServiceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3ServiceCreateCreated creates a S3ServiceCreateCreated with default headers values
func NewS3ServiceCreateCreated() *S3ServiceCreateCreated {
	return &S3ServiceCreateCreated{}
}

/*
S3ServiceCreateCreated describes a response with status code 201, with default header values.

Created
*/
type S3ServiceCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.S3ServicePostResponse
}

// IsSuccess returns true when this s3 service create created response has a 2xx status code
func (o *S3ServiceCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 service create created response has a 3xx status code
func (o *S3ServiceCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 service create created response has a 4xx status code
func (o *S3ServiceCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 service create created response has a 5xx status code
func (o *S3ServiceCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 service create created response a status code equal to that given
func (o *S3ServiceCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the s3 service create created response
func (o *S3ServiceCreateCreated) Code() int {
	return 201
}

func (o *S3ServiceCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services][%d] s3ServiceCreateCreated %s", 201, payload)
}

func (o *S3ServiceCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services][%d] s3ServiceCreateCreated %s", 201, payload)
}

func (o *S3ServiceCreateCreated) GetPayload() *models.S3ServicePostResponse {
	return o.Payload
}

func (o *S3ServiceCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.S3ServicePostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3ServiceCreateDefault creates a S3ServiceCreateDefault with default headers values
func NewS3ServiceCreateDefault(code int) *S3ServiceCreateDefault {
	return &S3ServiceCreateDefault{
		_statusCode: code,
	}
}

/*
	S3ServiceCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621706    | The specified \\\"{svm.uuid}\\\" and \\\"{svm.name}\\\" refer to different SVMs.
| 92405789   | The specified object server name contains invalid characters or not a fully qualified domain name. Valid characters for an object store server name are 0-9, A-Z, a-z, \".\", and \"-\". |
| 92405790   | Object store server names must have between 3 and 253 characters. |
| 92405839   | Creating an object store server requires an effective cluster version of data ONTAP 9.7.0 or later. Upgrade all the nodes to 9.7.0 or later and try the operation again. |
| 92405853   | Failed to create the object store server because Cloud Volumes ONTAP does not support object store servers. |
| 92405863   | An error occurs when creating an S3 user or bucket. The reason for failure is detailed in the error message. Follow the error codes specified for the user or bucket endpoints to see details for the failure. |
| 92405863   | Failed to create bucket \\\"{bucket name}\\\". Reason: \"Failed to create bucket \\\"{bucket name}\\\" for SVM \\\"{svm.name}\\\". Reason: Bucket name \\\"{bucket name}\\\" contains invalid characters or invalid character combinations. Valid characters for a bucket name are 0-9, a-z, \".\", and \"-\". Invalid character combinations are \".-\", \"-.\", and \"..\". \". Resolve all the issues and retry the operation. |
| 92405863   | Failed to create bucket \\\"{bucket name}\\\". Reason: \"Failed to create bucket \\\"{bucket name}\\\" for SVM \\\"{svm.name}\\\". Reason: Invalid QoS policy group specified \\\"{qos policy}\\\". The specified QoS policy group has a min-throughput value set, and the workload being assigned resides on a platform that does not support min-throughput or the cluster is in a mixed version state and the effective cluster version of ONTAP does not support min-throughput on this platform. Resolve all the issues and retry the operation.  |
| 92405863   | Failed to create bucket \\\"{bucket name}\\\". Reason: \"User(s) \"{user name(s)}\\\" specified in the principal list do not exist for SVM \\\"{svm.name}\\\". Use the \"object-store-server user create\" command to create a user.\". Resolve all the issues and retry the operation. |
| 92405863   | Failed to create user \\\"{user name}\\\". Reason: \"SVM \"Cluster\" is not a data SVM. Specify a data SVM.\". Resolve all the issues and retry the operation. |
| 92405884   | An object store server can only be created on a data SVM. An object store server can also be created on a system SVM on a mixed platform cluster. |
| 92405903   | Failed to configure HTTPS on an object store server for SVM \\\"{svm.name}\\\". Reason: {Reason of failure}.  |
| 92405900   | Certificate not found for SVM \\\"{svm.name}\\\".  |
| 92405917   | The specified certificate name and UUID do not refer to the same certificate.   |
| 92406020   | Only certificates of type \\\"server\\\" are supported.  |
| 92406044   | Failed to set default UNIX user for SVM \\\"{svm.name}\\\". Reason: UNIX user can only be created on a Data SVM.  |
| 92406071   | S3 protocol is not present in the allowed protocol list for SVM \"{svm.name}\".  |
| 92406196   | The specified value for the \"key_time_to_live\" field cannot be greater than the maximum limit specified for the \"max_key_time_to_live\" field in the object store server. |
| 92406197   | Object store user \"user-2\" must have a non-zero value for the \"key_time_to_live\" field because the maximum limit specified for the \"max_key_time_to_live\" field in the object store server is not zero.
| 92406230   | The value for \\\"retention.default_period\\\" parameter for object store bucket \\\"<bucket>\\\" cannot be greater than the maximum lock retention period set in the object store server for SVM \\\"<SVM>\\\". Check the maximum allowed lock retention period present in the object store server for SVM \\\"<SVM>\\\" and try the operation again.
| 92406231   | One or more object store buckets exist with a default retention period greater than the \\\"max_lock_retention_period\\\" specified. Check the default retention period set for each bucket in the specified SVM and try the operation again.
| 92406236   | The value for \\\"retention.default_period\\\" parameter for object store bucket \\\"<bucket>\\\" cannot be less than the minimum lock retention period set in the object store server for SVM \\\"<SVM>\\\". Check the minimum allowed lock retention period present in the object store server for SVM \\\"<SVM>\\\" and try the operation again.
| 92406237   | One or more object store buckets exist with a default retention period less than the \\\"min_lock_retention_period\\\" specified. Check the default retention period set for each bucket in the specified SVM and try the operation again.
| 92406238   | The value for the \\\"min_lock_retention_period\\\" parameter cannot be greater than the \\\"max_lock_retention_period\\\" parameter for the object store server for SVM \\\"vs1\\\".
| 92406217   | The specified \"-allowed-headers\" in not valid because it contains more than one wild card (\"*\") character.;
| 92406224   | A Cross-Origin Resource Sharing (CORS) rule must have an origin and HTTP method specified.;
| 92406211   | The specified method \"DONE\" is not valid. Valid methods are GET, PUT, DELETE, HEAD, and POST.;
| 92405863   | Failed to create CORS rules for bucket \"bb1\". Reason: \"Field \"index\" cannot be specified for this operation.\". Resolve all the issues and retry the operation.;
| 92406228   | Cannot exceed the maximum limit of 100 Cross-Origin Resource Sharing (CORS) rules per S3 bucket \\\"<bucket>\\\" in SVM \\\"<SVM>\\\".;
*/
type S3ServiceCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this s3 service create default response has a 2xx status code
func (o *S3ServiceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 service create default response has a 3xx status code
func (o *S3ServiceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 service create default response has a 4xx status code
func (o *S3ServiceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 service create default response has a 5xx status code
func (o *S3ServiceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 service create default response a status code equal to that given
func (o *S3ServiceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s3 service create default response
func (o *S3ServiceCreateDefault) Code() int {
	return o._statusCode
}

func (o *S3ServiceCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services][%d] s3_service_create default %s", o._statusCode, payload)
}

func (o *S3ServiceCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services][%d] s3_service_create default %s", o._statusCode, payload)
}

func (o *S3ServiceCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3ServiceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
