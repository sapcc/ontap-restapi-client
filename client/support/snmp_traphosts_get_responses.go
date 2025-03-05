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

	"github.com/sapcc/ontap-restapi/models"
)

// SnmpTraphostsGetReader is a Reader for the SnmpTraphostsGet structure.
type SnmpTraphostsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnmpTraphostsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnmpTraphostsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnmpTraphostsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnmpTraphostsGetOK creates a SnmpTraphostsGetOK with default headers values
func NewSnmpTraphostsGetOK() *SnmpTraphostsGetOK {
	return &SnmpTraphostsGetOK{}
}

/*
SnmpTraphostsGetOK describes a response with status code 200, with default header values.

OK
*/
type SnmpTraphostsGetOK struct {
	Payload *models.SnmpTraphost
}

// IsSuccess returns true when this snmp traphosts get o k response has a 2xx status code
func (o *SnmpTraphostsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snmp traphosts get o k response has a 3xx status code
func (o *SnmpTraphostsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snmp traphosts get o k response has a 4xx status code
func (o *SnmpTraphostsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snmp traphosts get o k response has a 5xx status code
func (o *SnmpTraphostsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snmp traphosts get o k response a status code equal to that given
func (o *SnmpTraphostsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snmp traphosts get o k response
func (o *SnmpTraphostsGetOK) Code() int {
	return 200
}

func (o *SnmpTraphostsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/snmp/traphosts/{host}][%d] snmpTraphostsGetOK %s", 200, payload)
}

func (o *SnmpTraphostsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/snmp/traphosts/{host}][%d] snmpTraphostsGetOK %s", 200, payload)
}

func (o *SnmpTraphostsGetOK) GetPayload() *models.SnmpTraphost {
	return o.Payload
}

func (o *SnmpTraphostsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnmpTraphost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnmpTraphostsGetDefault creates a SnmpTraphostsGetDefault with default headers values
func NewSnmpTraphostsGetDefault(code int) *SnmpTraphostsGetDefault {
	return &SnmpTraphostsGetDefault{
		_statusCode: code,
	}
}

/*
	SnmpTraphostsGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8847365 | Unknown host. |
| 9043970 | Traphost entry does not exist. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SnmpTraphostsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snmp traphosts get default response has a 2xx status code
func (o *SnmpTraphostsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snmp traphosts get default response has a 3xx status code
func (o *SnmpTraphostsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snmp traphosts get default response has a 4xx status code
func (o *SnmpTraphostsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snmp traphosts get default response has a 5xx status code
func (o *SnmpTraphostsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snmp traphosts get default response a status code equal to that given
func (o *SnmpTraphostsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snmp traphosts get default response
func (o *SnmpTraphostsGetDefault) Code() int {
	return o._statusCode
}

func (o *SnmpTraphostsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/snmp/traphosts/{host}][%d] snmp_traphosts_get default %s", o._statusCode, payload)
}

func (o *SnmpTraphostsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/snmp/traphosts/{host}][%d] snmp_traphosts_get default %s", o._statusCode, payload)
}

func (o *SnmpTraphostsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnmpTraphostsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
