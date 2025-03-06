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

// SecurityGroupGetReader is a Reader for the SecurityGroupGet structure.
type SecurityGroupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityGroupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityGroupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityGroupGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityGroupGetOK creates a SecurityGroupGetOK with default headers values
func NewSecurityGroupGetOK() *SecurityGroupGetOK {
	return &SecurityGroupGetOK{}
}

/*
SecurityGroupGetOK describes a response with status code 200, with default header values.

OK
*/
type SecurityGroupGetOK struct {
	Payload *models.SecurityGroup
}

// IsSuccess returns true when this security group get o k response has a 2xx status code
func (o *SecurityGroupGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security group get o k response has a 3xx status code
func (o *SecurityGroupGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security group get o k response has a 4xx status code
func (o *SecurityGroupGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security group get o k response has a 5xx status code
func (o *SecurityGroupGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security group get o k response a status code equal to that given
func (o *SecurityGroupGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security group get o k response
func (o *SecurityGroupGetOK) Code() int {
	return 200
}

func (o *SecurityGroupGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/groups/{owner.uuid}/{name}/{type}][%d] securityGroupGetOK %s", 200, payload)
}

func (o *SecurityGroupGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/groups/{owner.uuid}/{name}/{type}][%d] securityGroupGetOK %s", 200, payload)
}

func (o *SecurityGroupGetOK) GetPayload() *models.SecurityGroup {
	return o.Payload
}

func (o *SecurityGroupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityGroupGetDefault creates a SecurityGroupGetDefault with default headers values
func NewSecurityGroupGetDefault(code int) *SecurityGroupGetDefault {
	return &SecurityGroupGetDefault{
		_statusCode: code,
	}
}

/*
SecurityGroupGetDefault describes a response with status code -1, with default header values.

Error
*/
type SecurityGroupGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security group get default response has a 2xx status code
func (o *SecurityGroupGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security group get default response has a 3xx status code
func (o *SecurityGroupGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security group get default response has a 4xx status code
func (o *SecurityGroupGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security group get default response has a 5xx status code
func (o *SecurityGroupGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security group get default response a status code equal to that given
func (o *SecurityGroupGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security group get default response
func (o *SecurityGroupGetDefault) Code() int {
	return o._statusCode
}

func (o *SecurityGroupGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/groups/{owner.uuid}/{name}/{type}][%d] security_group_get default %s", o._statusCode, payload)
}

func (o *SecurityGroupGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/groups/{owner.uuid}/{name}/{type}][%d] security_group_get default %s", o._statusCode, payload)
}

func (o *SecurityGroupGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityGroupGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
