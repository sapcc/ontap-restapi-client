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

	"github.com/sapcc/ontap-restapi/models"
)

// MetroclusterCreateReader is a Reader for the MetroclusterCreate structure.
type MetroclusterCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroclusterCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewMetroclusterCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewMetroclusterCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMetroclusterCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMetroclusterCreateCreated creates a MetroclusterCreateCreated with default headers values
func NewMetroclusterCreateCreated() *MetroclusterCreateCreated {
	return &MetroclusterCreateCreated{}
}

/*
MetroclusterCreateCreated describes a response with status code 201, with default header values.

Created
*/
type MetroclusterCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this metrocluster create created response has a 2xx status code
func (o *MetroclusterCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this metrocluster create created response has a 3xx status code
func (o *MetroclusterCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this metrocluster create created response has a 4xx status code
func (o *MetroclusterCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this metrocluster create created response has a 5xx status code
func (o *MetroclusterCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this metrocluster create created response a status code equal to that given
func (o *MetroclusterCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the metrocluster create created response
func (o *MetroclusterCreateCreated) Code() int {
	return 201
}

func (o *MetroclusterCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/metrocluster][%d] metroclusterCreateCreated %s", 201, payload)
}

func (o *MetroclusterCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/metrocluster][%d] metroclusterCreateCreated %s", 201, payload)
}

func (o *MetroclusterCreateCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *MetroclusterCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroclusterCreateAccepted creates a MetroclusterCreateAccepted with default headers values
func NewMetroclusterCreateAccepted() *MetroclusterCreateAccepted {
	return &MetroclusterCreateAccepted{}
}

/*
MetroclusterCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type MetroclusterCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this metrocluster create accepted response has a 2xx status code
func (o *MetroclusterCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this metrocluster create accepted response has a 3xx status code
func (o *MetroclusterCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this metrocluster create accepted response has a 4xx status code
func (o *MetroclusterCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this metrocluster create accepted response has a 5xx status code
func (o *MetroclusterCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this metrocluster create accepted response a status code equal to that given
func (o *MetroclusterCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the metrocluster create accepted response
func (o *MetroclusterCreateAccepted) Code() int {
	return 202
}

func (o *MetroclusterCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/metrocluster][%d] metroclusterCreateAccepted %s", 202, payload)
}

func (o *MetroclusterCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/metrocluster][%d] metroclusterCreateAccepted %s", 202, payload)
}

func (o *MetroclusterCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *MetroclusterCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroclusterCreateDefault creates a MetroclusterCreateDefault with default headers values
func NewMetroclusterCreateDefault(code int) *MetroclusterCreateDefault {
	return &MetroclusterCreateDefault{
		_statusCode: code,
	}
}

/*
	MetroclusterCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2425734 | An internal error occurred. Wait a few minutes, and try the operation again. For further assistance, contact technical support. |
| 2432832 | Required environment variables are not set. |
| 2432833 | Operation is already running. |
| 2432834 | MetroCluster is already configured. |
| 2432835 | Operation not supported. |
| 2432836 | There are not enough disks in Pool1. Wait a few minutes, and try the operation again. For further assistance, contact technical support. |
| 2432839 | Required parameters not set. |
| 2432840 | Configuring DR Groups |
| 2432841 | Generating IP addresses |
| 2432843 | Running Aggregate Recommender |
| 2432844 | Checking remote storage pool |
| 2432845 | Mirroring aggregates |
| 2432846 | Configuring MetroCluster and DR mirroring |
| 2432848 | Setting up MetroCluster |
| 2432849 | MetroCluster setup is complete |
| 2432851 | Minimum number of required data aggregates for MetroCluster configuration are still not mirrored. Wait a few minutes, and try the operation again. For further assistance, contact technical support. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type MetroclusterCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this metrocluster create default response has a 2xx status code
func (o *MetroclusterCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this metrocluster create default response has a 3xx status code
func (o *MetroclusterCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this metrocluster create default response has a 4xx status code
func (o *MetroclusterCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this metrocluster create default response has a 5xx status code
func (o *MetroclusterCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this metrocluster create default response a status code equal to that given
func (o *MetroclusterCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the metrocluster create default response
func (o *MetroclusterCreateDefault) Code() int {
	return o._statusCode
}

func (o *MetroclusterCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/metrocluster][%d] metrocluster_create default %s", o._statusCode, payload)
}

func (o *MetroclusterCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/metrocluster][%d] metrocluster_create default %s", o._statusCode, payload)
}

func (o *MetroclusterCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MetroclusterCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
