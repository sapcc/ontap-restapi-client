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

// ClusterCollectionPerformanceMetricsGetReader is a Reader for the ClusterCollectionPerformanceMetricsGet structure.
type ClusterCollectionPerformanceMetricsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterCollectionPerformanceMetricsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterCollectionPerformanceMetricsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterCollectionPerformanceMetricsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterCollectionPerformanceMetricsGetOK creates a ClusterCollectionPerformanceMetricsGetOK with default headers values
func NewClusterCollectionPerformanceMetricsGetOK() *ClusterCollectionPerformanceMetricsGetOK {
	return &ClusterCollectionPerformanceMetricsGetOK{}
}

/*
ClusterCollectionPerformanceMetricsGetOK describes a response with status code 200, with default header values.

OK
*/
type ClusterCollectionPerformanceMetricsGetOK struct {
	Payload *models.ClusterMetricsResponse
}

// IsSuccess returns true when this cluster collection performance metrics get o k response has a 2xx status code
func (o *ClusterCollectionPerformanceMetricsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster collection performance metrics get o k response has a 3xx status code
func (o *ClusterCollectionPerformanceMetricsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster collection performance metrics get o k response has a 4xx status code
func (o *ClusterCollectionPerformanceMetricsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster collection performance metrics get o k response has a 5xx status code
func (o *ClusterCollectionPerformanceMetricsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster collection performance metrics get o k response a status code equal to that given
func (o *ClusterCollectionPerformanceMetricsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster collection performance metrics get o k response
func (o *ClusterCollectionPerformanceMetricsGetOK) Code() int {
	return 200
}

func (o *ClusterCollectionPerformanceMetricsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrics][%d] clusterCollectionPerformanceMetricsGetOK %s", 200, payload)
}

func (o *ClusterCollectionPerformanceMetricsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrics][%d] clusterCollectionPerformanceMetricsGetOK %s", 200, payload)
}

func (o *ClusterCollectionPerformanceMetricsGetOK) GetPayload() *models.ClusterMetricsResponse {
	return o.Payload
}

func (o *ClusterCollectionPerformanceMetricsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterMetricsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterCollectionPerformanceMetricsGetDefault creates a ClusterCollectionPerformanceMetricsGetDefault with default headers values
func NewClusterCollectionPerformanceMetricsGetDefault(code int) *ClusterCollectionPerformanceMetricsGetDefault {
	return &ClusterCollectionPerformanceMetricsGetDefault{
		_statusCode: code,
	}
}

/*
ClusterCollectionPerformanceMetricsGetDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterCollectionPerformanceMetricsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster collection performance metrics get default response has a 2xx status code
func (o *ClusterCollectionPerformanceMetricsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster collection performance metrics get default response has a 3xx status code
func (o *ClusterCollectionPerformanceMetricsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster collection performance metrics get default response has a 4xx status code
func (o *ClusterCollectionPerformanceMetricsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster collection performance metrics get default response has a 5xx status code
func (o *ClusterCollectionPerformanceMetricsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster collection performance metrics get default response a status code equal to that given
func (o *ClusterCollectionPerformanceMetricsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster collection performance metrics get default response
func (o *ClusterCollectionPerformanceMetricsGetDefault) Code() int {
	return o._statusCode
}

func (o *ClusterCollectionPerformanceMetricsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrics][%d] cluster_collection_performance_metrics_get default %s", o._statusCode, payload)
}

func (o *ClusterCollectionPerformanceMetricsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrics][%d] cluster_collection_performance_metrics_get default %s", o._statusCode, payload)
}

func (o *ClusterCollectionPerformanceMetricsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterCollectionPerformanceMetricsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
