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

// TotpCollectionGetReader is a Reader for the TotpCollectionGet structure.
type TotpCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TotpCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTotpCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTotpCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTotpCollectionGetOK creates a TotpCollectionGetOK with default headers values
func NewTotpCollectionGetOK() *TotpCollectionGetOK {
	return &TotpCollectionGetOK{}
}

/*
TotpCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type TotpCollectionGetOK struct {
	Payload *models.TotpResponse
}

// IsSuccess returns true when this totp collection get o k response has a 2xx status code
func (o *TotpCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this totp collection get o k response has a 3xx status code
func (o *TotpCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this totp collection get o k response has a 4xx status code
func (o *TotpCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this totp collection get o k response has a 5xx status code
func (o *TotpCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this totp collection get o k response a status code equal to that given
func (o *TotpCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the totp collection get o k response
func (o *TotpCollectionGetOK) Code() int {
	return 200
}

func (o *TotpCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/login/totps][%d] totpCollectionGetOK %s", 200, payload)
}

func (o *TotpCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/login/totps][%d] totpCollectionGetOK %s", 200, payload)
}

func (o *TotpCollectionGetOK) GetPayload() *models.TotpResponse {
	return o.Payload
}

func (o *TotpCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TotpResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTotpCollectionGetDefault creates a TotpCollectionGetDefault with default headers values
func NewTotpCollectionGetDefault(code int) *TotpCollectionGetDefault {
	return &TotpCollectionGetDefault{
		_statusCode: code,
	}
}

/*
TotpCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type TotpCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this totp collection get default response has a 2xx status code
func (o *TotpCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this totp collection get default response has a 3xx status code
func (o *TotpCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this totp collection get default response has a 4xx status code
func (o *TotpCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this totp collection get default response has a 5xx status code
func (o *TotpCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this totp collection get default response a status code equal to that given
func (o *TotpCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the totp collection get default response
func (o *TotpCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *TotpCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/login/totps][%d] totp_collection_get default %s", o._statusCode, payload)
}

func (o *TotpCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/login/totps][%d] totp_collection_get default %s", o._statusCode, payload)
}

func (o *TotpCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TotpCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
