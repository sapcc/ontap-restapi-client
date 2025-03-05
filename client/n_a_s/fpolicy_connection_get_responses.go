// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// FpolicyConnectionGetReader is a Reader for the FpolicyConnectionGet structure.
type FpolicyConnectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyConnectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyConnectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyConnectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyConnectionGetOK creates a FpolicyConnectionGetOK with default headers values
func NewFpolicyConnectionGetOK() *FpolicyConnectionGetOK {
	return &FpolicyConnectionGetOK{}
}

/*
FpolicyConnectionGetOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyConnectionGetOK struct {
	Payload *models.FpolicyConnection
}

// IsSuccess returns true when this fpolicy connection get o k response has a 2xx status code
func (o *FpolicyConnectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fpolicy connection get o k response has a 3xx status code
func (o *FpolicyConnectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fpolicy connection get o k response has a 4xx status code
func (o *FpolicyConnectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fpolicy connection get o k response has a 5xx status code
func (o *FpolicyConnectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fpolicy connection get o k response a status code equal to that given
func (o *FpolicyConnectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fpolicy connection get o k response
func (o *FpolicyConnectionGetOK) Code() int {
	return 200
}

func (o *FpolicyConnectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/connections/{node.uuid}/{policy.name}/{server}][%d] fpolicyConnectionGetOK %s", 200, payload)
}

func (o *FpolicyConnectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/connections/{node.uuid}/{policy.name}/{server}][%d] fpolicyConnectionGetOK %s", 200, payload)
}

func (o *FpolicyConnectionGetOK) GetPayload() *models.FpolicyConnection {
	return o.Payload
}

func (o *FpolicyConnectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FpolicyConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFpolicyConnectionGetDefault creates a FpolicyConnectionGetDefault with default headers values
func NewFpolicyConnectionGetDefault(code int) *FpolicyConnectionGetDefault {
	return &FpolicyConnectionGetDefault{
		_statusCode: code,
	}
}

/*
FpolicyConnectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type FpolicyConnectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fpolicy connection get default response has a 2xx status code
func (o *FpolicyConnectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fpolicy connection get default response has a 3xx status code
func (o *FpolicyConnectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fpolicy connection get default response has a 4xx status code
func (o *FpolicyConnectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fpolicy connection get default response has a 5xx status code
func (o *FpolicyConnectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fpolicy connection get default response a status code equal to that given
func (o *FpolicyConnectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fpolicy connection get default response
func (o *FpolicyConnectionGetDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyConnectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/connections/{node.uuid}/{policy.name}/{server}][%d] fpolicy_connection_get default %s", o._statusCode, payload)
}

func (o *FpolicyConnectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/connections/{node.uuid}/{policy.name}/{server}][%d] fpolicy_connection_get default %s", o._statusCode, payload)
}

func (o *FpolicyConnectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyConnectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
