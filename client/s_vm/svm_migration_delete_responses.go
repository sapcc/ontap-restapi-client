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

	"github.com/sapcc/ontap-restapi/models"
)

// SvmMigrationDeleteReader is a Reader for the SvmMigrationDelete structure.
type SvmMigrationDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmMigrationDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSvmMigrationDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSvmMigrationDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmMigrationDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmMigrationDeleteOK creates a SvmMigrationDeleteOK with default headers values
func NewSvmMigrationDeleteOK() *SvmMigrationDeleteOK {
	return &SvmMigrationDeleteOK{}
}

/*
SvmMigrationDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SvmMigrationDeleteOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this svm migration delete o k response has a 2xx status code
func (o *SvmMigrationDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm migration delete o k response has a 3xx status code
func (o *SvmMigrationDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm migration delete o k response has a 4xx status code
func (o *SvmMigrationDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm migration delete o k response has a 5xx status code
func (o *SvmMigrationDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this svm migration delete o k response a status code equal to that given
func (o *SvmMigrationDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the svm migration delete o k response
func (o *SvmMigrationDeleteOK) Code() int {
	return 200
}

func (o *SvmMigrationDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/migrations/{uuid}][%d] svmMigrationDeleteOK %s", 200, payload)
}

func (o *SvmMigrationDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/migrations/{uuid}][%d] svmMigrationDeleteOK %s", 200, payload)
}

func (o *SvmMigrationDeleteOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SvmMigrationDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmMigrationDeleteAccepted creates a SvmMigrationDeleteAccepted with default headers values
func NewSvmMigrationDeleteAccepted() *SvmMigrationDeleteAccepted {
	return &SvmMigrationDeleteAccepted{}
}

/*
SvmMigrationDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SvmMigrationDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this svm migration delete accepted response has a 2xx status code
func (o *SvmMigrationDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm migration delete accepted response has a 3xx status code
func (o *SvmMigrationDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm migration delete accepted response has a 4xx status code
func (o *SvmMigrationDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm migration delete accepted response has a 5xx status code
func (o *SvmMigrationDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this svm migration delete accepted response a status code equal to that given
func (o *SvmMigrationDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the svm migration delete accepted response
func (o *SvmMigrationDeleteAccepted) Code() int {
	return 202
}

func (o *SvmMigrationDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/migrations/{uuid}][%d] svmMigrationDeleteAccepted %s", 202, payload)
}

func (o *SvmMigrationDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/migrations/{uuid}][%d] svmMigrationDeleteAccepted %s", 202, payload)
}

func (o *SvmMigrationDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SvmMigrationDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmMigrationDeleteDefault creates a SvmMigrationDeleteDefault with default headers values
func NewSvmMigrationDeleteDefault(code int) *SvmMigrationDeleteDefault {
	return &SvmMigrationDeleteDefault{
		_statusCode: code,
	}
}

/*
	SvmMigrationDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 13172783 | Migrate RDB lookup failed |
| 13173738 | REST API DELETE method \\"/api/svm/migrations\\" is only supported on the destination cluster. Issue the REST API DELETE request to the destination cluster. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SvmMigrationDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this svm migration delete default response has a 2xx status code
func (o *SvmMigrationDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm migration delete default response has a 3xx status code
func (o *SvmMigrationDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm migration delete default response has a 4xx status code
func (o *SvmMigrationDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm migration delete default response has a 5xx status code
func (o *SvmMigrationDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm migration delete default response a status code equal to that given
func (o *SvmMigrationDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the svm migration delete default response
func (o *SvmMigrationDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SvmMigrationDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/migrations/{uuid}][%d] svm_migration_delete default %s", o._statusCode, payload)
}

func (o *SvmMigrationDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/migrations/{uuid}][%d] svm_migration_delete default %s", o._statusCode, payload)
}

func (o *SvmMigrationDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmMigrationDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
