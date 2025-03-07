// Code generated by go-swagger; DO NOT EDIT.

package support

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

// EmsDestinationCollectionGetReader is a Reader for the EmsDestinationCollectionGet structure.
type EmsDestinationCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsDestinationCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmsDestinationCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsDestinationCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsDestinationCollectionGetOK creates a EmsDestinationCollectionGetOK with default headers values
func NewEmsDestinationCollectionGetOK() *EmsDestinationCollectionGetOK {
	return &EmsDestinationCollectionGetOK{}
}

/*
EmsDestinationCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type EmsDestinationCollectionGetOK struct {
	Payload *models.EmsDestinationResponse
}

// IsSuccess returns true when this ems destination collection get o k response has a 2xx status code
func (o *EmsDestinationCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ems destination collection get o k response has a 3xx status code
func (o *EmsDestinationCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ems destination collection get o k response has a 4xx status code
func (o *EmsDestinationCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ems destination collection get o k response has a 5xx status code
func (o *EmsDestinationCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ems destination collection get o k response a status code equal to that given
func (o *EmsDestinationCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ems destination collection get o k response
func (o *EmsDestinationCollectionGetOK) Code() int {
	return 200
}

func (o *EmsDestinationCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/destinations][%d] emsDestinationCollectionGetOK %s", 200, payload)
}

func (o *EmsDestinationCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/destinations][%d] emsDestinationCollectionGetOK %s", 200, payload)
}

func (o *EmsDestinationCollectionGetOK) GetPayload() *models.EmsDestinationResponse {
	return o.Payload
}

func (o *EmsDestinationCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmsDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmsDestinationCollectionGetDefault creates a EmsDestinationCollectionGetDefault with default headers values
func NewEmsDestinationCollectionGetDefault(code int) *EmsDestinationCollectionGetDefault {
	return &EmsDestinationCollectionGetDefault{
		_statusCode: code,
	}
}

/*
EmsDestinationCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type EmsDestinationCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ems destination collection get default response has a 2xx status code
func (o *EmsDestinationCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ems destination collection get default response has a 3xx status code
func (o *EmsDestinationCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ems destination collection get default response has a 4xx status code
func (o *EmsDestinationCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ems destination collection get default response has a 5xx status code
func (o *EmsDestinationCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ems destination collection get default response a status code equal to that given
func (o *EmsDestinationCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ems destination collection get default response
func (o *EmsDestinationCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *EmsDestinationCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/destinations][%d] ems_destination_collection_get default %s", o._statusCode, payload)
}

func (o *EmsDestinationCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/destinations][%d] ems_destination_collection_get default %s", o._statusCode, payload)
}

func (o *EmsDestinationCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsDestinationCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
