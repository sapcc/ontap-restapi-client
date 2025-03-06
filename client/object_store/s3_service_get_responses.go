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

// S3ServiceGetReader is a Reader for the S3ServiceGet structure.
type S3ServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3ServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3ServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3ServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3ServiceGetOK creates a S3ServiceGetOK with default headers values
func NewS3ServiceGetOK() *S3ServiceGetOK {
	return &S3ServiceGetOK{}
}

/*
S3ServiceGetOK describes a response with status code 200, with default header values.

OK
*/
type S3ServiceGetOK struct {
	Payload *models.S3Service
}

// IsSuccess returns true when this s3 service get o k response has a 2xx status code
func (o *S3ServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 service get o k response has a 3xx status code
func (o *S3ServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 service get o k response has a 4xx status code
func (o *S3ServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 service get o k response has a 5xx status code
func (o *S3ServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 service get o k response a status code equal to that given
func (o *S3ServiceGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s3 service get o k response
func (o *S3ServiceGetOK) Code() int {
	return 200
}

func (o *S3ServiceGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}][%d] s3ServiceGetOK %s", 200, payload)
}

func (o *S3ServiceGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}][%d] s3ServiceGetOK %s", 200, payload)
}

func (o *S3ServiceGetOK) GetPayload() *models.S3Service {
	return o.Payload
}

func (o *S3ServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3ServiceGetDefault creates a S3ServiceGetDefault with default headers values
func NewS3ServiceGetDefault(code int) *S3ServiceGetDefault {
	return &S3ServiceGetDefault{
		_statusCode: code,
	}
}

/*
S3ServiceGetDefault describes a response with status code -1, with default header values.

Error
*/
type S3ServiceGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this s3 service get default response has a 2xx status code
func (o *S3ServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 service get default response has a 3xx status code
func (o *S3ServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 service get default response has a 4xx status code
func (o *S3ServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 service get default response has a 5xx status code
func (o *S3ServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 service get default response a status code equal to that given
func (o *S3ServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s3 service get default response
func (o *S3ServiceGetDefault) Code() int {
	return o._statusCode
}

func (o *S3ServiceGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}][%d] s3_service_get default %s", o._statusCode, payload)
}

func (o *S3ServiceGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}][%d] s3_service_get default %s", o._statusCode, payload)
}

func (o *S3ServiceGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3ServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
