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

// VscanCollectionGetReader is a Reader for the VscanCollectionGet structure.
type VscanCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanCollectionGetOK creates a VscanCollectionGetOK with default headers values
func NewVscanCollectionGetOK() *VscanCollectionGetOK {
	return &VscanCollectionGetOK{}
}

/*
VscanCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type VscanCollectionGetOK struct {
	Payload *models.VscanResponse
}

// IsSuccess returns true when this vscan collection get o k response has a 2xx status code
func (o *VscanCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vscan collection get o k response has a 3xx status code
func (o *VscanCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vscan collection get o k response has a 4xx status code
func (o *VscanCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vscan collection get o k response has a 5xx status code
func (o *VscanCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vscan collection get o k response a status code equal to that given
func (o *VscanCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vscan collection get o k response
func (o *VscanCollectionGetOK) Code() int {
	return 200
}

func (o *VscanCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan][%d] vscanCollectionGetOK %s", 200, payload)
}

func (o *VscanCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan][%d] vscanCollectionGetOK %s", 200, payload)
}

func (o *VscanCollectionGetOK) GetPayload() *models.VscanResponse {
	return o.Payload
}

func (o *VscanCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VscanResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVscanCollectionGetDefault creates a VscanCollectionGetDefault with default headers values
func NewVscanCollectionGetDefault(code int) *VscanCollectionGetDefault {
	return &VscanCollectionGetDefault{
		_statusCode: code,
	}
}

/*
VscanCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type VscanCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this vscan collection get default response has a 2xx status code
func (o *VscanCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vscan collection get default response has a 3xx status code
func (o *VscanCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vscan collection get default response has a 4xx status code
func (o *VscanCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vscan collection get default response has a 5xx status code
func (o *VscanCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vscan collection get default response a status code equal to that given
func (o *VscanCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vscan collection get default response
func (o *VscanCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *VscanCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan][%d] vscan_collection_get default %s", o._statusCode, payload)
}

func (o *VscanCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan][%d] vscan_collection_get default %s", o._statusCode, payload)
}

func (o *VscanCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
