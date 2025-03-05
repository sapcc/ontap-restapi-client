// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// FlexcacheGetReader is a Reader for the FlexcacheGet structure.
type FlexcacheGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FlexcacheGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFlexcacheGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFlexcacheGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFlexcacheGetOK creates a FlexcacheGetOK with default headers values
func NewFlexcacheGetOK() *FlexcacheGetOK {
	return &FlexcacheGetOK{}
}

/*
FlexcacheGetOK describes a response with status code 200, with default header values.

OK
*/
type FlexcacheGetOK struct {
	Payload *models.Flexcache
}

// IsSuccess returns true when this flexcache get o k response has a 2xx status code
func (o *FlexcacheGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this flexcache get o k response has a 3xx status code
func (o *FlexcacheGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flexcache get o k response has a 4xx status code
func (o *FlexcacheGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this flexcache get o k response has a 5xx status code
func (o *FlexcacheGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this flexcache get o k response a status code equal to that given
func (o *FlexcacheGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the flexcache get o k response
func (o *FlexcacheGetOK) Code() int {
	return 200
}

func (o *FlexcacheGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/flexcache/flexcaches/{uuid}][%d] flexcacheGetOK %s", 200, payload)
}

func (o *FlexcacheGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/flexcache/flexcaches/{uuid}][%d] flexcacheGetOK %s", 200, payload)
}

func (o *FlexcacheGetOK) GetPayload() *models.Flexcache {
	return o.Payload
}

func (o *FlexcacheGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Flexcache)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlexcacheGetDefault creates a FlexcacheGetDefault with default headers values
func NewFlexcacheGetDefault(code int) *FlexcacheGetDefault {
	return &FlexcacheGetDefault{
		_statusCode: code,
	}
}

/*
FlexcacheGetDefault describes a response with status code -1, with default header values.

Error
*/
type FlexcacheGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this flexcache get default response has a 2xx status code
func (o *FlexcacheGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this flexcache get default response has a 3xx status code
func (o *FlexcacheGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this flexcache get default response has a 4xx status code
func (o *FlexcacheGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this flexcache get default response has a 5xx status code
func (o *FlexcacheGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this flexcache get default response a status code equal to that given
func (o *FlexcacheGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the flexcache get default response
func (o *FlexcacheGetDefault) Code() int {
	return o._statusCode
}

func (o *FlexcacheGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/flexcache/flexcaches/{uuid}][%d] flexcache_get default %s", o._statusCode, payload)
}

func (o *FlexcacheGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/flexcache/flexcaches/{uuid}][%d] flexcache_get default %s", o._statusCode, payload)
}

func (o *FlexcacheGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FlexcacheGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
