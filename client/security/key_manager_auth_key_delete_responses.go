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

	"github.com/sapcc/ontap-restapi-client/models"
)

// KeyManagerAuthKeyDeleteReader is a Reader for the KeyManagerAuthKeyDelete structure.
type KeyManagerAuthKeyDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyManagerAuthKeyDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyManagerAuthKeyDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKeyManagerAuthKeyDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKeyManagerAuthKeyDeleteOK creates a KeyManagerAuthKeyDeleteOK with default headers values
func NewKeyManagerAuthKeyDeleteOK() *KeyManagerAuthKeyDeleteOK {
	return &KeyManagerAuthKeyDeleteOK{}
}

/*
KeyManagerAuthKeyDeleteOK describes a response with status code 200, with default header values.

OK
*/
type KeyManagerAuthKeyDeleteOK struct {
}

// IsSuccess returns true when this key manager auth key delete o k response has a 2xx status code
func (o *KeyManagerAuthKeyDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this key manager auth key delete o k response has a 3xx status code
func (o *KeyManagerAuthKeyDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this key manager auth key delete o k response has a 4xx status code
func (o *KeyManagerAuthKeyDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this key manager auth key delete o k response has a 5xx status code
func (o *KeyManagerAuthKeyDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this key manager auth key delete o k response a status code equal to that given
func (o *KeyManagerAuthKeyDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the key manager auth key delete o k response
func (o *KeyManagerAuthKeyDeleteOK) Code() int {
	return 200
}

func (o *KeyManagerAuthKeyDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/key-managers/{security_key_manager.uuid}/auth-keys/{key_id}][%d] keyManagerAuthKeyDeleteOK", 200)
}

func (o *KeyManagerAuthKeyDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /security/key-managers/{security_key_manager.uuid}/auth-keys/{key_id}][%d] keyManagerAuthKeyDeleteOK", 200)
}

func (o *KeyManagerAuthKeyDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewKeyManagerAuthKeyDeleteDefault creates a KeyManagerAuthKeyDeleteDefault with default headers values
func NewKeyManagerAuthKeyDeleteDefault(code int) *KeyManagerAuthKeyDeleteDefault {
	return &KeyManagerAuthKeyDeleteDefault{
		_statusCode: code,
	}
}

/*
	KeyManagerAuthKeyDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65536828 | External key management is not enabled for the SVM. |
| 65536859 | Authentication key-Id provided for deletion is in use. |
| 65536860 | Key-id provided for deletion is not an authentication key. |
| 65538800 | External KMIP DKMIP keymanager not configured on administrative Vserver. |
| 65538801 | Internal error while accessing keymanager database. |
| 65538802 | The UUID provided is not associated with the administrator SVM key manager. |
| 66060296 | Failed to delete key from an external key server. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type KeyManagerAuthKeyDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this key manager auth key delete default response has a 2xx status code
func (o *KeyManagerAuthKeyDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this key manager auth key delete default response has a 3xx status code
func (o *KeyManagerAuthKeyDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this key manager auth key delete default response has a 4xx status code
func (o *KeyManagerAuthKeyDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this key manager auth key delete default response has a 5xx status code
func (o *KeyManagerAuthKeyDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this key manager auth key delete default response a status code equal to that given
func (o *KeyManagerAuthKeyDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the key manager auth key delete default response
func (o *KeyManagerAuthKeyDeleteDefault) Code() int {
	return o._statusCode
}

func (o *KeyManagerAuthKeyDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/key-managers/{security_key_manager.uuid}/auth-keys/{key_id}][%d] key_manager_auth_key_delete default %s", o._statusCode, payload)
}

func (o *KeyManagerAuthKeyDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/key-managers/{security_key_manager.uuid}/auth-keys/{key_id}][%d] key_manager_auth_key_delete default %s", o._statusCode, payload)
}

func (o *KeyManagerAuthKeyDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *KeyManagerAuthKeyDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
