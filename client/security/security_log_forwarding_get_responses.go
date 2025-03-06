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

	"github.com/sapcc/ontap-restapi-client/models"
)

// SecurityLogForwardingGetReader is a Reader for the SecurityLogForwardingGet structure.
type SecurityLogForwardingGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityLogForwardingGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityLogForwardingGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityLogForwardingGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityLogForwardingGetOK creates a SecurityLogForwardingGetOK with default headers values
func NewSecurityLogForwardingGetOK() *SecurityLogForwardingGetOK {
	return &SecurityLogForwardingGetOK{}
}

/*
SecurityLogForwardingGetOK describes a response with status code 200, with default header values.

OK
*/
type SecurityLogForwardingGetOK struct {
	Payload *models.SecurityAuditLogForward
}

// IsSuccess returns true when this security log forwarding get o k response has a 2xx status code
func (o *SecurityLogForwardingGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security log forwarding get o k response has a 3xx status code
func (o *SecurityLogForwardingGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security log forwarding get o k response has a 4xx status code
func (o *SecurityLogForwardingGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security log forwarding get o k response has a 5xx status code
func (o *SecurityLogForwardingGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security log forwarding get o k response a status code equal to that given
func (o *SecurityLogForwardingGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security log forwarding get o k response
func (o *SecurityLogForwardingGetOK) Code() int {
	return 200
}

func (o *SecurityLogForwardingGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/audit/destinations/{address}/{port}][%d] securityLogForwardingGetOK %s", 200, payload)
}

func (o *SecurityLogForwardingGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/audit/destinations/{address}/{port}][%d] securityLogForwardingGetOK %s", 200, payload)
}

func (o *SecurityLogForwardingGetOK) GetPayload() *models.SecurityAuditLogForward {
	return o.Payload
}

func (o *SecurityLogForwardingGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityAuditLogForward)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityLogForwardingGetDefault creates a SecurityLogForwardingGetDefault with default headers values
func NewSecurityLogForwardingGetDefault(code int) *SecurityLogForwardingGetDefault {
	return &SecurityLogForwardingGetDefault{
		_statusCode: code,
	}
}

/*
SecurityLogForwardingGetDefault describes a response with status code -1, with default header values.

Error
*/
type SecurityLogForwardingGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security log forwarding get default response has a 2xx status code
func (o *SecurityLogForwardingGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security log forwarding get default response has a 3xx status code
func (o *SecurityLogForwardingGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security log forwarding get default response has a 4xx status code
func (o *SecurityLogForwardingGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security log forwarding get default response has a 5xx status code
func (o *SecurityLogForwardingGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security log forwarding get default response a status code equal to that given
func (o *SecurityLogForwardingGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security log forwarding get default response
func (o *SecurityLogForwardingGetDefault) Code() int {
	return o._statusCode
}

func (o *SecurityLogForwardingGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/audit/destinations/{address}/{port}][%d] security_log_forwarding_get default %s", o._statusCode, payload)
}

func (o *SecurityLogForwardingGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/audit/destinations/{address}/{port}][%d] security_log_forwarding_get default %s", o._statusCode, payload)
}

func (o *SecurityLogForwardingGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityLogForwardingGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
