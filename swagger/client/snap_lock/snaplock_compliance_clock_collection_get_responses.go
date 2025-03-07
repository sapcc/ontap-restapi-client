// Code generated by go-swagger; DO NOT EDIT.

package snap_lock

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

// SnaplockComplianceClockCollectionGetReader is a Reader for the SnaplockComplianceClockCollectionGet structure.
type SnaplockComplianceClockCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockComplianceClockCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockComplianceClockCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockComplianceClockCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockComplianceClockCollectionGetOK creates a SnaplockComplianceClockCollectionGetOK with default headers values
func NewSnaplockComplianceClockCollectionGetOK() *SnaplockComplianceClockCollectionGetOK {
	return &SnaplockComplianceClockCollectionGetOK{}
}

/*
SnaplockComplianceClockCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockComplianceClockCollectionGetOK struct {
	Payload *models.SnaplockComplianceClockResponse
}

// IsSuccess returns true when this snaplock compliance clock collection get o k response has a 2xx status code
func (o *SnaplockComplianceClockCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock compliance clock collection get o k response has a 3xx status code
func (o *SnaplockComplianceClockCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock compliance clock collection get o k response has a 4xx status code
func (o *SnaplockComplianceClockCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock compliance clock collection get o k response has a 5xx status code
func (o *SnaplockComplianceClockCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock compliance clock collection get o k response a status code equal to that given
func (o *SnaplockComplianceClockCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snaplock compliance clock collection get o k response
func (o *SnaplockComplianceClockCollectionGetOK) Code() int {
	return 200
}

func (o *SnaplockComplianceClockCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/compliance-clocks][%d] snaplockComplianceClockCollectionGetOK %s", 200, payload)
}

func (o *SnaplockComplianceClockCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/compliance-clocks][%d] snaplockComplianceClockCollectionGetOK %s", 200, payload)
}

func (o *SnaplockComplianceClockCollectionGetOK) GetPayload() *models.SnaplockComplianceClockResponse {
	return o.Payload
}

func (o *SnaplockComplianceClockCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockComplianceClockResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockComplianceClockCollectionGetDefault creates a SnaplockComplianceClockCollectionGetDefault with default headers values
func NewSnaplockComplianceClockCollectionGetDefault(code int) *SnaplockComplianceClockCollectionGetDefault {
	return &SnaplockComplianceClockCollectionGetDefault{
		_statusCode: code,
	}
}

/*
SnaplockComplianceClockCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SnaplockComplianceClockCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snaplock compliance clock collection get default response has a 2xx status code
func (o *SnaplockComplianceClockCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock compliance clock collection get default response has a 3xx status code
func (o *SnaplockComplianceClockCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock compliance clock collection get default response has a 4xx status code
func (o *SnaplockComplianceClockCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock compliance clock collection get default response has a 5xx status code
func (o *SnaplockComplianceClockCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock compliance clock collection get default response a status code equal to that given
func (o *SnaplockComplianceClockCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snaplock compliance clock collection get default response
func (o *SnaplockComplianceClockCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockComplianceClockCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/compliance-clocks][%d] snaplock_compliance_clock_collection_get default %s", o._statusCode, payload)
}

func (o *SnaplockComplianceClockCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/compliance-clocks][%d] snaplock_compliance_clock_collection_get default %s", o._statusCode, payload)
}

func (o *SnaplockComplianceClockCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockComplianceClockCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
