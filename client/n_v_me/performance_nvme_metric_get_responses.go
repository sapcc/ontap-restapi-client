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

	"github.com/sapcc/ontap-restapi-client/models"
)

// PerformanceNvmeMetricGetReader is a Reader for the PerformanceNvmeMetricGet structure.
type PerformanceNvmeMetricGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformanceNvmeMetricGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformanceNvmeMetricGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPerformanceNvmeMetricGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformanceNvmeMetricGetOK creates a PerformanceNvmeMetricGetOK with default headers values
func NewPerformanceNvmeMetricGetOK() *PerformanceNvmeMetricGetOK {
	return &PerformanceNvmeMetricGetOK{}
}

/*
PerformanceNvmeMetricGetOK describes a response with status code 200, with default header values.

OK
*/
type PerformanceNvmeMetricGetOK struct {
	Payload *models.PerformanceNvmeMetric
}

// IsSuccess returns true when this performance nvme metric get o k response has a 2xx status code
func (o *PerformanceNvmeMetricGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performance nvme metric get o k response has a 3xx status code
func (o *PerformanceNvmeMetricGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performance nvme metric get o k response has a 4xx status code
func (o *PerformanceNvmeMetricGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performance nvme metric get o k response has a 5xx status code
func (o *PerformanceNvmeMetricGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performance nvme metric get o k response a status code equal to that given
func (o *PerformanceNvmeMetricGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the performance nvme metric get o k response
func (o *PerformanceNvmeMetricGetOK) Code() int {
	return 200
}

func (o *PerformanceNvmeMetricGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/services/{svm.uuid}/metrics/{timestamp}][%d] performanceNvmeMetricGetOK %s", 200, payload)
}

func (o *PerformanceNvmeMetricGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/services/{svm.uuid}/metrics/{timestamp}][%d] performanceNvmeMetricGetOK %s", 200, payload)
}

func (o *PerformanceNvmeMetricGetOK) GetPayload() *models.PerformanceNvmeMetric {
	return o.Payload
}

func (o *PerformanceNvmeMetricGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceNvmeMetric)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformanceNvmeMetricGetDefault creates a PerformanceNvmeMetricGetDefault with default headers values
func NewPerformanceNvmeMetricGetDefault(code int) *PerformanceNvmeMetricGetDefault {
	return &PerformanceNvmeMetricGetDefault{
		_statusCode: code,
	}
}

/*
	PerformanceNvmeMetricGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8585947 | No metrics are available for the requested object. |
| 8586225 | An unexpected error occurred retrieving metrics for the requested object. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type PerformanceNvmeMetricGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this performance nvme metric get default response has a 2xx status code
func (o *PerformanceNvmeMetricGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this performance nvme metric get default response has a 3xx status code
func (o *PerformanceNvmeMetricGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this performance nvme metric get default response has a 4xx status code
func (o *PerformanceNvmeMetricGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this performance nvme metric get default response has a 5xx status code
func (o *PerformanceNvmeMetricGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this performance nvme metric get default response a status code equal to that given
func (o *PerformanceNvmeMetricGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the performance nvme metric get default response
func (o *PerformanceNvmeMetricGetDefault) Code() int {
	return o._statusCode
}

func (o *PerformanceNvmeMetricGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/services/{svm.uuid}/metrics/{timestamp}][%d] performance_nvme_metric_get default %s", o._statusCode, payload)
}

func (o *PerformanceNvmeMetricGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/services/{svm.uuid}/metrics/{timestamp}][%d] performance_nvme_metric_get default %s", o._statusCode, payload)
}

func (o *PerformanceNvmeMetricGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PerformanceNvmeMetricGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
