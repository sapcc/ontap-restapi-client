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

	"github.com/sapcc/ontap-restapi/models"
)

// SensorsGetReader is a Reader for the SensorsGet structure.
type SensorsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SensorsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSensorsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSensorsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSensorsGetOK creates a SensorsGetOK with default headers values
func NewSensorsGetOK() *SensorsGetOK {
	return &SensorsGetOK{}
}

/*
SensorsGetOK describes a response with status code 200, with default header values.

OK
*/
type SensorsGetOK struct {
	Payload *models.Sensors
}

// IsSuccess returns true when this sensors get o k response has a 2xx status code
func (o *SensorsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sensors get o k response has a 3xx status code
func (o *SensorsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sensors get o k response has a 4xx status code
func (o *SensorsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sensors get o k response has a 5xx status code
func (o *SensorsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sensors get o k response a status code equal to that given
func (o *SensorsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the sensors get o k response
func (o *SensorsGetOK) Code() int {
	return 200
}

func (o *SensorsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/sensors/{node.uuid}/{index}][%d] sensorsGetOK %s", 200, payload)
}

func (o *SensorsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/sensors/{node.uuid}/{index}][%d] sensorsGetOK %s", 200, payload)
}

func (o *SensorsGetOK) GetPayload() *models.Sensors {
	return o.Payload
}

func (o *SensorsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Sensors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSensorsGetDefault creates a SensorsGetDefault with default headers values
func NewSensorsGetDefault(code int) *SensorsGetDefault {
	return &SensorsGetDefault{
		_statusCode: code,
	}
}

/*
SensorsGetDefault describes a response with status code -1, with default header values.

Error
*/
type SensorsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this sensors get default response has a 2xx status code
func (o *SensorsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this sensors get default response has a 3xx status code
func (o *SensorsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this sensors get default response has a 4xx status code
func (o *SensorsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this sensors get default response has a 5xx status code
func (o *SensorsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this sensors get default response a status code equal to that given
func (o *SensorsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the sensors get default response
func (o *SensorsGetDefault) Code() int {
	return o._statusCode
}

func (o *SensorsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/sensors/{node.uuid}/{index}][%d] sensors_get default %s", o._statusCode, payload)
}

func (o *SensorsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/sensors/{node.uuid}/{index}][%d] sensors_get default %s", o._statusCode, payload)
}

func (o *SensorsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SensorsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
