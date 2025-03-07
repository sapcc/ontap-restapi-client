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

// AzureKeyVaultCollectionGetReader is a Reader for the AzureKeyVaultCollectionGet structure.
type AzureKeyVaultCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AzureKeyVaultCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAzureKeyVaultCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAzureKeyVaultCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAzureKeyVaultCollectionGetOK creates a AzureKeyVaultCollectionGetOK with default headers values
func NewAzureKeyVaultCollectionGetOK() *AzureKeyVaultCollectionGetOK {
	return &AzureKeyVaultCollectionGetOK{}
}

/*
AzureKeyVaultCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type AzureKeyVaultCollectionGetOK struct {
	Payload *models.AzureKeyVaultResponse
}

// IsSuccess returns true when this azure key vault collection get o k response has a 2xx status code
func (o *AzureKeyVaultCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this azure key vault collection get o k response has a 3xx status code
func (o *AzureKeyVaultCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure key vault collection get o k response has a 4xx status code
func (o *AzureKeyVaultCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure key vault collection get o k response has a 5xx status code
func (o *AzureKeyVaultCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this azure key vault collection get o k response a status code equal to that given
func (o *AzureKeyVaultCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the azure key vault collection get o k response
func (o *AzureKeyVaultCollectionGetOK) Code() int {
	return 200
}

func (o *AzureKeyVaultCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/azure-key-vaults][%d] azureKeyVaultCollectionGetOK %s", 200, payload)
}

func (o *AzureKeyVaultCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/azure-key-vaults][%d] azureKeyVaultCollectionGetOK %s", 200, payload)
}

func (o *AzureKeyVaultCollectionGetOK) GetPayload() *models.AzureKeyVaultResponse {
	return o.Payload
}

func (o *AzureKeyVaultCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureKeyVaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureKeyVaultCollectionGetDefault creates a AzureKeyVaultCollectionGetDefault with default headers values
func NewAzureKeyVaultCollectionGetDefault(code int) *AzureKeyVaultCollectionGetDefault {
	return &AzureKeyVaultCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	AzureKeyVaultCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65537549 | The Azure Key Vault Key Management Service is unreachable from one or more nodes. |
| 65537551 | Top-level internal key protection key (KEK) unavailable on one or more nodes. |
| 65537552 | Embedded KMIP server status not available. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AzureKeyVaultCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this azure key vault collection get default response has a 2xx status code
func (o *AzureKeyVaultCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this azure key vault collection get default response has a 3xx status code
func (o *AzureKeyVaultCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this azure key vault collection get default response has a 4xx status code
func (o *AzureKeyVaultCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this azure key vault collection get default response has a 5xx status code
func (o *AzureKeyVaultCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this azure key vault collection get default response a status code equal to that given
func (o *AzureKeyVaultCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the azure key vault collection get default response
func (o *AzureKeyVaultCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *AzureKeyVaultCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/azure-key-vaults][%d] azure_key_vault_collection_get default %s", o._statusCode, payload)
}

func (o *AzureKeyVaultCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/azure-key-vaults][%d] azure_key_vault_collection_get default %s", o._statusCode, payload)
}

func (o *AzureKeyVaultCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AzureKeyVaultCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
