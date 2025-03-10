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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// IscsiCredentialsCollectionGetReader is a Reader for the IscsiCredentialsCollectionGet structure.
type IscsiCredentialsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IscsiCredentialsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIscsiCredentialsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIscsiCredentialsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIscsiCredentialsCollectionGetOK creates a IscsiCredentialsCollectionGetOK with default headers values
func NewIscsiCredentialsCollectionGetOK() *IscsiCredentialsCollectionGetOK {
	return &IscsiCredentialsCollectionGetOK{}
}

/*
IscsiCredentialsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type IscsiCredentialsCollectionGetOK struct {
	Payload *models.IscsiCredentialsResponse
}

// IsSuccess returns true when this iscsi credentials collection get o k response has a 2xx status code
func (o *IscsiCredentialsCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iscsi credentials collection get o k response has a 3xx status code
func (o *IscsiCredentialsCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iscsi credentials collection get o k response has a 4xx status code
func (o *IscsiCredentialsCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iscsi credentials collection get o k response has a 5xx status code
func (o *IscsiCredentialsCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iscsi credentials collection get o k response a status code equal to that given
func (o *IscsiCredentialsCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the iscsi credentials collection get o k response
func (o *IscsiCredentialsCollectionGetOK) Code() int {
	return 200
}

func (o *IscsiCredentialsCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/credentials][%d] iscsiCredentialsCollectionGetOK %s", 200, payload)
}

func (o *IscsiCredentialsCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/credentials][%d] iscsiCredentialsCollectionGetOK %s", 200, payload)
}

func (o *IscsiCredentialsCollectionGetOK) GetPayload() *models.IscsiCredentialsResponse {
	return o.Payload
}

func (o *IscsiCredentialsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IscsiCredentialsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIscsiCredentialsCollectionGetDefault creates a IscsiCredentialsCollectionGetDefault with default headers values
func NewIscsiCredentialsCollectionGetDefault(code int) *IscsiCredentialsCollectionGetDefault {
	return &IscsiCredentialsCollectionGetDefault{
		_statusCode: code,
	}
}

/*
IscsiCredentialsCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type IscsiCredentialsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this iscsi credentials collection get default response has a 2xx status code
func (o *IscsiCredentialsCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iscsi credentials collection get default response has a 3xx status code
func (o *IscsiCredentialsCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iscsi credentials collection get default response has a 4xx status code
func (o *IscsiCredentialsCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iscsi credentials collection get default response has a 5xx status code
func (o *IscsiCredentialsCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iscsi credentials collection get default response a status code equal to that given
func (o *IscsiCredentialsCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the iscsi credentials collection get default response
func (o *IscsiCredentialsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *IscsiCredentialsCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/credentials][%d] iscsi_credentials_collection_get default %s", o._statusCode, payload)
}

func (o *IscsiCredentialsCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/credentials][%d] iscsi_credentials_collection_get default %s", o._statusCode, payload)
}

func (o *IscsiCredentialsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IscsiCredentialsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
