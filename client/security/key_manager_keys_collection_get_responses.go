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

// KeyManagerKeysCollectionGetReader is a Reader for the KeyManagerKeysCollectionGet structure.
type KeyManagerKeysCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyManagerKeysCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyManagerKeysCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKeyManagerKeysCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKeyManagerKeysCollectionGetOK creates a KeyManagerKeysCollectionGetOK with default headers values
func NewKeyManagerKeysCollectionGetOK() *KeyManagerKeysCollectionGetOK {
	return &KeyManagerKeysCollectionGetOK{}
}

/*
KeyManagerKeysCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type KeyManagerKeysCollectionGetOK struct {
	Payload *models.KeyManagerKeysResponse
}

// IsSuccess returns true when this key manager keys collection get o k response has a 2xx status code
func (o *KeyManagerKeysCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this key manager keys collection get o k response has a 3xx status code
func (o *KeyManagerKeysCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this key manager keys collection get o k response has a 4xx status code
func (o *KeyManagerKeysCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this key manager keys collection get o k response has a 5xx status code
func (o *KeyManagerKeysCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this key manager keys collection get o k response a status code equal to that given
func (o *KeyManagerKeysCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the key manager keys collection get o k response
func (o *KeyManagerKeysCollectionGetOK) Code() int {
	return 200
}

func (o *KeyManagerKeysCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers/{security_key_manager.uuid}/keys/{node.uuid}/key-ids][%d] keyManagerKeysCollectionGetOK %s", 200, payload)
}

func (o *KeyManagerKeysCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers/{security_key_manager.uuid}/keys/{node.uuid}/key-ids][%d] keyManagerKeysCollectionGetOK %s", 200, payload)
}

func (o *KeyManagerKeysCollectionGetOK) GetPayload() *models.KeyManagerKeysResponse {
	return o.Payload
}

func (o *KeyManagerKeysCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyManagerKeysResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyManagerKeysCollectionGetDefault creates a KeyManagerKeysCollectionGetDefault with default headers values
func NewKeyManagerKeysCollectionGetDefault(code int) *KeyManagerKeysCollectionGetDefault {
	return &KeyManagerKeysCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	KeyManagerKeysCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65537125 | The provided key manager UUID is not the UUID of a keymanager. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type KeyManagerKeysCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this key manager keys collection get default response has a 2xx status code
func (o *KeyManagerKeysCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this key manager keys collection get default response has a 3xx status code
func (o *KeyManagerKeysCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this key manager keys collection get default response has a 4xx status code
func (o *KeyManagerKeysCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this key manager keys collection get default response has a 5xx status code
func (o *KeyManagerKeysCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this key manager keys collection get default response a status code equal to that given
func (o *KeyManagerKeysCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the key manager keys collection get default response
func (o *KeyManagerKeysCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *KeyManagerKeysCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers/{security_key_manager.uuid}/keys/{node.uuid}/key-ids][%d] key_manager_keys_collection_get default %s", o._statusCode, payload)
}

func (o *KeyManagerKeysCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers/{security_key_manager.uuid}/keys/{node.uuid}/key-ids][%d] key_manager_keys_collection_get default %s", o._statusCode, payload)
}

func (o *KeyManagerKeysCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *KeyManagerKeysCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
