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

// SecurityExternalRoleMappingCreateReader is a Reader for the SecurityExternalRoleMappingCreate structure.
type SecurityExternalRoleMappingCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityExternalRoleMappingCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSecurityExternalRoleMappingCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityExternalRoleMappingCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityExternalRoleMappingCreateCreated creates a SecurityExternalRoleMappingCreateCreated with default headers values
func NewSecurityExternalRoleMappingCreateCreated() *SecurityExternalRoleMappingCreateCreated {
	return &SecurityExternalRoleMappingCreateCreated{}
}

/*
SecurityExternalRoleMappingCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SecurityExternalRoleMappingCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.SecurityExternalRoleMappingResponse
}

// IsSuccess returns true when this security external role mapping create created response has a 2xx status code
func (o *SecurityExternalRoleMappingCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security external role mapping create created response has a 3xx status code
func (o *SecurityExternalRoleMappingCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security external role mapping create created response has a 4xx status code
func (o *SecurityExternalRoleMappingCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this security external role mapping create created response has a 5xx status code
func (o *SecurityExternalRoleMappingCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this security external role mapping create created response a status code equal to that given
func (o *SecurityExternalRoleMappingCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the security external role mapping create created response
func (o *SecurityExternalRoleMappingCreateCreated) Code() int {
	return 201
}

func (o *SecurityExternalRoleMappingCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/external-role-mappings][%d] securityExternalRoleMappingCreateCreated %s", 201, payload)
}

func (o *SecurityExternalRoleMappingCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/external-role-mappings][%d] securityExternalRoleMappingCreateCreated %s", 201, payload)
}

func (o *SecurityExternalRoleMappingCreateCreated) GetPayload() *models.SecurityExternalRoleMappingResponse {
	return o.Payload
}

func (o *SecurityExternalRoleMappingCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.SecurityExternalRoleMappingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityExternalRoleMappingCreateDefault creates a SecurityExternalRoleMappingCreateDefault with default headers values
func NewSecurityExternalRoleMappingCreateDefault(code int) *SecurityExternalRoleMappingCreateDefault {
	return &SecurityExternalRoleMappingCreateDefault{
		_statusCode: code,
	}
}

/*
	SecurityExternalRoleMappingCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5636243 | Provided ONTAP role is not configured. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SecurityExternalRoleMappingCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security external role mapping create default response has a 2xx status code
func (o *SecurityExternalRoleMappingCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security external role mapping create default response has a 3xx status code
func (o *SecurityExternalRoleMappingCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security external role mapping create default response has a 4xx status code
func (o *SecurityExternalRoleMappingCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security external role mapping create default response has a 5xx status code
func (o *SecurityExternalRoleMappingCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security external role mapping create default response a status code equal to that given
func (o *SecurityExternalRoleMappingCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security external role mapping create default response
func (o *SecurityExternalRoleMappingCreateDefault) Code() int {
	return o._statusCode
}

func (o *SecurityExternalRoleMappingCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/external-role-mappings][%d] security_external_role_mapping_create default %s", o._statusCode, payload)
}

func (o *SecurityExternalRoleMappingCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/external-role-mappings][%d] security_external_role_mapping_create default %s", o._statusCode, payload)
}

func (o *SecurityExternalRoleMappingCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityExternalRoleMappingCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
