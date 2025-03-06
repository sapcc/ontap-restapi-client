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

// ApplicationCreateReader is a Reader for the ApplicationCreate structure.
type ApplicationCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewApplicationCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewApplicationCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationCreateCreated creates a ApplicationCreateCreated with default headers values
func NewApplicationCreateCreated() *ApplicationCreateCreated {
	return &ApplicationCreateCreated{}
}

/*
ApplicationCreateCreated describes a response with status code 201, with default header values.

Created
*/
type ApplicationCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this application create created response has a 2xx status code
func (o *ApplicationCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application create created response has a 3xx status code
func (o *ApplicationCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application create created response has a 4xx status code
func (o *ApplicationCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this application create created response has a 5xx status code
func (o *ApplicationCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this application create created response a status code equal to that given
func (o *ApplicationCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the application create created response
func (o *ApplicationCreateCreated) Code() int {
	return 201
}

func (o *ApplicationCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications][%d] applicationCreateCreated %s", 201, payload)
}

func (o *ApplicationCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications][%d] applicationCreateCreated %s", 201, payload)
}

func (o *ApplicationCreateCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ApplicationCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewApplicationCreateAccepted creates a ApplicationCreateAccepted with default headers values
func NewApplicationCreateAccepted() *ApplicationCreateAccepted {
	return &ApplicationCreateAccepted{}
}

/*
ApplicationCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ApplicationCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this application create accepted response has a 2xx status code
func (o *ApplicationCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application create accepted response has a 3xx status code
func (o *ApplicationCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application create accepted response has a 4xx status code
func (o *ApplicationCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this application create accepted response has a 5xx status code
func (o *ApplicationCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this application create accepted response a status code equal to that given
func (o *ApplicationCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the application create accepted response
func (o *ApplicationCreateAccepted) Code() int {
	return 202
}

func (o *ApplicationCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications][%d] applicationCreateAccepted %s", 202, payload)
}

func (o *ApplicationCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications][%d] applicationCreateAccepted %s", 202, payload)
}

func (o *ApplicationCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ApplicationCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewApplicationCreateDefault creates a ApplicationCreateDefault with default headers values
func NewApplicationCreateDefault(code int) *ApplicationCreateDefault {
	return &ApplicationCreateDefault{
		_statusCode: code,
	}
}

/*
	ApplicationCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65995775 | The size provided is too small. |
| 65995776 | The size provided is too large. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ApplicationCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this application create default response has a 2xx status code
func (o *ApplicationCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this application create default response has a 3xx status code
func (o *ApplicationCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this application create default response has a 4xx status code
func (o *ApplicationCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this application create default response has a 5xx status code
func (o *ApplicationCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this application create default response a status code equal to that given
func (o *ApplicationCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the application create default response
func (o *ApplicationCreateDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications][%d] application_create default %s", o._statusCode, payload)
}

func (o *ApplicationCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications][%d] application_create default %s", o._statusCode, payload)
}

func (o *ApplicationCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
