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

// FpolicyPersistentStoreDeleteReader is a Reader for the FpolicyPersistentStoreDelete structure.
type FpolicyPersistentStoreDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyPersistentStoreDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyPersistentStoreDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyPersistentStoreDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyPersistentStoreDeleteOK creates a FpolicyPersistentStoreDeleteOK with default headers values
func NewFpolicyPersistentStoreDeleteOK() *FpolicyPersistentStoreDeleteOK {
	return &FpolicyPersistentStoreDeleteOK{}
}

/*
FpolicyPersistentStoreDeleteOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyPersistentStoreDeleteOK struct {
}

// IsSuccess returns true when this fpolicy persistent store delete o k response has a 2xx status code
func (o *FpolicyPersistentStoreDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fpolicy persistent store delete o k response has a 3xx status code
func (o *FpolicyPersistentStoreDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fpolicy persistent store delete o k response has a 4xx status code
func (o *FpolicyPersistentStoreDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fpolicy persistent store delete o k response has a 5xx status code
func (o *FpolicyPersistentStoreDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fpolicy persistent store delete o k response a status code equal to that given
func (o *FpolicyPersistentStoreDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fpolicy persistent store delete o k response
func (o *FpolicyPersistentStoreDeleteOK) Code() int {
	return 200
}

func (o *FpolicyPersistentStoreDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/fpolicy/{svm.uuid}/persistent-stores/{name}][%d] fpolicyPersistentStoreDeleteOK", 200)
}

func (o *FpolicyPersistentStoreDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/fpolicy/{svm.uuid}/persistent-stores/{name}][%d] fpolicyPersistentStoreDeleteOK", 200)
}

func (o *FpolicyPersistentStoreDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFpolicyPersistentStoreDeleteDefault creates a FpolicyPersistentStoreDeleteDefault with default headers values
func NewFpolicyPersistentStoreDeleteDefault(code int) *FpolicyPersistentStoreDeleteDefault {
	return &FpolicyPersistentStoreDeleteDefault{
		_statusCode: code,
	}
}

/*
	FpolicyPersistentStoreDeleteDefault describes a response with status code -1, with default header values.

	| Error Code | Description |

| ---------- | ----------- |
| 9765064    | Persistent Store is active on at least one policy |
*/
type FpolicyPersistentStoreDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fpolicy persistent store delete default response has a 2xx status code
func (o *FpolicyPersistentStoreDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fpolicy persistent store delete default response has a 3xx status code
func (o *FpolicyPersistentStoreDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fpolicy persistent store delete default response has a 4xx status code
func (o *FpolicyPersistentStoreDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fpolicy persistent store delete default response has a 5xx status code
func (o *FpolicyPersistentStoreDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fpolicy persistent store delete default response a status code equal to that given
func (o *FpolicyPersistentStoreDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fpolicy persistent store delete default response
func (o *FpolicyPersistentStoreDeleteDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyPersistentStoreDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/fpolicy/{svm.uuid}/persistent-stores/{name}][%d] fpolicy_persistent_store_delete default %s", o._statusCode, payload)
}

func (o *FpolicyPersistentStoreDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/fpolicy/{svm.uuid}/persistent-stores/{name}][%d] fpolicy_persistent_store_delete default %s", o._statusCode, payload)
}

func (o *FpolicyPersistentStoreDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyPersistentStoreDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
