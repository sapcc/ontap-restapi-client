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

// GroupPolicyObjectCentralAccessPolicyCollectionGetReader is a Reader for the GroupPolicyObjectCentralAccessPolicyCollectionGet structure.
type GroupPolicyObjectCentralAccessPolicyCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupPolicyObjectCentralAccessPolicyCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGroupPolicyObjectCentralAccessPolicyCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGroupPolicyObjectCentralAccessPolicyCollectionGetOK creates a GroupPolicyObjectCentralAccessPolicyCollectionGetOK with default headers values
func NewGroupPolicyObjectCentralAccessPolicyCollectionGetOK() *GroupPolicyObjectCentralAccessPolicyCollectionGetOK {
	return &GroupPolicyObjectCentralAccessPolicyCollectionGetOK{}
}

/*
GroupPolicyObjectCentralAccessPolicyCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type GroupPolicyObjectCentralAccessPolicyCollectionGetOK struct {
	Payload *models.GroupPolicyObjectCentralAccessPolicyResponse
}

// IsSuccess returns true when this group policy object central access policy collection get o k response has a 2xx status code
func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this group policy object central access policy collection get o k response has a 3xx status code
func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group policy object central access policy collection get o k response has a 4xx status code
func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this group policy object central access policy collection get o k response has a 5xx status code
func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this group policy object central access policy collection get o k response a status code equal to that given
func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the group policy object central access policy collection get o k response
func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetOK) Code() int {
	return 200
}

func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/central-access-policies][%d] groupPolicyObjectCentralAccessPolicyCollectionGetOK %s", 200, payload)
}

func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/central-access-policies][%d] groupPolicyObjectCentralAccessPolicyCollectionGetOK %s", 200, payload)
}

func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetOK) GetPayload() *models.GroupPolicyObjectCentralAccessPolicyResponse {
	return o.Payload
}

func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupPolicyObjectCentralAccessPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupPolicyObjectCentralAccessPolicyCollectionGetDefault creates a GroupPolicyObjectCentralAccessPolicyCollectionGetDefault with default headers values
func NewGroupPolicyObjectCentralAccessPolicyCollectionGetDefault(code int) *GroupPolicyObjectCentralAccessPolicyCollectionGetDefault {
	return &GroupPolicyObjectCentralAccessPolicyCollectionGetDefault{
		_statusCode: code,
	}
}

/*
GroupPolicyObjectCentralAccessPolicyCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type GroupPolicyObjectCentralAccessPolicyCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this group policy object central access policy collection get default response has a 2xx status code
func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this group policy object central access policy collection get default response has a 3xx status code
func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this group policy object central access policy collection get default response has a 4xx status code
func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this group policy object central access policy collection get default response has a 5xx status code
func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this group policy object central access policy collection get default response a status code equal to that given
func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the group policy object central access policy collection get default response
func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/central-access-policies][%d] group_policy_object_central_access_policy_collection_get default %s", o._statusCode, payload)
}

func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/central-access-policies][%d] group_policy_object_central_access_policy_collection_get default %s", o._statusCode, payload)
}

func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GroupPolicyObjectCentralAccessPolicyCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
