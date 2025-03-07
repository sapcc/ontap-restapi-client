// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// MetroclusterDiagnosticsGetReader is a Reader for the MetroclusterDiagnosticsGet structure.
type MetroclusterDiagnosticsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroclusterDiagnosticsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroclusterDiagnosticsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMetroclusterDiagnosticsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMetroclusterDiagnosticsGetOK creates a MetroclusterDiagnosticsGetOK with default headers values
func NewMetroclusterDiagnosticsGetOK() *MetroclusterDiagnosticsGetOK {
	return &MetroclusterDiagnosticsGetOK{}
}

/*
MetroclusterDiagnosticsGetOK describes a response with status code 200, with default header values.

OK
*/
type MetroclusterDiagnosticsGetOK struct {
	Payload *models.MetroclusterDiagnostics
}

// IsSuccess returns true when this metrocluster diagnostics get o k response has a 2xx status code
func (o *MetroclusterDiagnosticsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this metrocluster diagnostics get o k response has a 3xx status code
func (o *MetroclusterDiagnosticsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this metrocluster diagnostics get o k response has a 4xx status code
func (o *MetroclusterDiagnosticsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this metrocluster diagnostics get o k response has a 5xx status code
func (o *MetroclusterDiagnosticsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this metrocluster diagnostics get o k response a status code equal to that given
func (o *MetroclusterDiagnosticsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the metrocluster diagnostics get o k response
func (o *MetroclusterDiagnosticsGetOK) Code() int {
	return 200
}

func (o *MetroclusterDiagnosticsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster/diagnostics][%d] metroclusterDiagnosticsGetOK %s", 200, payload)
}

func (o *MetroclusterDiagnosticsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster/diagnostics][%d] metroclusterDiagnosticsGetOK %s", 200, payload)
}

func (o *MetroclusterDiagnosticsGetOK) GetPayload() *models.MetroclusterDiagnostics {
	return o.Payload
}

func (o *MetroclusterDiagnosticsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetroclusterDiagnostics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroclusterDiagnosticsGetDefault creates a MetroclusterDiagnosticsGetDefault with default headers values
func NewMetroclusterDiagnosticsGetDefault(code int) *MetroclusterDiagnosticsGetDefault {
	return &MetroclusterDiagnosticsGetDefault{
		_statusCode: code,
	}
}

/*
	MetroclusterDiagnosticsGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2425734 | An internal error occurred. Wait a few minutes, and try the operation again. For further assistance, contact technical support. |
| 2426405 | MetroCluster is not configured in cluster "cluster". Run "metrocluster show" to verify both the local and the remote clusters are configured as MetroCluster, and then try the command again. |
| 2432856 | MetroCluster diagnostics result is not available. Use the REST API GET method on "/api/cluster/metrocluster/operations?type=check&fields=*" for more information. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type MetroclusterDiagnosticsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this metrocluster diagnostics get default response has a 2xx status code
func (o *MetroclusterDiagnosticsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this metrocluster diagnostics get default response has a 3xx status code
func (o *MetroclusterDiagnosticsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this metrocluster diagnostics get default response has a 4xx status code
func (o *MetroclusterDiagnosticsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this metrocluster diagnostics get default response has a 5xx status code
func (o *MetroclusterDiagnosticsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this metrocluster diagnostics get default response a status code equal to that given
func (o *MetroclusterDiagnosticsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the metrocluster diagnostics get default response
func (o *MetroclusterDiagnosticsGetDefault) Code() int {
	return o._statusCode
}

func (o *MetroclusterDiagnosticsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster/diagnostics][%d] metrocluster_diagnostics_get default %s", o._statusCode, payload)
}

func (o *MetroclusterDiagnosticsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster/diagnostics][%d] metrocluster_diagnostics_get default %s", o._statusCode, payload)
}

func (o *MetroclusterDiagnosticsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MetroclusterDiagnosticsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
