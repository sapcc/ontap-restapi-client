// Code generated by go-swagger; DO NOT EDIT.

package application

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

// ConsistencyGroupSnapshotCreateReader is a Reader for the ConsistencyGroupSnapshotCreate structure.
type ConsistencyGroupSnapshotCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsistencyGroupSnapshotCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewConsistencyGroupSnapshotCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewConsistencyGroupSnapshotCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConsistencyGroupSnapshotCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConsistencyGroupSnapshotCreateCreated creates a ConsistencyGroupSnapshotCreateCreated with default headers values
func NewConsistencyGroupSnapshotCreateCreated() *ConsistencyGroupSnapshotCreateCreated {
	return &ConsistencyGroupSnapshotCreateCreated{}
}

/*
ConsistencyGroupSnapshotCreateCreated describes a response with status code 201, with default header values.

Created
*/
type ConsistencyGroupSnapshotCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.ConsistencyGroupSnapshotResponse
}

// IsSuccess returns true when this consistency group snapshot create created response has a 2xx status code
func (o *ConsistencyGroupSnapshotCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this consistency group snapshot create created response has a 3xx status code
func (o *ConsistencyGroupSnapshotCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this consistency group snapshot create created response has a 4xx status code
func (o *ConsistencyGroupSnapshotCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this consistency group snapshot create created response has a 5xx status code
func (o *ConsistencyGroupSnapshotCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this consistency group snapshot create created response a status code equal to that given
func (o *ConsistencyGroupSnapshotCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the consistency group snapshot create created response
func (o *ConsistencyGroupSnapshotCreateCreated) Code() int {
	return 201
}

func (o *ConsistencyGroupSnapshotCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/consistency-groups/{consistency_group.uuid}/snapshots][%d] consistencyGroupSnapshotCreateCreated %s", 201, payload)
}

func (o *ConsistencyGroupSnapshotCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/consistency-groups/{consistency_group.uuid}/snapshots][%d] consistencyGroupSnapshotCreateCreated %s", 201, payload)
}

func (o *ConsistencyGroupSnapshotCreateCreated) GetPayload() *models.ConsistencyGroupSnapshotResponse {
	return o.Payload
}

func (o *ConsistencyGroupSnapshotCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.ConsistencyGroupSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsistencyGroupSnapshotCreateAccepted creates a ConsistencyGroupSnapshotCreateAccepted with default headers values
func NewConsistencyGroupSnapshotCreateAccepted() *ConsistencyGroupSnapshotCreateAccepted {
	return &ConsistencyGroupSnapshotCreateAccepted{}
}

/*
ConsistencyGroupSnapshotCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ConsistencyGroupSnapshotCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this consistency group snapshot create accepted response has a 2xx status code
func (o *ConsistencyGroupSnapshotCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this consistency group snapshot create accepted response has a 3xx status code
func (o *ConsistencyGroupSnapshotCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this consistency group snapshot create accepted response has a 4xx status code
func (o *ConsistencyGroupSnapshotCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this consistency group snapshot create accepted response has a 5xx status code
func (o *ConsistencyGroupSnapshotCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this consistency group snapshot create accepted response a status code equal to that given
func (o *ConsistencyGroupSnapshotCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the consistency group snapshot create accepted response
func (o *ConsistencyGroupSnapshotCreateAccepted) Code() int {
	return 202
}

func (o *ConsistencyGroupSnapshotCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/consistency-groups/{consistency_group.uuid}/snapshots][%d] consistencyGroupSnapshotCreateAccepted %s", 202, payload)
}

func (o *ConsistencyGroupSnapshotCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/consistency-groups/{consistency_group.uuid}/snapshots][%d] consistencyGroupSnapshotCreateAccepted %s", 202, payload)
}

func (o *ConsistencyGroupSnapshotCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ConsistencyGroupSnapshotCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewConsistencyGroupSnapshotCreateDefault creates a ConsistencyGroupSnapshotCreateDefault with default headers values
func NewConsistencyGroupSnapshotCreateDefault(code int) *ConsistencyGroupSnapshotCreateDefault {
	return &ConsistencyGroupSnapshotCreateDefault{
		_statusCode: code,
	}
}

/*
	ConsistencyGroupSnapshotCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 53411918 | Snapshot operation not permitted. |
| 53411921 | Failed to create snapshot for consistency group. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ConsistencyGroupSnapshotCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this consistency group snapshot create default response has a 2xx status code
func (o *ConsistencyGroupSnapshotCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this consistency group snapshot create default response has a 3xx status code
func (o *ConsistencyGroupSnapshotCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this consistency group snapshot create default response has a 4xx status code
func (o *ConsistencyGroupSnapshotCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this consistency group snapshot create default response has a 5xx status code
func (o *ConsistencyGroupSnapshotCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this consistency group snapshot create default response a status code equal to that given
func (o *ConsistencyGroupSnapshotCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the consistency group snapshot create default response
func (o *ConsistencyGroupSnapshotCreateDefault) Code() int {
	return o._statusCode
}

func (o *ConsistencyGroupSnapshotCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/consistency-groups/{consistency_group.uuid}/snapshots][%d] consistency_group_snapshot_create default %s", o._statusCode, payload)
}

func (o *ConsistencyGroupSnapshotCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/consistency-groups/{consistency_group.uuid}/snapshots][%d] consistency_group_snapshot_create default %s", o._statusCode, payload)
}

func (o *ConsistencyGroupSnapshotCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConsistencyGroupSnapshotCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
