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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// LunAttributeDeleteReader is a Reader for the LunAttributeDelete structure.
type LunAttributeDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LunAttributeDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLunAttributeDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLunAttributeDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLunAttributeDeleteOK creates a LunAttributeDeleteOK with default headers values
func NewLunAttributeDeleteOK() *LunAttributeDeleteOK {
	return &LunAttributeDeleteOK{}
}

/*
LunAttributeDeleteOK describes a response with status code 200, with default header values.

OK
*/
type LunAttributeDeleteOK struct {
}

// IsSuccess returns true when this lun attribute delete o k response has a 2xx status code
func (o *LunAttributeDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lun attribute delete o k response has a 3xx status code
func (o *LunAttributeDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lun attribute delete o k response has a 4xx status code
func (o *LunAttributeDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lun attribute delete o k response has a 5xx status code
func (o *LunAttributeDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lun attribute delete o k response a status code equal to that given
func (o *LunAttributeDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lun attribute delete o k response
func (o *LunAttributeDeleteOK) Code() int {
	return 200
}

func (o *LunAttributeDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /storage/luns/{lun.uuid}/attributes/{name}][%d] lunAttributeDeleteOK", 200)
}

func (o *LunAttributeDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /storage/luns/{lun.uuid}/attributes/{name}][%d] lunAttributeDeleteOK", 200)
}

func (o *LunAttributeDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLunAttributeDeleteDefault creates a LunAttributeDeleteDefault with default headers values
func NewLunAttributeDeleteDefault(code int) *LunAttributeDeleteDefault {
	return &LunAttributeDeleteDefault{
		_statusCode: code,
	}
}

/*
	LunAttributeDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5374875 | The specified LUN was not found. |
| 5374931 | The specified attribute was not found. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type LunAttributeDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this lun attribute delete default response has a 2xx status code
func (o *LunAttributeDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lun attribute delete default response has a 3xx status code
func (o *LunAttributeDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lun attribute delete default response has a 4xx status code
func (o *LunAttributeDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lun attribute delete default response has a 5xx status code
func (o *LunAttributeDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lun attribute delete default response a status code equal to that given
func (o *LunAttributeDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the lun attribute delete default response
func (o *LunAttributeDeleteDefault) Code() int {
	return o._statusCode
}

func (o *LunAttributeDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/luns/{lun.uuid}/attributes/{name}][%d] lun_attribute_delete default %s", o._statusCode, payload)
}

func (o *LunAttributeDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/luns/{lun.uuid}/attributes/{name}][%d] lun_attribute_delete default %s", o._statusCode, payload)
}

func (o *LunAttributeDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LunAttributeDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
