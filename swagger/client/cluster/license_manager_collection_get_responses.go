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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// LicenseManagerCollectionGetReader is a Reader for the LicenseManagerCollectionGet structure.
type LicenseManagerCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LicenseManagerCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLicenseManagerCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLicenseManagerCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLicenseManagerCollectionGetOK creates a LicenseManagerCollectionGetOK with default headers values
func NewLicenseManagerCollectionGetOK() *LicenseManagerCollectionGetOK {
	return &LicenseManagerCollectionGetOK{}
}

/*
LicenseManagerCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type LicenseManagerCollectionGetOK struct {
	Payload *models.LicenseManagerResponse
}

// IsSuccess returns true when this license manager collection get o k response has a 2xx status code
func (o *LicenseManagerCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this license manager collection get o k response has a 3xx status code
func (o *LicenseManagerCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this license manager collection get o k response has a 4xx status code
func (o *LicenseManagerCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this license manager collection get o k response has a 5xx status code
func (o *LicenseManagerCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this license manager collection get o k response a status code equal to that given
func (o *LicenseManagerCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the license manager collection get o k response
func (o *LicenseManagerCollectionGetOK) Code() int {
	return 200
}

func (o *LicenseManagerCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/licensing/license-managers][%d] licenseManagerCollectionGetOK %s", 200, payload)
}

func (o *LicenseManagerCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/licensing/license-managers][%d] licenseManagerCollectionGetOK %s", 200, payload)
}

func (o *LicenseManagerCollectionGetOK) GetPayload() *models.LicenseManagerResponse {
	return o.Payload
}

func (o *LicenseManagerCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicenseManagerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLicenseManagerCollectionGetDefault creates a LicenseManagerCollectionGetDefault with default headers values
func NewLicenseManagerCollectionGetDefault(code int) *LicenseManagerCollectionGetDefault {
	return &LicenseManagerCollectionGetDefault{
		_statusCode: code,
	}
}

/*
LicenseManagerCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type LicenseManagerCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this license manager collection get default response has a 2xx status code
func (o *LicenseManagerCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this license manager collection get default response has a 3xx status code
func (o *LicenseManagerCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this license manager collection get default response has a 4xx status code
func (o *LicenseManagerCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this license manager collection get default response has a 5xx status code
func (o *LicenseManagerCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this license manager collection get default response a status code equal to that given
func (o *LicenseManagerCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the license manager collection get default response
func (o *LicenseManagerCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *LicenseManagerCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/licensing/license-managers][%d] license_manager_collection_get default %s", o._statusCode, payload)
}

func (o *LicenseManagerCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/licensing/license-managers][%d] license_manager_collection_get default %s", o._statusCode, payload)
}

func (o *LicenseManagerCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LicenseManagerCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
