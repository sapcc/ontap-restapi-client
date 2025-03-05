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

// FpolicyPersistentStoreCollectionGetReader is a Reader for the FpolicyPersistentStoreCollectionGet structure.
type FpolicyPersistentStoreCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyPersistentStoreCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyPersistentStoreCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyPersistentStoreCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyPersistentStoreCollectionGetOK creates a FpolicyPersistentStoreCollectionGetOK with default headers values
func NewFpolicyPersistentStoreCollectionGetOK() *FpolicyPersistentStoreCollectionGetOK {
	return &FpolicyPersistentStoreCollectionGetOK{}
}

/*
FpolicyPersistentStoreCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyPersistentStoreCollectionGetOK struct {
	Payload *models.FpolicyPersistentStoreResponse
}

// IsSuccess returns true when this fpolicy persistent store collection get o k response has a 2xx status code
func (o *FpolicyPersistentStoreCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fpolicy persistent store collection get o k response has a 3xx status code
func (o *FpolicyPersistentStoreCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fpolicy persistent store collection get o k response has a 4xx status code
func (o *FpolicyPersistentStoreCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fpolicy persistent store collection get o k response has a 5xx status code
func (o *FpolicyPersistentStoreCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fpolicy persistent store collection get o k response a status code equal to that given
func (o *FpolicyPersistentStoreCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fpolicy persistent store collection get o k response
func (o *FpolicyPersistentStoreCollectionGetOK) Code() int {
	return 200
}

func (o *FpolicyPersistentStoreCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/persistent-stores][%d] fpolicyPersistentStoreCollectionGetOK %s", 200, payload)
}

func (o *FpolicyPersistentStoreCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/persistent-stores][%d] fpolicyPersistentStoreCollectionGetOK %s", 200, payload)
}

func (o *FpolicyPersistentStoreCollectionGetOK) GetPayload() *models.FpolicyPersistentStoreResponse {
	return o.Payload
}

func (o *FpolicyPersistentStoreCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FpolicyPersistentStoreResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFpolicyPersistentStoreCollectionGetDefault creates a FpolicyPersistentStoreCollectionGetDefault with default headers values
func NewFpolicyPersistentStoreCollectionGetDefault(code int) *FpolicyPersistentStoreCollectionGetDefault {
	return &FpolicyPersistentStoreCollectionGetDefault{
		_statusCode: code,
	}
}

/*
FpolicyPersistentStoreCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type FpolicyPersistentStoreCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fpolicy persistent store collection get default response has a 2xx status code
func (o *FpolicyPersistentStoreCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fpolicy persistent store collection get default response has a 3xx status code
func (o *FpolicyPersistentStoreCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fpolicy persistent store collection get default response has a 4xx status code
func (o *FpolicyPersistentStoreCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fpolicy persistent store collection get default response has a 5xx status code
func (o *FpolicyPersistentStoreCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fpolicy persistent store collection get default response a status code equal to that given
func (o *FpolicyPersistentStoreCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fpolicy persistent store collection get default response
func (o *FpolicyPersistentStoreCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyPersistentStoreCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/persistent-stores][%d] fpolicy_persistent_store_collection_get default %s", o._statusCode, payload)
}

func (o *FpolicyPersistentStoreCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/persistent-stores][%d] fpolicy_persistent_store_collection_get default %s", o._statusCode, payload)
}

func (o *FpolicyPersistentStoreCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyPersistentStoreCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
