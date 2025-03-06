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

	"github.com/sapcc/ontap-restapi-client/models"
)

// S3BucketSvmModifyReader is a Reader for the S3BucketSvmModify structure.
type S3BucketSvmModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3BucketSvmModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3BucketSvmModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewS3BucketSvmModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3BucketSvmModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3BucketSvmModifyOK creates a S3BucketSvmModifyOK with default headers values
func NewS3BucketSvmModifyOK() *S3BucketSvmModifyOK {
	return &S3BucketSvmModifyOK{}
}

/*
S3BucketSvmModifyOK describes a response with status code 200, with default header values.

OK
*/
type S3BucketSvmModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this s3 bucket svm modify o k response has a 2xx status code
func (o *S3BucketSvmModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 bucket svm modify o k response has a 3xx status code
func (o *S3BucketSvmModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 bucket svm modify o k response has a 4xx status code
func (o *S3BucketSvmModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 bucket svm modify o k response has a 5xx status code
func (o *S3BucketSvmModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 bucket svm modify o k response a status code equal to that given
func (o *S3BucketSvmModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s3 bucket svm modify o k response
func (o *S3BucketSvmModifyOK) Code() int {
	return 200
}

func (o *S3BucketSvmModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/buckets/{uuid}][%d] s3BucketSvmModifyOK %s", 200, payload)
}

func (o *S3BucketSvmModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/buckets/{uuid}][%d] s3BucketSvmModifyOK %s", 200, payload)
}

func (o *S3BucketSvmModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *S3BucketSvmModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3BucketSvmModifyAccepted creates a S3BucketSvmModifyAccepted with default headers values
func NewS3BucketSvmModifyAccepted() *S3BucketSvmModifyAccepted {
	return &S3BucketSvmModifyAccepted{}
}

/*
S3BucketSvmModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type S3BucketSvmModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this s3 bucket svm modify accepted response has a 2xx status code
func (o *S3BucketSvmModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 bucket svm modify accepted response has a 3xx status code
func (o *S3BucketSvmModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 bucket svm modify accepted response has a 4xx status code
func (o *S3BucketSvmModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 bucket svm modify accepted response has a 5xx status code
func (o *S3BucketSvmModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 bucket svm modify accepted response a status code equal to that given
func (o *S3BucketSvmModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the s3 bucket svm modify accepted response
func (o *S3BucketSvmModifyAccepted) Code() int {
	return 202
}

func (o *S3BucketSvmModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/buckets/{uuid}][%d] s3BucketSvmModifyAccepted %s", 202, payload)
}

func (o *S3BucketSvmModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/buckets/{uuid}][%d] s3BucketSvmModifyAccepted %s", 202, payload)
}

func (o *S3BucketSvmModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *S3BucketSvmModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3BucketSvmModifyDefault creates a S3BucketSvmModifyDefault with default headers values
func NewS3BucketSvmModifyDefault(code int) *S3BucketSvmModifyDefault {
	return &S3BucketSvmModifyDefault{
		_statusCode: code,
	}
}

/*
	S3BucketSvmModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error code | Message |
| ---------- | ------- |
| 92405778   | "Failed to modify bucket \\\"{bucket name}\\\" for SVM \\\"{svm.name}\\\". Reason: {Reason for failure}. ";
| 92405846   | "Failed to modify the object store volume. Reason: {Reason for failure}.";
| 92405811   | "Failed to modify bucket \\\"{bucket name}\\\" for SVM \\\"{svm.name}\\\". Wait a few minutes and try the operation again.";
| 92405858   | "Failed to \\\"modify\\\" the \\\"bucket\\\" because the operation is only supported on data SVMs.";
| 92405861   | "The specified \\\"{parameter field}\\\", \\\"{parameter value}\\\", does not exist.";
| 92405863   | "Failed to create access policies for bucket \\\"{bucket name}\\\". Reason: \"Principal list can have a maximum of 10 entries.\". Resolve all the issues and retry the operation.";
| 92405864   | "An error occurs when deleting an access policy. The reason for failure is detailed in the error message.";
| 92405891   | "The resources specified in the access policy are not valid. Valid ways to specify a resource are \\\"*\\\", \\\"<bucket-name>\\\", \\\"<bucket-name>/.../...\\\". Valid characters for a resource are 0-9, A-Z, a-z, \\\"_\\\", \\\"+\\\", \\\",\\\", \\\";\\\", \\\":\\\", \\\";\\\", \\\"=\\\", \\\".\\\", \\\"&\\\", \\\"@\\\", \\\"?\\\", \\\"(\\\", \\\")\\\", \\\"'\\\", \\\"*\\\", \\\"!\\\", \\\"-\\\" and \\\"\\$\\\".";
| 92405894   | "Statements, principals and resources list can have a maximum of 10 entries.";
| 92405897   | "The principals specified in the access policy are not in the correct format. User name must be in between 1 and 64 characters. Valid characters for a user name are 0-9, A-Z, a-z, \\\"_\\\", \\\"+\\\", \\\"=\\\", \\\",\\\", \\\".\\\", \\\"@\\\", and \\\"-\\\". ";
| 92405898   | "The SID specified in the access policy is not valid. Valid characters for a SID are 0-9, A-Z and a-z.";
| 92405940   | "The specified condition key is not valid for operator \"ip-address\". Valid choices of keys for this operator: source-ips.";
| 92406014   | "Failed to modify event selector for bucket \\\"{bucket name}\\\". If value of either access or permission is set to none, then the other must be set to none as well.";
| 92406032   | "Modifying the NAS path for a NAS bucket is not supported.";
| 92406230   | "The value for \\\"retention.default_period\\\" parameter for object store bucket \\\"<bucket>\\\" cannot be greater than the maximum lock retention period set in the object store server for SVM \\\"<SVM>\\\". Check the maximum allowed lock retention period present in the object store server for SVM \\\"<SVM>\\\" and try the operation again.";
| 92406236   | "The value for \\\"retention.default_period\\\" parameter for object store bucket \\\"<bucket>\\\" cannot be less than the minimum lock retention period set in the object store server for SVM \\\"<SVM>\\\". Check the minimum allowed lock retention period present in the object store server for SVM \\\"<SVM>\\\" and try the operation again.";
| 92406217   | "The specified \"allowed_headers\" is not valid because it contains more than one wild card (\"*\") character.";
| 92406224   | "A Cross-Origin Resource Sharing (CORS) rule must have an origin and HTTP method specified.";
| 92406222   | "Cannot specify Cross-Origin Resource Sharing (CORS) configuration for object store bucket \\\"<bucket>\\\" on SVM \\\"<SVM>\\\". Specifying such configuration is supported on object store volumes created in ONTAP 9.8 or later releases only.";
| 92406211   | "The specified method \"DONE\" is not valid. Valid methods are GET, PUT, DELETE, HEAD, and POST.";
| 92405863   | "Failed to create CORS rules for bucket \"bb1\". Reason: \"Field \"index\" cannot be specified for this operation.\". Resolve all the issues and retry the operation.";
| 92406228   | "Cannot exceed the maximum limit of 100 Cross-Origin Resource Sharing (CORS) rules per S3 bucket \\\"<bucket>\\\" in SVM \\\"<SVM>\\\".";;
*/
type S3BucketSvmModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this s3 bucket svm modify default response has a 2xx status code
func (o *S3BucketSvmModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 bucket svm modify default response has a 3xx status code
func (o *S3BucketSvmModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 bucket svm modify default response has a 4xx status code
func (o *S3BucketSvmModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 bucket svm modify default response has a 5xx status code
func (o *S3BucketSvmModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 bucket svm modify default response a status code equal to that given
func (o *S3BucketSvmModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s3 bucket svm modify default response
func (o *S3BucketSvmModifyDefault) Code() int {
	return o._statusCode
}

func (o *S3BucketSvmModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/buckets/{uuid}][%d] s3_bucket_svm_modify default %s", o._statusCode, payload)
}

func (o *S3BucketSvmModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/buckets/{uuid}][%d] s3_bucket_svm_modify default %s", o._statusCode, payload)
}

func (o *S3BucketSvmModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3BucketSvmModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
