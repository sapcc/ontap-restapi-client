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

// IpsecCaCertificateGetReader is a Reader for the IpsecCaCertificateGet structure.
type IpsecCaCertificateGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpsecCaCertificateGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpsecCaCertificateGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpsecCaCertificateGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpsecCaCertificateGetOK creates a IpsecCaCertificateGetOK with default headers values
func NewIpsecCaCertificateGetOK() *IpsecCaCertificateGetOK {
	return &IpsecCaCertificateGetOK{}
}

/*
IpsecCaCertificateGetOK describes a response with status code 200, with default header values.

OK
*/
type IpsecCaCertificateGetOK struct {
	Payload *models.IpsecCaCertificate
}

// IsSuccess returns true when this ipsec ca certificate get o k response has a 2xx status code
func (o *IpsecCaCertificateGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipsec ca certificate get o k response has a 3xx status code
func (o *IpsecCaCertificateGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipsec ca certificate get o k response has a 4xx status code
func (o *IpsecCaCertificateGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipsec ca certificate get o k response has a 5xx status code
func (o *IpsecCaCertificateGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipsec ca certificate get o k response a status code equal to that given
func (o *IpsecCaCertificateGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipsec ca certificate get o k response
func (o *IpsecCaCertificateGetOK) Code() int {
	return 200
}

func (o *IpsecCaCertificateGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/ipsec/ca-certificates/{certificate.uuid}][%d] ipsecCaCertificateGetOK %s", 200, payload)
}

func (o *IpsecCaCertificateGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/ipsec/ca-certificates/{certificate.uuid}][%d] ipsecCaCertificateGetOK %s", 200, payload)
}

func (o *IpsecCaCertificateGetOK) GetPayload() *models.IpsecCaCertificate {
	return o.Payload
}

func (o *IpsecCaCertificateGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpsecCaCertificate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpsecCaCertificateGetDefault creates a IpsecCaCertificateGetDefault with default headers values
func NewIpsecCaCertificateGetDefault(code int) *IpsecCaCertificateGetDefault {
	return &IpsecCaCertificateGetDefault{
		_statusCode: code,
	}
}

/*
IpsecCaCertificateGetDefault describes a response with status code -1, with default header values.

Error
*/
type IpsecCaCertificateGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ipsec ca certificate get default response has a 2xx status code
func (o *IpsecCaCertificateGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipsec ca certificate get default response has a 3xx status code
func (o *IpsecCaCertificateGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipsec ca certificate get default response has a 4xx status code
func (o *IpsecCaCertificateGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipsec ca certificate get default response has a 5xx status code
func (o *IpsecCaCertificateGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipsec ca certificate get default response a status code equal to that given
func (o *IpsecCaCertificateGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipsec ca certificate get default response
func (o *IpsecCaCertificateGetDefault) Code() int {
	return o._statusCode
}

func (o *IpsecCaCertificateGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/ipsec/ca-certificates/{certificate.uuid}][%d] ipsec_ca_certificate_get default %s", o._statusCode, payload)
}

func (o *IpsecCaCertificateGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/ipsec/ca-certificates/{certificate.uuid}][%d] ipsec_ca_certificate_get default %s", o._statusCode, payload)
}

func (o *IpsecCaCertificateGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IpsecCaCertificateGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
