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

// ApplicationTemplateCollectionGetReader is a Reader for the ApplicationTemplateCollectionGet structure.
type ApplicationTemplateCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationTemplateCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationTemplateCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationTemplateCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationTemplateCollectionGetOK creates a ApplicationTemplateCollectionGetOK with default headers values
func NewApplicationTemplateCollectionGetOK() *ApplicationTemplateCollectionGetOK {
	return &ApplicationTemplateCollectionGetOK{}
}

/*
ApplicationTemplateCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ApplicationTemplateCollectionGetOK struct {
	Payload *models.ApplicationTemplateResponse
}

// IsSuccess returns true when this application template collection get o k response has a 2xx status code
func (o *ApplicationTemplateCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application template collection get o k response has a 3xx status code
func (o *ApplicationTemplateCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application template collection get o k response has a 4xx status code
func (o *ApplicationTemplateCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this application template collection get o k response has a 5xx status code
func (o *ApplicationTemplateCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this application template collection get o k response a status code equal to that given
func (o *ApplicationTemplateCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the application template collection get o k response
func (o *ApplicationTemplateCollectionGetOK) Code() int {
	return 200
}

func (o *ApplicationTemplateCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/templates][%d] applicationTemplateCollectionGetOK %s", 200, payload)
}

func (o *ApplicationTemplateCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/templates][%d] applicationTemplateCollectionGetOK %s", 200, payload)
}

func (o *ApplicationTemplateCollectionGetOK) GetPayload() *models.ApplicationTemplateResponse {
	return o.Payload
}

func (o *ApplicationTemplateCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationTemplateCollectionGetDefault creates a ApplicationTemplateCollectionGetDefault with default headers values
func NewApplicationTemplateCollectionGetDefault(code int) *ApplicationTemplateCollectionGetDefault {
	return &ApplicationTemplateCollectionGetDefault{
		_statusCode: code,
	}
}

/*
ApplicationTemplateCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type ApplicationTemplateCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this application template collection get default response has a 2xx status code
func (o *ApplicationTemplateCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this application template collection get default response has a 3xx status code
func (o *ApplicationTemplateCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this application template collection get default response has a 4xx status code
func (o *ApplicationTemplateCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this application template collection get default response has a 5xx status code
func (o *ApplicationTemplateCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this application template collection get default response a status code equal to that given
func (o *ApplicationTemplateCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the application template collection get default response
func (o *ApplicationTemplateCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationTemplateCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/templates][%d] application_template_collection_get default %s", o._statusCode, payload)
}

func (o *ApplicationTemplateCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/templates][%d] application_template_collection_get default %s", o._statusCode, payload)
}

func (o *ApplicationTemplateCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationTemplateCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
