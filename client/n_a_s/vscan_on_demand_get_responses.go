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

// VscanOnDemandGetReader is a Reader for the VscanOnDemandGet structure.
type VscanOnDemandGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanOnDemandGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanOnDemandGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanOnDemandGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanOnDemandGetOK creates a VscanOnDemandGetOK with default headers values
func NewVscanOnDemandGetOK() *VscanOnDemandGetOK {
	return &VscanOnDemandGetOK{}
}

/*
VscanOnDemandGetOK describes a response with status code 200, with default header values.

OK
*/
type VscanOnDemandGetOK struct {
	Payload *models.VscanOnDemand
}

// IsSuccess returns true when this vscan on demand get o k response has a 2xx status code
func (o *VscanOnDemandGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vscan on demand get o k response has a 3xx status code
func (o *VscanOnDemandGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vscan on demand get o k response has a 4xx status code
func (o *VscanOnDemandGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vscan on demand get o k response has a 5xx status code
func (o *VscanOnDemandGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vscan on demand get o k response a status code equal to that given
func (o *VscanOnDemandGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vscan on demand get o k response
func (o *VscanOnDemandGetOK) Code() int {
	return 200
}

func (o *VscanOnDemandGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/on-demand-policies/{name}][%d] vscanOnDemandGetOK %s", 200, payload)
}

func (o *VscanOnDemandGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/on-demand-policies/{name}][%d] vscanOnDemandGetOK %s", 200, payload)
}

func (o *VscanOnDemandGetOK) GetPayload() *models.VscanOnDemand {
	return o.Payload
}

func (o *VscanOnDemandGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VscanOnDemand)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVscanOnDemandGetDefault creates a VscanOnDemandGetDefault with default headers values
func NewVscanOnDemandGetDefault(code int) *VscanOnDemandGetDefault {
	return &VscanOnDemandGetDefault{
		_statusCode: code,
	}
}

/*
VscanOnDemandGetDefault describes a response with status code -1, with default header values.

Error
*/
type VscanOnDemandGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this vscan on demand get default response has a 2xx status code
func (o *VscanOnDemandGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vscan on demand get default response has a 3xx status code
func (o *VscanOnDemandGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vscan on demand get default response has a 4xx status code
func (o *VscanOnDemandGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vscan on demand get default response has a 5xx status code
func (o *VscanOnDemandGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vscan on demand get default response a status code equal to that given
func (o *VscanOnDemandGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vscan on demand get default response
func (o *VscanOnDemandGetDefault) Code() int {
	return o._statusCode
}

func (o *VscanOnDemandGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/on-demand-policies/{name}][%d] vscan_on_demand_get default %s", o._statusCode, payload)
}

func (o *VscanOnDemandGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/on-demand-policies/{name}][%d] vscan_on_demand_get default %s", o._statusCode, payload)
}

func (o *VscanOnDemandGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanOnDemandGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
