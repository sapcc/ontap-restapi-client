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

// ShadowcopySetCollectionGetReader is a Reader for the ShadowcopySetCollectionGet structure.
type ShadowcopySetCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShadowcopySetCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShadowcopySetCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewShadowcopySetCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShadowcopySetCollectionGetOK creates a ShadowcopySetCollectionGetOK with default headers values
func NewShadowcopySetCollectionGetOK() *ShadowcopySetCollectionGetOK {
	return &ShadowcopySetCollectionGetOK{}
}

/*
ShadowcopySetCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ShadowcopySetCollectionGetOK struct {
	Payload *models.ShadowcopySetResponse
}

// IsSuccess returns true when this shadowcopy set collection get o k response has a 2xx status code
func (o *ShadowcopySetCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this shadowcopy set collection get o k response has a 3xx status code
func (o *ShadowcopySetCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this shadowcopy set collection get o k response has a 4xx status code
func (o *ShadowcopySetCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this shadowcopy set collection get o k response has a 5xx status code
func (o *ShadowcopySetCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this shadowcopy set collection get o k response a status code equal to that given
func (o *ShadowcopySetCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the shadowcopy set collection get o k response
func (o *ShadowcopySetCollectionGetOK) Code() int {
	return 200
}

func (o *ShadowcopySetCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/shadowcopy-sets][%d] shadowcopySetCollectionGetOK %s", 200, payload)
}

func (o *ShadowcopySetCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/shadowcopy-sets][%d] shadowcopySetCollectionGetOK %s", 200, payload)
}

func (o *ShadowcopySetCollectionGetOK) GetPayload() *models.ShadowcopySetResponse {
	return o.Payload
}

func (o *ShadowcopySetCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShadowcopySetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShadowcopySetCollectionGetDefault creates a ShadowcopySetCollectionGetDefault with default headers values
func NewShadowcopySetCollectionGetDefault(code int) *ShadowcopySetCollectionGetDefault {
	return &ShadowcopySetCollectionGetDefault{
		_statusCode: code,
	}
}

/*
ShadowcopySetCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type ShadowcopySetCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this shadowcopy set collection get default response has a 2xx status code
func (o *ShadowcopySetCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this shadowcopy set collection get default response has a 3xx status code
func (o *ShadowcopySetCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this shadowcopy set collection get default response has a 4xx status code
func (o *ShadowcopySetCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this shadowcopy set collection get default response has a 5xx status code
func (o *ShadowcopySetCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this shadowcopy set collection get default response a status code equal to that given
func (o *ShadowcopySetCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the shadowcopy set collection get default response
func (o *ShadowcopySetCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *ShadowcopySetCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/shadowcopy-sets][%d] shadowcopy_set_collection_get default %s", o._statusCode, payload)
}

func (o *ShadowcopySetCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/shadowcopy-sets][%d] shadowcopy_set_collection_get default %s", o._statusCode, payload)
}

func (o *ShadowcopySetCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ShadowcopySetCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
