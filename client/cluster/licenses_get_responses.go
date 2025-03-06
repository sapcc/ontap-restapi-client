// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// LicensesGetReader is a Reader for the LicensesGet structure.
type LicensesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LicensesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLicensesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLicensesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLicensesGetOK creates a LicensesGetOK with default headers values
func NewLicensesGetOK() *LicensesGetOK {
	return &LicensesGetOK{}
}

/*
LicensesGetOK describes a response with status code 200, with default header values.

OK
*/
type LicensesGetOK struct {
	Payload *models.LicensePackageResponse
}

// IsSuccess returns true when this licenses get o k response has a 2xx status code
func (o *LicensesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this licenses get o k response has a 3xx status code
func (o *LicensesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this licenses get o k response has a 4xx status code
func (o *LicensesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this licenses get o k response has a 5xx status code
func (o *LicensesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this licenses get o k response a status code equal to that given
func (o *LicensesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the licenses get o k response
func (o *LicensesGetOK) Code() int {
	return 200
}

func (o *LicensesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/licensing/licenses][%d] licensesGetOK %s", 200, payload)
}

func (o *LicensesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/licensing/licenses][%d] licensesGetOK %s", 200, payload)
}

func (o *LicensesGetOK) GetPayload() *models.LicensePackageResponse {
	return o.Payload
}

func (o *LicensesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicensePackageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLicensesGetDefault creates a LicensesGetDefault with default headers values
func NewLicensesGetDefault(code int) *LicensesGetDefault {
	return &LicensesGetDefault{
		_statusCode: code,
	}
}

/*
LicensesGetDefault describes a response with status code -1, with default header values.

Error
*/
type LicensesGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this licenses get default response has a 2xx status code
func (o *LicensesGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this licenses get default response has a 3xx status code
func (o *LicensesGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this licenses get default response has a 4xx status code
func (o *LicensesGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this licenses get default response has a 5xx status code
func (o *LicensesGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this licenses get default response a status code equal to that given
func (o *LicensesGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the licenses get default response
func (o *LicensesGetDefault) Code() int {
	return o._statusCode
}

func (o *LicensesGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/licensing/licenses][%d] licenses_get default %s", o._statusCode, payload)
}

func (o *LicensesGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/licensing/licenses][%d] licenses_get default %s", o._statusCode, payload)
}

func (o *LicensesGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LicensesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
