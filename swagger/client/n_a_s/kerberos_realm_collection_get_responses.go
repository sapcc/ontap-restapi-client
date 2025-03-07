// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// KerberosRealmCollectionGetReader is a Reader for the KerberosRealmCollectionGet structure.
type KerberosRealmCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KerberosRealmCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKerberosRealmCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKerberosRealmCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKerberosRealmCollectionGetOK creates a KerberosRealmCollectionGetOK with default headers values
func NewKerberosRealmCollectionGetOK() *KerberosRealmCollectionGetOK {
	return &KerberosRealmCollectionGetOK{}
}

/*
KerberosRealmCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type KerberosRealmCollectionGetOK struct {
	Payload *models.KerberosRealmResponse
}

// IsSuccess returns true when this kerberos realm collection get o k response has a 2xx status code
func (o *KerberosRealmCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kerberos realm collection get o k response has a 3xx status code
func (o *KerberosRealmCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kerberos realm collection get o k response has a 4xx status code
func (o *KerberosRealmCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kerberos realm collection get o k response has a 5xx status code
func (o *KerberosRealmCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kerberos realm collection get o k response a status code equal to that given
func (o *KerberosRealmCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kerberos realm collection get o k response
func (o *KerberosRealmCollectionGetOK) Code() int {
	return 200
}

func (o *KerberosRealmCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/kerberos/realms][%d] kerberosRealmCollectionGetOK %s", 200, payload)
}

func (o *KerberosRealmCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/kerberos/realms][%d] kerberosRealmCollectionGetOK %s", 200, payload)
}

func (o *KerberosRealmCollectionGetOK) GetPayload() *models.KerberosRealmResponse {
	return o.Payload
}

func (o *KerberosRealmCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KerberosRealmResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKerberosRealmCollectionGetDefault creates a KerberosRealmCollectionGetDefault with default headers values
func NewKerberosRealmCollectionGetDefault(code int) *KerberosRealmCollectionGetDefault {
	return &KerberosRealmCollectionGetDefault{
		_statusCode: code,
	}
}

/*
KerberosRealmCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type KerberosRealmCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this kerberos realm collection get default response has a 2xx status code
func (o *KerberosRealmCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kerberos realm collection get default response has a 3xx status code
func (o *KerberosRealmCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kerberos realm collection get default response has a 4xx status code
func (o *KerberosRealmCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kerberos realm collection get default response has a 5xx status code
func (o *KerberosRealmCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kerberos realm collection get default response a status code equal to that given
func (o *KerberosRealmCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kerberos realm collection get default response
func (o *KerberosRealmCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *KerberosRealmCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/kerberos/realms][%d] kerberos_realm_collection_get default %s", o._statusCode, payload)
}

func (o *KerberosRealmCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/kerberos/realms][%d] kerberos_realm_collection_get default %s", o._statusCode, payload)
}

func (o *KerberosRealmCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *KerberosRealmCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
