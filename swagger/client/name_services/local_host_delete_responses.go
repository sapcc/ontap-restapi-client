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

// LocalHostDeleteReader is a Reader for the LocalHostDelete structure.
type LocalHostDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalHostDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocalHostDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocalHostDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocalHostDeleteOK creates a LocalHostDeleteOK with default headers values
func NewLocalHostDeleteOK() *LocalHostDeleteOK {
	return &LocalHostDeleteOK{}
}

/*
LocalHostDeleteOK describes a response with status code 200, with default header values.

OK
*/
type LocalHostDeleteOK struct {
}

// IsSuccess returns true when this local host delete o k response has a 2xx status code
func (o *LocalHostDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this local host delete o k response has a 3xx status code
func (o *LocalHostDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this local host delete o k response has a 4xx status code
func (o *LocalHostDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this local host delete o k response has a 5xx status code
func (o *LocalHostDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this local host delete o k response a status code equal to that given
func (o *LocalHostDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the local host delete o k response
func (o *LocalHostDeleteOK) Code() int {
	return 200
}

func (o *LocalHostDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /name-services/local-hosts/{owner.uuid}/{address}][%d] localHostDeleteOK", 200)
}

func (o *LocalHostDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /name-services/local-hosts/{owner.uuid}/{address}][%d] localHostDeleteOK", 200)
}

func (o *LocalHostDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLocalHostDeleteDefault creates a LocalHostDeleteDefault with default headers values
func NewLocalHostDeleteDefault(code int) *LocalHostDeleteDefault {
	return &LocalHostDeleteDefault{
		_statusCode: code,
	}
}

/*
	LocalHostDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 23724055 | Internal error. Configuration for Vserver failed. Verify that the cluster is healthy, then try the command again. For further assistance, contact technical support. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type LocalHostDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this local host delete default response has a 2xx status code
func (o *LocalHostDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this local host delete default response has a 3xx status code
func (o *LocalHostDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this local host delete default response has a 4xx status code
func (o *LocalHostDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this local host delete default response has a 5xx status code
func (o *LocalHostDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this local host delete default response a status code equal to that given
func (o *LocalHostDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the local host delete default response
func (o *LocalHostDeleteDefault) Code() int {
	return o._statusCode
}

func (o *LocalHostDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /name-services/local-hosts/{owner.uuid}/{address}][%d] local_host_delete default %s", o._statusCode, payload)
}

func (o *LocalHostDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /name-services/local-hosts/{owner.uuid}/{address}][%d] local_host_delete default %s", o._statusCode, payload)
}

func (o *LocalHostDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LocalHostDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
