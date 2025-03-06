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

// KerberosInterfaceModifyReader is a Reader for the KerberosInterfaceModify structure.
type KerberosInterfaceModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KerberosInterfaceModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKerberosInterfaceModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKerberosInterfaceModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKerberosInterfaceModifyOK creates a KerberosInterfaceModifyOK with default headers values
func NewKerberosInterfaceModifyOK() *KerberosInterfaceModifyOK {
	return &KerberosInterfaceModifyOK{}
}

/*
KerberosInterfaceModifyOK describes a response with status code 200, with default header values.

OK
*/
type KerberosInterfaceModifyOK struct {
}

// IsSuccess returns true when this kerberos interface modify o k response has a 2xx status code
func (o *KerberosInterfaceModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kerberos interface modify o k response has a 3xx status code
func (o *KerberosInterfaceModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kerberos interface modify o k response has a 4xx status code
func (o *KerberosInterfaceModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kerberos interface modify o k response has a 5xx status code
func (o *KerberosInterfaceModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kerberos interface modify o k response a status code equal to that given
func (o *KerberosInterfaceModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kerberos interface modify o k response
func (o *KerberosInterfaceModifyOK) Code() int {
	return 200
}

func (o *KerberosInterfaceModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/nfs/kerberos/interfaces/{interface.uuid}][%d] kerberosInterfaceModifyOK", 200)
}

func (o *KerberosInterfaceModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/nfs/kerberos/interfaces/{interface.uuid}][%d] kerberosInterfaceModifyOK", 200)
}

func (o *KerberosInterfaceModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewKerberosInterfaceModifyDefault creates a KerberosInterfaceModifyDefault with default headers values
func NewKerberosInterfaceModifyDefault(code int) *KerberosInterfaceModifyDefault {
	return &KerberosInterfaceModifyDefault{
		_statusCode: code,
	}
}

/*
	KerberosInterfaceModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error codes | Description |
| ----------- | ----------- |
| 4           | The entry doesn't exist|
| 3276801     | Failed to bind service principal name on LIF.|
| 3276809     | Failed to disable NFS Kerberos on LIF.|
| 3276832     | Failed to insert Kerberos attributes to database.|
| 3276842     | Internal error. Failed to import Kerberos keytab file into the management databases. Contact technical support for assistance.|
| 3276861     | Kerberos is already enabled/disabled on this LIF.|
| 3276862     | Kerberos service principal name is required.|
| 3276889     | Failed to enable NFS Kerberos on LIF.|
| 3276937     | Failed to lookup the Vserver for the virtual interface.|
| 3276941     | Kerberos is a required field.|
| 3276942     | Service principal name is invalid. It must of the format:"nfs/<LIF-FQDN>@REALM"|
| 3276943     | Lookup of Kerberos realm failed. The entry doesn't exist.|
| 3276944     | Internal error. Reason: Failed to initialize the Kerberos context|
| 3276945     | Internal error. Reason: Failed to parse the service principal name|
| 3276951     | Warning: Skipping unsupported encryption type for service principal name|
| 3276952     | "organizational_unit" option cannot be used for "Other" vendor.|
| 3276965     | Account sharing across Vservers is not allowed. Use a different service principal name unique within the first 15 characters.|
| 3277019     | Cannot specify -force when enabling Kerberos.|
| 3277020     | Modifying the NFS Kerberos configuration for a LIF that is not configured for NFS is not supported.|
| 3277043     | Keytab import failed due to missing keys. Keys for encryption types are required for  Vserver but found no matching keys for service principal name. Generate the keytab file with all required keys and try again.|
| 3277133     | The _machine_account_ property can only be specified when the KDC vendor is Microsoft.|
*/
type KerberosInterfaceModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this kerberos interface modify default response has a 2xx status code
func (o *KerberosInterfaceModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kerberos interface modify default response has a 3xx status code
func (o *KerberosInterfaceModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kerberos interface modify default response has a 4xx status code
func (o *KerberosInterfaceModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kerberos interface modify default response has a 5xx status code
func (o *KerberosInterfaceModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kerberos interface modify default response a status code equal to that given
func (o *KerberosInterfaceModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kerberos interface modify default response
func (o *KerberosInterfaceModifyDefault) Code() int {
	return o._statusCode
}

func (o *KerberosInterfaceModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/nfs/kerberos/interfaces/{interface.uuid}][%d] kerberos_interface_modify default %s", o._statusCode, payload)
}

func (o *KerberosInterfaceModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/nfs/kerberos/interfaces/{interface.uuid}][%d] kerberos_interface_modify default %s", o._statusCode, payload)
}

func (o *KerberosInterfaceModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *KerberosInterfaceModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
