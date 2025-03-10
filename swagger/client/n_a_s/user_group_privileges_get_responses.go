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

// UserGroupPrivilegesGetReader is a Reader for the UserGroupPrivilegesGet structure.
type UserGroupPrivilegesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGroupPrivilegesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGroupPrivilegesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUserGroupPrivilegesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserGroupPrivilegesGetOK creates a UserGroupPrivilegesGetOK with default headers values
func NewUserGroupPrivilegesGetOK() *UserGroupPrivilegesGetOK {
	return &UserGroupPrivilegesGetOK{}
}

/*
UserGroupPrivilegesGetOK describes a response with status code 200, with default header values.

OK
*/
type UserGroupPrivilegesGetOK struct {
	Payload *models.UserGroupPrivileges
}

// IsSuccess returns true when this user group privileges get o k response has a 2xx status code
func (o *UserGroupPrivilegesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user group privileges get o k response has a 3xx status code
func (o *UserGroupPrivilegesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user group privileges get o k response has a 4xx status code
func (o *UserGroupPrivilegesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user group privileges get o k response has a 5xx status code
func (o *UserGroupPrivilegesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user group privileges get o k response a status code equal to that given
func (o *UserGroupPrivilegesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user group privileges get o k response
func (o *UserGroupPrivilegesGetOK) Code() int {
	return 200
}

func (o *UserGroupPrivilegesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/users-and-groups/privileges/{svm.uuid}/{name}][%d] userGroupPrivilegesGetOK %s", 200, payload)
}

func (o *UserGroupPrivilegesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/users-and-groups/privileges/{svm.uuid}/{name}][%d] userGroupPrivilegesGetOK %s", 200, payload)
}

func (o *UserGroupPrivilegesGetOK) GetPayload() *models.UserGroupPrivileges {
	return o.Payload
}

func (o *UserGroupPrivilegesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserGroupPrivileges)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupPrivilegesGetDefault creates a UserGroupPrivilegesGetDefault with default headers values
func NewUserGroupPrivilegesGetDefault(code int) *UserGroupPrivilegesGetDefault {
	return &UserGroupPrivilegesGetDefault{
		_statusCode: code,
	}
}

/*
UserGroupPrivilegesGetDefault describes a response with status code -1, with default header values.

Error
*/
type UserGroupPrivilegesGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this user group privileges get default response has a 2xx status code
func (o *UserGroupPrivilegesGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this user group privileges get default response has a 3xx status code
func (o *UserGroupPrivilegesGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this user group privileges get default response has a 4xx status code
func (o *UserGroupPrivilegesGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this user group privileges get default response has a 5xx status code
func (o *UserGroupPrivilegesGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this user group privileges get default response a status code equal to that given
func (o *UserGroupPrivilegesGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the user group privileges get default response
func (o *UserGroupPrivilegesGetDefault) Code() int {
	return o._statusCode
}

func (o *UserGroupPrivilegesGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/users-and-groups/privileges/{svm.uuid}/{name}][%d] user_group_privileges_get default %s", o._statusCode, payload)
}

func (o *UserGroupPrivilegesGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/users-and-groups/privileges/{svm.uuid}/{name}][%d] user_group_privileges_get default %s", o._statusCode, payload)
}

func (o *UserGroupPrivilegesGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UserGroupPrivilegesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
