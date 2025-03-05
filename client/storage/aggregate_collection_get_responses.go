// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// AggregateCollectionGetReader is a Reader for the AggregateCollectionGet structure.
type AggregateCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregateCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregateCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAggregateCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAggregateCollectionGetOK creates a AggregateCollectionGetOK with default headers values
func NewAggregateCollectionGetOK() *AggregateCollectionGetOK {
	return &AggregateCollectionGetOK{}
}

/*
AggregateCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type AggregateCollectionGetOK struct {
	Payload *models.AggregateResponse
}

// IsSuccess returns true when this aggregate collection get o k response has a 2xx status code
func (o *AggregateCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregate collection get o k response has a 3xx status code
func (o *AggregateCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate collection get o k response has a 4xx status code
func (o *AggregateCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate collection get o k response has a 5xx status code
func (o *AggregateCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate collection get o k response a status code equal to that given
func (o *AggregateCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregate collection get o k response
func (o *AggregateCollectionGetOK) Code() int {
	return 200
}

func (o *AggregateCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/aggregates][%d] aggregateCollectionGetOK %s", 200, payload)
}

func (o *AggregateCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/aggregates][%d] aggregateCollectionGetOK %s", 200, payload)
}

func (o *AggregateCollectionGetOK) GetPayload() *models.AggregateResponse {
	return o.Payload
}

func (o *AggregateCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AggregateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateCollectionGetDefault creates a AggregateCollectionGetDefault with default headers values
func NewAggregateCollectionGetDefault(code int) *AggregateCollectionGetDefault {
	return &AggregateCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	AggregateCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 787092 | The target field cannot be specified for this operation. |
| 918138 | Internal error. Failed to get encryption operation status. |
| 8586225 | Encountered unexpected error in retrieving metrics and statistics for an aggregate. |
| 19726341 | Not enough eligible spare disks are available on the node. |
| 19726344 | No recommendation can be made for this cluster. |
| 19726357 | Aggregate recommendations are not supported on MetroCluster with Fibre Channel (FC). |
| 19726358 | Aggregate recommendations are not supported on ONTAP Cloud. |
| 19726382 | Another provisioning operation is in progress on this cluster. Wait a few minutes, and try the operation again. |
| 19726386 | Encountered an error when retrieving licensing information on this cluster. |
| 19726387 | No recommendation can be provided for this cluster within the license capacity. |
| 19726401 | Aggregate recommendations are not supported when the DR group is not in the "normal" state. |
| 19726402 | Internal error. Unable to determine the MetroCluster configuration state. |
| 19726403 | Aggregate recommendation is not supported when there are no healthy target connections to remote storage. |
| 19726404 | The recommended mirrored aggregate couldn't use all the attached capacity in one of the SyncMirror pools. Make sure that the remote and local storage is symmetrically wired. |
| 19726405 | Not all local and remote disks attached to the node have been auto-partitioned. |
| 19726406 | Aggregate recommendations are not supported on this node because remote and local storage is not symmetrically wired. |
| 19726540 | The next tag is not supported for recommended aggregates. Retry the operation with a higher "return_timeout" value. |
| 196608055 | Aggregate recommendation is not supported on this node because it does not support NetApp Aggregate Encryption (NAE). |
| 196608206 | Internal error. Failed to get encryption operation status. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AggregateCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this aggregate collection get default response has a 2xx status code
func (o *AggregateCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this aggregate collection get default response has a 3xx status code
func (o *AggregateCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this aggregate collection get default response has a 4xx status code
func (o *AggregateCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this aggregate collection get default response has a 5xx status code
func (o *AggregateCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this aggregate collection get default response a status code equal to that given
func (o *AggregateCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the aggregate collection get default response
func (o *AggregateCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *AggregateCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/aggregates][%d] aggregate_collection_get default %s", o._statusCode, payload)
}

func (o *AggregateCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/aggregates][%d] aggregate_collection_get default %s", o._statusCode, payload)
}

func (o *AggregateCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AggregateCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
