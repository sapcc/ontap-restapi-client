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

// S3PolicyGetReader is a Reader for the S3PolicyGet structure.
type S3PolicyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3PolicyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3PolicyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3PolicyGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3PolicyGetOK creates a S3PolicyGetOK with default headers values
func NewS3PolicyGetOK() *S3PolicyGetOK {
	return &S3PolicyGetOK{}
}

/*
S3PolicyGetOK describes a response with status code 200, with default header values.

OK
*/
type S3PolicyGetOK struct {
	Payload *models.S3Policy
}

// IsSuccess returns true when this s3 policy get o k response has a 2xx status code
func (o *S3PolicyGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 policy get o k response has a 3xx status code
func (o *S3PolicyGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 policy get o k response has a 4xx status code
func (o *S3PolicyGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 policy get o k response has a 5xx status code
func (o *S3PolicyGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 policy get o k response a status code equal to that given
func (o *S3PolicyGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s3 policy get o k response
func (o *S3PolicyGetOK) Code() int {
	return 200
}

func (o *S3PolicyGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/policies/{name}][%d] s3PolicyGetOK %s", 200, payload)
}

func (o *S3PolicyGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/policies/{name}][%d] s3PolicyGetOK %s", 200, payload)
}

func (o *S3PolicyGetOK) GetPayload() *models.S3Policy {
	return o.Payload
}

func (o *S3PolicyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3Policy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3PolicyGetDefault creates a S3PolicyGetDefault with default headers values
func NewS3PolicyGetDefault(code int) *S3PolicyGetDefault {
	return &S3PolicyGetDefault{
		_statusCode: code,
	}
}

/*
S3PolicyGetDefault describes a response with status code -1, with default header values.

Error
*/
type S3PolicyGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this s3 policy get default response has a 2xx status code
func (o *S3PolicyGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 policy get default response has a 3xx status code
func (o *S3PolicyGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 policy get default response has a 4xx status code
func (o *S3PolicyGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 policy get default response has a 5xx status code
func (o *S3PolicyGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 policy get default response a status code equal to that given
func (o *S3PolicyGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s3 policy get default response
func (o *S3PolicyGetDefault) Code() int {
	return o._statusCode
}

func (o *S3PolicyGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/policies/{name}][%d] s3_policy_get default %s", o._statusCode, payload)
}

func (o *S3PolicyGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/policies/{name}][%d] s3_policy_get default %s", o._statusCode, payload)
}

func (o *S3PolicyGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3PolicyGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
