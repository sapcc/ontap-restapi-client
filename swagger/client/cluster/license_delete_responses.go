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

// LicenseDeleteReader is a Reader for the LicenseDelete structure.
type LicenseDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LicenseDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLicenseDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLicenseDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLicenseDeleteOK creates a LicenseDeleteOK with default headers values
func NewLicenseDeleteOK() *LicenseDeleteOK {
	return &LicenseDeleteOK{}
}

/*
LicenseDeleteOK describes a response with status code 200, with default header values.

OK
*/
type LicenseDeleteOK struct {
}

// IsSuccess returns true when this license delete o k response has a 2xx status code
func (o *LicenseDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this license delete o k response has a 3xx status code
func (o *LicenseDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this license delete o k response has a 4xx status code
func (o *LicenseDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this license delete o k response has a 5xx status code
func (o *LicenseDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this license delete o k response a status code equal to that given
func (o *LicenseDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the license delete o k response
func (o *LicenseDeleteOK) Code() int {
	return 200
}

func (o *LicenseDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /cluster/licensing/licenses/{name}][%d] licenseDeleteOK", 200)
}

func (o *LicenseDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /cluster/licensing/licenses/{name}][%d] licenseDeleteOK", 200)
}

func (o *LicenseDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLicenseDeleteDefault creates a LicenseDeleteDefault with default headers values
func NewLicenseDeleteDefault(code int) *LicenseDeleteDefault {
	return &LicenseDeleteDefault{
		_statusCode: code,
	}
}

/*
	LicenseDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 525028 | Error during volume limit check, cannot remove license |
| 525029 | Current volume use will exceed limits if license is removed |
| 1115137 | Cluster license requires a base license to be installed |
| 1115144 | Cloud licenses cannot be deleted |
| 1115178 | A tier license that is still in use cannot be deleted |
| 1115213 | License is still in use and cannot be removed |
| 1115406 | Capacity pool licenses cannot be deleted |
| 1115564 | Package is part of a NLFv2 license and cannot be removed individually |
| 66846823 | A FlexCache license that is still in use cannot be deleted |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type LicenseDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this license delete default response has a 2xx status code
func (o *LicenseDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this license delete default response has a 3xx status code
func (o *LicenseDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this license delete default response has a 4xx status code
func (o *LicenseDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this license delete default response has a 5xx status code
func (o *LicenseDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this license delete default response a status code equal to that given
func (o *LicenseDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the license delete default response
func (o *LicenseDeleteDefault) Code() int {
	return o._statusCode
}

func (o *LicenseDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/licensing/licenses/{name}][%d] license_delete default %s", o._statusCode, payload)
}

func (o *LicenseDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/licensing/licenses/{name}][%d] license_delete default %s", o._statusCode, payload)
}

func (o *LicenseDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LicenseDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
