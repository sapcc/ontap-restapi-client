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

	"github.com/sapcc/ontap-restapi/models"
)

// UnixGroupGetReader is a Reader for the UnixGroupGet structure.
type UnixGroupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnixGroupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnixGroupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnixGroupGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnixGroupGetOK creates a UnixGroupGetOK with default headers values
func NewUnixGroupGetOK() *UnixGroupGetOK {
	return &UnixGroupGetOK{}
}

/*
UnixGroupGetOK describes a response with status code 200, with default header values.

OK
*/
type UnixGroupGetOK struct {
	Payload *models.UnixGroup
}

// IsSuccess returns true when this unix group get o k response has a 2xx status code
func (o *UnixGroupGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unix group get o k response has a 3xx status code
func (o *UnixGroupGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unix group get o k response has a 4xx status code
func (o *UnixGroupGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unix group get o k response has a 5xx status code
func (o *UnixGroupGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unix group get o k response a status code equal to that given
func (o *UnixGroupGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the unix group get o k response
func (o *UnixGroupGetOK) Code() int {
	return 200
}

func (o *UnixGroupGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/unix-groups/{svm.uuid}/{name}][%d] unixGroupGetOK %s", 200, payload)
}

func (o *UnixGroupGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/unix-groups/{svm.uuid}/{name}][%d] unixGroupGetOK %s", 200, payload)
}

func (o *UnixGroupGetOK) GetPayload() *models.UnixGroup {
	return o.Payload
}

func (o *UnixGroupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnixGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnixGroupGetDefault creates a UnixGroupGetDefault with default headers values
func NewUnixGroupGetDefault(code int) *UnixGroupGetDefault {
	return &UnixGroupGetDefault{
		_statusCode: code,
	}
}

/*
UnixGroupGetDefault describes a response with status code -1, with default header values.

Error
*/
type UnixGroupGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this unix group get default response has a 2xx status code
func (o *UnixGroupGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unix group get default response has a 3xx status code
func (o *UnixGroupGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unix group get default response has a 4xx status code
func (o *UnixGroupGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unix group get default response has a 5xx status code
func (o *UnixGroupGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unix group get default response a status code equal to that given
func (o *UnixGroupGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the unix group get default response
func (o *UnixGroupGetDefault) Code() int {
	return o._statusCode
}

func (o *UnixGroupGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/unix-groups/{svm.uuid}/{name}][%d] unix_group_get default %s", o._statusCode, payload)
}

func (o *UnixGroupGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/unix-groups/{svm.uuid}/{name}][%d] unix_group_get default %s", o._statusCode, payload)
}

func (o *UnixGroupGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnixGroupGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
