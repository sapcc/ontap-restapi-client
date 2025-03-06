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

// HTTPProxyModifyReader is a Reader for the HTTPProxyModify structure.
type HTTPProxyModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HTTPProxyModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHTTPProxyModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHTTPProxyModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHTTPProxyModifyOK creates a HTTPProxyModifyOK with default headers values
func NewHTTPProxyModifyOK() *HTTPProxyModifyOK {
	return &HTTPProxyModifyOK{}
}

/*
HTTPProxyModifyOK describes a response with status code 200, with default header values.

OK
*/
type HTTPProxyModifyOK struct {
}

// IsSuccess returns true when this http proxy modify o k response has a 2xx status code
func (o *HTTPProxyModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this http proxy modify o k response has a 3xx status code
func (o *HTTPProxyModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this http proxy modify o k response has a 4xx status code
func (o *HTTPProxyModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this http proxy modify o k response has a 5xx status code
func (o *HTTPProxyModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this http proxy modify o k response a status code equal to that given
func (o *HTTPProxyModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the http proxy modify o k response
func (o *HTTPProxyModifyOK) Code() int {
	return 200
}

func (o *HTTPProxyModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /network/http-proxy/{uuid}][%d] httpProxyModifyOK", 200)
}

func (o *HTTPProxyModifyOK) String() string {
	return fmt.Sprintf("[PATCH /network/http-proxy/{uuid}][%d] httpProxyModifyOK", 200)
}

func (o *HTTPProxyModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHTTPProxyModifyDefault creates a HTTPProxyModifyDefault with default headers values
func NewHTTPProxyModifyDefault(code int) *HTTPProxyModifyDefault {
	return &HTTPProxyModifyDefault{
		_statusCode: code,
	}
}

/*
	HTTPProxyModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 26214473   | The HTTP proxy configuration is not valid. |
| 23724130   | Cannot use an IPv6 name server address because there are no IPv6 interfaces. |
| 26214481   | Username and password cannot be empty when HTTP proxy authentication is enabled. |
| 26214482   | Username and password cannot be specified when HTTP proxy authentication is disabled. |
*/
type HTTPProxyModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this http proxy modify default response has a 2xx status code
func (o *HTTPProxyModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this http proxy modify default response has a 3xx status code
func (o *HTTPProxyModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this http proxy modify default response has a 4xx status code
func (o *HTTPProxyModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this http proxy modify default response has a 5xx status code
func (o *HTTPProxyModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this http proxy modify default response a status code equal to that given
func (o *HTTPProxyModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the http proxy modify default response
func (o *HTTPProxyModifyDefault) Code() int {
	return o._statusCode
}

func (o *HTTPProxyModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/http-proxy/{uuid}][%d] http_proxy_modify default %s", o._statusCode, payload)
}

func (o *HTTPProxyModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/http-proxy/{uuid}][%d] http_proxy_modify default %s", o._statusCode, payload)
}

func (o *HTTPProxyModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *HTTPProxyModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
