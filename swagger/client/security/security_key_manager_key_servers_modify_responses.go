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

// SecurityKeyManagerKeyServersModifyReader is a Reader for the SecurityKeyManagerKeyServersModify structure.
type SecurityKeyManagerKeyServersModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeyManagerKeyServersModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityKeyManagerKeyServersModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeyManagerKeyServersModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeyManagerKeyServersModifyOK creates a SecurityKeyManagerKeyServersModifyOK with default headers values
func NewSecurityKeyManagerKeyServersModifyOK() *SecurityKeyManagerKeyServersModifyOK {
	return &SecurityKeyManagerKeyServersModifyOK{}
}

/*
SecurityKeyManagerKeyServersModifyOK describes a response with status code 200, with default header values.

OK
*/
type SecurityKeyManagerKeyServersModifyOK struct {
}

// IsSuccess returns true when this security key manager key servers modify o k response has a 2xx status code
func (o *SecurityKeyManagerKeyServersModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security key manager key servers modify o k response has a 3xx status code
func (o *SecurityKeyManagerKeyServersModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security key manager key servers modify o k response has a 4xx status code
func (o *SecurityKeyManagerKeyServersModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security key manager key servers modify o k response has a 5xx status code
func (o *SecurityKeyManagerKeyServersModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security key manager key servers modify o k response a status code equal to that given
func (o *SecurityKeyManagerKeyServersModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security key manager key servers modify o k response
func (o *SecurityKeyManagerKeyServersModifyOK) Code() int {
	return 200
}

func (o *SecurityKeyManagerKeyServersModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /security/key-managers/{uuid}/key-servers/{server}][%d] securityKeyManagerKeyServersModifyOK", 200)
}

func (o *SecurityKeyManagerKeyServersModifyOK) String() string {
	return fmt.Sprintf("[PATCH /security/key-managers/{uuid}/key-servers/{server}][%d] securityKeyManagerKeyServersModifyOK", 200)
}

func (o *SecurityKeyManagerKeyServersModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecurityKeyManagerKeyServersModifyDefault creates a SecurityKeyManagerKeyServersModifyDefault with default headers values
func NewSecurityKeyManagerKeyServersModifyDefault(code int) *SecurityKeyManagerKeyServersModifyDefault {
	return &SecurityKeyManagerKeyServersModifyDefault{
		_statusCode: code,
	}
}

/*
	SecurityKeyManagerKeyServersModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65536600 | Cannot modify a key server while a node is out quorum. |
| 65536824 | Multitenant key management is not supported in MetroCluster configurations. |
| 65536828 | External key management is not enabled for the SVM. |
| 65536843 | The key management server is not configured for the SVM. |
| 65536845 | Missing username. |
| 65536846 | Missing password. |
| 65536880 | One or more of the following values must be provided \"timeout\", \"username\", \"password\", \"secondary_key_servers\", \"create_remove_timeout\". |
| 65536921 | Unable to execute the command on the KMIP server. |
| 65537400 | Exceeded maximum number of secondary key servers. |
| 65538407 | A secondary key server is a duplicate of the associated primary key server. |
| 65538408 | The list of secondary key servers contains duplicates. |
| 65538413 | A secondary key server address is not formatted correctly. |
| 65538502 | A secondary key server is also a primary key server. |
| 65538503 | Support for adding secondary key servers requires an ECV of ONTAP 9.11.1 or later. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SecurityKeyManagerKeyServersModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security key manager key servers modify default response has a 2xx status code
func (o *SecurityKeyManagerKeyServersModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security key manager key servers modify default response has a 3xx status code
func (o *SecurityKeyManagerKeyServersModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security key manager key servers modify default response has a 4xx status code
func (o *SecurityKeyManagerKeyServersModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security key manager key servers modify default response has a 5xx status code
func (o *SecurityKeyManagerKeyServersModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security key manager key servers modify default response a status code equal to that given
func (o *SecurityKeyManagerKeyServersModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security key manager key servers modify default response
func (o *SecurityKeyManagerKeyServersModifyDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeyManagerKeyServersModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/key-managers/{uuid}/key-servers/{server}][%d] security_key_manager_key_servers_modify default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerKeyServersModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/key-managers/{uuid}/key-servers/{server}][%d] security_key_manager_key_servers_modify default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerKeyServersModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeyManagerKeyServersModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
