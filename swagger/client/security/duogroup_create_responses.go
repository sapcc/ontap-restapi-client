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

// DuogroupCreateReader is a Reader for the DuogroupCreate structure.
type DuogroupCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DuogroupCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDuogroupCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDuogroupCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDuogroupCreateCreated creates a DuogroupCreateCreated with default headers values
func NewDuogroupCreateCreated() *DuogroupCreateCreated {
	return &DuogroupCreateCreated{}
}

/*
DuogroupCreateCreated describes a response with status code 201, with default header values.

Created
*/
type DuogroupCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.DuogroupResponse
}

// IsSuccess returns true when this duogroup create created response has a 2xx status code
func (o *DuogroupCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this duogroup create created response has a 3xx status code
func (o *DuogroupCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this duogroup create created response has a 4xx status code
func (o *DuogroupCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this duogroup create created response has a 5xx status code
func (o *DuogroupCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this duogroup create created response a status code equal to that given
func (o *DuogroupCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the duogroup create created response
func (o *DuogroupCreateCreated) Code() int {
	return 201
}

func (o *DuogroupCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/duo/groups][%d] duogroupCreateCreated %s", 201, payload)
}

func (o *DuogroupCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/duo/groups][%d] duogroupCreateCreated %s", 201, payload)
}

func (o *DuogroupCreateCreated) GetPayload() *models.DuogroupResponse {
	return o.Payload
}

func (o *DuogroupCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.DuogroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDuogroupCreateDefault creates a DuogroupCreateDefault with default headers values
func NewDuogroupCreateDefault(code int) *DuogroupCreateDefault {
	return &DuogroupCreateDefault{
		_statusCode: code,
	}
}

/*
DuogroupCreateDefault describes a response with status code -1, with default header values.

Error
*/
type DuogroupCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this duogroup create default response has a 2xx status code
func (o *DuogroupCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this duogroup create default response has a 3xx status code
func (o *DuogroupCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this duogroup create default response has a 4xx status code
func (o *DuogroupCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this duogroup create default response has a 5xx status code
func (o *DuogroupCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this duogroup create default response a status code equal to that given
func (o *DuogroupCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the duogroup create default response
func (o *DuogroupCreateDefault) Code() int {
	return o._statusCode
}

func (o *DuogroupCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/duo/groups][%d] duogroup_create default %s", o._statusCode, payload)
}

func (o *DuogroupCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/duo/groups][%d] duogroup_create default %s", o._statusCode, payload)
}

func (o *DuogroupCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DuogroupCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
