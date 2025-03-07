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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// ActiveDirectoryGetReader is a Reader for the ActiveDirectoryGet structure.
type ActiveDirectoryGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActiveDirectoryGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActiveDirectoryGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewActiveDirectoryGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActiveDirectoryGetOK creates a ActiveDirectoryGetOK with default headers values
func NewActiveDirectoryGetOK() *ActiveDirectoryGetOK {
	return &ActiveDirectoryGetOK{}
}

/*
ActiveDirectoryGetOK describes a response with status code 200, with default header values.

OK
*/
type ActiveDirectoryGetOK struct {
	Payload *models.ActiveDirectory
}

// IsSuccess returns true when this active directory get o k response has a 2xx status code
func (o *ActiveDirectoryGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this active directory get o k response has a 3xx status code
func (o *ActiveDirectoryGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this active directory get o k response has a 4xx status code
func (o *ActiveDirectoryGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this active directory get o k response has a 5xx status code
func (o *ActiveDirectoryGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this active directory get o k response a status code equal to that given
func (o *ActiveDirectoryGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the active directory get o k response
func (o *ActiveDirectoryGetOK) Code() int {
	return 200
}

func (o *ActiveDirectoryGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/active-directory/{svm.uuid}][%d] activeDirectoryGetOK %s", 200, payload)
}

func (o *ActiveDirectoryGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/active-directory/{svm.uuid}][%d] activeDirectoryGetOK %s", 200, payload)
}

func (o *ActiveDirectoryGetOK) GetPayload() *models.ActiveDirectory {
	return o.Payload
}

func (o *ActiveDirectoryGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActiveDirectory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActiveDirectoryGetDefault creates a ActiveDirectoryGetDefault with default headers values
func NewActiveDirectoryGetDefault(code int) *ActiveDirectoryGetDefault {
	return &ActiveDirectoryGetDefault{
		_statusCode: code,
	}
}

/*
ActiveDirectoryGetDefault describes a response with status code -1, with default header values.

Error
*/
type ActiveDirectoryGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this active directory get default response has a 2xx status code
func (o *ActiveDirectoryGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this active directory get default response has a 3xx status code
func (o *ActiveDirectoryGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this active directory get default response has a 4xx status code
func (o *ActiveDirectoryGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this active directory get default response has a 5xx status code
func (o *ActiveDirectoryGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this active directory get default response a status code equal to that given
func (o *ActiveDirectoryGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the active directory get default response
func (o *ActiveDirectoryGetDefault) Code() int {
	return o._statusCode
}

func (o *ActiveDirectoryGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/active-directory/{svm.uuid}][%d] active_directory_get default %s", o._statusCode, payload)
}

func (o *ActiveDirectoryGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/active-directory/{svm.uuid}][%d] active_directory_get default %s", o._statusCode, payload)
}

func (o *ActiveDirectoryGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ActiveDirectoryGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
