// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// IscsiServiceCollectionGetReader is a Reader for the IscsiServiceCollectionGet structure.
type IscsiServiceCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IscsiServiceCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIscsiServiceCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIscsiServiceCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIscsiServiceCollectionGetOK creates a IscsiServiceCollectionGetOK with default headers values
func NewIscsiServiceCollectionGetOK() *IscsiServiceCollectionGetOK {
	return &IscsiServiceCollectionGetOK{}
}

/*
IscsiServiceCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type IscsiServiceCollectionGetOK struct {
	Payload *models.IscsiServiceResponse
}

// IsSuccess returns true when this iscsi service collection get o k response has a 2xx status code
func (o *IscsiServiceCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iscsi service collection get o k response has a 3xx status code
func (o *IscsiServiceCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iscsi service collection get o k response has a 4xx status code
func (o *IscsiServiceCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iscsi service collection get o k response has a 5xx status code
func (o *IscsiServiceCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iscsi service collection get o k response a status code equal to that given
func (o *IscsiServiceCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the iscsi service collection get o k response
func (o *IscsiServiceCollectionGetOK) Code() int {
	return 200
}

func (o *IscsiServiceCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/services][%d] iscsiServiceCollectionGetOK %s", 200, payload)
}

func (o *IscsiServiceCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/services][%d] iscsiServiceCollectionGetOK %s", 200, payload)
}

func (o *IscsiServiceCollectionGetOK) GetPayload() *models.IscsiServiceResponse {
	return o.Payload
}

func (o *IscsiServiceCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IscsiServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIscsiServiceCollectionGetDefault creates a IscsiServiceCollectionGetDefault with default headers values
func NewIscsiServiceCollectionGetDefault(code int) *IscsiServiceCollectionGetDefault {
	return &IscsiServiceCollectionGetDefault{
		_statusCode: code,
	}
}

/*
IscsiServiceCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type IscsiServiceCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this iscsi service collection get default response has a 2xx status code
func (o *IscsiServiceCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iscsi service collection get default response has a 3xx status code
func (o *IscsiServiceCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iscsi service collection get default response has a 4xx status code
func (o *IscsiServiceCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iscsi service collection get default response has a 5xx status code
func (o *IscsiServiceCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iscsi service collection get default response a status code equal to that given
func (o *IscsiServiceCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the iscsi service collection get default response
func (o *IscsiServiceCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *IscsiServiceCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/services][%d] iscsi_service_collection_get default %s", o._statusCode, payload)
}

func (o *IscsiServiceCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/services][%d] iscsi_service_collection_get default %s", o._statusCode, payload)
}

func (o *IscsiServiceCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IscsiServiceCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
