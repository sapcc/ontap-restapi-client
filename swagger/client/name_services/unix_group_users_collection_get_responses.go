// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// UnixGroupUsersCollectionGetReader is a Reader for the UnixGroupUsersCollectionGet structure.
type UnixGroupUsersCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnixGroupUsersCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnixGroupUsersCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnixGroupUsersCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnixGroupUsersCollectionGetOK creates a UnixGroupUsersCollectionGetOK with default headers values
func NewUnixGroupUsersCollectionGetOK() *UnixGroupUsersCollectionGetOK {
	return &UnixGroupUsersCollectionGetOK{}
}

/*
UnixGroupUsersCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type UnixGroupUsersCollectionGetOK struct {
	Payload *models.UnixGroupUsersResponse
}

// IsSuccess returns true when this unix group users collection get o k response has a 2xx status code
func (o *UnixGroupUsersCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unix group users collection get o k response has a 3xx status code
func (o *UnixGroupUsersCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unix group users collection get o k response has a 4xx status code
func (o *UnixGroupUsersCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unix group users collection get o k response has a 5xx status code
func (o *UnixGroupUsersCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unix group users collection get o k response a status code equal to that given
func (o *UnixGroupUsersCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the unix group users collection get o k response
func (o *UnixGroupUsersCollectionGetOK) Code() int {
	return 200
}

func (o *UnixGroupUsersCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/unix-groups/{svm.uuid}/{unix_group.name}/users][%d] unixGroupUsersCollectionGetOK %s", 200, payload)
}

func (o *UnixGroupUsersCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/unix-groups/{svm.uuid}/{unix_group.name}/users][%d] unixGroupUsersCollectionGetOK %s", 200, payload)
}

func (o *UnixGroupUsersCollectionGetOK) GetPayload() *models.UnixGroupUsersResponse {
	return o.Payload
}

func (o *UnixGroupUsersCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnixGroupUsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnixGroupUsersCollectionGetDefault creates a UnixGroupUsersCollectionGetDefault with default headers values
func NewUnixGroupUsersCollectionGetDefault(code int) *UnixGroupUsersCollectionGetDefault {
	return &UnixGroupUsersCollectionGetDefault{
		_statusCode: code,
	}
}

/*
UnixGroupUsersCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type UnixGroupUsersCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this unix group users collection get default response has a 2xx status code
func (o *UnixGroupUsersCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unix group users collection get default response has a 3xx status code
func (o *UnixGroupUsersCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unix group users collection get default response has a 4xx status code
func (o *UnixGroupUsersCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unix group users collection get default response has a 5xx status code
func (o *UnixGroupUsersCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unix group users collection get default response a status code equal to that given
func (o *UnixGroupUsersCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the unix group users collection get default response
func (o *UnixGroupUsersCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *UnixGroupUsersCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/unix-groups/{svm.uuid}/{unix_group.name}/users][%d] unix_group_users_collection_get default %s", o._statusCode, payload)
}

func (o *UnixGroupUsersCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/unix-groups/{svm.uuid}/{unix_group.name}/users][%d] unix_group_users_collection_get default %s", o._statusCode, payload)
}

func (o *UnixGroupUsersCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnixGroupUsersCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
