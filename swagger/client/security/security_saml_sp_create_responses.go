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

// SecuritySamlSpCreateReader is a Reader for the SecuritySamlSpCreate structure.
type SecuritySamlSpCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecuritySamlSpCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSecuritySamlSpCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSecuritySamlSpCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecuritySamlSpCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecuritySamlSpCreateCreated creates a SecuritySamlSpCreateCreated with default headers values
func NewSecuritySamlSpCreateCreated() *SecuritySamlSpCreateCreated {
	return &SecuritySamlSpCreateCreated{}
}

/*
SecuritySamlSpCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SecuritySamlSpCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this security saml sp create created response has a 2xx status code
func (o *SecuritySamlSpCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security saml sp create created response has a 3xx status code
func (o *SecuritySamlSpCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security saml sp create created response has a 4xx status code
func (o *SecuritySamlSpCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this security saml sp create created response has a 5xx status code
func (o *SecuritySamlSpCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this security saml sp create created response a status code equal to that given
func (o *SecuritySamlSpCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the security saml sp create created response
func (o *SecuritySamlSpCreateCreated) Code() int {
	return 201
}

func (o *SecuritySamlSpCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/cluster/saml-sp][%d] securitySamlSpCreateCreated %s", 201, payload)
}

func (o *SecuritySamlSpCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/cluster/saml-sp][%d] securitySamlSpCreateCreated %s", 201, payload)
}

func (o *SecuritySamlSpCreateCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SecuritySamlSpCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecuritySamlSpCreateAccepted creates a SecuritySamlSpCreateAccepted with default headers values
func NewSecuritySamlSpCreateAccepted() *SecuritySamlSpCreateAccepted {
	return &SecuritySamlSpCreateAccepted{}
}

/*
SecuritySamlSpCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SecuritySamlSpCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this security saml sp create accepted response has a 2xx status code
func (o *SecuritySamlSpCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security saml sp create accepted response has a 3xx status code
func (o *SecuritySamlSpCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security saml sp create accepted response has a 4xx status code
func (o *SecuritySamlSpCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this security saml sp create accepted response has a 5xx status code
func (o *SecuritySamlSpCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this security saml sp create accepted response a status code equal to that given
func (o *SecuritySamlSpCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the security saml sp create accepted response
func (o *SecuritySamlSpCreateAccepted) Code() int {
	return 202
}

func (o *SecuritySamlSpCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/cluster/saml-sp][%d] securitySamlSpCreateAccepted %s", 202, payload)
}

func (o *SecuritySamlSpCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/cluster/saml-sp][%d] securitySamlSpCreateAccepted %s", 202, payload)
}

func (o *SecuritySamlSpCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SecuritySamlSpCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecuritySamlSpCreateDefault creates a SecuritySamlSpCreateDefault with default headers values
func NewSecuritySamlSpCreateDefault(code int) *SecuritySamlSpCreateDefault {
	return &SecuritySamlSpCreateDefault{
		_statusCode: code,
	}
}

/*
	SecuritySamlSpCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 12320789 | Failed to download data file from specified URI. |
| 12320794 | The host parameter provided must be the cluster management interface's IP address.  If the cluster management interface is not available, the node management interface's IP address must be used. |
| 12320795 | A valid cluster or node management interface IP address must be provided. |
| 12320805 | The certificate information provided does not match any installed certificates. |
| 12320806 | The certificate information entered does not match any installed certificates. |
| 12320814 | An invalid IDP URI has been entered. |
| 12320815 | An IDP URI must be an HTTPS or FTPS URI. |
| 12320819 | Use the HTTPS scheme for the “idp_uri“ or set “verify_metadata_server“ to false. |
| 12320820 | No certificate is installed with the specified “certificate.ca“ and “certificate.serial“. |
| 12320821 | No certificate is installed with the specified “certificate.common_name“. |
| 12320823 | The host parameter provided must be the cluster management interface's IP address. If the cluster management interface is not available, the node management interface's IP address must be used. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SecuritySamlSpCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security saml sp create default response has a 2xx status code
func (o *SecuritySamlSpCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security saml sp create default response has a 3xx status code
func (o *SecuritySamlSpCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security saml sp create default response has a 4xx status code
func (o *SecuritySamlSpCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security saml sp create default response has a 5xx status code
func (o *SecuritySamlSpCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security saml sp create default response a status code equal to that given
func (o *SecuritySamlSpCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security saml sp create default response
func (o *SecuritySamlSpCreateDefault) Code() int {
	return o._statusCode
}

func (o *SecuritySamlSpCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/cluster/saml-sp][%d] security_saml_sp_create default %s", o._statusCode, payload)
}

func (o *SecuritySamlSpCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/cluster/saml-sp][%d] security_saml_sp_create default %s", o._statusCode, payload)
}

func (o *SecuritySamlSpCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecuritySamlSpCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
