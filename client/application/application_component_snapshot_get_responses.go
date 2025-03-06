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

// ApplicationComponentSnapshotGetReader is a Reader for the ApplicationComponentSnapshotGet structure.
type ApplicationComponentSnapshotGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationComponentSnapshotGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationComponentSnapshotGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationComponentSnapshotGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationComponentSnapshotGetOK creates a ApplicationComponentSnapshotGetOK with default headers values
func NewApplicationComponentSnapshotGetOK() *ApplicationComponentSnapshotGetOK {
	return &ApplicationComponentSnapshotGetOK{}
}

/*
ApplicationComponentSnapshotGetOK describes a response with status code 200, with default header values.

OK
*/
type ApplicationComponentSnapshotGetOK struct {
	Payload *models.ApplicationComponentSnapshot
}

// IsSuccess returns true when this application component snapshot get o k response has a 2xx status code
func (o *ApplicationComponentSnapshotGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application component snapshot get o k response has a 3xx status code
func (o *ApplicationComponentSnapshotGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application component snapshot get o k response has a 4xx status code
func (o *ApplicationComponentSnapshotGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this application component snapshot get o k response has a 5xx status code
func (o *ApplicationComponentSnapshotGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this application component snapshot get o k response a status code equal to that given
func (o *ApplicationComponentSnapshotGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the application component snapshot get o k response
func (o *ApplicationComponentSnapshotGetOK) Code() int {
	return 200
}

func (o *ApplicationComponentSnapshotGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/components/{component.uuid}/snapshots/{uuid}][%d] applicationComponentSnapshotGetOK %s", 200, payload)
}

func (o *ApplicationComponentSnapshotGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/components/{component.uuid}/snapshots/{uuid}][%d] applicationComponentSnapshotGetOK %s", 200, payload)
}

func (o *ApplicationComponentSnapshotGetOK) GetPayload() *models.ApplicationComponentSnapshot {
	return o.Payload
}

func (o *ApplicationComponentSnapshotGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationComponentSnapshot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationComponentSnapshotGetDefault creates a ApplicationComponentSnapshotGetDefault with default headers values
func NewApplicationComponentSnapshotGetDefault(code int) *ApplicationComponentSnapshotGetDefault {
	return &ApplicationComponentSnapshotGetDefault{
		_statusCode: code,
	}
}

/*
ApplicationComponentSnapshotGetDefault describes a response with status code -1, with default header values.

Error
*/
type ApplicationComponentSnapshotGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this application component snapshot get default response has a 2xx status code
func (o *ApplicationComponentSnapshotGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this application component snapshot get default response has a 3xx status code
func (o *ApplicationComponentSnapshotGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this application component snapshot get default response has a 4xx status code
func (o *ApplicationComponentSnapshotGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this application component snapshot get default response has a 5xx status code
func (o *ApplicationComponentSnapshotGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this application component snapshot get default response a status code equal to that given
func (o *ApplicationComponentSnapshotGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the application component snapshot get default response
func (o *ApplicationComponentSnapshotGetDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationComponentSnapshotGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/components/{component.uuid}/snapshots/{uuid}][%d] application_component_snapshot_get default %s", o._statusCode, payload)
}

func (o *ApplicationComponentSnapshotGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/components/{component.uuid}/snapshots/{uuid}][%d] application_component_snapshot_get default %s", o._statusCode, payload)
}

func (o *ApplicationComponentSnapshotGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationComponentSnapshotGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
