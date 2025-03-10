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

// SecurityGroupModifyReader is a Reader for the SecurityGroupModify structure.
type SecurityGroupModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityGroupModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityGroupModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityGroupModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityGroupModifyOK creates a SecurityGroupModifyOK with default headers values
func NewSecurityGroupModifyOK() *SecurityGroupModifyOK {
	return &SecurityGroupModifyOK{}
}

/*
SecurityGroupModifyOK describes a response with status code 200, with default header values.

OK
*/
type SecurityGroupModifyOK struct {
}

// IsSuccess returns true when this security group modify o k response has a 2xx status code
func (o *SecurityGroupModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security group modify o k response has a 3xx status code
func (o *SecurityGroupModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security group modify o k response has a 4xx status code
func (o *SecurityGroupModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security group modify o k response has a 5xx status code
func (o *SecurityGroupModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security group modify o k response a status code equal to that given
func (o *SecurityGroupModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security group modify o k response
func (o *SecurityGroupModifyOK) Code() int {
	return 200
}

func (o *SecurityGroupModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /security/groups/{owner.uuid}/{name}/{type}][%d] securityGroupModifyOK", 200)
}

func (o *SecurityGroupModifyOK) String() string {
	return fmt.Sprintf("[PATCH /security/groups/{owner.uuid}/{name}/{type}][%d] securityGroupModifyOK", 200)
}

func (o *SecurityGroupModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecurityGroupModifyDefault creates a SecurityGroupModifyDefault with default headers values
func NewSecurityGroupModifyDefault(code int) *SecurityGroupModifyDefault {
	return &SecurityGroupModifyDefault{
		_statusCode: code,
	}
}

/*
	SecurityGroupModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5636235 | Duplicate UUID. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SecurityGroupModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security group modify default response has a 2xx status code
func (o *SecurityGroupModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security group modify default response has a 3xx status code
func (o *SecurityGroupModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security group modify default response has a 4xx status code
func (o *SecurityGroupModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security group modify default response has a 5xx status code
func (o *SecurityGroupModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security group modify default response a status code equal to that given
func (o *SecurityGroupModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security group modify default response
func (o *SecurityGroupModifyDefault) Code() int {
	return o._statusCode
}

func (o *SecurityGroupModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/groups/{owner.uuid}/{name}/{type}][%d] security_group_modify default %s", o._statusCode, payload)
}

func (o *SecurityGroupModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/groups/{owner.uuid}/{name}/{type}][%d] security_group_modify default %s", o._statusCode, payload)
}

func (o *SecurityGroupModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityGroupModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
