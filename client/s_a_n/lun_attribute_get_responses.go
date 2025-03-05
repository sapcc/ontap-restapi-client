// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// LunAttributeGetReader is a Reader for the LunAttributeGet structure.
type LunAttributeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LunAttributeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLunAttributeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLunAttributeGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLunAttributeGetOK creates a LunAttributeGetOK with default headers values
func NewLunAttributeGetOK() *LunAttributeGetOK {
	return &LunAttributeGetOK{}
}

/*
LunAttributeGetOK describes a response with status code 200, with default header values.

OK
*/
type LunAttributeGetOK struct {
	Payload *models.LunAttribute
}

// IsSuccess returns true when this lun attribute get o k response has a 2xx status code
func (o *LunAttributeGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lun attribute get o k response has a 3xx status code
func (o *LunAttributeGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lun attribute get o k response has a 4xx status code
func (o *LunAttributeGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lun attribute get o k response has a 5xx status code
func (o *LunAttributeGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lun attribute get o k response a status code equal to that given
func (o *LunAttributeGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lun attribute get o k response
func (o *LunAttributeGetOK) Code() int {
	return 200
}

func (o *LunAttributeGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/luns/{lun.uuid}/attributes/{name}][%d] lunAttributeGetOK %s", 200, payload)
}

func (o *LunAttributeGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/luns/{lun.uuid}/attributes/{name}][%d] lunAttributeGetOK %s", 200, payload)
}

func (o *LunAttributeGetOK) GetPayload() *models.LunAttribute {
	return o.Payload
}

func (o *LunAttributeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LunAttribute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLunAttributeGetDefault creates a LunAttributeGetDefault with default headers values
func NewLunAttributeGetDefault(code int) *LunAttributeGetDefault {
	return &LunAttributeGetDefault{
		_statusCode: code,
	}
}

/*
	LunAttributeGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5374875 | The specified LUN was not found. |
| 5374931 | The specified attribute was not found. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type LunAttributeGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this lun attribute get default response has a 2xx status code
func (o *LunAttributeGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lun attribute get default response has a 3xx status code
func (o *LunAttributeGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lun attribute get default response has a 4xx status code
func (o *LunAttributeGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lun attribute get default response has a 5xx status code
func (o *LunAttributeGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lun attribute get default response a status code equal to that given
func (o *LunAttributeGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the lun attribute get default response
func (o *LunAttributeGetDefault) Code() int {
	return o._statusCode
}

func (o *LunAttributeGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/luns/{lun.uuid}/attributes/{name}][%d] lun_attribute_get default %s", o._statusCode, payload)
}

func (o *LunAttributeGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/luns/{lun.uuid}/attributes/{name}][%d] lun_attribute_get default %s", o._statusCode, payload)
}

func (o *LunAttributeGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LunAttributeGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
