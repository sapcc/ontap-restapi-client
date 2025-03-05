// Code generated by go-swagger; DO NOT EDIT.

package support

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

// EmsRoleConfigGetReader is a Reader for the EmsRoleConfigGet structure.
type EmsRoleConfigGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsRoleConfigGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmsRoleConfigGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsRoleConfigGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsRoleConfigGetOK creates a EmsRoleConfigGetOK with default headers values
func NewEmsRoleConfigGetOK() *EmsRoleConfigGetOK {
	return &EmsRoleConfigGetOK{}
}

/*
EmsRoleConfigGetOK describes a response with status code 200, with default header values.

OK
*/
type EmsRoleConfigGetOK struct {
	Payload *models.EmsRoleConfig
}

// IsSuccess returns true when this ems role config get o k response has a 2xx status code
func (o *EmsRoleConfigGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ems role config get o k response has a 3xx status code
func (o *EmsRoleConfigGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ems role config get o k response has a 4xx status code
func (o *EmsRoleConfigGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ems role config get o k response has a 5xx status code
func (o *EmsRoleConfigGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ems role config get o k response a status code equal to that given
func (o *EmsRoleConfigGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ems role config get o k response
func (o *EmsRoleConfigGetOK) Code() int {
	return 200
}

func (o *EmsRoleConfigGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/role-configs/{access_control_role.name}][%d] emsRoleConfigGetOK %s", 200, payload)
}

func (o *EmsRoleConfigGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/role-configs/{access_control_role.name}][%d] emsRoleConfigGetOK %s", 200, payload)
}

func (o *EmsRoleConfigGetOK) GetPayload() *models.EmsRoleConfig {
	return o.Payload
}

func (o *EmsRoleConfigGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmsRoleConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmsRoleConfigGetDefault creates a EmsRoleConfigGetDefault with default headers values
func NewEmsRoleConfigGetDefault(code int) *EmsRoleConfigGetDefault {
	return &EmsRoleConfigGetDefault{
		_statusCode: code,
	}
}

/*
EmsRoleConfigGetDefault describes a response with status code -1, with default header values.

Error
*/
type EmsRoleConfigGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ems role config get default response has a 2xx status code
func (o *EmsRoleConfigGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ems role config get default response has a 3xx status code
func (o *EmsRoleConfigGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ems role config get default response has a 4xx status code
func (o *EmsRoleConfigGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ems role config get default response has a 5xx status code
func (o *EmsRoleConfigGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ems role config get default response a status code equal to that given
func (o *EmsRoleConfigGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ems role config get default response
func (o *EmsRoleConfigGetDefault) Code() int {
	return o._statusCode
}

func (o *EmsRoleConfigGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/role-configs/{access_control_role.name}][%d] ems_role_config_get default %s", o._statusCode, payload)
}

func (o *EmsRoleConfigGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/role-configs/{access_control_role.name}][%d] ems_role_config_get default %s", o._statusCode, payload)
}

func (o *EmsRoleConfigGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsRoleConfigGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
