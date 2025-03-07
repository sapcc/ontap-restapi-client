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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// AutoUpdateStatusGetReader is a Reader for the AutoUpdateStatusGet structure.
type AutoUpdateStatusGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutoUpdateStatusGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAutoUpdateStatusGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAutoUpdateStatusGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAutoUpdateStatusGetOK creates a AutoUpdateStatusGetOK with default headers values
func NewAutoUpdateStatusGetOK() *AutoUpdateStatusGetOK {
	return &AutoUpdateStatusGetOK{}
}

/*
AutoUpdateStatusGetOK describes a response with status code 200, with default header values.

OK
*/
type AutoUpdateStatusGetOK struct {
	Payload *models.AutoUpdateStatus
}

// IsSuccess returns true when this auto update status get o k response has a 2xx status code
func (o *AutoUpdateStatusGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this auto update status get o k response has a 3xx status code
func (o *AutoUpdateStatusGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auto update status get o k response has a 4xx status code
func (o *AutoUpdateStatusGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this auto update status get o k response has a 5xx status code
func (o *AutoUpdateStatusGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this auto update status get o k response a status code equal to that given
func (o *AutoUpdateStatusGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the auto update status get o k response
func (o *AutoUpdateStatusGetOK) Code() int {
	return 200
}

func (o *AutoUpdateStatusGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/auto-update/updates/{uuid}][%d] autoUpdateStatusGetOK %s", 200, payload)
}

func (o *AutoUpdateStatusGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/auto-update/updates/{uuid}][%d] autoUpdateStatusGetOK %s", 200, payload)
}

func (o *AutoUpdateStatusGetOK) GetPayload() *models.AutoUpdateStatus {
	return o.Payload
}

func (o *AutoUpdateStatusGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AutoUpdateStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoUpdateStatusGetDefault creates a AutoUpdateStatusGetDefault with default headers values
func NewAutoUpdateStatusGetDefault(code int) *AutoUpdateStatusGetDefault {
	return &AutoUpdateStatusGetDefault{
		_statusCode: code,
	}
}

/*
AutoUpdateStatusGetDefault describes a response with status code -1, with default header values.

Error
*/
type AutoUpdateStatusGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this auto update status get default response has a 2xx status code
func (o *AutoUpdateStatusGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this auto update status get default response has a 3xx status code
func (o *AutoUpdateStatusGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this auto update status get default response has a 4xx status code
func (o *AutoUpdateStatusGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this auto update status get default response has a 5xx status code
func (o *AutoUpdateStatusGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this auto update status get default response a status code equal to that given
func (o *AutoUpdateStatusGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the auto update status get default response
func (o *AutoUpdateStatusGetDefault) Code() int {
	return o._statusCode
}

func (o *AutoUpdateStatusGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/auto-update/updates/{uuid}][%d] auto_update_status_get default %s", o._statusCode, payload)
}

func (o *AutoUpdateStatusGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/auto-update/updates/{uuid}][%d] auto_update_status_get default %s", o._statusCode, payload)
}

func (o *AutoUpdateStatusGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AutoUpdateStatusGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
