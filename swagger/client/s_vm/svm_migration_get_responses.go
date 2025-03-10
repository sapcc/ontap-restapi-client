// Code generated by go-swagger; DO NOT EDIT.

package s_vm

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

// SvmMigrationGetReader is a Reader for the SvmMigrationGet structure.
type SvmMigrationGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmMigrationGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSvmMigrationGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmMigrationGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmMigrationGetOK creates a SvmMigrationGetOK with default headers values
func NewSvmMigrationGetOK() *SvmMigrationGetOK {
	return &SvmMigrationGetOK{}
}

/*
SvmMigrationGetOK describes a response with status code 200, with default header values.

OK
*/
type SvmMigrationGetOK struct {
	Payload *models.SvmMigration
}

// IsSuccess returns true when this svm migration get o k response has a 2xx status code
func (o *SvmMigrationGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm migration get o k response has a 3xx status code
func (o *SvmMigrationGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm migration get o k response has a 4xx status code
func (o *SvmMigrationGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm migration get o k response has a 5xx status code
func (o *SvmMigrationGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this svm migration get o k response a status code equal to that given
func (o *SvmMigrationGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the svm migration get o k response
func (o *SvmMigrationGetOK) Code() int {
	return 200
}

func (o *SvmMigrationGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/migrations/{uuid}][%d] svmMigrationGetOK %s", 200, payload)
}

func (o *SvmMigrationGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/migrations/{uuid}][%d] svmMigrationGetOK %s", 200, payload)
}

func (o *SvmMigrationGetOK) GetPayload() *models.SvmMigration {
	return o.Payload
}

func (o *SvmMigrationGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SvmMigration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmMigrationGetDefault creates a SvmMigrationGetDefault with default headers values
func NewSvmMigrationGetDefault(code int) *SvmMigrationGetDefault {
	return &SvmMigrationGetDefault{
		_statusCode: code,
	}
}

/*
	SvmMigrationGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 13172783 | Migrate RDB lookup failed |
| 13173739 | Migrate pause operation failed. Retry pause operation using REST API PATCH method \\"/api/svm/migrations/<migration_uuid>?action=pause\\". Reason: {Reason} |
| 13173740 | Migrate abort operation failed. Retry abort operation by using REST API DELETE method \\"/api/svm/migrations/<migration_uuid>\\". Reason: {Reason} |
| 13173741 | Migrate failed. Retry the migrate by running the resume operation using REST API PATCH method \\"/api/svm/migrations/<migration_uuid>?action=resume\\". Reason: {Reason} |
| 13173742 | Migrate operation status: {Reason}. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SvmMigrationGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this svm migration get default response has a 2xx status code
func (o *SvmMigrationGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm migration get default response has a 3xx status code
func (o *SvmMigrationGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm migration get default response has a 4xx status code
func (o *SvmMigrationGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm migration get default response has a 5xx status code
func (o *SvmMigrationGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm migration get default response a status code equal to that given
func (o *SvmMigrationGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the svm migration get default response
func (o *SvmMigrationGetDefault) Code() int {
	return o._statusCode
}

func (o *SvmMigrationGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/migrations/{uuid}][%d] svm_migration_get default %s", o._statusCode, payload)
}

func (o *SvmMigrationGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/migrations/{uuid}][%d] svm_migration_get default %s", o._statusCode, payload)
}

func (o *SvmMigrationGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmMigrationGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
