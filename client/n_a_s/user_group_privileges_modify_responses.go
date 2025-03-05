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

	"github.com/sapcc/ontap-restapi/models"
)

// UserGroupPrivilegesModifyReader is a Reader for the UserGroupPrivilegesModify structure.
type UserGroupPrivilegesModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGroupPrivilegesModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGroupPrivilegesModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUserGroupPrivilegesModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserGroupPrivilegesModifyOK creates a UserGroupPrivilegesModifyOK with default headers values
func NewUserGroupPrivilegesModifyOK() *UserGroupPrivilegesModifyOK {
	return &UserGroupPrivilegesModifyOK{}
}

/*
UserGroupPrivilegesModifyOK describes a response with status code 200, with default header values.

OK
*/
type UserGroupPrivilegesModifyOK struct {
}

// IsSuccess returns true when this user group privileges modify o k response has a 2xx status code
func (o *UserGroupPrivilegesModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user group privileges modify o k response has a 3xx status code
func (o *UserGroupPrivilegesModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user group privileges modify o k response has a 4xx status code
func (o *UserGroupPrivilegesModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user group privileges modify o k response has a 5xx status code
func (o *UserGroupPrivilegesModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user group privileges modify o k response a status code equal to that given
func (o *UserGroupPrivilegesModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user group privileges modify o k response
func (o *UserGroupPrivilegesModifyOK) Code() int {
	return 200
}

func (o *UserGroupPrivilegesModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/cifs/users-and-groups/privileges/{svm.uuid}/{name}][%d] userGroupPrivilegesModifyOK", 200)
}

func (o *UserGroupPrivilegesModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/cifs/users-and-groups/privileges/{svm.uuid}/{name}][%d] userGroupPrivilegesModifyOK", 200)
}

func (o *UserGroupPrivilegesModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserGroupPrivilegesModifyDefault creates a UserGroupPrivilegesModifyDefault with default headers values
func NewUserGroupPrivilegesModifyDefault(code int) *UserGroupPrivilegesModifyDefault {
	return &UserGroupPrivilegesModifyDefault{
		_statusCode: code,
	}
}

/*
	UserGroupPrivilegesModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262196     | Field 'svm.name' is not supported in the body of PATCH request. |
| 262203     | Field 'svm.uuid' is not supported in the body of PATCH request. |
| 655673     | Failed to resolve the user or group. |
| 655730     | The specified local user to which privileges are to be associated to does not exist. |
*/
type UserGroupPrivilegesModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this user group privileges modify default response has a 2xx status code
func (o *UserGroupPrivilegesModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this user group privileges modify default response has a 3xx status code
func (o *UserGroupPrivilegesModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this user group privileges modify default response has a 4xx status code
func (o *UserGroupPrivilegesModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this user group privileges modify default response has a 5xx status code
func (o *UserGroupPrivilegesModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this user group privileges modify default response a status code equal to that given
func (o *UserGroupPrivilegesModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the user group privileges modify default response
func (o *UserGroupPrivilegesModifyDefault) Code() int {
	return o._statusCode
}

func (o *UserGroupPrivilegesModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/users-and-groups/privileges/{svm.uuid}/{name}][%d] user_group_privileges_modify default %s", o._statusCode, payload)
}

func (o *UserGroupPrivilegesModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/users-and-groups/privileges/{svm.uuid}/{name}][%d] user_group_privileges_modify default %s", o._statusCode, payload)
}

func (o *UserGroupPrivilegesModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UserGroupPrivilegesModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
