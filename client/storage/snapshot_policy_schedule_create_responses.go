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

// SnapshotPolicyScheduleCreateReader is a Reader for the SnapshotPolicyScheduleCreate structure.
type SnapshotPolicyScheduleCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotPolicyScheduleCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSnapshotPolicyScheduleCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotPolicyScheduleCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotPolicyScheduleCreateCreated creates a SnapshotPolicyScheduleCreateCreated with default headers values
func NewSnapshotPolicyScheduleCreateCreated() *SnapshotPolicyScheduleCreateCreated {
	return &SnapshotPolicyScheduleCreateCreated{}
}

/*
SnapshotPolicyScheduleCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SnapshotPolicyScheduleCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string
}

// IsSuccess returns true when this snapshot policy schedule create created response has a 2xx status code
func (o *SnapshotPolicyScheduleCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapshot policy schedule create created response has a 3xx status code
func (o *SnapshotPolicyScheduleCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapshot policy schedule create created response has a 4xx status code
func (o *SnapshotPolicyScheduleCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapshot policy schedule create created response has a 5xx status code
func (o *SnapshotPolicyScheduleCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this snapshot policy schedule create created response a status code equal to that given
func (o *SnapshotPolicyScheduleCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the snapshot policy schedule create created response
func (o *SnapshotPolicyScheduleCreateCreated) Code() int {
	return 201
}

func (o *SnapshotPolicyScheduleCreateCreated) Error() string {
	return fmt.Sprintf("[POST /storage/snapshot-policies/{snapshot_policy.uuid}/schedules][%d] snapshotPolicyScheduleCreateCreated", 201)
}

func (o *SnapshotPolicyScheduleCreateCreated) String() string {
	return fmt.Sprintf("[POST /storage/snapshot-policies/{snapshot_policy.uuid}/schedules][%d] snapshotPolicyScheduleCreateCreated", 201)
}

func (o *SnapshotPolicyScheduleCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewSnapshotPolicyScheduleCreateDefault creates a SnapshotPolicyScheduleCreateDefault with default headers values
func NewSnapshotPolicyScheduleCreateDefault(code int) *SnapshotPolicyScheduleCreateDefault {
	return &SnapshotPolicyScheduleCreateDefault{
		_statusCode: code,
	}
}

/*
	SnapshotPolicyScheduleCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1638407    | When adding schedule to a snapshot policy, the count for that schedule must be specified. |
| 1638410    | Specified schedule already exists in snapshot policy. |
| 1638413    | Schedule not found. |
| 1638451    | This operation would result in total snapshot count for the policy to exceed maximum supported count. |
| 1638508    | Another schedule has the same prefix within this policy. |
| 1638510    | Duplicate prefix. |
| 1638528    | This operation is not supported in a mixed-version cluster. |
| 1638531    | This operation is not supported because specified policy is owned by the cluster admin. |
| 918253     | Incorrect format for the retention period, duration must be in the ISO-8601 format. |
*/
type SnapshotPolicyScheduleCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snapshot policy schedule create default response has a 2xx status code
func (o *SnapshotPolicyScheduleCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapshot policy schedule create default response has a 3xx status code
func (o *SnapshotPolicyScheduleCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapshot policy schedule create default response has a 4xx status code
func (o *SnapshotPolicyScheduleCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapshot policy schedule create default response has a 5xx status code
func (o *SnapshotPolicyScheduleCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapshot policy schedule create default response a status code equal to that given
func (o *SnapshotPolicyScheduleCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snapshot policy schedule create default response
func (o *SnapshotPolicyScheduleCreateDefault) Code() int {
	return o._statusCode
}

func (o *SnapshotPolicyScheduleCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/snapshot-policies/{snapshot_policy.uuid}/schedules][%d] snapshot_policy_schedule_create default %s", o._statusCode, payload)
}

func (o *SnapshotPolicyScheduleCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/snapshot-policies/{snapshot_policy.uuid}/schedules][%d] snapshot_policy_schedule_create default %s", o._statusCode, payload)
}

func (o *SnapshotPolicyScheduleCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotPolicyScheduleCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
