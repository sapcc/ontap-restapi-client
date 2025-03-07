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

// HostsSettingsGetReader is a Reader for the HostsSettingsGet structure.
type HostsSettingsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostsSettingsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHostsSettingsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHostsSettingsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHostsSettingsGetOK creates a HostsSettingsGetOK with default headers values
func NewHostsSettingsGetOK() *HostsSettingsGetOK {
	return &HostsSettingsGetOK{}
}

/*
HostsSettingsGetOK describes a response with status code 200, with default header values.

OK
*/
type HostsSettingsGetOK struct {
	Payload *models.HostsSettings
}

// IsSuccess returns true when this hosts settings get o k response has a 2xx status code
func (o *HostsSettingsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hosts settings get o k response has a 3xx status code
func (o *HostsSettingsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hosts settings get o k response has a 4xx status code
func (o *HostsSettingsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hosts settings get o k response has a 5xx status code
func (o *HostsSettingsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hosts settings get o k response a status code equal to that given
func (o *HostsSettingsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the hosts settings get o k response
func (o *HostsSettingsGetOK) Code() int {
	return 200
}

func (o *HostsSettingsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/cache/host/settings/{uuid}][%d] hostsSettingsGetOK %s", 200, payload)
}

func (o *HostsSettingsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/cache/host/settings/{uuid}][%d] hostsSettingsGetOK %s", 200, payload)
}

func (o *HostsSettingsGetOK) GetPayload() *models.HostsSettings {
	return o.Payload
}

func (o *HostsSettingsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostsSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHostsSettingsGetDefault creates a HostsSettingsGetDefault with default headers values
func NewHostsSettingsGetDefault(code int) *HostsSettingsGetDefault {
	return &HostsSettingsGetDefault{
		_statusCode: code,
	}
}

/*
HostsSettingsGetDefault describes a response with status code -1, with default header values.

Error
*/
type HostsSettingsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this hosts settings get default response has a 2xx status code
func (o *HostsSettingsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hosts settings get default response has a 3xx status code
func (o *HostsSettingsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hosts settings get default response has a 4xx status code
func (o *HostsSettingsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hosts settings get default response has a 5xx status code
func (o *HostsSettingsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hosts settings get default response a status code equal to that given
func (o *HostsSettingsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the hosts settings get default response
func (o *HostsSettingsGetDefault) Code() int {
	return o._statusCode
}

func (o *HostsSettingsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/cache/host/settings/{uuid}][%d] hosts_settings_get default %s", o._statusCode, payload)
}

func (o *HostsSettingsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/cache/host/settings/{uuid}][%d] hosts_settings_get default %s", o._statusCode, payload)
}

func (o *HostsSettingsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *HostsSettingsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
