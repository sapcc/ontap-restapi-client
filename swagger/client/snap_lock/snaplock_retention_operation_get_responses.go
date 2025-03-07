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

// SnaplockRetentionOperationGetReader is a Reader for the SnaplockRetentionOperationGet structure.
type SnaplockRetentionOperationGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockRetentionOperationGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockRetentionOperationGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockRetentionOperationGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockRetentionOperationGetOK creates a SnaplockRetentionOperationGetOK with default headers values
func NewSnaplockRetentionOperationGetOK() *SnaplockRetentionOperationGetOK {
	return &SnaplockRetentionOperationGetOK{}
}

/*
SnaplockRetentionOperationGetOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockRetentionOperationGetOK struct {
	Payload *models.EbrOperation
}

// IsSuccess returns true when this snaplock retention operation get o k response has a 2xx status code
func (o *SnaplockRetentionOperationGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock retention operation get o k response has a 3xx status code
func (o *SnaplockRetentionOperationGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock retention operation get o k response has a 4xx status code
func (o *SnaplockRetentionOperationGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock retention operation get o k response has a 5xx status code
func (o *SnaplockRetentionOperationGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock retention operation get o k response a status code equal to that given
func (o *SnaplockRetentionOperationGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snaplock retention operation get o k response
func (o *SnaplockRetentionOperationGetOK) Code() int {
	return 200
}

func (o *SnaplockRetentionOperationGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/event-retention/operations/{id}][%d] snaplockRetentionOperationGetOK %s", 200, payload)
}

func (o *SnaplockRetentionOperationGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/event-retention/operations/{id}][%d] snaplockRetentionOperationGetOK %s", 200, payload)
}

func (o *SnaplockRetentionOperationGetOK) GetPayload() *models.EbrOperation {
	return o.Payload
}

func (o *SnaplockRetentionOperationGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EbrOperation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockRetentionOperationGetDefault creates a SnaplockRetentionOperationGetDefault with default headers values
func NewSnaplockRetentionOperationGetDefault(code int) *SnaplockRetentionOperationGetDefault {
	return &SnaplockRetentionOperationGetDefault{
		_statusCode: code,
	}
}

/*
	SnaplockRetentionOperationGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 14090242    | Only a user with the security login role \"vsadmin-snaplock\" is allowed to perform this operation. |
*/
type SnaplockRetentionOperationGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snaplock retention operation get default response has a 2xx status code
func (o *SnaplockRetentionOperationGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock retention operation get default response has a 3xx status code
func (o *SnaplockRetentionOperationGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock retention operation get default response has a 4xx status code
func (o *SnaplockRetentionOperationGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock retention operation get default response has a 5xx status code
func (o *SnaplockRetentionOperationGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock retention operation get default response a status code equal to that given
func (o *SnaplockRetentionOperationGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snaplock retention operation get default response
func (o *SnaplockRetentionOperationGetDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockRetentionOperationGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/event-retention/operations/{id}][%d] snaplock_retention_operation_get default %s", o._statusCode, payload)
}

func (o *SnaplockRetentionOperationGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/event-retention/operations/{id}][%d] snaplock_retention_operation_get default %s", o._statusCode, payload)
}

func (o *SnaplockRetentionOperationGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockRetentionOperationGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
