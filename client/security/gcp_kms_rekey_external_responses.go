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

// GcpKmsRekeyExternalReader is a Reader for the GcpKmsRekeyExternal structure.
type GcpKmsRekeyExternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GcpKmsRekeyExternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGcpKmsRekeyExternalCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewGcpKmsRekeyExternalAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGcpKmsRekeyExternalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGcpKmsRekeyExternalCreated creates a GcpKmsRekeyExternalCreated with default headers values
func NewGcpKmsRekeyExternalCreated() *GcpKmsRekeyExternalCreated {
	return &GcpKmsRekeyExternalCreated{}
}

/*
GcpKmsRekeyExternalCreated describes a response with status code 201, with default header values.

Created
*/
type GcpKmsRekeyExternalCreated struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this gcp kms rekey external created response has a 2xx status code
func (o *GcpKmsRekeyExternalCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this gcp kms rekey external created response has a 3xx status code
func (o *GcpKmsRekeyExternalCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gcp kms rekey external created response has a 4xx status code
func (o *GcpKmsRekeyExternalCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this gcp kms rekey external created response has a 5xx status code
func (o *GcpKmsRekeyExternalCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this gcp kms rekey external created response a status code equal to that given
func (o *GcpKmsRekeyExternalCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the gcp kms rekey external created response
func (o *GcpKmsRekeyExternalCreated) Code() int {
	return 201
}

func (o *GcpKmsRekeyExternalCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms/{gcp_kms.uuid}/rekey-external][%d] gcpKmsRekeyExternalCreated %s", 201, payload)
}

func (o *GcpKmsRekeyExternalCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms/{gcp_kms.uuid}/rekey-external][%d] gcpKmsRekeyExternalCreated %s", 201, payload)
}

func (o *GcpKmsRekeyExternalCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *GcpKmsRekeyExternalCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGcpKmsRekeyExternalAccepted creates a GcpKmsRekeyExternalAccepted with default headers values
func NewGcpKmsRekeyExternalAccepted() *GcpKmsRekeyExternalAccepted {
	return &GcpKmsRekeyExternalAccepted{}
}

/*
GcpKmsRekeyExternalAccepted describes a response with status code 202, with default header values.

Accepted
*/
type GcpKmsRekeyExternalAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this gcp kms rekey external accepted response has a 2xx status code
func (o *GcpKmsRekeyExternalAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this gcp kms rekey external accepted response has a 3xx status code
func (o *GcpKmsRekeyExternalAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gcp kms rekey external accepted response has a 4xx status code
func (o *GcpKmsRekeyExternalAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this gcp kms rekey external accepted response has a 5xx status code
func (o *GcpKmsRekeyExternalAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this gcp kms rekey external accepted response a status code equal to that given
func (o *GcpKmsRekeyExternalAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the gcp kms rekey external accepted response
func (o *GcpKmsRekeyExternalAccepted) Code() int {
	return 202
}

func (o *GcpKmsRekeyExternalAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms/{gcp_kms.uuid}/rekey-external][%d] gcpKmsRekeyExternalAccepted %s", 202, payload)
}

func (o *GcpKmsRekeyExternalAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms/{gcp_kms.uuid}/rekey-external][%d] gcpKmsRekeyExternalAccepted %s", 202, payload)
}

func (o *GcpKmsRekeyExternalAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *GcpKmsRekeyExternalAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGcpKmsRekeyExternalDefault creates a GcpKmsRekeyExternalDefault with default headers values
func NewGcpKmsRekeyExternalDefault(code int) *GcpKmsRekeyExternalDefault {
	return &GcpKmsRekeyExternalDefault{
		_statusCode: code,
	}
}

/*
	GcpKmsRekeyExternalDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65537547 | One or more volume encryption keys for encrypted volumes of this data SVM are stored in the key manager configured for the admin SVM. Use the REST API POST method to migrate this data SVM's keys from the admin SVM's key manager to this data SVM's key manager before running the rekey operation. |
| 65537556 | ONTAP is not able to successfully encrypt or decrypt because the configured external key manager for this SVM is in a blocked state. Possible reasons for a blocked state include the top-level external key protection key not found, disabled or having insufficient privileges. |
| 65537610 | Rekey cannot be performed on the SVM while the enabled keystore configuration is being switched. If a previous attempt to switch the keystore configuration failed, or was interrupted, the system will continue to prevent rekeying for the SVM. Use the REST API PATCH method "/api/security/key-stores/{uuid}" to re-run and complete the operation. |
| 65537721 | Google Cloud KMS is not configured for the given SVM. |
| 65537729 | External rekey failed on one or more nodes. Use the REST API POST method \"/api/security/gcp-kms/{uuid}/rekey-external\" to try the rekey operation again. |
| 65539436 | Rekey cannot be performed on the SVM while the enabled keystore configuration is being initialized. Wait until the keystore is in the active state, and rerun the rekey operation. |
| 65539437 | Rekey cannot be performed on the SVM while the enabled keystore configuration is being disabled. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type GcpKmsRekeyExternalDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this gcp kms rekey external default response has a 2xx status code
func (o *GcpKmsRekeyExternalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this gcp kms rekey external default response has a 3xx status code
func (o *GcpKmsRekeyExternalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this gcp kms rekey external default response has a 4xx status code
func (o *GcpKmsRekeyExternalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this gcp kms rekey external default response has a 5xx status code
func (o *GcpKmsRekeyExternalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this gcp kms rekey external default response a status code equal to that given
func (o *GcpKmsRekeyExternalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the gcp kms rekey external default response
func (o *GcpKmsRekeyExternalDefault) Code() int {
	return o._statusCode
}

func (o *GcpKmsRekeyExternalDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms/{gcp_kms.uuid}/rekey-external][%d] gcp_kms_rekey_external default %s", o._statusCode, payload)
}

func (o *GcpKmsRekeyExternalDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms/{gcp_kms.uuid}/rekey-external][%d] gcp_kms_rekey_external default %s", o._statusCode, payload)
}

func (o *GcpKmsRekeyExternalDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GcpKmsRekeyExternalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
