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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// LocalCifsUsersAndGroupsImportGetReader is a Reader for the LocalCifsUsersAndGroupsImportGet structure.
type LocalCifsUsersAndGroupsImportGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalCifsUsersAndGroupsImportGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocalCifsUsersAndGroupsImportGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocalCifsUsersAndGroupsImportGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocalCifsUsersAndGroupsImportGetOK creates a LocalCifsUsersAndGroupsImportGetOK with default headers values
func NewLocalCifsUsersAndGroupsImportGetOK() *LocalCifsUsersAndGroupsImportGetOK {
	return &LocalCifsUsersAndGroupsImportGetOK{}
}

/*
LocalCifsUsersAndGroupsImportGetOK describes a response with status code 200, with default header values.

OK
*/
type LocalCifsUsersAndGroupsImportGetOK struct {
	Payload *models.LocalCifsUsersAndGroupsImport
}

// IsSuccess returns true when this local cifs users and groups import get o k response has a 2xx status code
func (o *LocalCifsUsersAndGroupsImportGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this local cifs users and groups import get o k response has a 3xx status code
func (o *LocalCifsUsersAndGroupsImportGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this local cifs users and groups import get o k response has a 4xx status code
func (o *LocalCifsUsersAndGroupsImportGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this local cifs users and groups import get o k response has a 5xx status code
func (o *LocalCifsUsersAndGroupsImportGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this local cifs users and groups import get o k response a status code equal to that given
func (o *LocalCifsUsersAndGroupsImportGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the local cifs users and groups import get o k response
func (o *LocalCifsUsersAndGroupsImportGetOK) Code() int {
	return 200
}

func (o *LocalCifsUsersAndGroupsImportGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/users-and-groups/bulk-import/{svm.uuid}][%d] localCifsUsersAndGroupsImportGetOK %s", 200, payload)
}

func (o *LocalCifsUsersAndGroupsImportGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/users-and-groups/bulk-import/{svm.uuid}][%d] localCifsUsersAndGroupsImportGetOK %s", 200, payload)
}

func (o *LocalCifsUsersAndGroupsImportGetOK) GetPayload() *models.LocalCifsUsersAndGroupsImport {
	return o.Payload
}

func (o *LocalCifsUsersAndGroupsImportGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocalCifsUsersAndGroupsImport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocalCifsUsersAndGroupsImportGetDefault creates a LocalCifsUsersAndGroupsImportGetDefault with default headers values
func NewLocalCifsUsersAndGroupsImportGetDefault(code int) *LocalCifsUsersAndGroupsImportGetDefault {
	return &LocalCifsUsersAndGroupsImportGetDefault{
		_statusCode: code,
	}
}

/*
LocalCifsUsersAndGroupsImportGetDefault describes a response with status code -1, with default header values.

Error
*/
type LocalCifsUsersAndGroupsImportGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this local cifs users and groups import get default response has a 2xx status code
func (o *LocalCifsUsersAndGroupsImportGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this local cifs users and groups import get default response has a 3xx status code
func (o *LocalCifsUsersAndGroupsImportGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this local cifs users and groups import get default response has a 4xx status code
func (o *LocalCifsUsersAndGroupsImportGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this local cifs users and groups import get default response has a 5xx status code
func (o *LocalCifsUsersAndGroupsImportGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this local cifs users and groups import get default response a status code equal to that given
func (o *LocalCifsUsersAndGroupsImportGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the local cifs users and groups import get default response
func (o *LocalCifsUsersAndGroupsImportGetDefault) Code() int {
	return o._statusCode
}

func (o *LocalCifsUsersAndGroupsImportGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/users-and-groups/bulk-import/{svm.uuid}][%d] local_cifs_users_and_groups_import_get default %s", o._statusCode, payload)
}

func (o *LocalCifsUsersAndGroupsImportGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/users-and-groups/bulk-import/{svm.uuid}][%d] local_cifs_users_and_groups_import_get default %s", o._statusCode, payload)
}

func (o *LocalCifsUsersAndGroupsImportGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LocalCifsUsersAndGroupsImportGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
