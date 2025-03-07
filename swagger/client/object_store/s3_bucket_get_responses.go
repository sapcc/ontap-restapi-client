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

// S3BucketGetReader is a Reader for the S3BucketGet structure.
type S3BucketGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3BucketGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3BucketGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3BucketGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3BucketGetOK creates a S3BucketGetOK with default headers values
func NewS3BucketGetOK() *S3BucketGetOK {
	return &S3BucketGetOK{}
}

/*
S3BucketGetOK describes a response with status code 200, with default header values.

OK
*/
type S3BucketGetOK struct {
	Payload *models.S3Bucket
}

// IsSuccess returns true when this s3 bucket get o k response has a 2xx status code
func (o *S3BucketGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 bucket get o k response has a 3xx status code
func (o *S3BucketGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 bucket get o k response has a 4xx status code
func (o *S3BucketGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 bucket get o k response has a 5xx status code
func (o *S3BucketGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 bucket get o k response a status code equal to that given
func (o *S3BucketGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s3 bucket get o k response
func (o *S3BucketGetOK) Code() int {
	return 200
}

func (o *S3BucketGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/buckets/{svm.uuid}/{uuid}][%d] s3BucketGetOK %s", 200, payload)
}

func (o *S3BucketGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/buckets/{svm.uuid}/{uuid}][%d] s3BucketGetOK %s", 200, payload)
}

func (o *S3BucketGetOK) GetPayload() *models.S3Bucket {
	return o.Payload
}

func (o *S3BucketGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3Bucket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3BucketGetDefault creates a S3BucketGetDefault with default headers values
func NewS3BucketGetDefault(code int) *S3BucketGetDefault {
	return &S3BucketGetDefault{
		_statusCode: code,
	}
}

/*
S3BucketGetDefault describes a response with status code -1, with default header values.

Error
*/
type S3BucketGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this s3 bucket get default response has a 2xx status code
func (o *S3BucketGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 bucket get default response has a 3xx status code
func (o *S3BucketGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 bucket get default response has a 4xx status code
func (o *S3BucketGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 bucket get default response has a 5xx status code
func (o *S3BucketGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 bucket get default response a status code equal to that given
func (o *S3BucketGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s3 bucket get default response
func (o *S3BucketGetDefault) Code() int {
	return o._statusCode
}

func (o *S3BucketGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/buckets/{svm.uuid}/{uuid}][%d] s3_bucket_get default %s", o._statusCode, payload)
}

func (o *S3BucketGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/buckets/{svm.uuid}/{uuid}][%d] s3_bucket_get default %s", o._statusCode, payload)
}

func (o *S3BucketGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3BucketGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
