// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// NameMappingDeleteReader is a Reader for the NameMappingDelete structure.
type NameMappingDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NameMappingDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNameMappingDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNameMappingDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNameMappingDeleteOK creates a NameMappingDeleteOK with default headers values
func NewNameMappingDeleteOK() *NameMappingDeleteOK {
	return &NameMappingDeleteOK{}
}

/*
NameMappingDeleteOK describes a response with status code 200, with default header values.

OK
*/
type NameMappingDeleteOK struct {
}

// IsSuccess returns true when this name mapping delete o k response has a 2xx status code
func (o *NameMappingDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this name mapping delete o k response has a 3xx status code
func (o *NameMappingDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this name mapping delete o k response has a 4xx status code
func (o *NameMappingDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this name mapping delete o k response has a 5xx status code
func (o *NameMappingDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this name mapping delete o k response a status code equal to that given
func (o *NameMappingDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the name mapping delete o k response
func (o *NameMappingDeleteOK) Code() int {
	return 200
}

func (o *NameMappingDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /name-services/name-mappings/{svm.uuid}/{direction}/{index}][%d] nameMappingDeleteOK", 200)
}

func (o *NameMappingDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /name-services/name-mappings/{svm.uuid}/{direction}/{index}][%d] nameMappingDeleteOK", 200)
}

func (o *NameMappingDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNameMappingDeleteDefault creates a NameMappingDeleteDefault with default headers values
func NewNameMappingDeleteDefault(code int) *NameMappingDeleteDefault {
	return &NameMappingDeleteDefault{
		_statusCode: code,
	}
}

/*
NameMappingDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type NameMappingDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this name mapping delete default response has a 2xx status code
func (o *NameMappingDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this name mapping delete default response has a 3xx status code
func (o *NameMappingDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this name mapping delete default response has a 4xx status code
func (o *NameMappingDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this name mapping delete default response has a 5xx status code
func (o *NameMappingDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this name mapping delete default response a status code equal to that given
func (o *NameMappingDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the name mapping delete default response
func (o *NameMappingDeleteDefault) Code() int {
	return o._statusCode
}

func (o *NameMappingDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /name-services/name-mappings/{svm.uuid}/{direction}/{index}][%d] name_mapping_delete default %s", o._statusCode, payload)
}

func (o *NameMappingDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /name-services/name-mappings/{svm.uuid}/{direction}/{index}][%d] name_mapping_delete default %s", o._statusCode, payload)
}

func (o *NameMappingDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NameMappingDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
