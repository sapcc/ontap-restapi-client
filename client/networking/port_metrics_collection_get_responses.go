// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// PortMetricsCollectionGetReader is a Reader for the PortMetricsCollectionGet structure.
type PortMetricsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PortMetricsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPortMetricsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPortMetricsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPortMetricsCollectionGetOK creates a PortMetricsCollectionGetOK with default headers values
func NewPortMetricsCollectionGetOK() *PortMetricsCollectionGetOK {
	return &PortMetricsCollectionGetOK{}
}

/*
PortMetricsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type PortMetricsCollectionGetOK struct {
	Payload *models.PortMetricsResponse
}

// IsSuccess returns true when this port metrics collection get o k response has a 2xx status code
func (o *PortMetricsCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this port metrics collection get o k response has a 3xx status code
func (o *PortMetricsCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this port metrics collection get o k response has a 4xx status code
func (o *PortMetricsCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this port metrics collection get o k response has a 5xx status code
func (o *PortMetricsCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this port metrics collection get o k response a status code equal to that given
func (o *PortMetricsCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the port metrics collection get o k response
func (o *PortMetricsCollectionGetOK) Code() int {
	return 200
}

func (o *PortMetricsCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/ports/{uuid}/metrics][%d] portMetricsCollectionGetOK %s", 200, payload)
}

func (o *PortMetricsCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/ports/{uuid}/metrics][%d] portMetricsCollectionGetOK %s", 200, payload)
}

func (o *PortMetricsCollectionGetOK) GetPayload() *models.PortMetricsResponse {
	return o.Payload
}

func (o *PortMetricsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortMetricsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPortMetricsCollectionGetDefault creates a PortMetricsCollectionGetDefault with default headers values
func NewPortMetricsCollectionGetDefault(code int) *PortMetricsCollectionGetDefault {
	return &PortMetricsCollectionGetDefault{
		_statusCode: code,
	}
}

/*
PortMetricsCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type PortMetricsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this port metrics collection get default response has a 2xx status code
func (o *PortMetricsCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this port metrics collection get default response has a 3xx status code
func (o *PortMetricsCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this port metrics collection get default response has a 4xx status code
func (o *PortMetricsCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this port metrics collection get default response has a 5xx status code
func (o *PortMetricsCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this port metrics collection get default response a status code equal to that given
func (o *PortMetricsCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the port metrics collection get default response
func (o *PortMetricsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *PortMetricsCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/ports/{uuid}/metrics][%d] port_metrics_collection_get default %s", o._statusCode, payload)
}

func (o *PortMetricsCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/ports/{uuid}/metrics][%d] port_metrics_collection_get default %s", o._statusCode, payload)
}

func (o *PortMetricsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PortMetricsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
