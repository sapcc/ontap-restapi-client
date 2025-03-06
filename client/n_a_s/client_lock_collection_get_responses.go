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

	"github.com/sapcc/ontap-restapi-client/models"
)

// ClientLockCollectionGetReader is a Reader for the ClientLockCollectionGet structure.
type ClientLockCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClientLockCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClientLockCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClientLockCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClientLockCollectionGetOK creates a ClientLockCollectionGetOK with default headers values
func NewClientLockCollectionGetOK() *ClientLockCollectionGetOK {
	return &ClientLockCollectionGetOK{}
}

/*
ClientLockCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ClientLockCollectionGetOK struct {
	Payload *models.ClientLockResponse
}

// IsSuccess returns true when this client lock collection get o k response has a 2xx status code
func (o *ClientLockCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this client lock collection get o k response has a 3xx status code
func (o *ClientLockCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this client lock collection get o k response has a 4xx status code
func (o *ClientLockCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this client lock collection get o k response has a 5xx status code
func (o *ClientLockCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this client lock collection get o k response a status code equal to that given
func (o *ClientLockCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the client lock collection get o k response
func (o *ClientLockCollectionGetOK) Code() int {
	return 200
}

func (o *ClientLockCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/locks][%d] clientLockCollectionGetOK %s", 200, payload)
}

func (o *ClientLockCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/locks][%d] clientLockCollectionGetOK %s", 200, payload)
}

func (o *ClientLockCollectionGetOK) GetPayload() *models.ClientLockResponse {
	return o.Payload
}

func (o *ClientLockCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClientLockResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClientLockCollectionGetDefault creates a ClientLockCollectionGetDefault with default headers values
func NewClientLockCollectionGetDefault(code int) *ClientLockCollectionGetDefault {
	return &ClientLockCollectionGetDefault{
		_statusCode: code,
	}
}

/*
ClientLockCollectionGetDefault describes a response with status code -1, with default header values.

unexpected error_response
*/
type ClientLockCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this client lock collection get default response has a 2xx status code
func (o *ClientLockCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this client lock collection get default response has a 3xx status code
func (o *ClientLockCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this client lock collection get default response has a 4xx status code
func (o *ClientLockCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this client lock collection get default response has a 5xx status code
func (o *ClientLockCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this client lock collection get default response a status code equal to that given
func (o *ClientLockCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the client lock collection get default response
func (o *ClientLockCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *ClientLockCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/locks][%d] client_lock_collection_get default %s", o._statusCode, payload)
}

func (o *ClientLockCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/locks][%d] client_lock_collection_get default %s", o._statusCode, payload)
}

func (o *ClientLockCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClientLockCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
