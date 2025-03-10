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

// UserGroupPrivilegesCollectionGetReader is a Reader for the UserGroupPrivilegesCollectionGet structure.
type UserGroupPrivilegesCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGroupPrivilegesCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGroupPrivilegesCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUserGroupPrivilegesCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserGroupPrivilegesCollectionGetOK creates a UserGroupPrivilegesCollectionGetOK with default headers values
func NewUserGroupPrivilegesCollectionGetOK() *UserGroupPrivilegesCollectionGetOK {
	return &UserGroupPrivilegesCollectionGetOK{}
}

/*
UserGroupPrivilegesCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type UserGroupPrivilegesCollectionGetOK struct {
	Payload *models.UserGroupPrivilegesResponse
}

// IsSuccess returns true when this user group privileges collection get o k response has a 2xx status code
func (o *UserGroupPrivilegesCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user group privileges collection get o k response has a 3xx status code
func (o *UserGroupPrivilegesCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user group privileges collection get o k response has a 4xx status code
func (o *UserGroupPrivilegesCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user group privileges collection get o k response has a 5xx status code
func (o *UserGroupPrivilegesCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user group privileges collection get o k response a status code equal to that given
func (o *UserGroupPrivilegesCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user group privileges collection get o k response
func (o *UserGroupPrivilegesCollectionGetOK) Code() int {
	return 200
}

func (o *UserGroupPrivilegesCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/users-and-groups/privileges][%d] userGroupPrivilegesCollectionGetOK %s", 200, payload)
}

func (o *UserGroupPrivilegesCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/users-and-groups/privileges][%d] userGroupPrivilegesCollectionGetOK %s", 200, payload)
}

func (o *UserGroupPrivilegesCollectionGetOK) GetPayload() *models.UserGroupPrivilegesResponse {
	return o.Payload
}

func (o *UserGroupPrivilegesCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserGroupPrivilegesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupPrivilegesCollectionGetDefault creates a UserGroupPrivilegesCollectionGetDefault with default headers values
func NewUserGroupPrivilegesCollectionGetDefault(code int) *UserGroupPrivilegesCollectionGetDefault {
	return &UserGroupPrivilegesCollectionGetDefault{
		_statusCode: code,
	}
}

/*
UserGroupPrivilegesCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type UserGroupPrivilegesCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this user group privileges collection get default response has a 2xx status code
func (o *UserGroupPrivilegesCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this user group privileges collection get default response has a 3xx status code
func (o *UserGroupPrivilegesCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this user group privileges collection get default response has a 4xx status code
func (o *UserGroupPrivilegesCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this user group privileges collection get default response has a 5xx status code
func (o *UserGroupPrivilegesCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this user group privileges collection get default response a status code equal to that given
func (o *UserGroupPrivilegesCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the user group privileges collection get default response
func (o *UserGroupPrivilegesCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *UserGroupPrivilegesCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/users-and-groups/privileges][%d] user_group_privileges_collection_get default %s", o._statusCode, payload)
}

func (o *UserGroupPrivilegesCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/users-and-groups/privileges][%d] user_group_privileges_collection_get default %s", o._statusCode, payload)
}

func (o *UserGroupPrivilegesCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UserGroupPrivilegesCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
