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

// S3BucketSnapshotCollectionGetReader is a Reader for the S3BucketSnapshotCollectionGet structure.
type S3BucketSnapshotCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3BucketSnapshotCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3BucketSnapshotCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3BucketSnapshotCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3BucketSnapshotCollectionGetOK creates a S3BucketSnapshotCollectionGetOK with default headers values
func NewS3BucketSnapshotCollectionGetOK() *S3BucketSnapshotCollectionGetOK {
	return &S3BucketSnapshotCollectionGetOK{}
}

/*
S3BucketSnapshotCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type S3BucketSnapshotCollectionGetOK struct {
	Payload *models.S3BucketSnapshotResponse
}

// IsSuccess returns true when this s3 bucket snapshot collection get o k response has a 2xx status code
func (o *S3BucketSnapshotCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 bucket snapshot collection get o k response has a 3xx status code
func (o *S3BucketSnapshotCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 bucket snapshot collection get o k response has a 4xx status code
func (o *S3BucketSnapshotCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 bucket snapshot collection get o k response has a 5xx status code
func (o *S3BucketSnapshotCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 bucket snapshot collection get o k response a status code equal to that given
func (o *S3BucketSnapshotCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s3 bucket snapshot collection get o k response
func (o *S3BucketSnapshotCollectionGetOK) Code() int {
	return 200
}

func (o *S3BucketSnapshotCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/snapshots][%d] s3BucketSnapshotCollectionGetOK %s", 200, payload)
}

func (o *S3BucketSnapshotCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/snapshots][%d] s3BucketSnapshotCollectionGetOK %s", 200, payload)
}

func (o *S3BucketSnapshotCollectionGetOK) GetPayload() *models.S3BucketSnapshotResponse {
	return o.Payload
}

func (o *S3BucketSnapshotCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3BucketSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3BucketSnapshotCollectionGetDefault creates a S3BucketSnapshotCollectionGetDefault with default headers values
func NewS3BucketSnapshotCollectionGetDefault(code int) *S3BucketSnapshotCollectionGetDefault {
	return &S3BucketSnapshotCollectionGetDefault{
		_statusCode: code,
	}
}

/*
S3BucketSnapshotCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type S3BucketSnapshotCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this s3 bucket snapshot collection get default response has a 2xx status code
func (o *S3BucketSnapshotCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 bucket snapshot collection get default response has a 3xx status code
func (o *S3BucketSnapshotCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 bucket snapshot collection get default response has a 4xx status code
func (o *S3BucketSnapshotCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 bucket snapshot collection get default response has a 5xx status code
func (o *S3BucketSnapshotCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 bucket snapshot collection get default response a status code equal to that given
func (o *S3BucketSnapshotCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s3 bucket snapshot collection get default response
func (o *S3BucketSnapshotCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *S3BucketSnapshotCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/snapshots][%d] s3_bucket_snapshot_collection_get default %s", o._statusCode, payload)
}

func (o *S3BucketSnapshotCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/snapshots][%d] s3_bucket_snapshot_collection_get default %s", o._statusCode, payload)
}

func (o *S3BucketSnapshotCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3BucketSnapshotCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
