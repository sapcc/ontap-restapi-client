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

// SnaplockRetentionPolicyGetReader is a Reader for the SnaplockRetentionPolicyGet structure.
type SnaplockRetentionPolicyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockRetentionPolicyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockRetentionPolicyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockRetentionPolicyGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockRetentionPolicyGetOK creates a SnaplockRetentionPolicyGetOK with default headers values
func NewSnaplockRetentionPolicyGetOK() *SnaplockRetentionPolicyGetOK {
	return &SnaplockRetentionPolicyGetOK{}
}

/*
SnaplockRetentionPolicyGetOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockRetentionPolicyGetOK struct {
	Payload *models.SnaplockRetentionPolicy
}

// IsSuccess returns true when this snaplock retention policy get o k response has a 2xx status code
func (o *SnaplockRetentionPolicyGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock retention policy get o k response has a 3xx status code
func (o *SnaplockRetentionPolicyGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock retention policy get o k response has a 4xx status code
func (o *SnaplockRetentionPolicyGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock retention policy get o k response has a 5xx status code
func (o *SnaplockRetentionPolicyGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock retention policy get o k response a status code equal to that given
func (o *SnaplockRetentionPolicyGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snaplock retention policy get o k response
func (o *SnaplockRetentionPolicyGetOK) Code() int {
	return 200
}

func (o *SnaplockRetentionPolicyGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/event-retention/policies/{policy.name}][%d] snaplockRetentionPolicyGetOK %s", 200, payload)
}

func (o *SnaplockRetentionPolicyGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/event-retention/policies/{policy.name}][%d] snaplockRetentionPolicyGetOK %s", 200, payload)
}

func (o *SnaplockRetentionPolicyGetOK) GetPayload() *models.SnaplockRetentionPolicy {
	return o.Payload
}

func (o *SnaplockRetentionPolicyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockRetentionPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockRetentionPolicyGetDefault creates a SnaplockRetentionPolicyGetDefault with default headers values
func NewSnaplockRetentionPolicyGetDefault(code int) *SnaplockRetentionPolicyGetDefault {
	return &SnaplockRetentionPolicyGetDefault{
		_statusCode: code,
	}
}

/*
	SnaplockRetentionPolicyGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 13763280    | Only a user with security login role \"vsadmin-snaplock\" is allowed to perform this operation.  |
*/
type SnaplockRetentionPolicyGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snaplock retention policy get default response has a 2xx status code
func (o *SnaplockRetentionPolicyGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock retention policy get default response has a 3xx status code
func (o *SnaplockRetentionPolicyGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock retention policy get default response has a 4xx status code
func (o *SnaplockRetentionPolicyGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock retention policy get default response has a 5xx status code
func (o *SnaplockRetentionPolicyGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock retention policy get default response a status code equal to that given
func (o *SnaplockRetentionPolicyGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snaplock retention policy get default response
func (o *SnaplockRetentionPolicyGetDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockRetentionPolicyGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/event-retention/policies/{policy.name}][%d] snaplock_retention_policy_get default %s", o._statusCode, payload)
}

func (o *SnaplockRetentionPolicyGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/event-retention/policies/{policy.name}][%d] snaplock_retention_policy_get default %s", o._statusCode, payload)
}

func (o *SnaplockRetentionPolicyGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockRetentionPolicyGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
