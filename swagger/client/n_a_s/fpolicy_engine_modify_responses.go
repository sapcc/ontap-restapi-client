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

// FpolicyEngineModifyReader is a Reader for the FpolicyEngineModify structure.
type FpolicyEngineModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyEngineModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyEngineModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyEngineModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyEngineModifyOK creates a FpolicyEngineModifyOK with default headers values
func NewFpolicyEngineModifyOK() *FpolicyEngineModifyOK {
	return &FpolicyEngineModifyOK{}
}

/*
FpolicyEngineModifyOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyEngineModifyOK struct {
}

// IsSuccess returns true when this fpolicy engine modify o k response has a 2xx status code
func (o *FpolicyEngineModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fpolicy engine modify o k response has a 3xx status code
func (o *FpolicyEngineModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fpolicy engine modify o k response has a 4xx status code
func (o *FpolicyEngineModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fpolicy engine modify o k response has a 5xx status code
func (o *FpolicyEngineModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fpolicy engine modify o k response a status code equal to that given
func (o *FpolicyEngineModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fpolicy engine modify o k response
func (o *FpolicyEngineModifyOK) Code() int {
	return 200
}

func (o *FpolicyEngineModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/engines/{name}][%d] fpolicyEngineModifyOK", 200)
}

func (o *FpolicyEngineModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/engines/{name}][%d] fpolicyEngineModifyOK", 200)
}

func (o *FpolicyEngineModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFpolicyEngineModifyDefault creates a FpolicyEngineModifyDefault with default headers values
func NewFpolicyEngineModifyDefault(code int) *FpolicyEngineModifyDefault {
	return &FpolicyEngineModifyDefault{
		_statusCode: code,
	}
}

/*
	FpolicyEngineModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 9764922    | The primary and secondary server has a redundant IP address |
| 9764942    | At least one FPolicy policy is using the FPolicy engine |
| 9764886    | FPolicy engine is a cluster-level FPolicy engine |
| 9765011    | The resiliency feature is not supported with mandatory screening |
| 9765012    | The specified resiliency directory path does not exist|
| 9765042    | The specified send buffer size exceeds the maximum limit |
| 9765043    | The specified receive buffer size exceeds the maximum limit |
| 9765063    | Policy with Persistent Store feature does not support a "synchronous" |
*/
type FpolicyEngineModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fpolicy engine modify default response has a 2xx status code
func (o *FpolicyEngineModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fpolicy engine modify default response has a 3xx status code
func (o *FpolicyEngineModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fpolicy engine modify default response has a 4xx status code
func (o *FpolicyEngineModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fpolicy engine modify default response has a 5xx status code
func (o *FpolicyEngineModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fpolicy engine modify default response a status code equal to that given
func (o *FpolicyEngineModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fpolicy engine modify default response
func (o *FpolicyEngineModifyDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyEngineModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/engines/{name}][%d] fpolicy_engine_modify default %s", o._statusCode, payload)
}

func (o *FpolicyEngineModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/engines/{name}][%d] fpolicy_engine_modify default %s", o._statusCode, payload)
}

func (o *FpolicyEngineModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyEngineModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
