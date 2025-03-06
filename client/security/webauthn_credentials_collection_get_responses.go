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

	"github.com/sapcc/ontap-restapi-client/models"
)

// WebauthnCredentialsCollectionGetReader is a Reader for the WebauthnCredentialsCollectionGet structure.
type WebauthnCredentialsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WebauthnCredentialsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWebauthnCredentialsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWebauthnCredentialsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWebauthnCredentialsCollectionGetOK creates a WebauthnCredentialsCollectionGetOK with default headers values
func NewWebauthnCredentialsCollectionGetOK() *WebauthnCredentialsCollectionGetOK {
	return &WebauthnCredentialsCollectionGetOK{}
}

/*
WebauthnCredentialsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type WebauthnCredentialsCollectionGetOK struct {
	Payload *models.WebauthnCredentialsResponse
}

// IsSuccess returns true when this webauthn credentials collection get o k response has a 2xx status code
func (o *WebauthnCredentialsCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this webauthn credentials collection get o k response has a 3xx status code
func (o *WebauthnCredentialsCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webauthn credentials collection get o k response has a 4xx status code
func (o *WebauthnCredentialsCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this webauthn credentials collection get o k response has a 5xx status code
func (o *WebauthnCredentialsCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this webauthn credentials collection get o k response a status code equal to that given
func (o *WebauthnCredentialsCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the webauthn credentials collection get o k response
func (o *WebauthnCredentialsCollectionGetOK) Code() int {
	return 200
}

func (o *WebauthnCredentialsCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/credentials][%d] webauthnCredentialsCollectionGetOK %s", 200, payload)
}

func (o *WebauthnCredentialsCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/credentials][%d] webauthnCredentialsCollectionGetOK %s", 200, payload)
}

func (o *WebauthnCredentialsCollectionGetOK) GetPayload() *models.WebauthnCredentialsResponse {
	return o.Payload
}

func (o *WebauthnCredentialsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebauthnCredentialsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebauthnCredentialsCollectionGetDefault creates a WebauthnCredentialsCollectionGetDefault with default headers values
func NewWebauthnCredentialsCollectionGetDefault(code int) *WebauthnCredentialsCollectionGetDefault {
	return &WebauthnCredentialsCollectionGetDefault{
		_statusCode: code,
	}
}

/*
WebauthnCredentialsCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type WebauthnCredentialsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this webauthn credentials collection get default response has a 2xx status code
func (o *WebauthnCredentialsCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this webauthn credentials collection get default response has a 3xx status code
func (o *WebauthnCredentialsCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this webauthn credentials collection get default response has a 4xx status code
func (o *WebauthnCredentialsCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this webauthn credentials collection get default response has a 5xx status code
func (o *WebauthnCredentialsCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this webauthn credentials collection get default response a status code equal to that given
func (o *WebauthnCredentialsCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the webauthn credentials collection get default response
func (o *WebauthnCredentialsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *WebauthnCredentialsCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/credentials][%d] webauthn_credentials_collection_get default %s", o._statusCode, payload)
}

func (o *WebauthnCredentialsCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/credentials][%d] webauthn_credentials_collection_get default %s", o._statusCode, payload)
}

func (o *WebauthnCredentialsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *WebauthnCredentialsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
