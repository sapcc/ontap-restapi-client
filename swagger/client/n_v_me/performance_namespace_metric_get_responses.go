// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

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

// PerformanceNamespaceMetricGetReader is a Reader for the PerformanceNamespaceMetricGet structure.
type PerformanceNamespaceMetricGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformanceNamespaceMetricGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformanceNamespaceMetricGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPerformanceNamespaceMetricGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformanceNamespaceMetricGetOK creates a PerformanceNamespaceMetricGetOK with default headers values
func NewPerformanceNamespaceMetricGetOK() *PerformanceNamespaceMetricGetOK {
	return &PerformanceNamespaceMetricGetOK{}
}

/*
PerformanceNamespaceMetricGetOK describes a response with status code 200, with default header values.

OK
*/
type PerformanceNamespaceMetricGetOK struct {
	Payload *models.PerformanceNamespaceMetric
}

// IsSuccess returns true when this performance namespace metric get o k response has a 2xx status code
func (o *PerformanceNamespaceMetricGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performance namespace metric get o k response has a 3xx status code
func (o *PerformanceNamespaceMetricGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performance namespace metric get o k response has a 4xx status code
func (o *PerformanceNamespaceMetricGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performance namespace metric get o k response has a 5xx status code
func (o *PerformanceNamespaceMetricGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performance namespace metric get o k response a status code equal to that given
func (o *PerformanceNamespaceMetricGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the performance namespace metric get o k response
func (o *PerformanceNamespaceMetricGetOK) Code() int {
	return 200
}

func (o *PerformanceNamespaceMetricGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/namespaces/{nvme_namespace.uuid}/metrics/{timestamp}][%d] performanceNamespaceMetricGetOK %s", 200, payload)
}

func (o *PerformanceNamespaceMetricGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/namespaces/{nvme_namespace.uuid}/metrics/{timestamp}][%d] performanceNamespaceMetricGetOK %s", 200, payload)
}

func (o *PerformanceNamespaceMetricGetOK) GetPayload() *models.PerformanceNamespaceMetric {
	return o.Payload
}

func (o *PerformanceNamespaceMetricGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceNamespaceMetric)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformanceNamespaceMetricGetDefault creates a PerformanceNamespaceMetricGetDefault with default headers values
func NewPerformanceNamespaceMetricGetDefault(code int) *PerformanceNamespaceMetricGetDefault {
	return &PerformanceNamespaceMetricGetDefault{
		_statusCode: code,
	}
}

/*
	PerformanceNamespaceMetricGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8585947 | No metrics are available for the requested object. |
| 8586225 | An unexpected error occurred retrieving metrics for the requested object. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type PerformanceNamespaceMetricGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this performance namespace metric get default response has a 2xx status code
func (o *PerformanceNamespaceMetricGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this performance namespace metric get default response has a 3xx status code
func (o *PerformanceNamespaceMetricGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this performance namespace metric get default response has a 4xx status code
func (o *PerformanceNamespaceMetricGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this performance namespace metric get default response has a 5xx status code
func (o *PerformanceNamespaceMetricGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this performance namespace metric get default response a status code equal to that given
func (o *PerformanceNamespaceMetricGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the performance namespace metric get default response
func (o *PerformanceNamespaceMetricGetDefault) Code() int {
	return o._statusCode
}

func (o *PerformanceNamespaceMetricGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/namespaces/{nvme_namespace.uuid}/metrics/{timestamp}][%d] performance_namespace_metric_get default %s", o._statusCode, payload)
}

func (o *PerformanceNamespaceMetricGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/namespaces/{nvme_namespace.uuid}/metrics/{timestamp}][%d] performance_namespace_metric_get default %s", o._statusCode, payload)
}

func (o *PerformanceNamespaceMetricGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PerformanceNamespaceMetricGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
