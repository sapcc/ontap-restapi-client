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

// SupportedAlgorithmsCollectionGetReader is a Reader for the SupportedAlgorithmsCollectionGet structure.
type SupportedAlgorithmsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SupportedAlgorithmsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSupportedAlgorithmsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSupportedAlgorithmsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSupportedAlgorithmsCollectionGetOK creates a SupportedAlgorithmsCollectionGetOK with default headers values
func NewSupportedAlgorithmsCollectionGetOK() *SupportedAlgorithmsCollectionGetOK {
	return &SupportedAlgorithmsCollectionGetOK{}
}

/*
SupportedAlgorithmsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SupportedAlgorithmsCollectionGetOK struct {
	Payload *models.SupportedAlgorithmsResponse
}

// IsSuccess returns true when this supported algorithms collection get o k response has a 2xx status code
func (o *SupportedAlgorithmsCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this supported algorithms collection get o k response has a 3xx status code
func (o *SupportedAlgorithmsCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this supported algorithms collection get o k response has a 4xx status code
func (o *SupportedAlgorithmsCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this supported algorithms collection get o k response has a 5xx status code
func (o *SupportedAlgorithmsCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this supported algorithms collection get o k response a status code equal to that given
func (o *SupportedAlgorithmsCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the supported algorithms collection get o k response
func (o *SupportedAlgorithmsCollectionGetOK) Code() int {
	return 200
}

func (o *SupportedAlgorithmsCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/supported-algorithms][%d] supportedAlgorithmsCollectionGetOK %s", 200, payload)
}

func (o *SupportedAlgorithmsCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/supported-algorithms][%d] supportedAlgorithmsCollectionGetOK %s", 200, payload)
}

func (o *SupportedAlgorithmsCollectionGetOK) GetPayload() *models.SupportedAlgorithmsResponse {
	return o.Payload
}

func (o *SupportedAlgorithmsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupportedAlgorithmsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSupportedAlgorithmsCollectionGetDefault creates a SupportedAlgorithmsCollectionGetDefault with default headers values
func NewSupportedAlgorithmsCollectionGetDefault(code int) *SupportedAlgorithmsCollectionGetDefault {
	return &SupportedAlgorithmsCollectionGetDefault{
		_statusCode: code,
	}
}

/*
SupportedAlgorithmsCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SupportedAlgorithmsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this supported algorithms collection get default response has a 2xx status code
func (o *SupportedAlgorithmsCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this supported algorithms collection get default response has a 3xx status code
func (o *SupportedAlgorithmsCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this supported algorithms collection get default response has a 4xx status code
func (o *SupportedAlgorithmsCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this supported algorithms collection get default response has a 5xx status code
func (o *SupportedAlgorithmsCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this supported algorithms collection get default response a status code equal to that given
func (o *SupportedAlgorithmsCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the supported algorithms collection get default response
func (o *SupportedAlgorithmsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SupportedAlgorithmsCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/supported-algorithms][%d] supported_algorithms_collection_get default %s", o._statusCode, payload)
}

func (o *SupportedAlgorithmsCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/supported-algorithms][%d] supported_algorithms_collection_get default %s", o._statusCode, payload)
}

func (o *SupportedAlgorithmsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SupportedAlgorithmsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
