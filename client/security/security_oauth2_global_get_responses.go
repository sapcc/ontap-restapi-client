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

	"github.com/sapcc/ontap-restapi-client/models"
)

// SecurityOauth2GlobalGetReader is a Reader for the SecurityOauth2GlobalGet structure.
type SecurityOauth2GlobalGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityOauth2GlobalGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityOauth2GlobalGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityOauth2GlobalGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityOauth2GlobalGetOK creates a SecurityOauth2GlobalGetOK with default headers values
func NewSecurityOauth2GlobalGetOK() *SecurityOauth2GlobalGetOK {
	return &SecurityOauth2GlobalGetOK{}
}

/*
SecurityOauth2GlobalGetOK describes a response with status code 200, with default header values.

OK
*/
type SecurityOauth2GlobalGetOK struct {
	Payload *models.SecurityOauth2Global
}

// IsSuccess returns true when this security oauth2 global get o k response has a 2xx status code
func (o *SecurityOauth2GlobalGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security oauth2 global get o k response has a 3xx status code
func (o *SecurityOauth2GlobalGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security oauth2 global get o k response has a 4xx status code
func (o *SecurityOauth2GlobalGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security oauth2 global get o k response has a 5xx status code
func (o *SecurityOauth2GlobalGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security oauth2 global get o k response a status code equal to that given
func (o *SecurityOauth2GlobalGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security oauth2 global get o k response
func (o *SecurityOauth2GlobalGetOK) Code() int {
	return 200
}

func (o *SecurityOauth2GlobalGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/oauth2][%d] securityOauth2GlobalGetOK %s", 200, payload)
}

func (o *SecurityOauth2GlobalGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/oauth2][%d] securityOauth2GlobalGetOK %s", 200, payload)
}

func (o *SecurityOauth2GlobalGetOK) GetPayload() *models.SecurityOauth2Global {
	return o.Payload
}

func (o *SecurityOauth2GlobalGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityOauth2Global)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityOauth2GlobalGetDefault creates a SecurityOauth2GlobalGetDefault with default headers values
func NewSecurityOauth2GlobalGetDefault(code int) *SecurityOauth2GlobalGetDefault {
	return &SecurityOauth2GlobalGetDefault{
		_statusCode: code,
	}
}

/*
SecurityOauth2GlobalGetDefault describes a response with status code -1, with default header values.

Error
*/
type SecurityOauth2GlobalGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security oauth2 global get default response has a 2xx status code
func (o *SecurityOauth2GlobalGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security oauth2 global get default response has a 3xx status code
func (o *SecurityOauth2GlobalGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security oauth2 global get default response has a 4xx status code
func (o *SecurityOauth2GlobalGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security oauth2 global get default response has a 5xx status code
func (o *SecurityOauth2GlobalGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security oauth2 global get default response a status code equal to that given
func (o *SecurityOauth2GlobalGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security oauth2 global get default response
func (o *SecurityOauth2GlobalGetDefault) Code() int {
	return o._statusCode
}

func (o *SecurityOauth2GlobalGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/oauth2][%d] security_oauth2_global_get default %s", o._statusCode, payload)
}

func (o *SecurityOauth2GlobalGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/oauth2][%d] security_oauth2_global_get default %s", o._statusCode, payload)
}

func (o *SecurityOauth2GlobalGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityOauth2GlobalGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
