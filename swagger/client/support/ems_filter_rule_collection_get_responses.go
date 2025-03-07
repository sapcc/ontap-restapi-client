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

// EmsFilterRuleCollectionGetReader is a Reader for the EmsFilterRuleCollectionGet structure.
type EmsFilterRuleCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsFilterRuleCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmsFilterRuleCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsFilterRuleCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsFilterRuleCollectionGetOK creates a EmsFilterRuleCollectionGetOK with default headers values
func NewEmsFilterRuleCollectionGetOK() *EmsFilterRuleCollectionGetOK {
	return &EmsFilterRuleCollectionGetOK{}
}

/*
EmsFilterRuleCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type EmsFilterRuleCollectionGetOK struct {
	Payload *models.EmsFilterRuleResponse
}

// IsSuccess returns true when this ems filter rule collection get o k response has a 2xx status code
func (o *EmsFilterRuleCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ems filter rule collection get o k response has a 3xx status code
func (o *EmsFilterRuleCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ems filter rule collection get o k response has a 4xx status code
func (o *EmsFilterRuleCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ems filter rule collection get o k response has a 5xx status code
func (o *EmsFilterRuleCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ems filter rule collection get o k response a status code equal to that given
func (o *EmsFilterRuleCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ems filter rule collection get o k response
func (o *EmsFilterRuleCollectionGetOK) Code() int {
	return 200
}

func (o *EmsFilterRuleCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/filters/{name}/rules][%d] emsFilterRuleCollectionGetOK %s", 200, payload)
}

func (o *EmsFilterRuleCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/filters/{name}/rules][%d] emsFilterRuleCollectionGetOK %s", 200, payload)
}

func (o *EmsFilterRuleCollectionGetOK) GetPayload() *models.EmsFilterRuleResponse {
	return o.Payload
}

func (o *EmsFilterRuleCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmsFilterRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmsFilterRuleCollectionGetDefault creates a EmsFilterRuleCollectionGetDefault with default headers values
func NewEmsFilterRuleCollectionGetDefault(code int) *EmsFilterRuleCollectionGetDefault {
	return &EmsFilterRuleCollectionGetDefault{
		_statusCode: code,
	}
}

/*
EmsFilterRuleCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type EmsFilterRuleCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ems filter rule collection get default response has a 2xx status code
func (o *EmsFilterRuleCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ems filter rule collection get default response has a 3xx status code
func (o *EmsFilterRuleCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ems filter rule collection get default response has a 4xx status code
func (o *EmsFilterRuleCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ems filter rule collection get default response has a 5xx status code
func (o *EmsFilterRuleCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ems filter rule collection get default response a status code equal to that given
func (o *EmsFilterRuleCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ems filter rule collection get default response
func (o *EmsFilterRuleCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *EmsFilterRuleCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/filters/{name}/rules][%d] ems_filter_rule_collection_get default %s", o._statusCode, payload)
}

func (o *EmsFilterRuleCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/filters/{name}/rules][%d] ems_filter_rule_collection_get default %s", o._statusCode, payload)
}

func (o *EmsFilterRuleCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsFilterRuleCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
