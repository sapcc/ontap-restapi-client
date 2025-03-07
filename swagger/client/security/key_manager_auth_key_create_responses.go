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

// KeyManagerAuthKeyCreateReader is a Reader for the KeyManagerAuthKeyCreate structure.
type KeyManagerAuthKeyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyManagerAuthKeyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewKeyManagerAuthKeyCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKeyManagerAuthKeyCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKeyManagerAuthKeyCreateCreated creates a KeyManagerAuthKeyCreateCreated with default headers values
func NewKeyManagerAuthKeyCreateCreated() *KeyManagerAuthKeyCreateCreated {
	return &KeyManagerAuthKeyCreateCreated{}
}

/*
KeyManagerAuthKeyCreateCreated describes a response with status code 201, with default header values.

Created
*/
type KeyManagerAuthKeyCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.KeyManagerAuthKeyResponse
}

// IsSuccess returns true when this key manager auth key create created response has a 2xx status code
func (o *KeyManagerAuthKeyCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this key manager auth key create created response has a 3xx status code
func (o *KeyManagerAuthKeyCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this key manager auth key create created response has a 4xx status code
func (o *KeyManagerAuthKeyCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this key manager auth key create created response has a 5xx status code
func (o *KeyManagerAuthKeyCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this key manager auth key create created response a status code equal to that given
func (o *KeyManagerAuthKeyCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the key manager auth key create created response
func (o *KeyManagerAuthKeyCreateCreated) Code() int {
	return 201
}

func (o *KeyManagerAuthKeyCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{security_key_manager.uuid}/auth-keys][%d] keyManagerAuthKeyCreateCreated %s", 201, payload)
}

func (o *KeyManagerAuthKeyCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{security_key_manager.uuid}/auth-keys][%d] keyManagerAuthKeyCreateCreated %s", 201, payload)
}

func (o *KeyManagerAuthKeyCreateCreated) GetPayload() *models.KeyManagerAuthKeyResponse {
	return o.Payload
}

func (o *KeyManagerAuthKeyCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.KeyManagerAuthKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyManagerAuthKeyCreateDefault creates a KeyManagerAuthKeyCreateDefault with default headers values
func NewKeyManagerAuthKeyCreateDefault(code int) *KeyManagerAuthKeyCreateDefault {
	return &KeyManagerAuthKeyCreateDefault{
		_statusCode: code,
	}
}

/*
	KeyManagerAuthKeyCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65536048 | The maximum number of authentication keys are already configured. |
| 65536053 | Invalid passphrase length; passphrase must be 20-32 ASCII-range characters. |
| 65536056 | The key tag value provided contains invalid characters. |
| 65536056 | The key-tag option cannot contain any spaces, tabs or new lines. |
| 65536076 | Failed to push authentication key to any registered key servers. |
| 65536160 | Unable to determine the current number of configured authentication keys. |
| 65536264 | Failed to create authentication key. |
| 65536265 | Failed to create a key-id for the authentication key. |
| 65536828 | External key management is not enabled for the SVM. |
| 65536856 | No key servers found for the SVM. |
| 65536872 | Error cleaning up key database after key creation error. |
| 65536896 | External key management is not configured on the partner site. |
| 65536987 | One or more key servers are unavailable. |
| 65538800 | External KMIP DKMIP keymanager not configured on administrative Vserver. |
| 65538801 | Internal error while accessing keymanager database. |
| 65538802 | The UUID provided is not associated with the administrator SVM key manager. |
| 66060289 | Failed to store authentication key on key server. |
| 66060304 | Invalid key length. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type KeyManagerAuthKeyCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this key manager auth key create default response has a 2xx status code
func (o *KeyManagerAuthKeyCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this key manager auth key create default response has a 3xx status code
func (o *KeyManagerAuthKeyCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this key manager auth key create default response has a 4xx status code
func (o *KeyManagerAuthKeyCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this key manager auth key create default response has a 5xx status code
func (o *KeyManagerAuthKeyCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this key manager auth key create default response a status code equal to that given
func (o *KeyManagerAuthKeyCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the key manager auth key create default response
func (o *KeyManagerAuthKeyCreateDefault) Code() int {
	return o._statusCode
}

func (o *KeyManagerAuthKeyCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{security_key_manager.uuid}/auth-keys][%d] key_manager_auth_key_create default %s", o._statusCode, payload)
}

func (o *KeyManagerAuthKeyCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{security_key_manager.uuid}/auth-keys][%d] key_manager_auth_key_create default %s", o._statusCode, payload)
}

func (o *KeyManagerAuthKeyCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *KeyManagerAuthKeyCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
