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

	"github.com/sapcc/ontap-restapi-client/models"
)

// NodeMetricsCollectionGetReader is a Reader for the NodeMetricsCollectionGet structure.
type NodeMetricsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeMetricsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeMetricsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNodeMetricsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodeMetricsCollectionGetOK creates a NodeMetricsCollectionGetOK with default headers values
func NewNodeMetricsCollectionGetOK() *NodeMetricsCollectionGetOK {
	return &NodeMetricsCollectionGetOK{}
}

/*
NodeMetricsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type NodeMetricsCollectionGetOK struct {
	Payload *models.NodeMetricsResponse
}

// IsSuccess returns true when this node metrics collection get o k response has a 2xx status code
func (o *NodeMetricsCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this node metrics collection get o k response has a 3xx status code
func (o *NodeMetricsCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node metrics collection get o k response has a 4xx status code
func (o *NodeMetricsCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this node metrics collection get o k response has a 5xx status code
func (o *NodeMetricsCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this node metrics collection get o k response a status code equal to that given
func (o *NodeMetricsCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the node metrics collection get o k response
func (o *NodeMetricsCollectionGetOK) Code() int {
	return 200
}

func (o *NodeMetricsCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/nodes/{uuid}/metrics][%d] nodeMetricsCollectionGetOK %s", 200, payload)
}

func (o *NodeMetricsCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/nodes/{uuid}/metrics][%d] nodeMetricsCollectionGetOK %s", 200, payload)
}

func (o *NodeMetricsCollectionGetOK) GetPayload() *models.NodeMetricsResponse {
	return o.Payload
}

func (o *NodeMetricsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeMetricsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeMetricsCollectionGetDefault creates a NodeMetricsCollectionGetDefault with default headers values
func NewNodeMetricsCollectionGetDefault(code int) *NodeMetricsCollectionGetDefault {
	return &NodeMetricsCollectionGetDefault{
		_statusCode: code,
	}
}

/*
NodeMetricsCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type NodeMetricsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this node metrics collection get default response has a 2xx status code
func (o *NodeMetricsCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this node metrics collection get default response has a 3xx status code
func (o *NodeMetricsCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this node metrics collection get default response has a 4xx status code
func (o *NodeMetricsCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this node metrics collection get default response has a 5xx status code
func (o *NodeMetricsCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this node metrics collection get default response a status code equal to that given
func (o *NodeMetricsCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the node metrics collection get default response
func (o *NodeMetricsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *NodeMetricsCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/nodes/{uuid}/metrics][%d] node_metrics_collection_get default %s", o._statusCode, payload)
}

func (o *NodeMetricsCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/nodes/{uuid}/metrics][%d] node_metrics_collection_get default %s", o._statusCode, payload)
}

func (o *NodeMetricsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeMetricsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
