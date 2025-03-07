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

// MetroclusterGetReader is a Reader for the MetroclusterGet structure.
type MetroclusterGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroclusterGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroclusterGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMetroclusterGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMetroclusterGetOK creates a MetroclusterGetOK with default headers values
func NewMetroclusterGetOK() *MetroclusterGetOK {
	return &MetroclusterGetOK{}
}

/*
MetroclusterGetOK describes a response with status code 200, with default header values.

OK
*/
type MetroclusterGetOK struct {
	Payload *models.Metrocluster
}

// IsSuccess returns true when this metrocluster get o k response has a 2xx status code
func (o *MetroclusterGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this metrocluster get o k response has a 3xx status code
func (o *MetroclusterGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this metrocluster get o k response has a 4xx status code
func (o *MetroclusterGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this metrocluster get o k response has a 5xx status code
func (o *MetroclusterGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this metrocluster get o k response a status code equal to that given
func (o *MetroclusterGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the metrocluster get o k response
func (o *MetroclusterGetOK) Code() int {
	return 200
}

func (o *MetroclusterGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster][%d] metroclusterGetOK %s", 200, payload)
}

func (o *MetroclusterGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster][%d] metroclusterGetOK %s", 200, payload)
}

func (o *MetroclusterGetOK) GetPayload() *models.Metrocluster {
	return o.Payload
}

func (o *MetroclusterGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Metrocluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroclusterGetDefault creates a MetroclusterGetDefault with default headers values
func NewMetroclusterGetDefault(code int) *MetroclusterGetDefault {
	return &MetroclusterGetDefault{
		_statusCode: code,
	}
}

/*
	MetroclusterGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2425734 | An internal error occurred. Wait a few minutes, and try the operation again. For further assistance, contact technical support. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type MetroclusterGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this metrocluster get default response has a 2xx status code
func (o *MetroclusterGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this metrocluster get default response has a 3xx status code
func (o *MetroclusterGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this metrocluster get default response has a 4xx status code
func (o *MetroclusterGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this metrocluster get default response has a 5xx status code
func (o *MetroclusterGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this metrocluster get default response a status code equal to that given
func (o *MetroclusterGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the metrocluster get default response
func (o *MetroclusterGetDefault) Code() int {
	return o._statusCode
}

func (o *MetroclusterGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster][%d] metrocluster_get default %s", o._statusCode, payload)
}

func (o *MetroclusterGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster][%d] metrocluster_get default %s", o._statusCode, payload)
}

func (o *MetroclusterGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MetroclusterGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
