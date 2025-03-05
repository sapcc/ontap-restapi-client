// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// FcpServiceCollectionGetReader is a Reader for the FcpServiceCollectionGet structure.
type FcpServiceCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FcpServiceCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFcpServiceCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFcpServiceCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFcpServiceCollectionGetOK creates a FcpServiceCollectionGetOK with default headers values
func NewFcpServiceCollectionGetOK() *FcpServiceCollectionGetOK {
	return &FcpServiceCollectionGetOK{}
}

/*
FcpServiceCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type FcpServiceCollectionGetOK struct {
	Payload *models.FcpServiceResponse
}

// IsSuccess returns true when this fcp service collection get o k response has a 2xx status code
func (o *FcpServiceCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fcp service collection get o k response has a 3xx status code
func (o *FcpServiceCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fcp service collection get o k response has a 4xx status code
func (o *FcpServiceCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fcp service collection get o k response has a 5xx status code
func (o *FcpServiceCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fcp service collection get o k response a status code equal to that given
func (o *FcpServiceCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fcp service collection get o k response
func (o *FcpServiceCollectionGetOK) Code() int {
	return 200
}

func (o *FcpServiceCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/fcp/services][%d] fcpServiceCollectionGetOK %s", 200, payload)
}

func (o *FcpServiceCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/fcp/services][%d] fcpServiceCollectionGetOK %s", 200, payload)
}

func (o *FcpServiceCollectionGetOK) GetPayload() *models.FcpServiceResponse {
	return o.Payload
}

func (o *FcpServiceCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FcpServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFcpServiceCollectionGetDefault creates a FcpServiceCollectionGetDefault with default headers values
func NewFcpServiceCollectionGetDefault(code int) *FcpServiceCollectionGetDefault {
	return &FcpServiceCollectionGetDefault{
		_statusCode: code,
	}
}

/*
FcpServiceCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type FcpServiceCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fcp service collection get default response has a 2xx status code
func (o *FcpServiceCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fcp service collection get default response has a 3xx status code
func (o *FcpServiceCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fcp service collection get default response has a 4xx status code
func (o *FcpServiceCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fcp service collection get default response has a 5xx status code
func (o *FcpServiceCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fcp service collection get default response a status code equal to that given
func (o *FcpServiceCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fcp service collection get default response
func (o *FcpServiceCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *FcpServiceCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/fcp/services][%d] fcp_service_collection_get default %s", o._statusCode, payload)
}

func (o *FcpServiceCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/fcp/services][%d] fcp_service_collection_get default %s", o._statusCode, payload)
}

func (o *FcpServiceCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FcpServiceCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
