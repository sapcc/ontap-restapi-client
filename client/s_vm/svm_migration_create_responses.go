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

	"github.com/sapcc/ontap-restapi-client/models"
)

// SvmMigrationCreateReader is a Reader for the SvmMigrationCreate structure.
type SvmMigrationCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmMigrationCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSvmMigrationCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSvmMigrationCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmMigrationCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmMigrationCreateCreated creates a SvmMigrationCreateCreated with default headers values
func NewSvmMigrationCreateCreated() *SvmMigrationCreateCreated {
	return &SvmMigrationCreateCreated{}
}

/*
SvmMigrationCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SvmMigrationCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.SvmMigrationCreate
}

// IsSuccess returns true when this svm migration create created response has a 2xx status code
func (o *SvmMigrationCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm migration create created response has a 3xx status code
func (o *SvmMigrationCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm migration create created response has a 4xx status code
func (o *SvmMigrationCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm migration create created response has a 5xx status code
func (o *SvmMigrationCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this svm migration create created response a status code equal to that given
func (o *SvmMigrationCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the svm migration create created response
func (o *SvmMigrationCreateCreated) Code() int {
	return 201
}

func (o *SvmMigrationCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/migrations][%d] svmMigrationCreateCreated %s", 201, payload)
}

func (o *SvmMigrationCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/migrations][%d] svmMigrationCreateCreated %s", 201, payload)
}

func (o *SvmMigrationCreateCreated) GetPayload() *models.SvmMigrationCreate {
	return o.Payload
}

func (o *SvmMigrationCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.SvmMigrationCreate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmMigrationCreateAccepted creates a SvmMigrationCreateAccepted with default headers values
func NewSvmMigrationCreateAccepted() *SvmMigrationCreateAccepted {
	return &SvmMigrationCreateAccepted{}
}

/*
SvmMigrationCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SvmMigrationCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.SvmMigrationCreate
}

// IsSuccess returns true when this svm migration create accepted response has a 2xx status code
func (o *SvmMigrationCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm migration create accepted response has a 3xx status code
func (o *SvmMigrationCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm migration create accepted response has a 4xx status code
func (o *SvmMigrationCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm migration create accepted response has a 5xx status code
func (o *SvmMigrationCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this svm migration create accepted response a status code equal to that given
func (o *SvmMigrationCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the svm migration create accepted response
func (o *SvmMigrationCreateAccepted) Code() int {
	return 202
}

func (o *SvmMigrationCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/migrations][%d] svmMigrationCreateAccepted %s", 202, payload)
}

func (o *SvmMigrationCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/migrations][%d] svmMigrationCreateAccepted %s", 202, payload)
}

func (o *SvmMigrationCreateAccepted) GetPayload() *models.SvmMigrationCreate {
	return o.Payload
}

func (o *SvmMigrationCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.SvmMigrationCreate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmMigrationCreateDefault creates a SvmMigrationCreateDefault with default headers values
func NewSvmMigrationCreateDefault(code int) *SvmMigrationCreateDefault {
	return &SvmMigrationCreateDefault{
		_statusCode: code,
	}
}

/*
	SvmMigrationCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262245 | The value provided is invalid. |
| 13172746 | SVM migration cannot be started. This is a generic code, see the response message for details. |
| 13173748 | Migrate request cannot contain both \\"aggregates\\" and \\"volume_aggregate_pairs\\" within the \\"volume_placement\\" object. |
| 13173758 | The property \\"{property}\\" is not supported for this operation. |
| 13173759 | Migrate operation failed. To use LIF placement specify either the port UUID, or both the port name and the port node name. |
| 13173760 | Migrate operation failed. LIF placement requires either the port node name or the port UUID to be specified if the port name is specified. |
| 13173761 | Migrate operation failed. LIF placement requires either the port name or the port UUID to be specified if the port node name is specified. |
| 13173762 | Migrate operation failed. To use LIF placement specify at least one of the following: IP interface UUID, IP interface name, or IP interface IP. |
| 13173764 | Migrate operation failed because LIF placement is not supported on the destination cluster. Both clusters must have an effective cluster version of 9.12.1 or later. |
| 13173765 | Migrate operation failed. Unable to determine if LIF placement is supported.  Reason: \\"{Reason}\\" |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SvmMigrationCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this svm migration create default response has a 2xx status code
func (o *SvmMigrationCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm migration create default response has a 3xx status code
func (o *SvmMigrationCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm migration create default response has a 4xx status code
func (o *SvmMigrationCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm migration create default response has a 5xx status code
func (o *SvmMigrationCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm migration create default response a status code equal to that given
func (o *SvmMigrationCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the svm migration create default response
func (o *SvmMigrationCreateDefault) Code() int {
	return o._statusCode
}

func (o *SvmMigrationCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/migrations][%d] svm_migration_create default %s", o._statusCode, payload)
}

func (o *SvmMigrationCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/migrations][%d] svm_migration_create default %s", o._statusCode, payload)
}

func (o *SvmMigrationCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmMigrationCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
