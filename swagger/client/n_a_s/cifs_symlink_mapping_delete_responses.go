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

// CifsSymlinkMappingDeleteReader is a Reader for the CifsSymlinkMappingDelete structure.
type CifsSymlinkMappingDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsSymlinkMappingDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsSymlinkMappingDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsSymlinkMappingDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsSymlinkMappingDeleteOK creates a CifsSymlinkMappingDeleteOK with default headers values
func NewCifsSymlinkMappingDeleteOK() *CifsSymlinkMappingDeleteOK {
	return &CifsSymlinkMappingDeleteOK{}
}

/*
CifsSymlinkMappingDeleteOK describes a response with status code 200, with default header values.

OK
*/
type CifsSymlinkMappingDeleteOK struct {
}

// IsSuccess returns true when this cifs symlink mapping delete o k response has a 2xx status code
func (o *CifsSymlinkMappingDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs symlink mapping delete o k response has a 3xx status code
func (o *CifsSymlinkMappingDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs symlink mapping delete o k response has a 4xx status code
func (o *CifsSymlinkMappingDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs symlink mapping delete o k response has a 5xx status code
func (o *CifsSymlinkMappingDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs symlink mapping delete o k response a status code equal to that given
func (o *CifsSymlinkMappingDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs symlink mapping delete o k response
func (o *CifsSymlinkMappingDeleteOK) Code() int {
	return 200
}

func (o *CifsSymlinkMappingDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/unix-symlink-mapping/{svm.uuid}/{unix_path}][%d] cifsSymlinkMappingDeleteOK", 200)
}

func (o *CifsSymlinkMappingDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/unix-symlink-mapping/{svm.uuid}/{unix_path}][%d] cifsSymlinkMappingDeleteOK", 200)
}

func (o *CifsSymlinkMappingDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCifsSymlinkMappingDeleteDefault creates a CifsSymlinkMappingDeleteDefault with default headers values
func NewCifsSymlinkMappingDeleteDefault(code int) *CifsSymlinkMappingDeleteDefault {
	return &CifsSymlinkMappingDeleteDefault{
		_statusCode: code,
	}
}

/*
CifsSymlinkMappingDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type CifsSymlinkMappingDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs symlink mapping delete default response has a 2xx status code
func (o *CifsSymlinkMappingDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs symlink mapping delete default response has a 3xx status code
func (o *CifsSymlinkMappingDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs symlink mapping delete default response has a 4xx status code
func (o *CifsSymlinkMappingDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs symlink mapping delete default response has a 5xx status code
func (o *CifsSymlinkMappingDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs symlink mapping delete default response a status code equal to that given
func (o *CifsSymlinkMappingDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs symlink mapping delete default response
func (o *CifsSymlinkMappingDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CifsSymlinkMappingDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/unix-symlink-mapping/{svm.uuid}/{unix_path}][%d] cifs_symlink_mapping_delete default %s", o._statusCode, payload)
}

func (o *CifsSymlinkMappingDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/unix-symlink-mapping/{svm.uuid}/{unix_path}][%d] cifs_symlink_mapping_delete default %s", o._statusCode, payload)
}

func (o *CifsSymlinkMappingDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsSymlinkMappingDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
