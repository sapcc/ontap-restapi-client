// Code generated by go-swagger; DO NOT EDIT.

package security

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

// AwsKmsCollectionGetReader is a Reader for the AwsKmsCollectionGet structure.
type AwsKmsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AwsKmsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAwsKmsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAwsKmsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAwsKmsCollectionGetOK creates a AwsKmsCollectionGetOK with default headers values
func NewAwsKmsCollectionGetOK() *AwsKmsCollectionGetOK {
	return &AwsKmsCollectionGetOK{}
}

/*
AwsKmsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type AwsKmsCollectionGetOK struct {
	Payload *models.AwsKmsResponse
}

// IsSuccess returns true when this aws kms collection get o k response has a 2xx status code
func (o *AwsKmsCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aws kms collection get o k response has a 3xx status code
func (o *AwsKmsCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws kms collection get o k response has a 4xx status code
func (o *AwsKmsCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aws kms collection get o k response has a 5xx status code
func (o *AwsKmsCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aws kms collection get o k response a status code equal to that given
func (o *AwsKmsCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aws kms collection get o k response
func (o *AwsKmsCollectionGetOK) Code() int {
	return 200
}

func (o *AwsKmsCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/aws-kms][%d] awsKmsCollectionGetOK %s", 200, payload)
}

func (o *AwsKmsCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/aws-kms][%d] awsKmsCollectionGetOK %s", 200, payload)
}

func (o *AwsKmsCollectionGetOK) GetPayload() *models.AwsKmsResponse {
	return o.Payload
}

func (o *AwsKmsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AwsKmsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsKmsCollectionGetDefault creates a AwsKmsCollectionGetDefault with default headers values
func NewAwsKmsCollectionGetDefault(code int) *AwsKmsCollectionGetDefault {
	return &AwsKmsCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	AwsKmsCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65537551 | Top-level internal key protection key (KEK) unavailable on one or more nodes. |
| 65537552 | Embedded KMIP server status not available. |
| 65537915 | The Amazon Web Service Key Management Service is unreachable from one or more nodes. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AwsKmsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this aws kms collection get default response has a 2xx status code
func (o *AwsKmsCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this aws kms collection get default response has a 3xx status code
func (o *AwsKmsCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this aws kms collection get default response has a 4xx status code
func (o *AwsKmsCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this aws kms collection get default response has a 5xx status code
func (o *AwsKmsCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this aws kms collection get default response a status code equal to that given
func (o *AwsKmsCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the aws kms collection get default response
func (o *AwsKmsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *AwsKmsCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/aws-kms][%d] aws_kms_collection_get default %s", o._statusCode, payload)
}

func (o *AwsKmsCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/aws-kms][%d] aws_kms_collection_get default %s", o._statusCode, payload)
}

func (o *AwsKmsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AwsKmsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
