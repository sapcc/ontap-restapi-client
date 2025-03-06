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

	"github.com/sapcc/ontap-restapi-client/models"
)

// UserGroupPrivilegesCreateReader is a Reader for the UserGroupPrivilegesCreate structure.
type UserGroupPrivilegesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGroupPrivilegesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUserGroupPrivilegesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUserGroupPrivilegesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserGroupPrivilegesCreateCreated creates a UserGroupPrivilegesCreateCreated with default headers values
func NewUserGroupPrivilegesCreateCreated() *UserGroupPrivilegesCreateCreated {
	return &UserGroupPrivilegesCreateCreated{}
}

/*
UserGroupPrivilegesCreateCreated describes a response with status code 201, with default header values.

Created
*/
type UserGroupPrivilegesCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string
}

// IsSuccess returns true when this user group privileges create created response has a 2xx status code
func (o *UserGroupPrivilegesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user group privileges create created response has a 3xx status code
func (o *UserGroupPrivilegesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user group privileges create created response has a 4xx status code
func (o *UserGroupPrivilegesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this user group privileges create created response has a 5xx status code
func (o *UserGroupPrivilegesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this user group privileges create created response a status code equal to that given
func (o *UserGroupPrivilegesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the user group privileges create created response
func (o *UserGroupPrivilegesCreateCreated) Code() int {
	return 201
}

func (o *UserGroupPrivilegesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /protocols/cifs/users-and-groups/privileges][%d] userGroupPrivilegesCreateCreated", 201)
}

func (o *UserGroupPrivilegesCreateCreated) String() string {
	return fmt.Sprintf("[POST /protocols/cifs/users-and-groups/privileges][%d] userGroupPrivilegesCreateCreated", 201)
}

func (o *UserGroupPrivilegesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewUserGroupPrivilegesCreateDefault creates a UserGroupPrivilegesCreateDefault with default headers values
func NewUserGroupPrivilegesCreateDefault(code int) *UserGroupPrivilegesCreateDefault {
	return &UserGroupPrivilegesCreateDefault{
		_statusCode: code,
	}
}

/*
	UserGroupPrivilegesCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 655660     | The operation is allowed only on data SVMs. |
| 655673     | Failed to resolve the user or group. |
| 655730     | The specified local user to which privileges are to be associated to does not exist. |
*/
type UserGroupPrivilegesCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this user group privileges create default response has a 2xx status code
func (o *UserGroupPrivilegesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this user group privileges create default response has a 3xx status code
func (o *UserGroupPrivilegesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this user group privileges create default response has a 4xx status code
func (o *UserGroupPrivilegesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this user group privileges create default response has a 5xx status code
func (o *UserGroupPrivilegesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this user group privileges create default response a status code equal to that given
func (o *UserGroupPrivilegesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the user group privileges create default response
func (o *UserGroupPrivilegesCreateDefault) Code() int {
	return o._statusCode
}

func (o *UserGroupPrivilegesCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/cifs/users-and-groups/privileges][%d] user_group_privileges_create default %s", o._statusCode, payload)
}

func (o *UserGroupPrivilegesCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/cifs/users-and-groups/privileges][%d] user_group_privileges_create default %s", o._statusCode, payload)
}

func (o *UserGroupPrivilegesCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UserGroupPrivilegesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
