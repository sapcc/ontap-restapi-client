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

// GroupPolicyObjectCentralAccessPolicyGetReader is a Reader for the GroupPolicyObjectCentralAccessPolicyGet structure.
type GroupPolicyObjectCentralAccessPolicyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupPolicyObjectCentralAccessPolicyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupPolicyObjectCentralAccessPolicyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGroupPolicyObjectCentralAccessPolicyGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGroupPolicyObjectCentralAccessPolicyGetOK creates a GroupPolicyObjectCentralAccessPolicyGetOK with default headers values
func NewGroupPolicyObjectCentralAccessPolicyGetOK() *GroupPolicyObjectCentralAccessPolicyGetOK {
	return &GroupPolicyObjectCentralAccessPolicyGetOK{}
}

/*
GroupPolicyObjectCentralAccessPolicyGetOK describes a response with status code 200, with default header values.

OK
*/
type GroupPolicyObjectCentralAccessPolicyGetOK struct {
	Payload *models.GroupPolicyObjectCentralAccessPolicy
}

// IsSuccess returns true when this group policy object central access policy get o k response has a 2xx status code
func (o *GroupPolicyObjectCentralAccessPolicyGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this group policy object central access policy get o k response has a 3xx status code
func (o *GroupPolicyObjectCentralAccessPolicyGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group policy object central access policy get o k response has a 4xx status code
func (o *GroupPolicyObjectCentralAccessPolicyGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this group policy object central access policy get o k response has a 5xx status code
func (o *GroupPolicyObjectCentralAccessPolicyGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this group policy object central access policy get o k response a status code equal to that given
func (o *GroupPolicyObjectCentralAccessPolicyGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the group policy object central access policy get o k response
func (o *GroupPolicyObjectCentralAccessPolicyGetOK) Code() int {
	return 200
}

func (o *GroupPolicyObjectCentralAccessPolicyGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/central-access-policies/{name}][%d] groupPolicyObjectCentralAccessPolicyGetOK %s", 200, payload)
}

func (o *GroupPolicyObjectCentralAccessPolicyGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/central-access-policies/{name}][%d] groupPolicyObjectCentralAccessPolicyGetOK %s", 200, payload)
}

func (o *GroupPolicyObjectCentralAccessPolicyGetOK) GetPayload() *models.GroupPolicyObjectCentralAccessPolicy {
	return o.Payload
}

func (o *GroupPolicyObjectCentralAccessPolicyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupPolicyObjectCentralAccessPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupPolicyObjectCentralAccessPolicyGetDefault creates a GroupPolicyObjectCentralAccessPolicyGetDefault with default headers values
func NewGroupPolicyObjectCentralAccessPolicyGetDefault(code int) *GroupPolicyObjectCentralAccessPolicyGetDefault {
	return &GroupPolicyObjectCentralAccessPolicyGetDefault{
		_statusCode: code,
	}
}

/*
GroupPolicyObjectCentralAccessPolicyGetDefault describes a response with status code -1, with default header values.

Error
*/
type GroupPolicyObjectCentralAccessPolicyGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this group policy object central access policy get default response has a 2xx status code
func (o *GroupPolicyObjectCentralAccessPolicyGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this group policy object central access policy get default response has a 3xx status code
func (o *GroupPolicyObjectCentralAccessPolicyGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this group policy object central access policy get default response has a 4xx status code
func (o *GroupPolicyObjectCentralAccessPolicyGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this group policy object central access policy get default response has a 5xx status code
func (o *GroupPolicyObjectCentralAccessPolicyGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this group policy object central access policy get default response a status code equal to that given
func (o *GroupPolicyObjectCentralAccessPolicyGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the group policy object central access policy get default response
func (o *GroupPolicyObjectCentralAccessPolicyGetDefault) Code() int {
	return o._statusCode
}

func (o *GroupPolicyObjectCentralAccessPolicyGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/central-access-policies/{name}][%d] group_policy_object_central_access_policy_get default %s", o._statusCode, payload)
}

func (o *GroupPolicyObjectCentralAccessPolicyGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/central-access-policies/{name}][%d] group_policy_object_central_access_policy_get default %s", o._statusCode, payload)
}

func (o *GroupPolicyObjectCentralAccessPolicyGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GroupPolicyObjectCentralAccessPolicyGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
