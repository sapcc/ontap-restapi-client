// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// SplitStatusGetReader is a Reader for the SplitStatusGet structure.
type SplitStatusGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SplitStatusGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSplitStatusGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSplitStatusGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSplitStatusGetOK creates a SplitStatusGetOK with default headers values
func NewSplitStatusGetOK() *SplitStatusGetOK {
	return &SplitStatusGetOK{}
}

/*
SplitStatusGetOK describes a response with status code 200, with default header values.

OK
*/
type SplitStatusGetOK struct {
	Payload *models.SplitStatus
}

// IsSuccess returns true when this split status get o k response has a 2xx status code
func (o *SplitStatusGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this split status get o k response has a 3xx status code
func (o *SplitStatusGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this split status get o k response has a 4xx status code
func (o *SplitStatusGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this split status get o k response has a 5xx status code
func (o *SplitStatusGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this split status get o k response a status code equal to that given
func (o *SplitStatusGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the split status get o k response
func (o *SplitStatusGetOK) Code() int {
	return 200
}

func (o *SplitStatusGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/file/clone/split-status/{volume.uuid}][%d] splitStatusGetOK %s", 200, payload)
}

func (o *SplitStatusGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/file/clone/split-status/{volume.uuid}][%d] splitStatusGetOK %s", 200, payload)
}

func (o *SplitStatusGetOK) GetPayload() *models.SplitStatus {
	return o.Payload
}

func (o *SplitStatusGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SplitStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSplitStatusGetDefault creates a SplitStatusGetDefault with default headers values
func NewSplitStatusGetDefault(code int) *SplitStatusGetDefault {
	return &SplitStatusGetDefault{
		_statusCode: code,
	}
}

/*
SplitStatusGetDefault describes a response with status code -1, with default header values.

Error
*/
type SplitStatusGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this split status get default response has a 2xx status code
func (o *SplitStatusGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this split status get default response has a 3xx status code
func (o *SplitStatusGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this split status get default response has a 4xx status code
func (o *SplitStatusGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this split status get default response has a 5xx status code
func (o *SplitStatusGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this split status get default response a status code equal to that given
func (o *SplitStatusGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the split status get default response
func (o *SplitStatusGetDefault) Code() int {
	return o._statusCode
}

func (o *SplitStatusGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/file/clone/split-status/{volume.uuid}][%d] split_status_get default %s", o._statusCode, payload)
}

func (o *SplitStatusGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/file/clone/split-status/{volume.uuid}][%d] split_status_get default %s", o._statusCode, payload)
}

func (o *SplitStatusGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SplitStatusGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
