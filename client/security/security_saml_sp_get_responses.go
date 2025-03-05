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

// SecuritySamlSpGetReader is a Reader for the SecuritySamlSpGet structure.
type SecuritySamlSpGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecuritySamlSpGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecuritySamlSpGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecuritySamlSpGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecuritySamlSpGetOK creates a SecuritySamlSpGetOK with default headers values
func NewSecuritySamlSpGetOK() *SecuritySamlSpGetOK {
	return &SecuritySamlSpGetOK{}
}

/*
SecuritySamlSpGetOK describes a response with status code 200, with default header values.

OK
*/
type SecuritySamlSpGetOK struct {
	Payload *models.SecuritySamlSp
}

// IsSuccess returns true when this security saml sp get o k response has a 2xx status code
func (o *SecuritySamlSpGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security saml sp get o k response has a 3xx status code
func (o *SecuritySamlSpGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security saml sp get o k response has a 4xx status code
func (o *SecuritySamlSpGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security saml sp get o k response has a 5xx status code
func (o *SecuritySamlSpGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security saml sp get o k response a status code equal to that given
func (o *SecuritySamlSpGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security saml sp get o k response
func (o *SecuritySamlSpGetOK) Code() int {
	return 200
}

func (o *SecuritySamlSpGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/saml-sp][%d] securitySamlSpGetOK %s", 200, payload)
}

func (o *SecuritySamlSpGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/saml-sp][%d] securitySamlSpGetOK %s", 200, payload)
}

func (o *SecuritySamlSpGetOK) GetPayload() *models.SecuritySamlSp {
	return o.Payload
}

func (o *SecuritySamlSpGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecuritySamlSp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecuritySamlSpGetDefault creates a SecuritySamlSpGetDefault with default headers values
func NewSecuritySamlSpGetDefault(code int) *SecuritySamlSpGetDefault {
	return &SecuritySamlSpGetDefault{
		_statusCode: code,
	}
}

/*
SecuritySamlSpGetDefault describes a response with status code -1, with default header values.

Error
*/
type SecuritySamlSpGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security saml sp get default response has a 2xx status code
func (o *SecuritySamlSpGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security saml sp get default response has a 3xx status code
func (o *SecuritySamlSpGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security saml sp get default response has a 4xx status code
func (o *SecuritySamlSpGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security saml sp get default response has a 5xx status code
func (o *SecuritySamlSpGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security saml sp get default response a status code equal to that given
func (o *SecuritySamlSpGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security saml sp get default response
func (o *SecuritySamlSpGetDefault) Code() int {
	return o._statusCode
}

func (o *SecuritySamlSpGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/saml-sp][%d] security_saml_sp_get default %s", o._statusCode, payload)
}

func (o *SecuritySamlSpGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/saml-sp][%d] security_saml_sp_get default %s", o._statusCode, payload)
}

func (o *SecuritySamlSpGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecuritySamlSpGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
