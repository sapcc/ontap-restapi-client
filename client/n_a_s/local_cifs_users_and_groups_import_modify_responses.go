// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// LocalCifsUsersAndGroupsImportModifyReader is a Reader for the LocalCifsUsersAndGroupsImportModify structure.
type LocalCifsUsersAndGroupsImportModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalCifsUsersAndGroupsImportModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocalCifsUsersAndGroupsImportModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewLocalCifsUsersAndGroupsImportModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocalCifsUsersAndGroupsImportModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocalCifsUsersAndGroupsImportModifyOK creates a LocalCifsUsersAndGroupsImportModifyOK with default headers values
func NewLocalCifsUsersAndGroupsImportModifyOK() *LocalCifsUsersAndGroupsImportModifyOK {
	return &LocalCifsUsersAndGroupsImportModifyOK{}
}

/*
LocalCifsUsersAndGroupsImportModifyOK describes a response with status code 200, with default header values.

OK
*/
type LocalCifsUsersAndGroupsImportModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this local cifs users and groups import modify o k response has a 2xx status code
func (o *LocalCifsUsersAndGroupsImportModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this local cifs users and groups import modify o k response has a 3xx status code
func (o *LocalCifsUsersAndGroupsImportModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this local cifs users and groups import modify o k response has a 4xx status code
func (o *LocalCifsUsersAndGroupsImportModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this local cifs users and groups import modify o k response has a 5xx status code
func (o *LocalCifsUsersAndGroupsImportModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this local cifs users and groups import modify o k response a status code equal to that given
func (o *LocalCifsUsersAndGroupsImportModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the local cifs users and groups import modify o k response
func (o *LocalCifsUsersAndGroupsImportModifyOK) Code() int {
	return 200
}

func (o *LocalCifsUsersAndGroupsImportModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/users-and-groups/bulk-import/{svm.uuid}][%d] localCifsUsersAndGroupsImportModifyOK %s", 200, payload)
}

func (o *LocalCifsUsersAndGroupsImportModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/users-and-groups/bulk-import/{svm.uuid}][%d] localCifsUsersAndGroupsImportModifyOK %s", 200, payload)
}

func (o *LocalCifsUsersAndGroupsImportModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *LocalCifsUsersAndGroupsImportModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocalCifsUsersAndGroupsImportModifyAccepted creates a LocalCifsUsersAndGroupsImportModifyAccepted with default headers values
func NewLocalCifsUsersAndGroupsImportModifyAccepted() *LocalCifsUsersAndGroupsImportModifyAccepted {
	return &LocalCifsUsersAndGroupsImportModifyAccepted{}
}

/*
LocalCifsUsersAndGroupsImportModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type LocalCifsUsersAndGroupsImportModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this local cifs users and groups import modify accepted response has a 2xx status code
func (o *LocalCifsUsersAndGroupsImportModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this local cifs users and groups import modify accepted response has a 3xx status code
func (o *LocalCifsUsersAndGroupsImportModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this local cifs users and groups import modify accepted response has a 4xx status code
func (o *LocalCifsUsersAndGroupsImportModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this local cifs users and groups import modify accepted response has a 5xx status code
func (o *LocalCifsUsersAndGroupsImportModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this local cifs users and groups import modify accepted response a status code equal to that given
func (o *LocalCifsUsersAndGroupsImportModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the local cifs users and groups import modify accepted response
func (o *LocalCifsUsersAndGroupsImportModifyAccepted) Code() int {
	return 202
}

func (o *LocalCifsUsersAndGroupsImportModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/users-and-groups/bulk-import/{svm.uuid}][%d] localCifsUsersAndGroupsImportModifyAccepted %s", 202, payload)
}

func (o *LocalCifsUsersAndGroupsImportModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/users-and-groups/bulk-import/{svm.uuid}][%d] localCifsUsersAndGroupsImportModifyAccepted %s", 202, payload)
}

func (o *LocalCifsUsersAndGroupsImportModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *LocalCifsUsersAndGroupsImportModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocalCifsUsersAndGroupsImportModifyDefault creates a LocalCifsUsersAndGroupsImportModifyDefault with default headers values
func NewLocalCifsUsersAndGroupsImportModifyDefault(code int) *LocalCifsUsersAndGroupsImportModifyDefault {
	return &LocalCifsUsersAndGroupsImportModifyDefault{
		_statusCode: code,
	}
}

/*
LocalCifsUsersAndGroupsImportModifyDefault describes a response with status code -1, with default header values.

Error
*/
type LocalCifsUsersAndGroupsImportModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this local cifs users and groups import modify default response has a 2xx status code
func (o *LocalCifsUsersAndGroupsImportModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this local cifs users and groups import modify default response has a 3xx status code
func (o *LocalCifsUsersAndGroupsImportModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this local cifs users and groups import modify default response has a 4xx status code
func (o *LocalCifsUsersAndGroupsImportModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this local cifs users and groups import modify default response has a 5xx status code
func (o *LocalCifsUsersAndGroupsImportModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this local cifs users and groups import modify default response a status code equal to that given
func (o *LocalCifsUsersAndGroupsImportModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the local cifs users and groups import modify default response
func (o *LocalCifsUsersAndGroupsImportModifyDefault) Code() int {
	return o._statusCode
}

func (o *LocalCifsUsersAndGroupsImportModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/users-and-groups/bulk-import/{svm.uuid}][%d] local_cifs_users_and_groups_import_modify default %s", o._statusCode, payload)
}

func (o *LocalCifsUsersAndGroupsImportModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/users-and-groups/bulk-import/{svm.uuid}][%d] local_cifs_users_and_groups_import_modify default %s", o._statusCode, payload)
}

func (o *LocalCifsUsersAndGroupsImportModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LocalCifsUsersAndGroupsImportModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
