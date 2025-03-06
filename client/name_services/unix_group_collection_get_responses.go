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

	"github.com/sapcc/ontap-restapi-client/models"
)

// UnixGroupCollectionGetReader is a Reader for the UnixGroupCollectionGet structure.
type UnixGroupCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnixGroupCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnixGroupCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnixGroupCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnixGroupCollectionGetOK creates a UnixGroupCollectionGetOK with default headers values
func NewUnixGroupCollectionGetOK() *UnixGroupCollectionGetOK {
	return &UnixGroupCollectionGetOK{}
}

/*
UnixGroupCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type UnixGroupCollectionGetOK struct {
	Payload *models.UnixGroupResponse
}

// IsSuccess returns true when this unix group collection get o k response has a 2xx status code
func (o *UnixGroupCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unix group collection get o k response has a 3xx status code
func (o *UnixGroupCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unix group collection get o k response has a 4xx status code
func (o *UnixGroupCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unix group collection get o k response has a 5xx status code
func (o *UnixGroupCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unix group collection get o k response a status code equal to that given
func (o *UnixGroupCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the unix group collection get o k response
func (o *UnixGroupCollectionGetOK) Code() int {
	return 200
}

func (o *UnixGroupCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/unix-groups][%d] unixGroupCollectionGetOK %s", 200, payload)
}

func (o *UnixGroupCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/unix-groups][%d] unixGroupCollectionGetOK %s", 200, payload)
}

func (o *UnixGroupCollectionGetOK) GetPayload() *models.UnixGroupResponse {
	return o.Payload
}

func (o *UnixGroupCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnixGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnixGroupCollectionGetDefault creates a UnixGroupCollectionGetDefault with default headers values
func NewUnixGroupCollectionGetDefault(code int) *UnixGroupCollectionGetDefault {
	return &UnixGroupCollectionGetDefault{
		_statusCode: code,
	}
}

/*
UnixGroupCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type UnixGroupCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this unix group collection get default response has a 2xx status code
func (o *UnixGroupCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unix group collection get default response has a 3xx status code
func (o *UnixGroupCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unix group collection get default response has a 4xx status code
func (o *UnixGroupCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unix group collection get default response has a 5xx status code
func (o *UnixGroupCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unix group collection get default response a status code equal to that given
func (o *UnixGroupCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the unix group collection get default response
func (o *UnixGroupCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *UnixGroupCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/unix-groups][%d] unix_group_collection_get default %s", o._statusCode, payload)
}

func (o *UnixGroupCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/unix-groups][%d] unix_group_collection_get default %s", o._statusCode, payload)
}

func (o *UnixGroupCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnixGroupCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
