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

// SecurityCertificateCreateReader is a Reader for the SecurityCertificateCreate structure.
type SecurityCertificateCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityCertificateCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSecurityCertificateCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityCertificateCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityCertificateCreateCreated creates a SecurityCertificateCreateCreated with default headers values
func NewSecurityCertificateCreateCreated() *SecurityCertificateCreateCreated {
	return &SecurityCertificateCreateCreated{}
}

/*
SecurityCertificateCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SecurityCertificateCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.SecurityCertificateResponse
}

// IsSuccess returns true when this security certificate create created response has a 2xx status code
func (o *SecurityCertificateCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security certificate create created response has a 3xx status code
func (o *SecurityCertificateCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security certificate create created response has a 4xx status code
func (o *SecurityCertificateCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this security certificate create created response has a 5xx status code
func (o *SecurityCertificateCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this security certificate create created response a status code equal to that given
func (o *SecurityCertificateCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the security certificate create created response
func (o *SecurityCertificateCreateCreated) Code() int {
	return 201
}

func (o *SecurityCertificateCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/certificates][%d] securityCertificateCreateCreated %s", 201, payload)
}

func (o *SecurityCertificateCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/certificates][%d] securityCertificateCreateCreated %s", 201, payload)
}

func (o *SecurityCertificateCreateCreated) GetPayload() *models.SecurityCertificateResponse {
	return o.Payload
}

func (o *SecurityCertificateCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.SecurityCertificateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityCertificateCreateDefault creates a SecurityCertificateCreateDefault with default headers values
func NewSecurityCertificateCreateDefault(code int) *SecurityCertificateCreateDefault {
	return &SecurityCertificateCreateDefault{
		_statusCode: code,
	}
}

/*
	SecurityCertificateCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 3735645    |  Cannot specify a value for serial. It is generated automatically. |
| 3735622    |  The certificate type is not supported. |
| 3735664    |  The specified key size is not supported in FIPS mode. |
| 3735665    |  The specified hash function is not supported in FIPS mode. |
| 3735553    |  Failed to create self-signed Certificate. |
| 3735646    |  Failed to store the certificates. |
| 3735693    |  The certificate installation failed as private key was empty. |
| 3735618    |  Cannot accept private key for server_ca or client_ca. |
| 52363365   |  Failed to allocate memory. |
| 52559975   |  Failed to read the certificate due to incorrect formatting. |
| 52363366   |  Unsupported key type. |
| 52560123   |  Failed to read the key due to incorrect formatting. |
| 52559972   |  The certificates start date is later than the current date. |
| 52559976   |  The certificate and private key do not match. |
| 52559973   |  The certificate has expired. |
| 52363366   |  Logic error: use of a dead object. |
| 3735696    |  Intermediate certificates are not supported with client_ca and server_ca type certificates. |
| 52559974   |  The certificate is not supported in FIPS mode. |
| 3735676    |  Cannot continue the installation without a value for the common name. Since the subject field in the certificate is empty, the field "common_name" must have a value to continue with the installation. |
| 3735558    |  Failed to extract information about Common Name from the certificate. |
| 3735588    |  The common name (CN) extracted from the certificate is not valid. |
| 3735632    |  Failed to extract Certificate Authority Information from the certificate. |
| 3735700    |  The specified key size is not supported. |
| 52560173   |  The hash function is not supported for digital signatures. |
| 3735751    |  Failed to authenticate and fetch the access token from Azure OAuth host. |
| 3735752    |  Failed to extract the private key from the Azure Key Vault certificate.
| 3735753    |  Unsupported content_type in the Azure secrets response.
| 3735754    |  Internal error. Failed to parse the JSON response from Azure Key Vault.
| 3735755    |  REST call to Azure failed.
| 3735756    |  Invalid client certificate.
| 3735757    |  Internal error. Failed to generate client assertion.
| 3735762    |  Provided Azure Key Vault configuration is incorrect.
| 3735763    |  Provided Azure Key Vault configuration is incomplete.
| 3735764    |  Request to Azure failed. Reason - Azure error code and Azure error message.
*/
type SecurityCertificateCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security certificate create default response has a 2xx status code
func (o *SecurityCertificateCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security certificate create default response has a 3xx status code
func (o *SecurityCertificateCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security certificate create default response has a 4xx status code
func (o *SecurityCertificateCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security certificate create default response has a 5xx status code
func (o *SecurityCertificateCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security certificate create default response a status code equal to that given
func (o *SecurityCertificateCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security certificate create default response
func (o *SecurityCertificateCreateDefault) Code() int {
	return o._statusCode
}

func (o *SecurityCertificateCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/certificates][%d] security_certificate_create default %s", o._statusCode, payload)
}

func (o *SecurityCertificateCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/certificates][%d] security_certificate_create default %s", o._statusCode, payload)
}

func (o *SecurityCertificateCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityCertificateCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
