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

// ApplicationComponentSnapshotCreateReader is a Reader for the ApplicationComponentSnapshotCreate structure.
type ApplicationComponentSnapshotCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationComponentSnapshotCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewApplicationComponentSnapshotCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewApplicationComponentSnapshotCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationComponentSnapshotCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationComponentSnapshotCreateCreated creates a ApplicationComponentSnapshotCreateCreated with default headers values
func NewApplicationComponentSnapshotCreateCreated() *ApplicationComponentSnapshotCreateCreated {
	return &ApplicationComponentSnapshotCreateCreated{}
}

/*
ApplicationComponentSnapshotCreateCreated describes a response with status code 201, with default header values.

Created
*/
type ApplicationComponentSnapshotCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this application component snapshot create created response has a 2xx status code
func (o *ApplicationComponentSnapshotCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application component snapshot create created response has a 3xx status code
func (o *ApplicationComponentSnapshotCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application component snapshot create created response has a 4xx status code
func (o *ApplicationComponentSnapshotCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this application component snapshot create created response has a 5xx status code
func (o *ApplicationComponentSnapshotCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this application component snapshot create created response a status code equal to that given
func (o *ApplicationComponentSnapshotCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the application component snapshot create created response
func (o *ApplicationComponentSnapshotCreateCreated) Code() int {
	return 201
}

func (o *ApplicationComponentSnapshotCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/components/{component.uuid}/snapshots][%d] applicationComponentSnapshotCreateCreated %s", 201, payload)
}

func (o *ApplicationComponentSnapshotCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/components/{component.uuid}/snapshots][%d] applicationComponentSnapshotCreateCreated %s", 201, payload)
}

func (o *ApplicationComponentSnapshotCreateCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ApplicationComponentSnapshotCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewApplicationComponentSnapshotCreateAccepted creates a ApplicationComponentSnapshotCreateAccepted with default headers values
func NewApplicationComponentSnapshotCreateAccepted() *ApplicationComponentSnapshotCreateAccepted {
	return &ApplicationComponentSnapshotCreateAccepted{}
}

/*
ApplicationComponentSnapshotCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ApplicationComponentSnapshotCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this application component snapshot create accepted response has a 2xx status code
func (o *ApplicationComponentSnapshotCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application component snapshot create accepted response has a 3xx status code
func (o *ApplicationComponentSnapshotCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application component snapshot create accepted response has a 4xx status code
func (o *ApplicationComponentSnapshotCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this application component snapshot create accepted response has a 5xx status code
func (o *ApplicationComponentSnapshotCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this application component snapshot create accepted response a status code equal to that given
func (o *ApplicationComponentSnapshotCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the application component snapshot create accepted response
func (o *ApplicationComponentSnapshotCreateAccepted) Code() int {
	return 202
}

func (o *ApplicationComponentSnapshotCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/components/{component.uuid}/snapshots][%d] applicationComponentSnapshotCreateAccepted %s", 202, payload)
}

func (o *ApplicationComponentSnapshotCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/components/{component.uuid}/snapshots][%d] applicationComponentSnapshotCreateAccepted %s", 202, payload)
}

func (o *ApplicationComponentSnapshotCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ApplicationComponentSnapshotCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewApplicationComponentSnapshotCreateDefault creates a ApplicationComponentSnapshotCreateDefault with default headers values
func NewApplicationComponentSnapshotCreateDefault(code int) *ApplicationComponentSnapshotCreateDefault {
	return &ApplicationComponentSnapshotCreateDefault{
		_statusCode: code,
	}
}

/*
ApplicationComponentSnapshotCreateDefault describes a response with status code -1, with default header values.

Error
*/
type ApplicationComponentSnapshotCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this application component snapshot create default response has a 2xx status code
func (o *ApplicationComponentSnapshotCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this application component snapshot create default response has a 3xx status code
func (o *ApplicationComponentSnapshotCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this application component snapshot create default response has a 4xx status code
func (o *ApplicationComponentSnapshotCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this application component snapshot create default response has a 5xx status code
func (o *ApplicationComponentSnapshotCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this application component snapshot create default response a status code equal to that given
func (o *ApplicationComponentSnapshotCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the application component snapshot create default response
func (o *ApplicationComponentSnapshotCreateDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationComponentSnapshotCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/components/{component.uuid}/snapshots][%d] application_component_snapshot_create default %s", o._statusCode, payload)
}

func (o *ApplicationComponentSnapshotCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/components/{component.uuid}/snapshots][%d] application_component_snapshot_create default %s", o._statusCode, payload)
}

func (o *ApplicationComponentSnapshotCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationComponentSnapshotCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
