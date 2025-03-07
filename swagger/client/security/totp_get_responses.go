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

// TotpGetReader is a Reader for the TotpGet structure.
type TotpGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TotpGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTotpGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTotpGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTotpGetOK creates a TotpGetOK with default headers values
func NewTotpGetOK() *TotpGetOK {
	return &TotpGetOK{}
}

/*
TotpGetOK describes a response with status code 200, with default header values.

OK
*/
type TotpGetOK struct {
	Payload *models.Totp
}

// IsSuccess returns true when this totp get o k response has a 2xx status code
func (o *TotpGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this totp get o k response has a 3xx status code
func (o *TotpGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this totp get o k response has a 4xx status code
func (o *TotpGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this totp get o k response has a 5xx status code
func (o *TotpGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this totp get o k response a status code equal to that given
func (o *TotpGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the totp get o k response
func (o *TotpGetOK) Code() int {
	return 200
}

func (o *TotpGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/login/totps/{owner.uuid}/{account.name}][%d] totpGetOK %s", 200, payload)
}

func (o *TotpGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/login/totps/{owner.uuid}/{account.name}][%d] totpGetOK %s", 200, payload)
}

func (o *TotpGetOK) GetPayload() *models.Totp {
	return o.Payload
}

func (o *TotpGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Totp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTotpGetDefault creates a TotpGetDefault with default headers values
func NewTotpGetDefault(code int) *TotpGetDefault {
	return &TotpGetDefault{
		_statusCode: code,
	}
}

/*
TotpGetDefault describes a response with status code -1, with default header values.

Error
*/
type TotpGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this totp get default response has a 2xx status code
func (o *TotpGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this totp get default response has a 3xx status code
func (o *TotpGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this totp get default response has a 4xx status code
func (o *TotpGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this totp get default response has a 5xx status code
func (o *TotpGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this totp get default response a status code equal to that given
func (o *TotpGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the totp get default response
func (o *TotpGetDefault) Code() int {
	return o._statusCode
}

func (o *TotpGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/login/totps/{owner.uuid}/{account.name}][%d] totp_get default %s", o._statusCode, payload)
}

func (o *TotpGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/login/totps/{owner.uuid}/{account.name}][%d] totp_get default %s", o._statusCode, payload)
}

func (o *TotpGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TotpGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
