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

// SecurityLogForwardingCreateReader is a Reader for the SecurityLogForwardingCreate structure.
type SecurityLogForwardingCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityLogForwardingCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSecurityLogForwardingCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSecurityLogForwardingCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityLogForwardingCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityLogForwardingCreateCreated creates a SecurityLogForwardingCreateCreated with default headers values
func NewSecurityLogForwardingCreateCreated() *SecurityLogForwardingCreateCreated {
	return &SecurityLogForwardingCreateCreated{}
}

/*
SecurityLogForwardingCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SecurityLogForwardingCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.SecurityAuditLogForwardResponse
}

// IsSuccess returns true when this security log forwarding create created response has a 2xx status code
func (o *SecurityLogForwardingCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security log forwarding create created response has a 3xx status code
func (o *SecurityLogForwardingCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security log forwarding create created response has a 4xx status code
func (o *SecurityLogForwardingCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this security log forwarding create created response has a 5xx status code
func (o *SecurityLogForwardingCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this security log forwarding create created response a status code equal to that given
func (o *SecurityLogForwardingCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the security log forwarding create created response
func (o *SecurityLogForwardingCreateCreated) Code() int {
	return 201
}

func (o *SecurityLogForwardingCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/audit/destinations][%d] securityLogForwardingCreateCreated %s", 201, payload)
}

func (o *SecurityLogForwardingCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/audit/destinations][%d] securityLogForwardingCreateCreated %s", 201, payload)
}

func (o *SecurityLogForwardingCreateCreated) GetPayload() *models.SecurityAuditLogForwardResponse {
	return o.Payload
}

func (o *SecurityLogForwardingCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.SecurityAuditLogForwardResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityLogForwardingCreateAccepted creates a SecurityLogForwardingCreateAccepted with default headers values
func NewSecurityLogForwardingCreateAccepted() *SecurityLogForwardingCreateAccepted {
	return &SecurityLogForwardingCreateAccepted{}
}

/*
SecurityLogForwardingCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SecurityLogForwardingCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.SecurityAuditLogForwardResponse
}

// IsSuccess returns true when this security log forwarding create accepted response has a 2xx status code
func (o *SecurityLogForwardingCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security log forwarding create accepted response has a 3xx status code
func (o *SecurityLogForwardingCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security log forwarding create accepted response has a 4xx status code
func (o *SecurityLogForwardingCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this security log forwarding create accepted response has a 5xx status code
func (o *SecurityLogForwardingCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this security log forwarding create accepted response a status code equal to that given
func (o *SecurityLogForwardingCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the security log forwarding create accepted response
func (o *SecurityLogForwardingCreateAccepted) Code() int {
	return 202
}

func (o *SecurityLogForwardingCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/audit/destinations][%d] securityLogForwardingCreateAccepted %s", 202, payload)
}

func (o *SecurityLogForwardingCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/audit/destinations][%d] securityLogForwardingCreateAccepted %s", 202, payload)
}

func (o *SecurityLogForwardingCreateAccepted) GetPayload() *models.SecurityAuditLogForwardResponse {
	return o.Payload
}

func (o *SecurityLogForwardingCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.SecurityAuditLogForwardResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityLogForwardingCreateDefault creates a SecurityLogForwardingCreateDefault with default headers values
func NewSecurityLogForwardingCreateDefault(code int) *SecurityLogForwardingCreateDefault {
	return &SecurityLogForwardingCreateDefault{
		_statusCode: code,
	}
}

/*
	SecurityLogForwardingCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 15661      | The object specified could not be found |
| 13114      | Internal error |
| 13115      | Invalid input |
| 4522285    | Server verification cannot be enabled because it requires a protocol with encryption. Encryption can be selected using the protocol field.|
| 9240603    | Cannot ping destination host. Verify connectivity to desired host or skip the connectivity check with the -force parameter. |
| 327698     | Failed to create RPC client to destination host |
| 9240609    | Cannot connect to destination host. |
| 9240604    | Cannot resolve the destination host. |
*/
type SecurityLogForwardingCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security log forwarding create default response has a 2xx status code
func (o *SecurityLogForwardingCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security log forwarding create default response has a 3xx status code
func (o *SecurityLogForwardingCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security log forwarding create default response has a 4xx status code
func (o *SecurityLogForwardingCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security log forwarding create default response has a 5xx status code
func (o *SecurityLogForwardingCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security log forwarding create default response a status code equal to that given
func (o *SecurityLogForwardingCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security log forwarding create default response
func (o *SecurityLogForwardingCreateDefault) Code() int {
	return o._statusCode
}

func (o *SecurityLogForwardingCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/audit/destinations][%d] security_log_forwarding_create default %s", o._statusCode, payload)
}

func (o *SecurityLogForwardingCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/audit/destinations][%d] security_log_forwarding_create default %s", o._statusCode, payload)
}

func (o *SecurityLogForwardingCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityLogForwardingCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
