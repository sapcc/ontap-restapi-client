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

// NfsClientsGetReader is a Reader for the NfsClientsGet structure.
type NfsClientsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NfsClientsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNfsClientsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNfsClientsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNfsClientsGetOK creates a NfsClientsGetOK with default headers values
func NewNfsClientsGetOK() *NfsClientsGetOK {
	return &NfsClientsGetOK{}
}

/*
NfsClientsGetOK describes a response with status code 200, with default header values.

OK
*/
type NfsClientsGetOK struct {
	Payload *models.NfsClientsResponse
}

// IsSuccess returns true when this nfs clients get o k response has a 2xx status code
func (o *NfsClientsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nfs clients get o k response has a 3xx status code
func (o *NfsClientsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nfs clients get o k response has a 4xx status code
func (o *NfsClientsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nfs clients get o k response has a 5xx status code
func (o *NfsClientsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nfs clients get o k response a status code equal to that given
func (o *NfsClientsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nfs clients get o k response
func (o *NfsClientsGetOK) Code() int {
	return 200
}

func (o *NfsClientsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/connected-clients][%d] nfsClientsGetOK %s", 200, payload)
}

func (o *NfsClientsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/connected-clients][%d] nfsClientsGetOK %s", 200, payload)
}

func (o *NfsClientsGetOK) GetPayload() *models.NfsClientsResponse {
	return o.Payload
}

func (o *NfsClientsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NfsClientsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNfsClientsGetDefault creates a NfsClientsGetDefault with default headers values
func NewNfsClientsGetDefault(code int) *NfsClientsGetDefault {
	return &NfsClientsGetDefault{
		_statusCode: code,
	}
}

/*
	NfsClientsGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 917887     | The first character of the volume name must be a letter or underscore.|
| 917888     | Invalid character in volume name. It can have '_' and alphanumeric characters.|
*/
type NfsClientsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nfs clients get default response has a 2xx status code
func (o *NfsClientsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nfs clients get default response has a 3xx status code
func (o *NfsClientsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nfs clients get default response has a 4xx status code
func (o *NfsClientsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nfs clients get default response has a 5xx status code
func (o *NfsClientsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nfs clients get default response a status code equal to that given
func (o *NfsClientsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nfs clients get default response
func (o *NfsClientsGetDefault) Code() int {
	return o._statusCode
}

func (o *NfsClientsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/connected-clients][%d] nfs_clients_get default %s", o._statusCode, payload)
}

func (o *NfsClientsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/connected-clients][%d] nfs_clients_get default %s", o._statusCode, payload)
}

func (o *NfsClientsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NfsClientsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
