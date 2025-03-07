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

// S3BucketModifyReader is a Reader for the S3BucketModify structure.
type S3BucketModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3BucketModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3BucketModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewS3BucketModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3BucketModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3BucketModifyOK creates a S3BucketModifyOK with default headers values
func NewS3BucketModifyOK() *S3BucketModifyOK {
	return &S3BucketModifyOK{}
}

/*
S3BucketModifyOK describes a response with status code 200, with default header values.

OK
*/
type S3BucketModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this s3 bucket modify o k response has a 2xx status code
func (o *S3BucketModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 bucket modify o k response has a 3xx status code
func (o *S3BucketModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 bucket modify o k response has a 4xx status code
func (o *S3BucketModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 bucket modify o k response has a 5xx status code
func (o *S3BucketModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 bucket modify o k response a status code equal to that given
func (o *S3BucketModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s3 bucket modify o k response
func (o *S3BucketModifyOK) Code() int {
	return 200
}

func (o *S3BucketModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/buckets/{svm.uuid}/{uuid}][%d] s3BucketModifyOK %s", 200, payload)
}

func (o *S3BucketModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/buckets/{svm.uuid}/{uuid}][%d] s3BucketModifyOK %s", 200, payload)
}

func (o *S3BucketModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *S3BucketModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3BucketModifyAccepted creates a S3BucketModifyAccepted with default headers values
func NewS3BucketModifyAccepted() *S3BucketModifyAccepted {
	return &S3BucketModifyAccepted{}
}

/*
S3BucketModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type S3BucketModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this s3 bucket modify accepted response has a 2xx status code
func (o *S3BucketModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 bucket modify accepted response has a 3xx status code
func (o *S3BucketModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 bucket modify accepted response has a 4xx status code
func (o *S3BucketModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 bucket modify accepted response has a 5xx status code
func (o *S3BucketModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 bucket modify accepted response a status code equal to that given
func (o *S3BucketModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the s3 bucket modify accepted response
func (o *S3BucketModifyAccepted) Code() int {
	return 202
}

func (o *S3BucketModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/buckets/{svm.uuid}/{uuid}][%d] s3BucketModifyAccepted %s", 202, payload)
}

func (o *S3BucketModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/buckets/{svm.uuid}/{uuid}][%d] s3BucketModifyAccepted %s", 202, payload)
}

func (o *S3BucketModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *S3BucketModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3BucketModifyDefault creates a S3BucketModifyDefault with default headers values
func NewS3BucketModifyDefault(code int) *S3BucketModifyDefault {
	return &S3BucketModifyDefault{
		_statusCode: code,
	}
}

/*
	S3BucketModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error code | Message |
| ---------- | ------- |
| 92405778   | "Failed to modify bucket \\\"{bucket name}\\\" for SVM \\\"{svm.name}\\\". Reason: {Reason for failure}. ";
| 92405846   | "Failed to modify the object store volume. Reason: {Reason for failure}.";
| 92405811   | "Failed to modify bucket \\\"{bucket name}\\\" for SVM \\\"{svm.name}\\\". Wait a few minutes and try the operation again.";
| 92405858   | "Failed to \\\"modify\\\" the \\\"bucket\\\" because the operation is only supported on data SVMs.";
| 92405861   | "The specified SVM UUID or bucket UUID does not exist.";
| 92405863   | "An error occurs when creating an access policy. The reason for failure is detailed in the error message.";
| 92405864   | "An error occurs when deleting an access policy. The reason for failure is detailed in the error message.";
| 92405891   | "The resources specified in the access policy are not valid. Valid ways to specify a resource are \\\"*\\\", \\\"<bucket-name>\\\", \\\"<bucket-name>/.../...\\\". Valid characters for a resource are 0-9, A-Z, a-z, \\\"_\\\", \\\"+\\\", \\\",\\\", \\\";\\\", \\\":\\\", \\\";\\\", \\\"=\\\", \\\".\\\", \\\"&\\\", \\\"@\\\", \\\"?\\\", \\\"(\\\", \\\")\\\", \\\"'\\\", \\\"*\\\", \\\"!\\\", \\\"-\\\" and \\\"\\$\\\".";
| 92405894   | "Statements, principals and resources list can have a maximum of 10 entries.";
| 92405897   | "The principals specified in the access policy are not in the correct format. User name must be between 1 and 64 characters. Valid characters for a user name are 0-9, A-Z, a-z, \\\"_\\\", \\\"+\\\", \\\"=\\\", \\\",\\\", \\\".\\\", \\\"@\\\", and \\\"-\\\". ";
| 92405898   | "The SID specified in the access policy is not valid. Valid characters for a SID are 0-9, A-Z and a-z.";
| 92406014   | "Failed to modify event selector for bucket \\\"{bucket name}\\\". If the value of either access or permission is set to none, they both must be set to none.";
| 92733458   | "[Job job number] Job failed: Failed to modify bucket "s3bucket1" for SVM "vs1". Reason: {Reason for failure}. ";
| 8454236    | "Could not assign qtree "qtree1" to QoS policy group "group1". Invalid QoS policy group specified "group1". The specified QoS policy group has a min-throughput value set, and the workload being assigned resides on a platform that does not support min-throughput or the cluster is in a mixed version state and the effective cluster version of ONTAP does not support min-throughput on this platform.";
| 8454323    | "Policy group with UUID "23bwegew-8eqg-121r-bjad-0050e628wq732" does not exist."
| 92406230   | "The value for \\\"retention.default_period\\\" parameter for object store bucket \\\"<bucket>\\\" cannot be greater than the maximum lock retention period set in the object store server for SVM \\\"<SVM>\\\". Check the maximum allowed lock retention period present in the object store server for SVM \\\"<SVM>\\\" and try the operation again.";
| 92406236   | "The value for \\\"retention.default_period\\\" parameter for object store bucket \\\"<bucket>\\\" cannot be less than the minimum lock retention period set in the object store server for SVM \\\"<SVM>\\\". Check the minimum allowed lock retention period present in the object store server for SVM \\\"<SVM>\\\" and try the operation again.";
| 92406217   | "The specified \"allowed_headers\" is not valid because it contains more than one wild card (\"*\") character.";
| 92406224   | "A Cross-Origin Resource Sharing (CORS) rule must have an origin and HTTP method specified.";
| 92406222   | "Cannot specify Cross-Origin Resource Sharing (CORS) configuration for object store bucket \\\"<bucket>\\\" on SVM \\\"<SVM>\\\". Specifying such configuration is supported on object store volumes created in ONTAP 9.8 or later releases only.";
| 92406211   | "The specified method \"DONE\" is not valid. Valid methods are GET, PUT, DELETE, HEAD, and POST.";
| 92405863   | "Failed to create CORS rules for bucket \"bb1\". Reason: \"Field \"index\" cannot be specified for this operation.\". Resolve all the issues and retry the operation.";
| 92406228   | "Cannot exceed the maximum limit of 100 Cross-Origin Resource Sharing (CORS) rules per S3 bucket \\\"<bucket>\\\" in SVM \\\"<SVM>\\\".";;
*/
type S3BucketModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this s3 bucket modify default response has a 2xx status code
func (o *S3BucketModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 bucket modify default response has a 3xx status code
func (o *S3BucketModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 bucket modify default response has a 4xx status code
func (o *S3BucketModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 bucket modify default response has a 5xx status code
func (o *S3BucketModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 bucket modify default response a status code equal to that given
func (o *S3BucketModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s3 bucket modify default response
func (o *S3BucketModifyDefault) Code() int {
	return o._statusCode
}

func (o *S3BucketModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/buckets/{svm.uuid}/{uuid}][%d] s3_bucket_modify default %s", o._statusCode, payload)
}

func (o *S3BucketModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/buckets/{svm.uuid}/{uuid}][%d] s3_bucket_modify default %s", o._statusCode, payload)
}

func (o *S3BucketModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3BucketModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
