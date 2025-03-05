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

// AzureKeyVaultRestoreReader is a Reader for the AzureKeyVaultRestore structure.
type AzureKeyVaultRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AzureKeyVaultRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAzureKeyVaultRestoreCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewAzureKeyVaultRestoreAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAzureKeyVaultRestoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAzureKeyVaultRestoreCreated creates a AzureKeyVaultRestoreCreated with default headers values
func NewAzureKeyVaultRestoreCreated() *AzureKeyVaultRestoreCreated {
	return &AzureKeyVaultRestoreCreated{}
}

/*
AzureKeyVaultRestoreCreated describes a response with status code 201, with default header values.

Created
*/
type AzureKeyVaultRestoreCreated struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this azure key vault restore created response has a 2xx status code
func (o *AzureKeyVaultRestoreCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this azure key vault restore created response has a 3xx status code
func (o *AzureKeyVaultRestoreCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure key vault restore created response has a 4xx status code
func (o *AzureKeyVaultRestoreCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure key vault restore created response has a 5xx status code
func (o *AzureKeyVaultRestoreCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this azure key vault restore created response a status code equal to that given
func (o *AzureKeyVaultRestoreCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the azure key vault restore created response
func (o *AzureKeyVaultRestoreCreated) Code() int {
	return 201
}

func (o *AzureKeyVaultRestoreCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/azure-key-vaults/{uuid}/restore][%d] azureKeyVaultRestoreCreated %s", 201, payload)
}

func (o *AzureKeyVaultRestoreCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/azure-key-vaults/{uuid}/restore][%d] azureKeyVaultRestoreCreated %s", 201, payload)
}

func (o *AzureKeyVaultRestoreCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *AzureKeyVaultRestoreCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureKeyVaultRestoreAccepted creates a AzureKeyVaultRestoreAccepted with default headers values
func NewAzureKeyVaultRestoreAccepted() *AzureKeyVaultRestoreAccepted {
	return &AzureKeyVaultRestoreAccepted{}
}

/*
AzureKeyVaultRestoreAccepted describes a response with status code 202, with default header values.

Accepted
*/
type AzureKeyVaultRestoreAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this azure key vault restore accepted response has a 2xx status code
func (o *AzureKeyVaultRestoreAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this azure key vault restore accepted response has a 3xx status code
func (o *AzureKeyVaultRestoreAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure key vault restore accepted response has a 4xx status code
func (o *AzureKeyVaultRestoreAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure key vault restore accepted response has a 5xx status code
func (o *AzureKeyVaultRestoreAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this azure key vault restore accepted response a status code equal to that given
func (o *AzureKeyVaultRestoreAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the azure key vault restore accepted response
func (o *AzureKeyVaultRestoreAccepted) Code() int {
	return 202
}

func (o *AzureKeyVaultRestoreAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/azure-key-vaults/{uuid}/restore][%d] azureKeyVaultRestoreAccepted %s", 202, payload)
}

func (o *AzureKeyVaultRestoreAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/azure-key-vaults/{uuid}/restore][%d] azureKeyVaultRestoreAccepted %s", 202, payload)
}

func (o *AzureKeyVaultRestoreAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *AzureKeyVaultRestoreAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureKeyVaultRestoreDefault creates a AzureKeyVaultRestoreDefault with default headers values
func NewAzureKeyVaultRestoreDefault(code int) *AzureKeyVaultRestoreDefault {
	return &AzureKeyVaultRestoreDefault{
		_statusCode: code,
	}
}

/*
	AzureKeyVaultRestoreDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65537120 | Azure Key Vault is not configured for the given SVM. |
| 65537515 | Failed to restore keys on some nodes in the cluster. |
| 65537544 | Missing wrapped top-level internal key protection key (KEK) from internal database. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AzureKeyVaultRestoreDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this azure key vault restore default response has a 2xx status code
func (o *AzureKeyVaultRestoreDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this azure key vault restore default response has a 3xx status code
func (o *AzureKeyVaultRestoreDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this azure key vault restore default response has a 4xx status code
func (o *AzureKeyVaultRestoreDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this azure key vault restore default response has a 5xx status code
func (o *AzureKeyVaultRestoreDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this azure key vault restore default response a status code equal to that given
func (o *AzureKeyVaultRestoreDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the azure key vault restore default response
func (o *AzureKeyVaultRestoreDefault) Code() int {
	return o._statusCode
}

func (o *AzureKeyVaultRestoreDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/azure-key-vaults/{uuid}/restore][%d] azure_key_vault_restore default %s", o._statusCode, payload)
}

func (o *AzureKeyVaultRestoreDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/azure-key-vaults/{uuid}/restore][%d] azure_key_vault_restore default %s", o._statusCode, payload)
}

func (o *AzureKeyVaultRestoreDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AzureKeyVaultRestoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
