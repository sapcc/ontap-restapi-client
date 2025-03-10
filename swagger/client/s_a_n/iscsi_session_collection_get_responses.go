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

// IscsiSessionCollectionGetReader is a Reader for the IscsiSessionCollectionGet structure.
type IscsiSessionCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IscsiSessionCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIscsiSessionCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIscsiSessionCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIscsiSessionCollectionGetOK creates a IscsiSessionCollectionGetOK with default headers values
func NewIscsiSessionCollectionGetOK() *IscsiSessionCollectionGetOK {
	return &IscsiSessionCollectionGetOK{}
}

/*
IscsiSessionCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type IscsiSessionCollectionGetOK struct {
	Payload *models.IscsiSessionResponse
}

// IsSuccess returns true when this iscsi session collection get o k response has a 2xx status code
func (o *IscsiSessionCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iscsi session collection get o k response has a 3xx status code
func (o *IscsiSessionCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iscsi session collection get o k response has a 4xx status code
func (o *IscsiSessionCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iscsi session collection get o k response has a 5xx status code
func (o *IscsiSessionCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iscsi session collection get o k response a status code equal to that given
func (o *IscsiSessionCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the iscsi session collection get o k response
func (o *IscsiSessionCollectionGetOK) Code() int {
	return 200
}

func (o *IscsiSessionCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/sessions][%d] iscsiSessionCollectionGetOK %s", 200, payload)
}

func (o *IscsiSessionCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/sessions][%d] iscsiSessionCollectionGetOK %s", 200, payload)
}

func (o *IscsiSessionCollectionGetOK) GetPayload() *models.IscsiSessionResponse {
	return o.Payload
}

func (o *IscsiSessionCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IscsiSessionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIscsiSessionCollectionGetDefault creates a IscsiSessionCollectionGetDefault with default headers values
func NewIscsiSessionCollectionGetDefault(code int) *IscsiSessionCollectionGetDefault {
	return &IscsiSessionCollectionGetDefault{
		_statusCode: code,
	}
}

/*
IscsiSessionCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type IscsiSessionCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this iscsi session collection get default response has a 2xx status code
func (o *IscsiSessionCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iscsi session collection get default response has a 3xx status code
func (o *IscsiSessionCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iscsi session collection get default response has a 4xx status code
func (o *IscsiSessionCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iscsi session collection get default response has a 5xx status code
func (o *IscsiSessionCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iscsi session collection get default response a status code equal to that given
func (o *IscsiSessionCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the iscsi session collection get default response
func (o *IscsiSessionCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *IscsiSessionCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/sessions][%d] iscsi_session_collection_get default %s", o._statusCode, payload)
}

func (o *IscsiSessionCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/sessions][%d] iscsi_session_collection_get default %s", o._statusCode, payload)
}

func (o *IscsiSessionCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IscsiSessionCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
