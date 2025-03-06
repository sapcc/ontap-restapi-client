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

// CoredumpCollectionGetReader is a Reader for the CoredumpCollectionGet structure.
type CoredumpCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CoredumpCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCoredumpCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCoredumpCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCoredumpCollectionGetOK creates a CoredumpCollectionGetOK with default headers values
func NewCoredumpCollectionGetOK() *CoredumpCollectionGetOK {
	return &CoredumpCollectionGetOK{}
}

/*
CoredumpCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type CoredumpCollectionGetOK struct {
	Payload *models.CoredumpResponse
}

// IsSuccess returns true when this coredump collection get o k response has a 2xx status code
func (o *CoredumpCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this coredump collection get o k response has a 3xx status code
func (o *CoredumpCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this coredump collection get o k response has a 4xx status code
func (o *CoredumpCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this coredump collection get o k response has a 5xx status code
func (o *CoredumpCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this coredump collection get o k response a status code equal to that given
func (o *CoredumpCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the coredump collection get o k response
func (o *CoredumpCollectionGetOK) Code() int {
	return 200
}

func (o *CoredumpCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/coredump/coredumps][%d] coredumpCollectionGetOK %s", 200, payload)
}

func (o *CoredumpCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/coredump/coredumps][%d] coredumpCollectionGetOK %s", 200, payload)
}

func (o *CoredumpCollectionGetOK) GetPayload() *models.CoredumpResponse {
	return o.Payload
}

func (o *CoredumpCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoredumpResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCoredumpCollectionGetDefault creates a CoredumpCollectionGetDefault with default headers values
func NewCoredumpCollectionGetDefault(code int) *CoredumpCollectionGetDefault {
	return &CoredumpCollectionGetDefault{
		_statusCode: code,
	}
}

/*
CoredumpCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type CoredumpCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this coredump collection get default response has a 2xx status code
func (o *CoredumpCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this coredump collection get default response has a 3xx status code
func (o *CoredumpCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this coredump collection get default response has a 4xx status code
func (o *CoredumpCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this coredump collection get default response has a 5xx status code
func (o *CoredumpCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this coredump collection get default response a status code equal to that given
func (o *CoredumpCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the coredump collection get default response
func (o *CoredumpCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *CoredumpCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/coredump/coredumps][%d] coredump_collection_get default %s", o._statusCode, payload)
}

func (o *CoredumpCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/coredump/coredumps][%d] coredump_collection_get default %s", o._statusCode, payload)
}

func (o *CoredumpCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CoredumpCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
