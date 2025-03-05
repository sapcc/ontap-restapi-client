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

// SecurityKeyManagerCollectionGetReader is a Reader for the SecurityKeyManagerCollectionGet structure.
type SecurityKeyManagerCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeyManagerCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityKeyManagerCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeyManagerCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeyManagerCollectionGetOK creates a SecurityKeyManagerCollectionGetOK with default headers values
func NewSecurityKeyManagerCollectionGetOK() *SecurityKeyManagerCollectionGetOK {
	return &SecurityKeyManagerCollectionGetOK{}
}

/*
SecurityKeyManagerCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SecurityKeyManagerCollectionGetOK struct {
	Payload *models.SecurityKeyManagerResponse
}

// IsSuccess returns true when this security key manager collection get o k response has a 2xx status code
func (o *SecurityKeyManagerCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security key manager collection get o k response has a 3xx status code
func (o *SecurityKeyManagerCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security key manager collection get o k response has a 4xx status code
func (o *SecurityKeyManagerCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security key manager collection get o k response has a 5xx status code
func (o *SecurityKeyManagerCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security key manager collection get o k response a status code equal to that given
func (o *SecurityKeyManagerCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security key manager collection get o k response
func (o *SecurityKeyManagerCollectionGetOK) Code() int {
	return 200
}

func (o *SecurityKeyManagerCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers][%d] securityKeyManagerCollectionGetOK %s", 200, payload)
}

func (o *SecurityKeyManagerCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers][%d] securityKeyManagerCollectionGetOK %s", 200, payload)
}

func (o *SecurityKeyManagerCollectionGetOK) GetPayload() *models.SecurityKeyManagerResponse {
	return o.Payload
}

func (o *SecurityKeyManagerCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityKeyManagerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityKeyManagerCollectionGetDefault creates a SecurityKeyManagerCollectionGetDefault with default headers values
func NewSecurityKeyManagerCollectionGetDefault(code int) *SecurityKeyManagerCollectionGetDefault {
	return &SecurityKeyManagerCollectionGetDefault{
		_statusCode: code,
	}
}

/*
SecurityKeyManagerCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SecurityKeyManagerCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security key manager collection get default response has a 2xx status code
func (o *SecurityKeyManagerCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security key manager collection get default response has a 3xx status code
func (o *SecurityKeyManagerCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security key manager collection get default response has a 4xx status code
func (o *SecurityKeyManagerCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security key manager collection get default response has a 5xx status code
func (o *SecurityKeyManagerCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security key manager collection get default response a status code equal to that given
func (o *SecurityKeyManagerCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security key manager collection get default response
func (o *SecurityKeyManagerCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeyManagerCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers][%d] security_key_manager_collection_get default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers][%d] security_key_manager_collection_get default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeyManagerCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
