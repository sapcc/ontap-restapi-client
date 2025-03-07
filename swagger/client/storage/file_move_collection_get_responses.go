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

// FileMoveCollectionGetReader is a Reader for the FileMoveCollectionGet structure.
type FileMoveCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileMoveCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFileMoveCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFileMoveCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileMoveCollectionGetOK creates a FileMoveCollectionGetOK with default headers values
func NewFileMoveCollectionGetOK() *FileMoveCollectionGetOK {
	return &FileMoveCollectionGetOK{}
}

/*
FileMoveCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type FileMoveCollectionGetOK struct {
	Payload *models.FileMoveResponse
}

// IsSuccess returns true when this file move collection get o k response has a 2xx status code
func (o *FileMoveCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file move collection get o k response has a 3xx status code
func (o *FileMoveCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file move collection get o k response has a 4xx status code
func (o *FileMoveCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this file move collection get o k response has a 5xx status code
func (o *FileMoveCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this file move collection get o k response a status code equal to that given
func (o *FileMoveCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the file move collection get o k response
func (o *FileMoveCollectionGetOK) Code() int {
	return 200
}

func (o *FileMoveCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/file/moves][%d] fileMoveCollectionGetOK %s", 200, payload)
}

func (o *FileMoveCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/file/moves][%d] fileMoveCollectionGetOK %s", 200, payload)
}

func (o *FileMoveCollectionGetOK) GetPayload() *models.FileMoveResponse {
	return o.Payload
}

func (o *FileMoveCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileMoveResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileMoveCollectionGetDefault creates a FileMoveCollectionGetDefault with default headers values
func NewFileMoveCollectionGetDefault(code int) *FileMoveCollectionGetDefault {
	return &FileMoveCollectionGetDefault{
		_statusCode: code,
	}
}

/*
FileMoveCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type FileMoveCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this file move collection get default response has a 2xx status code
func (o *FileMoveCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this file move collection get default response has a 3xx status code
func (o *FileMoveCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this file move collection get default response has a 4xx status code
func (o *FileMoveCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this file move collection get default response has a 5xx status code
func (o *FileMoveCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this file move collection get default response a status code equal to that given
func (o *FileMoveCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the file move collection get default response
func (o *FileMoveCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *FileMoveCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/file/moves][%d] file_move_collection_get default %s", o._statusCode, payload)
}

func (o *FileMoveCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/file/moves][%d] file_move_collection_get default %s", o._statusCode, payload)
}

func (o *FileMoveCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FileMoveCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
