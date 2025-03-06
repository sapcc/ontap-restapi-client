// Code generated by go-swagger; DO NOT EDIT.

package snap_lock

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

// SnaplockLogDeleteReader is a Reader for the SnaplockLogDelete structure.
type SnaplockLogDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockLogDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockLogDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSnaplockLogDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockLogDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockLogDeleteOK creates a SnaplockLogDeleteOK with default headers values
func NewSnaplockLogDeleteOK() *SnaplockLogDeleteOK {
	return &SnaplockLogDeleteOK{}
}

/*
SnaplockLogDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockLogDeleteOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this snaplock log delete o k response has a 2xx status code
func (o *SnaplockLogDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock log delete o k response has a 3xx status code
func (o *SnaplockLogDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock log delete o k response has a 4xx status code
func (o *SnaplockLogDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock log delete o k response has a 5xx status code
func (o *SnaplockLogDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock log delete o k response a status code equal to that given
func (o *SnaplockLogDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snaplock log delete o k response
func (o *SnaplockLogDeleteOK) Code() int {
	return 200
}

func (o *SnaplockLogDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplockLogDeleteOK %s", 200, payload)
}

func (o *SnaplockLogDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplockLogDeleteOK %s", 200, payload)
}

func (o *SnaplockLogDeleteOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnaplockLogDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLogDeleteAccepted creates a SnaplockLogDeleteAccepted with default headers values
func NewSnaplockLogDeleteAccepted() *SnaplockLogDeleteAccepted {
	return &SnaplockLogDeleteAccepted{}
}

/*
SnaplockLogDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnaplockLogDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this snaplock log delete accepted response has a 2xx status code
func (o *SnaplockLogDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock log delete accepted response has a 3xx status code
func (o *SnaplockLogDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock log delete accepted response has a 4xx status code
func (o *SnaplockLogDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock log delete accepted response has a 5xx status code
func (o *SnaplockLogDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock log delete accepted response a status code equal to that given
func (o *SnaplockLogDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the snaplock log delete accepted response
func (o *SnaplockLogDeleteAccepted) Code() int {
	return 202
}

func (o *SnaplockLogDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplockLogDeleteAccepted %s", 202, payload)
}

func (o *SnaplockLogDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplockLogDeleteAccepted %s", 202, payload)
}

func (o *SnaplockLogDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnaplockLogDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLogDeleteDefault creates a SnaplockLogDeleteDefault with default headers values
func NewSnaplockLogDeleteDefault(code int) *SnaplockLogDeleteDefault {
	return &SnaplockLogDeleteDefault{
		_statusCode: code,
	}
}

/*
	SnaplockLogDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 14090346    | Internal Error. Wait a few minutes, then try the command again  |
*/
type SnaplockLogDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snaplock log delete default response has a 2xx status code
func (o *SnaplockLogDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock log delete default response has a 3xx status code
func (o *SnaplockLogDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock log delete default response has a 4xx status code
func (o *SnaplockLogDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock log delete default response has a 5xx status code
func (o *SnaplockLogDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock log delete default response a status code equal to that given
func (o *SnaplockLogDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snaplock log delete default response
func (o *SnaplockLogDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockLogDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplock_log_delete default %s", o._statusCode, payload)
}

func (o *SnaplockLogDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplock_log_delete default %s", o._statusCode, payload)
}

func (o *SnaplockLogDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockLogDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
