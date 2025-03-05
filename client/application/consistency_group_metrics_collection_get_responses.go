// Code generated by go-swagger; DO NOT EDIT.

package application

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

// ConsistencyGroupMetricsCollectionGetReader is a Reader for the ConsistencyGroupMetricsCollectionGet structure.
type ConsistencyGroupMetricsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsistencyGroupMetricsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConsistencyGroupMetricsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConsistencyGroupMetricsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConsistencyGroupMetricsCollectionGetOK creates a ConsistencyGroupMetricsCollectionGetOK with default headers values
func NewConsistencyGroupMetricsCollectionGetOK() *ConsistencyGroupMetricsCollectionGetOK {
	return &ConsistencyGroupMetricsCollectionGetOK{}
}

/*
ConsistencyGroupMetricsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ConsistencyGroupMetricsCollectionGetOK struct {
	Payload *models.ConsistencyGroupMetricsResponse
}

// IsSuccess returns true when this consistency group metrics collection get o k response has a 2xx status code
func (o *ConsistencyGroupMetricsCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this consistency group metrics collection get o k response has a 3xx status code
func (o *ConsistencyGroupMetricsCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this consistency group metrics collection get o k response has a 4xx status code
func (o *ConsistencyGroupMetricsCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this consistency group metrics collection get o k response has a 5xx status code
func (o *ConsistencyGroupMetricsCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this consistency group metrics collection get o k response a status code equal to that given
func (o *ConsistencyGroupMetricsCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the consistency group metrics collection get o k response
func (o *ConsistencyGroupMetricsCollectionGetOK) Code() int {
	return 200
}

func (o *ConsistencyGroupMetricsCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/consistency-groups/{consistency_group.uuid}/metrics][%d] consistencyGroupMetricsCollectionGetOK %s", 200, payload)
}

func (o *ConsistencyGroupMetricsCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/consistency-groups/{consistency_group.uuid}/metrics][%d] consistencyGroupMetricsCollectionGetOK %s", 200, payload)
}

func (o *ConsistencyGroupMetricsCollectionGetOK) GetPayload() *models.ConsistencyGroupMetricsResponse {
	return o.Payload
}

func (o *ConsistencyGroupMetricsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsistencyGroupMetricsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsistencyGroupMetricsCollectionGetDefault creates a ConsistencyGroupMetricsCollectionGetDefault with default headers values
func NewConsistencyGroupMetricsCollectionGetDefault(code int) *ConsistencyGroupMetricsCollectionGetDefault {
	return &ConsistencyGroupMetricsCollectionGetDefault{
		_statusCode: code,
	}
}

/*
ConsistencyGroupMetricsCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type ConsistencyGroupMetricsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this consistency group metrics collection get default response has a 2xx status code
func (o *ConsistencyGroupMetricsCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this consistency group metrics collection get default response has a 3xx status code
func (o *ConsistencyGroupMetricsCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this consistency group metrics collection get default response has a 4xx status code
func (o *ConsistencyGroupMetricsCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this consistency group metrics collection get default response has a 5xx status code
func (o *ConsistencyGroupMetricsCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this consistency group metrics collection get default response a status code equal to that given
func (o *ConsistencyGroupMetricsCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the consistency group metrics collection get default response
func (o *ConsistencyGroupMetricsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *ConsistencyGroupMetricsCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/consistency-groups/{consistency_group.uuid}/metrics][%d] consistency_group_metrics_collection_get default %s", o._statusCode, payload)
}

func (o *ConsistencyGroupMetricsCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/consistency-groups/{consistency_group.uuid}/metrics][%d] consistency_group_metrics_collection_get default %s", o._statusCode, payload)
}

func (o *ConsistencyGroupMetricsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConsistencyGroupMetricsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
