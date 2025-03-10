// Code generated by go-swagger; DO NOT EDIT.

package snap_mirror

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

// SnapmirrorPolicyGetReader is a Reader for the SnapmirrorPolicyGet structure.
type SnapmirrorPolicyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapmirrorPolicyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapmirrorPolicyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapmirrorPolicyGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapmirrorPolicyGetOK creates a SnapmirrorPolicyGetOK with default headers values
func NewSnapmirrorPolicyGetOK() *SnapmirrorPolicyGetOK {
	return &SnapmirrorPolicyGetOK{}
}

/*
SnapmirrorPolicyGetOK describes a response with status code 200, with default header values.

OK
*/
type SnapmirrorPolicyGetOK struct {
	Payload *models.SnapmirrorPolicy
}

// IsSuccess returns true when this snapmirror policy get o k response has a 2xx status code
func (o *SnapmirrorPolicyGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapmirror policy get o k response has a 3xx status code
func (o *SnapmirrorPolicyGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapmirror policy get o k response has a 4xx status code
func (o *SnapmirrorPolicyGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapmirror policy get o k response has a 5xx status code
func (o *SnapmirrorPolicyGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snapmirror policy get o k response a status code equal to that given
func (o *SnapmirrorPolicyGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snapmirror policy get o k response
func (o *SnapmirrorPolicyGetOK) Code() int {
	return 200
}

func (o *SnapmirrorPolicyGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /snapmirror/policies/{uuid}][%d] snapmirrorPolicyGetOK %s", 200, payload)
}

func (o *SnapmirrorPolicyGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /snapmirror/policies/{uuid}][%d] snapmirrorPolicyGetOK %s", 200, payload)
}

func (o *SnapmirrorPolicyGetOK) GetPayload() *models.SnapmirrorPolicy {
	return o.Payload
}

func (o *SnapmirrorPolicyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapmirrorPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapmirrorPolicyGetDefault creates a SnapmirrorPolicyGetDefault with default headers values
func NewSnapmirrorPolicyGetDefault(code int) *SnapmirrorPolicyGetDefault {
	return &SnapmirrorPolicyGetDefault{
		_statusCode: code,
	}
}

/*
	SnapmirrorPolicyGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 13303842    | SnapMirror policy is not supported.|
*/
type SnapmirrorPolicyGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snapmirror policy get default response has a 2xx status code
func (o *SnapmirrorPolicyGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapmirror policy get default response has a 3xx status code
func (o *SnapmirrorPolicyGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapmirror policy get default response has a 4xx status code
func (o *SnapmirrorPolicyGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapmirror policy get default response has a 5xx status code
func (o *SnapmirrorPolicyGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapmirror policy get default response a status code equal to that given
func (o *SnapmirrorPolicyGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snapmirror policy get default response
func (o *SnapmirrorPolicyGetDefault) Code() int {
	return o._statusCode
}

func (o *SnapmirrorPolicyGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /snapmirror/policies/{uuid}][%d] snapmirror_policy_get default %s", o._statusCode, payload)
}

func (o *SnapmirrorPolicyGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /snapmirror/policies/{uuid}][%d] snapmirror_policy_get default %s", o._statusCode, payload)
}

func (o *SnapmirrorPolicyGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapmirrorPolicyGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
