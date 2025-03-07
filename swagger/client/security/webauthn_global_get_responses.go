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

// WebauthnGlobalGetReader is a Reader for the WebauthnGlobalGet structure.
type WebauthnGlobalGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WebauthnGlobalGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWebauthnGlobalGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWebauthnGlobalGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWebauthnGlobalGetOK creates a WebauthnGlobalGetOK with default headers values
func NewWebauthnGlobalGetOK() *WebauthnGlobalGetOK {
	return &WebauthnGlobalGetOK{}
}

/*
WebauthnGlobalGetOK describes a response with status code 200, with default header values.

OK
*/
type WebauthnGlobalGetOK struct {
	Payload *models.WebauthnGlobal
}

// IsSuccess returns true when this webauthn global get o k response has a 2xx status code
func (o *WebauthnGlobalGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this webauthn global get o k response has a 3xx status code
func (o *WebauthnGlobalGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webauthn global get o k response has a 4xx status code
func (o *WebauthnGlobalGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this webauthn global get o k response has a 5xx status code
func (o *WebauthnGlobalGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this webauthn global get o k response a status code equal to that given
func (o *WebauthnGlobalGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the webauthn global get o k response
func (o *WebauthnGlobalGetOK) Code() int {
	return 200
}

func (o *WebauthnGlobalGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/global-settings/{owner.uuid}][%d] webauthnGlobalGetOK %s", 200, payload)
}

func (o *WebauthnGlobalGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/global-settings/{owner.uuid}][%d] webauthnGlobalGetOK %s", 200, payload)
}

func (o *WebauthnGlobalGetOK) GetPayload() *models.WebauthnGlobal {
	return o.Payload
}

func (o *WebauthnGlobalGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebauthnGlobal)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebauthnGlobalGetDefault creates a WebauthnGlobalGetDefault with default headers values
func NewWebauthnGlobalGetDefault(code int) *WebauthnGlobalGetDefault {
	return &WebauthnGlobalGetDefault{
		_statusCode: code,
	}
}

/*
WebauthnGlobalGetDefault describes a response with status code -1, with default header values.

Error
*/
type WebauthnGlobalGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this webauthn global get default response has a 2xx status code
func (o *WebauthnGlobalGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this webauthn global get default response has a 3xx status code
func (o *WebauthnGlobalGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this webauthn global get default response has a 4xx status code
func (o *WebauthnGlobalGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this webauthn global get default response has a 5xx status code
func (o *WebauthnGlobalGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this webauthn global get default response a status code equal to that given
func (o *WebauthnGlobalGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the webauthn global get default response
func (o *WebauthnGlobalGetDefault) Code() int {
	return o._statusCode
}

func (o *WebauthnGlobalGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/global-settings/{owner.uuid}][%d] webauthn_global_get default %s", o._statusCode, payload)
}

func (o *WebauthnGlobalGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/global-settings/{owner.uuid}][%d] webauthn_global_get default %s", o._statusCode, payload)
}

func (o *WebauthnGlobalGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *WebauthnGlobalGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
