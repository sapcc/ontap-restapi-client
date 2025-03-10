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

// GroupRoleMappingsModifyReader is a Reader for the GroupRoleMappingsModify structure.
type GroupRoleMappingsModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupRoleMappingsModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupRoleMappingsModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGroupRoleMappingsModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGroupRoleMappingsModifyOK creates a GroupRoleMappingsModifyOK with default headers values
func NewGroupRoleMappingsModifyOK() *GroupRoleMappingsModifyOK {
	return &GroupRoleMappingsModifyOK{}
}

/*
GroupRoleMappingsModifyOK describes a response with status code 200, with default header values.

OK
*/
type GroupRoleMappingsModifyOK struct {
}

// IsSuccess returns true when this group role mappings modify o k response has a 2xx status code
func (o *GroupRoleMappingsModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this group role mappings modify o k response has a 3xx status code
func (o *GroupRoleMappingsModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group role mappings modify o k response has a 4xx status code
func (o *GroupRoleMappingsModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this group role mappings modify o k response has a 5xx status code
func (o *GroupRoleMappingsModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this group role mappings modify o k response a status code equal to that given
func (o *GroupRoleMappingsModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the group role mappings modify o k response
func (o *GroupRoleMappingsModifyOK) Code() int {
	return 200
}

func (o *GroupRoleMappingsModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /security/group/role-mappings/{group_id}/{ontap_role.name}][%d] groupRoleMappingsModifyOK", 200)
}

func (o *GroupRoleMappingsModifyOK) String() string {
	return fmt.Sprintf("[PATCH /security/group/role-mappings/{group_id}/{ontap_role.name}][%d] groupRoleMappingsModifyOK", 200)
}

func (o *GroupRoleMappingsModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGroupRoleMappingsModifyDefault creates a GroupRoleMappingsModifyDefault with default headers values
func NewGroupRoleMappingsModifyDefault(code int) *GroupRoleMappingsModifyDefault {
	return &GroupRoleMappingsModifyDefault{
		_statusCode: code,
	}
}

/*
GroupRoleMappingsModifyDefault describes a response with status code -1, with default header values.

Error
*/
type GroupRoleMappingsModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this group role mappings modify default response has a 2xx status code
func (o *GroupRoleMappingsModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this group role mappings modify default response has a 3xx status code
func (o *GroupRoleMappingsModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this group role mappings modify default response has a 4xx status code
func (o *GroupRoleMappingsModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this group role mappings modify default response has a 5xx status code
func (o *GroupRoleMappingsModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this group role mappings modify default response a status code equal to that given
func (o *GroupRoleMappingsModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the group role mappings modify default response
func (o *GroupRoleMappingsModifyDefault) Code() int {
	return o._statusCode
}

func (o *GroupRoleMappingsModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/group/role-mappings/{group_id}/{ontap_role.name}][%d] group_role_mappings_modify default %s", o._statusCode, payload)
}

func (o *GroupRoleMappingsModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/group/role-mappings/{group_id}/{ontap_role.name}][%d] group_role_mappings_modify default %s", o._statusCode, payload)
}

func (o *GroupRoleMappingsModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GroupRoleMappingsModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
