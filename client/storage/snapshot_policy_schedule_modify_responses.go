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

	"github.com/sapcc/ontap-restapi-client/models"
)

// SnapshotPolicyScheduleModifyReader is a Reader for the SnapshotPolicyScheduleModify structure.
type SnapshotPolicyScheduleModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotPolicyScheduleModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotPolicyScheduleModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotPolicyScheduleModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotPolicyScheduleModifyOK creates a SnapshotPolicyScheduleModifyOK with default headers values
func NewSnapshotPolicyScheduleModifyOK() *SnapshotPolicyScheduleModifyOK {
	return &SnapshotPolicyScheduleModifyOK{}
}

/*
SnapshotPolicyScheduleModifyOK describes a response with status code 200, with default header values.

OK
*/
type SnapshotPolicyScheduleModifyOK struct {
}

// IsSuccess returns true when this snapshot policy schedule modify o k response has a 2xx status code
func (o *SnapshotPolicyScheduleModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapshot policy schedule modify o k response has a 3xx status code
func (o *SnapshotPolicyScheduleModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapshot policy schedule modify o k response has a 4xx status code
func (o *SnapshotPolicyScheduleModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapshot policy schedule modify o k response has a 5xx status code
func (o *SnapshotPolicyScheduleModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snapshot policy schedule modify o k response a status code equal to that given
func (o *SnapshotPolicyScheduleModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snapshot policy schedule modify o k response
func (o *SnapshotPolicyScheduleModifyOK) Code() int {
	return 200
}

func (o *SnapshotPolicyScheduleModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /storage/snapshot-policies/{snapshot_policy.uuid}/schedules/{schedule.uuid}][%d] snapshotPolicyScheduleModifyOK", 200)
}

func (o *SnapshotPolicyScheduleModifyOK) String() string {
	return fmt.Sprintf("[PATCH /storage/snapshot-policies/{snapshot_policy.uuid}/schedules/{schedule.uuid}][%d] snapshotPolicyScheduleModifyOK", 200)
}

func (o *SnapshotPolicyScheduleModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSnapshotPolicyScheduleModifyDefault creates a SnapshotPolicyScheduleModifyDefault with default headers values
func NewSnapshotPolicyScheduleModifyDefault(code int) *SnapshotPolicyScheduleModifyDefault {
	return &SnapshotPolicyScheduleModifyDefault{
		_statusCode: code,
	}
}

/*
	SnapshotPolicyScheduleModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Code

| Error Code | Description |
| ---------- | ----------- |
| 1638451    | This operation would result in total snapshot count for the policy to exceed maximum supported count. |
| 918253     | Incorrect format for the retention period, duration must be in the ISO-8601 format. |
*/
type SnapshotPolicyScheduleModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snapshot policy schedule modify default response has a 2xx status code
func (o *SnapshotPolicyScheduleModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapshot policy schedule modify default response has a 3xx status code
func (o *SnapshotPolicyScheduleModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapshot policy schedule modify default response has a 4xx status code
func (o *SnapshotPolicyScheduleModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapshot policy schedule modify default response has a 5xx status code
func (o *SnapshotPolicyScheduleModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapshot policy schedule modify default response a status code equal to that given
func (o *SnapshotPolicyScheduleModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snapshot policy schedule modify default response
func (o *SnapshotPolicyScheduleModifyDefault) Code() int {
	return o._statusCode
}

func (o *SnapshotPolicyScheduleModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/snapshot-policies/{snapshot_policy.uuid}/schedules/{schedule.uuid}][%d] snapshot_policy_schedule_modify default %s", o._statusCode, payload)
}

func (o *SnapshotPolicyScheduleModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/snapshot-policies/{snapshot_policy.uuid}/schedules/{schedule.uuid}][%d] snapshot_policy_schedule_modify default %s", o._statusCode, payload)
}

func (o *SnapshotPolicyScheduleModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotPolicyScheduleModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
