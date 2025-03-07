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

// GroupPolicyObjectCollectionGetReader is a Reader for the GroupPolicyObjectCollectionGet structure.
type GroupPolicyObjectCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupPolicyObjectCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupPolicyObjectCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGroupPolicyObjectCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGroupPolicyObjectCollectionGetOK creates a GroupPolicyObjectCollectionGetOK with default headers values
func NewGroupPolicyObjectCollectionGetOK() *GroupPolicyObjectCollectionGetOK {
	return &GroupPolicyObjectCollectionGetOK{}
}

/*
GroupPolicyObjectCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type GroupPolicyObjectCollectionGetOK struct {
	Payload *models.GroupPolicyObjectResponse
}

// IsSuccess returns true when this group policy object collection get o k response has a 2xx status code
func (o *GroupPolicyObjectCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this group policy object collection get o k response has a 3xx status code
func (o *GroupPolicyObjectCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group policy object collection get o k response has a 4xx status code
func (o *GroupPolicyObjectCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this group policy object collection get o k response has a 5xx status code
func (o *GroupPolicyObjectCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this group policy object collection get o k response a status code equal to that given
func (o *GroupPolicyObjectCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the group policy object collection get o k response
func (o *GroupPolicyObjectCollectionGetOK) Code() int {
	return 200
}

func (o *GroupPolicyObjectCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/objects][%d] groupPolicyObjectCollectionGetOK %s", 200, payload)
}

func (o *GroupPolicyObjectCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/objects][%d] groupPolicyObjectCollectionGetOK %s", 200, payload)
}

func (o *GroupPolicyObjectCollectionGetOK) GetPayload() *models.GroupPolicyObjectResponse {
	return o.Payload
}

func (o *GroupPolicyObjectCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupPolicyObjectResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupPolicyObjectCollectionGetDefault creates a GroupPolicyObjectCollectionGetDefault with default headers values
func NewGroupPolicyObjectCollectionGetDefault(code int) *GroupPolicyObjectCollectionGetDefault {
	return &GroupPolicyObjectCollectionGetDefault{
		_statusCode: code,
	}
}

/*
GroupPolicyObjectCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type GroupPolicyObjectCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this group policy object collection get default response has a 2xx status code
func (o *GroupPolicyObjectCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this group policy object collection get default response has a 3xx status code
func (o *GroupPolicyObjectCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this group policy object collection get default response has a 4xx status code
func (o *GroupPolicyObjectCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this group policy object collection get default response has a 5xx status code
func (o *GroupPolicyObjectCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this group policy object collection get default response a status code equal to that given
func (o *GroupPolicyObjectCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the group policy object collection get default response
func (o *GroupPolicyObjectCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *GroupPolicyObjectCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/objects][%d] group_policy_object_collection_get default %s", o._statusCode, payload)
}

func (o *GroupPolicyObjectCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies/{svm.uuid}/objects][%d] group_policy_object_collection_get default %s", o._statusCode, payload)
}

func (o *GroupPolicyObjectCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GroupPolicyObjectCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
