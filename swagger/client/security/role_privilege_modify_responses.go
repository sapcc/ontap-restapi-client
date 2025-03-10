// Code generated by go-swagger; DO NOT EDIT.

package security

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

// RolePrivilegeModifyReader is a Reader for the RolePrivilegeModify structure.
type RolePrivilegeModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RolePrivilegeModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRolePrivilegeModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRolePrivilegeModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRolePrivilegeModifyOK creates a RolePrivilegeModifyOK with default headers values
func NewRolePrivilegeModifyOK() *RolePrivilegeModifyOK {
	return &RolePrivilegeModifyOK{}
}

/*
RolePrivilegeModifyOK describes a response with status code 200, with default header values.

OK
*/
type RolePrivilegeModifyOK struct {
}

// IsSuccess returns true when this role privilege modify o k response has a 2xx status code
func (o *RolePrivilegeModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this role privilege modify o k response has a 3xx status code
func (o *RolePrivilegeModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this role privilege modify o k response has a 4xx status code
func (o *RolePrivilegeModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this role privilege modify o k response has a 5xx status code
func (o *RolePrivilegeModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this role privilege modify o k response a status code equal to that given
func (o *RolePrivilegeModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the role privilege modify o k response
func (o *RolePrivilegeModifyOK) Code() int {
	return 200
}

func (o *RolePrivilegeModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /security/roles/{owner.uuid}/{name}/privileges/{path}][%d] rolePrivilegeModifyOK", 200)
}

func (o *RolePrivilegeModifyOK) String() string {
	return fmt.Sprintf("[PATCH /security/roles/{owner.uuid}/{name}/privileges/{path}][%d] rolePrivilegeModifyOK", 200)
}

func (o *RolePrivilegeModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRolePrivilegeModifyDefault creates a RolePrivilegeModifyDefault with default headers values
func NewRolePrivilegeModifyDefault(code int) *RolePrivilegeModifyDefault {
	return &RolePrivilegeModifyDefault{
		_statusCode: code,
	}
}

/*
	RolePrivilegeModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5636168 | This role is mapped to a rest-role and cannot be modified directly. Modifications must be done with rest-role. |
| 5636192 | The query parameter cannot be specified for the privileges tuple with API endpoint entries. |
| 5636200 | The specified value of the access parameter is invalid, if a command or command directory is specified in the path parameter. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type RolePrivilegeModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this role privilege modify default response has a 2xx status code
func (o *RolePrivilegeModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this role privilege modify default response has a 3xx status code
func (o *RolePrivilegeModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this role privilege modify default response has a 4xx status code
func (o *RolePrivilegeModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this role privilege modify default response has a 5xx status code
func (o *RolePrivilegeModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this role privilege modify default response a status code equal to that given
func (o *RolePrivilegeModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the role privilege modify default response
func (o *RolePrivilegeModifyDefault) Code() int {
	return o._statusCode
}

func (o *RolePrivilegeModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/roles/{owner.uuid}/{name}/privileges/{path}][%d] role_privilege_modify default %s", o._statusCode, payload)
}

func (o *RolePrivilegeModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/roles/{owner.uuid}/{name}/privileges/{path}][%d] role_privilege_modify default %s", o._statusCode, payload)
}

func (o *RolePrivilegeModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RolePrivilegeModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
