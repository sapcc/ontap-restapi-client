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

// S3PolicyCreateReader is a Reader for the S3PolicyCreate structure.
type S3PolicyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3PolicyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewS3PolicyCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3PolicyCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3PolicyCreateCreated creates a S3PolicyCreateCreated with default headers values
func NewS3PolicyCreateCreated() *S3PolicyCreateCreated {
	return &S3PolicyCreateCreated{}
}

/*
S3PolicyCreateCreated describes a response with status code 201, with default header values.

Created
*/
type S3PolicyCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.S3PolicyResponse
}

// IsSuccess returns true when this s3 policy create created response has a 2xx status code
func (o *S3PolicyCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 policy create created response has a 3xx status code
func (o *S3PolicyCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 policy create created response has a 4xx status code
func (o *S3PolicyCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 policy create created response has a 5xx status code
func (o *S3PolicyCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 policy create created response a status code equal to that given
func (o *S3PolicyCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the s3 policy create created response
func (o *S3PolicyCreateCreated) Code() int {
	return 201
}

func (o *S3PolicyCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services/{svm.uuid}/policies][%d] s3PolicyCreateCreated %s", 201, payload)
}

func (o *S3PolicyCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services/{svm.uuid}/policies][%d] s3PolicyCreateCreated %s", 201, payload)
}

func (o *S3PolicyCreateCreated) GetPayload() *models.S3PolicyResponse {
	return o.Payload
}

func (o *S3PolicyCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.S3PolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3PolicyCreateDefault creates a S3PolicyCreateDefault with default headers values
func NewS3PolicyCreateDefault(code int) *S3PolicyCreateDefault {
	return &S3PolicyCreateDefault{
		_statusCode: code,
	}
}

/*
	S3PolicyCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 92405906   | The specified action name is invalid.
| 92405947   | Creating an object store server policy or statement requires an effective cluster version of 9.8 or later.
| 92405948   | Policy name is not valid. Policy names must have between 1 and 128 characters.
| 92405949   | Policy name contains invalid characters. Valid characters: 0-9, A-Z, a-z, "_", "+", "=", ",", ".", "@", and "-".
| 92405950   | Policy name already exists for SVM.
| 92405954   | Policy name is reserved for read-only policies. Cannot be used for custom policy creation.
| 92405963   | Failed to create policy statements for policy. Reason: "{reason of failure}". Resolve all issues and retry the operation.
| 92405863   | Failed to create s3 policy statements. Reason: "{reason of failure}". Valid ways to specify a resource are \"*\", \"<bucket-name>\", \"<bucket-name>/.../...\".\". Resolve all the issues and retry the operation.
*/
type S3PolicyCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this s3 policy create default response has a 2xx status code
func (o *S3PolicyCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 policy create default response has a 3xx status code
func (o *S3PolicyCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 policy create default response has a 4xx status code
func (o *S3PolicyCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 policy create default response has a 5xx status code
func (o *S3PolicyCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 policy create default response a status code equal to that given
func (o *S3PolicyCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s3 policy create default response
func (o *S3PolicyCreateDefault) Code() int {
	return o._statusCode
}

func (o *S3PolicyCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services/{svm.uuid}/policies][%d] s3_policy_create default %s", o._statusCode, payload)
}

func (o *S3PolicyCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services/{svm.uuid}/policies][%d] s3_policy_create default %s", o._statusCode, payload)
}

func (o *S3PolicyCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3PolicyCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
