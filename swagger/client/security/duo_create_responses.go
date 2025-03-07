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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// DuoCreateReader is a Reader for the DuoCreate structure.
type DuoCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DuoCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDuoCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDuoCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDuoCreateCreated creates a DuoCreateCreated with default headers values
func NewDuoCreateCreated() *DuoCreateCreated {
	return &DuoCreateCreated{}
}

/*
DuoCreateCreated describes a response with status code 201, with default header values.

Created
*/
type DuoCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.DuoResponse
}

// IsSuccess returns true when this duo create created response has a 2xx status code
func (o *DuoCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this duo create created response has a 3xx status code
func (o *DuoCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this duo create created response has a 4xx status code
func (o *DuoCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this duo create created response has a 5xx status code
func (o *DuoCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this duo create created response a status code equal to that given
func (o *DuoCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the duo create created response
func (o *DuoCreateCreated) Code() int {
	return 201
}

func (o *DuoCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/duo/profiles][%d] duoCreateCreated %s", 201, payload)
}

func (o *DuoCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/duo/profiles][%d] duoCreateCreated %s", 201, payload)
}

func (o *DuoCreateCreated) GetPayload() *models.DuoResponse {
	return o.Payload
}

func (o *DuoCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.DuoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDuoCreateDefault creates a DuoCreateDefault with default headers values
func NewDuoCreateDefault(code int) *DuoCreateDefault {
	return &DuoCreateDefault{
		_statusCode: code,
	}
}

/*
DuoCreateDefault describes a response with status code -1, with default header values.

Error
*/
type DuoCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this duo create default response has a 2xx status code
func (o *DuoCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this duo create default response has a 3xx status code
func (o *DuoCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this duo create default response has a 4xx status code
func (o *DuoCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this duo create default response has a 5xx status code
func (o *DuoCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this duo create default response a status code equal to that given
func (o *DuoCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the duo create default response
func (o *DuoCreateDefault) Code() int {
	return o._statusCode
}

func (o *DuoCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/duo/profiles][%d] duo_create default %s", o._statusCode, payload)
}

func (o *DuoCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/duo/profiles][%d] duo_create default %s", o._statusCode, payload)
}

func (o *DuoCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DuoCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
