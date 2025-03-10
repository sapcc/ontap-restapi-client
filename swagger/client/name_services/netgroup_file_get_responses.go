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

// NetgroupFileGetReader is a Reader for the NetgroupFileGet structure.
type NetgroupFileGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetgroupFileGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetgroupFileGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetgroupFileGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetgroupFileGetOK creates a NetgroupFileGetOK with default headers values
func NewNetgroupFileGetOK() *NetgroupFileGetOK {
	return &NetgroupFileGetOK{}
}

/*
NetgroupFileGetOK describes a response with status code 200, with default header values.

OK
*/
type NetgroupFileGetOK struct {
	Payload *models.NetgroupFile
}

// IsSuccess returns true when this netgroup file get o k response has a 2xx status code
func (o *NetgroupFileGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this netgroup file get o k response has a 3xx status code
func (o *NetgroupFileGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this netgroup file get o k response has a 4xx status code
func (o *NetgroupFileGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this netgroup file get o k response has a 5xx status code
func (o *NetgroupFileGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this netgroup file get o k response a status code equal to that given
func (o *NetgroupFileGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the netgroup file get o k response
func (o *NetgroupFileGetOK) Code() int {
	return 200
}

func (o *NetgroupFileGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/netgroup-files/{svm.uuid}][%d] netgroupFileGetOK %s", 200, payload)
}

func (o *NetgroupFileGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/netgroup-files/{svm.uuid}][%d] netgroupFileGetOK %s", 200, payload)
}

func (o *NetgroupFileGetOK) GetPayload() *models.NetgroupFile {
	return o.Payload
}

func (o *NetgroupFileGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetgroupFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetgroupFileGetDefault creates a NetgroupFileGetDefault with default headers values
func NewNetgroupFileGetDefault(code int) *NetgroupFileGetDefault {
	return &NetgroupFileGetDefault{
		_statusCode: code,
	}
}

/*
NetgroupFileGetDefault describes a response with status code -1, with default header values.

Error
*/
type NetgroupFileGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this netgroup file get default response has a 2xx status code
func (o *NetgroupFileGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this netgroup file get default response has a 3xx status code
func (o *NetgroupFileGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this netgroup file get default response has a 4xx status code
func (o *NetgroupFileGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this netgroup file get default response has a 5xx status code
func (o *NetgroupFileGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this netgroup file get default response a status code equal to that given
func (o *NetgroupFileGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the netgroup file get default response
func (o *NetgroupFileGetDefault) Code() int {
	return o._statusCode
}

func (o *NetgroupFileGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/netgroup-files/{svm.uuid}][%d] netgroup_file_get default %s", o._statusCode, payload)
}

func (o *NetgroupFileGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/netgroup-files/{svm.uuid}][%d] netgroup_file_get default %s", o._statusCode, payload)
}

func (o *NetgroupFileGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetgroupFileGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
