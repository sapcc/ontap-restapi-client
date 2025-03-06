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

	"github.com/sapcc/ontap-restapi-client/models"
)

// SnaplockLegalHoldInstanceGetReader is a Reader for the SnaplockLegalHoldInstanceGet structure.
type SnaplockLegalHoldInstanceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockLegalHoldInstanceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockLegalHoldInstanceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockLegalHoldInstanceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockLegalHoldInstanceGetOK creates a SnaplockLegalHoldInstanceGetOK with default headers values
func NewSnaplockLegalHoldInstanceGetOK() *SnaplockLegalHoldInstanceGetOK {
	return &SnaplockLegalHoldInstanceGetOK{}
}

/*
SnaplockLegalHoldInstanceGetOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockLegalHoldInstanceGetOK struct {
	Payload *models.SnaplockLitigation
}

// IsSuccess returns true when this snaplock legal hold instance get o k response has a 2xx status code
func (o *SnaplockLegalHoldInstanceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock legal hold instance get o k response has a 3xx status code
func (o *SnaplockLegalHoldInstanceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock legal hold instance get o k response has a 4xx status code
func (o *SnaplockLegalHoldInstanceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock legal hold instance get o k response has a 5xx status code
func (o *SnaplockLegalHoldInstanceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock legal hold instance get o k response a status code equal to that given
func (o *SnaplockLegalHoldInstanceGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snaplock legal hold instance get o k response
func (o *SnaplockLegalHoldInstanceGetOK) Code() int {
	return 200
}

func (o *SnaplockLegalHoldInstanceGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/litigations/{id}][%d] snaplockLegalHoldInstanceGetOK %s", 200, payload)
}

func (o *SnaplockLegalHoldInstanceGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/litigations/{id}][%d] snaplockLegalHoldInstanceGetOK %s", 200, payload)
}

func (o *SnaplockLegalHoldInstanceGetOK) GetPayload() *models.SnaplockLitigation {
	return o.Payload
}

func (o *SnaplockLegalHoldInstanceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockLitigation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLegalHoldInstanceGetDefault creates a SnaplockLegalHoldInstanceGetDefault with default headers values
func NewSnaplockLegalHoldInstanceGetDefault(code int) *SnaplockLegalHoldInstanceGetDefault {
	return &SnaplockLegalHoldInstanceGetDefault{
		_statusCode: code,
	}
}

/*
	SnaplockLegalHoldInstanceGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 13763280    | Only a user with security login role \"vsadmin-snaplock\" is allowed to perform this operation.  |
| 14090346    | Internal Error. Wait a few minutes, then try the command again  |
| 14090343    | Invalid Field  |
*/
type SnaplockLegalHoldInstanceGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snaplock legal hold instance get default response has a 2xx status code
func (o *SnaplockLegalHoldInstanceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock legal hold instance get default response has a 3xx status code
func (o *SnaplockLegalHoldInstanceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock legal hold instance get default response has a 4xx status code
func (o *SnaplockLegalHoldInstanceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock legal hold instance get default response has a 5xx status code
func (o *SnaplockLegalHoldInstanceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock legal hold instance get default response a status code equal to that given
func (o *SnaplockLegalHoldInstanceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snaplock legal hold instance get default response
func (o *SnaplockLegalHoldInstanceGetDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockLegalHoldInstanceGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/litigations/{id}][%d] snaplock_legal_hold_instance_get default %s", o._statusCode, payload)
}

func (o *SnaplockLegalHoldInstanceGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/litigations/{id}][%d] snaplock_legal_hold_instance_get default %s", o._statusCode, payload)
}

func (o *SnaplockLegalHoldInstanceGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockLegalHoldInstanceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
