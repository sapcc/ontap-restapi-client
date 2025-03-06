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

// AutoUpdateInfoGetReader is a Reader for the AutoUpdateInfoGet structure.
type AutoUpdateInfoGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutoUpdateInfoGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAutoUpdateInfoGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAutoUpdateInfoGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAutoUpdateInfoGetOK creates a AutoUpdateInfoGetOK with default headers values
func NewAutoUpdateInfoGetOK() *AutoUpdateInfoGetOK {
	return &AutoUpdateInfoGetOK{}
}

/*
AutoUpdateInfoGetOK describes a response with status code 200, with default header values.

OK
*/
type AutoUpdateInfoGetOK struct {
	Payload *models.AutoUpdateInfo
}

// IsSuccess returns true when this auto update info get o k response has a 2xx status code
func (o *AutoUpdateInfoGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this auto update info get o k response has a 3xx status code
func (o *AutoUpdateInfoGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auto update info get o k response has a 4xx status code
func (o *AutoUpdateInfoGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this auto update info get o k response has a 5xx status code
func (o *AutoUpdateInfoGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this auto update info get o k response a status code equal to that given
func (o *AutoUpdateInfoGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the auto update info get o k response
func (o *AutoUpdateInfoGetOK) Code() int {
	return 200
}

func (o *AutoUpdateInfoGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/auto-update][%d] autoUpdateInfoGetOK %s", 200, payload)
}

func (o *AutoUpdateInfoGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/auto-update][%d] autoUpdateInfoGetOK %s", 200, payload)
}

func (o *AutoUpdateInfoGetOK) GetPayload() *models.AutoUpdateInfo {
	return o.Payload
}

func (o *AutoUpdateInfoGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AutoUpdateInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoUpdateInfoGetDefault creates a AutoUpdateInfoGetDefault with default headers values
func NewAutoUpdateInfoGetDefault(code int) *AutoUpdateInfoGetDefault {
	return &AutoUpdateInfoGetDefault{
		_statusCode: code,
	}
}

/*
AutoUpdateInfoGetDefault describes a response with status code -1, with default header values.

Error
*/
type AutoUpdateInfoGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this auto update info get default response has a 2xx status code
func (o *AutoUpdateInfoGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this auto update info get default response has a 3xx status code
func (o *AutoUpdateInfoGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this auto update info get default response has a 4xx status code
func (o *AutoUpdateInfoGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this auto update info get default response has a 5xx status code
func (o *AutoUpdateInfoGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this auto update info get default response a status code equal to that given
func (o *AutoUpdateInfoGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the auto update info get default response
func (o *AutoUpdateInfoGetDefault) Code() int {
	return o._statusCode
}

func (o *AutoUpdateInfoGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/auto-update][%d] auto_update_info_get default %s", o._statusCode, payload)
}

func (o *AutoUpdateInfoGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/auto-update][%d] auto_update_info_get default %s", o._statusCode, payload)
}

func (o *AutoUpdateInfoGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AutoUpdateInfoGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
