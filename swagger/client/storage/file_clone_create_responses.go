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

// FileCloneCreateReader is a Reader for the FileCloneCreate structure.
type FileCloneCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileCloneCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewFileCloneCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewFileCloneCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFileCloneCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileCloneCreateCreated creates a FileCloneCreateCreated with default headers values
func NewFileCloneCreateCreated() *FileCloneCreateCreated {
	return &FileCloneCreateCreated{}
}

/*
FileCloneCreateCreated describes a response with status code 201, with default header values.

Created
*/
type FileCloneCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this file clone create created response has a 2xx status code
func (o *FileCloneCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file clone create created response has a 3xx status code
func (o *FileCloneCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file clone create created response has a 4xx status code
func (o *FileCloneCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this file clone create created response has a 5xx status code
func (o *FileCloneCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this file clone create created response a status code equal to that given
func (o *FileCloneCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the file clone create created response
func (o *FileCloneCreateCreated) Code() int {
	return 201
}

func (o *FileCloneCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/file/clone][%d] fileCloneCreateCreated %s", 201, payload)
}

func (o *FileCloneCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/file/clone][%d] fileCloneCreateCreated %s", 201, payload)
}

func (o *FileCloneCreateCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *FileCloneCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileCloneCreateAccepted creates a FileCloneCreateAccepted with default headers values
func NewFileCloneCreateAccepted() *FileCloneCreateAccepted {
	return &FileCloneCreateAccepted{}
}

/*
FileCloneCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type FileCloneCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this file clone create accepted response has a 2xx status code
func (o *FileCloneCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file clone create accepted response has a 3xx status code
func (o *FileCloneCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file clone create accepted response has a 4xx status code
func (o *FileCloneCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this file clone create accepted response has a 5xx status code
func (o *FileCloneCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this file clone create accepted response a status code equal to that given
func (o *FileCloneCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the file clone create accepted response
func (o *FileCloneCreateAccepted) Code() int {
	return 202
}

func (o *FileCloneCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/file/clone][%d] fileCloneCreateAccepted %s", 202, payload)
}

func (o *FileCloneCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/file/clone][%d] fileCloneCreateAccepted %s", 202, payload)
}

func (o *FileCloneCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *FileCloneCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileCloneCreateDefault creates a FileCloneCreateDefault with default headers values
func NewFileCloneCreateDefault(code int) *FileCloneCreateDefault {
	return &FileCloneCreateDefault{
		_statusCode: code,
	}
}

/*
	FileCloneCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 917864 | Failed to get the file handle. |
| 917898 | Invalid `value` for the field -range. Value must be provided in the format \<source start block\>:\<destination start block\>:\<block length\>. |
| 918236 | The specified `volume.uuid` and `volume.name` refer to different volumes. |
| 13565952 | Clone start failed. |
| 13565988 | The destination file cannot be marked as a backup clone as only a full-file clone can be a backup clone. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type FileCloneCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this file clone create default response has a 2xx status code
func (o *FileCloneCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this file clone create default response has a 3xx status code
func (o *FileCloneCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this file clone create default response has a 4xx status code
func (o *FileCloneCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this file clone create default response has a 5xx status code
func (o *FileCloneCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this file clone create default response a status code equal to that given
func (o *FileCloneCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the file clone create default response
func (o *FileCloneCreateDefault) Code() int {
	return o._statusCode
}

func (o *FileCloneCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/file/clone][%d] file_clone_create default %s", o._statusCode, payload)
}

func (o *FileCloneCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/file/clone][%d] file_clone_create default %s", o._statusCode, payload)
}

func (o *FileCloneCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FileCloneCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
