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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// SecurityExternalRoleMappingGetReader is a Reader for the SecurityExternalRoleMappingGet structure.
type SecurityExternalRoleMappingGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityExternalRoleMappingGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityExternalRoleMappingGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityExternalRoleMappingGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityExternalRoleMappingGetOK creates a SecurityExternalRoleMappingGetOK with default headers values
func NewSecurityExternalRoleMappingGetOK() *SecurityExternalRoleMappingGetOK {
	return &SecurityExternalRoleMappingGetOK{}
}

/*
SecurityExternalRoleMappingGetOK describes a response with status code 200, with default header values.

OK
*/
type SecurityExternalRoleMappingGetOK struct {
	Payload *models.SecurityExternalRoleMapping
}

// IsSuccess returns true when this security external role mapping get o k response has a 2xx status code
func (o *SecurityExternalRoleMappingGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security external role mapping get o k response has a 3xx status code
func (o *SecurityExternalRoleMappingGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security external role mapping get o k response has a 4xx status code
func (o *SecurityExternalRoleMappingGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security external role mapping get o k response has a 5xx status code
func (o *SecurityExternalRoleMappingGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security external role mapping get o k response a status code equal to that given
func (o *SecurityExternalRoleMappingGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security external role mapping get o k response
func (o *SecurityExternalRoleMappingGetOK) Code() int {
	return 200
}

func (o *SecurityExternalRoleMappingGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/external-role-mappings/{external_role}/{provider}][%d] securityExternalRoleMappingGetOK %s", 200, payload)
}

func (o *SecurityExternalRoleMappingGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/external-role-mappings/{external_role}/{provider}][%d] securityExternalRoleMappingGetOK %s", 200, payload)
}

func (o *SecurityExternalRoleMappingGetOK) GetPayload() *models.SecurityExternalRoleMapping {
	return o.Payload
}

func (o *SecurityExternalRoleMappingGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityExternalRoleMapping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityExternalRoleMappingGetDefault creates a SecurityExternalRoleMappingGetDefault with default headers values
func NewSecurityExternalRoleMappingGetDefault(code int) *SecurityExternalRoleMappingGetDefault {
	return &SecurityExternalRoleMappingGetDefault{
		_statusCode: code,
	}
}

/*
SecurityExternalRoleMappingGetDefault describes a response with status code -1, with default header values.

Error
*/
type SecurityExternalRoleMappingGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security external role mapping get default response has a 2xx status code
func (o *SecurityExternalRoleMappingGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security external role mapping get default response has a 3xx status code
func (o *SecurityExternalRoleMappingGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security external role mapping get default response has a 4xx status code
func (o *SecurityExternalRoleMappingGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security external role mapping get default response has a 5xx status code
func (o *SecurityExternalRoleMappingGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security external role mapping get default response a status code equal to that given
func (o *SecurityExternalRoleMappingGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security external role mapping get default response
func (o *SecurityExternalRoleMappingGetDefault) Code() int {
	return o._statusCode
}

func (o *SecurityExternalRoleMappingGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/external-role-mappings/{external_role}/{provider}][%d] security_external_role_mapping_get default %s", o._statusCode, payload)
}

func (o *SecurityExternalRoleMappingGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/external-role-mappings/{external_role}/{provider}][%d] security_external_role_mapping_get default %s", o._statusCode, payload)
}

func (o *SecurityExternalRoleMappingGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityExternalRoleMappingGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
