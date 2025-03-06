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

// PerformanceFcPortMetricCollectionGetReader is a Reader for the PerformanceFcPortMetricCollectionGet structure.
type PerformanceFcPortMetricCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformanceFcPortMetricCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformanceFcPortMetricCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPerformanceFcPortMetricCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformanceFcPortMetricCollectionGetOK creates a PerformanceFcPortMetricCollectionGetOK with default headers values
func NewPerformanceFcPortMetricCollectionGetOK() *PerformanceFcPortMetricCollectionGetOK {
	return &PerformanceFcPortMetricCollectionGetOK{}
}

/*
PerformanceFcPortMetricCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type PerformanceFcPortMetricCollectionGetOK struct {
	Payload *models.PerformanceFcPortMetricResponse
}

// IsSuccess returns true when this performance fc port metric collection get o k response has a 2xx status code
func (o *PerformanceFcPortMetricCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performance fc port metric collection get o k response has a 3xx status code
func (o *PerformanceFcPortMetricCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performance fc port metric collection get o k response has a 4xx status code
func (o *PerformanceFcPortMetricCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performance fc port metric collection get o k response has a 5xx status code
func (o *PerformanceFcPortMetricCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performance fc port metric collection get o k response a status code equal to that given
func (o *PerformanceFcPortMetricCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the performance fc port metric collection get o k response
func (o *PerformanceFcPortMetricCollectionGetOK) Code() int {
	return 200
}

func (o *PerformanceFcPortMetricCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/ports/{fc_port.uuid}/metrics][%d] performanceFcPortMetricCollectionGetOK %s", 200, payload)
}

func (o *PerformanceFcPortMetricCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/ports/{fc_port.uuid}/metrics][%d] performanceFcPortMetricCollectionGetOK %s", 200, payload)
}

func (o *PerformanceFcPortMetricCollectionGetOK) GetPayload() *models.PerformanceFcPortMetricResponse {
	return o.Payload
}

func (o *PerformanceFcPortMetricCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceFcPortMetricResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformanceFcPortMetricCollectionGetDefault creates a PerformanceFcPortMetricCollectionGetDefault with default headers values
func NewPerformanceFcPortMetricCollectionGetDefault(code int) *PerformanceFcPortMetricCollectionGetDefault {
	return &PerformanceFcPortMetricCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	PerformanceFcPortMetricCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8585947 | No metrics are available for the requested object. |
| 8586225 | An unexpected error occurred retrieving metrics for the requested object. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type PerformanceFcPortMetricCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this performance fc port metric collection get default response has a 2xx status code
func (o *PerformanceFcPortMetricCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this performance fc port metric collection get default response has a 3xx status code
func (o *PerformanceFcPortMetricCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this performance fc port metric collection get default response has a 4xx status code
func (o *PerformanceFcPortMetricCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this performance fc port metric collection get default response has a 5xx status code
func (o *PerformanceFcPortMetricCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this performance fc port metric collection get default response a status code equal to that given
func (o *PerformanceFcPortMetricCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the performance fc port metric collection get default response
func (o *PerformanceFcPortMetricCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *PerformanceFcPortMetricCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/ports/{fc_port.uuid}/metrics][%d] performance_fc_port_metric_collection_get default %s", o._statusCode, payload)
}

func (o *PerformanceFcPortMetricCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/ports/{fc_port.uuid}/metrics][%d] performance_fc_port_metric_collection_get default %s", o._statusCode, payload)
}

func (o *PerformanceFcPortMetricCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PerformanceFcPortMetricCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
