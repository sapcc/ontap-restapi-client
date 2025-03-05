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

// LocalHostCreateReader is a Reader for the LocalHostCreate structure.
type LocalHostCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalHostCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewLocalHostCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocalHostCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocalHostCreateCreated creates a LocalHostCreateCreated with default headers values
func NewLocalHostCreateCreated() *LocalHostCreateCreated {
	return &LocalHostCreateCreated{}
}

/*
LocalHostCreateCreated describes a response with status code 201, with default header values.

Created
*/
type LocalHostCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string
}

// IsSuccess returns true when this local host create created response has a 2xx status code
func (o *LocalHostCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this local host create created response has a 3xx status code
func (o *LocalHostCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this local host create created response has a 4xx status code
func (o *LocalHostCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this local host create created response has a 5xx status code
func (o *LocalHostCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this local host create created response a status code equal to that given
func (o *LocalHostCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the local host create created response
func (o *LocalHostCreateCreated) Code() int {
	return 201
}

func (o *LocalHostCreateCreated) Error() string {
	return fmt.Sprintf("[POST /name-services/local-hosts][%d] localHostCreateCreated", 201)
}

func (o *LocalHostCreateCreated) String() string {
	return fmt.Sprintf("[POST /name-services/local-hosts][%d] localHostCreateCreated", 201)
}

func (o *LocalHostCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewLocalHostCreateDefault creates a LocalHostCreateDefault with default headers values
func NewLocalHostCreateDefault(code int) *LocalHostCreateDefault {
	return &LocalHostCreateDefault{
		_statusCode: code,
	}
}

/*
	LocalHostCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1377682 | IPv6 is not enabled in the cluster. |
| 1966253 | IPv6 is not enabled in the cluster. Enable IPv6 using command "network options ipv6 modify -enabled true" and try again. |
| 2621706 | The specified owner UUID is incorrect for the specified owner name. |
| 8912896 | Only admin or data Vservers allowed. |
| 23724055 | Internal error. Configuration for Vserver failed. Verify that the cluster is healthy, then try the command again. For further assistance, contact technical support. |
| 23724155 | The specified IPv4 address is not supported because it is one of the following: multicast, loopback, 0.0.0.0, or broadcast. |
| 23724156 | The specified IPv6 address is not supported because it is one of the following: ::, link-local, multicast, v4-compatible, v4-mapped, or loopback. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type LocalHostCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this local host create default response has a 2xx status code
func (o *LocalHostCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this local host create default response has a 3xx status code
func (o *LocalHostCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this local host create default response has a 4xx status code
func (o *LocalHostCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this local host create default response has a 5xx status code
func (o *LocalHostCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this local host create default response a status code equal to that given
func (o *LocalHostCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the local host create default response
func (o *LocalHostCreateDefault) Code() int {
	return o._statusCode
}

func (o *LocalHostCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /name-services/local-hosts][%d] local_host_create default %s", o._statusCode, payload)
}

func (o *LocalHostCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /name-services/local-hosts][%d] local_host_create default %s", o._statusCode, payload)
}

func (o *LocalHostCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LocalHostCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
