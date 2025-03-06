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

// FileInfoCollectionGetReader is a Reader for the FileInfoCollectionGet structure.
type FileInfoCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileInfoCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFileInfoCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFileInfoCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileInfoCollectionGetOK creates a FileInfoCollectionGetOK with default headers values
func NewFileInfoCollectionGetOK() *FileInfoCollectionGetOK {
	return &FileInfoCollectionGetOK{}
}

/*
FileInfoCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type FileInfoCollectionGetOK struct {
	Payload *models.FileInfoResponse
}

// IsSuccess returns true when this file info collection get o k response has a 2xx status code
func (o *FileInfoCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file info collection get o k response has a 3xx status code
func (o *FileInfoCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file info collection get o k response has a 4xx status code
func (o *FileInfoCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this file info collection get o k response has a 5xx status code
func (o *FileInfoCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this file info collection get o k response a status code equal to that given
func (o *FileInfoCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the file info collection get o k response
func (o *FileInfoCollectionGetOK) Code() int {
	return 200
}

func (o *FileInfoCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volumes/{volume.uuid}/files/{path}][%d] fileInfoCollectionGetOK %s", 200, payload)
}

func (o *FileInfoCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volumes/{volume.uuid}/files/{path}][%d] fileInfoCollectionGetOK %s", 200, payload)
}

func (o *FileInfoCollectionGetOK) GetPayload() *models.FileInfoResponse {
	return o.Payload
}

func (o *FileInfoCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileInfoCollectionGetDefault creates a FileInfoCollectionGetDefault with default headers values
func NewFileInfoCollectionGetDefault(code int) *FileInfoCollectionGetDefault {
	return &FileInfoCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	FileInfoCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 917752 | Volume is offline. |
| 918235 | A volume with UUID {volume.uuid} was not found. |
| 6488109 | Operation not supported on FlexCache volumes. |
| 6684674 | No such file or directory. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type FileInfoCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this file info collection get default response has a 2xx status code
func (o *FileInfoCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this file info collection get default response has a 3xx status code
func (o *FileInfoCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this file info collection get default response has a 4xx status code
func (o *FileInfoCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this file info collection get default response has a 5xx status code
func (o *FileInfoCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this file info collection get default response a status code equal to that given
func (o *FileInfoCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the file info collection get default response
func (o *FileInfoCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *FileInfoCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volumes/{volume.uuid}/files/{path}][%d] file_info_collection_get default %s", o._statusCode, payload)
}

func (o *FileInfoCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volumes/{volume.uuid}/files/{path}][%d] file_info_collection_get default %s", o._statusCode, payload)
}

func (o *FileInfoCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FileInfoCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
