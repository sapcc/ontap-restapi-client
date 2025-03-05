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

// SecurityKeystoreGetReader is a Reader for the SecurityKeystoreGet structure.
type SecurityKeystoreGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeystoreGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityKeystoreGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeystoreGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeystoreGetOK creates a SecurityKeystoreGetOK with default headers values
func NewSecurityKeystoreGetOK() *SecurityKeystoreGetOK {
	return &SecurityKeystoreGetOK{}
}

/*
SecurityKeystoreGetOK describes a response with status code 200, with default header values.

OK
*/
type SecurityKeystoreGetOK struct {
	Payload *models.SecurityKeystore
}

// IsSuccess returns true when this security keystore get o k response has a 2xx status code
func (o *SecurityKeystoreGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security keystore get o k response has a 3xx status code
func (o *SecurityKeystoreGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security keystore get o k response has a 4xx status code
func (o *SecurityKeystoreGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security keystore get o k response has a 5xx status code
func (o *SecurityKeystoreGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security keystore get o k response a status code equal to that given
func (o *SecurityKeystoreGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security keystore get o k response
func (o *SecurityKeystoreGetOK) Code() int {
	return 200
}

func (o *SecurityKeystoreGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-stores/{uuid}][%d] securityKeystoreGetOK %s", 200, payload)
}

func (o *SecurityKeystoreGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-stores/{uuid}][%d] securityKeystoreGetOK %s", 200, payload)
}

func (o *SecurityKeystoreGetOK) GetPayload() *models.SecurityKeystore {
	return o.Payload
}

func (o *SecurityKeystoreGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityKeystore)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityKeystoreGetDefault creates a SecurityKeystoreGetDefault with default headers values
func NewSecurityKeystoreGetDefault(code int) *SecurityKeystoreGetDefault {
	return &SecurityKeystoreGetDefault{
		_statusCode: code,
	}
}

/*
	SecurityKeystoreGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65539523 | Internal error. The keymanager_config table is missing. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SecurityKeystoreGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security keystore get default response has a 2xx status code
func (o *SecurityKeystoreGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security keystore get default response has a 3xx status code
func (o *SecurityKeystoreGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security keystore get default response has a 4xx status code
func (o *SecurityKeystoreGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security keystore get default response has a 5xx status code
func (o *SecurityKeystoreGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security keystore get default response a status code equal to that given
func (o *SecurityKeystoreGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security keystore get default response
func (o *SecurityKeystoreGetDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeystoreGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-stores/{uuid}][%d] security_keystore_get default %s", o._statusCode, payload)
}

func (o *SecurityKeystoreGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-stores/{uuid}][%d] security_keystore_get default %s", o._statusCode, payload)
}

func (o *SecurityKeystoreGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeystoreGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
