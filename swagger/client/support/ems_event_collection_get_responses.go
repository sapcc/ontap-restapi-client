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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// EmsEventCollectionGetReader is a Reader for the EmsEventCollectionGet structure.
type EmsEventCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsEventCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmsEventCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsEventCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsEventCollectionGetOK creates a EmsEventCollectionGetOK with default headers values
func NewEmsEventCollectionGetOK() *EmsEventCollectionGetOK {
	return &EmsEventCollectionGetOK{}
}

/*
EmsEventCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type EmsEventCollectionGetOK struct {
	Payload *models.EmsEventResponse
}

// IsSuccess returns true when this ems event collection get o k response has a 2xx status code
func (o *EmsEventCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ems event collection get o k response has a 3xx status code
func (o *EmsEventCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ems event collection get o k response has a 4xx status code
func (o *EmsEventCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ems event collection get o k response has a 5xx status code
func (o *EmsEventCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ems event collection get o k response a status code equal to that given
func (o *EmsEventCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ems event collection get o k response
func (o *EmsEventCollectionGetOK) Code() int {
	return 200
}

func (o *EmsEventCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/events][%d] emsEventCollectionGetOK %s", 200, payload)
}

func (o *EmsEventCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/events][%d] emsEventCollectionGetOK %s", 200, payload)
}

func (o *EmsEventCollectionGetOK) GetPayload() *models.EmsEventResponse {
	return o.Payload
}

func (o *EmsEventCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmsEventResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmsEventCollectionGetDefault creates a EmsEventCollectionGetDefault with default headers values
func NewEmsEventCollectionGetDefault(code int) *EmsEventCollectionGetDefault {
	return &EmsEventCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	EmsEventCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 983093 | The provided filter is unknown. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type EmsEventCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ems event collection get default response has a 2xx status code
func (o *EmsEventCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ems event collection get default response has a 3xx status code
func (o *EmsEventCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ems event collection get default response has a 4xx status code
func (o *EmsEventCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ems event collection get default response has a 5xx status code
func (o *EmsEventCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ems event collection get default response a status code equal to that given
func (o *EmsEventCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ems event collection get default response
func (o *EmsEventCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *EmsEventCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/events][%d] ems_event_collection_get default %s", o._statusCode, payload)
}

func (o *EmsEventCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/events][%d] ems_event_collection_get default %s", o._statusCode, payload)
}

func (o *EmsEventCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsEventCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
