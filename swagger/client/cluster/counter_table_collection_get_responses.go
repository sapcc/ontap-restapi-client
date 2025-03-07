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

// CounterTableCollectionGetReader is a Reader for the CounterTableCollectionGet structure.
type CounterTableCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CounterTableCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCounterTableCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCounterTableCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCounterTableCollectionGetOK creates a CounterTableCollectionGetOK with default headers values
func NewCounterTableCollectionGetOK() *CounterTableCollectionGetOK {
	return &CounterTableCollectionGetOK{}
}

/*
CounterTableCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type CounterTableCollectionGetOK struct {
	Payload *models.CounterTableResponse
}

// IsSuccess returns true when this counter table collection get o k response has a 2xx status code
func (o *CounterTableCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this counter table collection get o k response has a 3xx status code
func (o *CounterTableCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this counter table collection get o k response has a 4xx status code
func (o *CounterTableCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this counter table collection get o k response has a 5xx status code
func (o *CounterTableCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this counter table collection get o k response a status code equal to that given
func (o *CounterTableCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the counter table collection get o k response
func (o *CounterTableCollectionGetOK) Code() int {
	return 200
}

func (o *CounterTableCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/counter/tables][%d] counterTableCollectionGetOK %s", 200, payload)
}

func (o *CounterTableCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/counter/tables][%d] counterTableCollectionGetOK %s", 200, payload)
}

func (o *CounterTableCollectionGetOK) GetPayload() *models.CounterTableResponse {
	return o.Payload
}

func (o *CounterTableCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CounterTableResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCounterTableCollectionGetDefault creates a CounterTableCollectionGetDefault with default headers values
func NewCounterTableCollectionGetDefault(code int) *CounterTableCollectionGetDefault {
	return &CounterTableCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	CounterTableCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8585368 | The system has not completed it's initialization |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type CounterTableCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this counter table collection get default response has a 2xx status code
func (o *CounterTableCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this counter table collection get default response has a 3xx status code
func (o *CounterTableCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this counter table collection get default response has a 4xx status code
func (o *CounterTableCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this counter table collection get default response has a 5xx status code
func (o *CounterTableCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this counter table collection get default response a status code equal to that given
func (o *CounterTableCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the counter table collection get default response
func (o *CounterTableCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *CounterTableCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/counter/tables][%d] counter_table_collection_get default %s", o._statusCode, payload)
}

func (o *CounterTableCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/counter/tables][%d] counter_table_collection_get default %s", o._statusCode, payload)
}

func (o *CounterTableCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CounterTableCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
