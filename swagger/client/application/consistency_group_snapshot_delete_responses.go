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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// ConsistencyGroupSnapshotDeleteReader is a Reader for the ConsistencyGroupSnapshotDelete structure.
type ConsistencyGroupSnapshotDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsistencyGroupSnapshotDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConsistencyGroupSnapshotDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewConsistencyGroupSnapshotDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConsistencyGroupSnapshotDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConsistencyGroupSnapshotDeleteOK creates a ConsistencyGroupSnapshotDeleteOK with default headers values
func NewConsistencyGroupSnapshotDeleteOK() *ConsistencyGroupSnapshotDeleteOK {
	return &ConsistencyGroupSnapshotDeleteOK{}
}

/*
ConsistencyGroupSnapshotDeleteOK describes a response with status code 200, with default header values.

OK
*/
type ConsistencyGroupSnapshotDeleteOK struct {
}

// IsSuccess returns true when this consistency group snapshot delete o k response has a 2xx status code
func (o *ConsistencyGroupSnapshotDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this consistency group snapshot delete o k response has a 3xx status code
func (o *ConsistencyGroupSnapshotDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this consistency group snapshot delete o k response has a 4xx status code
func (o *ConsistencyGroupSnapshotDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this consistency group snapshot delete o k response has a 5xx status code
func (o *ConsistencyGroupSnapshotDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this consistency group snapshot delete o k response a status code equal to that given
func (o *ConsistencyGroupSnapshotDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the consistency group snapshot delete o k response
func (o *ConsistencyGroupSnapshotDeleteOK) Code() int {
	return 200
}

func (o *ConsistencyGroupSnapshotDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /application/consistency-groups/{consistency_group.uuid}/snapshots/{uuid}][%d] consistencyGroupSnapshotDeleteOK", 200)
}

func (o *ConsistencyGroupSnapshotDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /application/consistency-groups/{consistency_group.uuid}/snapshots/{uuid}][%d] consistencyGroupSnapshotDeleteOK", 200)
}

func (o *ConsistencyGroupSnapshotDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConsistencyGroupSnapshotDeleteAccepted creates a ConsistencyGroupSnapshotDeleteAccepted with default headers values
func NewConsistencyGroupSnapshotDeleteAccepted() *ConsistencyGroupSnapshotDeleteAccepted {
	return &ConsistencyGroupSnapshotDeleteAccepted{}
}

/*
ConsistencyGroupSnapshotDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ConsistencyGroupSnapshotDeleteAccepted struct {
}

// IsSuccess returns true when this consistency group snapshot delete accepted response has a 2xx status code
func (o *ConsistencyGroupSnapshotDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this consistency group snapshot delete accepted response has a 3xx status code
func (o *ConsistencyGroupSnapshotDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this consistency group snapshot delete accepted response has a 4xx status code
func (o *ConsistencyGroupSnapshotDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this consistency group snapshot delete accepted response has a 5xx status code
func (o *ConsistencyGroupSnapshotDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this consistency group snapshot delete accepted response a status code equal to that given
func (o *ConsistencyGroupSnapshotDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the consistency group snapshot delete accepted response
func (o *ConsistencyGroupSnapshotDeleteAccepted) Code() int {
	return 202
}

func (o *ConsistencyGroupSnapshotDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /application/consistency-groups/{consistency_group.uuid}/snapshots/{uuid}][%d] consistencyGroupSnapshotDeleteAccepted", 202)
}

func (o *ConsistencyGroupSnapshotDeleteAccepted) String() string {
	return fmt.Sprintf("[DELETE /application/consistency-groups/{consistency_group.uuid}/snapshots/{uuid}][%d] consistencyGroupSnapshotDeleteAccepted", 202)
}

func (o *ConsistencyGroupSnapshotDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConsistencyGroupSnapshotDeleteDefault creates a ConsistencyGroupSnapshotDeleteDefault with default headers values
func NewConsistencyGroupSnapshotDeleteDefault(code int) *ConsistencyGroupSnapshotDeleteDefault {
	return &ConsistencyGroupSnapshotDeleteDefault{
		_statusCode: code,
	}
}

/*
	ConsistencyGroupSnapshotDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 53412007 | Failed to delete the snapshot because it has not expired or is locked. |
| 53412008 | Failed to delete the consistency group snapshot because it is currently used as a reference for a replication relationship. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ConsistencyGroupSnapshotDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this consistency group snapshot delete default response has a 2xx status code
func (o *ConsistencyGroupSnapshotDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this consistency group snapshot delete default response has a 3xx status code
func (o *ConsistencyGroupSnapshotDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this consistency group snapshot delete default response has a 4xx status code
func (o *ConsistencyGroupSnapshotDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this consistency group snapshot delete default response has a 5xx status code
func (o *ConsistencyGroupSnapshotDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this consistency group snapshot delete default response a status code equal to that given
func (o *ConsistencyGroupSnapshotDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the consistency group snapshot delete default response
func (o *ConsistencyGroupSnapshotDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ConsistencyGroupSnapshotDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/consistency-groups/{consistency_group.uuid}/snapshots/{uuid}][%d] consistency_group_snapshot_delete default %s", o._statusCode, payload)
}

func (o *ConsistencyGroupSnapshotDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/consistency-groups/{consistency_group.uuid}/snapshots/{uuid}][%d] consistency_group_snapshot_delete default %s", o._statusCode, payload)
}

func (o *ConsistencyGroupSnapshotDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConsistencyGroupSnapshotDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
