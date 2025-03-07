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

// VscanEventCollectionGetReader is a Reader for the VscanEventCollectionGet structure.
type VscanEventCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanEventCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanEventCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanEventCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanEventCollectionGetOK creates a VscanEventCollectionGetOK with default headers values
func NewVscanEventCollectionGetOK() *VscanEventCollectionGetOK {
	return &VscanEventCollectionGetOK{}
}

/*
VscanEventCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type VscanEventCollectionGetOK struct {
	Payload *models.VscanEventResponse
}

// IsSuccess returns true when this vscan event collection get o k response has a 2xx status code
func (o *VscanEventCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vscan event collection get o k response has a 3xx status code
func (o *VscanEventCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vscan event collection get o k response has a 4xx status code
func (o *VscanEventCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vscan event collection get o k response has a 5xx status code
func (o *VscanEventCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vscan event collection get o k response a status code equal to that given
func (o *VscanEventCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vscan event collection get o k response
func (o *VscanEventCollectionGetOK) Code() int {
	return 200
}

func (o *VscanEventCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/events][%d] vscanEventCollectionGetOK %s", 200, payload)
}

func (o *VscanEventCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/events][%d] vscanEventCollectionGetOK %s", 200, payload)
}

func (o *VscanEventCollectionGetOK) GetPayload() *models.VscanEventResponse {
	return o.Payload
}

func (o *VscanEventCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VscanEventResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVscanEventCollectionGetDefault creates a VscanEventCollectionGetDefault with default headers values
func NewVscanEventCollectionGetDefault(code int) *VscanEventCollectionGetDefault {
	return &VscanEventCollectionGetDefault{
		_statusCode: code,
	}
}

/*
VscanEventCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type VscanEventCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this vscan event collection get default response has a 2xx status code
func (o *VscanEventCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vscan event collection get default response has a 3xx status code
func (o *VscanEventCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vscan event collection get default response has a 4xx status code
func (o *VscanEventCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vscan event collection get default response has a 5xx status code
func (o *VscanEventCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vscan event collection get default response a status code equal to that given
func (o *VscanEventCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vscan event collection get default response
func (o *VscanEventCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *VscanEventCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/events][%d] vscan_event_collection_get default %s", o._statusCode, payload)
}

func (o *VscanEventCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/events][%d] vscan_event_collection_get default %s", o._statusCode, payload)
}

func (o *VscanEventCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanEventCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
