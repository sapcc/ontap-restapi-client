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

// VscanScannerCollectionGetReader is a Reader for the VscanScannerCollectionGet structure.
type VscanScannerCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanScannerCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanScannerCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanScannerCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanScannerCollectionGetOK creates a VscanScannerCollectionGetOK with default headers values
func NewVscanScannerCollectionGetOK() *VscanScannerCollectionGetOK {
	return &VscanScannerCollectionGetOK{}
}

/*
VscanScannerCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type VscanScannerCollectionGetOK struct {
	Payload *models.VscanScannerPoolResponse
}

// IsSuccess returns true when this vscan scanner collection get o k response has a 2xx status code
func (o *VscanScannerCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vscan scanner collection get o k response has a 3xx status code
func (o *VscanScannerCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vscan scanner collection get o k response has a 4xx status code
func (o *VscanScannerCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vscan scanner collection get o k response has a 5xx status code
func (o *VscanScannerCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vscan scanner collection get o k response a status code equal to that given
func (o *VscanScannerCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vscan scanner collection get o k response
func (o *VscanScannerCollectionGetOK) Code() int {
	return 200
}

func (o *VscanScannerCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/scanner-pools][%d] vscanScannerCollectionGetOK %s", 200, payload)
}

func (o *VscanScannerCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/scanner-pools][%d] vscanScannerCollectionGetOK %s", 200, payload)
}

func (o *VscanScannerCollectionGetOK) GetPayload() *models.VscanScannerPoolResponse {
	return o.Payload
}

func (o *VscanScannerCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VscanScannerPoolResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVscanScannerCollectionGetDefault creates a VscanScannerCollectionGetDefault with default headers values
func NewVscanScannerCollectionGetDefault(code int) *VscanScannerCollectionGetDefault {
	return &VscanScannerCollectionGetDefault{
		_statusCode: code,
	}
}

/*
VscanScannerCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type VscanScannerCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this vscan scanner collection get default response has a 2xx status code
func (o *VscanScannerCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vscan scanner collection get default response has a 3xx status code
func (o *VscanScannerCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vscan scanner collection get default response has a 4xx status code
func (o *VscanScannerCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vscan scanner collection get default response has a 5xx status code
func (o *VscanScannerCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vscan scanner collection get default response a status code equal to that given
func (o *VscanScannerCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vscan scanner collection get default response
func (o *VscanScannerCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *VscanScannerCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/scanner-pools][%d] vscan_scanner_collection_get default %s", o._statusCode, payload)
}

func (o *VscanScannerCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/scanner-pools][%d] vscan_scanner_collection_get default %s", o._statusCode, payload)
}

func (o *VscanScannerCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanScannerCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
