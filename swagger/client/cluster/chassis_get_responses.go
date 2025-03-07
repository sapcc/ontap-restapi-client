// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// ChassisGetReader is a Reader for the ChassisGet structure.
type ChassisGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChassisGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChassisGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChassisGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChassisGetOK creates a ChassisGetOK with default headers values
func NewChassisGetOK() *ChassisGetOK {
	return &ChassisGetOK{}
}

/*
ChassisGetOK describes a response with status code 200, with default header values.

OK
*/
type ChassisGetOK struct {
	Payload *models.Chassis
}

// IsSuccess returns true when this chassis get o k response has a 2xx status code
func (o *ChassisGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this chassis get o k response has a 3xx status code
func (o *ChassisGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chassis get o k response has a 4xx status code
func (o *ChassisGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this chassis get o k response has a 5xx status code
func (o *ChassisGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this chassis get o k response a status code equal to that given
func (o *ChassisGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the chassis get o k response
func (o *ChassisGetOK) Code() int {
	return 200
}

func (o *ChassisGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/chassis/{id}][%d] chassisGetOK %s", 200, payload)
}

func (o *ChassisGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/chassis/{id}][%d] chassisGetOK %s", 200, payload)
}

func (o *ChassisGetOK) GetPayload() *models.Chassis {
	return o.Payload
}

func (o *ChassisGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Chassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChassisGetDefault creates a ChassisGetDefault with default headers values
func NewChassisGetDefault(code int) *ChassisGetDefault {
	return &ChassisGetDefault{
		_statusCode: code,
	}
}

/*
ChassisGetDefault describes a response with status code -1, with default header values.

Error
*/
type ChassisGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this chassis get default response has a 2xx status code
func (o *ChassisGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this chassis get default response has a 3xx status code
func (o *ChassisGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this chassis get default response has a 4xx status code
func (o *ChassisGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this chassis get default response has a 5xx status code
func (o *ChassisGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this chassis get default response a status code equal to that given
func (o *ChassisGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the chassis get default response
func (o *ChassisGetDefault) Code() int {
	return o._statusCode
}

func (o *ChassisGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/chassis/{id}][%d] chassis_get default %s", o._statusCode, payload)
}

func (o *ChassisGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/chassis/{id}][%d] chassis_get default %s", o._statusCode, payload)
}

func (o *ChassisGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ChassisGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
