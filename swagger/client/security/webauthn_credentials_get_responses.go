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

// WebauthnCredentialsGetReader is a Reader for the WebauthnCredentialsGet structure.
type WebauthnCredentialsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WebauthnCredentialsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWebauthnCredentialsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWebauthnCredentialsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWebauthnCredentialsGetOK creates a WebauthnCredentialsGetOK with default headers values
func NewWebauthnCredentialsGetOK() *WebauthnCredentialsGetOK {
	return &WebauthnCredentialsGetOK{}
}

/*
WebauthnCredentialsGetOK describes a response with status code 200, with default header values.

OK
*/
type WebauthnCredentialsGetOK struct {
	Payload *models.WebauthnCredentials
}

// IsSuccess returns true when this webauthn credentials get o k response has a 2xx status code
func (o *WebauthnCredentialsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this webauthn credentials get o k response has a 3xx status code
func (o *WebauthnCredentialsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webauthn credentials get o k response has a 4xx status code
func (o *WebauthnCredentialsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this webauthn credentials get o k response has a 5xx status code
func (o *WebauthnCredentialsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this webauthn credentials get o k response a status code equal to that given
func (o *WebauthnCredentialsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the webauthn credentials get o k response
func (o *WebauthnCredentialsGetOK) Code() int {
	return 200
}

func (o *WebauthnCredentialsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/credentials/{owner.uuid}/{username}/{index}/{relying_party.id}][%d] webauthnCredentialsGetOK %s", 200, payload)
}

func (o *WebauthnCredentialsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/credentials/{owner.uuid}/{username}/{index}/{relying_party.id}][%d] webauthnCredentialsGetOK %s", 200, payload)
}

func (o *WebauthnCredentialsGetOK) GetPayload() *models.WebauthnCredentials {
	return o.Payload
}

func (o *WebauthnCredentialsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebauthnCredentials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebauthnCredentialsGetDefault creates a WebauthnCredentialsGetDefault with default headers values
func NewWebauthnCredentialsGetDefault(code int) *WebauthnCredentialsGetDefault {
	return &WebauthnCredentialsGetDefault{
		_statusCode: code,
	}
}

/*
WebauthnCredentialsGetDefault describes a response with status code -1, with default header values.

Error
*/
type WebauthnCredentialsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this webauthn credentials get default response has a 2xx status code
func (o *WebauthnCredentialsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this webauthn credentials get default response has a 3xx status code
func (o *WebauthnCredentialsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this webauthn credentials get default response has a 4xx status code
func (o *WebauthnCredentialsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this webauthn credentials get default response has a 5xx status code
func (o *WebauthnCredentialsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this webauthn credentials get default response a status code equal to that given
func (o *WebauthnCredentialsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the webauthn credentials get default response
func (o *WebauthnCredentialsGetDefault) Code() int {
	return o._statusCode
}

func (o *WebauthnCredentialsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/credentials/{owner.uuid}/{username}/{index}/{relying_party.id}][%d] webauthn_credentials_get default %s", o._statusCode, payload)
}

func (o *WebauthnCredentialsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/credentials/{owner.uuid}/{username}/{index}/{relying_party.id}][%d] webauthn_credentials_get default %s", o._statusCode, payload)
}

func (o *WebauthnCredentialsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *WebauthnCredentialsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
