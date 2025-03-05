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

// ClientLockDeleteReader is a Reader for the ClientLockDelete structure.
type ClientLockDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClientLockDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClientLockDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClientLockDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClientLockDeleteOK creates a ClientLockDeleteOK with default headers values
func NewClientLockDeleteOK() *ClientLockDeleteOK {
	return &ClientLockDeleteOK{}
}

/*
ClientLockDeleteOK describes a response with status code 200, with default header values.

OK
*/
type ClientLockDeleteOK struct {
}

// IsSuccess returns true when this client lock delete o k response has a 2xx status code
func (o *ClientLockDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this client lock delete o k response has a 3xx status code
func (o *ClientLockDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this client lock delete o k response has a 4xx status code
func (o *ClientLockDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this client lock delete o k response has a 5xx status code
func (o *ClientLockDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this client lock delete o k response a status code equal to that given
func (o *ClientLockDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the client lock delete o k response
func (o *ClientLockDeleteOK) Code() int {
	return 200
}

func (o *ClientLockDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/locks/{uuid}][%d] clientLockDeleteOK", 200)
}

func (o *ClientLockDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/locks/{uuid}][%d] clientLockDeleteOK", 200)
}

func (o *ClientLockDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClientLockDeleteDefault creates a ClientLockDeleteDefault with default headers values
func NewClientLockDeleteDefault(code int) *ClientLockDeleteDefault {
	return &ClientLockDeleteDefault{
		_statusCode: code,
	}
}

/*
ClientLockDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type ClientLockDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this client lock delete default response has a 2xx status code
func (o *ClientLockDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this client lock delete default response has a 3xx status code
func (o *ClientLockDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this client lock delete default response has a 4xx status code
func (o *ClientLockDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this client lock delete default response has a 5xx status code
func (o *ClientLockDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this client lock delete default response a status code equal to that given
func (o *ClientLockDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the client lock delete default response
func (o *ClientLockDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ClientLockDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/locks/{uuid}][%d] client_lock_delete default %s", o._statusCode, payload)
}

func (o *ClientLockDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/locks/{uuid}][%d] client_lock_delete default %s", o._statusCode, payload)
}

func (o *ClientLockDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClientLockDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
