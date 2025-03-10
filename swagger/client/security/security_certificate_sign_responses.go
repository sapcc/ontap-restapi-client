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

// SecurityCertificateSignReader is a Reader for the SecurityCertificateSign structure.
type SecurityCertificateSignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityCertificateSignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityCertificateSignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityCertificateSignDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityCertificateSignOK creates a SecurityCertificateSignOK with default headers values
func NewSecurityCertificateSignOK() *SecurityCertificateSignOK {
	return &SecurityCertificateSignOK{}
}

/*
SecurityCertificateSignOK describes a response with status code 200, with default header values.

OK
*/
type SecurityCertificateSignOK struct {
	Payload *models.SecurityCertificateSignResponse
}

// IsSuccess returns true when this security certificate sign o k response has a 2xx status code
func (o *SecurityCertificateSignOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security certificate sign o k response has a 3xx status code
func (o *SecurityCertificateSignOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security certificate sign o k response has a 4xx status code
func (o *SecurityCertificateSignOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security certificate sign o k response has a 5xx status code
func (o *SecurityCertificateSignOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security certificate sign o k response a status code equal to that given
func (o *SecurityCertificateSignOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security certificate sign o k response
func (o *SecurityCertificateSignOK) Code() int {
	return 200
}

func (o *SecurityCertificateSignOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/certificates/{ca.uuid}/sign][%d] securityCertificateSignOK %s", 200, payload)
}

func (o *SecurityCertificateSignOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/certificates/{ca.uuid}/sign][%d] securityCertificateSignOK %s", 200, payload)
}

func (o *SecurityCertificateSignOK) GetPayload() *models.SecurityCertificateSignResponse {
	return o.Payload
}

func (o *SecurityCertificateSignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityCertificateSignResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityCertificateSignDefault creates a SecurityCertificateSignDefault with default headers values
func NewSecurityCertificateSignDefault(code int) *SecurityCertificateSignDefault {
	return &SecurityCertificateSignDefault{
		_statusCode: code,
	}
}

/*
	SecurityCertificateSignDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 3735628    | Failed to use CA certificate for signing. |
| 3735665    | The specified hash function is not supported in FIPS mode. |
| 52559974   | The certificate is not supported in FIPS mode. |
| 3735626    | Failed to generate signed Certificate. |
| 3735558    | Failed to extract information about Common Name from the certificate. |
| 3735588    | The common name (CN) extracted from the certificate is not valid. |
| 3735632    | Failed to extract Certificate Authority Information from the certificate. |
| 3735629    | Failed to sign the certificate because Common Name of signing certificate and Common Name of CA certificate are same. |
| 3735630    | Failed to sign the certificate because expiry date of signing certificate exceeds the expiry date of CA certificate. |
| 3735701    | Invalid expiration period. The allowed range for expiration time is between 1 and 3652 days. |
*/
type SecurityCertificateSignDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security certificate sign default response has a 2xx status code
func (o *SecurityCertificateSignDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security certificate sign default response has a 3xx status code
func (o *SecurityCertificateSignDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security certificate sign default response has a 4xx status code
func (o *SecurityCertificateSignDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security certificate sign default response has a 5xx status code
func (o *SecurityCertificateSignDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security certificate sign default response a status code equal to that given
func (o *SecurityCertificateSignDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security certificate sign default response
func (o *SecurityCertificateSignDefault) Code() int {
	return o._statusCode
}

func (o *SecurityCertificateSignDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/certificates/{ca.uuid}/sign][%d] security_certificate_sign default %s", o._statusCode, payload)
}

func (o *SecurityCertificateSignDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/certificates/{ca.uuid}/sign][%d] security_certificate_sign default %s", o._statusCode, payload)
}

func (o *SecurityCertificateSignDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityCertificateSignDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
