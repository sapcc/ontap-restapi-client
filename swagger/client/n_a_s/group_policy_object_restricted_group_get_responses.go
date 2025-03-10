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

// GroupPolicyObjectRestrictedGroupGetReader is a Reader for the GroupPolicyObjectRestrictedGroupGet structure.
type GroupPolicyObjectRestrictedGroupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupPolicyObjectRestrictedGroupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupPolicyObjectRestrictedGroupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGroupPolicyObjectRestrictedGroupGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGroupPolicyObjectRestrictedGroupGetOK creates a GroupPolicyObjectRestrictedGroupGetOK with default headers values
func NewGroupPolicyObjectRestrictedGroupGetOK() *GroupPolicyObjectRestrictedGroupGetOK {
	return &GroupPolicyObjectRestrictedGroupGetOK{}
}

/*
GroupPolicyObjectRestrictedGroupGetOK describes a response with status code 200, with default header values.

OK
*/
type GroupPolicyObjectRestrictedGroupGetOK struct {
	Payload *models.GroupPolicyObjectRestrictedGroup
}

// IsSuccess returns true when this group policy object restricted group get o k response has a 2xx status code
func (o *GroupPolicyObjectRestrictedGroupGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this group policy object restricted group get o k response has a 3xx status code
func (o *GroupPolicyObjectRestrictedGroupGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group policy object restricted group get o k response has a 4xx status code
func (o *GroupPolicyObjectRestrictedGroupGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this group policy object restricted group get o k response has a 5xx status code
func (o *GroupPolicyObjectRestrictedGroupGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this group policy object restricted group get o k response a status code equal to that given
func (o *GroupPolicyObjectRestrictedGroupGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the group policy object restricted group get o k response
func (o *GroupPolicyObjectRestrictedGroupGetOK) Code() int {
	return 200
}

func (o *GroupPolicyObjectRestrictedGroupGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/restricted-groups/{policy_index}/{group_name}][%d] groupPolicyObjectRestrictedGroupGetOK %s", 200, payload)
}

func (o *GroupPolicyObjectRestrictedGroupGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/restricted-groups/{policy_index}/{group_name}][%d] groupPolicyObjectRestrictedGroupGetOK %s", 200, payload)
}

func (o *GroupPolicyObjectRestrictedGroupGetOK) GetPayload() *models.GroupPolicyObjectRestrictedGroup {
	return o.Payload
}

func (o *GroupPolicyObjectRestrictedGroupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupPolicyObjectRestrictedGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupPolicyObjectRestrictedGroupGetDefault creates a GroupPolicyObjectRestrictedGroupGetDefault with default headers values
func NewGroupPolicyObjectRestrictedGroupGetDefault(code int) *GroupPolicyObjectRestrictedGroupGetDefault {
	return &GroupPolicyObjectRestrictedGroupGetDefault{
		_statusCode: code,
	}
}

/*
GroupPolicyObjectRestrictedGroupGetDefault describes a response with status code -1, with default header values.

Error
*/
type GroupPolicyObjectRestrictedGroupGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this group policy object restricted group get default response has a 2xx status code
func (o *GroupPolicyObjectRestrictedGroupGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this group policy object restricted group get default response has a 3xx status code
func (o *GroupPolicyObjectRestrictedGroupGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this group policy object restricted group get default response has a 4xx status code
func (o *GroupPolicyObjectRestrictedGroupGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this group policy object restricted group get default response has a 5xx status code
func (o *GroupPolicyObjectRestrictedGroupGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this group policy object restricted group get default response a status code equal to that given
func (o *GroupPolicyObjectRestrictedGroupGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the group policy object restricted group get default response
func (o *GroupPolicyObjectRestrictedGroupGetDefault) Code() int {
	return o._statusCode
}

func (o *GroupPolicyObjectRestrictedGroupGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/restricted-groups/{policy_index}/{group_name}][%d] group_policy_object_restricted_group_get default %s", o._statusCode, payload)
}

func (o *GroupPolicyObjectRestrictedGroupGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/restricted-groups/{policy_index}/{group_name}][%d] group_policy_object_restricted_group_get default %s", o._statusCode, payload)
}

func (o *GroupPolicyObjectRestrictedGroupGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GroupPolicyObjectRestrictedGroupGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
