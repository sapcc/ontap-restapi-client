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

	"github.com/sapcc/ontap-restapi/models"
)

// ApplicationComponentGetReader is a Reader for the ApplicationComponentGet structure.
type ApplicationComponentGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationComponentGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationComponentGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationComponentGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationComponentGetOK creates a ApplicationComponentGetOK with default headers values
func NewApplicationComponentGetOK() *ApplicationComponentGetOK {
	return &ApplicationComponentGetOK{}
}

/*
ApplicationComponentGetOK describes a response with status code 200, with default header values.

OK
*/
type ApplicationComponentGetOK struct {
	Payload *models.ApplicationComponent
}

// IsSuccess returns true when this application component get o k response has a 2xx status code
func (o *ApplicationComponentGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application component get o k response has a 3xx status code
func (o *ApplicationComponentGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application component get o k response has a 4xx status code
func (o *ApplicationComponentGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this application component get o k response has a 5xx status code
func (o *ApplicationComponentGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this application component get o k response a status code equal to that given
func (o *ApplicationComponentGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the application component get o k response
func (o *ApplicationComponentGetOK) Code() int {
	return 200
}

func (o *ApplicationComponentGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/components/{uuid}][%d] applicationComponentGetOK %s", 200, payload)
}

func (o *ApplicationComponentGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/components/{uuid}][%d] applicationComponentGetOK %s", 200, payload)
}

func (o *ApplicationComponentGetOK) GetPayload() *models.ApplicationComponent {
	return o.Payload
}

func (o *ApplicationComponentGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationComponent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationComponentGetDefault creates a ApplicationComponentGetDefault with default headers values
func NewApplicationComponentGetDefault(code int) *ApplicationComponentGetDefault {
	return &ApplicationComponentGetDefault{
		_statusCode: code,
	}
}

/*
ApplicationComponentGetDefault describes a response with status code -1, with default header values.

Error
*/
type ApplicationComponentGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this application component get default response has a 2xx status code
func (o *ApplicationComponentGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this application component get default response has a 3xx status code
func (o *ApplicationComponentGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this application component get default response has a 4xx status code
func (o *ApplicationComponentGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this application component get default response has a 5xx status code
func (o *ApplicationComponentGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this application component get default response a status code equal to that given
func (o *ApplicationComponentGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the application component get default response
func (o *ApplicationComponentGetDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationComponentGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/components/{uuid}][%d] application_component_get default %s", o._statusCode, payload)
}

func (o *ApplicationComponentGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/components/{uuid}][%d] application_component_get default %s", o._statusCode, payload)
}

func (o *ApplicationComponentGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationComponentGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
