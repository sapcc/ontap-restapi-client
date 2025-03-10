// Code generated by go-swagger; DO NOT EDIT.

package snap_lock

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

// SnaplockLegalHoldGetReader is a Reader for the SnaplockLegalHoldGet structure.
type SnaplockLegalHoldGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockLegalHoldGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockLegalHoldGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockLegalHoldGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockLegalHoldGetOK creates a SnaplockLegalHoldGetOK with default headers values
func NewSnaplockLegalHoldGetOK() *SnaplockLegalHoldGetOK {
	return &SnaplockLegalHoldGetOK{}
}

/*
SnaplockLegalHoldGetOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockLegalHoldGetOK struct {
	Payload *models.SnaplockLegalHoldOperation
}

// IsSuccess returns true when this snaplock legal hold get o k response has a 2xx status code
func (o *SnaplockLegalHoldGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock legal hold get o k response has a 3xx status code
func (o *SnaplockLegalHoldGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock legal hold get o k response has a 4xx status code
func (o *SnaplockLegalHoldGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock legal hold get o k response has a 5xx status code
func (o *SnaplockLegalHoldGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock legal hold get o k response a status code equal to that given
func (o *SnaplockLegalHoldGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snaplock legal hold get o k response
func (o *SnaplockLegalHoldGetOK) Code() int {
	return 200
}

func (o *SnaplockLegalHoldGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/litigations/{litigation.id}/operations/{id}][%d] snaplockLegalHoldGetOK %s", 200, payload)
}

func (o *SnaplockLegalHoldGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/litigations/{litigation.id}/operations/{id}][%d] snaplockLegalHoldGetOK %s", 200, payload)
}

func (o *SnaplockLegalHoldGetOK) GetPayload() *models.SnaplockLegalHoldOperation {
	return o.Payload
}

func (o *SnaplockLegalHoldGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockLegalHoldOperation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLegalHoldGetDefault creates a SnaplockLegalHoldGetDefault with default headers values
func NewSnaplockLegalHoldGetDefault(code int) *SnaplockLegalHoldGetDefault {
	return &SnaplockLegalHoldGetDefault{
		_statusCode: code,
	}
}

/*
	SnaplockLegalHoldGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 13763280    | Only a user with security login role \"vsadmin-snaplock\" is allowed to perform this operation.  |
| 14090343    | Invalid Field  |
| 14090346    | Internal Error. Wait a few minutes, then try the command again  |
*/
type SnaplockLegalHoldGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snaplock legal hold get default response has a 2xx status code
func (o *SnaplockLegalHoldGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock legal hold get default response has a 3xx status code
func (o *SnaplockLegalHoldGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock legal hold get default response has a 4xx status code
func (o *SnaplockLegalHoldGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock legal hold get default response has a 5xx status code
func (o *SnaplockLegalHoldGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock legal hold get default response a status code equal to that given
func (o *SnaplockLegalHoldGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snaplock legal hold get default response
func (o *SnaplockLegalHoldGetDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockLegalHoldGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/litigations/{litigation.id}/operations/{id}][%d] snaplock_legal_hold_get default %s", o._statusCode, payload)
}

func (o *SnaplockLegalHoldGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/litigations/{litigation.id}/operations/{id}][%d] snaplock_legal_hold_get default %s", o._statusCode, payload)
}

func (o *SnaplockLegalHoldGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockLegalHoldGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
