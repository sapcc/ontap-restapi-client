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

// NfsTLSInterfaceModifyReader is a Reader for the NfsTLSInterfaceModify structure.
type NfsTLSInterfaceModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NfsTLSInterfaceModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNfsTLSInterfaceModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNfsTLSInterfaceModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNfsTLSInterfaceModifyOK creates a NfsTLSInterfaceModifyOK with default headers values
func NewNfsTLSInterfaceModifyOK() *NfsTLSInterfaceModifyOK {
	return &NfsTLSInterfaceModifyOK{}
}

/*
NfsTLSInterfaceModifyOK describes a response with status code 200, with default header values.

OK
*/
type NfsTLSInterfaceModifyOK struct {
}

// IsSuccess returns true when this nfs Tls interface modify o k response has a 2xx status code
func (o *NfsTLSInterfaceModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nfs Tls interface modify o k response has a 3xx status code
func (o *NfsTLSInterfaceModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nfs Tls interface modify o k response has a 4xx status code
func (o *NfsTLSInterfaceModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nfs Tls interface modify o k response has a 5xx status code
func (o *NfsTLSInterfaceModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nfs Tls interface modify o k response a status code equal to that given
func (o *NfsTLSInterfaceModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nfs Tls interface modify o k response
func (o *NfsTLSInterfaceModifyOK) Code() int {
	return 200
}

func (o *NfsTLSInterfaceModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/nfs/tls/interfaces/{interface.uuid}][%d] nfsTlsInterfaceModifyOK", 200)
}

func (o *NfsTLSInterfaceModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/nfs/tls/interfaces/{interface.uuid}][%d] nfsTlsInterfaceModifyOK", 200)
}

func (o *NfsTLSInterfaceModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNfsTLSInterfaceModifyDefault creates a NfsTLSInterfaceModifyDefault with default headers values
func NewNfsTLSInterfaceModifyDefault(code int) *NfsTLSInterfaceModifyDefault {
	return &NfsTLSInterfaceModifyDefault{
		_statusCode: code,
	}
}

/*
	NfsTLSInterfaceModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error codes | Description |
| ----------- | ----------- |
| 2           | The value provided is an invalid value for field \"interface.uuid\".|
| 4           | The entry doesn't exist|
| 3277204     | TLS already enabled/disabled on this LIF.|
| 3277205     | Cannot enable TLS because no certificate was provided.|
| 3277206     | The "enabled" is a required field.|
| 3277210     | The FQDN of the LIF IP address and the common name present in the certificate do not match. Update the certificate so that the common name and LIF FQDN match and retry the operation.|
| 3277217     | Failed to enable TLS on LIF on vserver because the certificate does not have LIF as subject alternate name.|
| 92405900    | Certificate not found for the SVM.|
| 92406020    | Only certificates of type 'server' are supported on the SVM.|
*/
type NfsTLSInterfaceModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nfs tls interface modify default response has a 2xx status code
func (o *NfsTLSInterfaceModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nfs tls interface modify default response has a 3xx status code
func (o *NfsTLSInterfaceModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nfs tls interface modify default response has a 4xx status code
func (o *NfsTLSInterfaceModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nfs tls interface modify default response has a 5xx status code
func (o *NfsTLSInterfaceModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nfs tls interface modify default response a status code equal to that given
func (o *NfsTLSInterfaceModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nfs tls interface modify default response
func (o *NfsTLSInterfaceModifyDefault) Code() int {
	return o._statusCode
}

func (o *NfsTLSInterfaceModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/nfs/tls/interfaces/{interface.uuid}][%d] nfs_tls_interface_modify default %s", o._statusCode, payload)
}

func (o *NfsTLSInterfaceModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/nfs/tls/interfaces/{interface.uuid}][%d] nfs_tls_interface_modify default %s", o._statusCode, payload)
}

func (o *NfsTLSInterfaceModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NfsTLSInterfaceModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
