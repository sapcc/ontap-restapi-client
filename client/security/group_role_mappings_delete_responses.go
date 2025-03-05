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

// GroupRoleMappingsDeleteReader is a Reader for the GroupRoleMappingsDelete structure.
type GroupRoleMappingsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupRoleMappingsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupRoleMappingsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGroupRoleMappingsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGroupRoleMappingsDeleteOK creates a GroupRoleMappingsDeleteOK with default headers values
func NewGroupRoleMappingsDeleteOK() *GroupRoleMappingsDeleteOK {
	return &GroupRoleMappingsDeleteOK{}
}

/*
GroupRoleMappingsDeleteOK describes a response with status code 200, with default header values.

OK
*/
type GroupRoleMappingsDeleteOK struct {
}

// IsSuccess returns true when this group role mappings delete o k response has a 2xx status code
func (o *GroupRoleMappingsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this group role mappings delete o k response has a 3xx status code
func (o *GroupRoleMappingsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group role mappings delete o k response has a 4xx status code
func (o *GroupRoleMappingsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this group role mappings delete o k response has a 5xx status code
func (o *GroupRoleMappingsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this group role mappings delete o k response a status code equal to that given
func (o *GroupRoleMappingsDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the group role mappings delete o k response
func (o *GroupRoleMappingsDeleteOK) Code() int {
	return 200
}

func (o *GroupRoleMappingsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/group/role-mappings/{group_id}/{ontap_role.name}][%d] groupRoleMappingsDeleteOK", 200)
}

func (o *GroupRoleMappingsDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /security/group/role-mappings/{group_id}/{ontap_role.name}][%d] groupRoleMappingsDeleteOK", 200)
}

func (o *GroupRoleMappingsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGroupRoleMappingsDeleteDefault creates a GroupRoleMappingsDeleteDefault with default headers values
func NewGroupRoleMappingsDeleteDefault(code int) *GroupRoleMappingsDeleteDefault {
	return &GroupRoleMappingsDeleteDefault{
		_statusCode: code,
	}
}

/*
GroupRoleMappingsDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type GroupRoleMappingsDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this group role mappings delete default response has a 2xx status code
func (o *GroupRoleMappingsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this group role mappings delete default response has a 3xx status code
func (o *GroupRoleMappingsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this group role mappings delete default response has a 4xx status code
func (o *GroupRoleMappingsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this group role mappings delete default response has a 5xx status code
func (o *GroupRoleMappingsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this group role mappings delete default response a status code equal to that given
func (o *GroupRoleMappingsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the group role mappings delete default response
func (o *GroupRoleMappingsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *GroupRoleMappingsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/group/role-mappings/{group_id}/{ontap_role.name}][%d] group_role_mappings_delete default %s", o._statusCode, payload)
}

func (o *GroupRoleMappingsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/group/role-mappings/{group_id}/{ontap_role.name}][%d] group_role_mappings_delete default %s", o._statusCode, payload)
}

func (o *GroupRoleMappingsDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GroupRoleMappingsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
