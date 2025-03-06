// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// HTTPProxyCreateReader is a Reader for the HTTPProxyCreate structure.
type HTTPProxyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HTTPProxyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewHTTPProxyCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHTTPProxyCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHTTPProxyCreateCreated creates a HTTPProxyCreateCreated with default headers values
func NewHTTPProxyCreateCreated() *HTTPProxyCreateCreated {
	return &HTTPProxyCreateCreated{}
}

/*
HTTPProxyCreateCreated describes a response with status code 201, with default header values.

Created
*/
type HTTPProxyCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.NetworkHTTPProxyResponse
}

// IsSuccess returns true when this http proxy create created response has a 2xx status code
func (o *HTTPProxyCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this http proxy create created response has a 3xx status code
func (o *HTTPProxyCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this http proxy create created response has a 4xx status code
func (o *HTTPProxyCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this http proxy create created response has a 5xx status code
func (o *HTTPProxyCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this http proxy create created response a status code equal to that given
func (o *HTTPProxyCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the http proxy create created response
func (o *HTTPProxyCreateCreated) Code() int {
	return 201
}

func (o *HTTPProxyCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/http-proxy][%d] httpProxyCreateCreated %s", 201, payload)
}

func (o *HTTPProxyCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/http-proxy][%d] httpProxyCreateCreated %s", 201, payload)
}

func (o *HTTPProxyCreateCreated) GetPayload() *models.NetworkHTTPProxyResponse {
	return o.Payload
}

func (o *HTTPProxyCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.NetworkHTTPProxyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHTTPProxyCreateDefault creates a HTTPProxyCreateDefault with default headers values
func NewHTTPProxyCreateDefault(code int) *HTTPProxyCreateDefault {
	return &HTTPProxyCreateDefault{
		_statusCode: code,
	}
}

/*
	HTTPProxyCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 26214473   | HTTP proxy configuration is not valid. |
| 26214476   | The "IPspace" parameter should not be specified in the SVM context. |
| 26214477   | The specified IPspace does not exist. |
| 23724130   | Cannot use an IPv6 name server address because there are no IPv6 interfaces. |
| 26214481   | Username and password cannot be empty when HTTP proxy authentication is enabled. |
| 26214482   | Username and password cannot be specified when HTTP proxy authentication is disabled. |
| 26214480   | One of \"svm.name\", \"svm.uuid\", \"ipspace.name\" or \"ipspace.uuid\" must be specified. |
| 2621462    | SVM \"vs0\" does not exist. |
| 262186     | Field \"svm.name\" cannot be used with field \"ipspace.name\". |
*/
type HTTPProxyCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this http proxy create default response has a 2xx status code
func (o *HTTPProxyCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this http proxy create default response has a 3xx status code
func (o *HTTPProxyCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this http proxy create default response has a 4xx status code
func (o *HTTPProxyCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this http proxy create default response has a 5xx status code
func (o *HTTPProxyCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this http proxy create default response a status code equal to that given
func (o *HTTPProxyCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the http proxy create default response
func (o *HTTPProxyCreateDefault) Code() int {
	return o._statusCode
}

func (o *HTTPProxyCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/http-proxy][%d] http_proxy_create default %s", o._statusCode, payload)
}

func (o *HTTPProxyCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/http-proxy][%d] http_proxy_create default %s", o._statusCode, payload)
}

func (o *HTTPProxyCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *HTTPProxyCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
