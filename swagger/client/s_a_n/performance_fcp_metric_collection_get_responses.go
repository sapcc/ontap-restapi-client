// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// PerformanceFcpMetricCollectionGetReader is a Reader for the PerformanceFcpMetricCollectionGet structure.
type PerformanceFcpMetricCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformanceFcpMetricCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformanceFcpMetricCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPerformanceFcpMetricCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformanceFcpMetricCollectionGetOK creates a PerformanceFcpMetricCollectionGetOK with default headers values
func NewPerformanceFcpMetricCollectionGetOK() *PerformanceFcpMetricCollectionGetOK {
	return &PerformanceFcpMetricCollectionGetOK{}
}

/*
PerformanceFcpMetricCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type PerformanceFcpMetricCollectionGetOK struct {
	Payload *models.PerformanceFcpMetricResponse
}

// IsSuccess returns true when this performance fcp metric collection get o k response has a 2xx status code
func (o *PerformanceFcpMetricCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performance fcp metric collection get o k response has a 3xx status code
func (o *PerformanceFcpMetricCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performance fcp metric collection get o k response has a 4xx status code
func (o *PerformanceFcpMetricCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performance fcp metric collection get o k response has a 5xx status code
func (o *PerformanceFcpMetricCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performance fcp metric collection get o k response a status code equal to that given
func (o *PerformanceFcpMetricCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the performance fcp metric collection get o k response
func (o *PerformanceFcpMetricCollectionGetOK) Code() int {
	return 200
}

func (o *PerformanceFcpMetricCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/fcp/services/{svm.uuid}/metrics][%d] performanceFcpMetricCollectionGetOK %s", 200, payload)
}

func (o *PerformanceFcpMetricCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/fcp/services/{svm.uuid}/metrics][%d] performanceFcpMetricCollectionGetOK %s", 200, payload)
}

func (o *PerformanceFcpMetricCollectionGetOK) GetPayload() *models.PerformanceFcpMetricResponse {
	return o.Payload
}

func (o *PerformanceFcpMetricCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceFcpMetricResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformanceFcpMetricCollectionGetDefault creates a PerformanceFcpMetricCollectionGetDefault with default headers values
func NewPerformanceFcpMetricCollectionGetDefault(code int) *PerformanceFcpMetricCollectionGetDefault {
	return &PerformanceFcpMetricCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	PerformanceFcpMetricCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8585947 | No metrics are available for the requested object. |
| 8586225 | An unexpected error occurred retrieving metrics for the requested object. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type PerformanceFcpMetricCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this performance fcp metric collection get default response has a 2xx status code
func (o *PerformanceFcpMetricCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this performance fcp metric collection get default response has a 3xx status code
func (o *PerformanceFcpMetricCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this performance fcp metric collection get default response has a 4xx status code
func (o *PerformanceFcpMetricCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this performance fcp metric collection get default response has a 5xx status code
func (o *PerformanceFcpMetricCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this performance fcp metric collection get default response a status code equal to that given
func (o *PerformanceFcpMetricCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the performance fcp metric collection get default response
func (o *PerformanceFcpMetricCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *PerformanceFcpMetricCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/fcp/services/{svm.uuid}/metrics][%d] performance_fcp_metric_collection_get default %s", o._statusCode, payload)
}

func (o *PerformanceFcpMetricCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/fcp/services/{svm.uuid}/metrics][%d] performance_fcp_metric_collection_get default %s", o._statusCode, payload)
}

func (o *PerformanceFcpMetricCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PerformanceFcpMetricCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
