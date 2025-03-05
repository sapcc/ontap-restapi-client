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

	"github.com/sapcc/ontap-restapi/models"
)

// RolePrivilegeDeleteReader is a Reader for the RolePrivilegeDelete structure.
type RolePrivilegeDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RolePrivilegeDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRolePrivilegeDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRolePrivilegeDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRolePrivilegeDeleteOK creates a RolePrivilegeDeleteOK with default headers values
func NewRolePrivilegeDeleteOK() *RolePrivilegeDeleteOK {
	return &RolePrivilegeDeleteOK{}
}

/*
RolePrivilegeDeleteOK describes a response with status code 200, with default header values.

OK
*/
type RolePrivilegeDeleteOK struct {
}

// IsSuccess returns true when this role privilege delete o k response has a 2xx status code
func (o *RolePrivilegeDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this role privilege delete o k response has a 3xx status code
func (o *RolePrivilegeDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this role privilege delete o k response has a 4xx status code
func (o *RolePrivilegeDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this role privilege delete o k response has a 5xx status code
func (o *RolePrivilegeDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this role privilege delete o k response a status code equal to that given
func (o *RolePrivilegeDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the role privilege delete o k response
func (o *RolePrivilegeDeleteOK) Code() int {
	return 200
}

func (o *RolePrivilegeDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/roles/{owner.uuid}/{name}/privileges/{path}][%d] rolePrivilegeDeleteOK", 200)
}

func (o *RolePrivilegeDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /security/roles/{owner.uuid}/{name}/privileges/{path}][%d] rolePrivilegeDeleteOK", 200)
}

func (o *RolePrivilegeDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRolePrivilegeDeleteDefault creates a RolePrivilegeDeleteDefault with default headers values
func NewRolePrivilegeDeleteDefault(code int) *RolePrivilegeDeleteDefault {
	return &RolePrivilegeDeleteDefault{
		_statusCode: code,
	}
}

/*
	RolePrivilegeDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1263347 | Cannot modify pre-defined roles. |
| 5636168 | This role is mapped to a rest-role and cannot be modified directly. Modifications must be done with rest-role. |
| 5636169 | Specified URI path is invalid or not supported. Resource-qualified endpoints are not supported. |
| 5636170 | URI does not exist. |
| 5636172 | User accounts detected with this role assigned. Update or delete those accounts before deleting this role. |
| 5636173 | This feature requires an effective cluster version of 9.6 or later. |
| 5636184 | Expanded REST roles for granular resource control feature is currently disabled. |
| 5636185 | The specified UUID was not found. |
| 5636186 | Expanded REST roles for granular resource control requires an effective cluster version of 9.10.1 or later. |
| 13434890 | Vserver-ID failed for Vserver roles. |
| 13434893 | The SVM does not exist. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type RolePrivilegeDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this role privilege delete default response has a 2xx status code
func (o *RolePrivilegeDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this role privilege delete default response has a 3xx status code
func (o *RolePrivilegeDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this role privilege delete default response has a 4xx status code
func (o *RolePrivilegeDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this role privilege delete default response has a 5xx status code
func (o *RolePrivilegeDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this role privilege delete default response a status code equal to that given
func (o *RolePrivilegeDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the role privilege delete default response
func (o *RolePrivilegeDeleteDefault) Code() int {
	return o._statusCode
}

func (o *RolePrivilegeDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/roles/{owner.uuid}/{name}/privileges/{path}][%d] role_privilege_delete default %s", o._statusCode, payload)
}

func (o *RolePrivilegeDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/roles/{owner.uuid}/{name}/privileges/{path}][%d] role_privilege_delete default %s", o._statusCode, payload)
}

func (o *RolePrivilegeDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RolePrivilegeDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
