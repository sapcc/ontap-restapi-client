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

// NfsGetReader is a Reader for the NfsGet structure.
type NfsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NfsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNfsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNfsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNfsGetOK creates a NfsGetOK with default headers values
func NewNfsGetOK() *NfsGetOK {
	return &NfsGetOK{}
}

/*
NfsGetOK describes a response with status code 200, with default header values.

OK
*/
type NfsGetOK struct {
	Payload *models.NfsService
}

// IsSuccess returns true when this nfs get o k response has a 2xx status code
func (o *NfsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nfs get o k response has a 3xx status code
func (o *NfsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nfs get o k response has a 4xx status code
func (o *NfsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nfs get o k response has a 5xx status code
func (o *NfsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nfs get o k response a status code equal to that given
func (o *NfsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nfs get o k response
func (o *NfsGetOK) Code() int {
	return 200
}

func (o *NfsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/services/{svm.uuid}][%d] nfsGetOK %s", 200, payload)
}

func (o *NfsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/services/{svm.uuid}][%d] nfsGetOK %s", 200, payload)
}

func (o *NfsGetOK) GetPayload() *models.NfsService {
	return o.Payload
}

func (o *NfsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NfsService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNfsGetDefault creates a NfsGetDefault with default headers values
func NewNfsGetDefault(code int) *NfsGetDefault {
	return &NfsGetDefault{
		_statusCode: code,
	}
}

/*
NfsGetDefault describes a response with status code -1, with default header values.

Error
*/
type NfsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nfs get default response has a 2xx status code
func (o *NfsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nfs get default response has a 3xx status code
func (o *NfsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nfs get default response has a 4xx status code
func (o *NfsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nfs get default response has a 5xx status code
func (o *NfsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nfs get default response a status code equal to that given
func (o *NfsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nfs get default response
func (o *NfsGetDefault) Code() int {
	return o._statusCode
}

func (o *NfsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/services/{svm.uuid}][%d] nfs_get default %s", o._statusCode, payload)
}

func (o *NfsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/services/{svm.uuid}][%d] nfs_get default %s", o._statusCode, payload)
}

func (o *NfsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NfsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
