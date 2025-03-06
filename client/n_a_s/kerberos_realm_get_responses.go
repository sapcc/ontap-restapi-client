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

	"github.com/sapcc/ontap-restapi-client/models"
)

// KerberosRealmGetReader is a Reader for the KerberosRealmGet structure.
type KerberosRealmGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KerberosRealmGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKerberosRealmGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKerberosRealmGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKerberosRealmGetOK creates a KerberosRealmGetOK with default headers values
func NewKerberosRealmGetOK() *KerberosRealmGetOK {
	return &KerberosRealmGetOK{}
}

/*
KerberosRealmGetOK describes a response with status code 200, with default header values.

OK
*/
type KerberosRealmGetOK struct {
	Payload *models.KerberosRealm
}

// IsSuccess returns true when this kerberos realm get o k response has a 2xx status code
func (o *KerberosRealmGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kerberos realm get o k response has a 3xx status code
func (o *KerberosRealmGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kerberos realm get o k response has a 4xx status code
func (o *KerberosRealmGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kerberos realm get o k response has a 5xx status code
func (o *KerberosRealmGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kerberos realm get o k response a status code equal to that given
func (o *KerberosRealmGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kerberos realm get o k response
func (o *KerberosRealmGetOK) Code() int {
	return 200
}

func (o *KerberosRealmGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/kerberos/realms/{svm.uuid}/{name}][%d] kerberosRealmGetOK %s", 200, payload)
}

func (o *KerberosRealmGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/kerberos/realms/{svm.uuid}/{name}][%d] kerberosRealmGetOK %s", 200, payload)
}

func (o *KerberosRealmGetOK) GetPayload() *models.KerberosRealm {
	return o.Payload
}

func (o *KerberosRealmGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KerberosRealm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKerberosRealmGetDefault creates a KerberosRealmGetDefault with default headers values
func NewKerberosRealmGetDefault(code int) *KerberosRealmGetDefault {
	return &KerberosRealmGetDefault{
		_statusCode: code,
	}
}

/*
KerberosRealmGetDefault describes a response with status code -1, with default header values.

Error
*/
type KerberosRealmGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this kerberos realm get default response has a 2xx status code
func (o *KerberosRealmGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kerberos realm get default response has a 3xx status code
func (o *KerberosRealmGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kerberos realm get default response has a 4xx status code
func (o *KerberosRealmGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kerberos realm get default response has a 5xx status code
func (o *KerberosRealmGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kerberos realm get default response a status code equal to that given
func (o *KerberosRealmGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kerberos realm get default response
func (o *KerberosRealmGetDefault) Code() int {
	return o._statusCode
}

func (o *KerberosRealmGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/kerberos/realms/{svm.uuid}/{name}][%d] kerberos_realm_get default %s", o._statusCode, payload)
}

func (o *KerberosRealmGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/kerberos/realms/{svm.uuid}/{name}][%d] kerberos_realm_get default %s", o._statusCode, payload)
}

func (o *KerberosRealmGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *KerberosRealmGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
