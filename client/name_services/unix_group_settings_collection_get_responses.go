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

	"github.com/sapcc/ontap-restapi-client/models"
)

// UnixGroupSettingsCollectionGetReader is a Reader for the UnixGroupSettingsCollectionGet structure.
type UnixGroupSettingsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnixGroupSettingsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnixGroupSettingsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnixGroupSettingsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnixGroupSettingsCollectionGetOK creates a UnixGroupSettingsCollectionGetOK with default headers values
func NewUnixGroupSettingsCollectionGetOK() *UnixGroupSettingsCollectionGetOK {
	return &UnixGroupSettingsCollectionGetOK{}
}

/*
UnixGroupSettingsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type UnixGroupSettingsCollectionGetOK struct {
	Payload *models.UnixGroupSettingsResponse
}

// IsSuccess returns true when this unix group settings collection get o k response has a 2xx status code
func (o *UnixGroupSettingsCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unix group settings collection get o k response has a 3xx status code
func (o *UnixGroupSettingsCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unix group settings collection get o k response has a 4xx status code
func (o *UnixGroupSettingsCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unix group settings collection get o k response has a 5xx status code
func (o *UnixGroupSettingsCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unix group settings collection get o k response a status code equal to that given
func (o *UnixGroupSettingsCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the unix group settings collection get o k response
func (o *UnixGroupSettingsCollectionGetOK) Code() int {
	return 200
}

func (o *UnixGroupSettingsCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/cache/unix-group/settings][%d] unixGroupSettingsCollectionGetOK %s", 200, payload)
}

func (o *UnixGroupSettingsCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/cache/unix-group/settings][%d] unixGroupSettingsCollectionGetOK %s", 200, payload)
}

func (o *UnixGroupSettingsCollectionGetOK) GetPayload() *models.UnixGroupSettingsResponse {
	return o.Payload
}

func (o *UnixGroupSettingsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnixGroupSettingsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnixGroupSettingsCollectionGetDefault creates a UnixGroupSettingsCollectionGetDefault with default headers values
func NewUnixGroupSettingsCollectionGetDefault(code int) *UnixGroupSettingsCollectionGetDefault {
	return &UnixGroupSettingsCollectionGetDefault{
		_statusCode: code,
	}
}

/*
UnixGroupSettingsCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type UnixGroupSettingsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this unix group settings collection get default response has a 2xx status code
func (o *UnixGroupSettingsCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unix group settings collection get default response has a 3xx status code
func (o *UnixGroupSettingsCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unix group settings collection get default response has a 4xx status code
func (o *UnixGroupSettingsCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unix group settings collection get default response has a 5xx status code
func (o *UnixGroupSettingsCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unix group settings collection get default response a status code equal to that given
func (o *UnixGroupSettingsCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the unix group settings collection get default response
func (o *UnixGroupSettingsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *UnixGroupSettingsCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/cache/unix-group/settings][%d] unix_group_settings_collection_get default %s", o._statusCode, payload)
}

func (o *UnixGroupSettingsCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/cache/unix-group/settings][%d] unix_group_settings_collection_get default %s", o._statusCode, payload)
}

func (o *UnixGroupSettingsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnixGroupSettingsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
