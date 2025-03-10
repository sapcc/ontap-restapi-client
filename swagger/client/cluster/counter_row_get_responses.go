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

// CounterRowGetReader is a Reader for the CounterRowGet structure.
type CounterRowGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CounterRowGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCounterRowGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCounterRowGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCounterRowGetOK creates a CounterRowGetOK with default headers values
func NewCounterRowGetOK() *CounterRowGetOK {
	return &CounterRowGetOK{}
}

/*
CounterRowGetOK describes a response with status code 200, with default header values.

OK
*/
type CounterRowGetOK struct {
	Payload *models.CounterRow
}

// IsSuccess returns true when this counter row get o k response has a 2xx status code
func (o *CounterRowGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this counter row get o k response has a 3xx status code
func (o *CounterRowGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this counter row get o k response has a 4xx status code
func (o *CounterRowGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this counter row get o k response has a 5xx status code
func (o *CounterRowGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this counter row get o k response a status code equal to that given
func (o *CounterRowGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the counter row get o k response
func (o *CounterRowGetOK) Code() int {
	return 200
}

func (o *CounterRowGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/counter/tables/{counter_table.name}/rows/{id}][%d] counterRowGetOK %s", 200, payload)
}

func (o *CounterRowGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/counter/tables/{counter_table.name}/rows/{id}][%d] counterRowGetOK %s", 200, payload)
}

func (o *CounterRowGetOK) GetPayload() *models.CounterRow {
	return o.Payload
}

func (o *CounterRowGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CounterRow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCounterRowGetDefault creates a CounterRowGetDefault with default headers values
func NewCounterRowGetDefault(code int) *CounterRowGetDefault {
	return &CounterRowGetDefault{
		_statusCode: code,
	}
}

/*
	CounterRowGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8585320 | Table requested is not found |
| 8586228 | Invalid counter name request. |
| 8586229 | Invalid counter property request. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type CounterRowGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this counter row get default response has a 2xx status code
func (o *CounterRowGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this counter row get default response has a 3xx status code
func (o *CounterRowGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this counter row get default response has a 4xx status code
func (o *CounterRowGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this counter row get default response has a 5xx status code
func (o *CounterRowGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this counter row get default response a status code equal to that given
func (o *CounterRowGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the counter row get default response
func (o *CounterRowGetDefault) Code() int {
	return o._statusCode
}

func (o *CounterRowGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/counter/tables/{counter_table.name}/rows/{id}][%d] counter_row_get default %s", o._statusCode, payload)
}

func (o *CounterRowGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/counter/tables/{counter_table.name}/rows/{id}][%d] counter_row_get default %s", o._statusCode, payload)
}

func (o *CounterRowGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CounterRowGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
