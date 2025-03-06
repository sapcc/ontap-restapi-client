// Code generated by go-swagger; DO NOT EDIT.

package support

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

// AutosupportCreateReader is a Reader for the AutosupportCreate structure.
type AutosupportCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutosupportCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAutosupportCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAutosupportCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAutosupportCreateCreated creates a AutosupportCreateCreated with default headers values
func NewAutosupportCreateCreated() *AutosupportCreateCreated {
	return &AutosupportCreateCreated{}
}

/*
AutosupportCreateCreated describes a response with status code 201, with default header values.

Created
*/
type AutosupportCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.AutosupportMessageResponse
}

// IsSuccess returns true when this autosupport create created response has a 2xx status code
func (o *AutosupportCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this autosupport create created response has a 3xx status code
func (o *AutosupportCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autosupport create created response has a 4xx status code
func (o *AutosupportCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this autosupport create created response has a 5xx status code
func (o *AutosupportCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this autosupport create created response a status code equal to that given
func (o *AutosupportCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the autosupport create created response
func (o *AutosupportCreateCreated) Code() int {
	return 201
}

func (o *AutosupportCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/autosupport/messages][%d] autosupportCreateCreated %s", 201, payload)
}

func (o *AutosupportCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/autosupport/messages][%d] autosupportCreateCreated %s", 201, payload)
}

func (o *AutosupportCreateCreated) GetPayload() *models.AutosupportMessageResponse {
	return o.Payload
}

func (o *AutosupportCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.AutosupportMessageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutosupportCreateDefault creates a AutosupportCreateDefault with default headers values
func NewAutosupportCreateDefault(code int) *AutosupportCreateDefault {
	return &AutosupportCreateDefault{
		_statusCode: code,
	}
}

/*
	AutosupportCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8650851 | AutoSupport invocation failed because AutoSupport collection is disabled |
| 8650852 | AutoSupport invocation failed because the trigger is disabled |
| 8650853 | AutoSupport invocation failed because the trigger rate is higher than the maximum supported value |
| 8650866 | The message parameter is not supported with performance AutoSupports |
| 8650884 | The specified field is not allowed for the current operation |
| 8650885 | The specified field cannot be empty |
| 53149748 | The destination URI provided for the invoked AutoSupport is invalid |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AutosupportCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this autosupport create default response has a 2xx status code
func (o *AutosupportCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this autosupport create default response has a 3xx status code
func (o *AutosupportCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this autosupport create default response has a 4xx status code
func (o *AutosupportCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this autosupport create default response has a 5xx status code
func (o *AutosupportCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this autosupport create default response a status code equal to that given
func (o *AutosupportCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the autosupport create default response
func (o *AutosupportCreateDefault) Code() int {
	return o._statusCode
}

func (o *AutosupportCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/autosupport/messages][%d] autosupport_create default %s", o._statusCode, payload)
}

func (o *AutosupportCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/autosupport/messages][%d] autosupport_create default %s", o._statusCode, payload)
}

func (o *AutosupportCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AutosupportCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
