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

// S3BucketSnapshotGetReader is a Reader for the S3BucketSnapshotGet structure.
type S3BucketSnapshotGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3BucketSnapshotGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3BucketSnapshotGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3BucketSnapshotGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3BucketSnapshotGetOK creates a S3BucketSnapshotGetOK with default headers values
func NewS3BucketSnapshotGetOK() *S3BucketSnapshotGetOK {
	return &S3BucketSnapshotGetOK{}
}

/*
S3BucketSnapshotGetOK describes a response with status code 200, with default header values.

OK
*/
type S3BucketSnapshotGetOK struct {
	Payload *models.S3BucketSnapshot
}

// IsSuccess returns true when this s3 bucket snapshot get o k response has a 2xx status code
func (o *S3BucketSnapshotGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 bucket snapshot get o k response has a 3xx status code
func (o *S3BucketSnapshotGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 bucket snapshot get o k response has a 4xx status code
func (o *S3BucketSnapshotGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 bucket snapshot get o k response has a 5xx status code
func (o *S3BucketSnapshotGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 bucket snapshot get o k response a status code equal to that given
func (o *S3BucketSnapshotGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s3 bucket snapshot get o k response
func (o *S3BucketSnapshotGetOK) Code() int {
	return 200
}

func (o *S3BucketSnapshotGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/snapshots/{uuid}][%d] s3BucketSnapshotGetOK %s", 200, payload)
}

func (o *S3BucketSnapshotGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/snapshots/{uuid}][%d] s3BucketSnapshotGetOK %s", 200, payload)
}

func (o *S3BucketSnapshotGetOK) GetPayload() *models.S3BucketSnapshot {
	return o.Payload
}

func (o *S3BucketSnapshotGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3BucketSnapshot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3BucketSnapshotGetDefault creates a S3BucketSnapshotGetDefault with default headers values
func NewS3BucketSnapshotGetDefault(code int) *S3BucketSnapshotGetDefault {
	return &S3BucketSnapshotGetDefault{
		_statusCode: code,
	}
}

/*
S3BucketSnapshotGetDefault describes a response with status code -1, with default header values.

Error
*/
type S3BucketSnapshotGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this s3 bucket snapshot get default response has a 2xx status code
func (o *S3BucketSnapshotGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 bucket snapshot get default response has a 3xx status code
func (o *S3BucketSnapshotGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 bucket snapshot get default response has a 4xx status code
func (o *S3BucketSnapshotGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 bucket snapshot get default response has a 5xx status code
func (o *S3BucketSnapshotGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 bucket snapshot get default response a status code equal to that given
func (o *S3BucketSnapshotGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s3 bucket snapshot get default response
func (o *S3BucketSnapshotGetDefault) Code() int {
	return o._statusCode
}

func (o *S3BucketSnapshotGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/snapshots/{uuid}][%d] s3_bucket_snapshot_get default %s", o._statusCode, payload)
}

func (o *S3BucketSnapshotGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/snapshots/{uuid}][%d] s3_bucket_snapshot_get default %s", o._statusCode, payload)
}

func (o *S3BucketSnapshotGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3BucketSnapshotGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
