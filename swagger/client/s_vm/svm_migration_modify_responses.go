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

// SvmMigrationModifyReader is a Reader for the SvmMigrationModify structure.
type SvmMigrationModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmMigrationModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSvmMigrationModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSvmMigrationModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmMigrationModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmMigrationModifyOK creates a SvmMigrationModifyOK with default headers values
func NewSvmMigrationModifyOK() *SvmMigrationModifyOK {
	return &SvmMigrationModifyOK{}
}

/*
SvmMigrationModifyOK describes a response with status code 200, with default header values.

OK
*/
type SvmMigrationModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this svm migration modify o k response has a 2xx status code
func (o *SvmMigrationModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm migration modify o k response has a 3xx status code
func (o *SvmMigrationModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm migration modify o k response has a 4xx status code
func (o *SvmMigrationModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm migration modify o k response has a 5xx status code
func (o *SvmMigrationModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this svm migration modify o k response a status code equal to that given
func (o *SvmMigrationModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the svm migration modify o k response
func (o *SvmMigrationModifyOK) Code() int {
	return 200
}

func (o *SvmMigrationModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /svm/migrations/{uuid}][%d] svmMigrationModifyOK %s", 200, payload)
}

func (o *SvmMigrationModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /svm/migrations/{uuid}][%d] svmMigrationModifyOK %s", 200, payload)
}

func (o *SvmMigrationModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SvmMigrationModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmMigrationModifyAccepted creates a SvmMigrationModifyAccepted with default headers values
func NewSvmMigrationModifyAccepted() *SvmMigrationModifyAccepted {
	return &SvmMigrationModifyAccepted{}
}

/*
SvmMigrationModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SvmMigrationModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this svm migration modify accepted response has a 2xx status code
func (o *SvmMigrationModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm migration modify accepted response has a 3xx status code
func (o *SvmMigrationModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm migration modify accepted response has a 4xx status code
func (o *SvmMigrationModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm migration modify accepted response has a 5xx status code
func (o *SvmMigrationModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this svm migration modify accepted response a status code equal to that given
func (o *SvmMigrationModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the svm migration modify accepted response
func (o *SvmMigrationModifyAccepted) Code() int {
	return 202
}

func (o *SvmMigrationModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /svm/migrations/{uuid}][%d] svmMigrationModifyAccepted %s", 202, payload)
}

func (o *SvmMigrationModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /svm/migrations/{uuid}][%d] svmMigrationModifyAccepted %s", 202, payload)
}

func (o *SvmMigrationModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SvmMigrationModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmMigrationModifyDefault creates a SvmMigrationModifyDefault with default headers values
func NewSvmMigrationModifyDefault(code int) *SvmMigrationModifyDefault {
	return &SvmMigrationModifyDefault{
		_statusCode: code,
	}
}

/*
	SvmMigrationModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 13172783 | Migrate RDB lookup failed |
| 13173737 | REST API PATCH method \\"/api/svm/migrations\\" is only supported on the destination cluster. Issue the REST API PATCH request to the destination cluster. |
| 13173746 | Migrate resume operation failed. Cannot specify volume granular placement during resume if aggregate placement was specified during start operation. |
| 13173747 | Migrate operation failed. Volume placement can only be specified on PATCH with an action of \\"resume\\". |
| 13173748 | Migrate request cannot contain both \\"aggregates\\" and \\"volume_aggregate_pairs\\" within the \\"volume_placement\\" object. |
| 13173763 | Migrate operation failed. LIF placement is not supported in PATCH operations. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SvmMigrationModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this svm migration modify default response has a 2xx status code
func (o *SvmMigrationModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm migration modify default response has a 3xx status code
func (o *SvmMigrationModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm migration modify default response has a 4xx status code
func (o *SvmMigrationModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm migration modify default response has a 5xx status code
func (o *SvmMigrationModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm migration modify default response a status code equal to that given
func (o *SvmMigrationModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the svm migration modify default response
func (o *SvmMigrationModifyDefault) Code() int {
	return o._statusCode
}

func (o *SvmMigrationModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /svm/migrations/{uuid}][%d] svm_migration_modify default %s", o._statusCode, payload)
}

func (o *SvmMigrationModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /svm/migrations/{uuid}][%d] svm_migration_modify default %s", o._statusCode, payload)
}

func (o *SvmMigrationModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmMigrationModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
