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

// LunGetReader is a Reader for the LunGet structure.
type LunGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LunGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLunGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLunGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLunGetOK creates a LunGetOK with default headers values
func NewLunGetOK() *LunGetOK {
	return &LunGetOK{}
}

/*
LunGetOK describes a response with status code 200, with default header values.

OK
*/
type LunGetOK struct {
	Payload *models.Lun
}

// IsSuccess returns true when this lun get o k response has a 2xx status code
func (o *LunGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lun get o k response has a 3xx status code
func (o *LunGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lun get o k response has a 4xx status code
func (o *LunGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lun get o k response has a 5xx status code
func (o *LunGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lun get o k response a status code equal to that given
func (o *LunGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lun get o k response
func (o *LunGetOK) Code() int {
	return 200
}

func (o *LunGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/luns/{uuid}][%d] lunGetOK %s", 200, payload)
}

func (o *LunGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/luns/{uuid}][%d] lunGetOK %s", 200, payload)
}

func (o *LunGetOK) GetPayload() *models.Lun {
	return o.Payload
}

func (o *LunGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Lun)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLunGetDefault creates a LunGetDefault with default headers values
func NewLunGetDefault(code int) *LunGetDefault {
	return &LunGetDefault{
		_statusCode: code,
	}
}

/*
	LunGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5374875 | The specified LUN was not found. |
| 5374876 | The specified LUN was not found. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type LunGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this lun get default response has a 2xx status code
func (o *LunGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lun get default response has a 3xx status code
func (o *LunGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lun get default response has a 4xx status code
func (o *LunGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lun get default response has a 5xx status code
func (o *LunGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lun get default response a status code equal to that given
func (o *LunGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the lun get default response
func (o *LunGetDefault) Code() int {
	return o._statusCode
}

func (o *LunGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/luns/{uuid}][%d] lun_get default %s", o._statusCode, payload)
}

func (o *LunGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/luns/{uuid}][%d] lun_get default %s", o._statusCode, payload)
}

func (o *LunGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LunGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
