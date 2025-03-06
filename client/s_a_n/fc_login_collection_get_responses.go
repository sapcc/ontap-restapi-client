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

	"github.com/sapcc/ontap-restapi-client/models"
)

// FcLoginCollectionGetReader is a Reader for the FcLoginCollectionGet structure.
type FcLoginCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FcLoginCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFcLoginCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFcLoginCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFcLoginCollectionGetOK creates a FcLoginCollectionGetOK with default headers values
func NewFcLoginCollectionGetOK() *FcLoginCollectionGetOK {
	return &FcLoginCollectionGetOK{}
}

/*
FcLoginCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type FcLoginCollectionGetOK struct {
	Payload *models.FcLoginResponse
}

// IsSuccess returns true when this fc login collection get o k response has a 2xx status code
func (o *FcLoginCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fc login collection get o k response has a 3xx status code
func (o *FcLoginCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fc login collection get o k response has a 4xx status code
func (o *FcLoginCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fc login collection get o k response has a 5xx status code
func (o *FcLoginCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fc login collection get o k response a status code equal to that given
func (o *FcLoginCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fc login collection get o k response
func (o *FcLoginCollectionGetOK) Code() int {
	return 200
}

func (o *FcLoginCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/logins][%d] fcLoginCollectionGetOK %s", 200, payload)
}

func (o *FcLoginCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/logins][%d] fcLoginCollectionGetOK %s", 200, payload)
}

func (o *FcLoginCollectionGetOK) GetPayload() *models.FcLoginResponse {
	return o.Payload
}

func (o *FcLoginCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FcLoginResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFcLoginCollectionGetDefault creates a FcLoginCollectionGetDefault with default headers values
func NewFcLoginCollectionGetDefault(code int) *FcLoginCollectionGetDefault {
	return &FcLoginCollectionGetDefault{
		_statusCode: code,
	}
}

/*
FcLoginCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type FcLoginCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fc login collection get default response has a 2xx status code
func (o *FcLoginCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fc login collection get default response has a 3xx status code
func (o *FcLoginCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fc login collection get default response has a 4xx status code
func (o *FcLoginCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fc login collection get default response has a 5xx status code
func (o *FcLoginCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fc login collection get default response a status code equal to that given
func (o *FcLoginCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fc login collection get default response
func (o *FcLoginCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *FcLoginCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/logins][%d] fc_login_collection_get default %s", o._statusCode, payload)
}

func (o *FcLoginCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/logins][%d] fc_login_collection_get default %s", o._statusCode, payload)
}

func (o *FcLoginCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FcLoginCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
