// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// LdapSchemaCreateReader is a Reader for the LdapSchemaCreate structure.
type LdapSchemaCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LdapSchemaCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewLdapSchemaCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLdapSchemaCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLdapSchemaCreateCreated creates a LdapSchemaCreateCreated with default headers values
func NewLdapSchemaCreateCreated() *LdapSchemaCreateCreated {
	return &LdapSchemaCreateCreated{}
}

/*
LdapSchemaCreateCreated describes a response with status code 201, with default header values.

Created
*/
type LdapSchemaCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.LdapSchema
}

// IsSuccess returns true when this ldap schema create created response has a 2xx status code
func (o *LdapSchemaCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ldap schema create created response has a 3xx status code
func (o *LdapSchemaCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ldap schema create created response has a 4xx status code
func (o *LdapSchemaCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ldap schema create created response has a 5xx status code
func (o *LdapSchemaCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ldap schema create created response a status code equal to that given
func (o *LdapSchemaCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ldap schema create created response
func (o *LdapSchemaCreateCreated) Code() int {
	return 201
}

func (o *LdapSchemaCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /name-services/ldap-schemas][%d] ldapSchemaCreateCreated %s", 201, payload)
}

func (o *LdapSchemaCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /name-services/ldap-schemas][%d] ldapSchemaCreateCreated %s", 201, payload)
}

func (o *LdapSchemaCreateCreated) GetPayload() *models.LdapSchema {
	return o.Payload
}

func (o *LdapSchemaCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.LdapSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLdapSchemaCreateDefault creates a LdapSchemaCreateDefault with default headers values
func NewLdapSchemaCreateDefault(code int) *LdapSchemaCreateDefault {
	return &LdapSchemaCreateDefault{
		_statusCode: code,
	}
}

/*
	LdapSchemaCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621706    | The specified SVM UUID is incorrect for the specified SVM name. |
| 4915221    | LDAP schema name in use in data SVM. |
| 4915222    | LDAP schema name in use in admin SVM. |
*/
type LdapSchemaCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ldap schema create default response has a 2xx status code
func (o *LdapSchemaCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ldap schema create default response has a 3xx status code
func (o *LdapSchemaCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ldap schema create default response has a 4xx status code
func (o *LdapSchemaCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ldap schema create default response has a 5xx status code
func (o *LdapSchemaCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ldap schema create default response a status code equal to that given
func (o *LdapSchemaCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ldap schema create default response
func (o *LdapSchemaCreateDefault) Code() int {
	return o._statusCode
}

func (o *LdapSchemaCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /name-services/ldap-schemas][%d] ldap_schema_create default %s", o._statusCode, payload)
}

func (o *LdapSchemaCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /name-services/ldap-schemas][%d] ldap_schema_create default %s", o._statusCode, payload)
}

func (o *LdapSchemaCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LdapSchemaCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
