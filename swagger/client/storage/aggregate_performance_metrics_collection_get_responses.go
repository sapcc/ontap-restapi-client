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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// AggregatePerformanceMetricsCollectionGetReader is a Reader for the AggregatePerformanceMetricsCollectionGet structure.
type AggregatePerformanceMetricsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregatePerformanceMetricsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregatePerformanceMetricsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAggregatePerformanceMetricsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAggregatePerformanceMetricsCollectionGetOK creates a AggregatePerformanceMetricsCollectionGetOK with default headers values
func NewAggregatePerformanceMetricsCollectionGetOK() *AggregatePerformanceMetricsCollectionGetOK {
	return &AggregatePerformanceMetricsCollectionGetOK{}
}

/*
AggregatePerformanceMetricsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type AggregatePerformanceMetricsCollectionGetOK struct {
	Payload *models.PerformanceMetricResponse
}

// IsSuccess returns true when this aggregate performance metrics collection get o k response has a 2xx status code
func (o *AggregatePerformanceMetricsCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregate performance metrics collection get o k response has a 3xx status code
func (o *AggregatePerformanceMetricsCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate performance metrics collection get o k response has a 4xx status code
func (o *AggregatePerformanceMetricsCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate performance metrics collection get o k response has a 5xx status code
func (o *AggregatePerformanceMetricsCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate performance metrics collection get o k response a status code equal to that given
func (o *AggregatePerformanceMetricsCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregate performance metrics collection get o k response
func (o *AggregatePerformanceMetricsCollectionGetOK) Code() int {
	return 200
}

func (o *AggregatePerformanceMetricsCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/aggregates/{uuid}/metrics][%d] aggregatePerformanceMetricsCollectionGetOK %s", 200, payload)
}

func (o *AggregatePerformanceMetricsCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/aggregates/{uuid}/metrics][%d] aggregatePerformanceMetricsCollectionGetOK %s", 200, payload)
}

func (o *AggregatePerformanceMetricsCollectionGetOK) GetPayload() *models.PerformanceMetricResponse {
	return o.Payload
}

func (o *AggregatePerformanceMetricsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceMetricResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregatePerformanceMetricsCollectionGetDefault creates a AggregatePerformanceMetricsCollectionGetDefault with default headers values
func NewAggregatePerformanceMetricsCollectionGetDefault(code int) *AggregatePerformanceMetricsCollectionGetDefault {
	return &AggregatePerformanceMetricsCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	AggregatePerformanceMetricsCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8586225 | Encountered unexpected error in retrieving metrics for the requested aggregate. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AggregatePerformanceMetricsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this aggregate performance metrics collection get default response has a 2xx status code
func (o *AggregatePerformanceMetricsCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this aggregate performance metrics collection get default response has a 3xx status code
func (o *AggregatePerformanceMetricsCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this aggregate performance metrics collection get default response has a 4xx status code
func (o *AggregatePerformanceMetricsCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this aggregate performance metrics collection get default response has a 5xx status code
func (o *AggregatePerformanceMetricsCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this aggregate performance metrics collection get default response a status code equal to that given
func (o *AggregatePerformanceMetricsCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the aggregate performance metrics collection get default response
func (o *AggregatePerformanceMetricsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *AggregatePerformanceMetricsCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/aggregates/{uuid}/metrics][%d] aggregate_performance_metrics_collection_get default %s", o._statusCode, payload)
}

func (o *AggregatePerformanceMetricsCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/aggregates/{uuid}/metrics][%d] aggregate_performance_metrics_collection_get default %s", o._statusCode, payload)
}

func (o *AggregatePerformanceMetricsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AggregatePerformanceMetricsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
