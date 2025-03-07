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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// InterfacesMetricsCollectionGetReader is a Reader for the InterfacesMetricsCollectionGet structure.
type InterfacesMetricsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterfacesMetricsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterfacesMetricsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewInterfacesMetricsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInterfacesMetricsCollectionGetOK creates a InterfacesMetricsCollectionGetOK with default headers values
func NewInterfacesMetricsCollectionGetOK() *InterfacesMetricsCollectionGetOK {
	return &InterfacesMetricsCollectionGetOK{}
}

/*
InterfacesMetricsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type InterfacesMetricsCollectionGetOK struct {
	Payload *models.InterfaceMetricsResponse
}

// IsSuccess returns true when this interfaces metrics collection get o k response has a 2xx status code
func (o *InterfacesMetricsCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this interfaces metrics collection get o k response has a 3xx status code
func (o *InterfacesMetricsCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this interfaces metrics collection get o k response has a 4xx status code
func (o *InterfacesMetricsCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this interfaces metrics collection get o k response has a 5xx status code
func (o *InterfacesMetricsCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this interfaces metrics collection get o k response a status code equal to that given
func (o *InterfacesMetricsCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the interfaces metrics collection get o k response
func (o *InterfacesMetricsCollectionGetOK) Code() int {
	return 200
}

func (o *InterfacesMetricsCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/interfaces/{uuid}/metrics][%d] interfacesMetricsCollectionGetOK %s", 200, payload)
}

func (o *InterfacesMetricsCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/interfaces/{uuid}/metrics][%d] interfacesMetricsCollectionGetOK %s", 200, payload)
}

func (o *InterfacesMetricsCollectionGetOK) GetPayload() *models.InterfaceMetricsResponse {
	return o.Payload
}

func (o *InterfacesMetricsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceMetricsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterfacesMetricsCollectionGetDefault creates a InterfacesMetricsCollectionGetDefault with default headers values
func NewInterfacesMetricsCollectionGetDefault(code int) *InterfacesMetricsCollectionGetDefault {
	return &InterfacesMetricsCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	InterfacesMetricsCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8585947 | No metrics are available for the requested object. |
| 8586225 | An unexpected error occurred retrieving metrics for the requested object. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type InterfacesMetricsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this interfaces metrics collection get default response has a 2xx status code
func (o *InterfacesMetricsCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this interfaces metrics collection get default response has a 3xx status code
func (o *InterfacesMetricsCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this interfaces metrics collection get default response has a 4xx status code
func (o *InterfacesMetricsCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this interfaces metrics collection get default response has a 5xx status code
func (o *InterfacesMetricsCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this interfaces metrics collection get default response a status code equal to that given
func (o *InterfacesMetricsCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the interfaces metrics collection get default response
func (o *InterfacesMetricsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *InterfacesMetricsCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/interfaces/{uuid}/metrics][%d] interfaces_metrics_collection_get default %s", o._statusCode, payload)
}

func (o *InterfacesMetricsCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/interfaces/{uuid}/metrics][%d] interfaces_metrics_collection_get default %s", o._statusCode, payload)
}

func (o *InterfacesMetricsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *InterfacesMetricsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
