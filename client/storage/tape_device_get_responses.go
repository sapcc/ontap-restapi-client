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

// TapeDeviceGetReader is a Reader for the TapeDeviceGet structure.
type TapeDeviceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TapeDeviceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTapeDeviceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTapeDeviceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTapeDeviceGetOK creates a TapeDeviceGetOK with default headers values
func NewTapeDeviceGetOK() *TapeDeviceGetOK {
	return &TapeDeviceGetOK{}
}

/*
TapeDeviceGetOK describes a response with status code 200, with default header values.

OK
*/
type TapeDeviceGetOK struct {
	Payload *models.TapeDevice
}

// IsSuccess returns true when this tape device get o k response has a 2xx status code
func (o *TapeDeviceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tape device get o k response has a 3xx status code
func (o *TapeDeviceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tape device get o k response has a 4xx status code
func (o *TapeDeviceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tape device get o k response has a 5xx status code
func (o *TapeDeviceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tape device get o k response a status code equal to that given
func (o *TapeDeviceGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tape device get o k response
func (o *TapeDeviceGetOK) Code() int {
	return 200
}

func (o *TapeDeviceGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/tape-devices/{node.uuid}/{device_id}][%d] tapeDeviceGetOK %s", 200, payload)
}

func (o *TapeDeviceGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/tape-devices/{node.uuid}/{device_id}][%d] tapeDeviceGetOK %s", 200, payload)
}

func (o *TapeDeviceGetOK) GetPayload() *models.TapeDevice {
	return o.Payload
}

func (o *TapeDeviceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TapeDevice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTapeDeviceGetDefault creates a TapeDeviceGetDefault with default headers values
func NewTapeDeviceGetDefault(code int) *TapeDeviceGetDefault {
	return &TapeDeviceGetDefault{
		_statusCode: code,
	}
}

/*
TapeDeviceGetDefault describes a response with status code -1, with default header values.

Error
*/
type TapeDeviceGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this tape device get default response has a 2xx status code
func (o *TapeDeviceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tape device get default response has a 3xx status code
func (o *TapeDeviceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tape device get default response has a 4xx status code
func (o *TapeDeviceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tape device get default response has a 5xx status code
func (o *TapeDeviceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tape device get default response a status code equal to that given
func (o *TapeDeviceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tape device get default response
func (o *TapeDeviceGetDefault) Code() int {
	return o._statusCode
}

func (o *TapeDeviceGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/tape-devices/{node.uuid}/{device_id}][%d] tape_device_get default %s", o._statusCode, payload)
}

func (o *TapeDeviceGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/tape-devices/{node.uuid}/{device_id}][%d] tape_device_get default %s", o._statusCode, payload)
}

func (o *TapeDeviceGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TapeDeviceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
