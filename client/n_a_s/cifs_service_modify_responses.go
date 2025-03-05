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

	"github.com/sapcc/ontap-restapi/models"
)

// CifsServiceModifyReader is a Reader for the CifsServiceModify structure.
type CifsServiceModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsServiceModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsServiceModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCifsServiceModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsServiceModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsServiceModifyOK creates a CifsServiceModifyOK with default headers values
func NewCifsServiceModifyOK() *CifsServiceModifyOK {
	return &CifsServiceModifyOK{}
}

/*
CifsServiceModifyOK describes a response with status code 200, with default header values.

OK
*/
type CifsServiceModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cifs service modify o k response has a 2xx status code
func (o *CifsServiceModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs service modify o k response has a 3xx status code
func (o *CifsServiceModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs service modify o k response has a 4xx status code
func (o *CifsServiceModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs service modify o k response has a 5xx status code
func (o *CifsServiceModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs service modify o k response a status code equal to that given
func (o *CifsServiceModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs service modify o k response
func (o *CifsServiceModifyOK) Code() int {
	return 200
}

func (o *CifsServiceModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/services/{svm.uuid}][%d] cifsServiceModifyOK %s", 200, payload)
}

func (o *CifsServiceModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/services/{svm.uuid}][%d] cifsServiceModifyOK %s", 200, payload)
}

func (o *CifsServiceModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *CifsServiceModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsServiceModifyAccepted creates a CifsServiceModifyAccepted with default headers values
func NewCifsServiceModifyAccepted() *CifsServiceModifyAccepted {
	return &CifsServiceModifyAccepted{}
}

/*
CifsServiceModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type CifsServiceModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cifs service modify accepted response has a 2xx status code
func (o *CifsServiceModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs service modify accepted response has a 3xx status code
func (o *CifsServiceModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs service modify accepted response has a 4xx status code
func (o *CifsServiceModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs service modify accepted response has a 5xx status code
func (o *CifsServiceModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs service modify accepted response a status code equal to that given
func (o *CifsServiceModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the cifs service modify accepted response
func (o *CifsServiceModifyAccepted) Code() int {
	return 202
}

func (o *CifsServiceModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/services/{svm.uuid}][%d] cifsServiceModifyAccepted %s", 202, payload)
}

func (o *CifsServiceModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/services/{svm.uuid}][%d] cifsServiceModifyAccepted %s", 202, payload)
}

func (o *CifsServiceModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *CifsServiceModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsServiceModifyDefault creates a CifsServiceModifyDefault with default headers values
func NewCifsServiceModifyDefault(code int) *CifsServiceModifyDefault {
	return &CifsServiceModifyDefault{
		_statusCode: code,
	}
}

/*
	CifsServiceModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 3735751    | Failed to authenticate and retrieve the access token from the Azure OAuth host. |
| 3735752    | Failed to extract the private key from the Azure Key Vault certificate. |
| 3735753    | Unsupported content_type in the Azure secrets response. |
| 3735754    | Failed to parse the JSON response from Azure Key Vault. |
| 3735755    | REST API call to Azure failed. |
| 3735756    | Invalid client certificate. |
| 3735757    | Failed to generate client assertion. |
| 3735762    | The provided Azure Key Vault configuration is incorrect. |
| 3735763    | The provided Azure Key Vault configuration is incomplete. |
| 3735764    | Request to Azure failed. Reason - Azure error code and Azure error message. |
| 655390     | STARTTLS and LDAPS cannot be used together.|
| 655562     | NetBIOS name is longer than 15 characters. |
| 655538     | Unable to modify the CIFS server. The server name is already used by another SVM.|
| 655563     | NetBIOS name contains characters that are not allowed. |
| 655771     | The number of NetBIOS aliases cannot exceed the maximum supported number of '200'. |
| 655923     | Retrieving credentials from AKV is not supported because the effective cluster version is not ONTAP 9.15.0 or later. |
| 656473     | Fields security.kdc_encryption and security.advertised_kdc_encryptions are mutually exclusive. Specify only one of the two.|
*/
type CifsServiceModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs service modify default response has a 2xx status code
func (o *CifsServiceModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs service modify default response has a 3xx status code
func (o *CifsServiceModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs service modify default response has a 4xx status code
func (o *CifsServiceModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs service modify default response has a 5xx status code
func (o *CifsServiceModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs service modify default response a status code equal to that given
func (o *CifsServiceModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs service modify default response
func (o *CifsServiceModifyDefault) Code() int {
	return o._statusCode
}

func (o *CifsServiceModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/services/{svm.uuid}][%d] cifs_service_modify default %s", o._statusCode, payload)
}

func (o *CifsServiceModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/services/{svm.uuid}][%d] cifs_service_modify default %s", o._statusCode, payload)
}

func (o *CifsServiceModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsServiceModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
