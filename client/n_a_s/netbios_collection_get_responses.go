// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NetbiosCollectionGetReader is a Reader for the NetbiosCollectionGet structure.
type NetbiosCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetbiosCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetbiosCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetbiosCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetbiosCollectionGetOK creates a NetbiosCollectionGetOK with default headers values
func NewNetbiosCollectionGetOK() *NetbiosCollectionGetOK {
	return &NetbiosCollectionGetOK{}
}

/*
NetbiosCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type NetbiosCollectionGetOK struct {
	Payload *models.NetbiosResponse
}

// IsSuccess returns true when this netbios collection get o k response has a 2xx status code
func (o *NetbiosCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this netbios collection get o k response has a 3xx status code
func (o *NetbiosCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this netbios collection get o k response has a 4xx status code
func (o *NetbiosCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this netbios collection get o k response has a 5xx status code
func (o *NetbiosCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this netbios collection get o k response a status code equal to that given
func (o *NetbiosCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the netbios collection get o k response
func (o *NetbiosCollectionGetOK) Code() int {
	return 200
}

func (o *NetbiosCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/netbios][%d] netbiosCollectionGetOK %s", 200, payload)
}

func (o *NetbiosCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/netbios][%d] netbiosCollectionGetOK %s", 200, payload)
}

func (o *NetbiosCollectionGetOK) GetPayload() *models.NetbiosResponse {
	return o.Payload
}

func (o *NetbiosCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetbiosResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetbiosCollectionGetDefault creates a NetbiosCollectionGetDefault with default headers values
func NewNetbiosCollectionGetDefault(code int) *NetbiosCollectionGetDefault {
	return &NetbiosCollectionGetDefault{
		_statusCode: code,
	}
}

/*
NetbiosCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type NetbiosCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this netbios collection get default response has a 2xx status code
func (o *NetbiosCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this netbios collection get default response has a 3xx status code
func (o *NetbiosCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this netbios collection get default response has a 4xx status code
func (o *NetbiosCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this netbios collection get default response has a 5xx status code
func (o *NetbiosCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this netbios collection get default response a status code equal to that given
func (o *NetbiosCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the netbios collection get default response
func (o *NetbiosCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *NetbiosCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/netbios][%d] netbios_collection_get default %s", o._statusCode, payload)
}

func (o *NetbiosCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/netbios][%d] netbios_collection_get default %s", o._statusCode, payload)
}

func (o *NetbiosCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetbiosCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
