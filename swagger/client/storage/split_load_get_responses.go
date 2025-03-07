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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// SplitLoadGetReader is a Reader for the SplitLoadGet structure.
type SplitLoadGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SplitLoadGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSplitLoadGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSplitLoadGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSplitLoadGetOK creates a SplitLoadGetOK with default headers values
func NewSplitLoadGetOK() *SplitLoadGetOK {
	return &SplitLoadGetOK{}
}

/*
SplitLoadGetOK describes a response with status code 200, with default header values.

OK
*/
type SplitLoadGetOK struct {
	Payload *models.SplitLoad
}

// IsSuccess returns true when this split load get o k response has a 2xx status code
func (o *SplitLoadGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this split load get o k response has a 3xx status code
func (o *SplitLoadGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this split load get o k response has a 4xx status code
func (o *SplitLoadGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this split load get o k response has a 5xx status code
func (o *SplitLoadGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this split load get o k response a status code equal to that given
func (o *SplitLoadGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the split load get o k response
func (o *SplitLoadGetOK) Code() int {
	return 200
}

func (o *SplitLoadGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/file/clone/split-loads/{node.uuid}][%d] splitLoadGetOK %s", 200, payload)
}

func (o *SplitLoadGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/file/clone/split-loads/{node.uuid}][%d] splitLoadGetOK %s", 200, payload)
}

func (o *SplitLoadGetOK) GetPayload() *models.SplitLoad {
	return o.Payload
}

func (o *SplitLoadGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SplitLoad)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSplitLoadGetDefault creates a SplitLoadGetDefault with default headers values
func NewSplitLoadGetDefault(code int) *SplitLoadGetDefault {
	return &SplitLoadGetDefault{
		_statusCode: code,
	}
}

/*
SplitLoadGetDefault describes a response with status code -1, with default header values.

Error
*/
type SplitLoadGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this split load get default response has a 2xx status code
func (o *SplitLoadGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this split load get default response has a 3xx status code
func (o *SplitLoadGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this split load get default response has a 4xx status code
func (o *SplitLoadGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this split load get default response has a 5xx status code
func (o *SplitLoadGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this split load get default response a status code equal to that given
func (o *SplitLoadGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the split load get default response
func (o *SplitLoadGetDefault) Code() int {
	return o._statusCode
}

func (o *SplitLoadGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/file/clone/split-loads/{node.uuid}][%d] split_load_get default %s", o._statusCode, payload)
}

func (o *SplitLoadGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/file/clone/split-loads/{node.uuid}][%d] split_load_get default %s", o._statusCode, payload)
}

func (o *SplitLoadGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SplitLoadGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
