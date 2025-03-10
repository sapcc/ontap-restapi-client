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

// WebauthnGlobalCollectionGetReader is a Reader for the WebauthnGlobalCollectionGet structure.
type WebauthnGlobalCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WebauthnGlobalCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWebauthnGlobalCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWebauthnGlobalCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWebauthnGlobalCollectionGetOK creates a WebauthnGlobalCollectionGetOK with default headers values
func NewWebauthnGlobalCollectionGetOK() *WebauthnGlobalCollectionGetOK {
	return &WebauthnGlobalCollectionGetOK{}
}

/*
WebauthnGlobalCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type WebauthnGlobalCollectionGetOK struct {
	Payload *models.WebauthnGlobalResponse
}

// IsSuccess returns true when this webauthn global collection get o k response has a 2xx status code
func (o *WebauthnGlobalCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this webauthn global collection get o k response has a 3xx status code
func (o *WebauthnGlobalCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webauthn global collection get o k response has a 4xx status code
func (o *WebauthnGlobalCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this webauthn global collection get o k response has a 5xx status code
func (o *WebauthnGlobalCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this webauthn global collection get o k response a status code equal to that given
func (o *WebauthnGlobalCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the webauthn global collection get o k response
func (o *WebauthnGlobalCollectionGetOK) Code() int {
	return 200
}

func (o *WebauthnGlobalCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/global-settings][%d] webauthnGlobalCollectionGetOK %s", 200, payload)
}

func (o *WebauthnGlobalCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/global-settings][%d] webauthnGlobalCollectionGetOK %s", 200, payload)
}

func (o *WebauthnGlobalCollectionGetOK) GetPayload() *models.WebauthnGlobalResponse {
	return o.Payload
}

func (o *WebauthnGlobalCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebauthnGlobalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebauthnGlobalCollectionGetDefault creates a WebauthnGlobalCollectionGetDefault with default headers values
func NewWebauthnGlobalCollectionGetDefault(code int) *WebauthnGlobalCollectionGetDefault {
	return &WebauthnGlobalCollectionGetDefault{
		_statusCode: code,
	}
}

/*
WebauthnGlobalCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type WebauthnGlobalCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this webauthn global collection get default response has a 2xx status code
func (o *WebauthnGlobalCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this webauthn global collection get default response has a 3xx status code
func (o *WebauthnGlobalCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this webauthn global collection get default response has a 4xx status code
func (o *WebauthnGlobalCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this webauthn global collection get default response has a 5xx status code
func (o *WebauthnGlobalCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this webauthn global collection get default response a status code equal to that given
func (o *WebauthnGlobalCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the webauthn global collection get default response
func (o *WebauthnGlobalCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *WebauthnGlobalCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/global-settings][%d] webauthn_global_collection_get default %s", o._statusCode, payload)
}

func (o *WebauthnGlobalCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/global-settings][%d] webauthn_global_collection_get default %s", o._statusCode, payload)
}

func (o *WebauthnGlobalCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *WebauthnGlobalCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
