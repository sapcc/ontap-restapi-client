// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// WebModifyReader is a Reader for the WebModify structure.
type WebModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WebModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWebModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewWebModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWebModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWebModifyOK creates a WebModifyOK with default headers values
func NewWebModifyOK() *WebModifyOK {
	return &WebModifyOK{}
}

/*
WebModifyOK describes a response with status code 200, with default header values.

OK
*/
type WebModifyOK struct {
}

// IsSuccess returns true when this web modify o k response has a 2xx status code
func (o *WebModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this web modify o k response has a 3xx status code
func (o *WebModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this web modify o k response has a 4xx status code
func (o *WebModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this web modify o k response has a 5xx status code
func (o *WebModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this web modify o k response a status code equal to that given
func (o *WebModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the web modify o k response
func (o *WebModifyOK) Code() int {
	return 200
}

func (o *WebModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /cluster/web][%d] webModifyOK", 200)
}

func (o *WebModifyOK) String() string {
	return fmt.Sprintf("[PATCH /cluster/web][%d] webModifyOK", 200)
}

func (o *WebModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWebModifyAccepted creates a WebModifyAccepted with default headers values
func NewWebModifyAccepted() *WebModifyAccepted {
	return &WebModifyAccepted{}
}

/*
WebModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type WebModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this web modify accepted response has a 2xx status code
func (o *WebModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this web modify accepted response has a 3xx status code
func (o *WebModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this web modify accepted response has a 4xx status code
func (o *WebModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this web modify accepted response has a 5xx status code
func (o *WebModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this web modify accepted response a status code equal to that given
func (o *WebModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the web modify accepted response
func (o *WebModifyAccepted) Code() int {
	return 202
}

func (o *WebModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/web][%d] webModifyAccepted %s", 202, payload)
}

func (o *WebModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/web][%d] webModifyAccepted %s", 202, payload)
}

func (o *WebModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *WebModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebModifyDefault creates a WebModifyDefault with default headers values
func NewWebModifyDefault(code int) *WebModifyDefault {
	return &WebModifyDefault{
		_statusCode: code,
	}
}

/*
	WebModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 9830406 | Reconfiguration of the web services failed. |
| 9830407 | The web services failed to restart. |
| 9830408 | Reconfiguration and/or restart of the web services failed. |
| 9830442 | Client authentication cannot be enabled without a client ca certificate. |
| 9830463 | The cluster must be fully upgraded before modifying this resource. |
| 9830464 | HTTP cannot be enabled when FIPS is also enabled. |
| 9830483 | The CSRF token timeout is invalid. |
| 9830484 | The maximum concurrent CSRF token count cannot be lower than 100. |
| 9830485 | The CSRF idle timeout cannot be greater than the CSRF absolute timeout. |
| 9830486 | CSRF requires an effective cluster version of 9.7 or later. |
| 9830487 | The HTTP and HTTPS ports must not have the same value. |
| 9830488 | The certificate is not a "server" certificate. |
| 9830489 | The certificate does not exist for the given SVM. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type WebModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this web modify default response has a 2xx status code
func (o *WebModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this web modify default response has a 3xx status code
func (o *WebModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this web modify default response has a 4xx status code
func (o *WebModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this web modify default response has a 5xx status code
func (o *WebModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this web modify default response a status code equal to that given
func (o *WebModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the web modify default response
func (o *WebModifyDefault) Code() int {
	return o._statusCode
}

func (o *WebModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/web][%d] web_modify default %s", o._statusCode, payload)
}

func (o *WebModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/web][%d] web_modify default %s", o._statusCode, payload)
}

func (o *WebModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *WebModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
