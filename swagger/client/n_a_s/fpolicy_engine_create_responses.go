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

// FpolicyEngineCreateReader is a Reader for the FpolicyEngineCreate structure.
type FpolicyEngineCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyEngineCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewFpolicyEngineCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyEngineCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyEngineCreateCreated creates a FpolicyEngineCreateCreated with default headers values
func NewFpolicyEngineCreateCreated() *FpolicyEngineCreateCreated {
	return &FpolicyEngineCreateCreated{}
}

/*
FpolicyEngineCreateCreated describes a response with status code 201, with default header values.

Created
*/
type FpolicyEngineCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.FpolicyEngineResponse
}

// IsSuccess returns true when this fpolicy engine create created response has a 2xx status code
func (o *FpolicyEngineCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fpolicy engine create created response has a 3xx status code
func (o *FpolicyEngineCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fpolicy engine create created response has a 4xx status code
func (o *FpolicyEngineCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this fpolicy engine create created response has a 5xx status code
func (o *FpolicyEngineCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this fpolicy engine create created response a status code equal to that given
func (o *FpolicyEngineCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the fpolicy engine create created response
func (o *FpolicyEngineCreateCreated) Code() int {
	return 201
}

func (o *FpolicyEngineCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/fpolicy/{svm.uuid}/engines][%d] fpolicyEngineCreateCreated %s", 201, payload)
}

func (o *FpolicyEngineCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/fpolicy/{svm.uuid}/engines][%d] fpolicyEngineCreateCreated %s", 201, payload)
}

func (o *FpolicyEngineCreateCreated) GetPayload() *models.FpolicyEngineResponse {
	return o.Payload
}

func (o *FpolicyEngineCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.FpolicyEngineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFpolicyEngineCreateDefault creates a FpolicyEngineCreateDefault with default headers values
func NewFpolicyEngineCreateDefault(code int) *FpolicyEngineCreateDefault {
	return &FpolicyEngineCreateDefault{
		_statusCode: code,
	}
}

/*
	FpolicyEngineCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 9764885    | The primary secondary server has a redundant IP address |
| 9764953    | The name of the FPolicy engine is "native" which is reserved by the system |
| 9765011    | The resiliency feature is not supported with mandatory screening |
| 9765012    | The specified resiliency directory path does not exist |
*/
type FpolicyEngineCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fpolicy engine create default response has a 2xx status code
func (o *FpolicyEngineCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fpolicy engine create default response has a 3xx status code
func (o *FpolicyEngineCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fpolicy engine create default response has a 4xx status code
func (o *FpolicyEngineCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fpolicy engine create default response has a 5xx status code
func (o *FpolicyEngineCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fpolicy engine create default response a status code equal to that given
func (o *FpolicyEngineCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fpolicy engine create default response
func (o *FpolicyEngineCreateDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyEngineCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/fpolicy/{svm.uuid}/engines][%d] fpolicy_engine_create default %s", o._statusCode, payload)
}

func (o *FpolicyEngineCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/fpolicy/{svm.uuid}/engines][%d] fpolicy_engine_create default %s", o._statusCode, payload)
}

func (o *FpolicyEngineCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyEngineCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
