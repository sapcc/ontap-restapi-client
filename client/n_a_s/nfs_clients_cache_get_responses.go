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

// NfsClientsCacheGetReader is a Reader for the NfsClientsCacheGet structure.
type NfsClientsCacheGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NfsClientsCacheGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNfsClientsCacheGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNfsClientsCacheGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNfsClientsCacheGetOK creates a NfsClientsCacheGetOK with default headers values
func NewNfsClientsCacheGetOK() *NfsClientsCacheGetOK {
	return &NfsClientsCacheGetOK{}
}

/*
NfsClientsCacheGetOK describes a response with status code 200, with default header values.

OK
*/
type NfsClientsCacheGetOK struct {
	Payload *models.NfsClientsCache
}

// IsSuccess returns true when this nfs clients cache get o k response has a 2xx status code
func (o *NfsClientsCacheGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nfs clients cache get o k response has a 3xx status code
func (o *NfsClientsCacheGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nfs clients cache get o k response has a 4xx status code
func (o *NfsClientsCacheGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nfs clients cache get o k response has a 5xx status code
func (o *NfsClientsCacheGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nfs clients cache get o k response a status code equal to that given
func (o *NfsClientsCacheGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nfs clients cache get o k response
func (o *NfsClientsCacheGetOK) Code() int {
	return 200
}

func (o *NfsClientsCacheGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/connected-client-settings][%d] nfsClientsCacheGetOK %s", 200, payload)
}

func (o *NfsClientsCacheGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/connected-client-settings][%d] nfsClientsCacheGetOK %s", 200, payload)
}

func (o *NfsClientsCacheGetOK) GetPayload() *models.NfsClientsCache {
	return o.Payload
}

func (o *NfsClientsCacheGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NfsClientsCache)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNfsClientsCacheGetDefault creates a NfsClientsCacheGetDefault with default headers values
func NewNfsClientsCacheGetDefault(code int) *NfsClientsCacheGetDefault {
	return &NfsClientsCacheGetDefault{
		_statusCode: code,
	}
}

/*
NfsClientsCacheGetDefault describes a response with status code -1, with default header values.

Error
*/
type NfsClientsCacheGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nfs clients cache get default response has a 2xx status code
func (o *NfsClientsCacheGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nfs clients cache get default response has a 3xx status code
func (o *NfsClientsCacheGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nfs clients cache get default response has a 4xx status code
func (o *NfsClientsCacheGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nfs clients cache get default response has a 5xx status code
func (o *NfsClientsCacheGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nfs clients cache get default response a status code equal to that given
func (o *NfsClientsCacheGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nfs clients cache get default response
func (o *NfsClientsCacheGetDefault) Code() int {
	return o._statusCode
}

func (o *NfsClientsCacheGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/connected-client-settings][%d] nfs_clients_cache_get default %s", o._statusCode, payload)
}

func (o *NfsClientsCacheGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/connected-client-settings][%d] nfs_clients_cache_get default %s", o._statusCode, payload)
}

func (o *NfsClientsCacheGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NfsClientsCacheGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
