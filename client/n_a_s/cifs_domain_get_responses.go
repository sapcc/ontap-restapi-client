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

// CifsDomainGetReader is a Reader for the CifsDomainGet structure.
type CifsDomainGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsDomainGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsDomainGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsDomainGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsDomainGetOK creates a CifsDomainGetOK with default headers values
func NewCifsDomainGetOK() *CifsDomainGetOK {
	return &CifsDomainGetOK{}
}

/*
CifsDomainGetOK describes a response with status code 200, with default header values.

OK
*/
type CifsDomainGetOK struct {
	Payload *models.CifsDomain
}

// IsSuccess returns true when this cifs domain get o k response has a 2xx status code
func (o *CifsDomainGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs domain get o k response has a 3xx status code
func (o *CifsDomainGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs domain get o k response has a 4xx status code
func (o *CifsDomainGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs domain get o k response has a 5xx status code
func (o *CifsDomainGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs domain get o k response a status code equal to that given
func (o *CifsDomainGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs domain get o k response
func (o *CifsDomainGetOK) Code() int {
	return 200
}

func (o *CifsDomainGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/domains/{svm.uuid}][%d] cifsDomainGetOK %s", 200, payload)
}

func (o *CifsDomainGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/domains/{svm.uuid}][%d] cifsDomainGetOK %s", 200, payload)
}

func (o *CifsDomainGetOK) GetPayload() *models.CifsDomain {
	return o.Payload
}

func (o *CifsDomainGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CifsDomain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsDomainGetDefault creates a CifsDomainGetDefault with default headers values
func NewCifsDomainGetDefault(code int) *CifsDomainGetDefault {
	return &CifsDomainGetDefault{
		_statusCode: code,
	}
}

/*
	CifsDomainGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
|  656463    | SVM UUID must be provided for a query on the field rediscover_trusts and reset_discovered_servers. |
*/
type CifsDomainGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs domain get default response has a 2xx status code
func (o *CifsDomainGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs domain get default response has a 3xx status code
func (o *CifsDomainGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs domain get default response has a 4xx status code
func (o *CifsDomainGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs domain get default response has a 5xx status code
func (o *CifsDomainGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs domain get default response a status code equal to that given
func (o *CifsDomainGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs domain get default response
func (o *CifsDomainGetDefault) Code() int {
	return o._statusCode
}

func (o *CifsDomainGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/domains/{svm.uuid}][%d] cifs_domain_get default %s", o._statusCode, payload)
}

func (o *CifsDomainGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/domains/{svm.uuid}][%d] cifs_domain_get default %s", o._statusCode, payload)
}

func (o *CifsDomainGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsDomainGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
