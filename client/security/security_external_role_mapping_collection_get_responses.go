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

// SecurityExternalRoleMappingCollectionGetReader is a Reader for the SecurityExternalRoleMappingCollectionGet structure.
type SecurityExternalRoleMappingCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityExternalRoleMappingCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityExternalRoleMappingCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityExternalRoleMappingCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityExternalRoleMappingCollectionGetOK creates a SecurityExternalRoleMappingCollectionGetOK with default headers values
func NewSecurityExternalRoleMappingCollectionGetOK() *SecurityExternalRoleMappingCollectionGetOK {
	return &SecurityExternalRoleMappingCollectionGetOK{}
}

/*
SecurityExternalRoleMappingCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SecurityExternalRoleMappingCollectionGetOK struct {
	Payload *models.SecurityExternalRoleMappingResponse
}

// IsSuccess returns true when this security external role mapping collection get o k response has a 2xx status code
func (o *SecurityExternalRoleMappingCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security external role mapping collection get o k response has a 3xx status code
func (o *SecurityExternalRoleMappingCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security external role mapping collection get o k response has a 4xx status code
func (o *SecurityExternalRoleMappingCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security external role mapping collection get o k response has a 5xx status code
func (o *SecurityExternalRoleMappingCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security external role mapping collection get o k response a status code equal to that given
func (o *SecurityExternalRoleMappingCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security external role mapping collection get o k response
func (o *SecurityExternalRoleMappingCollectionGetOK) Code() int {
	return 200
}

func (o *SecurityExternalRoleMappingCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/external-role-mappings][%d] securityExternalRoleMappingCollectionGetOK %s", 200, payload)
}

func (o *SecurityExternalRoleMappingCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/external-role-mappings][%d] securityExternalRoleMappingCollectionGetOK %s", 200, payload)
}

func (o *SecurityExternalRoleMappingCollectionGetOK) GetPayload() *models.SecurityExternalRoleMappingResponse {
	return o.Payload
}

func (o *SecurityExternalRoleMappingCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityExternalRoleMappingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityExternalRoleMappingCollectionGetDefault creates a SecurityExternalRoleMappingCollectionGetDefault with default headers values
func NewSecurityExternalRoleMappingCollectionGetDefault(code int) *SecurityExternalRoleMappingCollectionGetDefault {
	return &SecurityExternalRoleMappingCollectionGetDefault{
		_statusCode: code,
	}
}

/*
SecurityExternalRoleMappingCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SecurityExternalRoleMappingCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security external role mapping collection get default response has a 2xx status code
func (o *SecurityExternalRoleMappingCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security external role mapping collection get default response has a 3xx status code
func (o *SecurityExternalRoleMappingCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security external role mapping collection get default response has a 4xx status code
func (o *SecurityExternalRoleMappingCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security external role mapping collection get default response has a 5xx status code
func (o *SecurityExternalRoleMappingCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security external role mapping collection get default response a status code equal to that given
func (o *SecurityExternalRoleMappingCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security external role mapping collection get default response
func (o *SecurityExternalRoleMappingCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SecurityExternalRoleMappingCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/external-role-mappings][%d] security_external_role_mapping_collection_get default %s", o._statusCode, payload)
}

func (o *SecurityExternalRoleMappingCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/external-role-mappings][%d] security_external_role_mapping_collection_get default %s", o._statusCode, payload)
}

func (o *SecurityExternalRoleMappingCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityExternalRoleMappingCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
