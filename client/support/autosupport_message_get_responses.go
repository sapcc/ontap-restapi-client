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

	"github.com/sapcc/ontap-restapi/models"
)

// AutosupportMessageGetReader is a Reader for the AutosupportMessageGet structure.
type AutosupportMessageGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutosupportMessageGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAutosupportMessageGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAutosupportMessageGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAutosupportMessageGetOK creates a AutosupportMessageGetOK with default headers values
func NewAutosupportMessageGetOK() *AutosupportMessageGetOK {
	return &AutosupportMessageGetOK{}
}

/*
AutosupportMessageGetOK describes a response with status code 200, with default header values.

OK
*/
type AutosupportMessageGetOK struct {
	Payload *models.AutosupportMessage
}

// IsSuccess returns true when this autosupport message get o k response has a 2xx status code
func (o *AutosupportMessageGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this autosupport message get o k response has a 3xx status code
func (o *AutosupportMessageGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autosupport message get o k response has a 4xx status code
func (o *AutosupportMessageGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this autosupport message get o k response has a 5xx status code
func (o *AutosupportMessageGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this autosupport message get o k response a status code equal to that given
func (o *AutosupportMessageGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the autosupport message get o k response
func (o *AutosupportMessageGetOK) Code() int {
	return 200
}

func (o *AutosupportMessageGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/autosupport/messages/{node.uuid}/{index}/{destination}][%d] autosupportMessageGetOK %s", 200, payload)
}

func (o *AutosupportMessageGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/autosupport/messages/{node.uuid}/{index}/{destination}][%d] autosupportMessageGetOK %s", 200, payload)
}

func (o *AutosupportMessageGetOK) GetPayload() *models.AutosupportMessage {
	return o.Payload
}

func (o *AutosupportMessageGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AutosupportMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutosupportMessageGetDefault creates a AutosupportMessageGetDefault with default headers values
func NewAutosupportMessageGetDefault(code int) *AutosupportMessageGetDefault {
	return &AutosupportMessageGetDefault{
		_statusCode: code,
	}
}

/*
AutosupportMessageGetDefault describes a response with status code -1, with default header values.

Error
*/
type AutosupportMessageGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this autosupport message get default response has a 2xx status code
func (o *AutosupportMessageGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this autosupport message get default response has a 3xx status code
func (o *AutosupportMessageGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this autosupport message get default response has a 4xx status code
func (o *AutosupportMessageGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this autosupport message get default response has a 5xx status code
func (o *AutosupportMessageGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this autosupport message get default response a status code equal to that given
func (o *AutosupportMessageGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the autosupport message get default response
func (o *AutosupportMessageGetDefault) Code() int {
	return o._statusCode
}

func (o *AutosupportMessageGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/autosupport/messages/{node.uuid}/{index}/{destination}][%d] autosupport_message_get default %s", o._statusCode, payload)
}

func (o *AutosupportMessageGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/autosupport/messages/{node.uuid}/{index}/{destination}][%d] autosupport_message_get default %s", o._statusCode, payload)
}

func (o *AutosupportMessageGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AutosupportMessageGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
