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

// VscanOnDemandDeleteReader is a Reader for the VscanOnDemandDelete structure.
type VscanOnDemandDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanOnDemandDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanOnDemandDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanOnDemandDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanOnDemandDeleteOK creates a VscanOnDemandDeleteOK with default headers values
func NewVscanOnDemandDeleteOK() *VscanOnDemandDeleteOK {
	return &VscanOnDemandDeleteOK{}
}

/*
VscanOnDemandDeleteOK describes a response with status code 200, with default header values.

OK
*/
type VscanOnDemandDeleteOK struct {
}

// IsSuccess returns true when this vscan on demand delete o k response has a 2xx status code
func (o *VscanOnDemandDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vscan on demand delete o k response has a 3xx status code
func (o *VscanOnDemandDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vscan on demand delete o k response has a 4xx status code
func (o *VscanOnDemandDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vscan on demand delete o k response has a 5xx status code
func (o *VscanOnDemandDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vscan on demand delete o k response a status code equal to that given
func (o *VscanOnDemandDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vscan on demand delete o k response
func (o *VscanOnDemandDeleteOK) Code() int {
	return 200
}

func (o *VscanOnDemandDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/vscan/{svm.uuid}/on-demand-policies/{name}][%d] vscanOnDemandDeleteOK", 200)
}

func (o *VscanOnDemandDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/vscan/{svm.uuid}/on-demand-policies/{name}][%d] vscanOnDemandDeleteOK", 200)
}

func (o *VscanOnDemandDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVscanOnDemandDeleteDefault creates a VscanOnDemandDeleteDefault with default headers values
func NewVscanOnDemandDeleteDefault(code int) *VscanOnDemandDeleteDefault {
	return &VscanOnDemandDeleteDefault{
		_statusCode: code,
	}
}

/*
VscanOnDemandDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type VscanOnDemandDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this vscan on demand delete default response has a 2xx status code
func (o *VscanOnDemandDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vscan on demand delete default response has a 3xx status code
func (o *VscanOnDemandDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vscan on demand delete default response has a 4xx status code
func (o *VscanOnDemandDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vscan on demand delete default response has a 5xx status code
func (o *VscanOnDemandDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vscan on demand delete default response a status code equal to that given
func (o *VscanOnDemandDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vscan on demand delete default response
func (o *VscanOnDemandDeleteDefault) Code() int {
	return o._statusCode
}

func (o *VscanOnDemandDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/vscan/{svm.uuid}/on-demand-policies/{name}][%d] vscan_on_demand_delete default %s", o._statusCode, payload)
}

func (o *VscanOnDemandDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/vscan/{svm.uuid}/on-demand-policies/{name}][%d] vscan_on_demand_delete default %s", o._statusCode, payload)
}

func (o *VscanOnDemandDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanOnDemandDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
