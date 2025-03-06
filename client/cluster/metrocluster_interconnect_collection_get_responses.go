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

// MetroclusterInterconnectCollectionGetReader is a Reader for the MetroclusterInterconnectCollectionGet structure.
type MetroclusterInterconnectCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroclusterInterconnectCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroclusterInterconnectCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMetroclusterInterconnectCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMetroclusterInterconnectCollectionGetOK creates a MetroclusterInterconnectCollectionGetOK with default headers values
func NewMetroclusterInterconnectCollectionGetOK() *MetroclusterInterconnectCollectionGetOK {
	return &MetroclusterInterconnectCollectionGetOK{}
}

/*
MetroclusterInterconnectCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type MetroclusterInterconnectCollectionGetOK struct {
	Payload *models.MetroclusterInterconnectResponse
}

// IsSuccess returns true when this metrocluster interconnect collection get o k response has a 2xx status code
func (o *MetroclusterInterconnectCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this metrocluster interconnect collection get o k response has a 3xx status code
func (o *MetroclusterInterconnectCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this metrocluster interconnect collection get o k response has a 4xx status code
func (o *MetroclusterInterconnectCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this metrocluster interconnect collection get o k response has a 5xx status code
func (o *MetroclusterInterconnectCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this metrocluster interconnect collection get o k response a status code equal to that given
func (o *MetroclusterInterconnectCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the metrocluster interconnect collection get o k response
func (o *MetroclusterInterconnectCollectionGetOK) Code() int {
	return 200
}

func (o *MetroclusterInterconnectCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster/interconnects][%d] metroclusterInterconnectCollectionGetOK %s", 200, payload)
}

func (o *MetroclusterInterconnectCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster/interconnects][%d] metroclusterInterconnectCollectionGetOK %s", 200, payload)
}

func (o *MetroclusterInterconnectCollectionGetOK) GetPayload() *models.MetroclusterInterconnectResponse {
	return o.Payload
}

func (o *MetroclusterInterconnectCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetroclusterInterconnectResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroclusterInterconnectCollectionGetDefault creates a MetroclusterInterconnectCollectionGetDefault with default headers values
func NewMetroclusterInterconnectCollectionGetDefault(code int) *MetroclusterInterconnectCollectionGetDefault {
	return &MetroclusterInterconnectCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	MetroclusterInterconnectCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2424994 | MetroCluster is not supported when HAOSC is not set to "mcc", "mcc-2n" or "mccip". Run the "ha-config modify chassis" and "ha-config modify controller" commands in maintenance mode to change the HAOSC settings on node "node" in cluster "cluster". |
| 2425734 | An internal error occurred. Wait a few minutes, and try the operation again. For further assistance, contact technical support. |
| 2427132 | MetroCluster is not configured on this cluster. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type MetroclusterInterconnectCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this metrocluster interconnect collection get default response has a 2xx status code
func (o *MetroclusterInterconnectCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this metrocluster interconnect collection get default response has a 3xx status code
func (o *MetroclusterInterconnectCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this metrocluster interconnect collection get default response has a 4xx status code
func (o *MetroclusterInterconnectCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this metrocluster interconnect collection get default response has a 5xx status code
func (o *MetroclusterInterconnectCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this metrocluster interconnect collection get default response a status code equal to that given
func (o *MetroclusterInterconnectCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the metrocluster interconnect collection get default response
func (o *MetroclusterInterconnectCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *MetroclusterInterconnectCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster/interconnects][%d] metrocluster_interconnect_collection_get default %s", o._statusCode, payload)
}

func (o *MetroclusterInterconnectCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster/interconnects][%d] metrocluster_interconnect_collection_get default %s", o._statusCode, payload)
}

func (o *MetroclusterInterconnectCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MetroclusterInterconnectCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
