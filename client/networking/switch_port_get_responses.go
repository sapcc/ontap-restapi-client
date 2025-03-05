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

	"github.com/sapcc/ontap-restapi/models"
)

// SwitchPortGetReader is a Reader for the SwitchPortGet structure.
type SwitchPortGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SwitchPortGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSwitchPortGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSwitchPortGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSwitchPortGetOK creates a SwitchPortGetOK with default headers values
func NewSwitchPortGetOK() *SwitchPortGetOK {
	return &SwitchPortGetOK{}
}

/*
SwitchPortGetOK describes a response with status code 200, with default header values.

OK
*/
type SwitchPortGetOK struct {
	Payload *models.SwitchPort
}

// IsSuccess returns true when this switch port get o k response has a 2xx status code
func (o *SwitchPortGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this switch port get o k response has a 3xx status code
func (o *SwitchPortGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this switch port get o k response has a 4xx status code
func (o *SwitchPortGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this switch port get o k response has a 5xx status code
func (o *SwitchPortGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this switch port get o k response a status code equal to that given
func (o *SwitchPortGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the switch port get o k response
func (o *SwitchPortGetOK) Code() int {
	return 200
}

func (o *SwitchPortGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/switch/ports/{switch}/{identity.name}/{identity.index}][%d] switchPortGetOK %s", 200, payload)
}

func (o *SwitchPortGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/switch/ports/{switch}/{identity.name}/{identity.index}][%d] switchPortGetOK %s", 200, payload)
}

func (o *SwitchPortGetOK) GetPayload() *models.SwitchPort {
	return o.Payload
}

func (o *SwitchPortGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SwitchPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSwitchPortGetDefault creates a SwitchPortGetDefault with default headers values
func NewSwitchPortGetDefault(code int) *SwitchPortGetDefault {
	return &SwitchPortGetDefault{
		_statusCode: code,
	}
}

/*
SwitchPortGetDefault describes a response with status code -1, with default header values.

Error
*/
type SwitchPortGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this switch port get default response has a 2xx status code
func (o *SwitchPortGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this switch port get default response has a 3xx status code
func (o *SwitchPortGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this switch port get default response has a 4xx status code
func (o *SwitchPortGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this switch port get default response has a 5xx status code
func (o *SwitchPortGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this switch port get default response a status code equal to that given
func (o *SwitchPortGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the switch port get default response
func (o *SwitchPortGetDefault) Code() int {
	return o._statusCode
}

func (o *SwitchPortGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/switch/ports/{switch}/{identity.name}/{identity.index}][%d] switch_port_get default %s", o._statusCode, payload)
}

func (o *SwitchPortGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/switch/ports/{switch}/{identity.name}/{identity.index}][%d] switch_port_get default %s", o._statusCode, payload)
}

func (o *SwitchPortGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SwitchPortGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
