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

// PoliciesAndRulesToBeAppliedCollectionGetReader is a Reader for the PoliciesAndRulesToBeAppliedCollectionGet structure.
type PoliciesAndRulesToBeAppliedCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PoliciesAndRulesToBeAppliedCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPoliciesAndRulesToBeAppliedCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPoliciesAndRulesToBeAppliedCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPoliciesAndRulesToBeAppliedCollectionGetOK creates a PoliciesAndRulesToBeAppliedCollectionGetOK with default headers values
func NewPoliciesAndRulesToBeAppliedCollectionGetOK() *PoliciesAndRulesToBeAppliedCollectionGetOK {
	return &PoliciesAndRulesToBeAppliedCollectionGetOK{}
}

/*
PoliciesAndRulesToBeAppliedCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type PoliciesAndRulesToBeAppliedCollectionGetOK struct {
	Payload *models.PoliciesAndRulesToBeAppliedResponse
}

// IsSuccess returns true when this policies and rules to be applied collection get o k response has a 2xx status code
func (o *PoliciesAndRulesToBeAppliedCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this policies and rules to be applied collection get o k response has a 3xx status code
func (o *PoliciesAndRulesToBeAppliedCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this policies and rules to be applied collection get o k response has a 4xx status code
func (o *PoliciesAndRulesToBeAppliedCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this policies and rules to be applied collection get o k response has a 5xx status code
func (o *PoliciesAndRulesToBeAppliedCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this policies and rules to be applied collection get o k response a status code equal to that given
func (o *PoliciesAndRulesToBeAppliedCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the policies and rules to be applied collection get o k response
func (o *PoliciesAndRulesToBeAppliedCollectionGetOK) Code() int {
	return 200
}

func (o *PoliciesAndRulesToBeAppliedCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies][%d] policiesAndRulesToBeAppliedCollectionGetOK %s", 200, payload)
}

func (o *PoliciesAndRulesToBeAppliedCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies][%d] policiesAndRulesToBeAppliedCollectionGetOK %s", 200, payload)
}

func (o *PoliciesAndRulesToBeAppliedCollectionGetOK) GetPayload() *models.PoliciesAndRulesToBeAppliedResponse {
	return o.Payload
}

func (o *PoliciesAndRulesToBeAppliedCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PoliciesAndRulesToBeAppliedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPoliciesAndRulesToBeAppliedCollectionGetDefault creates a PoliciesAndRulesToBeAppliedCollectionGetDefault with default headers values
func NewPoliciesAndRulesToBeAppliedCollectionGetDefault(code int) *PoliciesAndRulesToBeAppliedCollectionGetDefault {
	return &PoliciesAndRulesToBeAppliedCollectionGetDefault{
		_statusCode: code,
	}
}

/*
PoliciesAndRulesToBeAppliedCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type PoliciesAndRulesToBeAppliedCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this policies and rules to be applied collection get default response has a 2xx status code
func (o *PoliciesAndRulesToBeAppliedCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this policies and rules to be applied collection get default response has a 3xx status code
func (o *PoliciesAndRulesToBeAppliedCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this policies and rules to be applied collection get default response has a 4xx status code
func (o *PoliciesAndRulesToBeAppliedCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this policies and rules to be applied collection get default response has a 5xx status code
func (o *PoliciesAndRulesToBeAppliedCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this policies and rules to be applied collection get default response a status code equal to that given
func (o *PoliciesAndRulesToBeAppliedCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the policies and rules to be applied collection get default response
func (o *PoliciesAndRulesToBeAppliedCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *PoliciesAndRulesToBeAppliedCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies][%d] policies_and_rules_to_be_applied_collection_get default %s", o._statusCode, payload)
}

func (o *PoliciesAndRulesToBeAppliedCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/group-policies][%d] policies_and_rules_to_be_applied_collection_get default %s", o._statusCode, payload)
}

func (o *PoliciesAndRulesToBeAppliedCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PoliciesAndRulesToBeAppliedCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
