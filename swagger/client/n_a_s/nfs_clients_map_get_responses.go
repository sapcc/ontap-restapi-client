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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NfsClientsMapGetReader is a Reader for the NfsClientsMapGet structure.
type NfsClientsMapGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NfsClientsMapGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNfsClientsMapGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNfsClientsMapGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNfsClientsMapGetOK creates a NfsClientsMapGetOK with default headers values
func NewNfsClientsMapGetOK() *NfsClientsMapGetOK {
	return &NfsClientsMapGetOK{}
}

/*
NfsClientsMapGetOK describes a response with status code 200, with default header values.

OK
*/
type NfsClientsMapGetOK struct {
	Payload *models.NfsClientsMapResponse
}

// IsSuccess returns true when this nfs clients map get o k response has a 2xx status code
func (o *NfsClientsMapGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nfs clients map get o k response has a 3xx status code
func (o *NfsClientsMapGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nfs clients map get o k response has a 4xx status code
func (o *NfsClientsMapGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nfs clients map get o k response has a 5xx status code
func (o *NfsClientsMapGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nfs clients map get o k response a status code equal to that given
func (o *NfsClientsMapGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nfs clients map get o k response
func (o *NfsClientsMapGetOK) Code() int {
	return 200
}

func (o *NfsClientsMapGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/connected-client-maps][%d] nfsClientsMapGetOK %s", 200, payload)
}

func (o *NfsClientsMapGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/connected-client-maps][%d] nfsClientsMapGetOK %s", 200, payload)
}

func (o *NfsClientsMapGetOK) GetPayload() *models.NfsClientsMapResponse {
	return o.Payload
}

func (o *NfsClientsMapGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NfsClientsMapResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNfsClientsMapGetDefault creates a NfsClientsMapGetDefault with default headers values
func NewNfsClientsMapGetDefault(code int) *NfsClientsMapGetDefault {
	return &NfsClientsMapGetDefault{
		_statusCode: code,
	}
}

/*
NfsClientsMapGetDefault describes a response with status code -1, with default header values.

Error
*/
type NfsClientsMapGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nfs clients map get default response has a 2xx status code
func (o *NfsClientsMapGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nfs clients map get default response has a 3xx status code
func (o *NfsClientsMapGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nfs clients map get default response has a 4xx status code
func (o *NfsClientsMapGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nfs clients map get default response has a 5xx status code
func (o *NfsClientsMapGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nfs clients map get default response a status code equal to that given
func (o *NfsClientsMapGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nfs clients map get default response
func (o *NfsClientsMapGetDefault) Code() int {
	return o._statusCode
}

func (o *NfsClientsMapGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/connected-client-maps][%d] nfs_clients_map_get default %s", o._statusCode, payload)
}

func (o *NfsClientsMapGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/connected-client-maps][%d] nfs_clients_map_get default %s", o._statusCode, payload)
}

func (o *NfsClientsMapGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NfsClientsMapGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
