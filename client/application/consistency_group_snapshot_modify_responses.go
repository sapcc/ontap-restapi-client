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

	"github.com/sapcc/ontap-restapi-client/models"
)

// ConsistencyGroupSnapshotModifyReader is a Reader for the ConsistencyGroupSnapshotModify structure.
type ConsistencyGroupSnapshotModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsistencyGroupSnapshotModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConsistencyGroupSnapshotModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewConsistencyGroupSnapshotModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConsistencyGroupSnapshotModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConsistencyGroupSnapshotModifyOK creates a ConsistencyGroupSnapshotModifyOK with default headers values
func NewConsistencyGroupSnapshotModifyOK() *ConsistencyGroupSnapshotModifyOK {
	return &ConsistencyGroupSnapshotModifyOK{}
}

/*
ConsistencyGroupSnapshotModifyOK describes a response with status code 200, with default header values.

OK
*/
type ConsistencyGroupSnapshotModifyOK struct {
}

// IsSuccess returns true when this consistency group snapshot modify o k response has a 2xx status code
func (o *ConsistencyGroupSnapshotModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this consistency group snapshot modify o k response has a 3xx status code
func (o *ConsistencyGroupSnapshotModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this consistency group snapshot modify o k response has a 4xx status code
func (o *ConsistencyGroupSnapshotModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this consistency group snapshot modify o k response has a 5xx status code
func (o *ConsistencyGroupSnapshotModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this consistency group snapshot modify o k response a status code equal to that given
func (o *ConsistencyGroupSnapshotModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the consistency group snapshot modify o k response
func (o *ConsistencyGroupSnapshotModifyOK) Code() int {
	return 200
}

func (o *ConsistencyGroupSnapshotModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /application/consistency-groups/{consistency_group.uuid}/snapshots/{uuid}][%d] consistencyGroupSnapshotModifyOK", 200)
}

func (o *ConsistencyGroupSnapshotModifyOK) String() string {
	return fmt.Sprintf("[PATCH /application/consistency-groups/{consistency_group.uuid}/snapshots/{uuid}][%d] consistencyGroupSnapshotModifyOK", 200)
}

func (o *ConsistencyGroupSnapshotModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConsistencyGroupSnapshotModifyAccepted creates a ConsistencyGroupSnapshotModifyAccepted with default headers values
func NewConsistencyGroupSnapshotModifyAccepted() *ConsistencyGroupSnapshotModifyAccepted {
	return &ConsistencyGroupSnapshotModifyAccepted{}
}

/*
ConsistencyGroupSnapshotModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ConsistencyGroupSnapshotModifyAccepted struct {
}

// IsSuccess returns true when this consistency group snapshot modify accepted response has a 2xx status code
func (o *ConsistencyGroupSnapshotModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this consistency group snapshot modify accepted response has a 3xx status code
func (o *ConsistencyGroupSnapshotModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this consistency group snapshot modify accepted response has a 4xx status code
func (o *ConsistencyGroupSnapshotModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this consistency group snapshot modify accepted response has a 5xx status code
func (o *ConsistencyGroupSnapshotModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this consistency group snapshot modify accepted response a status code equal to that given
func (o *ConsistencyGroupSnapshotModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the consistency group snapshot modify accepted response
func (o *ConsistencyGroupSnapshotModifyAccepted) Code() int {
	return 202
}

func (o *ConsistencyGroupSnapshotModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /application/consistency-groups/{consistency_group.uuid}/snapshots/{uuid}][%d] consistencyGroupSnapshotModifyAccepted", 202)
}

func (o *ConsistencyGroupSnapshotModifyAccepted) String() string {
	return fmt.Sprintf("[PATCH /application/consistency-groups/{consistency_group.uuid}/snapshots/{uuid}][%d] consistencyGroupSnapshotModifyAccepted", 202)
}

func (o *ConsistencyGroupSnapshotModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConsistencyGroupSnapshotModifyDefault creates a ConsistencyGroupSnapshotModifyDefault with default headers values
func NewConsistencyGroupSnapshotModifyDefault(code int) *ConsistencyGroupSnapshotModifyDefault {
	return &ConsistencyGroupSnapshotModifyDefault{
		_statusCode: code,
	}
}

/*
	ConsistencyGroupSnapshotModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 53411925 | Failed to find a previously initiated two-phase snapshot operation for consistency group. |
| 53412015 | SnapLock expiry time cannot be before the current expiry time. |
| 53412016 | Snapshot is not retained by SnapLock. |
| 53412017 | The operation is not supported on this platform. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ConsistencyGroupSnapshotModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this consistency group snapshot modify default response has a 2xx status code
func (o *ConsistencyGroupSnapshotModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this consistency group snapshot modify default response has a 3xx status code
func (o *ConsistencyGroupSnapshotModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this consistency group snapshot modify default response has a 4xx status code
func (o *ConsistencyGroupSnapshotModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this consistency group snapshot modify default response has a 5xx status code
func (o *ConsistencyGroupSnapshotModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this consistency group snapshot modify default response a status code equal to that given
func (o *ConsistencyGroupSnapshotModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the consistency group snapshot modify default response
func (o *ConsistencyGroupSnapshotModifyDefault) Code() int {
	return o._statusCode
}

func (o *ConsistencyGroupSnapshotModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /application/consistency-groups/{consistency_group.uuid}/snapshots/{uuid}][%d] consistency_group_snapshot_modify default %s", o._statusCode, payload)
}

func (o *ConsistencyGroupSnapshotModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /application/consistency-groups/{consistency_group.uuid}/snapshots/{uuid}][%d] consistency_group_snapshot_modify default %s", o._statusCode, payload)
}

func (o *ConsistencyGroupSnapshotModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConsistencyGroupSnapshotModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
