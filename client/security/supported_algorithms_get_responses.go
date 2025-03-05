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

	"github.com/sapcc/ontap-restapi/models"
)

// SupportedAlgorithmsGetReader is a Reader for the SupportedAlgorithmsGet structure.
type SupportedAlgorithmsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SupportedAlgorithmsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSupportedAlgorithmsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSupportedAlgorithmsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSupportedAlgorithmsGetOK creates a SupportedAlgorithmsGetOK with default headers values
func NewSupportedAlgorithmsGetOK() *SupportedAlgorithmsGetOK {
	return &SupportedAlgorithmsGetOK{}
}

/*
SupportedAlgorithmsGetOK describes a response with status code 200, with default header values.

OK
*/
type SupportedAlgorithmsGetOK struct {
	Payload *models.SupportedAlgorithms
}

// IsSuccess returns true when this supported algorithms get o k response has a 2xx status code
func (o *SupportedAlgorithmsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this supported algorithms get o k response has a 3xx status code
func (o *SupportedAlgorithmsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this supported algorithms get o k response has a 4xx status code
func (o *SupportedAlgorithmsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this supported algorithms get o k response has a 5xx status code
func (o *SupportedAlgorithmsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this supported algorithms get o k response a status code equal to that given
func (o *SupportedAlgorithmsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the supported algorithms get o k response
func (o *SupportedAlgorithmsGetOK) Code() int {
	return 200
}

func (o *SupportedAlgorithmsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/supported-algorithms/{owner.uuid}/{algorithm.name}][%d] supportedAlgorithmsGetOK %s", 200, payload)
}

func (o *SupportedAlgorithmsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/supported-algorithms/{owner.uuid}/{algorithm.name}][%d] supportedAlgorithmsGetOK %s", 200, payload)
}

func (o *SupportedAlgorithmsGetOK) GetPayload() *models.SupportedAlgorithms {
	return o.Payload
}

func (o *SupportedAlgorithmsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupportedAlgorithms)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSupportedAlgorithmsGetDefault creates a SupportedAlgorithmsGetDefault with default headers values
func NewSupportedAlgorithmsGetDefault(code int) *SupportedAlgorithmsGetDefault {
	return &SupportedAlgorithmsGetDefault{
		_statusCode: code,
	}
}

/*
SupportedAlgorithmsGetDefault describes a response with status code -1, with default header values.

Error
*/
type SupportedAlgorithmsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this supported algorithms get default response has a 2xx status code
func (o *SupportedAlgorithmsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this supported algorithms get default response has a 3xx status code
func (o *SupportedAlgorithmsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this supported algorithms get default response has a 4xx status code
func (o *SupportedAlgorithmsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this supported algorithms get default response has a 5xx status code
func (o *SupportedAlgorithmsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this supported algorithms get default response a status code equal to that given
func (o *SupportedAlgorithmsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the supported algorithms get default response
func (o *SupportedAlgorithmsGetDefault) Code() int {
	return o._statusCode
}

func (o *SupportedAlgorithmsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/supported-algorithms/{owner.uuid}/{algorithm.name}][%d] supported_algorithms_get default %s", o._statusCode, payload)
}

func (o *SupportedAlgorithmsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/webauthn/supported-algorithms/{owner.uuid}/{algorithm.name}][%d] supported_algorithms_get default %s", o._statusCode, payload)
}

func (o *SupportedAlgorithmsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SupportedAlgorithmsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
