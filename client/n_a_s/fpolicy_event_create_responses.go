// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// FpolicyEventCreateReader is a Reader for the FpolicyEventCreate structure.
type FpolicyEventCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyEventCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewFpolicyEventCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyEventCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyEventCreateCreated creates a FpolicyEventCreateCreated with default headers values
func NewFpolicyEventCreateCreated() *FpolicyEventCreateCreated {
	return &FpolicyEventCreateCreated{}
}

/*
FpolicyEventCreateCreated describes a response with status code 201, with default header values.

Created
*/
type FpolicyEventCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.FpolicyEventResponse
}

// IsSuccess returns true when this fpolicy event create created response has a 2xx status code
func (o *FpolicyEventCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fpolicy event create created response has a 3xx status code
func (o *FpolicyEventCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fpolicy event create created response has a 4xx status code
func (o *FpolicyEventCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this fpolicy event create created response has a 5xx status code
func (o *FpolicyEventCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this fpolicy event create created response a status code equal to that given
func (o *FpolicyEventCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the fpolicy event create created response
func (o *FpolicyEventCreateCreated) Code() int {
	return 201
}

func (o *FpolicyEventCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/fpolicy/{svm.uuid}/events][%d] fpolicyEventCreateCreated %s", 201, payload)
}

func (o *FpolicyEventCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/fpolicy/{svm.uuid}/events][%d] fpolicyEventCreateCreated %s", 201, payload)
}

func (o *FpolicyEventCreateCreated) GetPayload() *models.FpolicyEventResponse {
	return o.Payload
}

func (o *FpolicyEventCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.FpolicyEventResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFpolicyEventCreateDefault creates a FpolicyEventCreateDefault with default headers values
func NewFpolicyEventCreateDefault(code int) *FpolicyEventCreateDefault {
	return &FpolicyEventCreateDefault{
		_statusCode: code,
	}
}

/*
	FpolicyEventCreateDefault describes a response with status code -1, with default header values.

	| Error Code | Description |

| ---------- | ----------- |
| 9764929    | The file operation is not supported by the protocol |
| 9764955    | The filter is not supported by the protocol |
| 9764930    | The filter is not supported by any of the file operations |
| 9764946    | The protocol is specified without a file operation or a file operation and filter pair |
*/
type FpolicyEventCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fpolicy event create default response has a 2xx status code
func (o *FpolicyEventCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fpolicy event create default response has a 3xx status code
func (o *FpolicyEventCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fpolicy event create default response has a 4xx status code
func (o *FpolicyEventCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fpolicy event create default response has a 5xx status code
func (o *FpolicyEventCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fpolicy event create default response a status code equal to that given
func (o *FpolicyEventCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fpolicy event create default response
func (o *FpolicyEventCreateDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyEventCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/fpolicy/{svm.uuid}/events][%d] fpolicy_event_create default %s", o._statusCode, payload)
}

func (o *FpolicyEventCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/fpolicy/{svm.uuid}/events][%d] fpolicy_event_create default %s", o._statusCode, payload)
}

func (o *FpolicyEventCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyEventCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
