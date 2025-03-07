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

// AzureKeyVaultRekeyInternalReader is a Reader for the AzureKeyVaultRekeyInternal structure.
type AzureKeyVaultRekeyInternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AzureKeyVaultRekeyInternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAzureKeyVaultRekeyInternalCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewAzureKeyVaultRekeyInternalAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAzureKeyVaultRekeyInternalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAzureKeyVaultRekeyInternalCreated creates a AzureKeyVaultRekeyInternalCreated with default headers values
func NewAzureKeyVaultRekeyInternalCreated() *AzureKeyVaultRekeyInternalCreated {
	return &AzureKeyVaultRekeyInternalCreated{}
}

/*
AzureKeyVaultRekeyInternalCreated describes a response with status code 201, with default header values.

Created
*/
type AzureKeyVaultRekeyInternalCreated struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this azure key vault rekey internal created response has a 2xx status code
func (o *AzureKeyVaultRekeyInternalCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this azure key vault rekey internal created response has a 3xx status code
func (o *AzureKeyVaultRekeyInternalCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure key vault rekey internal created response has a 4xx status code
func (o *AzureKeyVaultRekeyInternalCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure key vault rekey internal created response has a 5xx status code
func (o *AzureKeyVaultRekeyInternalCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this azure key vault rekey internal created response a status code equal to that given
func (o *AzureKeyVaultRekeyInternalCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the azure key vault rekey internal created response
func (o *AzureKeyVaultRekeyInternalCreated) Code() int {
	return 201
}

func (o *AzureKeyVaultRekeyInternalCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/azure-key-vaults/{uuid}/rekey-internal][%d] azureKeyVaultRekeyInternalCreated %s", 201, payload)
}

func (o *AzureKeyVaultRekeyInternalCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/azure-key-vaults/{uuid}/rekey-internal][%d] azureKeyVaultRekeyInternalCreated %s", 201, payload)
}

func (o *AzureKeyVaultRekeyInternalCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *AzureKeyVaultRekeyInternalCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureKeyVaultRekeyInternalAccepted creates a AzureKeyVaultRekeyInternalAccepted with default headers values
func NewAzureKeyVaultRekeyInternalAccepted() *AzureKeyVaultRekeyInternalAccepted {
	return &AzureKeyVaultRekeyInternalAccepted{}
}

/*
AzureKeyVaultRekeyInternalAccepted describes a response with status code 202, with default header values.

Accepted
*/
type AzureKeyVaultRekeyInternalAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this azure key vault rekey internal accepted response has a 2xx status code
func (o *AzureKeyVaultRekeyInternalAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this azure key vault rekey internal accepted response has a 3xx status code
func (o *AzureKeyVaultRekeyInternalAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure key vault rekey internal accepted response has a 4xx status code
func (o *AzureKeyVaultRekeyInternalAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure key vault rekey internal accepted response has a 5xx status code
func (o *AzureKeyVaultRekeyInternalAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this azure key vault rekey internal accepted response a status code equal to that given
func (o *AzureKeyVaultRekeyInternalAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the azure key vault rekey internal accepted response
func (o *AzureKeyVaultRekeyInternalAccepted) Code() int {
	return 202
}

func (o *AzureKeyVaultRekeyInternalAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/azure-key-vaults/{uuid}/rekey-internal][%d] azureKeyVaultRekeyInternalAccepted %s", 202, payload)
}

func (o *AzureKeyVaultRekeyInternalAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/azure-key-vaults/{uuid}/rekey-internal][%d] azureKeyVaultRekeyInternalAccepted %s", 202, payload)
}

func (o *AzureKeyVaultRekeyInternalAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *AzureKeyVaultRekeyInternalAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureKeyVaultRekeyInternalDefault creates a AzureKeyVaultRekeyInternalDefault with default headers values
func NewAzureKeyVaultRekeyInternalDefault(code int) *AzureKeyVaultRekeyInternalDefault {
	return &AzureKeyVaultRekeyInternalDefault{
		_statusCode: code,
	}
}

/*
	AzureKeyVaultRekeyInternalDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65537120 | Azure Key Vault is not configured for the given SVM. |
| 65537547 | One or more volume encryption keys for encrypted volumes of this data SVM are stored in the key manager configured for the admin SVM. Use the REST API POST method to migrate this data SVM's keys from the admin SVM's key manager to this data SVM's key manager before running the rekey operation. |
| 65537559 | There are no existing internal keys for the SVM. A rekey operation is allowed for an SVM with one or more encryption keys. |
| 65537610 | Rekey cannot be performed on the SVM while the enabled keystore configuration is being switched. If a previous attempt to switch the keystore configuration failed, or was interrupted, the system will continue to prevent rekeying for the SVM. Use the REST API PATCH method "/api/security/key-stores/{uuid}" to re-run and complete the operation. |
| 65539436 | Rekey cannot be performed on the SVM while the enabled keystore configuration is being initialized. Wait until the keystore is in the active state, and rerun the rekey operation. |
| 65539437 | Rekey cannot be performed on the SVM while the enabled keystore configuration is being disabled. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AzureKeyVaultRekeyInternalDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this azure key vault rekey internal default response has a 2xx status code
func (o *AzureKeyVaultRekeyInternalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this azure key vault rekey internal default response has a 3xx status code
func (o *AzureKeyVaultRekeyInternalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this azure key vault rekey internal default response has a 4xx status code
func (o *AzureKeyVaultRekeyInternalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this azure key vault rekey internal default response has a 5xx status code
func (o *AzureKeyVaultRekeyInternalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this azure key vault rekey internal default response a status code equal to that given
func (o *AzureKeyVaultRekeyInternalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the azure key vault rekey internal default response
func (o *AzureKeyVaultRekeyInternalDefault) Code() int {
	return o._statusCode
}

func (o *AzureKeyVaultRekeyInternalDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/azure-key-vaults/{uuid}/rekey-internal][%d] azure_key_vault_rekey_internal default %s", o._statusCode, payload)
}

func (o *AzureKeyVaultRekeyInternalDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/azure-key-vaults/{uuid}/rekey-internal][%d] azure_key_vault_rekey_internal default %s", o._statusCode, payload)
}

func (o *AzureKeyVaultRekeyInternalDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AzureKeyVaultRekeyInternalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
