// Code generated by go-swagger; DO NOT EDIT.

package security

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

// MultiAdminVerifyRequestCreateReader is a Reader for the MultiAdminVerifyRequestCreate structure.
type MultiAdminVerifyRequestCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MultiAdminVerifyRequestCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewMultiAdminVerifyRequestCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMultiAdminVerifyRequestCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMultiAdminVerifyRequestCreateCreated creates a MultiAdminVerifyRequestCreateCreated with default headers values
func NewMultiAdminVerifyRequestCreateCreated() *MultiAdminVerifyRequestCreateCreated {
	return &MultiAdminVerifyRequestCreateCreated{}
}

/*
MultiAdminVerifyRequestCreateCreated describes a response with status code 201, with default header values.

Created
*/
type MultiAdminVerifyRequestCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.MultiAdminVerifyRequestResponse
}

// IsSuccess returns true when this multi admin verify request create created response has a 2xx status code
func (o *MultiAdminVerifyRequestCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this multi admin verify request create created response has a 3xx status code
func (o *MultiAdminVerifyRequestCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this multi admin verify request create created response has a 4xx status code
func (o *MultiAdminVerifyRequestCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this multi admin verify request create created response has a 5xx status code
func (o *MultiAdminVerifyRequestCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this multi admin verify request create created response a status code equal to that given
func (o *MultiAdminVerifyRequestCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the multi admin verify request create created response
func (o *MultiAdminVerifyRequestCreateCreated) Code() int {
	return 201
}

func (o *MultiAdminVerifyRequestCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/multi-admin-verify/requests][%d] multiAdminVerifyRequestCreateCreated %s", 201, payload)
}

func (o *MultiAdminVerifyRequestCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/multi-admin-verify/requests][%d] multiAdminVerifyRequestCreateCreated %s", 201, payload)
}

func (o *MultiAdminVerifyRequestCreateCreated) GetPayload() *models.MultiAdminVerifyRequestResponse {
	return o.Payload
}

func (o *MultiAdminVerifyRequestCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.MultiAdminVerifyRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiAdminVerifyRequestCreateDefault creates a MultiAdminVerifyRequestCreateDefault with default headers values
func NewMultiAdminVerifyRequestCreateDefault(code int) *MultiAdminVerifyRequestCreateDefault {
	return &MultiAdminVerifyRequestCreateDefault{
		_statusCode: code,
	}
}

/*
	MultiAdminVerifyRequestCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262148 | The specified command is not recognized. |
| 262279 | Invalid field set. |
| 262304 | Too many requests. Delete one before creating another. |
| 262305 | Can't approve non-pending request. |
| 262306 | Can't veto an expired request. |
| 262308 | The specified command is not supported by this feature. |
| 262309 | The feature must be enabled first. |
| 262311 | Value must be greater than zero. |
| 262312 | Number of required approvers must be less than the total number of unique approvers in the approval-groups. |
| 262313 | Number of unique approvers in the approval-groups must be greater than the number of required approvers. |
| 262326 | Failed to parse query. |
| 262327 | Failed to crate the request. |
| 262328 | There is no matching rule for this request. |
| 262330 | Cannot approve/veto a request multiple times. |
| 262334 | The parameter specified in the command is not supported. |
| 262337 | Cannot approve/veto the user's own request. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type MultiAdminVerifyRequestCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this multi admin verify request create default response has a 2xx status code
func (o *MultiAdminVerifyRequestCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this multi admin verify request create default response has a 3xx status code
func (o *MultiAdminVerifyRequestCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this multi admin verify request create default response has a 4xx status code
func (o *MultiAdminVerifyRequestCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this multi admin verify request create default response has a 5xx status code
func (o *MultiAdminVerifyRequestCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this multi admin verify request create default response a status code equal to that given
func (o *MultiAdminVerifyRequestCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the multi admin verify request create default response
func (o *MultiAdminVerifyRequestCreateDefault) Code() int {
	return o._statusCode
}

func (o *MultiAdminVerifyRequestCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/multi-admin-verify/requests][%d] multi_admin_verify_request_create default %s", o._statusCode, payload)
}

func (o *MultiAdminVerifyRequestCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/multi-admin-verify/requests][%d] multi_admin_verify_request_create default %s", o._statusCode, payload)
}

func (o *MultiAdminVerifyRequestCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MultiAdminVerifyRequestCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
