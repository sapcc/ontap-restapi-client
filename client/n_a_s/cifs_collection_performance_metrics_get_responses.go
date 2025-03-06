// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// CifsCollectionPerformanceMetricsGetReader is a Reader for the CifsCollectionPerformanceMetricsGet structure.
type CifsCollectionPerformanceMetricsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsCollectionPerformanceMetricsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsCollectionPerformanceMetricsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsCollectionPerformanceMetricsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsCollectionPerformanceMetricsGetOK creates a CifsCollectionPerformanceMetricsGetOK with default headers values
func NewCifsCollectionPerformanceMetricsGetOK() *CifsCollectionPerformanceMetricsGetOK {
	return &CifsCollectionPerformanceMetricsGetOK{}
}

/*
CifsCollectionPerformanceMetricsGetOK describes a response with status code 200, with default header values.

OK
*/
type CifsCollectionPerformanceMetricsGetOK struct {
	Payload *models.PerformanceCifsMetricResponse
}

// IsSuccess returns true when this cifs collection performance metrics get o k response has a 2xx status code
func (o *CifsCollectionPerformanceMetricsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs collection performance metrics get o k response has a 3xx status code
func (o *CifsCollectionPerformanceMetricsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs collection performance metrics get o k response has a 4xx status code
func (o *CifsCollectionPerformanceMetricsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs collection performance metrics get o k response has a 5xx status code
func (o *CifsCollectionPerformanceMetricsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs collection performance metrics get o k response a status code equal to that given
func (o *CifsCollectionPerformanceMetricsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs collection performance metrics get o k response
func (o *CifsCollectionPerformanceMetricsGetOK) Code() int {
	return 200
}

func (o *CifsCollectionPerformanceMetricsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/services/{svm.uuid}/metrics][%d] cifsCollectionPerformanceMetricsGetOK %s", 200, payload)
}

func (o *CifsCollectionPerformanceMetricsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/services/{svm.uuid}/metrics][%d] cifsCollectionPerformanceMetricsGetOK %s", 200, payload)
}

func (o *CifsCollectionPerformanceMetricsGetOK) GetPayload() *models.PerformanceCifsMetricResponse {
	return o.Payload
}

func (o *CifsCollectionPerformanceMetricsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceCifsMetricResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsCollectionPerformanceMetricsGetDefault creates a CifsCollectionPerformanceMetricsGetDefault with default headers values
func NewCifsCollectionPerformanceMetricsGetDefault(code int) *CifsCollectionPerformanceMetricsGetDefault {
	return &CifsCollectionPerformanceMetricsGetDefault{
		_statusCode: code,
	}
}

/*
CifsCollectionPerformanceMetricsGetDefault describes a response with status code -1, with default header values.

Error
*/
type CifsCollectionPerformanceMetricsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs collection performance metrics get default response has a 2xx status code
func (o *CifsCollectionPerformanceMetricsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs collection performance metrics get default response has a 3xx status code
func (o *CifsCollectionPerformanceMetricsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs collection performance metrics get default response has a 4xx status code
func (o *CifsCollectionPerformanceMetricsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs collection performance metrics get default response has a 5xx status code
func (o *CifsCollectionPerformanceMetricsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs collection performance metrics get default response a status code equal to that given
func (o *CifsCollectionPerformanceMetricsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs collection performance metrics get default response
func (o *CifsCollectionPerformanceMetricsGetDefault) Code() int {
	return o._statusCode
}

func (o *CifsCollectionPerformanceMetricsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/services/{svm.uuid}/metrics][%d] cifs_collection_performance_metrics_get default %s", o._statusCode, payload)
}

func (o *CifsCollectionPerformanceMetricsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/services/{svm.uuid}/metrics][%d] cifs_collection_performance_metrics_get default %s", o._statusCode, payload)
}

func (o *CifsCollectionPerformanceMetricsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsCollectionPerformanceMetricsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
