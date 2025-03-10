// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// DNSModifyReader is a Reader for the DNSModify structure.
type DNSModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DNSModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDNSModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDNSModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDNSModifyOK creates a DNSModifyOK with default headers values
func NewDNSModifyOK() *DNSModifyOK {
	return &DNSModifyOK{}
}

/*
DNSModifyOK describes a response with status code 200, with default header values.

OK
*/
type DNSModifyOK struct {
}

// IsSuccess returns true when this dns modify o k response has a 2xx status code
func (o *DNSModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dns modify o k response has a 3xx status code
func (o *DNSModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns modify o k response has a 4xx status code
func (o *DNSModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dns modify o k response has a 5xx status code
func (o *DNSModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dns modify o k response a status code equal to that given
func (o *DNSModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dns modify o k response
func (o *DNSModifyOK) Code() int {
	return 200
}

func (o *DNSModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /name-services/dns/{uuid}][%d] dnsModifyOK", 200)
}

func (o *DNSModifyOK) String() string {
	return fmt.Sprintf("[PATCH /name-services/dns/{uuid}][%d] dnsModifyOK", 200)
}

func (o *DNSModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDNSModifyDefault creates a DNSModifyDefault with default headers values
func NewDNSModifyDefault(code int) *DNSModifyDefault {
	return &DNSModifyDefault{
		_statusCode: code,
	}
}

/*
	DNSModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8847360    | Only admin or data SVMs allowed |
| 8847361    | Exceeded the maximum number of domains allowed. Maximum of six domains only |
| 8847362    | Exceeded the maximum number of name servers allowed. Maximum of three name servers only |
| 8847376    | FQDN is mandatory if dynamic DNS update is being enabled. |
| 8847380    | Secure updates can be enabled only after a CIFS server or an Active Directory account has been created for the SVM. |
| 8847381    | A unique FQDN must be specified for each SVM. |
| 8847383    | The specified TTL exceeds the maximum supported value of 720 hours. |
| 8847392    | Domain name cannot be an IP address |
| 8847393    | Top level domain name is invalid |
| 8847394    | FQDN name violated the limitations |
| 8847399    | One or more of the specified DNS servers do not exist or cannot be reached |
| 8847404    | Dynamic DNS is applicable only for data SVMs |
| 8847405    | DNS parameters updated successfully; however the update of Dynamic DNS-related parameters has failed. |
| 9240587    | FQDN name cannot be empty |
| 9240588    | FQDN name is too long. Maximum supported length: 255 characters  |
| 9240590    | FQDN name is reserved. Following names are reserved: "all", "local" and "localhost" |
| 9240607    | One of the FQDN labels is too long. Maximum supported length is 63 characters |
| 9240608    | FQDN does not contain a ".". At least one "." is mandatory for an FQDN. |
| 9240607    | A label of the FQDN is too long (73 characters). Maximum supported length for each label is 63 characters. |
| 23724130   | Cannot use an IPv6 name server address because there are no IPv6 LIFs |
*/
type DNSModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this dns modify default response has a 2xx status code
func (o *DNSModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dns modify default response has a 3xx status code
func (o *DNSModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dns modify default response has a 4xx status code
func (o *DNSModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dns modify default response has a 5xx status code
func (o *DNSModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dns modify default response a status code equal to that given
func (o *DNSModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dns modify default response
func (o *DNSModifyDefault) Code() int {
	return o._statusCode
}

func (o *DNSModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /name-services/dns/{uuid}][%d] dns_modify default %s", o._statusCode, payload)
}

func (o *DNSModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /name-services/dns/{uuid}][%d] dns_modify default %s", o._statusCode, payload)
}

func (o *DNSModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DNSModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
