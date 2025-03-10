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

// ActiveDirectoryPreferredDcCollectionGetReader is a Reader for the ActiveDirectoryPreferredDcCollectionGet structure.
type ActiveDirectoryPreferredDcCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActiveDirectoryPreferredDcCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActiveDirectoryPreferredDcCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewActiveDirectoryPreferredDcCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActiveDirectoryPreferredDcCollectionGetOK creates a ActiveDirectoryPreferredDcCollectionGetOK with default headers values
func NewActiveDirectoryPreferredDcCollectionGetOK() *ActiveDirectoryPreferredDcCollectionGetOK {
	return &ActiveDirectoryPreferredDcCollectionGetOK{}
}

/*
ActiveDirectoryPreferredDcCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ActiveDirectoryPreferredDcCollectionGetOK struct {
	Payload *models.ActiveDirectoryPreferredDcResponse
}

// IsSuccess returns true when this active directory preferred dc collection get o k response has a 2xx status code
func (o *ActiveDirectoryPreferredDcCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this active directory preferred dc collection get o k response has a 3xx status code
func (o *ActiveDirectoryPreferredDcCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this active directory preferred dc collection get o k response has a 4xx status code
func (o *ActiveDirectoryPreferredDcCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this active directory preferred dc collection get o k response has a 5xx status code
func (o *ActiveDirectoryPreferredDcCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this active directory preferred dc collection get o k response a status code equal to that given
func (o *ActiveDirectoryPreferredDcCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the active directory preferred dc collection get o k response
func (o *ActiveDirectoryPreferredDcCollectionGetOK) Code() int {
	return 200
}

func (o *ActiveDirectoryPreferredDcCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/active-directory/{svm.uuid}/preferred-domain-controllers][%d] activeDirectoryPreferredDcCollectionGetOK %s", 200, payload)
}

func (o *ActiveDirectoryPreferredDcCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/active-directory/{svm.uuid}/preferred-domain-controllers][%d] activeDirectoryPreferredDcCollectionGetOK %s", 200, payload)
}

func (o *ActiveDirectoryPreferredDcCollectionGetOK) GetPayload() *models.ActiveDirectoryPreferredDcResponse {
	return o.Payload
}

func (o *ActiveDirectoryPreferredDcCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActiveDirectoryPreferredDcResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActiveDirectoryPreferredDcCollectionGetDefault creates a ActiveDirectoryPreferredDcCollectionGetDefault with default headers values
func NewActiveDirectoryPreferredDcCollectionGetDefault(code int) *ActiveDirectoryPreferredDcCollectionGetDefault {
	return &ActiveDirectoryPreferredDcCollectionGetDefault{
		_statusCode: code,
	}
}

/*
ActiveDirectoryPreferredDcCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type ActiveDirectoryPreferredDcCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this active directory preferred dc collection get default response has a 2xx status code
func (o *ActiveDirectoryPreferredDcCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this active directory preferred dc collection get default response has a 3xx status code
func (o *ActiveDirectoryPreferredDcCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this active directory preferred dc collection get default response has a 4xx status code
func (o *ActiveDirectoryPreferredDcCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this active directory preferred dc collection get default response has a 5xx status code
func (o *ActiveDirectoryPreferredDcCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this active directory preferred dc collection get default response a status code equal to that given
func (o *ActiveDirectoryPreferredDcCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the active directory preferred dc collection get default response
func (o *ActiveDirectoryPreferredDcCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *ActiveDirectoryPreferredDcCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/active-directory/{svm.uuid}/preferred-domain-controllers][%d] active_directory_preferred_dc_collection_get default %s", o._statusCode, payload)
}

func (o *ActiveDirectoryPreferredDcCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/active-directory/{svm.uuid}/preferred-domain-controllers][%d] active_directory_preferred_dc_collection_get default %s", o._statusCode, payload)
}

func (o *ActiveDirectoryPreferredDcCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ActiveDirectoryPreferredDcCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
