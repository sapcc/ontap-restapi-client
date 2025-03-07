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

// SplitLoadModifyReader is a Reader for the SplitLoadModify structure.
type SplitLoadModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SplitLoadModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSplitLoadModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSplitLoadModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSplitLoadModifyOK creates a SplitLoadModifyOK with default headers values
func NewSplitLoadModifyOK() *SplitLoadModifyOK {
	return &SplitLoadModifyOK{}
}

/*
SplitLoadModifyOK describes a response with status code 200, with default header values.

OK
*/
type SplitLoadModifyOK struct {
}

// IsSuccess returns true when this split load modify o k response has a 2xx status code
func (o *SplitLoadModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this split load modify o k response has a 3xx status code
func (o *SplitLoadModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this split load modify o k response has a 4xx status code
func (o *SplitLoadModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this split load modify o k response has a 5xx status code
func (o *SplitLoadModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this split load modify o k response a status code equal to that given
func (o *SplitLoadModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the split load modify o k response
func (o *SplitLoadModifyOK) Code() int {
	return 200
}

func (o *SplitLoadModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /storage/file/clone/split-loads/{node.uuid}][%d] splitLoadModifyOK", 200)
}

func (o *SplitLoadModifyOK) String() string {
	return fmt.Sprintf("[PATCH /storage/file/clone/split-loads/{node.uuid}][%d] splitLoadModifyOK", 200)
}

func (o *SplitLoadModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSplitLoadModifyDefault creates a SplitLoadModifyDefault with default headers values
func NewSplitLoadModifyDefault(code int) *SplitLoadModifyDefault {
	return &SplitLoadModifyDefault{
		_statusCode: code,
	}
}

/*
	SplitLoadModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 13565966 | Failed to modify the maximum split load on the node due to an invalid field value. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SplitLoadModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this split load modify default response has a 2xx status code
func (o *SplitLoadModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this split load modify default response has a 3xx status code
func (o *SplitLoadModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this split load modify default response has a 4xx status code
func (o *SplitLoadModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this split load modify default response has a 5xx status code
func (o *SplitLoadModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this split load modify default response a status code equal to that given
func (o *SplitLoadModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the split load modify default response
func (o *SplitLoadModifyDefault) Code() int {
	return o._statusCode
}

func (o *SplitLoadModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/file/clone/split-loads/{node.uuid}][%d] split_load_modify default %s", o._statusCode, payload)
}

func (o *SplitLoadModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/file/clone/split-loads/{node.uuid}][%d] split_load_modify default %s", o._statusCode, payload)
}

func (o *SplitLoadModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SplitLoadModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
