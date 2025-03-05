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

// ScheduleDeleteReader is a Reader for the ScheduleDelete structure.
type ScheduleDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduleDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScheduleDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewScheduleDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScheduleDeleteOK creates a ScheduleDeleteOK with default headers values
func NewScheduleDeleteOK() *ScheduleDeleteOK {
	return &ScheduleDeleteOK{}
}

/*
ScheduleDeleteOK describes a response with status code 200, with default header values.

OK
*/
type ScheduleDeleteOK struct {
}

// IsSuccess returns true when this schedule delete o k response has a 2xx status code
func (o *ScheduleDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schedule delete o k response has a 3xx status code
func (o *ScheduleDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule delete o k response has a 4xx status code
func (o *ScheduleDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this schedule delete o k response has a 5xx status code
func (o *ScheduleDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule delete o k response a status code equal to that given
func (o *ScheduleDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the schedule delete o k response
func (o *ScheduleDeleteOK) Code() int {
	return 200
}

func (o *ScheduleDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /cluster/schedules/{uuid}][%d] scheduleDeleteOK", 200)
}

func (o *ScheduleDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /cluster/schedules/{uuid}][%d] scheduleDeleteOK", 200)
}

func (o *ScheduleDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScheduleDeleteDefault creates a ScheduleDeleteDefault with default headers values
func NewScheduleDeleteDefault(code int) *ScheduleDeleteDefault {
	return &ScheduleDeleteDefault{
		_statusCode: code,
	}
}

/*
	ScheduleDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 459758 | Cannot delete a job schedule that is in use. Remove all references to the schedule, and then try to delete again. |
| 459761 | Schedule cannot be deleted on this cluster because it is replicated from the remote cluster. |
| 459762 | The schedule cannot be deleted because it is a system-level schedule. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ScheduleDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this schedule delete default response has a 2xx status code
func (o *ScheduleDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this schedule delete default response has a 3xx status code
func (o *ScheduleDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this schedule delete default response has a 4xx status code
func (o *ScheduleDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this schedule delete default response has a 5xx status code
func (o *ScheduleDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this schedule delete default response a status code equal to that given
func (o *ScheduleDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the schedule delete default response
func (o *ScheduleDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ScheduleDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/schedules/{uuid}][%d] schedule_delete default %s", o._statusCode, payload)
}

func (o *ScheduleDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/schedules/{uuid}][%d] schedule_delete default %s", o._statusCode, payload)
}

func (o *ScheduleDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ScheduleDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
