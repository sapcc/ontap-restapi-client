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

// UnixUserGetReader is a Reader for the UnixUserGet structure.
type UnixUserGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnixUserGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnixUserGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnixUserGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnixUserGetOK creates a UnixUserGetOK with default headers values
func NewUnixUserGetOK() *UnixUserGetOK {
	return &UnixUserGetOK{}
}

/*
UnixUserGetOK describes a response with status code 200, with default header values.

OK
*/
type UnixUserGetOK struct {
	Payload *models.UnixUser
}

// IsSuccess returns true when this unix user get o k response has a 2xx status code
func (o *UnixUserGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unix user get o k response has a 3xx status code
func (o *UnixUserGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unix user get o k response has a 4xx status code
func (o *UnixUserGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unix user get o k response has a 5xx status code
func (o *UnixUserGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unix user get o k response a status code equal to that given
func (o *UnixUserGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the unix user get o k response
func (o *UnixUserGetOK) Code() int {
	return 200
}

func (o *UnixUserGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/unix-users/{svm.uuid}/{name}][%d] unixUserGetOK %s", 200, payload)
}

func (o *UnixUserGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/unix-users/{svm.uuid}/{name}][%d] unixUserGetOK %s", 200, payload)
}

func (o *UnixUserGetOK) GetPayload() *models.UnixUser {
	return o.Payload
}

func (o *UnixUserGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnixUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnixUserGetDefault creates a UnixUserGetDefault with default headers values
func NewUnixUserGetDefault(code int) *UnixUserGetDefault {
	return &UnixUserGetDefault{
		_statusCode: code,
	}
}

/*
UnixUserGetDefault describes a response with status code -1, with default header values.

Error
*/
type UnixUserGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this unix user get default response has a 2xx status code
func (o *UnixUserGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unix user get default response has a 3xx status code
func (o *UnixUserGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unix user get default response has a 4xx status code
func (o *UnixUserGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unix user get default response has a 5xx status code
func (o *UnixUserGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unix user get default response a status code equal to that given
func (o *UnixUserGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the unix user get default response
func (o *UnixUserGetDefault) Code() int {
	return o._statusCode
}

func (o *UnixUserGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/unix-users/{svm.uuid}/{name}][%d] unix_user_get default %s", o._statusCode, payload)
}

func (o *UnixUserGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/unix-users/{svm.uuid}/{name}][%d] unix_user_get default %s", o._statusCode, payload)
}

func (o *UnixUserGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnixUserGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
