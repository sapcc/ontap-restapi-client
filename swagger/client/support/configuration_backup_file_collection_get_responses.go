// Code generated by go-swagger; DO NOT EDIT.

package support

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

// ConfigurationBackupFileCollectionGetReader is a Reader for the ConfigurationBackupFileCollectionGet structure.
type ConfigurationBackupFileCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigurationBackupFileCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigurationBackupFileCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConfigurationBackupFileCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConfigurationBackupFileCollectionGetOK creates a ConfigurationBackupFileCollectionGetOK with default headers values
func NewConfigurationBackupFileCollectionGetOK() *ConfigurationBackupFileCollectionGetOK {
	return &ConfigurationBackupFileCollectionGetOK{}
}

/*
ConfigurationBackupFileCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ConfigurationBackupFileCollectionGetOK struct {
	Payload *models.ConfigurationBackupFileResponse
}

// IsSuccess returns true when this configuration backup file collection get o k response has a 2xx status code
func (o *ConfigurationBackupFileCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this configuration backup file collection get o k response has a 3xx status code
func (o *ConfigurationBackupFileCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this configuration backup file collection get o k response has a 4xx status code
func (o *ConfigurationBackupFileCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this configuration backup file collection get o k response has a 5xx status code
func (o *ConfigurationBackupFileCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this configuration backup file collection get o k response a status code equal to that given
func (o *ConfigurationBackupFileCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the configuration backup file collection get o k response
func (o *ConfigurationBackupFileCollectionGetOK) Code() int {
	return 200
}

func (o *ConfigurationBackupFileCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/configuration-backup/backups][%d] configurationBackupFileCollectionGetOK %s", 200, payload)
}

func (o *ConfigurationBackupFileCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/configuration-backup/backups][%d] configurationBackupFileCollectionGetOK %s", 200, payload)
}

func (o *ConfigurationBackupFileCollectionGetOK) GetPayload() *models.ConfigurationBackupFileResponse {
	return o.Payload
}

func (o *ConfigurationBackupFileCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigurationBackupFileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigurationBackupFileCollectionGetDefault creates a ConfigurationBackupFileCollectionGetDefault with default headers values
func NewConfigurationBackupFileCollectionGetDefault(code int) *ConfigurationBackupFileCollectionGetDefault {
	return &ConfigurationBackupFileCollectionGetDefault{
		_statusCode: code,
	}
}

/*
ConfigurationBackupFileCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type ConfigurationBackupFileCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this configuration backup file collection get default response has a 2xx status code
func (o *ConfigurationBackupFileCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this configuration backup file collection get default response has a 3xx status code
func (o *ConfigurationBackupFileCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this configuration backup file collection get default response has a 4xx status code
func (o *ConfigurationBackupFileCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this configuration backup file collection get default response has a 5xx status code
func (o *ConfigurationBackupFileCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this configuration backup file collection get default response a status code equal to that given
func (o *ConfigurationBackupFileCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the configuration backup file collection get default response
func (o *ConfigurationBackupFileCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *ConfigurationBackupFileCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/configuration-backup/backups][%d] configuration_backup_file_collection_get default %s", o._statusCode, payload)
}

func (o *ConfigurationBackupFileCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/configuration-backup/backups][%d] configuration_backup_file_collection_get default %s", o._statusCode, payload)
}

func (o *ConfigurationBackupFileCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConfigurationBackupFileCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
