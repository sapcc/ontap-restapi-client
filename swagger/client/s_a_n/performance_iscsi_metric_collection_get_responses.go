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

// PerformanceIscsiMetricCollectionGetReader is a Reader for the PerformanceIscsiMetricCollectionGet structure.
type PerformanceIscsiMetricCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformanceIscsiMetricCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformanceIscsiMetricCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPerformanceIscsiMetricCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformanceIscsiMetricCollectionGetOK creates a PerformanceIscsiMetricCollectionGetOK with default headers values
func NewPerformanceIscsiMetricCollectionGetOK() *PerformanceIscsiMetricCollectionGetOK {
	return &PerformanceIscsiMetricCollectionGetOK{}
}

/*
PerformanceIscsiMetricCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type PerformanceIscsiMetricCollectionGetOK struct {
	Payload *models.PerformanceIscsiMetricResponse
}

// IsSuccess returns true when this performance iscsi metric collection get o k response has a 2xx status code
func (o *PerformanceIscsiMetricCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performance iscsi metric collection get o k response has a 3xx status code
func (o *PerformanceIscsiMetricCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performance iscsi metric collection get o k response has a 4xx status code
func (o *PerformanceIscsiMetricCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performance iscsi metric collection get o k response has a 5xx status code
func (o *PerformanceIscsiMetricCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performance iscsi metric collection get o k response a status code equal to that given
func (o *PerformanceIscsiMetricCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the performance iscsi metric collection get o k response
func (o *PerformanceIscsiMetricCollectionGetOK) Code() int {
	return 200
}

func (o *PerformanceIscsiMetricCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/services/{svm.uuid}/metrics][%d] performanceIscsiMetricCollectionGetOK %s", 200, payload)
}

func (o *PerformanceIscsiMetricCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/services/{svm.uuid}/metrics][%d] performanceIscsiMetricCollectionGetOK %s", 200, payload)
}

func (o *PerformanceIscsiMetricCollectionGetOK) GetPayload() *models.PerformanceIscsiMetricResponse {
	return o.Payload
}

func (o *PerformanceIscsiMetricCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceIscsiMetricResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformanceIscsiMetricCollectionGetDefault creates a PerformanceIscsiMetricCollectionGetDefault with default headers values
func NewPerformanceIscsiMetricCollectionGetDefault(code int) *PerformanceIscsiMetricCollectionGetDefault {
	return &PerformanceIscsiMetricCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	PerformanceIscsiMetricCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8585947 | No metrics are available for the requested object. |
| 8586225 | An unexpected error occurred retrieving metrics for the requested object. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type PerformanceIscsiMetricCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this performance iscsi metric collection get default response has a 2xx status code
func (o *PerformanceIscsiMetricCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this performance iscsi metric collection get default response has a 3xx status code
func (o *PerformanceIscsiMetricCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this performance iscsi metric collection get default response has a 4xx status code
func (o *PerformanceIscsiMetricCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this performance iscsi metric collection get default response has a 5xx status code
func (o *PerformanceIscsiMetricCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this performance iscsi metric collection get default response a status code equal to that given
func (o *PerformanceIscsiMetricCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the performance iscsi metric collection get default response
func (o *PerformanceIscsiMetricCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *PerformanceIscsiMetricCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/services/{svm.uuid}/metrics][%d] performance_iscsi_metric_collection_get default %s", o._statusCode, payload)
}

func (o *PerformanceIscsiMetricCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/services/{svm.uuid}/metrics][%d] performance_iscsi_metric_collection_get default %s", o._statusCode, payload)
}

func (o *PerformanceIscsiMetricCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PerformanceIscsiMetricCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
