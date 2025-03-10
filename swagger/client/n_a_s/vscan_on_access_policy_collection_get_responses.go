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

// VscanOnAccessPolicyCollectionGetReader is a Reader for the VscanOnAccessPolicyCollectionGet structure.
type VscanOnAccessPolicyCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanOnAccessPolicyCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanOnAccessPolicyCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanOnAccessPolicyCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanOnAccessPolicyCollectionGetOK creates a VscanOnAccessPolicyCollectionGetOK with default headers values
func NewVscanOnAccessPolicyCollectionGetOK() *VscanOnAccessPolicyCollectionGetOK {
	return &VscanOnAccessPolicyCollectionGetOK{}
}

/*
VscanOnAccessPolicyCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type VscanOnAccessPolicyCollectionGetOK struct {
	Payload *models.VscanOnAccessResponse
}

// IsSuccess returns true when this vscan on access policy collection get o k response has a 2xx status code
func (o *VscanOnAccessPolicyCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vscan on access policy collection get o k response has a 3xx status code
func (o *VscanOnAccessPolicyCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vscan on access policy collection get o k response has a 4xx status code
func (o *VscanOnAccessPolicyCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vscan on access policy collection get o k response has a 5xx status code
func (o *VscanOnAccessPolicyCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vscan on access policy collection get o k response a status code equal to that given
func (o *VscanOnAccessPolicyCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vscan on access policy collection get o k response
func (o *VscanOnAccessPolicyCollectionGetOK) Code() int {
	return 200
}

func (o *VscanOnAccessPolicyCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/on-access-policies][%d] vscanOnAccessPolicyCollectionGetOK %s", 200, payload)
}

func (o *VscanOnAccessPolicyCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/on-access-policies][%d] vscanOnAccessPolicyCollectionGetOK %s", 200, payload)
}

func (o *VscanOnAccessPolicyCollectionGetOK) GetPayload() *models.VscanOnAccessResponse {
	return o.Payload
}

func (o *VscanOnAccessPolicyCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VscanOnAccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVscanOnAccessPolicyCollectionGetDefault creates a VscanOnAccessPolicyCollectionGetDefault with default headers values
func NewVscanOnAccessPolicyCollectionGetDefault(code int) *VscanOnAccessPolicyCollectionGetDefault {
	return &VscanOnAccessPolicyCollectionGetDefault{
		_statusCode: code,
	}
}

/*
VscanOnAccessPolicyCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type VscanOnAccessPolicyCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this vscan on access policy collection get default response has a 2xx status code
func (o *VscanOnAccessPolicyCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vscan on access policy collection get default response has a 3xx status code
func (o *VscanOnAccessPolicyCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vscan on access policy collection get default response has a 4xx status code
func (o *VscanOnAccessPolicyCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vscan on access policy collection get default response has a 5xx status code
func (o *VscanOnAccessPolicyCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vscan on access policy collection get default response a status code equal to that given
func (o *VscanOnAccessPolicyCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vscan on access policy collection get default response
func (o *VscanOnAccessPolicyCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *VscanOnAccessPolicyCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/on-access-policies][%d] vscan_on_access_policy_collection_get default %s", o._statusCode, payload)
}

func (o *VscanOnAccessPolicyCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/on-access-policies][%d] vscan_on_access_policy_collection_get default %s", o._statusCode, payload)
}

func (o *VscanOnAccessPolicyCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanOnAccessPolicyCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
