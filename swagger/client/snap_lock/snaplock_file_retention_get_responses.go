// Code generated by go-swagger; DO NOT EDIT.

package snap_lock

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

// SnaplockFileRetentionGetReader is a Reader for the SnaplockFileRetentionGet structure.
type SnaplockFileRetentionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockFileRetentionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockFileRetentionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockFileRetentionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockFileRetentionGetOK creates a SnaplockFileRetentionGetOK with default headers values
func NewSnaplockFileRetentionGetOK() *SnaplockFileRetentionGetOK {
	return &SnaplockFileRetentionGetOK{}
}

/*
SnaplockFileRetentionGetOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockFileRetentionGetOK struct {
	Payload *models.SnaplockFileRetention
}

// IsSuccess returns true when this snaplock file retention get o k response has a 2xx status code
func (o *SnaplockFileRetentionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock file retention get o k response has a 3xx status code
func (o *SnaplockFileRetentionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock file retention get o k response has a 4xx status code
func (o *SnaplockFileRetentionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock file retention get o k response has a 5xx status code
func (o *SnaplockFileRetentionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock file retention get o k response a status code equal to that given
func (o *SnaplockFileRetentionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snaplock file retention get o k response
func (o *SnaplockFileRetentionGetOK) Code() int {
	return 200
}

func (o *SnaplockFileRetentionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/file/{volume.uuid}/{path}][%d] snaplockFileRetentionGetOK %s", 200, payload)
}

func (o *SnaplockFileRetentionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/file/{volume.uuid}/{path}][%d] snaplockFileRetentionGetOK %s", 200, payload)
}

func (o *SnaplockFileRetentionGetOK) GetPayload() *models.SnaplockFileRetention {
	return o.Payload
}

func (o *SnaplockFileRetentionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockFileRetention)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockFileRetentionGetDefault creates a SnaplockFileRetentionGetDefault with default headers values
func NewSnaplockFileRetentionGetDefault(code int) *SnaplockFileRetentionGetDefault {
	return &SnaplockFileRetentionGetDefault{
		_statusCode: code,
	}
}

/*
	SnaplockFileRetentionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 262179      | Unexpected argument \"<file_name>\" |
| 917864      | Operation not supported |
| 6691623     | User is not authorized  |
| 13762592    | Operation not supported on non-SnapLock volume |
| 14090347    | File path must be in the format \"\/\<dir\>\/\<file path\>\" |
| 917804      | Path should be given in the format \"\/\vol\/\<volume name>\/\<file path>\". |
*/
type SnaplockFileRetentionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snaplock file retention get default response has a 2xx status code
func (o *SnaplockFileRetentionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock file retention get default response has a 3xx status code
func (o *SnaplockFileRetentionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock file retention get default response has a 4xx status code
func (o *SnaplockFileRetentionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock file retention get default response has a 5xx status code
func (o *SnaplockFileRetentionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock file retention get default response a status code equal to that given
func (o *SnaplockFileRetentionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snaplock file retention get default response
func (o *SnaplockFileRetentionGetDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockFileRetentionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/file/{volume.uuid}/{path}][%d] snaplock_file_retention_get default %s", o._statusCode, payload)
}

func (o *SnaplockFileRetentionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/file/{volume.uuid}/{path}][%d] snaplock_file_retention_get default %s", o._statusCode, payload)
}

func (o *SnaplockFileRetentionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockFileRetentionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
