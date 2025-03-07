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

// LocalCifsGroupMembersGetReader is a Reader for the LocalCifsGroupMembersGet structure.
type LocalCifsGroupMembersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalCifsGroupMembersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocalCifsGroupMembersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocalCifsGroupMembersGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocalCifsGroupMembersGetOK creates a LocalCifsGroupMembersGetOK with default headers values
func NewLocalCifsGroupMembersGetOK() *LocalCifsGroupMembersGetOK {
	return &LocalCifsGroupMembersGetOK{}
}

/*
LocalCifsGroupMembersGetOK describes a response with status code 200, with default header values.

OK
*/
type LocalCifsGroupMembersGetOK struct {
	Payload *models.LocalCifsGroupMembers
}

// IsSuccess returns true when this local cifs group members get o k response has a 2xx status code
func (o *LocalCifsGroupMembersGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this local cifs group members get o k response has a 3xx status code
func (o *LocalCifsGroupMembersGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this local cifs group members get o k response has a 4xx status code
func (o *LocalCifsGroupMembersGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this local cifs group members get o k response has a 5xx status code
func (o *LocalCifsGroupMembersGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this local cifs group members get o k response a status code equal to that given
func (o *LocalCifsGroupMembersGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the local cifs group members get o k response
func (o *LocalCifsGroupMembersGetOK) Code() int {
	return 200
}

func (o *LocalCifsGroupMembersGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/local-groups/{svm.uuid}/{local_cifs_group.sid}/members/{name}][%d] localCifsGroupMembersGetOK %s", 200, payload)
}

func (o *LocalCifsGroupMembersGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/local-groups/{svm.uuid}/{local_cifs_group.sid}/members/{name}][%d] localCifsGroupMembersGetOK %s", 200, payload)
}

func (o *LocalCifsGroupMembersGetOK) GetPayload() *models.LocalCifsGroupMembers {
	return o.Payload
}

func (o *LocalCifsGroupMembersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocalCifsGroupMembers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocalCifsGroupMembersGetDefault creates a LocalCifsGroupMembersGetDefault with default headers values
func NewLocalCifsGroupMembersGetDefault(code int) *LocalCifsGroupMembersGetDefault {
	return &LocalCifsGroupMembersGetDefault{
		_statusCode: code,
	}
}

/*
LocalCifsGroupMembersGetDefault describes a response with status code -1, with default header values.

Error
*/
type LocalCifsGroupMembersGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this local cifs group members get default response has a 2xx status code
func (o *LocalCifsGroupMembersGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this local cifs group members get default response has a 3xx status code
func (o *LocalCifsGroupMembersGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this local cifs group members get default response has a 4xx status code
func (o *LocalCifsGroupMembersGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this local cifs group members get default response has a 5xx status code
func (o *LocalCifsGroupMembersGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this local cifs group members get default response a status code equal to that given
func (o *LocalCifsGroupMembersGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the local cifs group members get default response
func (o *LocalCifsGroupMembersGetDefault) Code() int {
	return o._statusCode
}

func (o *LocalCifsGroupMembersGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/local-groups/{svm.uuid}/{local_cifs_group.sid}/members/{name}][%d] local_cifs_group_members_get default %s", o._statusCode, payload)
}

func (o *LocalCifsGroupMembersGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/local-groups/{svm.uuid}/{local_cifs_group.sid}/members/{name}][%d] local_cifs_group_members_get default %s", o._statusCode, payload)
}

func (o *LocalCifsGroupMembersGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LocalCifsGroupMembersGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
