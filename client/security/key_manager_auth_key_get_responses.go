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

// KeyManagerAuthKeyGetReader is a Reader for the KeyManagerAuthKeyGet structure.
type KeyManagerAuthKeyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyManagerAuthKeyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyManagerAuthKeyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKeyManagerAuthKeyGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKeyManagerAuthKeyGetOK creates a KeyManagerAuthKeyGetOK with default headers values
func NewKeyManagerAuthKeyGetOK() *KeyManagerAuthKeyGetOK {
	return &KeyManagerAuthKeyGetOK{}
}

/*
KeyManagerAuthKeyGetOK describes a response with status code 200, with default header values.

OK
*/
type KeyManagerAuthKeyGetOK struct {
	Payload *models.KeyManagerAuthKey
}

// IsSuccess returns true when this key manager auth key get o k response has a 2xx status code
func (o *KeyManagerAuthKeyGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this key manager auth key get o k response has a 3xx status code
func (o *KeyManagerAuthKeyGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this key manager auth key get o k response has a 4xx status code
func (o *KeyManagerAuthKeyGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this key manager auth key get o k response has a 5xx status code
func (o *KeyManagerAuthKeyGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this key manager auth key get o k response a status code equal to that given
func (o *KeyManagerAuthKeyGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the key manager auth key get o k response
func (o *KeyManagerAuthKeyGetOK) Code() int {
	return 200
}

func (o *KeyManagerAuthKeyGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers/{security_key_manager.uuid}/auth-keys/{key_id}][%d] keyManagerAuthKeyGetOK %s", 200, payload)
}

func (o *KeyManagerAuthKeyGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers/{security_key_manager.uuid}/auth-keys/{key_id}][%d] keyManagerAuthKeyGetOK %s", 200, payload)
}

func (o *KeyManagerAuthKeyGetOK) GetPayload() *models.KeyManagerAuthKey {
	return o.Payload
}

func (o *KeyManagerAuthKeyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyManagerAuthKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyManagerAuthKeyGetDefault creates a KeyManagerAuthKeyGetDefault with default headers values
func NewKeyManagerAuthKeyGetDefault(code int) *KeyManagerAuthKeyGetDefault {
	return &KeyManagerAuthKeyGetDefault{
		_statusCode: code,
	}
}

/*
	KeyManagerAuthKeyGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65536828 | External key management is not enabled for the SVM. |
| 65536856 | No key servers found for the SVM. |
| 65536896 | External key management is not configured on the partner site. |
| 65538800 | External KMIP DKMIP keymanager not configured on administrative Vserver. |
| 65538801 | Internal error while accessing keymanager database. |
| 65538802 | The UUID provided is not associated with the administrator SVM key manager. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type KeyManagerAuthKeyGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this key manager auth key get default response has a 2xx status code
func (o *KeyManagerAuthKeyGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this key manager auth key get default response has a 3xx status code
func (o *KeyManagerAuthKeyGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this key manager auth key get default response has a 4xx status code
func (o *KeyManagerAuthKeyGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this key manager auth key get default response has a 5xx status code
func (o *KeyManagerAuthKeyGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this key manager auth key get default response a status code equal to that given
func (o *KeyManagerAuthKeyGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the key manager auth key get default response
func (o *KeyManagerAuthKeyGetDefault) Code() int {
	return o._statusCode
}

func (o *KeyManagerAuthKeyGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers/{security_key_manager.uuid}/auth-keys/{key_id}][%d] key_manager_auth_key_get default %s", o._statusCode, payload)
}

func (o *KeyManagerAuthKeyGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers/{security_key_manager.uuid}/auth-keys/{key_id}][%d] key_manager_auth_key_get default %s", o._statusCode, payload)
}

func (o *KeyManagerAuthKeyGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *KeyManagerAuthKeyGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
