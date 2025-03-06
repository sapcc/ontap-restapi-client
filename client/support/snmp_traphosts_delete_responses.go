// Code generated by go-swagger; DO NOT EDIT.

package support

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

// SnmpTraphostsDeleteReader is a Reader for the SnmpTraphostsDelete structure.
type SnmpTraphostsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnmpTraphostsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnmpTraphostsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnmpTraphostsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnmpTraphostsDeleteOK creates a SnmpTraphostsDeleteOK with default headers values
func NewSnmpTraphostsDeleteOK() *SnmpTraphostsDeleteOK {
	return &SnmpTraphostsDeleteOK{}
}

/*
SnmpTraphostsDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SnmpTraphostsDeleteOK struct {
}

// IsSuccess returns true when this snmp traphosts delete o k response has a 2xx status code
func (o *SnmpTraphostsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snmp traphosts delete o k response has a 3xx status code
func (o *SnmpTraphostsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snmp traphosts delete o k response has a 4xx status code
func (o *SnmpTraphostsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snmp traphosts delete o k response has a 5xx status code
func (o *SnmpTraphostsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snmp traphosts delete o k response a status code equal to that given
func (o *SnmpTraphostsDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snmp traphosts delete o k response
func (o *SnmpTraphostsDeleteOK) Code() int {
	return 200
}

func (o *SnmpTraphostsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /support/snmp/traphosts/{host}][%d] snmpTraphostsDeleteOK", 200)
}

func (o *SnmpTraphostsDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /support/snmp/traphosts/{host}][%d] snmpTraphostsDeleteOK", 200)
}

func (o *SnmpTraphostsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSnmpTraphostsDeleteDefault creates a SnmpTraphostsDeleteDefault with default headers values
func NewSnmpTraphostsDeleteDefault(code int) *SnmpTraphostsDeleteDefault {
	return &SnmpTraphostsDeleteDefault{
		_statusCode: code,
	}
}

/*
	SnmpTraphostsDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8847365 | Unknown host. |
| 9043970 | Traphost entry does not exist. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SnmpTraphostsDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snmp traphosts delete default response has a 2xx status code
func (o *SnmpTraphostsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snmp traphosts delete default response has a 3xx status code
func (o *SnmpTraphostsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snmp traphosts delete default response has a 4xx status code
func (o *SnmpTraphostsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snmp traphosts delete default response has a 5xx status code
func (o *SnmpTraphostsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snmp traphosts delete default response a status code equal to that given
func (o *SnmpTraphostsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snmp traphosts delete default response
func (o *SnmpTraphostsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SnmpTraphostsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /support/snmp/traphosts/{host}][%d] snmp_traphosts_delete default %s", o._statusCode, payload)
}

func (o *SnmpTraphostsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /support/snmp/traphosts/{host}][%d] snmp_traphosts_delete default %s", o._statusCode, payload)
}

func (o *SnmpTraphostsDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnmpTraphostsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
