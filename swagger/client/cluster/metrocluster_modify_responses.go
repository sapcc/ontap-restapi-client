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

// MetroclusterModifyReader is a Reader for the MetroclusterModify structure.
type MetroclusterModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroclusterModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroclusterModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewMetroclusterModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMetroclusterModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMetroclusterModifyOK creates a MetroclusterModifyOK with default headers values
func NewMetroclusterModifyOK() *MetroclusterModifyOK {
	return &MetroclusterModifyOK{}
}

/*
MetroclusterModifyOK describes a response with status code 200, with default header values.

OK
*/
type MetroclusterModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this metrocluster modify o k response has a 2xx status code
func (o *MetroclusterModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this metrocluster modify o k response has a 3xx status code
func (o *MetroclusterModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this metrocluster modify o k response has a 4xx status code
func (o *MetroclusterModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this metrocluster modify o k response has a 5xx status code
func (o *MetroclusterModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this metrocluster modify o k response a status code equal to that given
func (o *MetroclusterModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the metrocluster modify o k response
func (o *MetroclusterModifyOK) Code() int {
	return 200
}

func (o *MetroclusterModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/metrocluster][%d] metroclusterModifyOK %s", 200, payload)
}

func (o *MetroclusterModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/metrocluster][%d] metroclusterModifyOK %s", 200, payload)
}

func (o *MetroclusterModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *MetroclusterModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroclusterModifyAccepted creates a MetroclusterModifyAccepted with default headers values
func NewMetroclusterModifyAccepted() *MetroclusterModifyAccepted {
	return &MetroclusterModifyAccepted{}
}

/*
MetroclusterModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type MetroclusterModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this metrocluster modify accepted response has a 2xx status code
func (o *MetroclusterModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this metrocluster modify accepted response has a 3xx status code
func (o *MetroclusterModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this metrocluster modify accepted response has a 4xx status code
func (o *MetroclusterModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this metrocluster modify accepted response has a 5xx status code
func (o *MetroclusterModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this metrocluster modify accepted response a status code equal to that given
func (o *MetroclusterModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the metrocluster modify accepted response
func (o *MetroclusterModifyAccepted) Code() int {
	return 202
}

func (o *MetroclusterModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/metrocluster][%d] metroclusterModifyAccepted %s", 202, payload)
}

func (o *MetroclusterModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/metrocluster][%d] metroclusterModifyAccepted %s", 202, payload)
}

func (o *MetroclusterModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *MetroclusterModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroclusterModifyDefault creates a MetroclusterModifyDefault with default headers values
func NewMetroclusterModifyDefault(code int) *MetroclusterModifyDefault {
	return &MetroclusterModifyDefault{
		_statusCode: code,
	}
}

/*
	MetroclusterModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2424873 | Failed to validate the node and cluster components before the "action" operation. |
| 2425138 | Switchover cannot be performed in the current DR mode. Run "metrocluster show" to view the DR mode of the local cluster, and run "switchover" only in one of the following situations: 1. The DR mode of the local cluster is "normal". 2. The DR mode of the local cluster is "partial-switchover". |
| 2425333 | Heal DR data aggregates cannot be performed in the current DR mode. Run "metrocluster show" and "metrocluster node show" to view the DR mode of the local cluster and the DR mode of the local nodes respectively, and run "heal aggregates" only in one of the following situations: 1. The DR mode of the local cluster is "switchover", and the DR mode of all local nodes is "switchover completed". 2. The DR mode of at least one local node is "heal aggrs failed". |
| 2425335 | Heal DR root aggregates cannot be performed in the current DR mode. Run "metrocluster node show" to view the DR mode of the local nodes, and run "heal root-aggregates" only in one of the following situations: 1. The DR mode of all local nodes is "heal aggrs completed". 2. The DR mode of at least one local node is "heal roots failed". |
| 2425734 | An internal error occurred. Wait a few minutes, and try the operation again. For further assistance, contact technical support. |
| 2427558 | This MetroCluster operation cannot be run because another "action" operation is currently in progress. Run "metrocluster operation history show -job-id id -instance" to view the status of the currently running operation. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type MetroclusterModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this metrocluster modify default response has a 2xx status code
func (o *MetroclusterModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this metrocluster modify default response has a 3xx status code
func (o *MetroclusterModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this metrocluster modify default response has a 4xx status code
func (o *MetroclusterModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this metrocluster modify default response has a 5xx status code
func (o *MetroclusterModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this metrocluster modify default response a status code equal to that given
func (o *MetroclusterModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the metrocluster modify default response
func (o *MetroclusterModifyDefault) Code() int {
	return o._statusCode
}

func (o *MetroclusterModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/metrocluster][%d] metrocluster_modify default %s", o._statusCode, payload)
}

func (o *MetroclusterModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/metrocluster][%d] metrocluster_modify default %s", o._statusCode, payload)
}

func (o *MetroclusterModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MetroclusterModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
