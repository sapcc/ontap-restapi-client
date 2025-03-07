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

// FileDeleteReader is a Reader for the FileDelete structure.
type FileDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFileDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewFileDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFileDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileDeleteOK creates a FileDeleteOK with default headers values
func NewFileDeleteOK() *FileDeleteOK {
	return &FileDeleteOK{}
}

/*
FileDeleteOK describes a response with status code 200, with default header values.

OK
*/
type FileDeleteOK struct {
}

// IsSuccess returns true when this file delete o k response has a 2xx status code
func (o *FileDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file delete o k response has a 3xx status code
func (o *FileDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file delete o k response has a 4xx status code
func (o *FileDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this file delete o k response has a 5xx status code
func (o *FileDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this file delete o k response a status code equal to that given
func (o *FileDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the file delete o k response
func (o *FileDeleteOK) Code() int {
	return 200
}

func (o *FileDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /storage/volumes/{volume.uuid}/files/{path}][%d] fileDeleteOK", 200)
}

func (o *FileDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /storage/volumes/{volume.uuid}/files/{path}][%d] fileDeleteOK", 200)
}

func (o *FileDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFileDeleteAccepted creates a FileDeleteAccepted with default headers values
func NewFileDeleteAccepted() *FileDeleteAccepted {
	return &FileDeleteAccepted{}
}

/*
FileDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type FileDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this file delete accepted response has a 2xx status code
func (o *FileDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file delete accepted response has a 3xx status code
func (o *FileDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file delete accepted response has a 4xx status code
func (o *FileDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this file delete accepted response has a 5xx status code
func (o *FileDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this file delete accepted response a status code equal to that given
func (o *FileDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the file delete accepted response
func (o *FileDeleteAccepted) Code() int {
	return 202
}

func (o *FileDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/volumes/{volume.uuid}/files/{path}][%d] fileDeleteAccepted %s", 202, payload)
}

func (o *FileDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/volumes/{volume.uuid}/files/{path}][%d] fileDeleteAccepted %s", 202, payload)
}

func (o *FileDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *FileDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileDeleteDefault creates a FileDeleteDefault with default headers values
func NewFileDeleteDefault(code int) *FileDeleteDefault {
	return &FileDeleteDefault{
		_statusCode: code,
	}
}

/*
	FileDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 131102 | Read-only file system. |
| 131138 | Directory not empty. |
| 918235 | A volume with UUID {volume.uuid} was not found. |
| 6488081 | The {field} field is not supported for DELETE operations. |
| 6488110 | A volume delete is not supported on this endpoint. |
| 6684674 | No such file or directory. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type FileDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this file delete default response has a 2xx status code
func (o *FileDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this file delete default response has a 3xx status code
func (o *FileDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this file delete default response has a 4xx status code
func (o *FileDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this file delete default response has a 5xx status code
func (o *FileDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this file delete default response a status code equal to that given
func (o *FileDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the file delete default response
func (o *FileDeleteDefault) Code() int {
	return o._statusCode
}

func (o *FileDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/volumes/{volume.uuid}/files/{path}][%d] file_delete default %s", o._statusCode, payload)
}

func (o *FileDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/volumes/{volume.uuid}/files/{path}][%d] file_delete default %s", o._statusCode, payload)
}

func (o *FileDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FileDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
