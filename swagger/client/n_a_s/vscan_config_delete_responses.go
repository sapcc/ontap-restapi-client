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

// VscanConfigDeleteReader is a Reader for the VscanConfigDelete structure.
type VscanConfigDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanConfigDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanConfigDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanConfigDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanConfigDeleteOK creates a VscanConfigDeleteOK with default headers values
func NewVscanConfigDeleteOK() *VscanConfigDeleteOK {
	return &VscanConfigDeleteOK{}
}

/*
VscanConfigDeleteOK describes a response with status code 200, with default header values.

OK
*/
type VscanConfigDeleteOK struct {
}

// IsSuccess returns true when this vscan config delete o k response has a 2xx status code
func (o *VscanConfigDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vscan config delete o k response has a 3xx status code
func (o *VscanConfigDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vscan config delete o k response has a 4xx status code
func (o *VscanConfigDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vscan config delete o k response has a 5xx status code
func (o *VscanConfigDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vscan config delete o k response a status code equal to that given
func (o *VscanConfigDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vscan config delete o k response
func (o *VscanConfigDeleteOK) Code() int {
	return 200
}

func (o *VscanConfigDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/vscan/{svm.uuid}][%d] vscanConfigDeleteOK", 200)
}

func (o *VscanConfigDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/vscan/{svm.uuid}][%d] vscanConfigDeleteOK", 200)
}

func (o *VscanConfigDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVscanConfigDeleteDefault creates a VscanConfigDeleteDefault with default headers values
func NewVscanConfigDeleteDefault(code int) *VscanConfigDeleteDefault {
	return &VscanConfigDeleteDefault{
		_statusCode: code,
	}
}

/*
	VscanConfigDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 10027259   | A scanner-pool, an On-Access policy, or an On-Demand policy might fail to get deleted due to either a systematic error or some hardware failure. The error code returned details the failure along with the reason for the failure. For example, \"Failed to delete On-Access policy \"sp1\". Reason: \"Failed to delete policy. Reason: policy must be disabled before being deleted.\". Retry the operation.\"
*/
type VscanConfigDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this vscan config delete default response has a 2xx status code
func (o *VscanConfigDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vscan config delete default response has a 3xx status code
func (o *VscanConfigDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vscan config delete default response has a 4xx status code
func (o *VscanConfigDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vscan config delete default response has a 5xx status code
func (o *VscanConfigDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vscan config delete default response a status code equal to that given
func (o *VscanConfigDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vscan config delete default response
func (o *VscanConfigDeleteDefault) Code() int {
	return o._statusCode
}

func (o *VscanConfigDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/vscan/{svm.uuid}][%d] vscan_config_delete default %s", o._statusCode, payload)
}

func (o *VscanConfigDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/vscan/{svm.uuid}][%d] vscan_config_delete default %s", o._statusCode, payload)
}

func (o *VscanConfigDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanConfigDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
