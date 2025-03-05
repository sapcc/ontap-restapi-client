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

	"github.com/sapcc/ontap-restapi/models"
)

// DNSCreateReader is a Reader for the DNSCreate structure.
type DNSCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DNSCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDNSCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDNSCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDNSCreateCreated creates a DNSCreateCreated with default headers values
func NewDNSCreateCreated() *DNSCreateCreated {
	return &DNSCreateCreated{}
}

/*
DNSCreateCreated describes a response with status code 201, with default header values.

Created
*/
type DNSCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.DNSResponse
}

// IsSuccess returns true when this dns create created response has a 2xx status code
func (o *DNSCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dns create created response has a 3xx status code
func (o *DNSCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns create created response has a 4xx status code
func (o *DNSCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dns create created response has a 5xx status code
func (o *DNSCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dns create created response a status code equal to that given
func (o *DNSCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dns create created response
func (o *DNSCreateCreated) Code() int {
	return 201
}

func (o *DNSCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /name-services/dns][%d] dnsCreateCreated %s", 201, payload)
}

func (o *DNSCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /name-services/dns][%d] dnsCreateCreated %s", 201, payload)
}

func (o *DNSCreateCreated) GetPayload() *models.DNSResponse {
	return o.Payload
}

func (o *DNSCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.DNSResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSCreateDefault creates a DNSCreateDefault with default headers values
func NewDNSCreateDefault(code int) *DNSCreateDefault {
	return &DNSCreateDefault{
		_statusCode: code,
	}
}

/*
	DNSCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621706    | The specified SVM UUID is incorrect for the specified SVM name |
| 8847360    | Only admin or data SVMs allowed |
| 8847361    | Exceeded the maximum number of domains allowed. Maximum of six domains only |
| 8847362    | Exceeded the maximum number of name servers allowed. Maximum of three name servers only |
| 8847392    | Domain name cannot be an IP address |
| 8847393    | Top level domain name is invalid |
| 8847399    | One or more of the specified DNS servers do not exist or cannot be reached |
| 8847394    | FQDN name violated the limitations |
| 8847403    | Scope specified is invalid for the specified SVM |
| 9240587    | FQDN name cannot be empty |
| 9240588    | FQDN name is too long. Maximum supported length: 255 characters  |
| 9240590    | FQDN name is reserved. Following names are reserved: "all", "local" and "localhost" |
| 9240607    | One of the FQDN labels is too long. Maximum supported length is 63 characters |
| 13434916   | The SVM is in the process of being created. Wait a few minutes, and then try the command again. |
| 23724130   | Cannot use an IPv6 name server address because there are no IPv6 LIFs |
| 1377682    | IPv6 is not enabled in the cluster |
*/
type DNSCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this dns create default response has a 2xx status code
func (o *DNSCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dns create default response has a 3xx status code
func (o *DNSCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dns create default response has a 4xx status code
func (o *DNSCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dns create default response has a 5xx status code
func (o *DNSCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dns create default response a status code equal to that given
func (o *DNSCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dns create default response
func (o *DNSCreateDefault) Code() int {
	return o._statusCode
}

func (o *DNSCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /name-services/dns][%d] dns_create default %s", o._statusCode, payload)
}

func (o *DNSCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /name-services/dns][%d] dns_create default %s", o._statusCode, payload)
}

func (o *DNSCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DNSCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
